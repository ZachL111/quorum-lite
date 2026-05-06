package policy

import "testing"

func TestFixtureDecisions(t *testing.T) {
	tests := []struct {
		name         string
		signal       Signal
		wantScore    int
		wantDecision string
	}{
		{name: "case_1", signal: Signal{Demand: 69, Capacity: 107, Latency: 25, Risk: 8, Weight: 11}, wantScore: 155, wantDecision: "review"},
		{name: "case_2", signal: Signal{Demand: 85, Capacity: 94, Latency: 14, Risk: 20, Weight: 9}, wantScore: 122, wantDecision: "review"},
		{name: "case_3", signal: Signal{Demand: 75, Capacity: 83, Latency: 25, Risk: 20, Weight: 4}, wantScore: 17, wantDecision: "review"},
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
