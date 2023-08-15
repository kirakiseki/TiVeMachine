package setup

import (
	"context"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/utils"
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
	zerologLogger := zerolog.New(output).With().Timestamp().Logger()
	zerologLogger.Info().Msg("Initialized zerolog")
	return &zerologLogger
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

// GORM Logger

type GORMLogger struct {
	logLevel logger.LogLevel
}

func (l *GORMLogger) Info(ctx context.Context, s string, i ...interface{}) {
	if l.logLevel >= logger.Info {
		Inst.Logger.Info().Msg(fmt.Sprintf(s, i...))
	}
}

func (l *GORMLogger) Warn(ctx context.Context, s string, i ...interface{}) {
	if l.logLevel >= logger.Warn {
		Inst.Logger.Warn().Msg(fmt.Sprintf(s, i...))
	}
}

func (l *GORMLogger) Error(ctx context.Context, s string, i ...interface{}) {
	if l.logLevel >= logger.Error {
		Inst.Logger.Error().Msg(fmt.Sprintf(s, i...))
	}
}

func (l *GORMLogger) Trace(ctx context.Context, begin time.Time, fc func() (sql string, rowsAffected int64), err error) {
	if l.logLevel <= logger.Silent {
		return
	}

	elapsed := time.Since(begin)
	switch {
	case err != nil && l.logLevel >= logger.Error && (!errors.Is(err, logger.ErrRecordNotFound)):
		sql, rows := fc()
		if rows == -1 {
			Inst.Logger.Error().
				Err(err).
				Dur("elapsed", elapsed).
				Str("sql", sql).
				Str("rows", "-").
				Str("caller", utils.FileWithLineNum()).
				Msg("GORM")
		} else {
			Inst.Logger.Error().
				Err(err).
				Dur("elapsed", elapsed).
				Str("sql", sql).
				Int64("rows", rows).
				Str("caller", utils.FileWithLineNum()).
				Msg("GORM")
		}
	case l.logLevel >= logger.Warn:
		sql, rows := fc()
		if rows == -1 {
			Inst.Logger.Warn().
				Dur("elapsed", elapsed).
				Str("sql", sql).
				Str("rows", "-").
				Str("caller", utils.FileWithLineNum()).
				Msg("GORM")
		} else {
			Inst.Logger.Warn().
				Dur("elapsed", elapsed).
				Str("sql", sql).
				Int64("rows", rows).
				Str("caller", utils.FileWithLineNum()).
				Msg("GORM")
		}
	case l.logLevel == logger.Info:
		sql, rows := fc()
		if rows == -1 {
			Inst.Logger.Info().
				Dur("elapsed", elapsed).
				Str("sql", sql).
				Str("rows", "-").
				Str("caller", utils.FileWithLineNum()).
				Msg("GORM")
		} else {
			Inst.Logger.Info().
				Dur("elapsed", elapsed).
				Str("sql", sql).
				Int64("rows", rows).
				Str("caller", utils.FileWithLineNum()).
				Msg("GORM")
		}
	}
}

func (l *GORMLogger) LogMode(level logger.LogLevel) logger.Interface {
	newLogger := *l
	newLogger.logLevel = level
	return &newLogger
}
