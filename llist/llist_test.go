package llist

import (
	"fmt"
	"strings"
	"testing"
)

func TestLList(t *testing.T) {
	text := `The approach will not be easy. You are required to maneuver
      straight down this trench and skim the surface to this point. The
      target area is only two meters wide. It’s a small thermal exhaust
      port, right below the main port. The shaft leads directly to the
      reactor system. A precise hit will start a chain reaction which should
      destroy the station. Only a precise hit will set up a chain reaction.
      The shaft is ray-shielded, so you’ll have to use proton torpedoes.
      That’s impossible, even for a computer. It’s not impossible. I used
      to bull’s-eye womp rats in my T-sixteen back home. They’re not much
      bigger than two meters. Man your ships! And may the Force be with you!`

	words := strings.Fields(text)
	// convert to []interface{}
	wordz := make([]interface{}, len(words))
	for i := range wordz {
		wordz[i] = words[i]
	}

	ll := LListCreateWith(wordz)
	fmt.Println(ll)

	if ll.First().(string) != "The" {
		t.Error(`"The" should be first`)
	}

	upper := func(strIn interface{}) interface{} {
		return strings.ToUpper(strIn.(string))
	}

	llCaps := ll.Map(upper)
	fmt.Println(llCaps)

	if llCaps.First().(string) != "THE" {
		t.Error(`"The" should be first`)
	}
}
