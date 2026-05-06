package policy

import "testing"

func TestDomainReviewLane(t *testing.T) {
	item := DomainReview{Signal: 59, Slack: 51, Drag: 18, Confidence: 85}
	if got := DomainReviewScore(item); got != 200 {
		t.Fatalf("score = %d", got)
	}
	if got := DomainReviewLane(item); got != "ship" {
		t.Fatalf("lane = %s", got)
	}
}
