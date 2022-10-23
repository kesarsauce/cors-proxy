package main

import (
 "context"
 "log"
 "net/http"

 "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
 "google.golang.org/grpc"

 pb "kesarsauce/music-albums-server/proto"
)



func main() {
    ctx := context.Background()
    gwmux := runtime.NewServeMux()
    opts2 := []grpc.DialOption{grpc.WithInsecure()}
    err2 := pb.RegisterInventoryHandlerFromEndpoint(ctx, gwmux,  "0.0.0.0:50051", opts2)
    if err2 != nil {
        log.Fatalf("failed to listen in http: %v", err2)
    }

    log.Println("Listening on port 8081")
  port := ":8081"
  http.ListenAndServe(port, gwmux)
}