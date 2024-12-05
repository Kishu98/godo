# Godo

A multi-tool CLI app with few tools which I use almost everyday.


## Features

- CLI/TUI
- Add tasks to todo list
- Write journals
- Check Weather
- Check top weekly dev news

## Usage/Examples
To get the top weekly dev news:
```bash
go run main.go news
```
To get all the tools:
```bash
go run main.go --help
```
To get all the flags for any tool:
```bash
go run main.go todo -h
```
To run TUI, run any tool without any flags:
```bash
go run main.go todo
```


## Run Locally

Clone the project

```bash
  git clone https://github.com/Kishu98/godo.git
```

Go to the project directory

```bash
  cd godo
```

Install dependencies

```bash
  go run main.go
```


## Tech Stack

Go, bubbletea framework, Cobra

## Roadmap

- Add database integration to store data.
- Keep adding new tools
- Deploy this app to the world

