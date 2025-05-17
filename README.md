# Student-API

This repository contains a REST API for performing various student-related operations such as creating, retrieving, updating, and deleting student records. The data is stored using SQLite, a lightweight file-based relational database.

### Features

- Create a new student
- Get student details
- Update existing student data
- Delete student records

### Tech Stack

- **Language**: Go
- **Database**: SQLite
- **Framework**: net/http (standard library)

### Requirements:
Code Developed and Tested on: 
go version 1.23.6 linux/amd64

### Running the API
```bash
git clone https://github.com/greninja517/student-api.git
cd student-api
go run main.go
```
The API-endpoint will be available at http://localhost:9999/.