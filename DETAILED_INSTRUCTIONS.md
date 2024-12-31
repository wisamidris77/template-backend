# Detailed Instructions

## Project Setup

1. **Clone the Repository**:
   ```sh
   git clone https://github.com/wisamidris7/template-backend.git
   cd template-backend
   ```

2. **Install Dependencies**:
   ```sh
   go mod tidy
   ```

3. **Run the Server**:
   ```sh
   go run main.go
   ```

## Project Structure

### main.go
The entry point of the application. It initializes the database connection, sets up the router, and starts the server.
```go
// filepath: /C:/Users/wisam/Desktop/template/backend/main.go
// ...existing code...
```

### config/database.go
Contains the logic for connecting to the database and running migrations.
```go
// filepath: /C:/Users/wisam/Desktop/template/backend/config/database.go
// ...existing code...
```

### config/seeder.go
Seeds the database with initial data.
```go
// filepath: /C:/Users/wisam/Desktop/template/backend/config/seeder.go
// ...existing code...
```

### models/user.go
Defines the User model and its methods for password hashing and validation.
```go
// filepath: /c:/Users/wisam/Desktop/template/backend/models/user.go
// ...existing code...
```

### router/router.go
Sets up the main application routes and middleware.
```go
// filepath: /c:/Users/wisam/Desktop/template/backend/router/router.go
// ...existing code...
```

### router/dtos/userDto.go
Defines the data transfer object for the User model.
```go
// filepath: /c:/Users/wisam/Desktop/template/backend/router/dtos/userDto.go
// ...existing code...
```

### auth/router.go
Sets up the authentication routes.
```go
// filepath: /c:/Users/wisam/Desktop/template/backend/auth/router.go
// ...existing code...
```

### auth/models.go
Defines the request and response models for authentication.
```go
// filepath: /c:/Users/wisam/Desktop/template/backend/auth/models.go
// ...existing code...
```

### auth/auth_jwt.go
Contains the logic for generating and validating JWT tokens and the authentication middleware.
```go
// filepath: /c:/Users/wisam/Desktop/template/backend/auth/auth_jwt.go
// ...existing code...
```

## Running the Application

1. **Start the Server**:
   ```sh
   go run main.go
   ```

2. **Access the Application**:
   Open your browser and navigate to `http://127.0.0.1:8080`.

## API Endpoints

### User Endpoints
- `GET /api/users`: Fetch all users.
- `POST /api/users`: Create a new user.

### Authentication Endpoints
- `POST /auth/login`: Login and receive a JWT token.

## Contributing

Feel free to contribute and improve the code. Fork the repository, make your changes, and submit a pull request.

## Disclaimer

This is my first app made in Go, so don't blame me if there is anything bad or wrong.
