# Wallet Service

This project is a wallet service that allows users to create wallets, perform transactions, and check wallet balances. It uses Docker for containerization and PostgreSQL as the database.

## Project Structure

- **cmd/main.go**: The entry point of the application.
- **pkg**: Contains the main logic for handling wallets and transactions.
- **Makefile**: Defines useful commands for managing the development environment and Docker containers.
- **docker-compose.yml**: Defines the services and how they interact.
- **.env.example**: Example environment file for configuring the application.

## Prerequisites

Ensure you have the following tools installed on your machine:

- [Go](https://golang.org/dl/) (version 1.16+)
- [Docker](https://www.docker.com/get-started)
- [Docker Compose](https://docs.docker.com/compose/)

## Setup

1. Clone the repository:

   ```bash
   git clone https://github.com/asror96/wallet.git
   cd wallet
   
2. The project has a fold for convenience. To find out all the commands type:
    ```bash
   make help
3. Build the project
    ```bash
   make build