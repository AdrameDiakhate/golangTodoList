# golangTodoList

Ceci est un petit projet de Todo List écrit avec golang.

Pour l'utiliser, il faut:

A)Avec docker

1)Cloner le projet

2)Créer un fichier .env à la racine du projet

3)Copier les informations dans .env.example dans le fichier .env que vous venez de créer et modifier-les avec vos propres informations.

4)Lancer le projet golang et mysql ensemble en faisant: 

  docker-compose up --build
  
5) Voir les containers et son état:

  docker ps

6)Aller dans postman pour tester les endpoints ou utiliser curl 😊

7)Arréter les containers

  docker compose down

B) Sans docker

1)Cloner le projet

2)Créer un fichier .env à la racine du projet

3)Copier les informations dans .env.example dans le fichier .env que vous venez de créer et modifier-les avec vos propres informations. (DB_HOST=127.0.0.1)

4)Installer les dépendances en exécutant dans votre terminal :

    go mod download

5) Lancez l'application avec la commande :

    go run .

