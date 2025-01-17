package customMiddleware

import (
	logger "github.com/kizmey/order_management_system/observability/logs"
	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
	"time"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {

		//if c.Path() == "/metricsx" {
		//	return nil
		//}

		start := time.Now()

		fields := logrus.Fields{
			"method":   c.Request().Method,
			"path":     c.Path(),
			"query":    c.QueryString(),
			"remoteIP": c.RealIP(),
		}

		// Log HTTP request
		//logger.LogInfo("HTTP request", fields)

		if err := next(c); err != nil {
			c.Error(err)
		}

		// Log HTTP response
		fields["status"] = c.Response().Status
		fields["latency"] = time.Since(start).Seconds()

		logger.LogInfo("HTTP response", fields)
		return nil
	}
}
