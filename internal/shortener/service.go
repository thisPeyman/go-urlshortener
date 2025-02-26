package shortener

import (
	"context"

	"github.com/redis/go-redis/v9"
	"github.com/thisPeyman/go-urlshortner/api"
	"github.com/thisPeyman/go-urlshortner/internal/pkg/url_shortener/repository"
	"github.com/thisPeyman/go-urlshortner/pkg/utils"
)

type ShortenerService struct {
	api.UnimplementedShortenerServiceServer
	idGenService api.IDGeneratorServiceClient
	redis        *redis.Client
	db           *repository.Queries
}

func NewShortenerService(redis *redis.Client, idGenService api.IDGeneratorServiceClient, db *repository.Queries) *ShortenerService {
	return &ShortenerService{
		redis:        redis,
		idGenService: idGenService,
		db:           db,
	}
}

func (s *ShortenerService) ExpandURL(ctx context.Context, req *api.ExpandURLRequest) (*api.ExpandURLResponse, error) {
	longUrl, err := s.db.GetLongURL(ctx, req.ShortUrl)
	if err != nil {
		return nil, err
	}

	return &api.ExpandURLResponse{
		LongUrl: longUrl,
	}, nil
}

func (s *ShortenerService) ShortenUrl(ctx context.Context, req *api.ShortenURLRequest) (*api.ShortenURLResponse, error) {
	idResponse, err := s.idGenService.GenerateID(ctx, &api.GenerateIDRequest{})
	if err != nil {
		return nil, err
	}
	randomID := idResponse.RandomID

	shortUrl := utils.EncodeToBase62(randomID)

	err = s.db.CreateShortURL(ctx, repository.CreateShortURLParams{
		ShortUrl: shortUrl,
		LongUrl:  req.LongUrl,
	})
	if err != nil {
		return nil, err
	}

	return &api.ShortenURLResponse{ShortUrl: shortUrl}, nil
}
