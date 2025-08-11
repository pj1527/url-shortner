# URL Shortener

A simple URL shortener service built with Go.

## Getting Started

### Prerequisites

- [Go](https://golang.org/doc/install) installed on your machine.

### Setup

1.  Clone the repository:
    ```sh
    git clone https://github.com/pj1527/url-shortner.git
    cd url-shortner
    ```

2.  Tidy up the modules:
    ```sh
    go mod tidy
    ```

## Usage

1.  **Run the server:**
    ```sh
    go run .
    ```
    The server will start on `http://localhost:8080`.

2.  **Shorten a URL:**
    Use `curl` to send a `POST` request with the long URL in the body:
    ```sh
    curl -X POST -d "https://www.google.com" http://localhost:8080/shorten
    ```
    The response will be the shortened URL, for example: `http://localhost:8080/P7Cv`

3.  **Use the shortened URL:**
    Open the shortened URL in your browser, and you will be redirected to the original long URL.

