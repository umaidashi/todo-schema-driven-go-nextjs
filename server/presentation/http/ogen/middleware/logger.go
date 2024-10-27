package middlewares

import (
	"fmt"
	"server/infra/applogger"

	"github.com/ogen-go/ogen/middleware"
)

func Logging() middleware.Middleware {
	return func(
		req middleware.Request,
		next func(req middleware.Request) (middleware.Response, error),
	) (middleware.Response, error) {
		applogger.Info(fmt.Sprintf("%s %s", req.Raw.Method, req.Raw.URL))

		res, err := next(req)
		if err != nil {
			applogger.Error(fmt.Sprintf("Error %v", err))
		} else {
			applogger.Info(fmt.Sprintf("Response %T", res.Type))
		}

		return res, err
	}
}
