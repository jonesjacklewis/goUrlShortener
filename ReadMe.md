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