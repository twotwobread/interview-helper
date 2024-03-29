package health

import (
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v3"
	"github.com/stretchr/testify/assert"
	"github.com/twotwobread/interview-helper/api/router"
)

func TestHealthController(t *testing.T) {
	app := fiber.New()
	router.Setup(app)

	req, _ := app.Test(httptest.NewRequest("GET", "/health", nil))
	assert.Equal(t, fiber.StatusOK, req.StatusCode, "Health Expected status code 200")
}
