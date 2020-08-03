# R2R Social :: Social Network Backend 

Social Network Backend Project developed with Golang.

Go version: 1.14.2

### Go packages used:

  - Mongo Driver
  - Mongo Options
  - Mongo BSON
  - Mongo BSON Primitive
  - Bcrypt
  - Gorilla Mux
  - Cors
  - JWT-Go

### Packages Installation

Install the dependencies and devDependencies and start the server.

```sh
$ cd demo-social-project

$ go get go.mongodb.org/mongo-driver/mongo
$ go get go.mongodb.org/mongo-driver/mongo/options
$ go get go.mongodb.org/mongo-driver/bson
$ go get go.mongodb.org/mongo-driver/bson/primitive
$ go get golang.org/x/crypto/bcrypt
$ go get github.com/gorilla/mux
$ go get github.com/rs/cors
$ go get github.com/dgrijalva/jwt-go
```

### APIs 

This project expose the next APIs

| API | Method | Description |
| ------ | ------ | ------ |
| Login | POST | Login to obtain an access token |
| Register | POST | Register a new user |
| View Profile | GET | Retrieve information from user |
| Update Profile | PUT | Update personal info of user |
| New Post | POST | Publish a post |
| Get Posts | GET | Get published posts from user |
| Delete Post | DELETE | Delete a certain post |
| Upload Avatar | POST | Upload avatar image for user's profile  |
| Upload Banner | POST | Upload banner image for user's profile |
| New Relationship | POST | Create a relation between 2 user accounts |
| Delete Relationship | DELETE | Delete a relation between 2 user accounts |
| Get Relationship | GET | Get all relations from an user |
| User List | GET | List of users registered |
| Get Posts By Relationship | GET | Get Post of users related |

### Running Project
To run project just run follow commands

```sh
$ cd demo-social-project
$ go run main.go
```

Verify the deployment by navigating to your server address in your preferred browser.

```sh
127.0.0.1:8000
```
Have fun :)
