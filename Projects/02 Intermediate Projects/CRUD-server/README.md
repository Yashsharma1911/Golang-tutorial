1)Download official mongoDB repository for Linux/Ubuntu
https://www.vultr.com/docs/install-and-configure-mongodb-database-server-on-ubuntu-20-04/


2)To test mongoDB locally Create instance mongoDB database on your local Machine (This project is based on Linux/Ubuntu)

Similar to the process here by following these commands

account_db> db.info.insertOne({"name":"John Doe","age":"30","email":"john@example.com"})

account_db> db.info.insertOne({"name":"Jane Smith","age":"25","email":"jane@example.com"})

account_db> db.info.find()

2)Import the MongoDB driver for your Golang application.

$ go get go.mongodb.org/mongo-driver/mongo

3)To test endpoints
POST: curl -X POST localhost:8080/api/v1/info -H "Content-Type: application/json" -d '{"name":"John Doe","age":"20","email":"example@domain.com"}'

GET: curl -X GET localhost:8080/api/v1/info

PUT: curl -X PUT localhost:8080/api/v1/info -H "Content-Type: application/json" -d '{"name":"Jane Smith","age":"20","email":"example@domain.com"}'

DELETE: curl -X DELETE localhost:8080/api/v1/info -H "Content-Type: application/json" -d '{"name":"John Doe"}'
