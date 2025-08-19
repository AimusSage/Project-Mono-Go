# Student–Tutor Matching API (Go)

As a test with Go I wrote this simple REST API that stores students and tutors in memory and provides a naive "match".  
This project is intended as a learning exercise to explore Go.

---

## 🚀 Features
- Add students (`POST /students`)
- Add tutors (`POST /tutors`)
- Get a naive student–tutor match (`GET /match`)
- Health check (`GET /health`)
- In-memory data storage (no database needed)
- Thread-safe with a basic mutex
- Ready to run locally or in Docker

---

## 🛠 Requirements
- [Go 1.22+](https://go.dev/dl/)
- [Docker](https://www.docker.com/) (optional, for containerized run but runs just fine in terminal)

---

## ▶️ Run locally

```bash
# clone project
git clone https://github.com/aimussage/student-tutor-api-go.git
cd student-tutor-api

# initialize module (first time only)
go mod init example.com/student-tutor-api

# run server
go run .
````

Server runs at:

```
http://localhost:8080
```

---

## 🧪 API Usage

### Health check

```bash
curl http://localhost:8080/health
# -> ok
```

### Add a student

```bash
curl -X POST http://localhost:8080/students \
  -H "Content-Type: application/json" \
  -d '{"name":"Aimus","age":15}'
```

### Add a tutor

```bash
curl -X POST http://localhost:8080/tutors \
  -H "Content-Type: application/json" \
  -d '{"name":"Sage","subject":"Math"}'
```

### Get a match

```bash
curl http://localhost:8080/match
```

Example response:

```json
{
  "student": { "id": 1, "name": "Aimus", "age": 15 },
  "tutor": { "id": 2, "name": "Sage", "subject": "Math" }
}
```

---

## 🐳 Run with Docker

Build image:

```bash
docker build -t student-tutor-api .
```

Run container:

```bash
docker run --rm -p 8080:8080 student-tutor-api
```

Test:

```bash
curl http://localhost:8080/health
# -> ok
```

---

## 🧹 Project Structure

```
student-tutor-api/
 ├── main.go        # API implementation
 ├── go.mod         # Go module definition
 ├── main_test.go   # unit tests
 └── Dockerfile     # container definition
```

---

## 📌 Notes

* Data is stored **in memory** and will be lost when the server restarts.
* This project is a minimal learning exercise; real apps should use a database and stronger matching logic.

```
