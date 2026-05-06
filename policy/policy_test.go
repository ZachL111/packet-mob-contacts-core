package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 71, Capacity: 88, Latency: 15, Risk: 23, Weight: 13}, wantScore: 35, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 73, Capacity: 79, Latency: 27, Risk: 14, Weight: 8}, wantScore: 35, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 92, Capacity: 102, Latency: 22, Risk: 11, Weight: 7}, wantScore: 135, wantDecision: "review"},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := Score(tc.signal); got != tc.wantScore {
				t.Fatalf("score = %d, want %d", got, tc.wantScore)
			}
			if got := Classify(tc.signal); got != tc.wantDecision {
				t.Fatalf("decision = %s, want %s", got, tc.wantDecision)
			}
		})
	}
}
