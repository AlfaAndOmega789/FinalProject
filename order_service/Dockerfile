FROM golang:1.22

ENV GOTOOLCHAIN=auto

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd

RUN go build -o main .

CMD ["./main"]