package calculate

import (
	"github.com/htw-swa-jk-nk-ns/service-raw-data/vote"
	"sort"
)

type Result struct {
	Candidate  string `json:"candidate" xml:"candidate"`
	TotalVotes int    `json:"totalVotes" xml:"totalVotes"`
}

type VotesByCountry struct {
	Country    string `json:"country" xml:"country"`
	TotalVotes int    `json:"totalVotes" xml:"totalVotes"`
}

type CandidatesByCountry struct {
	Country    string   `json:"country" xml:"country"`
	Candidates []Result `json:"candidates" xml:"candidates"`
}

type valueGetStringFunc func(vote *vote.Vote) string

func GetResults(v vote.Votes) []Result {
	var results []Result
	m := getNumOfXGroupedByY(v, GetCandidateForVote, nil)
	for measurement, totalVotes := range m[""] {
		results = append(results, Result{
			Candidate:  measurement,
			TotalVotes: totalVotes,
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].TotalVotes > results[j].TotalVotes
	})
	return results
}

func GetVotesByCountry(v vote.Votes) []VotesByCountry {
	var results []VotesByCountry
	m := getNumOfXGroupedByY(v, GetCountryForVote, nil)
	for measurement, totalVotes := range m[""] {
		results = append(results, VotesByCountry{
			Country:    measurement,
			TotalVotes: totalVotes,
		})
	}
	sort.Slice(results, func(i, j int) bool {
		return results[i].TotalVotes > results[j].TotalVotes
	})
	return results
}

func GetCandidatesByCountry(v vote.Votes) []CandidatesByCountry {
	var candidatesByCountry []CandidatesByCountry
	m := getNumOfXGroupedByY(v, GetCandidateForVote, GetCountryForVote)
	for country, results := range m {
		var countryResult CandidatesByCountry
		countryResult.Country = country
		for candidate, totalVotes := range results {
			countryResult.Candidates = append(countryResult.Candidates, Result{
				Candidate:  candidate,
				TotalVotes: totalVotes,
			})
		}
		sort.Slice(countryResult.Candidates, func(i, j int) bool {
			return countryResult.Candidates[i].TotalVotes > countryResult.Candidates[j].TotalVotes
		})
		candidatesByCountry = append(candidatesByCountry, countryResult)
	}
	return candidatesByCountry
}

func GetTop5Candidates(v vote.Votes) []Result {
	return GetResults(v)[:5]
}

func GetTop5Countries(v vote.Votes) []VotesByCountry {
	return GetVotesByCountry(v)[:5]
}

func getNumOfXGroupedByY(v vote.Votes, getFunc valueGetStringFunc, groupedByFunc valueGetStringFunc) map[string]map[string]int {
	results := make(map[string]map[string]int)
	if groupedByFunc == nil {
		groupedByFunc = func(vote *vote.Vote) string {
			return ""
		}
	}

	for _, vt := range v {
		filter := groupedByFunc(&vt)
		get := getFunc(&vt)
		if _, ok := results[filter]; !ok {
			results[filter] = make(map[string]int)
		}
		if _, ok := results[filter][get]; !ok {
			results[filter][get] = 0
		}
		results[filter][get]++
	}
	return results
}

func GetCandidateForVote(v *vote.Vote) string {
	return v.Candidate
}

func GetCountryForVote(v *vote.Vote) string {
	return v.Country
}
