module github.com/adrianfulla/Proyecto1-Redes

go 1.22.1

replace github.com/adrianfulla/Proyecto1-Redes/server/xmpp => ./server/xmpp

require github.com/adrianfulla/Proyecto1-Redes/server/xmpp-functions v0.0.0-00010101000000-000000000000

require github.com/adrianfulla/Proyecto1-Redes/server/xmpp v0.0.0-00010101000000-000000000000 // indirect

replace github.com/adrianfulla/Proyecto1-Redes/server/xmpp-functions => ./server/xmpp-functions
