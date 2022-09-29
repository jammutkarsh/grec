
# CLI Data Entry grec

This a simple data entry grec based on CLI.

An experimental exploratory project for me to get familiar with golang and
CLI grec building process.

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
  go build -o grec main.go json-operations.go
```

Run executable

```bash
  ./grec
```

Commands & flags

```bash
  ./grec help
  # For help 
  ./grec add -d "Let's Goooo"   
  # Only multi-word statements need "double quotes".
  ./grec ls   
  # List all the entries.
  ./grec paste 
  # outputs the last clip. 
  ./grec delete -d 4
  # Delete a specific entry with its index.
  ./grec clear 
  # Delete the whole list.
```
