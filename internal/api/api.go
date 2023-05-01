package api

import (
	"github.com/ansrivas/fiberprometheus/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"
	_ "github.com/zhenisduissekov/http-3rdparty-task-service/docs"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
	"github.com/zhenisduissekov/http-3rdparty-task-service/internal/handler"
)

const (
	reqTimeFormat  = "15:04:05"
	reqLogFormat   = "[${time}] ${status} - ${latency} ${method} ${path} ${ip} ${url} in ${bytesReceived} bytes/ out ${bytesSent} bytes\n"
	prometheusPath = "/metrics"
	healthPath     = "/health"
	timeZone       = "Local"
)

func New(h *handler.Handler, conf *config.Conf) *fiber.App {
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
	
	//app.Use(basicauth.New(basicauth.Config{ //todo: basic auth needs to be uncommented in production
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
		api.Post("/task", h.AssignTask)   // add a task to the queue
    	api.Get("/task/:id", h.CheckTask) // get the status of a task by using its id  [pending/in_process/done/new/failed]
    	api.Get("/task", h.GetAllTasks)   // get all tasks -- for debugging purposes only	
	}

	return app
}
