package calculate

import (
	"github.com/htw-swa-jk-nk-ns/service-raw-data/vote"
	"github.com/rs/xid"
	"math/rand"
	"time"
)

type testData struct {
	votes      vote.Votes
	candidates map[string]testDataCandidate
}

func (t testData) getResultsByCountry() map[string]int {
	m := make(map[string]int)
	m["germany"] = 0
	m["france"] = 0
	m["england"] = 0
	for _, candidate := range t.candidates {
		m["germany"] += candidate.germanyVotes
		m["france"] += candidate.franceVotes
		m["england"] += candidate.englandVotes
	}
	return m
}

func (t testData) getCandidatesByCountry() map[string]map[string]int {
	m := map[string]map[string]int{
		"germany": make(map[string]int),
		"france":  make(map[string]int),
		"england": make(map[string]int),
	}
	for _, candidate := range t.candidates {
		for _, country := range []string{"germany", "france", "england"} {
			if _, ok := m[country][candidate.candidate]; !ok {
				m[country][candidate.candidate] = 0
			}
		}
		m["germany"][candidate.candidate] += candidate.germanyVotes
		m["france"][candidate.candidate] += candidate.franceVotes
		m["england"][candidate.candidate] += candidate.englandVotes
	}
	return m
}

type testDataCandidate struct {
	candidate    string
	germanyVotes int
	franceVotes  int
	englandVotes int
}

func (t testDataCandidate) getTotalVotes() int {
	return t.germanyVotes + t.franceVotes + t.englandVotes
}

var currentTestData = getDefaultTestData()

// helper functions to generate testData

// getDefaultTestData generates testData including 100 votes with predefined results
func getDefaultTestData() testData {
	var defaultTestData testData
	defaultTestData.candidates = make(map[string]testDataCandidate)

	var candidate testDataCandidate
	// Peter 25 votes (10 germany, 9 france, 6 england)
	candidate = testDataCandidate{
		candidate:    "Peter",
		germanyVotes: 10,
		franceVotes:  9,
		englandVotes: 6,
	}
	defaultTestData.candidates[candidate.candidate] = candidate
	defaultTestData.votes = append(defaultTestData.votes, generateVotesByTestDataCandidate(candidate)...)

	// Anna 20 votes (6 germany, 11 france, 3 england)
	candidate = testDataCandidate{
		candidate:    "Anna",
		germanyVotes: 6,
		franceVotes:  11,
		englandVotes: 3,
	}
	defaultTestData.candidates[candidate.candidate] = candidate
	defaultTestData.votes = append(defaultTestData.votes, generateVotesByTestDataCandidate(candidate)...)

	// Lisa 17 votes (7 germany, 0 france, 10 england)
	candidate = testDataCandidate{
		candidate:    "Lisa",
		germanyVotes: 7,
		franceVotes:  0,
		englandVotes: 10,
	}
	defaultTestData.candidates[candidate.candidate] = candidate
	defaultTestData.votes = append(defaultTestData.votes, generateVotesByTestDataCandidate(candidate)...)

	// Paul 16 votes (6 germany, 6 france, 4 england)
	candidate = testDataCandidate{
		candidate:    "Paul",
		germanyVotes: 6,
		franceVotes:  6,
		englandVotes: 4,
	}
	defaultTestData.candidates[candidate.candidate] = candidate
	defaultTestData.votes = append(defaultTestData.votes, generateVotesByTestDataCandidate(candidate)...)

	// Simon 12 votes (4 germany, 4 france, 4 england)
	candidate = testDataCandidate{
		candidate:    "Simon",
		germanyVotes: 4,
		franceVotes:  4,
		englandVotes: 4,
	}
	defaultTestData.candidates[candidate.candidate] = candidate
	defaultTestData.votes = append(defaultTestData.votes, generateVotesByTestDataCandidate(candidate)...)

	// Hans 10 votes (0 germany, 7 france, 3 england)
	candidate = testDataCandidate{
		candidate:    "Hans",
		germanyVotes: 0,
		franceVotes:  7,
		englandVotes: 3,
	}
	defaultTestData.candidates[candidate.candidate] = candidate
	defaultTestData.votes = append(defaultTestData.votes, generateVotesByTestDataCandidate(candidate)...)

	// randomize votes
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(defaultTestData.votes), func(i, j int) {
		defaultTestData.votes[i], defaultTestData.votes[j] = defaultTestData.votes[j], defaultTestData.votes[i]
	})

	return defaultTestData
}

func generateVotesByTestDataCandidate(candidate testDataCandidate) vote.Votes {
	var votes vote.Votes
	votes = append(votes, generateRandomVotes(candidate.candidate, "germany", candidate.germanyVotes)...)
	votes = append(votes, generateRandomVotes(candidate.candidate, "france", candidate.franceVotes)...)
	votes = append(votes, generateRandomVotes(candidate.candidate, "england", candidate.englandVotes)...)
	return votes
}

func generateRandomVotes(candidate string, country string, num int) vote.Votes {
	var votes vote.Votes
	for i := 0; i < num; i++ {
		votes = append(votes, generateRandomVote(candidate, country))
	}
	return votes
}

func generateRandomVote(candidate string, country string) vote.Vote {
	id := xid.New().String()
	if candidate == "" {
		candidate = id
	}
	if country == "" {
		country = id
	}
	return vote.Vote{
		ID:        id,
		Name:      id,
		Country:   country,
		Candidate: candidate,
		Date:      1,
	}
}
