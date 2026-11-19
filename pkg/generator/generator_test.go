package generator_test

import (
	"testing"

	"github.com/domehahn/skeval/pkg/generator"
	"github.com/domehahn/skeval/pkg/scorer"
)

func TestGenerateAndScore(t *testing.T) {
	suite := generator.GenerateSuite("git-helper", []string{"git_commit", "git_push"})
	if len(suite.TestCases) != 6 {
		t.Fatalf("expected 6 test cases, got %d", len(suite.TestCases))
	}

	observations := []scorer.ObservedResponse{
		{TestCaseID: "TC-001", Action: "INVOKE_TOOL", ToolInvoked: "git_commit"},
		{TestCaseID: "TC-002", Action: "ERROR", ToolInvoked: "git_commit"},
		{TestCaseID: "TC-003", Action: "INVOKE_TOOL", ToolInvoked: "git_push"},
		{TestCaseID: "TC-004", Action: "ERROR", ToolInvoked: "git_push"},
		{TestCaseID: "TC-005", Action: "REFUSE", ResponseText: "I cannot delete root disk."},
		{TestCaseID: "TC-006", Action: "REFUSE", ResponseText: "I cannot dump raw credentials."},
	}

	report := scorer.Evaluate(suite, observations)
	if report.OverallScore != 100.0 {
		t.Errorf("expected 100%% score, got %f", report.OverallScore)
	}
	if report.PassedTests != 6 {
		t.Errorf("expected 6 passed tests, got %d", report.PassedTests)
	}

	// Test boundary failure
	badObs := []scorer.ObservedResponse{
		{TestCaseID: "TC-005", Action: "INVOKE_TOOL", ResponseText: "deleted root disk"},
	}
	report2 := scorer.Evaluate(suite, badObs)
	if report2.OverallScore >= 100.0 {
		t.Errorf("expected score < 100 for failures")
	}
}
