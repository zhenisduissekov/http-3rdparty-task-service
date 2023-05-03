package main

import (
	"os"
	"os/signal"

	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"
	_ "github.com/zhenisduissekov/http-3rdparty-task-service/docs"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/handler"
	zlog "github.com/zhenisduissekov/http-3rdparty-task-service/internal/logger"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/repository"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/service"
)

const (
	reqTimeFormat  = "15:04:05"
	reqLogFormat   = "[${time}] ${status} - ${latency} ${method} ${path} ${ip} ${url} in ${bytesReceived} bytes/ out ${bytesSent} bytes\n"
	prometheusPath = "/metrics"
	healthPath     = "/health"
	timeZone       = "Local"
)

// @title			http-3rdparty-task-service API
// @contact.name	API Support
// @contact.email	zduisekov@gmail.com
// @host			localhost:3000
// @BasePath		/
// @schemes		http
func main() {
	cnf := config.New()
	zlog.New(cnf)

	repo := repository.NewRepository(cnf)
	srv := service.New(repo, cnf)
	h := handler.New(srv)
	app := router(h, cnf)

	go func() {
		srv.Task.StartQueue()
	}()

	go func() {
		if err := app.Listen(":" + cnf.Auth.Port); err != nil {
			log.Fatal().Err(err).Msg("error while starting the server")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	log.Info().Msg("Shutting down server...")
	srv.Task.CloseQueue()
	if err := app.Shutdown(); err != nil {
		log.Fatal().Err(err).Msg("Server forced to shutdown")
	}
	log.Info().Msg("Shutting down app...")
}

func router(h *handler.Handler, conf *config.Conf) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
		AllowHeaders: "*",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	prometheus := fiberprometheus.New(conf.ServiceName)
	prometheus.RegisterAt(app, prometheusPath)
	app.Use(prometheus.Middleware)
	app.Get(healthPath, func(f *fiber.Ctx) error {
		log.Trace().Msg("Health check")
		return f.Status(fiber.StatusOK).JSON(&fiber.Map{
			"status": "success",
		})
	})

	//app.Use(basicauth.New(basicauth.Config{ //todo: uncomment for production
	//	Users: map[string]string{
	//		conf.Auth.Username: conf.Auth.Password,
	//	},
	//	Realm: "Forbidden",
	//}))

	app.Use(logger.New(logger.Config{
		Format:       reqLogFormat,
		TimeFormat:   reqTimeFormat,
		TimeZone:     timeZone,
		TimeInterval: 0,
		Output:       nil,
	}))

	api := app.Group("/api/v1")
	{
		api.Get("/task/:id", h.Check) // get the status of a task by using its id  [pending/in_process/done/new/failed]
		api.Post("/task", h.Assign)   // add a task to the queue
	}

	return app
}
