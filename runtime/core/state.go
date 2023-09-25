package core

type State string

const (
	NONE          State = ""
	LOOKUP        State = "lookup"
	PASS          State = "pass"
	HASH          State = "hash"
	ERROR         State = "error"
	RESTART       State = "restart"
	DELIVER       State = "deliver"
	FETCH         State = "fetch"
	DELIVER_STALE State = "deliver_stale"
	LOG           State = "log"
)
