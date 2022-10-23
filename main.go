package main

import (
	"context"
	"flag"
	pb "kesarsauce/cors-proxy/proto"
	"log"
	"net"
	"google.golang.org/grpc"
)

// album represents data about a record album.
type album struct {
    ID     string  `json:"id"`
    Title  string  `json:"title"`
    Artist string  `json:"artist"`
    Price  float32 `json:"price"`
}

// albums slice to seed record album data.
var albums = []album{
    {ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 57.99},
    {ID: "2", Title: "Queen", Artist: "Gerry Mulligan", Price: 17.99},
    {ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
    {ID: "4", Title: "Hocus Pocus", Artist: "Focus", Price: 49.99},
    {ID: "5", Title: "Sylvia", Artist: "Focus", Price: 89.99},
    {ID: "6", Title: "Bohemian Rhapsody", Artist: "Queen", Price: 69.99},
}

type inventoryServer struct {
	pb.UnimplementedInventoryServer
}

func (s *inventoryServer) GetAlbumList(ctx context.Context, req *pb.GetAlbumListRequest) (*pb.GetAlbumListResponse, error) {
    albumsResp := []*pb.Album{}
    for _, album := range albums {
        albumsResp = append(albumsResp, &pb.Album{Id: album.ID, Title: album.Title, Artist: album.Artist, Price: album.Price})
    }
    return &pb.GetAlbumListResponse{Albums: albumsResp}, nil
}

func (s *inventoryServer) GetAlbumById(ctx context.Context, req *pb.GetAlbumByIdRequest) (*pb.GetAlbumByIdResponse, error) {
    id := req.GetId()
    for _, album := range albums {
        if album.ID == id {
            return &pb.GetAlbumByIdResponse{Album: &pb.Album{Id: album.ID, Title: album.Title, Artist: album.Artist, Price: album.Price}}, nil
        }
    }
    return &pb.GetAlbumByIdResponse{}, nil
}

func newServer() *inventoryServer {
    s := &inventoryServer{}
    return s
}

func main() {
    flag.Parse()
    lis, err := net.Listen("tcp", "localhost:50051")
    if err != nil {
    log.Fatalf("failed to listen: %v", err)
    }
    var opts []grpc.ServerOption
    grpcServer := grpc.NewServer(opts...)
    
    pb.RegisterInventoryServer(grpcServer, newServer())
    grpcServer.Serve(lis)
}
/*
// getAlbums responds with the list of all albums as JSON.
func getAlbums(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, albums)
}

// postAlbums adds an album from JSON received in the request body.
func postAlbums(c *gin.Context) {
    var newAlbum album

    // Call BindJSON to bind the received JSON to
    // newAlbum.
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    albums = append(albums, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

// getAlbumByID locates the album whose ID value matches the id
// parameter sent by the client, then returns that album as a response.
func getAlbumByID(c *gin.Context) {
    id := c.Param("id")

    // Loop through the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range albums {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}
*/