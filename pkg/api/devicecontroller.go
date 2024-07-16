package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
	"github.com/mole-squad/soq-api/api"
	"github.com/mole-squad/soq-api/pkg/auth"
	"github.com/mole-squad/soq-api/pkg/common"
	"github.com/mole-squad/soq-api/pkg/interfaces"
	"github.com/mole-squad/soq-api/pkg/models"
	"go.uber.org/fx"
	"gorm.io/gorm"
)

type DeviceControllerParams struct {
	fx.In

	AuthService   interfaces.AuthService
	DeviceService interfaces.DeviceService
	Router        *chi.Mux
}

type DeviceControllerResult struct {
	fx.Out

	DeviceController DeviceController
}

type DeviceController struct {
	deviceService interfaces.DeviceService
}

func NewDeviceController(params DeviceControllerParams) (DeviceControllerResult, error) {
	ctrl := DeviceController{deviceService: params.DeviceService}

	deviceRouter := chi.NewRouter()
	deviceRouter.Use(params.AuthService.AuthRequired())

	deviceRouter.Get("/", ctrl.ListDevices)
	deviceRouter.Post("/", ctrl.CreateDevice)

	deviceRouter.Route("/{deviceID}", func(r chi.Router) {
		r.Use(ctrl.deviceContextMiddleware)

		r.Patch("/", ctrl.UpdateDevice)
		r.Delete("/", ctrl.DeleteDevice)
	})

	params.Router.Mount("/devices", deviceRouter)

	return DeviceControllerResult{DeviceController: ctrl}, nil
}

func (ctrl *DeviceController) ListDevices(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	devices, err := ctrl.deviceService.ListUserDevices(ctx, user.ID)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	respList := []render.Renderer{}
	for _, device := range devices {
		respList = append(respList, device.AsDTO())
	}

	render.RenderList(w, r, respList)
}

func (ctrl *DeviceController) CreateDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	user, err := auth.GetUserFromCtx(ctx)
	if err != nil {
		render.Render(w, r, common.ErrUnauthorized(err))
		return
	}

	dto := &api.CreateDeviceRequestDTO{}
	if err = render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	newDevice := models.Device{
		UserKey:  dto.UserKey,
		DeviceID: dto.DeviceID,
	}

	device, err := ctrl.deviceService.CreateUserDevice(ctx, user, &newDevice)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, device.AsDTO())
}

func (ctrl *DeviceController) UpdateDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	device := r.Context().Value(deviceContextKey).(*models.Device)

	dto := &api.UpdateDeviceRequestDTO{}
	if err := render.Bind(r, dto); err != nil {
		render.Render(w, r, common.ErrInvalidRequest(err))
		return
	}

	update := models.Device{
		Model:    gorm.Model{ID: device.ID},
		UserKey:  dto.UserKey,
		DeviceID: dto.DeviceID,
	}

	updatedDevice, err := ctrl.deviceService.UpdateUserDevice(ctx, &update)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.Render(w, r, updatedDevice.AsDTO())
}

func (ctrl *DeviceController) DeleteDevice(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	device := r.Context().Value(deviceContextKey).(*models.Device)

	err := ctrl.deviceService.DeleteUserDevice(ctx, device.ID)
	if err != nil {
		render.Render(w, r, common.ErrUnknown(err))
		return
	}

	render.NoContent(w, r)
}

func (ctrl *DeviceController) deviceContextMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		deviceID, err := strconv.Atoi(chi.URLParam(r, "deviceID"))
		if err != nil {
			render.Render(w, r, common.ErrInvalidRequest(fmt.Errorf("failed to parse deviceID: %w", err)))
			return
		}

		user, err := auth.GetUserFromCtx(ctx)
		if err != nil {
			render.Render(w, r, common.ErrUnauthorized(err))
			return
		}

		device, err := ctrl.deviceService.GetUserDevice(r.Context(), user.ID, uint(deviceID))
		if err != nil {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				render.Render(w, r, common.ErrNotFound)
			} else {
				render.Render(w, r, common.ErrUnknown(err))
			}

			return
		}

		ctxWithDevice := context.WithValue(r.Context(), deviceContextKey, &device)

		next.ServeHTTP(w, r.WithContext(ctxWithDevice))
	})
}
