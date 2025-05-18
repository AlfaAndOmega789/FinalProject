FROM golang:1.22

ENV GOTOOLCHAIN=auto

WORKDIR /app

COPY . .

RUN go build -o main .

CMD ["./main"]
