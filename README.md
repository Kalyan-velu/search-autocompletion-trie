# Autocomplete Editor

A web-based editor with real-time autocomplete suggestions using a Trie data structure, Go, and HTMX.

## Project Structure

├── .github/
│ └── copilot-instructions.md
├── view/
│ ├── index.css
│ └── index.html
├── go.mod
├── main.go
├── trie.go
└── README.md

## Components

- [`main.go`](main.go): HTTP server and request handlers
- [`trie.go`](trie.go): Trie data structure implementation
- [`view/index.html`](view/index.html): Frontend interface with HTMX integration

## Features

- Real-time autocomplete suggestions using a Trie data structure
- Case-insensitive word suggestions
- Dynamic word learning from user input
- Responsive design for mobile devices
- Server-side implementation in Go
- Frontend using HTMX for seamless interactions

## Getting Started

1. Ensure Go 1.23.3 or later is installed
2. Clone the repository
3. Run the server:
4. Open [http://localhost:8080](http://localhost:8080) in your browser

### How It Works

The application uses a `Trie data structure` to store and search words efficiently. As users type in the editor, HTMX sends requests to the autocompleteHandler which searches the trie for matching prefixes and returns suggestions.

The server automatically learns new words from user input (words longer than 2 characters) and adds them to the trie for future suggestions.

### API Endpoints

- `GET /`: Serves the main editor interface
- `GET /autocomplete?q={query}`: Returns autocomplete suggestions for the given query

### Technologies Used

Go 1.23.3
HTMX 1.5.0
TailwindCSS 2.2.19
