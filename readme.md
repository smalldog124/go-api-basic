# Go API Basic
เขียน api ด้วย Golang โดยใช้ [Gin Framework](https://github.com/gin-gonic/gin)
 และเชื่อมต่อกับ MongoDB,Postgres

 ## Run Database ผ่าน Docker
[MongoDB](https://hub.docker.com/_/mongo)
```sh
docker run --name liza-shop -p 27017:27017 -d mongo:3.6
```
[Postgres](https://hub.docker.com/_/postgres)
```sh
docker run --name liza-shop -e POSTGRES_PASSWORD=example -e POSTGRES_USER=admin -p 5432:5432 -d postgres:11
```

## Library
Gin :: https://github.com/gin-gonic/gin

mgo :: https://pkg.go.dev/gopkg.in/mgo.v2

postgres :: https://github.com/jmoiron/sqlx