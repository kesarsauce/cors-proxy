package main

import (
	"context"
	"fmt"
	pb "kesarsauce/music-albums-server/proto/albums"
	"net/http"
	"os"

	"github.com/twitchtv/twirp"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// album represents data about a record album.
type Album struct {
	gorm.Model
	ID     string  `json:"id"`
	Title  string  `json:"title"`
	Artist string  `json:"artist"`
	Price  float32 `json:"price"`
}

var DB *gorm.DB

type inventoryServer struct {
}

func (s *inventoryServer) GetAlbumList(ctx context.Context, req *pb.GetAlbumListRequest) (*pb.GetAlbumListResponse, error) {
	var albumsResp []*pb.Album
	result := DB.Find(&albumsResp)
    if result.Error != nil {
		return &pb.GetAlbumListResponse{}, twirp.NotFound.Error("Albums not found")
	}
	return &pb.GetAlbumListResponse{Albums: albumsResp}, nil
}

func (s *inventoryServer) GetAlbumById(ctx context.Context, req *pb.GetAlbumByIdRequest) (*pb.GetAlbumByIdResponse, error) {
	id := req.GetId()
	var album *pb.Album
	result := DB.First(&album, id)
	if result.Error != nil {
		return &pb.GetAlbumByIdResponse{}, twirp.NotFound.Error("Album not found")
	}
	return &pb.GetAlbumByIdResponse{Album: album}, nil
}

func (s *inventoryServer) AddAlbum(ctx context.Context, req *pb.AddAlbumRequest) (*pb.AddAlbumResponse, error) {
	album := req.GetAlbum()
	result := DB.Create(album)
	if result.Error != nil {
		return &pb.AddAlbumResponse{Success: false}, nil
	}
	return &pb.AddAlbumResponse{Success: true}, nil
}

func newServer() *inventoryServer {
	s := &inventoryServer{}
	return s
}

func main() {
	dsn := fmt.Sprintf(
        "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai",
        os.Getenv("PGSQL_HOST"),
        os.Getenv("PGSQL_USER"),
        os.Getenv("PGSQL_PASS"),
        os.Getenv("PGSQL_DB"),
        os.Getenv("PGSQL_PORT"),
    )
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	db.AutoMigrate(&Album{})

	DB = db

	twirpHandler := pb.NewInventoryServer(newServer())
	mux := http.NewServeMux()
	mux.Handle(twirpHandler.PathPrefix(), twirpHandler)
	http.ListenAndServe(":8080", mux)
}
