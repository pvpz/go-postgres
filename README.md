# Go-Postgres
This project is simple CRUD application built in golang and using PostgreSQL as DB.

## PostgreSQL Table
```sql
CREATE TABLE users
(
    id SERIAL PRIMARY KEY,
    name varchar(30),
    rating INTEGER
);
```
## Installation
```
docker-compose up
```

## HTTP Requests
Get all users
```
http://localhost:8080/api/user
```
Get user
```
http://localhost:8080/api/user/1
```
Create user
```
curl -X POST -H "Content-Type: application/json" \    -d '{"name": "petya", "rating": 17}' \    http://localhost:8080/api/user
```
Update user
```
curl -X PUT -H "Content-Type: application/json" \    -d '{"name": "petya", "rating": 117}' \    http://localhost:8080/api/user/1
```
Delete user
```
curl -X DELETE -H "Content-Type: application/json" \    -d '{}' \    http://localhost:8080/api/user/1
```