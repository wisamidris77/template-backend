# template-backend

A template you can use for backend made with Go.

## Project Structure

- **main.go**: The entry point of the application. It initializes the database connection, sets up the router, and starts the server.
- **config/database.go**: Contains the logic for connecting to the database and running migrations.
- **config/seeder.go**: Seeds the database with initial data.
- **models/user.go**: Defines the User model and its methods for password hashing and validation.
- **router/router.go**: Sets up the main application routes and middleware.
- **router/dtos/userDto.go**: Defines the data transfer object for the User model.
- **auth/router.go**: Sets up the authentication routes.
- **auth/models.go**: Defines the request and response models for authentication.
- **auth/auth_jwt.go**: Contains the logic for generating and validating JWT tokens and the authentication middleware.

## Disclaimer

This is my first app made in Go, so don't blame me if there is anything bad or wrong. Feel free to contribute and improve the code.

And at last >> Star us 🤩
## Getting Started

1. Clone the repository.
2. Run `go mod tidy` to install dependencies.
3. Run `go run main.go` to start the server.
4. If you got error like set CGO_ENABLED=1 or something like that just run the init.cmd/init.sh to fix it
5. You should be install gcc because CGO_ENABLED=1 is required by sqlite3 go package and CGO_ENABLED=1 needs gcc

## Detailed Instructions

For more detailed instructions, refer to the [Detailed Instructions](DETAILED_INSTRUCTIONS.md) file.