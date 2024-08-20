package main

import (
	"fmt"
	"log"
	"strings"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	xmpp "github.com/adrianfulla/Proyecto1-Redes/server/xmpp"
	xmppfunctions "github.com/adrianfulla/Proyecto1-Redes/server/xmpp-functions"
)

func main() {
    ShowLoginWindow()
}

func ShowLoginWindow() {
    myApp := app.New()
    myWindow := myApp.NewWindow("XMPP Chat Client")

    serverEntry := widget.NewEntry()
    serverEntry.SetPlaceHolder("Server (e.g., alumchat.lol:5222)")

    usernameEntry := widget.NewEntry()
    usernameEntry.SetPlaceHolder("Username")

    passwordEntry := widget.NewPasswordEntry()
    passwordEntry.SetPlaceHolder("Password")

    loginButton := widget.NewButton("Login", func() {
        server := serverEntry.Text
        username := usernameEntry.Text
        password := passwordEntry.Text

        hostPort := strings.Split(server, ":")
        if len(hostPort) != 2 {
            dialog.ShowError(fmt.Errorf("invalid server format"), myWindow)
            return
        }

        handler, err := xmppfunctions.Login(hostPort[0], hostPort[1], username, password)
        if err != nil {
            log.Printf("Login failed: %v", err)
            dialog.ShowError(err, myWindow)
            return
        }

        // If login is successful, move to the chat window
        ShowContactsWindow(myApp, handler)
        myWindow.Close()
    })

    createAccountButton := widget.NewButton("Create Account", func() {
        ShowCreateAccountDialog(myApp, myWindow)
    })

    myWindow.SetContent(container.NewVBox(
        widget.NewLabel("Login to XMPP Server"),
        serverEntry,
        usernameEntry,
        passwordEntry,
        loginButton,
        createAccountButton,
    ))

    myWindow.ShowAndRun()
}

// ShowChatWindow opens a new chat window with the selected contact.
func ShowChatWindow(app fyne.App, handler *xmpp.XMPPHandler, recipient string) *xmpp.ChatWindow {
    chatWindow := app.NewWindow("Chat with " + recipient)

    messageEntry := widget.NewEntry()
    messageEntry.SetPlaceHolder("Type your message...")

    chatContent := container.NewVBox()

	if queuedMessages, ok := handler.MessageQueue[recipient]; ok {
        for _, msg := range queuedMessages {
            chatContent.Add(widget.NewLabel(fmt.Sprintf("%s: %s", strings.Split(msg.From,"/")[0], msg.Body)))
        }
        delete(handler.MessageQueue, recipient) // Clear the queue after displaying
    }

    sendMessageButton := widget.NewButton("Send", func() {
        message := messageEntry.Text
        if message != "" {
            err := xmppfunctions.SendMessage(handler, recipient, message)
            if err != nil {
                log.Printf("Failed to send message: %v", err)
            } else {
                chatContent.Add(widget.NewLabel("Me: " + message))
                messageEntry.SetText("")
            }
        }
    })

    // Create a grid layout with two columns
    messageRow := container.New(layout.NewGridLayoutWithColumns(2), messageEntry, sendMessageButton)

	chatWindow.Resize(fyne.NewSize(400, 500))


    chatWindow.SetContent(container.NewBorder(
        chatContent,
        messageRow,
        nil, nil,
    ))

	chatWindow.Canvas().SetOnTypedKey(func(key *fyne.KeyEvent) {
        if key.Name == fyne.KeyReturn {
            sendMessageButton.OnTapped()
        }
    })

	chatWindow.Canvas().SetOnTypedRune(func(r rune) {
        messageEntry.TypedRune(r)
    })

	chatWindow.Canvas().SetOnTypedKey(func(*fyne.KeyEvent) {
        width := chatWindow.Canvas().Size().Width
        messageEntry.Resize(fyne.NewSize(width*0.8, messageEntry.MinSize().Height))
        sendMessageButton.Resize(fyne.NewSize(width*0.2, sendMessageButton.MinSize().Height))
    })

    // Handle incoming messages in a separate goroutine
    go func() {
        for {
            err := handler.HandleIncomingStanzas()
            if err != nil {
                log.Printf("Error handling stanzas: %v", err)
                continue
            }

            // After `HandleIncomingStanzas`, you might need to trigger the UI update.
            chatContent.Add(widget.NewLabel("Received a new message...")) // Example, replace with actual data
            chatWindow.Content().Refresh()
        }
    }()

    chatWindow.Resize(fyne.NewSize(400, 500))
    chatWindow.Show()

	return &xmpp.ChatWindow{
		Window:      chatWindow,
		ChatContent: chatContent,
		Handler:     handler,
		Recipient:   recipient,
	}
}

// ShowContactsWindow displays the user's contact list.
func ShowContactsWindow(app fyne.App, handler *xmpp.XMPPHandler) {
    contactWindow := app.NewWindow("Contacts - " + handler.Username)

    var contacts []xmppfunctions.Contact

    // Function to refresh the contact list
    refreshContactList := func(contactList *widget.List) {
        var err error
        contacts, err = xmppfunctions.GetContacts(handler)
        if err != nil {
            log.Printf("Failed to get contacts: %v", err)
            dialog.ShowError(err, contactWindow)
            return
        }

        contactList.Length = func() int {
            return len(contacts)
        }
        contactList.UpdateItem = func(i widget.ListItemID, o fyne.CanvasObject) {
			jid := contacts[i].JID
			queuedMessages := len(handler.MessageQueue[jid])
            displayText := fmt.Sprintf("%s - %s",jid,contacts[i].Status)

            if queuedMessages > 0 {
                displayText = fmt.Sprintf("%s (%d) - %s", jid, queuedMessages, contacts[i].Status)
            }

            o.(*widget.Label).SetText(displayText)
			
        }

        contactList.Refresh()
    }

    contactList := widget.NewList(
        func() int {
            return len(contacts)
        },
        func() fyne.CanvasObject {
            return widget.NewLabel("")
        },
        func(i widget.ListItemID, o fyne.CanvasObject) {},
    )

    refreshContactList(contactList)

    contactList.OnSelected = func(id widget.ListItemID) {
        if id >= 0 && id < len(contacts) {
            selectedContact := contacts[id]
            chatWindow := ShowChatWindow(app, handler, selectedContact.JID)
            handler.ChatWindows[selectedContact.JID] = chatWindow
        } else {
            log.Printf("Invalid selection: %d", id)
        }
    }

    addContactButton := widget.NewButton("Add Contact", func() {
        jidEntry := widget.NewEntry()
        jidEntry.SetPlaceHolder("Enter contact JID (e.g., user@example.com)")

        dialog.ShowCustomConfirm("Add Contact", "Add", "Cancel", container.NewVBox(
            widget.NewLabel("Add a new contact"),
            jidEntry,
        ), func(ok bool) {
            if ok {
                newJID := jidEntry.Text
                if newJID != "" {
                    err := xmppfunctions.AddContact(handler, newJID)
                    if err != nil {
                        log.Printf("Failed to add contact: %v", err)
                        dialog.ShowError(err, contactWindow)
                    } else {
                        log.Printf("Contact added: %s", newJID)
                        refreshContactList(contactList)
                    }
                }
            }
        }, contactWindow)
    })

    settingsButton := widget.NewButton("Settings", func() {
        ShowUserSettingsWindow(app, handler)
    })

    contactWindow.SetContent(
        container.NewBorder(
            container.NewVBox(settingsButton, addContactButton, widget.NewLabel("Your Contacts")),
            nil, nil, nil,
            container.NewVScroll(contactList),
        ),
    )

    contactWindow.Resize(fyne.NewSize(300, 400))
    contactWindow.Show()
	
	go func() {
        for {
            for _, queuedMessages := range handler.MessageQueue {
                if len(queuedMessages) > 0 {
                    contactList.Refresh()
                }
            }
            time.Sleep(2 * time.Second)
        }
    }()

    go func() {
        for msg := range handler.MessageChan {
            handler.DispatchMessage(msg)
        }
    }()

    handler.ListenForIncomingStanzas()
}




func ShowUserSettingsWindow(app fyne.App, handler *xmpp.XMPPHandler) {
    settingsWindow := app.NewWindow("User Settings")

    logoutButton := widget.NewButton("Logout", func() {
        err := xmppfunctions.Logout(handler)
        if err != nil {
            log.Printf("Logout failed: %v", err)
            dialog.ShowError(err, settingsWindow)
        } else {
            log.Println("Logged out successfully")
			CloseAllWindows(app)
            settingsWindow.Close()
			app.Quit()
        }
    })

    deleteAccountButton := widget.NewButton("Delete Account", func() {
        confirmDialog := dialog.NewConfirm("Delete Account", "Are you sure you want to delete your account?", func(confirm bool) {
            if confirm {
                err := xmppfunctions.RemoveAccount(handler)
                if err != nil {
                    log.Printf("Account removal failed: %v", err)
                    dialog.ShowError(err, settingsWindow)
					app.Quit()
                } else {
                    log.Println("Account removed successfully")
                    settingsWindow.Close()
                    // Optionally close all windows or return to the login window
                    app.Quit()
                }
            }
        }, settingsWindow)
        confirmDialog.SetDismissText("Cancel")
        confirmDialog.Show()
    })

    settingsWindow.SetContent(container.NewVBox(
        widget.NewLabel("User Settings"),
        logoutButton,
        deleteAccountButton,
    ))

    settingsWindow.Resize(fyne.NewSize(300, 200))
    settingsWindow.Show()
}


func CloseAllWindows(app fyne.App) {
    for _, window := range app.Driver().AllWindows() {
        window.Close()
    }
}


func ShowCreateAccountDialog(app fyne.App, parent fyne.Window) {
    serverEntry := widget.NewEntry()
    serverEntry.SetPlaceHolder("Server (e.g., alumchat.lol:5222)")

    usernameEntry := widget.NewEntry()
    usernameEntry.SetPlaceHolder("Desired Username")

    passwordEntry := widget.NewPasswordEntry()
    passwordEntry.SetPlaceHolder("Desired Password")

    errorLabel := widget.NewLabel("")

    dialogWindow := app.NewWindow("Create Account")

    confirmButton := widget.NewButton("Create Account", func() {
        server := serverEntry.Text
        username := usernameEntry.Text
        password := passwordEntry.Text

        hostPort := strings.Split(server, ":")
        if len(hostPort) != 2 {
            errorLabel.SetText("Invalid server format")
            return
        }

        err := xmppfunctions.CreateUser(hostPort[0], hostPort[1], username, password)
        if err != nil {
            log.Printf("Account creation failed: %v", err)
            errorLabel.SetText(fmt.Sprintf("Error: %v", err))
            return
        }

        errorLabel.SetText("Account created successfully!")
        log.Println("Account created successfully")
        dialogWindow.Close() // Close the account creation window on success
    })

    content := container.NewVBox(
        widget.NewLabel("Create a New XMPP Account"),
        serverEntry,
        usernameEntry,
        passwordEntry,
        errorLabel,
        confirmButton,
    )

    dialogWindow.SetContent(content)
    dialogWindow.Resize(fyne.NewSize(300, 200))
    dialogWindow.Show()
}


