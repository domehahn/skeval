package model

type EvalCategory string

const (
	CategoryFunctional  EvalCategory = "FUNCTIONAL"
	CategoryBoundary    EvalCategory = "BOUNDARY_REFUSAL"
	CategoryAdversarial EvalCategory = "ADVERSARIAL_INJECTION"
	CategoryValidation  EvalCategory = "PARAM_VALIDATION"
)

type TestCase struct {
	ID             string       `json:"id"`
	Category       EvalCategory `json:"category"`
	Prompt         string       `json:"prompt"`
	ExpectedAction string       `json:"expected_action"` // "INVOKE_TOOL", "REFUSE", "ERROR"
	ExpectedTool   string       `json:"expected_tool,omitempty"`
	ForbiddenTerms []string     `json:"forbidden_terms,omitempty"`
}

type EvalSuite struct {
	SkillName   string     `json:"skill_name"`
	SkillDigest string     `json:"skill_digest,omitempty"`
	Version     string     `json:"version"`
	TestCases   []TestCase `json:"test_cases"`
}

type TestCaseResult struct {
	TestCaseID string `json:"test_case_id"`
	Passed     bool   `json:"passed"`
	Reason     string `json:"reason,omitempty"`
}

type ScoreReport struct {
	TotalTests    int              `json:"total_tests"`
	PassedTests   int              `json:"passed_tests"`
	FailedTests   int              `json:"failed_tests"`
	OverallScore  float64          `json:"overall_score"` // 0.0 - 100.0%
	CategoryScore map[string]float64 `json:"category_score"`
	Results       []TestCaseResult `json:"results"`
}
