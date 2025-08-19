# 1. Build stage
FROM golang:1.22 AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app

# 2. Minimal runtime
FROM scratch
COPY --from=build /app/app /app
EXPOSE 8080
ENTRYPOINT ["/app"]