package routes

import (
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gofiber/fiber/v2"
	"github.com/jayanthkrishna/url-shortening-service/helpers"
)

type request struct {
	URL         string        `json:"url"`
	CustomShort string        `json:"short"`
	Expiry      time.Duration `json:"expiry"`
}

type response struct {
	URL             string        `json:"url"`
	CustomShort     string        `json:"short"`
	Expiry          time.Duration `json:"expiry"`
	XRateRemaining  int           `json:"rate_limit"`
	XRateLimitReset time.Duration `json:"rate_limit_reset"`
}

func ShortenURL(c *fiber.Ctx) error {

	body := new(request)

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "cannot parse JSON",
		})
	}

	//implement rate limiting

	//check if the url is an actual URL

	if !govalidator.IsURL(body.URL) {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Invalid URL",
		})
	}

	//check for domain error

	if !helpers.RemoveDomainError(body.URL) {
		return c.Status(fiber.StatusServiceUnavailable).JSON(fiber.Map{
			"error": "U cant access",
		})
	}

	//enforce https

	body.URL = helpers.EnforceHttps(body.URL)
	return nil
}
