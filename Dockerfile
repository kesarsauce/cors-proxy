FROM golang:latest

WORKDIR $GOPATH/src/kesarsauce/cors-proxy

COPY . $GOPATH/src/kesarsauce/cors-proxy

RUN go build -o /cors-proxy

EXPOSE 8080

CMD [ "/cors-proxy" ]