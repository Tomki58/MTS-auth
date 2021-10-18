package middlewares

import (
	"MTS/auth/common"
	"net/http"

	"go.uber.org/zap"
)

func Logging(next http.Handler) http.Handler {
	// write your middleware definition here
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// request log
		// logger.WithFields(logrus.Fields{
		// 	"clientIp":      r.RemoteAddr,
		// 	"path":          r.URL.Path,
		// 	"method":        r.Method,
		// 	"contentLength": r.ContentLength,
		// }).Info("Incoming request")

		common.Logger.Info("Incoming request",
			zap.String("method", r.Method))

		next.ServeHTTP(w, r)

		// logger.WithFields(logrus.Fields{
		// 	"statusCode": w.Header().Get("statusCode"),
		// }).Info("Outgoing response")
		common.Logger.Info("Outgoing response")
	})
}
