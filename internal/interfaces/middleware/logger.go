package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/labstack/echo/v4"
	"go.uber.org/zap"
)

func Logger(logger *zap.Logger) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			start := time.Now()
			req := c.Request()
			res := c.Response()

			bodyBytes, _ := io.ReadAll(req.Body)
			req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

			logger.Info("Request started",
				zap.String("method", req.Method),
				zap.String("url", req.URL.String()),
				zap.ByteString("body", bodyBytes),
			)

			err := next(c)

			logger.Info("Request completed",
				zap.Int("status", res.Status),
				zap.Duration("duration", time.Since(start)),
			)

			if err != nil {
				c.Error(err)
				logger.Error("Request error",
					zap.Error(err),
				)
			}

			return err
		}
	}
}
