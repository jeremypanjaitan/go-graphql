# About
Implement clean architecture on graphql project. 
# How to run ?
- to download dependencies run `go mod download` 
- then run the project `go run .` 
# Routes
I have also include the postman collection if you want to use postman.
## Create
 ```
 curl --location --request POST 'http://localhost:3000/api/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation {\r\n    create(name:\"Inca Kola\",info:\"Inca Kola is a soft drink that was created in Peru in 1935 by British immigrant Joseph Robinson Lindley using lemon verbena (wiki)\",price:1.99){\r\n        id,name,info,price\r\n    }\r\n}","variables":{}}'
```
## Read Specific
```
curl --location --request POST 'http://localhost:3000/api/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"query {\r\n    product(id: 1){\r\n        id\r\n        name\r\n        price\r\n    }\r\n}","variables":{}}'
```
## Read All
```
curl --location --request POST 'http://localhost:3000/api/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"{\r\n    productList {\r\n        id\r\n        name\r\n        \r\n    }\r\n}","variables":{}}'
```
## Update
```
curl --location --request POST 'http://localhost:3000/api/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation {\r\n   update(id:1,name:\"test\",info:\"test info\",price:3.95){\r\n       id,name,info,price\r\n    }\r\n}","variables":{}}'
```
## Delete
```
curl --location --request POST 'http://localhost:3000/api/graphql' \
--header 'Content-Type: application/json' \
--data-raw '{"query":"mutation {\r\n   delete(id:1){\r\n       id,name,info,price\r\n    }\r\n}","variables":{}}'
```
# Requirements
- go version >= 1.17.3