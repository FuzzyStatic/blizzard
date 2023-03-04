package blizzard

import (
	"context"
	"fmt"
	"testing"
)

func TestOWLSummaryData(t *testing.T) {
	dat, _, err := usClient.OWLSummaryData(context.Background())
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestOWLPlayersAPI(t *testing.T) {
	dat, _, err := usClient.OWLPlayersAPI(context.Background(), 15401)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestOWLMatchesAPI(t *testing.T) {
	dat, _, err := usClient.OWLMatchesAPI(context.Background(), 38971)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestOWLSegmentsAPI(t *testing.T) {
	dat, _, err := usClient.OWLSegmentsAPI(context.Background(), "owl2-2022-kickoff-clash-qualifiers")
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}

func TestOWLTeamsAPI(t *testing.T) {
	dat, _, err := usClient.OWLTeamsAPI(context.Background(), 15170)
	if err != nil {
		fmt.Println(err)
		t.Fail()
	}

	if printOutput != "" {
		fmt.Printf("%+v\n", dat)
	}
}
