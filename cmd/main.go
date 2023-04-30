package main

import (
	"context"
	"flag"

	"github.com/gofiber/fiber/v2"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/api"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/connection"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/handler"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/logger"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/repository"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/service"
	"go.uber.org/fx"
)

var (
	listenAddress = flag.String("port", ":3000", "Ports ")
)

//	@title			http-3rdparty-task-service API
//	@contact.name	API Support
//	@contact.email	zduisekov@gmail.com
//	@host			localhost:3000
//	@BasePath		/
//	@schemes		http
func main() {
	
	flag.Parse()
	//besides being short and concise, uber fx DI provides modularity and composition, also testability for future development
	app := fx.New(
		fx.Provide(
			config.New,
			logger.New,
			connection.New,
			service.New,
			repository.New,
			handler.New,
			api.New,
		),
		fx.Invoke(setupLifeCycle),
	)
	app.Run()
}

func setupLifeCycle(lc fx.Lifecycle, app *fiber.App) {
	var cancel context.CancelFunc
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			ctx, cancel = context.WithCancel(ctx)

			var err error
			srv := service.New()
			go srv.TaskQueue()
			go func(ctx context.Context) {
				flag.Parse()
				err = app.Listen(*listenAddress)
				ctx.Done()
			}(ctx)

			return err
		},
		OnStop: func(ctx context.Context) error {
			cancel()
			return app.Shutdown()
		},
	})
}


