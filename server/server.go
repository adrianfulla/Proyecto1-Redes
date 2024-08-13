package main

import (
	"encoding/json"
	"net/http"
    "fmt"
    "log"

	"github.com/adrianfulla/Proyecto1-Redes/server/xmpp"

)


func main() {
	// http.HandleFunc("/", homePageHandler)

	// http.HandleFunc("/api/thumbnail", thumbnailHandler)

	// fs := http.FileServer(http.Dir("./frontend/dist"))

	// http.Handle("/", fs)

	// fmt.Println("Server listening on port 3000")
	// log.Panic(
	// 	http.ListenAndServe(":3000", nil),
	// )

	// MelliumTest()

	CreateUserTest()
	// LogInTest()

}

func LogInTest(){
	server := "alumchat.lol:5222"
    username := "aa-test4"
    password := "12345"

    handler, err := xmpp.NewXMPPHandler(server, username, password)
    if err != nil {
        log.Fatalf("Failed to initialize XMPP handler: %v", err)
    }

    handler.SendPresence("available", "Ready to chat!")
    handler.SendMessage("adrianfulla21592-test1@alumchat.lol", "Hello, how are you?")
    handler.HandleIncomingStanzas()

    handler.WaitForShutdown()
}

func CreateUserTest(){
	// server := "alumchat.lol:7070"
	server := "alumchat.lol:5222"

    // Establish a connection to the XMPP server
    conn, err := xmpp.NewXMPPConnection(server, false) 
    if err != nil {
        log.Fatalf("Failed to connect: %v", err)
    }
    defer conn.Close()

    // Start an XMPP stream without authentication (for registration purposes)
    if err := conn.StartStream(""); err != nil {
        log.Fatalf("Failed to start stream: %v", err)
    }

    // Attempt to create a new user account
    username := "aa-test11"
    password := "12345"
    err = xmpp.CreateUser(conn, username, password)
    if err != nil {
        log.Fatalf("Failed to create user: %v", err)
    } else {
        log.Println("User created successfully")
    }
}

func homePageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}

type thumbnailRequest struct {
	Url string `json:"url"`
}

func thumbnailHandler(w http.ResponseWriter, r *http.Request) {
	var decoded thumbnailRequest

	err := json.NewDecoder(r.Body).Decode(&decoded)
	if err != nil {
		fmt.Printf("Got a thumbnail")
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	fmt.Printf("Got the following url: %s\n", decoded.Url)
}
