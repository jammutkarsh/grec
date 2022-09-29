# CLI To-Do App

A simple CLI to-do application built in go.

An experimental exploratory project for me to get familiar with golang and building CLI apps

## System Requirements

- [Go](https://golang.org/doc/install)

## Install & Run

```bash
  go install github.com/JammUtkarsh/grec@latest
```

Commands & flags

```bash
  grec help
  # For help 
  grec add -d "Let's Goooo"   
  # Only multi-word statements need "double quotes".
  grec ls   
  # List all the entries.
  grec paste 
  # outputs the last clip. 
  grec delete -d 4
  # Delete a specific entry with its index.
  grec clear 
  # Delete the whole list.
```

## Run Locally

Clone the project

```bash
  git clone https://github.com/JammUtkarsh/grec
  ```

  ```bash
  cd grec
  ```

  ```bash
  go build .
```
