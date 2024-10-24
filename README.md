# Golang Authentication and Authorization Project

This project demonstrates the implementation of authentication and authorization in a Golang application, following the guide from [Tanmay Vaish's tutorial](https://tanmay-vaish.hashnode.dev/how-to-implement-authentication-and-authorization-in-golang).

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
    git clone https://github.com/your-username/your-repo-name.git
    cd your-repo-name
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

## API Endpoints

- **POST /login**: Authenticate a user and generate a JWT token.
- **GET /profile**: Protected route that returns user profile information (requires valid JWT).
- **POST /admin**: Protected route that can be accessed only by Admin users.

## License
This project is open-source and available under the [MIT License](LICENSE).

## References
- Tutorial: [How to Implement Authentication and Authorization in Golang](https://tanmay-vaish.hashnode.dev/how-to-implement-authentication-and-authorization-in-golang)

