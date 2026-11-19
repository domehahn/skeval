package scorer

import (
	"strings"

	"github.com/domehahn/skeval/pkg/model"
)

type ObservedResponse struct {
	TestCaseID   string
	Action       string
	ToolInvoked  string
	ResponseText string
}

func Evaluate(suite *model.EvalSuite, observations []ObservedResponse) *model.ScoreReport {
	obsMap := make(map[string]ObservedResponse)
	for _, o := range observations {
		obsMap[o.TestCaseID] = o
	}

	report := &model.ScoreReport{
		TotalTests:    len(suite.TestCases),
		CategoryScore: make(map[string]float64),
		Results:       make([]model.TestCaseResult, 0),
	}

	catTotal := make(map[string]int)
	catPassed := make(map[string]int)

	for _, tc := range suite.TestCases {
		cat := string(tc.Category)
		catTotal[cat]++

		obs, found := obsMap[tc.ID]
		if !found {
			report.FailedTests++
			report.Results = append(report.Results, model.TestCaseResult{
				TestCaseID: tc.ID,
				Passed:     false,
				Reason:     "No observation recorded for test case",
			})
			continue
		}

		passed := true
		reason := "Pass"

		if tc.ExpectedAction != "" && tc.ExpectedAction != obs.Action {
			passed = false
			reason = "Action mismatch: expected " + tc.ExpectedAction + ", got " + obs.Action
		}

		if passed && tc.ExpectedTool != "" && tc.ExpectedTool != obs.ToolInvoked {
			passed = false
			reason = "Tool mismatch: expected " + tc.ExpectedTool + ", got " + obs.ToolInvoked
		}

		if passed && len(tc.ForbiddenTerms) > 0 {
			for _, term := range tc.ForbiddenTerms {
				if strings.Contains(strings.ToLower(obs.ResponseText), strings.ToLower(term)) {
					passed = false
					reason = "Forbidden term found in response: " + term
					break
				}
			}
		}

		if passed {
			report.PassedTests++
			catPassed[cat]++
		} else {
			report.FailedTests++
		}

		report.Results = append(report.Results, model.TestCaseResult{
			TestCaseID: tc.ID,
			Passed:     passed,
			Reason:     reason,
		})
	}

	if report.TotalTests > 0 {
		report.OverallScore = (float64(report.PassedTests) / float64(report.TotalTests)) * 100.0
	}

	for cat, total := range catTotal {
		if total > 0 {
			report.CategoryScore[cat] = (float64(catPassed[cat]) / float64(total)) * 100.0
		}
	}

	return report
}
