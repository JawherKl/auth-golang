# Golang Authentication and Authorization Project

This project demonstrates the implementation of authentication and authorization in a Golang application.

## Features
- **User Authentication**: Secure login system using JWT (JSON Web Tokens).
- **Authorization**: Role-based access control for different user types (Admin, User).
- **Password Hashing**: Implements bcrypt for secure password storage.
- **REST API**: Exposes APIs for user login and restricted access to protected routes.

## Prerequisites
- Golang installed (Go 1.15 or higher)
- A running database (PostgreSQL, MySQL, etc.)

## Installation

1. Clone the repository:
    ```bash
    git clone https://github.com/JawherKl/auth-golang
    cd auth-golang
    ```

2. Install dependencies:
    ```bash
    go mod tidy
    ```

3. Configure the environment variables:
    - `JWT_SECRET`: The secret key for signing JWTs.
    - `DATABASE_URL`: Connection string for the database.

4. Run the project:
    ```bash
    go run main.go
    ```
## Setting up server directories/packages
For the organization of our server, we will categorize the directories into the following sections:
[tree-project](tree-project.png)

1. routes - This directory will contain the definitions of the different routes for the server, which map URLs to specific functions.
2. controllers - This directory will hold the functions that will handle the incoming requests and interact with the models and other components.
3. models - This directory will store the data models that represent the objects in the application and their relationships.
4. middlewares - This directory will contain any middleware functions that process requests and responses, such as authentication or logging.
5. utils - This directory will house various utility functions that are used throughout the application, such as helper functions or custom data types.

These directories will also serve as packages.

## API Endpoints

- **POST /login**: Authenticate a user and generate a JWT token.
- **POST /signup**: Create a user and inserted to database.
- **GET /home**: Protected route that returns home profile information (requires valid JWT).
- **GET /premium**: Protected route that returns premium information (requires valid JWT).
- **GET /logout**: Protected route that can be logout session user.

## License
This project is open-source and available under the [MIT License](LICENSE).

