package api

import (
	"github.com/htw-swa-jk-nk-ns/service-calculate/calculate"
	"github.com/htw-swa-jk-nk-ns/service-raw-data/vote"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"net/http"
)

// StartAPI starts the API.
func StartAPI() {
	e := echo.New()

	e.GET("/results", getResults)

	e.GET("/resultsByCountry", getResultsByCountry)

	e.GET("/candidatesByCountry", getCandidatesByCountry)

	e.GET("/getTop5Candidates", getTop5Candidates)

	e.GET("/getTop5Countries", getTop5Countries)

	if viper.GetString("api.certfile") != "" && viper.GetString("api.keyfile") != "" {
		e.Logger.Fatal(e.StartTLS(":"+viper.GetString("api.port"), viper.GetString("api.certfile"), viper.GetString("api.keyfile")))
	} else {
		e.Logger.Fatal(e.Start(":" + viper.GetString("api.port")))
	}
}

func getResults(ctx echo.Context) error {
	var votes vote.Votes
	if err := ctx.Bind(&votes); err != nil {
		return getApiResoponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to bind input")))
	}
	return getApiResoponse(ctx, http.StatusOK, calculate.GetResults(votes))
}

func getResultsByCountry(ctx echo.Context) error {
	var votes vote.Votes
	if err := ctx.Bind(&votes); err != nil {
		return getApiResoponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to bind input")))
	}
	return getApiResoponse(ctx, http.StatusOK, calculate.GetResultsByCountry(votes))
}

func getCandidatesByCountry(ctx echo.Context) error {
	var votes vote.Votes
	if err := ctx.Bind(&votes); err != nil {
		return getApiResoponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to bind input")))
	}
	return getApiResoponse(ctx, http.StatusOK, calculate.GetCandidatesByCountry(votes))
}

func getTop5Candidates(ctx echo.Context) error {
	var votes vote.Votes
	if err := ctx.Bind(&votes); err != nil {
		return getApiResoponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to bind input")))
	}
	return getApiResoponse(ctx, http.StatusOK, calculate.GetTop5Candidates(votes))
}

func getTop5Countries(ctx echo.Context) error {
	var votes vote.Votes
	if err := ctx.Bind(&votes); err != nil {
		return getApiResoponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to bind input")))
	}
	return getApiResoponse(ctx, http.StatusOK, calculate.GetTop5Countries(votes))
}

func getApiResoponse(ctx echo.Context, statusCode int, response interface{}) error {
	switch format := viper.GetString("api.format"); format {
	case "json":
		return ctx.JSON(statusCode, response)
	case "xml":
		return ctx.XML(statusCode, response)
	default:
		return ctx.String(http.StatusInternalServerError, "invalid output format '"+format+"'")
	}
}

type OutputError struct {
	message string
}

func newOutputError(err error) OutputError {
	return OutputError{
		message: err.Error(),
	}
}
