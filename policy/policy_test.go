package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	signal := Signal{Demand: 71, Capacity: 88, Latency: 15, Risk: 23, Weight: 13}
	if got := Score(signal); got != 35 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 73, Capacity: 79, Latency: 27, Risk: 14, Weight: 8}
	if got := Score(signal); got != 35 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
	signal := Signal{Demand: 92, Capacity: 102, Latency: 22, Risk: 11, Weight: 7}
	if got := Score(signal); got != 135 { t.Fatalf("score = %d", got) }
	if got := Classify(signal); got != "review" { t.Fatalf("decision = %s", got) }
}
