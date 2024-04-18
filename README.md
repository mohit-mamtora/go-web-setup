# go-web-setup
Welcome to **go-web-setup**, a rapid development template for Golang web applications. This setup is fortified with popular libraries and tools like Echo, sqlx, sql-migrate, and zap⚡️ to promise an extraordinary journey for your next project.


## Features

- **User Schema**: Comes with a predefined user schema.
- **Default APIs**: Includes standard APIs for user registration, login, and CRUD operations.
- **JWT Authentication**: Uses JSON Web Tokens for authentication.

## Setup

Follow these steps to get started with the project:

1. Clone the repository:
    ```bash
    git clone <repository_url>
    cd go-web-setup
    ```

2. Install dependencies:
    ```bash
    go get
    ```

3. Setup environment variables:
    - Create a `.env` file (refer to `.env.example` for the required variables).
    - Source the `.env` file:
        ```bash
        source .env
        ```

4. Start the application:
    ```bash
    docker-compose up -d
    go run cmd/prod/main.go
    ```
## Project Structure

The project follows a layered architecture:

    ├── app
    │   ├── dependencies.go
    │   ├── logger
    │   ├── model
    │   ├── repository
    │   ├── routes
    │   └── services
    ├── cmd
    │   ├── migrate
    │   ├── prod
    │   └── staging
    ├── config
    ├── logs
    └── migrations

## To-Do

- Update the readme with more detailed instructions.
- Write test cases to ensure code quality and reliability.
- Implement queue management (e.g., Horizon) for asynchronous processing.
- Add CSRF support for enhanced security.

## Contribution

Feel free to contribute to the project by forking the repository and raising pull requests. Your contributions are highly appreciated!