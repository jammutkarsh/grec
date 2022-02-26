
# CLI Data Entry App

This a simple data entry app based on CLI.

An experimental exploratory project for me to get familiar with golang and
CLI app building process.

## Requirements

- [Go](https://golang.org/doc/install)

## Run Locally

Clone the project

```bash
  git clone https://github.com/JammUtkarsh/grec
```

Go to the project directory

```bash
  cd go-cli-records
```

Build executable

```bash
  go build -o app main.go json-operations.go
```

Run executable

```bash
  ./app
```

Commands & flags

```bash
  # For help 
  ./app help
  # Only multi-word statements need "double quotes".
  ./app add -d "Let's Goooo"   
  # List all the entries.
  ./app ls   
  # Delete a specific entry with its index.
  ./app delete -d 4
  # Delete the whole list.
  ./app clear 
```
