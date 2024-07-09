FROM golang:1.22

WORKDIR /src/app

COPY go.mod ./

RUN go mod tidy

COPY . .

RUN go build -v -o main .

EXPOSE 8080

CMD ["./main"]