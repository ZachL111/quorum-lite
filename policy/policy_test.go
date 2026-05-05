package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 69, Capacity: 107, Latency: 25, Risk: 8, Weight: 11}
	if got := Score(signal); got != 155 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 85, Capacity: 94, Latency: 14, Risk: 20, Weight: 9}
	if got := Score(signal); got != 122 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 75, Capacity: 83, Latency: 25, Risk: 20, Weight: 4}
	if got := Score(signal); got != 17 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
