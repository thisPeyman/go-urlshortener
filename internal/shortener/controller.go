package shortener

import (
	"net/http"

	sentryecho "github.com/getsentry/sentry-go/echo"
	"github.com/labstack/echo/v4"
	"github.com/thisPeyman/go-urlshortner/api"
)

type ShortenerController struct {
	shortenerService *ShortenerService
}

type ShortenUrlRequest struct {
	LongUrl string `json:"longUrl" validate:"required,url"`
}

func RegisterController(e *echo.Echo, shortenerService *ShortenerService) {
	controller := &ShortenerController{
		shortenerService: shortenerService,
	}

	e.POST("/shorten", controller.ShortenUrl)
	e.GET("/:shortUrl", controller.ExpandUrl)
}

func (c *ShortenerController) ShortenUrl(e echo.Context) error {

	var data ShortenUrlRequest
	if err := e.Bind(&data); err != nil {
		return err
	}
	if err := e.Validate(data); err != nil {
		return err
	}

	res, err := c.shortenerService.ShortenUrl(e.Request().Context(), &api.ShortenURLRequest{LongUrl: data.LongUrl})
	if err != nil {
		if hub := sentryecho.GetHubFromContext(e); hub != nil {
			hub.CaptureException(err)
		}
		return err
	}

	return e.JSON(http.StatusCreated, res)
}

func (c *ShortenerController) ExpandUrl(e echo.Context) error {
	shortUrl := e.Param("shortUrl")

	res, err := c.shortenerService.ExpandURL(e.Request().Context(), &api.ExpandURLRequest{ShortUrl: shortUrl})
	if err != nil {
		return err
	}

	return e.Redirect(http.StatusFound, res.LongUrl)
}
