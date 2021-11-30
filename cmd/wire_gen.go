// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//+build !wireinject

package cmd

import (
	"context"
	"github.com/comeonjy/go-kit/pkg/xlog"
	"github.com/comeonjy/working/configs"
	"github.com/comeonjy/working/internal/data"
	"github.com/comeonjy/working/internal/server"
	"github.com/comeonjy/working/internal/service"
)

import (
	_ "net/http/pprof"
)

// Injectors from wire.go:

func InitApp(ctx context.Context, logger *xlog.Logger) *App {
	configsInterface := configs.NewConfig(ctx)
	dataData := data.NewData(configsInterface)
	workRepo := data.NewWorkRepo(dataData)
	workingService := service.NewWorkingService(configsInterface, logger, workRepo)
	grpcServer := server.NewGrpcServer(workingService, configsInterface, logger)
	httpServer := server.NewHttpServer(ctx, configsInterface, logger)
	app := newApp(ctx, grpcServer, httpServer, configsInterface)
	return app
}
