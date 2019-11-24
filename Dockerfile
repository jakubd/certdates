
FROM golang:latest
ADD . /go/src/certdates
WORKDIR /go/src/certdates
RUN go get all
RUN go build certdates.go
CMD ["./certdates", "--domains", "domains.txt", "--threshold", "60"]