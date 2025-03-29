# Architecture and Design Choices

## Introduction

This project balances KISS (Keep It Simple, Stupid) principles with Clean Architecture best practices, even if it required additional time to implement. In a rapid brainstorming scenario for a startup, I would have likely kept it minimal with:
- main.go handling the server, logging, and routing.
- A separate file for loading test data.
- A file for business logic.

## Architecture

The repository follows a Clean Architecture approach with three layers. However, the domain also serves as the DTO, and all layers use the same types, for simplicity.

The repository/infrastructure layer was introduced to separate HTTP handlers from the core logic for clarity. The same applies to the fake json DB loading logic.

Additionally, the project includes:
- Basic testing
- A Makefile
- CI setup

Docker and Docker Compose might be an overkill for this technical test or a brainstorming session(specially my multi-stage Docker build) but I wanted to showcase good practices and simplify the testing. In real-world projects, I'd also include Swagger for API documentation.

## API

I adhered to best practices for API naming conventions and kept things simple by avoiding third-party routers or frameworks like Echo.

## Golang Style Guide

To maintain readability and consistency, I followed Go’s idiomatic style and referred to:

>[Uber Go Style Guide](https://github.com/uber-go/guide/blob/master/style.md)

I enforced best practices using GolangCI-Lint and focused on meaningful, concise naming.

## Documentation

Good documentation is critical. I documented the API using Bruno and added inline comments where necessary. Bruno, an open-source alternative to Postman, allows API collections to be stored within the codebase, ensuring synchronization.

For convenience, I also provided cURL commands in the README for command-line API testing.

## Docker and Docker Compose

Following Docker best practices, I used multi-stage builds to keep the final image lightweight. Docker Compose is included for local deployment and reproducibility.

## Logging and Configuration

Given the project scope, I opted for Go’s built-in slog over third-party loggers like Zap or Zerolog. Similarly, I avoided .env files as just the port and the JSON file path were necessary. In larger projects, I like using [Viper](https://github.com/spf13/viper) for configuration management.

## Error handling & Graceful shutdown

I avoided sentinel errors to enhance extendibility and prevent error comparison pitfalls (==). Instead, I used Go’s built-in errors package to create contextual errors and ensured errors are either logged or returned—not both.

The application supports graceful shutdown on SIGTERM, ensuring resource cleanup. 

## Testing

For testing, I used Testify, as I believe a testing suite is a good practice.

If I had to expand the project with more test cases, I would consider using Mockery to generate mocks for interfaces. And for integration tests, I would use [Testcontainers](https://golang.testcontainers.org/quickstart/) if needed. And depending of what to test I would add [table driven tests](https://dave.cheney.net/2019/05/07/prefer-table-driven-tests).

Due to time constraints, I provided basic tests at the handler level without mocks. Additionally, Bruno is included for manual API validation.

## Project lifecycle

I used Git for version control, Gitlab CI and automated tasks with Makefile. 

## Optimizations

At the end of the project, I applied several optimizations, such as ordering user actions to make it easier to calculate the probabilities of the next action. I perform the ordering at the end since we load the data in bulk only once. I believe that adding ordering at the insertion stage would have increased complexity due to the overhead of binary search and value shifting.

For the fourth endpoint, we require a referral graph, which I moved to the database loading phase to avoid recreating it on each request.

Additionally, in a real-world scenario where large amounts of data are sent in responses, I would add middleware to compress the response.

I would also consider using caching, adding indexes to the database, and other performance optimizations.

I would also consider adding rlocks and locks if a DB was added depending of the query.
