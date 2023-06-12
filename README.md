# gorss

## Goal

The goal with gorss is to practice Golang

## ⚙️ Installation

### Prerequistes

- Postgres DB
- Create a .env file with PORT and DB_URL environment variables

#### Suggestion: 

You can create a Postgres container instead of installing it. 

```bash
docker run --name postgresql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d postgres
```

For PORT and DB_URL you can use:

```
PORT=8000
DB_URL=postgres://root:12345@localhost:5432/rss?sslmode=disable
```


Inside a Go module:

```bash
go build && ./goRSS 
```
