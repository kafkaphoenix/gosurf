# gosurf

## Description
gosurf is a simple API microservice that allows operations over users and actions.

## Architecture
This project follows clean architecture:
- domain: Contains the models.
- usecases: Operations that can be performed over the models.
- repository: Contains our fake database and http server

Additionaly, `cmd` is our entry point.

## Requirements
To run this project, please follow the steps listed below:

### 1. Go installation
Follow the official [Go installation guide](https://go.dev/doc/install) to install Go 1.24.

### 2. Install Docker & Docker Compose
Ensure that Docker and Docker Compose are installed:
- [Docker Installation](https://docs.docker.com/get-docker/)
- [Docker Compose Installation](https://docs.docker.com/compose/install/)
> Docker has to be running to start the application.

### 3. Makefile
Ensure that Makefile is available to use the commands provided in the project,
otherwise run the commands manually.

## Try it out!

### 1. Start application:
```sh
make app
```

### 2. Clean any docker resource:
```sh
make clean
```

### 3. View application logs:
```sh
make logs
```

### 4. Testing endpoints:

#### Get User
```sh
curl -X GET http://localhost:8081/v1/users/1
```

### Get Total Actions given a User
```sh
curl -X GET http://localhost:8081/v1/users/1/actions/total
```

### Get Next Action probabilities given an Action type
```sh
curl -X GET "http://localhost:8081/v1/actions/next-probabilities?type=REFER_USER"
```
