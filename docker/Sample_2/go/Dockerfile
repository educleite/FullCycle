FROM golang:1.17

WORKDIR /go/src

COPY main.go . 

RUN go build -o server main.go

CMD [ "./server" ] 