FROM golang:1.12.0 
WORKDIR $GOPATH/src/noritubusi-map/server

COPY . .

# To check mysql's healthceck
RUN apt-get update && apt-get install -y mysql-client

RUN go get -u github.com/golang/dep/cmd/dep

CMD ["./delayed_start"]
