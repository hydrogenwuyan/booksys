FROM golang:latest

ENV GOPROXY https://goproxy.cn,direct
WORKDIR $GOPATH/src/project/booksys
COPY . $GOPATH/src/project/booksys
RUN go build .

EXPOSE 12019
ENTRYPOINT ["./booksys"]