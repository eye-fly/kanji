## Overview

This is a web-based application for word learning/memorization. It's similar to [AnkiWeb](https://apps.ankiweb.net/), but was designed primarily for local, small-scale use. Despite its limited scope, optimization and security were important consideration.

## Features

- Scheduled data backups using `crontab` and a simple Bash script
- Caching of statick files on client side
- Optimized database structure to minimize the amount of user-specific data stored
- multi-level logging, many-to-many relationships, robust error handling based on good practises, modular code/clean code.

## Implemented Functionalities 
- User authentication (login and registration)
- Review and lesson sessions
- Individual user assignment management
- Unlocking New worlds based on current progress
- Ability to search and add new words before default queue
- Integration with the [WaniKani API](https://docs.api.wanikani.com/20170710/#introduction)

## Technologies Used

- **Go** (Golang) for Backend Development
- **SQLite** for Database Management
- **Gorilla Mux** for Routing
- **SamOnzeWeb/godb** for ORM-like Database Interaction

## Project Structure

```
│── subject/         # Subject data processing and integration with DB
│── waniAPI/         # Integration with WaniKani API
│── webAPI/
│   │── json/     # JSON handling
│   │── learning/ # Endpoint to add new items to the learning queue or change their order
│   │── lesson/   # Lesson session management
│   │── review/   # Review session management
│   │── user/     # User authentication and session management
│── main.go          # Main server entry point
│── subjects.db      # SQLite database file
```

# Some implementation Details
## Security
User credentials are being hashed(sha256) on client side and onlt then sent to the server. Session Management is done with use of session cookies/token and has token expiry, are invalidated after logout and cookies served as session tokens are http-only.

things that are still missing: randomizing hashes for same passwords(of two different ussers), input whitelisting/validation.

## Unit tests
there are [tests](./subjects/assignment/sqlInterface_test.go) for some parts of the project that have more complex logic.

## Optimization
Usage of composite primary key (user_id, subject_id), no data duplication in tables, fetching only relevant data, updating user data is done in batches when its possible

# Installation

**Note:** Currently, the authentication logic is not enforced. All requests are treated as if they belong to user ID 101. To enable authentication and multiple users, uncomment the following [function](./webAPI/user/func.go)
  ```go
  func CheckCookieAndGetUserId(db *godb.DB, r *http.Request) (int, error)
  ```

### Steps:
1. Install dependencies:
   ```sh
   go mod tidy
   ```
2. Ensure that `subjects.db` exists, or create a new one using
   ```
   sqlite3 subjects.db < schema.sql
   ```
3. Run the server:
   ```sh
   go run main.go
   ```

## Configuration

- Change the server port if needed:
  ```go
  err = http.ListenAndServe(":8080", router)
  ```

