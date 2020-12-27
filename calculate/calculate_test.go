package calculate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetResults(t *testing.T) {
	results := GetResults(currentTestData.votes)

	// check if total votes match the expectations
	for _, result := range results {
		testDataForCandidate, ok := currentTestData.candidates[result.Candidate]
		if !assert.True(t, ok, "a candidate was returned by GetResults() that does not exist in test data: '%s'", result.Candidate) {
			return
		}
		if !assert.Equal(t, testDataForCandidate.getTotalVotes(), result.TotalVotes, "returned total votes for candidate '%s' did not match with expected test data", result.Candidate) {
			return
		}
	}

	// check if slice results are sorted correctly
	for k, result := range results {
		if k != len(results)-1 && !assert.GreaterOrEqual(t, result.TotalVotes, results[k+1].TotalVotes, "results are not sorted correctly") {
			return
		}
	}
}

func TestGetResultsByCountry(t *testing.T) {
	resultsByCountry := GetResultsByCountry(currentTestData.votes)
	expectations := currentTestData.getResultsByCountry()

	// check if total votes match the expectations
	for _, result := range resultsByCountry {
		expectedVotes, ok := expectations[result.Country]
		if !assert.True(t, ok, "a country was returned by GetResultsByCountry that does not exist in test data: '%s'", result.Country) {
			return
		}
		if !assert.Equal(t, expectedVotes, result.TotalVotes, "returned total votes for country '%s' did not match with expected test data", result.Country) {
			return
		}
	}

	// check if slice results are sorted correctly
	for k, result := range resultsByCountry {
		if k != len(resultsByCountry)-1 && !assert.GreaterOrEqual(t, result.TotalVotes, resultsByCountry[k+1].TotalVotes, "results are not sorted correctly") {
			return
		}
	}
}

func TestGetCandidatesByCountry(t *testing.T) {
	candidatesByCountry := GetCandidatesByCountry(currentTestData.votes)
	expectations := currentTestData.getCandidatesByCountry()

	// check if total votes per country match the expectations
	for _, candidatesForCountry := range candidatesByCountry {
		_, ok := expectations[candidatesForCountry.Country]
		if !assert.True(t, ok, "a country was returned by GetCandidatesByCountry that does not exist in test data: '%s'", candidatesForCountry.Country) {
			return
		}
		for _, result := range candidatesForCountry.Candidates {
			expectation, ok := expectations[candidatesForCountry.Country][result.Candidate]
			if !assert.True(t, ok, "a candidate was returned by GetCandidatesForCountry() for country '%s' that does not exist in test data: '%s'", candidatesForCountry.Country, result.Candidate) {
				return
			}
			if !assert.Equal(t, expectation, result.TotalVotes, "returned total votes for candidate '%s' in country '%s' did not match with expected test data", result.Candidate, candidatesForCountry.Country) {
				return
			}
		}
	}

	// check if slice results are sorted correctly
	for _, candidatesForCountry := range candidatesByCountry {
		for k, result := range candidatesForCountry.Candidates {
			if k != len(candidatesForCountry.Candidates)-1 && !assert.GreaterOrEqualf(t, result.TotalVotes, candidatesForCountry.Candidates[k+1].TotalVotes, "results are not sorted correctly for country '%s'", candidatesForCountry.Country) {
				return
			}
		}
	}
}
