package api

import (
	"encoding/json"
	"github.com/htw-swa-jk-nk-ns/service-calculate/calculate"
	"github.com/htw-swa-jk-nk-ns/service-raw-data/vote"
	"github.com/labstack/echo"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
	"io/ioutil"
	"net/http"
)

// StartAPI starts the API.
func StartAPI() {
	e := echo.New()

	e.GET("/results", getResults)

	e.GET("/votesByCountry", getVotesByCountry)

	e.GET("/getResultsByCountry", getResultsByCountry)

	e.GET("/top5Candidates", getTop5Candidates)

	e.GET("/top5Countries", getTop5Countries)

	if viper.GetString("api.certfile") != "" && viper.GetString("api.keyfile") != "" {
		e.Logger.Fatal(e.StartTLS(":"+viper.GetString("api.port"), viper.GetString("api.certfile"), viper.GetString("api.keyfile")))
	} else {
		e.Logger.Fatal(e.Start(":" + viper.GetString("api.port")))
	}
}

func getResults(ctx echo.Context) error {
	votes, err := getAllVotes()
	if err != nil {
		return getApiResponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to get all votes")))
	}
	return getApiResponse(ctx, http.StatusOK, calculate.GetResults(votes))
}

func getVotesByCountry(ctx echo.Context) error {
	votes, err := getAllVotes()
	if err != nil {
		return getApiResponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to get all votes")))
	}
	return getApiResponse(ctx, http.StatusOK, calculate.GetVotesByCountry(votes))
}

func getResultsByCountry(ctx echo.Context) error {
	votes, err := getAllVotes()
	if err != nil {
		return getApiResponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to get all votes")))
	}
	return getApiResponse(ctx, http.StatusOK, calculate.GetCandidatesByCountry(votes))
}

func getTop5Candidates(ctx echo.Context) error {
	votes, err := getAllVotes()
	if err != nil {
		return getApiResponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to get all votes")))
	}
	return getApiResponse(ctx, http.StatusOK, calculate.GetTop5Candidates(votes))
}

func getTop5Countries(ctx echo.Context) error {
	votes, err := getAllVotes()
	if err != nil {
		return getApiResponse(ctx, http.StatusBadRequest, newOutputError(errors.Wrap(err, "failed to get all votes")))
	}
	return getApiResponse(ctx, http.StatusOK, calculate.GetTop5Countries(votes))
}

func getAllVotes() (vote.Votes, error) {
	res, err := http.Get(viper.GetString("votes.api"))
	if err != nil {
		return nil, errors.Wrap(err, "http get failed")
	}
	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read body")
	}

	var v vote.Votes
	err = json.Unmarshal(body, &v)
	if err != nil {
		log.Error().Err(err).Msgf("failed to unmarshal json to vote.Votes, json: '%s'", string(body))
		return nil, errors.Wrapf(err, "failed to unmarshal json to vote.Votes, json: '%s'", string(body))
	}
	return v, nil
}

func getApiResponse(ctx echo.Context, statusCode int, response interface{}) error {
	if statusCode == http.StatusOK && response == nil {
		response = []struct{}{}
	}
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
	Message string `json:"message" xml:"message"`
}

func newOutputError(err error) OutputError {
	return OutputError{
		Message: err.Error(),
	}
}
