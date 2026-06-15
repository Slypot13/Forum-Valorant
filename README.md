# 🎮 ValoRoom

## 📌 Présentation

ValoRoom est un forum communautaire inspiré de Reddit, dédié à l'univers du jeu Valorant.

## 🚀 Fonctionnalités

* Gestion des utilisateurs
* Gestion des fils de discussion
* Gestion des messages
* Réactions
* Tri des messages
* Pagination
* Catégories
* Recherche
* Administration
---

## 🏗️ Architecture du projet

Le projet suit une architecture MVC :

### Models

Contiennent les structures de données.

### Repositories

Gèrent les requêtes SQL et l'accès à la base de données.

### Services

Contiennent la logique métier et les règles de gestion.

### Controllers

Gèrent les requêtes HTTP et les interactions avec les vues.

### Templates

Pages HTML affichées aux utilisateurs.

---

## 🛠️ Technologies utilisées

* Go (Golang)
* HTML
* CSS
* MySQL
* JWT
* Git
* GitHub

---

## 📂 Base de données

Principales tables :

* users
* threads
* messages
* reactions
* tags
* thread_tags

---

## ▶️ Lancer le projet

### Installer les dépendances

```bash
go mod tidy
```

### Lancer le serveur

```bash
go run main.go
```

### Accéder à l'application

```txt
http://localhost:8080
```

---

## 👨‍💻 Auteurs

POTTIER SYLVAIN
