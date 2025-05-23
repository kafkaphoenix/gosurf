# gosurf

## Description
gosurf is a simple API microservice that allows operations over users and actions.

## Architecture
This project follows clean architecture:
- domain: Contains the models.
- usecases: Operations that can be performed over the models.
- repository: Contains our fake database and http server

Additionaly, `cmd` is our entry point.

Please refer to [/docs/README.md](./docs/README.md) for more information about the project structure and design choices.

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

### 4. Testing endpoints with Bruno:
[Bruno](https://www.usebruno.com/) is a Fast and Git-Friendly Opensource API client. Collection can be found
in the `bruno` folder.

### 5. Testing endpoints with CURL:

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

### Get Referral index
```sh
curl -X GET "http://localhost:8081/v1/referral-index"
```
