# Projet 1 Réseaux - Utilisation d'un protocole existant
[![es](https://img.shields.io/badge/lang-es-yellow.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.md)
[![en](https://img.shields.io/badge/lang-en-red.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.en.md)
[![fr](https://img.shields.io/badge/lang-fr-gr.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.fr.md)
[![pt](https://img.shields.io/badge/lang-pt-blue.svg)](https://github.com/adrianfulla/Proyecto1-Redes/blob/main/README.pt.md)


## Objectifs du projet
- Implémenter un protocole basé sur les normes.
- Comprendre l'objectif du protocole XMPP.
- Comprendre le fonctionnement des services du protocole XMPP.
- Appliquer les connaissances acquises en programmation Web et mobile.

## Objectifs atteints
- Développement d'un client de chat basé sur XMPP entièrement en langage de programmation Go.
- Le client comprend les fonctionnalités suivantes :
    - Enregistrer un nouveau compte sur le serveur.
    - Se connecter avec un compte.
    - Se déconnecter avec un compte.
    - Supprimer le compte du serveur.
    - Afficher les contacts et leur statut.
    - Ajouter des contacts.
    - Afficher les détails des contacts.
    - Communication 1 à 1 avec un contact.
    - Définir la présence sur le serveur.
- De plus, un protocole XMPP a été implémenté à partir de zéro pour faciliter la communication avec le serveur.
    
## Exigences d'installation
- GoLang -> https://go.dev/doc/install

## Instalación
1. Cloner le dépôt:
```bash
    git clone https://github.com/adrianfulla/Proyecto1-Redes.git
```
2. Accéder au répertoire du dépôt :
```bash
    cd Proyecto1-Redes/
```

3. Exécuter la commande suivante :
```bash
    go mod tidy
```
Cette commande installera toutes les dépendances nécessaires.

4. Exécuter la commande go run pour compiler et exécuter le client :
```bash
    go run ./server
```

## Remarques
- La bibliothèque Fyne et sa documentation ont été utilisées pour créer la visualisation du client. https://github.com/fyne-io/fyne
- -ChatGPT, utilisant le modèle Go Golang, a été utilisé à la fois pour créer le protocole XMPP et pour développer l'interface utilisateur. La conversation est disponible ici : https://chatgpt.com/share/2923f8f4-38c2-44c0-b8d2-bc6538e32ba8
- Le livre XMPP: The Definitive Guide de Peter Saint-Andre, Kevin Smith et Remko Tronçon a été utilisé comme référence pour implémenter le protocole XMPP.