# URL Shortener

A simple, high-performance URL shortener service built with Go and Gin web framework, featuring multiple storage backends and a clean architecture.

## Features

- **URL Shortening**: Convert long URLs to short, memorable links
- **Automatic Redirection**: Short URLs automatically redirect to their original destination
- **Multiple Storage Backends**:
  - In-memory (default, non-persistent)
  - Redis (persistent, recommended for production)
- **RESTful API**: Simple and intuitive API endpoints
- **Health Check**: Built-in health check endpoint
- **Environment-based Configuration**: Easy configuration through environment variables
- **Concurrent-safe**: Thread-safe implementation for high concurrency

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) 1.16 or higher
- [Git](https://git-scm.com/) (for cloning the repository)
- (Optional) [Redis](https://redis.io/) (for persistent storage)

### Environment Variables

Create a `.env` file in the `cmd/url-shortener` directory with the following variables:

```
PORT=8080
REDIS_ADDR=localhost:6379  # Only required if using Redis
```
- (Optional) [Redis](https://redis.io/docs/getting-started/) server if using Redis storage

### Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/pj1527/url-shortner.git
   cd url-shortner
   ```

2. Install dependencies:
   ```sh
   go mod download
   ```

### Running the Application

#### Using In-Memory Storage (Default)
```sh
go run cmd/url-shortener/main.go
```

#### Using Redis Storage
1. Make sure Redis is running
2. Set up your `.env` file with Redis configuration
3. Run the application:
   ```sh
   go run cmd/url-shortener/main.go
   ```

## API Documentation

### 1. Shorten a URL

**Request**
```bash
curl -X POST http://localhost:8080/api/shorten \
  -H "Content-Type: application/json" \
  -d '{"url": "https://example.com/very/long/url"}'
```

**Response**
```json
{
  "short_url": "http://localhost:8080/abc123"
}
```

### 2. Redirect to Original URL

**Request**
```bash
curl -L http://localhost:8080/abc123
```

### 3. Health Check

**Request**
```bash
curl http://localhost:8080/health
```
   
   **Response**:
   ```json
   {
     "status": "UP"
   }
   ```

## Architecture

The application follows a clean architecture with clear separation of concerns:

- **Handler**: HTTP request handling and response formatting
- **Service**: Business logic and use cases
- **Repository**: Data access layer with support for multiple storage backends
- **Config**: Environment configuration management
- **Utils**: Helper functions and utilities

## Configuration

| Environment Variable | Default | Description |
|----------------------|---------|-------------|
| `PORT` | `8080` | Port to run the server on |
| `REDIS_ADDR` | `localhost:6379` | Redis server address (if using Redis) |