# URL Shortener

A simple URL shortener service built with Go and Gin web framework.

## Features

- Shorten long URLs to short, memorable links
- Redirect short URLs to their original destination
- In-memory storage (data is lost on server restart)
- Simple REST API

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or higher
- Git (for cloning the repository)

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/pj1527/url-shortner.git
   cd url-shortner
   ```

2. Download dependencies:
   ```sh
   go mod download
   ```

### Running the Server

Start the server with the following command:

```sh
go run cmd/url-shortener/main.go
```

The server will start on `http://localhost:8080`.

## API Reference

### Shorten a URL

**Request**

```http
POST /shorten
Content-Type: application/json

{
    "url": "https://example.com/very/long/url"
}
```

**Example using curl:**

```sh
curl -X POST http://localhost:8080/shorten \
  -H "Content-Type: application/json" \
  -d '{"url":"https://example.com/very/long/url"}'
```

**Response**

```json
{
  "short_url": "http://localhost:8080/8M0Kx"
}
```

### Redirect to Original URL

**Request**

```http
GET /:shortKey
```

**Example using curl:**

```sh
curl -v http://localhost:8080/8M0Kx
```

**Response**

- Status: 302 Found
- Location: Original URL (e.g., `https://example.com/very/long/url`)

## Development

### Project Structure

```
.
├── cmd/
│   └── url-shortener/
│       └── main.go          # Application entry point
├── internal/
│   ├── handler/            # HTTP request handlers
│   ├── repository/         # Data access layer
│   └── service/            # Business logic
└── pkg/
    └── utils/              # Utility functions
```

### Notes

- This is an in-memory implementation, so all shortened URLs will be lost when the server restarts.
- The service runs on port 8080 by default.
- The base URL for shortened links is hardcoded to `http://localhost:8080`.

