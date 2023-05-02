package api
import (
"io"
"net/http"
"testing"

"github.com/gofiber/fiber/v2"
"github.com/zhenisduissekov/http-3rdparty-task-service/internal/config"
)


func Test_app(t *testing.T) {
	app := New(nil, &config.Conf{ServiceName: "test"})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	req, err := http.NewRequest(http.MethodGet, "/hello", nil)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := app.Test(req, -1)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status code %d but got %d", http.StatusOK, resp.StatusCode)
	}

	expected := "Hello, World!"
	actual, _ := io.ReadAll(resp.Body)
	if string(actual) != expected {
		t.Errorf("Expected response body %q but got %q", expected, actual)
	}
}
