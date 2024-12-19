# Go URL Shortener

A simple URL shortener written in Go, designed as a learning project to explore Go's features and build a functional web service. This project includes core functionality for URL shortening and redirection, unit tests, and persistent storage using SQLite.

---

## Features

- Shorten long URLs into unique, compact hashes.
- Redirect shortened URLs back to their original targets.
- Persistent storage using SQLite for scalability and reliability.
- Comprehensive unit tests for robust validation.
- Modular code structure for easy extensibility.

---

## Endpoints

### **1. POST /shorten**
Shortens a given URL and returns a shortened version.

#### **Request**
```http
POST /shorten HTTP/1.1
Content-Type: application/json

{
  "url": "https://www.google.com"
}
```

#### **Response**
```http
HTTP/1.1 200 OK
Content-Type: application/json

{
  "short": "http://localhost:8080/long/aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8="
}
```


### **2. GET /long/{short}**
Redirects to the original URL associated with the given shortened hash.

#### **Request**
```http
GET /long/aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8= HTTP/1.1
```

#### **Response**
Redirects to `https://www.google.com`.

## Requirements

- Go 1.16 or higher
- SQLite

## Running the Project

1. Clone the repository:

   ```bash
   git clone https://github.com/jonesjacklewis/goUrlShortener
   cd goUrlShortener
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Run the application:

   ```bash
    go run short.go
    ```
4. The application will start on `http://localhost:8080`.


## Testing

The project includes a comprehensive suite of unit tests to validate functionality.

To run the tests, use the following command:

```bash
go test ./...
```

## Potential Enhancements

- Expiration time for shortened URLs.
- Custom short URLs.
- Rate limiting and security features.
- Gurantee unique hashes as opposed to base64 encoding.

## Technologies Used

- Go: Programming language for the backend.
- SQLite: Lightweight database for persistent storage.

