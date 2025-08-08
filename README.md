# Go SOCKS5 Proxy

A simple, cross-platform SOCKS5 proxy server written in Go.

## Features

* Easy to set up and run
* Cross-platform (Linux, Windows, macOS)
* Docker support
* Optional username/password authentication

## Installation

### From Source

1. **Clone the repository:**
   ```sh
   git clone https://github.com/your-username/go-socks5-proxy.git
   cd go-socks5-proxy
   ```

2. **Build the binary:**
   ```sh
   go build main.go
   ```

### Pre-compiled Binaries

1.  Download the appropriate binary for your operating system from the [releases page](https://github.com/your-username/go-socks5-proxy/releases).
2.  Run the binary from your terminal.

## Usage

### Basic Usage

To start the proxy on the default port `1080`:

```sh
./go-socks5-proxy
```

### With Authentication

To start the proxy with a username and password:

```sh
./go-socks5-proxy -user myuser -pass mypassword
```

### Docker

1. **Build the Docker image:**
   ```sh
   docker build -t go-socks5-proxy .
   ```

2. **Run the Docker container:**
   ```sh
   docker run -p 1080:1080 go-socks5-proxy
   ```

   To run with authentication:
   ```sh
   docker run -p 1080:1080 go-socks5-proxy -user myuser -pass mypassword
   ```

## Mobile Support (iOS/Android)

Running a standalone Go binary on mobile is not the standard approach. Instead, you can use `gomobile` to compile the Go code into a library that can be embedded in a native application.

1.  **Install `gomobile`:**
    ```sh
    go get golang.org/x/mobile/cmd/gomobile
    gomobile init
    ```

2.  **Create a Go package for mobile:**

    Create a new directory `mobile/` and a file `mobile/proxy.go` with a function to start the proxy. This function can then be exported for use in a native application.

3.  **Bind the package for iOS/Android:**
    ```sh
    # For Android (creates an .aar file)
    gomobile bind -target=android go-socks5-proxy/mobile

    # For iOS (creates a .framework file)
    gomobile bind -target=ios go-socks5-proxy/mobile
    ```

4.  **Integrate the library:**

    Integrate the generated `.aar` or `.framework` file into your Android or iOS project.
