package htable

import (
	"fmt"
	"strings"
	"testing"
)

func TestHTable(t *testing.T) {
	table := New()
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

	for _, w := range words {
		val := table.Get(w)
		if val == nil {
			table = table.Set(w, 1)
		} else {
			table = table.Set(w, val.(int)+1)
		}
	}

	fmt.Println(table)

	if table.Get("approach").(int) != 1 {
		t.Error(`"approach" should appear once`)
	}
	if table.Get("The").(int) != 4 {
		t.Error(`"The" should appear 4 times`)
	}
	//HTPrint(table)
}
