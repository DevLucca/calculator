package main

import (
	"net/http"

	"github.com/apaxa-go/eval"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	e.HideBanner = true
	e.POST("/", calculate)
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Logger.Fatal(e.Start(":1323"))
}

func calculate(c echo.Context) (err error) {
	eq := &struct {
		Eq string `json:"equation"`
	}{}
	if err = c.Bind(&eq); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to parse equation 1")
	}
	expr, err := eval.ParseString(eq.Eq, "")
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to parse equation 2")
	}
	r, err := expr.EvalToInterface(nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, "Unable to convert Eval to Interface")
	}
	return c.JSON(http.StatusOK, struct {
		Result interface{} `json:"result"`
	}{
		Result: r,
	})
}
