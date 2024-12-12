# Backend Challenge

This project is a JSON API built with **Golang** to manage Farms and their associated Crop Productions. The API is designed to be modular, robust, and production-ready.

## Features

- Create a Farm with nested Crop Productions.
- Delete a Farm by its ID.
- List all Farms with associated productions.
- Persistent data storage using PostgreSQL.
- Containerized using Docker for easy deployment.
- Comprehensive testing.

## Tech Stack

- **Golang**: Main programming language.
- **PostgreSQL**: Database for storing farms and productions.
- **Gorilla Mux**: Router for HTTP requests.
- **Docker**: Containerization for deployment.

## Prerequisites

- [Go](https://golang.org/dl/) 1.23
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)
- Makefile installed locally if you want to execute short commands defined
- PostgreSQL installed locally (if running without Docker).

## Setup

### Running Locally

1. Clone the repository:

   ```bash
   git clone https://github.com/gezielelyon/backend-challenge.git
   cd backend-challenge
   ```

1. Run the services:
   ```bash
   make prepare
   ```
   or
   ```bash
   docker compose up --build -d
   ```
