# Lenslocked

A Go web application for managing photo galleries.

## Project Structure

```
.
├── cmd/
│   └── exp/           # Experimental code and examples
├── templates/         # HTML templates
│   ├── home.gohtml    # Home page template
│   ├── contact.gohtml # Contact page template
│   └── faq.gohtml     # FAQ page template
├── main.go           # Main application entry point
├── go.mod            # Go module definition
├── go.sum            # Go module checksums
├── modd.conf         # Development configuration
└── README.md         # Project documentation
```

## Getting Started

1. Clone the repository
2. Run `go mod download` to install dependencies
3. Run `go run main.go` to start the server
4. Visit `http://localhost:3000` in your browser

## Features

- Home page
- Contact page
- FAQ page
- Gallery management