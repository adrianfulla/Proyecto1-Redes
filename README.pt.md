# Projeto 1 Redes - Uso de um Protocolo Existente
[![es](https://img.shields.io/badge/lang-es-yellow.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.md)
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.en.md)
[![fr](https://img.shields.io/badge/lang-fr-gr.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.fr.md)
[![pt](https://img.shields.io/badge/lang-pt-blue.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.pt.md)

## Objetivos do Projeto
- Implementar um protocolo baseado nos padrões.
- Compreender o propósito do protocolo XMPP.
- Compreender como funcionam os serviços do protocolo XMPP.
- Aplicar os conhecimentos adquiridos em programação Web e móvel.

## Objetivos Alcançados
-Desenvolvido um cliente de chat baseado em XMPP inteiramente na linguagem de programação Go.
- O cliente inclui as seguintes funcionalidades:
    - Registrar uma nova conta no servidor.
    - Fazer login com uma conta.
    - Fazer logout com uma conta.
    - Excluir a conta do servidor.
    - Exibir contatos e seus status.
    - Adicionar contatos.
    - Exibir detalhes do contato.
    - Comunicação 1 a 1 com um contato.
    - Definição de presença no servidor.
- Além disso, um protocolo XMPP foi implementado do zero para facilitar a comunicação com o servidor.

## Requisitos de Instalação
- GoLang -> https://go.dev/doc/install

## Instalação
1. Clonar o repositório:
```bash
    git clone https://github.com/adrianfulla/Proyecto1-Redes.git
```
2. Acessar o diretório do repositório:
```bash
    cd Proyecto1-Redes/
```

3. Executar o comando:
```bash
    go mod tidy
```
Este comando instalará todas as dependências necessárias.

4. Executar o comando go run para compilar e executar o cliente:
```bash
    go run ./server
```

## Notas
- A biblioteca Fyne e sua documentação foram utilizadas para criar a visualização do cliente. https://github.com/fyne-io/fyne
- ChatGPT, utilizando o modelo Go Golang, foi utilizado tanto para criar o protocolo XMPP quanto para desenvolver a interface do usuário. A conversa pode ser encontrada aqui: https://chatgpt.com/share/2923f8f4-38c2-44c0-b8d2-bc6538e32ba8
- O livro XMPP: The Definitive Guide por Peter Saint-Andre, Kevin Smith e Remko Tronçon foi usado como referência para implementar o protocolo XMPP.