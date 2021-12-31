# ENDPOINT

|  ENDPOINT | METHOD  | ACCESS  | 
|---|---|---|
|  /register |  POST |  all |
|  /login |  POST | all  |
|  /articles | GET  | all  |
|  /articles | POST  | all |
|  /articles/{articleId} | GET  |  all |
|  /articles/{articleId}| DELETE  | author  |
|  /articles/{articleId} | PUT  | author  |
|  /articles/{articleId}/comments | POST  | all  |
|  /articles/{articleId}/comments/{commentId} | DELETE  | author & commentator  |
|  /articles/{articleId}/comments/{commentId} | PUT  | commentator  |


## HOW TO RUN THIS PROJECT
1. clone this project 
2. move to project folder 
```shell
cd /path/to/project
```
3. run docker-compose
```javascript
docker-compose -f docker-compose.yaml up
```
4. set username name and password in mongodb
```javascript
// run this if you want create user and this user can read write any where
db.createUser(
    {
        user: "root",
        pwd: "root",
        roles: ["userAdminAnyDatabase", "readWriteAnyDatabase"]
    }
)

// run this if you want to create user for specific database
db.createUser({
    user: "root",
    pwd: "root",
    roles: [
        {
            role: "readWrite",
            db: "dbName"
        }
    ]
})
```
5. stop mongo container
```shell
docker-compose -f docker-compose.yaml stop
```
6. re-create container
```shell
docker-compose -f docker-compose.yaml up
```

7create index on user collection
```shell
db.users.createIndex({username:1,email:1},{unique:true})
```

8. create schema validation, run schema on /app/schema/articles.js

9. run http server
```shell
go run main.go
```