# golangTodoList

Ceci est un petit projet de Todo List écrit avec golang.

Pour l'utiliser, il faut:

A)Avec docker
1)Cloner le projet

2)Lancer le projet golang et mysql ensemble en faisant: 

  docker-compose up --build
  
3) Voir le container et son état:

  docker ps
4)Aller dans postman pour tester les endpoints ou utiliser curl 😊

B) Sans docker

1)Cloner le projet

2)Configurez les variables d'environnement en renommant le .env.example
  en .env et remplissez-le avec vos informations.

3)Installer les dépendances en exécutant dans votre terminal :

    go mod download

4) Lancez l'application avec la commande :

    go run .

