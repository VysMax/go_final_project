FROM golang:1.22.4

WORKDIR /app

COPY . .

RUN go mod download

ENV GOOS=linux GOARCH=amd64 TODO_PORT=7540 TODO_DBFILE=scheduler.db TODO_PASSWORD=highPiglet%

EXPOSE ${TODO_PORT}

RUN go build -o /startApp main.go

CMD ["/startApp"]



