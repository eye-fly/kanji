## Overview

SQL Filler is a web-based API project that provides functionalities for user authentication, review sessions, lesson sessions, and learning module management. It integrates with a SQLite database and uses the `gorilla/mux` router for handling HTTP requests.

## Features

- User Authentication (Login & Registration)
- Review Sessions
- Lesson Sessions
- Learning Module Management
- API for Data Exchange with [WaniKani API](https://docs.api.wanikani.com/20170710/#introduction)
- Static File Handling
- SQLite Database Integration

## Technologies Used

- **Go** (Golang) for Backend Development
- **SQLite** for Database Management
- **Gorilla Mux** for Routing
- **SamOnzeWeb/godb** for ORM-like Database Interaction

## Project Structure

```
/sql_filler
│── subject/         # Subject data processing/integration with DB
│── waniAPI/         # Integration with WaniKani API
│── webAPI/json/     # JSON handling
│── webAPI/learning/ # Endpoint to add new items to the learning queue or change order
│── webAPI/lesson/   # Lesson session management
│── webAPI/review/   # Review session management
│── webAPI/user/     # User authentication and session management
│── main.go          # Main server entry point
│── subjects.db      # SQLite database file
```

## Installation

**Note:** Currently, the authentication logic is not enforced. All requests are treated as if they belong to user ID 101. To enable authentication and multiple users, uncomment the following function in `webAPI/user/func.go`:
  ```go
  func CheckCookieAndGetUserId(db *godb.DB, r *http.Request) (int, error)
  ```

### Steps:
1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Ensure that `subjects.db` exists, or create a new one using `table.txt`.
3. Run the server:
   ```sh
   go run main.go
   ```

## Configuration

- Change the server port if needed:
  ```go
  err = http.ListenAndServe(":8080", router)
  ```

