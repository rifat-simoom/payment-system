package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/rifat-simoom/payment-system/internal/card/ports"
	"github.com/rifat-simoom/payment-system/internal/card/service"
	"github.com/rifat-simoom/payment-system/internal/common/logs"
	"github.com/rifat-simoom/payment-system/internal/common/server"
	"google.golang.org/grpc"
)

func main() {
	logs.Init()

	ctx := context.Background()

	application := service.NewApplication(ctx)

	serverType := strings.ToLower(os.Getenv("SERVER_TO_RUN"))
	switch serverType {
	case "http":
		go loadFixtures(application)

		server.RunHTTPServer(func(router chi.Router) http.Handler {
			return ports.HandlerFromMux(
				ports.NewHttpServer(application),
				router,
			)
		})
	case "grpc":
		server.RunGRPCServer(func(server *grpc.Server) {
			svc := ports.NewGrpcServer(application)
			trainer.RegisterTrainerServiceServer(server, svc)
		})
	default:
		panic(fmt.Sprintf("server type '%s' is not supported", serverType))
	}
}
