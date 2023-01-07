# golangTodoList

Ceci est un petit projet de Todo List écrit avec golang.

Pour l'utiliser, il faut:

A)Avec docker

1)Cloner le projet

2)Créer un fichier .env à la racine du projet

3)Copier les informations du fichier .env.example dans le fichier .env que vous venez de créer et modifier-les avec vos propres informations.

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

3)Copier les informations du fichier .env.example dans le fichier .env que vous venez de créer et modifier-les avec vos propres informations. (DB_HOST=127.0.0.1)

4)Installer les dépendances en exécutant dans votre terminal :

    go mod download

5) Lancez l'application avec la commande :

    go run .

C) Avec kubernetes ( minikube )

1) Installer minikube :  https://kubernetes.io/docs/tasks/tools/install-minikube/

2) Démarrer un cluster Kubernetes local sur votre ordinateur avec la commande: minikube start

3) Créer un namespace: kubectl create ns nom-du-namespace (dans mon cas le namespace s'appelle todolist-golang-mysql)

4) Créer les différents objets du manifest en lançant la commande: kubectl apply -f statefulset.yml -n nom-du-namespace 

5)Afficher tout ce qui vient d'être créer: kubectl get all -n nom-du-namespace

6) Afficher les pods et voir leurs états: kubectl get pod -n nom-du-namespace

7) Voyer votre api en action en lançant la commande: minikube service -n nom-du-namespace golang-service

Puis cliquer sur l'url généré.


NB: Pour les secrets, c'est en base 64. Pour changer les informations de connexion à la bdd, vous pouvez procéder comme suit:

Entrer dans votre terminal et taper par exemple: echo root | base64 

