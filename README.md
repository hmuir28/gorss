# gorss

## Goal

The goal with gorss is to practice Golang

## ⚙️ Installation

### Prerequistes

- Postgres DB

#### Suggestion: 

You can create a Postgres container instead of installing it. 

```bash
docker run --name postgresql -e POSTGRES_USER=root -e POSTGRES_PASSWORD=12345 -p 5432:5432 -d postgres
```

Inside a Go module:

```bash
go get github.com/hmuir28/gorss
go build && ./goRSS 
```
