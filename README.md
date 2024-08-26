# Proyecto 1 Redes - Uso de un protocolo existente
[![es](https://img.shields.io/badge/lang-es-yellow.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.md)
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.en.md)
[![fr](https://img.shields.io/badge/lang-fr-gr.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.fr.md)
[![pt](https://img.shields.io/badge/lang-pt-blue.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.pt.md)


## Objetivos del proyecto
- Implementar un protocolo en base a los estándares.
- Comprender el propósito del protocolo XMPP.
- Comprender cómo funcionan los servicios del protocolo XMPP.
- Aplicar los conocimientos adquiridos en programacion Web y móvil

## Objetivos Logrados
- Se desarrollo un cliente para un chat basado en XMPP completamente en el lenguaje de programación Go
- El cliente cuenta con las siguientes funcionalidades
    - Registrar una nueva cuenta en el servidor
    - Iniciar sesión con una cuenta
    - Cerrar sesión con una cuenta
    - Eliminar la cuenta del servidor
    - Mostrar contactos y sus estados
    - Adicionar contactos 
    - Mostrar detalles del contacto
    - Comunicación 1 a 1 con un contacto
    - Definición de presencia en servidor
    
- Adicionalmente se implemento un protocolo XMPP desde cero para realizar la comunicación con el servidor.

## Requisitos de instalación
- GoLang -> https://go.dev/doc/install

## Instalación
1. Clonar repositorio:
```bash
    git clone https://github.com/adrianfulla/Proyecto1-Redes.git
```
2. Acceder al directorio del repositorio:
```bash
    cd Proyecto1-Redes/
```

3. Ejecutar comando:
```bash
    go mod tidy
```
Este comando instalara todas las dependencias necesarias

4. Ejecutar el comando go run para compilar y ejecutar el cliente:
```bash
    go run ./server
```

## Notas
- Se utilizó la librería de Fyne y su documentación para crear la visualización del cliente. https://github.com/fyne-io/fyne
- Se utilizó el recurso de ChatGPT con el modelo de Go Golang tanto para la creación del protocolo XMPP como para la creación de la interfaz de usuario. La conversación se puede encontrar aca: 
https://chatgpt.com/share/2923f8f4-38c2-44c0-b8d2-bc6538e32ba8
- Se utilizó el libro XMPP: The Definitive Guide por Peter Saint-Andre, Kevin Smith y Remko Tronçon como referencia para la implementación del protocolo XMPP.