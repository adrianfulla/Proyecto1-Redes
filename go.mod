module github.com/adrianfulla/Proyecto1-Redes

go 1.22.1

replace github.com/adrianfulla/Proyecto1-Redes/server/xmpp => ./server/xmpp

require (
	github.com/adrianfulla/Proyecto1-Redes/server/xmpp v0.0.0-00010101000000-000000000000
	mellium.im/sasl v0.3.1
	mellium.im/xmlstream v0.15.4
	mellium.im/xmpp v0.21.4
)

require (
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/mod v0.7.0 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	golang.org/x/tools v0.5.0 // indirect
	mellium.im/reader v0.1.0 // indirect
)
