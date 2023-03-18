// README.md

# Basic-Go-App

### Introduction

Basic-Go-App is an application REST API that enable users to CREATE, READ, UPDATE, and DELETE of all users daily activities data. These activities data including Title, Content, Image, and amount of user's content Likes.

### Project Support Features

- Users can post all users daily activity
- Users can update all users daily activity
- Users can get all users daily activity for general, and/or based on likes, and/or based on amount of data or pages they want to
- Users can get specific user daily activity based on ID
- Users can delete specific user's daily activity

### Installation Guide

- Clone this repository [here](https://github.com/relieyanhilman/basic-go-app).
- Run go mod tidy
- Setup the app.env configuration file and docker-compose.ymml file

### Usage

- Make sure postgre container is running (by running docker compose up)
- Run go run main.go to start the application
- Connect to the API using Postman on port 8000 (based on PORT configuration in app.env file).

### API Endpoints

| HTTP Verbs | Endpoints                                          | Action                                         |
| ---------- | -------------------------------------------------- | ---------------------------------------------- |
| POST       | /api/posts/                                        | To create user daily activity                  |
| GET        | /api/posts/?page=1&limit=10&likesGt=20&likesLt=100 | To retrieve all users posts                    |
| GET        | /api/posts/:postId                                 | To retrieve details of a user's daily activity |
| PUT        | /api/posts/:postId                                 | To edit the details of a user's daily activity |
| DELETE     | /api/posts/:postId                                 | To delete a user's daily activity              |

### Technologies Used

- [Go](https://go.dev/) Go is a statically typed, compiled high-level programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. It is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency.
- [GORM](https://gorm.io/) GORM is an Object Relational Mapping (ORM) library for Golang.
- [PostgreSQL](https://www.postgresql.org/) PostgreSQL, also known as Postgres, is a free and open-source relational database management system emphasizing extensibility and SQL compliance

### Authors

- [Relieyan Ramadhan Hilman](https://github.com/relieyanhilman)

### Credits

- [CODEVO](https://github.com/wpcodevo) for [this](https://codevoweb.com/build-restful-crud-api-with-golang/) article
- [article](https://www.moesif.com/blog/technical/api-design/REST-API-Design-Filtering-Sorting-and-Pagination/)
- [article](https://devpress.csdn.net/postgresql/62f228b17e66823466184be6.html)
- [article](https://medium.easyread.co/how-to-do-pagination-in-postgres-with-golang-in-4-common-ways-12365b9fb528)
- and other articles that I'cant write here
