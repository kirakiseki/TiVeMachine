package setup

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"net/http"
	"os"
	"time"
)

const (
	green   = "\033[97;42m"
	white   = "\033[90;47m"
	yellow  = "\033[90;43m"
	red     = "\033[97;41m"
	blue    = "\033[97;44m"
	magenta = "\033[97;45m"
	cyan    = "\033[97;46m"
	reset   = "\033[0m"
)

func Zerolog() *zerolog.Logger {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339}
	logger := zerolog.New(output).With().Timestamp().Logger()
	logger.Info().Msg("Initialized zerolog")
	return &logger
}

func GinLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		end := time.Now()
		latency := end.Sub(start)
		path := c.Request.URL.Path
		method := c.Request.Method
		statusCode := c.Writer.Status()

		var statusColor, methodColor string

		switch {
		case statusCode >= http.StatusOK && statusCode < http.StatusMultipleChoices:
			statusColor = green
		case statusCode >= http.StatusMultipleChoices && statusCode < http.StatusBadRequest:
			statusColor = white
		case statusCode >= http.StatusBadRequest && statusCode < http.StatusInternalServerError:
			statusColor = yellow
		default:
			statusColor = red
		}

		switch method {
		case http.MethodGet:
			methodColor = blue
		case http.MethodPost:
			methodColor = cyan
		default:
			methodColor = reset
		}

		Inst.Logger.Info().
			Str("method", method).
			Str("path", path).
			Int("status_code", statusCode).
			Dur("latency", latency).
			Msg(fmt.Sprintf("%s %3d %s%s %-4s %s", statusColor, statusCode, reset, methodColor, method, reset))
	}
}
