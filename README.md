# User Management System

## Introduction

This is a backend project for a User Management System, written in Go. It provides a RESTful API for creating, reading, updating, and deleting users in a SQLite database.

## Built With

- Go - The programming language used
- Gin - The web framework used
- GORM - The ORM library used
- SQLite - The database used

## Getting Started

### Prerequisites

- Go 1.16 or later
- SQLite

### Installation

1. Clone the repository to your local machine.
2. Install the necessary Go packages by running `go get`.

## Running the Application

1. Run `go run main.go` in the root directory of the project.
2. The server will start on `localhost:8008`.

## Usage

The application provides the following endpoints:

- `GET /users`: Fetch all users.
- `GET /users/:id`: Fetch a user by ID.
- `POST /users`: Create a new user.
- `PATCH /users/:id`: Update a user by ID.
- `DELETE /users/:id`: Delete a user by ID.
