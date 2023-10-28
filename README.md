# Hands on GoLang 

## Setup
1. `go mod init github.com/kararnab/handsongo`
2. `go mod tidy` -> Install external dependencies, and remove unused ones (if any)
3. `make server-run` or `go run .` -> Starts dev server
4. `make server-build` -> Build the server executable

## Generate a random JWT Secret
```bash
node -e "console.log(require('crypto').randomBytes(32).toString('hex'))"
``` 
This script will generate a random 64 bit ASCII string, e.g. that can be used for encrypting JWT tokens.
Explanation is as follows-
1. `node -e` tells Node.js to evaluate a script, in this case, a Javascript string
2. `crypto` is the cryptographic module forming part of Node.js core. It is already installed as part of Node.js, no extra npm package is involved.
3. `randomBytes()` is a <u>function</u> that generates cryptographically strong pseudo-random data. It will return a Buffer object.
4. `toString()` is a method of the Buffer class that decodes the object to a string according to the specified character encoding, which, in this case, is hex, viz. hexadecimal.
5. Each of the 64 characters can be 0-9 and A,B,C,D,E or F


## Libraries Used
1. DotEnv `go get github.com/joho/godotenv`
2. JWT `github.com/golang-jwt/jwt/v5 v5.0.0`
3. PASETO `github.com/o1egl/paseto v1.0.0`
4. GIN `go get -u github.com/gin-gonic/gin`
5. [Structured Logging](https://github.com/topics/structured-logging) using `Zerolog` -> `go get -u github.com/rs/zerolog/log`
6. Stringer `go get -u golang.org/x/tools/cmd/stringer`


### To Study
- [ ] Channel and Goroutine
- [ ] Transaction Isolation level in PostGRES

## Popular Web Frameworks in GO
Although we can use the standard net/http package to implement APIs, it would be much easier to just take advantage of some existing web frameworks as they come loaded with features like routing, parameter binding, validation, middleware, built in ORM etc. Some of the most popular goLang web framework are-
- Gin (42.4k ★)
- Beego (25.1k ★)
- Echo (18.k ★)
- Revel (11.9k ★)
- Martini (11.1k ★)
- Fiber (8.6k ★)
- Buffalo (5.9k ★)

If you prefer a lightweight package with only routing feature, these are some of the most popular HTTP routers-
- FastHttp (13.6 ★)
- Gorilla Mux (12.9 ★)
- HttpRouter (11.9 ★)
- Chi (8.3 ★)


## Database Interaction Comparisons

### Tools

* Use [DB Diagram](https://dbdiagram.io/) to draw ER diagrams by just writing code.
* If you want to create any flow diagrams, try out [Excalidraw](https://excalidraw.com/), you can also check its [source code.](https://github.com/excalidraw/excalidraw)

### Database/SQL
- Very fast & straightforward
- Manual mapping SQL fields to variable
- Easy to make mistakes, not caught until runtime

### GORM
- CRUD functions already implemented, very short production code
- Must learn to write queries using gorm's function
- Run slowly on high load, 3-5 times slower than standard library

### SQLX
- Quite fast & easy to use, faster than GORM and near to standard library performance
- Field mapping via query text & struct tags
- Failure won't occur until runtime

### SQLC [link](https://sqlc.dev/)
- Very fast and easy to use
- Automatic code generation to native library CRUD code
- Catch SQL query errors before generating codes
- Full support Postgres. MySQL is experimental.

## Minikube & kubectl commands
- Minikube (Starting up/deleting the cluster)
- kubectl (configuring the minikube cluster)
1. `minikube start`
2. `minikube stop`
3. `minikube delete`
4. `minikube status`
5. `minikube dashboard` -> Access k8s dashboard running within the cluster
6. `kubectl version` -> Show client and server version of k8s
7. `kubectl logs [podname]` -> Log to console
8. `kubectl exec -it [podname] -- bin/bash` -> Get interactive terminal
9. `kubectl get nodes | pod | services | replicaset | deployment` [-o yaml] -> Status of different k8s component, as per etcd
10. `kubectl describe service service_name` -> Describes the service with selector, port, endpoint etc
11. `kubectl create | edit | delete deployment image_name` -> Create/edit/delete deployments
12. `kubectl apply -f nginx-deployment.yaml`
13. `kubectl delete -f nginx-deployment.yaml`


## Docker Commands frequently used

1. `docker pull postgres:12-alpine` -> A minimal Docker image of Postgres
2. `docker images`
3. `docker ps -a` -> To show all the containers, whether running or not
4. `docker ps` -> To show all running containers
5. `docker start|stop [containerId]`
6. `docker exec -it postgres12 /bin/sh` -> Go inside postgres shell

## Most frequent Git commands
Here are few of the important commands in Git
1. `git add .` -> Adds all files to be staged for next commit
2. `git commit --amend` -> Amends changes to the last commit, note that this changes commit SHA
3. `git commit --amend --date="now"` -> Updates the date timestamp of the commit
4. `git fetch -ap` -> Fetch with Append and Prune
5. `git log branchB..branchA` -> show the commits on branchA that are not on branchB
6. `git rebase master` -> apply any commits of current branch on top of master (chronologically)
7. `git reset HEAD{0}` -> Soft reset to the head i.e. most recent commit of current branch
8. `git branch -D <branch_name>` | `git branch -d <branch_name>` -> Deletes branch from local (D is force, d is normal)
9. `git push origin --delete old-branch` -> Deletes branch from remote
10. `git tag -a v1.4 -m "my version 1.4"` > `git push origin <tag_name>` ->  new annotated tag identified with v1.4

## Mocking, Testing and 🤞
For mocking, we use [Mockgen](https://github.com/golang/mock).
