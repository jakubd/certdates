
FROM golang:latest
ADD . /go/src/certdates
WORKDIR /go/src/certdates
RUN go get all
RUN go build main.go
CMD ["./main", "--domains", "domains.txt"]