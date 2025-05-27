FROM golang:1.22

ENV GOTOOLCHAIN=auto

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY .env ./

RUN go build -o main .

CMD ["./main"]
