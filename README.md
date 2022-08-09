# MongoDB CRUD operations using REST API in Golang.
Implementation of REST API containing "GET", "POST", "PUT", "DELETE" endpoints to perform CRUD operation in MongoDB using only Golang.

# Folder Structure
## controllers
This file contains all the functions we need to perform operations when endpoints are hit.
## routers 
This function routes all the incoming request to specified functions present in controllers.
## structure
Defines our Database Structure

## go.mod
Has all the required dependencies
## go.sum
Validated hash of dependencies
## main.go
Initiate our server

## Installation

Requires [Go](https://go.dev/dl/) 1.18+

Install the dependencies start the server.

```sh
1. Once you download the zip, Open the directory in Termial.
2. Run "go mod tidy" - All the required dependencies will be downloaded and installed
3. In controllers.go kindly replace const connectionString = " with your mongoDB database URL"
```

# Working

1. Run **"go run ./main.go"** to start the server. (localhost:1111)
2. Will register some users before performing all the operations. Refer **structure.go** for input json structure (**Cred**) and **routers.go** for endpoints.
3. After registrations we can do Login. Program will end if no username found else Login succesful. Refer **structure.go** for input json structure (**Cred**)
4. Now we can perform CRUD operations. Refer **structure.go** for input json structure (**Employee**)
5. Program ends if username is entered incorrect, although this can be handled in different ways.
6. In this project we have performed MongoDB CRUD operations using REST API in Golang. We can register new users. Authenticate their login.
If the username is present in the Employee DB we can perform all the operations. **We can set the flag on successfull login and skip the checking username in database.- not implemented.**
7. Operations as follows - Adding new employee, Updating employee details, Get all employee present in database, Delete one emplyee details, Delete whole database.

