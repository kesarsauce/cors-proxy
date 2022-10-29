FROM golang:latest

WORKDIR $GOPATH/src/kesarsauce/music-albums-server

COPY . $GOPATH/src/kesarsauce/music-albums-server

RUN go build -o /music-albums-server

EXPOSE 8080

CMD [ "/music-albums-server" ]
