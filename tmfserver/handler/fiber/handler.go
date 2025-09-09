package fiber

import (
	"net/url"
	"strings"

	"encoding/json"
	"log/slog"

	"github.com/gofiber/fiber/v2"
	svc "github.com/hesusruiz/isbetmf/tmfserver/service"
)

// Handler is the handler for the TMF API (both V4 and V5).
type Handler struct {
	service *svc.Service
}

// extractAPIVersion extracts the API version from the URL path
func extractAPIVersion(path string) string {
	// Expected path format: /tmf-api/{apiFamily}/v{version}/...
	parts := strings.Split(path, "/")
	for i, part := range parts {
		if strings.HasPrefix(part, "v") && len(part) > 1 {
			// Check if this is likely a version (v4 or v5)
			if part == "v4" || part == "v5" {
				return part
			}
		}
		// Also check for the pattern where version comes after apiFamily
		if i > 0 && parts[i-1] != "" && strings.HasPrefix(part, "v") {
			return part
		}
	}
	// Default to v5 if not found
	return "v5"
}

// NewHandler creates a new handler.
func NewHandler(s *svc.Service) *Handler {
	return &Handler{service: s}
}

// Health is a simple hello world handler.
func (h *Handler) Health(c *fiber.Ctx) error {
	resp := &svc.Response{
		StatusCode: 200,
		Body:       "I am good, thanks",
	}
	return SendResponse(c, resp)
}

// CreateHubSubscription creates a new notification subscription (hub)
func (h *Handler) CreateHubSubscription(c *fiber.Ctx) error {
	jwtToken := svc.ExtractJWTToken(c.Get("Authorization"))

	// Extract API version from the URL path
	apiVersion := extractAPIVersion(c.Path())

	req := &svc.Request{
		Method:      c.Method(),
		Action:      svc.HttpActions[c.Method()],
		APIfamily:   c.Params("apiFamily"),
		APIVersion:  apiVersion,
		Body:        c.Body(),
		AccessToken: jwtToken,
	}

	resp := h.service.CreateHubSubscription(req)
	return SendResponse(c, resp)
}

// DeleteHubSubscription deletes an existing notification subscription (hub)
func (h *Handler) DeleteHubSubscription(c *fiber.Ctx) error {
	jwtToken := svc.ExtractJWTToken(c.Get("Authorization"))

	// Extract API version from the URL path
	apiVersion := extractAPIVersion(c.Path())

	idParam, _ := url.QueryUnescape(c.Params("id"))
	req := &svc.Request{
		Method:      c.Method(),
		Action:      svc.HttpActions[c.Method()],
		APIfamily:   c.Params("apiFamily"),
		APIVersion:  apiVersion,
		ID:          idParam,
		AccessToken: jwtToken,
	}

	resp := h.service.DeleteHubSubscription(req)
	return SendResponse(c, resp)
}

// MockListener is a minimal endpoint to receive notifications locally for testing
func (h *Handler) MockListener(c *fiber.Ctx) error {
	path := string(c.Request().URI().Path())
	body := c.Body()
	if len(body) > 0 {
		var payload any
		if err := json.Unmarshal(body, &payload); err == nil {
			slog.Info("listener received event", slog.String("path", path), slog.Int("bytes", len(body)), slog.Any("body", payload))
		} else {
			slog.Info("listener received event", slog.String("path", path), slog.Int("bytes", len(body)), slog.String("bodyRaw", string(body)))
		}
	} else {
		slog.Info("listener received event", slog.String("path", path), slog.Int("bytes", 0))
	}
	return c.SendStatus(fiber.StatusNoContent)
}

// CreateGenericObject creates a new TMF object using generalized parameters.
func (h *Handler) CreateGenericObject(c *fiber.Ctx) error {
	jwtToken := svc.ExtractJWTToken(c.Get("Authorization"))

	resourceName := c.Params("resourceName")
	if resourceName == "hub" {
		slog.Debug("handling hub creation")
		return h.CreateHubSubscription(c)
	}

	// Extract API version from the URL path
	apiVersion := extractAPIVersion(c.Path())

	req := &svc.Request{
		Method:       c.Method(),
		Action:       svc.HttpActions[c.Method()],
		APIfamily:    c.Params("apiFamily"),
		APIVersion:   apiVersion,
		ResourceName: c.Params("resourceName"),
		Body:         c.Body(),
		AccessToken:  jwtToken, // Store the raw JWT token
	}

	resp := h.service.CreateGenericObject(req)
	return SendResponse(c, resp)
}

// GetGenericObject retrieves a TMF object using generalized parameters.
func (h *Handler) GetGenericObject(c *fiber.Ctx) error {
	jwtToken := svc.ExtractJWTToken(c.Get("Authorization"))

	queryParams, _ := url.ParseQuery(string(c.Request().URI().QueryString()))
	idParam, _ := url.QueryUnescape(c.Params("id"))
	req := &svc.Request{
		Method:       c.Method(),
		Action:       svc.HttpActions[c.Method()],
		ResourceName: c.Params("resourceName"),
		ID:           idParam,
		QueryParams:  queryParams,
		AccessToken:  jwtToken, // Store the raw JWT token
	}

	resp := h.service.GetGenericObject(req)
	return SendResponse(c, resp)
}

// UpdateGenericObject updates an existing TMF object using generalized parameters.
func (h *Handler) UpdateGenericObject(c *fiber.Ctx) error {
	jwtToken := svc.ExtractJWTToken(c.Get("Authorization"))

	// Extract API version from the URL path
	apiVersion := extractAPIVersion(c.Path())

	idParam, _ := url.QueryUnescape(c.Params("id"))
	req := &svc.Request{
		Method:       c.Method(),
		Action:       svc.HttpActions[c.Method()],
		APIVersion:   apiVersion,
		ResourceName: c.Params("resourceName"),
		ID:           idParam,
		Body:         c.Body(),
		AccessToken:  jwtToken, // Store the raw JWT token
	}

	resp := h.service.UpdateGenericObject(req)
	return SendResponse(c, resp)
}

// DeleteGenericObject deletes a TMF object using generalized parameters.
func (h *Handler) DeleteGenericObject(c *fiber.Ctx) error {
	jwtToken := svc.ExtractJWTToken(c.Get("Authorization"))

	resourceName := c.Params("resourceName")
	if resourceName == "hub" {
		slog.Debug("handling hub creation")
		return h.DeleteHubSubscription(c)
	}

	// Extract API version from the URL path
	apiVersion := extractAPIVersion(c.Path())

	idParam, _ := url.QueryUnescape(c.Params("id"))
	req := &svc.Request{
		Method:       c.Method(),
		Action:       svc.HttpActions[c.Method()],
		APIVersion:   apiVersion,
		ResourceName: c.Params("resourceName"),
		ID:           idParam,
		AccessToken:  jwtToken, // Store the raw JWT token
	}

	resp := h.service.DeleteGenericObject(req)
	return SendResponse(c, resp)
}

// ListGenericObjects retrieves all TMF objects of a given type using generalized parameters.
func (h *Handler) ListGenericObjects(c *fiber.Ctx) error {
	jwtToken := svc.ExtractJWTToken(c.Get("Authorization"))

	// Extract API version from the URL path
	apiVersion := extractAPIVersion(c.Path())

	queryParams, _ := url.ParseQuery(string(c.Request().URI().QueryString()))
	req := &svc.Request{
		Method:       c.Method(),
		Action:       svc.HttpActions["LIST"],
		APIVersion:   apiVersion,
		ResourceName: c.Params("resourceName"),
		QueryParams:  queryParams,
		AccessToken:  jwtToken, // Store the raw JWT token
	}

	resp := h.service.ListGenericObjects(req)
	return SendResponse(c, resp)
}

func SendResponse(c *fiber.Ctx, resp *svc.Response) error {
	for key, value := range resp.Headers {
		c.Set(key, value)
	}
	if resp.Body != nil {
		return c.Status(resp.StatusCode).JSON(resp.Body)
	}
	return c.SendStatus(resp.StatusCode)
}
