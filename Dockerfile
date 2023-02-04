FROM golang:1.15

WORKDIR /go/src

CMD ["tail", "-f", "/dev/null"]