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
)

func New(h *handler.Handler, conf *config.Conf) *fiber.App {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
		AllowHeaders: "*",
	}))

	app.Static("/", "./")

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
	//app.Use(basicauth.New(basicauth.Config{
	//	Users: map[string]string{
	//		conf.Auth.Username: conf.Auth.Password,
	//	},
	//	Realm: "Forbidden",
	//}))
	app.Use(logger.New(logger.Config{
		Format:       reqLogFormat,
		TimeFormat:   reqTimeFormat,
		TimeZone:     "Local",
		TimeInterval: 0,
		Output:       nil,
	}))
	app.Post("/api/v1/task", h.AssignTask)   // assign a task
	app.Get("/api/v1/task/:id", h.CheckTask) // check task status
	app.Get("/dummy", somethiung)
	return app
}

// ShowAccount godoc
//	@Summary		Show an account
//	@Description	get string by ID
//	@Tags			accounts
//	@Accept			json
//	@Produce		json
//	@Param			id	path		int	true	"Account ID"
//	@Success		200	{object}	string
//	@Failure		400	{object}	string
//	@Failure		404	{object}	string
//	@Failure		500	{object}	string
//	@Router			/accounts/{id} [get]
func somethiung(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"status": "success",
	})
}