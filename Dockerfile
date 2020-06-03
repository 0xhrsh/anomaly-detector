FROM golang:1.14

COPY . $GOPATH/
# RUN go mod init bitbucket.org/harsh-not-haarsh/anomaly-detector/src/master
# RUN go mod verify
RUN go get github.com/lib/pq
RUN go get github.com/go-kit/kit/transport/http
RUN go get github.com/go-kit/kit/endpoint

RUN go build *.go

EXPOSE 8080

CMD ["./detector"]

