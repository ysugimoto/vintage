package core

// Currently Edge runtime does not support WAF
// Can be set and get but always returns zero value or set value
type Waf struct {
	AnomalyScore         int64
	Blocked              bool
	Counter              int64
	Executed             bool
	Failures             int64
	HttpViolationScore   int64
	InboundAnomalyScore  int64
	LFIScore             int64
	Logdata              string
	Logged               bool
	Message              string
	Passed               bool
	PHPInjectionScore    int64
	RCEScore             int64
	RFIScore             int64
	RuleId               int64
	SessionFixationScore int64
	SQLInjectionScore    int64
	Severity             int64
	XSSScore             int64
}
