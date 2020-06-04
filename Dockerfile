FROM golang:1.14

WORKDIR /
COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .
RUN go build *.go

EXPOSE 8080

CMD ["./detector"]

