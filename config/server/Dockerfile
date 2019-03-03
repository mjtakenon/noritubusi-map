FROM golang:1.12.0 
WORKDIR $GOPATH/src/noritubusi-map/server

COPY . .

RUN go get -u github.com/golang/dep/cmd/dep

CMD ["./start"]
