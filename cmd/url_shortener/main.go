package main

import (
	"context"
	"net"

	"github.com/jackc/pgx/v5"
	"github.com/labstack/echo/v4"
	"github.com/redis/go-redis/v9"
	"github.com/thisPeyman/go-urlshortner/api"
	"github.com/thisPeyman/go-urlshortner/internal/pkg/url_shortener/repository"
	"github.com/thisPeyman/go-urlshortner/internal/shortener"
	"github.com/thisPeyman/go-urlshortner/pkg/echoext"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func ProvideGrpcServer() *grpc.Server {
	return grpc.NewServer()
}

func ProvideHttpRouter() *echo.Echo {
	e := echo.New()

	e.Validator = echoext.NewCustomValidator()

	return e
}

func ProvideRedisClient() *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     "localhost:6370",
		Password: "",
		DB:       0,
	})
}

func ProvideDatabase(lc fx.Lifecycle) (*pgx.Conn, error) {
	conn, err := pgx.Connect(context.Background(), "host=localhost port=5430 user=postgres password=postgres dbname=postgres sslmode=disable")
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close(ctx)
		},
	})

	return conn, nil
}

func ProvideRepository(conn *pgx.Conn) *repository.Queries {
	return repository.New(conn)
}

func RegisterGrpcServices(server *grpc.Server, shortenerService *shortener.ShortenerService) {
	api.RegisterShortenerServiceServer(server, shortenerService)
}

func ProvideIDGeneratorService(lc fx.Lifecycle) (api.IDGeneratorServiceClient, error) {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			return conn.Close()
		},
	})

	return api.NewIDGeneratorServiceClient(conn), nil
}

func StartGrpcServer(lc fx.Lifecycle, server *grpc.Server, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", ":50053")
			if err != nil {
				return err
			}
			log.Info("🚀 gRPC Server running on port 50053")
			go func() {
				if err := server.Serve(listener); err != nil {
					log.Fatal("Failed to serve", zap.Error(err))
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			server.GracefulStop()
			log.Info("🛑 gRPC Server stopped")
			return nil
		},
	})
}

func StartHttpServer(lc fx.Lifecycle, e *echo.Echo, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				if err := e.Start(":40010"); err != nil {
					log.Fatal("Failed to serve", zap.Error(err))
				}
			}()
			return nil
		},
	})
}

func main() {
	app := fx.New(
		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log}
		// }),
		fx.Provide(
			ProvideGrpcServer,
			shortener.NewShortenerService,
			zap.NewExample,
			ProvideRedisClient,
			ProvideIDGeneratorService,
			ProvideDatabase,
			ProvideRepository,
			ProvideHttpRouter,
		),
		fx.Invoke(
			RegisterGrpcServices,
			shortener.RegisterController,
			StartGrpcServer,
			StartHttpServer,
		),
	)

	app.Run()
}
