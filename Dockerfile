FROM golang:latest

WORKDIR $GOPATH/src/kesarsauce/cors-proxy

COPY . $GOPATH/src/kesarsauce/cors-proxy

RUN go build -o /cors-proxy

EXPOSE 50051

CMD [ "/cors-proxy" ]
