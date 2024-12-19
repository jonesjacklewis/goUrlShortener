# Go URL Shortener

A simple URL shortener written in Go to help me learn the language.

Includes unit tests and two endpoints.

## Endpoints

### POST /shorten

Shortens a URL.

#### Request

```json
{
  "url": "https://www.google.com"
}
```

#### Response

```json
{
    "short": "http://localhost:8080/long/aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8="
}
```

### GET /long/{short}

Redirects to the original URL.

#### Request

```
GET /long/aHR0cHM6Ly93d3cuZ29vZ2xlLmNvbS8=
```

#### Response

Redirects to `https://www.google.com`.

## Running

```bash
go run short.go
```