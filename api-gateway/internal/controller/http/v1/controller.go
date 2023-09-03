package v1

import (
	"api-gateway/internal/config"
	dto2 "api-gateway/internal/domain/queue/dto"
	"api-gateway/internal/domain/queue/service"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log/slog"
	"strconv"
)

type Handler struct {
	service service.Service
	app     *fiber.App
	config  *config.Config
}

func NewHandler(service service.Service, app *fiber.App, config *config.Config) *Handler {
	return &Handler{
		service: service,
		app:     app,
		config:  config,
	}
}

func (h *Handler) Register() {
	routes := h.app.Group("/")
	routes.Put("/:category", h.Put)
	routes.Get("/:category", h.Get)
}

func (h *Handler) Put(ctx *fiber.Ctx) error {
	category := ctx.Path()
	v := ctx.Query("v", "")
	if v == "" {
		ctx.Status(404)
		ctx.Send(nil)
		return fmt.Errorf("can't add new value to queue without value")
	}
	dto := dto2.NewPutQueueInput(category, v)
	res, err := h.service.Put(ctx.Context(), dto)
	if err != nil {
		ctx.Status(404)
		ctx.Send(nil)
		return fmt.Errorf("can't add new value to queue, try again later")
	}
	if res.Error != "" {
		ctx.Status(int(res.Status))
		ctx.Send(nil)
		return fmt.Errorf(res.Error)
	}
	ctx.Status(int(res.Status))
	return nil
}

func (h *Handler) Get(ctx *fiber.Ctx) error {
	category := ctx.Path()
	timeout := ctx.Query("timeout", "")
	if timeout == "" {
		timeout = "0"
	}
	timeoutInt, err := strconv.Atoi(timeout)
	if err != nil {
		ctx.Status(404)
		ctx.Send(nil)
		slog.Error(err.Error())
		return err
	}
	dto := dto2.NewGetQueueInput(int64(timeoutInt), category)
	res, err := h.service.Get(ctx.Context(), dto)
	if err != nil {
		ctx.Status(404)
		ctx.Send(nil)
		return fmt.Errorf("can't get value from queue, try again later")
	}
	if res.Error != "" {
		ctx.Status(int(res.Status))
		return fmt.Errorf(res.Error)
	}
	ctx.Status(int(res.Status))
	ctx.Send([]byte(res.Item))
	return nil
}
