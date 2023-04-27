# Instagram Web Scraper with Golang and Chat GPT Comment Generator
This is a project that uses Golang and Chat GPT to create an Instagram Web Scraper that can comment on a publication with a generated comment.

## Installation
- Golang (https://golang.org/)
- Chat GPT Key
- Instagram Account
- or Run It with Docker

## Libs

- Colly
- Urfave CLI

## Usage
To use this application locally, you need to follow the steps below:

1. Clone the repository: **git clone https://github.com/roronoadiogo/insta-bot-comment.git**
1. Navigate to the directory: cd **insta-bot-comment**
1. Install the dependencies: 
```go
user@machine:~$ go mod download .
```
1. Add your Instagram username and password in **config.yaml** file
1. Run the application: 
```go
user@machine:~$ go run main.go
```

If you desire, run it with Docker

```console
user@machine:~$ docker build -t insta-bot-image .
```
and run 

```console
user@machine:~$ docker run insta-bot-image
```

The application run in CLI the commands below to execute: (In development)

## Contributing
Contributions are always welcome! If you find any bugs or want to add new features to this application, feel free to create a pull request or open a issue.

