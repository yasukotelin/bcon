FROM golang:1.11.5-stretch

RUN go get -u github.com/yasukotelin/bcon

ENTRYPOINT ["bcon"]