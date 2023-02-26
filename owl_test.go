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
