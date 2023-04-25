FROM golang:1.20-alpine

WORKDIR /app

COPY . .

RUN go mod download && go mod tidy

RUN go build -o /insta-bot-comment

CMD ["/insta-bot-comment"]