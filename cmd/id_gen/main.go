package main

import (
	"context"
	"fmt"
	"net"

	"github.com/thisPeyman/go-urlshortner/api"
	idgenerator "github.com/thisPeyman/go-urlshortner/internal/id_generator"
	snowflakego "github.com/thisPeyman/snowflake-go"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func ProvideSnowFlake() (*snowflakego.Snowflake, error) {
	return snowflakego.New(1)
}

func ProvideGrpcServer() *grpc.Server {
	return grpc.NewServer()
}

func RegisterGrpcService(server *grpc.Server, idService *idgenerator.IDGeneratorService) {
	api.RegisterIDGeneratorServiceServer(server, idService)
}

func StartGrpcServer(lc fx.Lifecycle, server *grpc.Server, log *zap.Logger) {
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			listener, err := net.Listen("tcp", ":50052")
			if err != nil {
				return err
			}
			fmt.Println("🚀 gRPC Server running on port 50052")
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

func main() {
	app := fx.New(
		// fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
		// 	return &fxevent.ZapLogger{Logger: log}
		// }),
		fx.Provide(
			ProvideGrpcServer,
			idgenerator.NewIDGeneratorService,
			ProvideSnowFlake,
			zap.NewExample,
		),
		fx.Invoke(
			RegisterGrpcService,
			StartGrpcServer,
		),
	)

	app.Run()
}
