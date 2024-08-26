# Project 1 Networks - Using an Existing Protocol
[![es](https://img.shields.io/badge/lang-es-yellow.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.md)
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.en.md)

## Project Objectives
- Implement a protocol based on standards.
- Understand the purpose of the XMPP protocol.
- Understand how XMPP protocol services work.
- Apply the knowledge acquired in Web and mobile programming.

## Achieved Objectives
- Developed a chat client based on XMPP entirely in the Go programming language.
- The client includes the following features:
    - Register a new account on the server.
    - Log in with an account.
    - Log out with an account.
    - Delete the account from the server.
    - Display contacts and their statuses.
    - Add contacts.
    - Display contact details.
    - 1-on-1 communication with a contact.
    - Define presence on the server.
- Additionally, an XMPP protocol was implemented from scratch to facilitate communication with the server.

## Installation Requirements
- GoLang -> https://go.dev/doc/install

## Installation
1. Clone the repository:
```bash
    git clone https://github.com/adrianfulla/Proyecto1-Redes.git
```
2. Access the repository directory:
```bash
    cd Proyecto1-Redes/
```

3. Run the following command:
```bash
    go mod tidy
```
This command will install all the necessary dependencies.

4. Run the `go run` command to compile and execute the client:
```bash
    go run ./server
```

## Notes
- The Fyne library and its documentation were used to create the client’s visualization. https://github.com/fyne-io/fyne
- ChatGPT, using the Go Golang model, was utilized both for creating the XMPP protocol and for developing the user interface. The conversation can be found here: https://chatgpt.com/share/2923f8f4-38c2-44c0-b8d2-bc6538e32ba8
- The book XMPP: The Definitive Guide by Peter Saint-Andre, Kevin Smith, and Remko Tronçon was used as a reference for implementing the XMPP protocol.