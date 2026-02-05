# Status Code Testing Checklist

Use this checklist to manually verify the HTTP status codes requirements.

## 1. Verify 200 OK (Success)
- [ ] **Home Page**: Open `http://localhost:8080/`.
  - **Expected**: Page loads, standard banner selected. (Status 200)
- [ ] **Creation**: Enter text "test", click "Submit".
  - **Expected**: ASCII art appears in the Result box.

## 2. Verify 404 Not Found (Missing Resources)
- [ ] **Invalid Path**: Navigate to `http://localhost:8080/nothinghere`.
  - **Expected**: Message "404 Not Found" appears on screen.
- [ ] **Missing Banner**: 
  - Stop server (`Ctrl+C`).
  - Rename `banners/standard.txt` to `banners/standard.bak`.
  - Start server (`go run main.go`).
  - Try to generate art with "standard" banner.
  - **Expected**: Message "404 Not Found: Banner file not found".
  - *Restore the file afterwards!*

## 3. Verify 400 Bad Request (Invalid Request)
- [ ] **POST to Home**: Use Curl or Postman to send POST to `/`.
  - Command: `curl -X POST http://localhost:8080/`
  - **Expected**: "400 Bad Request"
- [ ] **GET to ASCII-Art**: Navigate to `http://localhost:8080/ascii-art` in browser.
  - **Expected**: "400 Bad Request"
- [ ] **Invalid Banner**: Use Curl to send invalid banner name.
  - Command: `curl -d "text=hi&banner=fake" -X POST http://localhost:8080/ascii-art`
  - **Expected**: "400 Bad Request: Invalid banner"

## 4. Verify 500 Internal Server Error
- [ ] **Template Error**:
  - Stop server.
  - Create a temporary bug in `handlers/handlers.go` (e.g., change template name to "missing.html").
  - Start server and reload home page.
  - **Expected**: "500 Internal Server Error".
  - *Undo the change afterwards!*
