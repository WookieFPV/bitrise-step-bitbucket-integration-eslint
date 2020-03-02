package main

//Message type
type Message struct {
	RuleID    string `json:"ruleId"`
	Severity  int    `json:"severity"`
	Message   string `json:"message"`
	Line      int    `json:"line"`
	Column    int    `json:"column"`
	NodeType  string `json:"nodeType"`
	Source    string `json:"source"`
	MessageID string `json:"messageId"`
	EndLine   int    `json:"endLine"`
	EndColumn int    `json:"endColumn"`
}

// Issues Eslint error/warning element
type Issues struct {
	FilePath            string    `json:"filePath"`
	Messages            []Message `json:"messages"`
	ErrorCount          int       `json:"errorCount"`
	WarningCount        int       `json:"warningCount"`
	FixableErrorCount   int       `json:"fixableErrorCount"`
	FixableWarningCount int       `json:"fixableWarningCount"`
}

// Annotation which will be used to send to bitbucket server
type Annotation struct {
	Path     string `json:"path"`
	Line     int    `json:"line"`
	Message  string `json:"message"`
	Severity string `json:"severity"` // one of: "LOW"/"MEDIUM"/"HIGH"
}

//BitbucketAnnotations type which needs to get posted to bitbucket
type BitbucketAnnotations struct {
	Annotations []Annotation `json:"annotations"`
}

// EslintReportData which contains the summary for the eslint report
type EslintReportData struct {
	Title string `json:"title"`
	Value int    `json:"value"`
}

// Report which will be used to send to bitbucket server
type Report struct {
	Title   string             `json:"title"`
	Vendor  string             `json:"vendor"`
	LogoURL string             `json:"logoUrl"`
	Data    []EslintReportData `json:"data"`
	Result  string             `json:"result"` // "FAIL" or "PASS"
	Link    string             `json:"link"`
	Details string             `json:"details"`
}
