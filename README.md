# ASCII Art Web

A web application that generates ASCII art from text using different banner styles.

## Description

This web app converts text input into ASCII art using Go's `net/http` package. It supports three banner styles (standard, shadow, thinkertoy) and displays results in real-time.

## Author

**Dixon Osure**

## Usage

```bash
# Run the server
go run main.go

# Open browser
http://localhost:8080
```

## Implementation Details

### Algorithm

1. **Input Processing**: Text is split by newlines and each line is processed separately
2. **Character Mapping**: Each ASCII character (32-126) maps to a position in the banner file: `index = (char - 32) * 9`
3. **Art Generation**: For each character, 8 lines are extracted from the banner and concatenated horizontally
4. **Real-time Display**: JavaScript sends POST requests on input change with 300ms debounce

### Endpoints

| Method | Route | Description |
|--------|-------|-------------|
| GET | `/` | Serves the main HTML page |
| POST | `/ascii-art` | Generates ASCII art from text and banner |

### HTTP Status Codes

- `200` - Success
- `400` - Bad request (invalid input/method)
- `404` - Not found (missing banner/template)
- `500` - Internal server error

### Project Structure

```
ascii-art/
├── main.go              # Server entry point
├── handlers/
│   ├── handlers.go      # HTTP handlers
│   └── handlers_test.go # Handler tests
├── Artprinter/
│   ├── artprinter.go    # ASCII art generation
│   └── artprinter_test.go
├── templates/
│   └── index.html       # Web interface
└── banners/
    ├── standard.txt
    ├── shadow.txt
    └── thinkertoy.txt
```
