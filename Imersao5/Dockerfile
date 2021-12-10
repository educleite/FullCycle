FROM golang:1.17

WORKDIR /go/src

RUN apt-get update && apt-get install build-essential librdkafka-dev -y 
# && apt-get install gccgo-go    

CMD ["tail", "-f", "/dev/null"]