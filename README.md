# Go HTTP Server

This is a reusable Go HTTP server with routing capabilities. It utilizes the `router` package for routing and the `logger` package for logging.

## Features

- Routing using the `router` package
- Logging using the `logger` package
- Graceful shutdown on SIGINT and SIGTERM signals

## Installation

To use the HTTP server in your Go project, you need to import the required packages:

```go
import (
	"github.com/sdpsagarpawar/router"
	"github.com/sdpsagarpawar/logger"
	"github.com/your-username/your-package/myserver"
)
```

## Usage

1. Create a new instance of the server with the router, logger, and timeout:
```
r := router.NewRouter()
l := logger.NewLogger()

// Create a new instance of the server
server := myserver.NewServer(r, l, 10*time.Second)
```

2. Define your route handlers using the router package:
```
// Define route handlers
helloHandler := func(w http.ResponseWriter, r *http.Request) {
    // Your route handler logic here
}

// Add routes to the router
r.AddRoute("GET", "/hello", helloHandler)

```

3. Start the server by calling the Start method:
```
server.Start("8080")

```

This will start the server on port 8080 and listen for incoming HTTP requests.

## Graceful shutdown
The server will gracefully shut down when a SIGINT or SIGTERM signal is received. You can also trigger the shutdown programmatically if needed. The server will attempt to finish processing any ongoing requests before shutting down.

## Contributing
Contributions are welcome! If you find any issues or have suggestions for improvements, please feel free to open an issue or submit a pull request.

## License
This project is licensed under the MIT License.