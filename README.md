## Description
This is the server component of the music storage application, built using Go. It handles data storage and serves the client application. Additionally, it provides APIs that allow users to download music from YouTube (usign an api key). 

All music files are saved in the `/etc/music` directory within the server container. You can update this path if necessary to match your specific setup.

## Installation
Before you begin, make sure you have Go installed.

```shell
go build cmd/main.go
```

### Run with Docker compose

You can use Docker Compose to run the server alongside the client. Here's a Docker Compose configuration:

```yaml
version: '3.8'
services:
  db:
    image: postgres:latest
    container_name: postgres
    restart: always
    environment:
      POSTGRES_USER: ${DB_USER}
      POSTGRES_PASSWORD: ${DB_PASSWORD}
      POSTGRES_DB: ${DB_NAME}
    volumes:
      - ./postgres-data:/var/lib/postgresql/data
    ports:
      - "5432:5432"

  frontend:
    build:
      context: ./swiftFrontend
      dockerfile: Dockerfile
    container_name: frontend
    ports:
      - "3000:3000"
    depends_on:
      - backend

  backend:
    build:
      context: ./swiftBackend
      dockerfile: Dockerfile
    container_name: backend
    environment:
      DB_HOST: ${DB_HOST}
      DB_PORT: ${DB_PORT}
      DB_USER: ${DB_USER}
      DB_PASSWORD: ${DB_PASSWORD}
      DB_NAME: ${DB_NAME}
    volumes:
      - /etc/music:/etc/music
    ports:
      - "4000:4000"
    depends_on:
      - db

volumes:
  postgres-data:
```
