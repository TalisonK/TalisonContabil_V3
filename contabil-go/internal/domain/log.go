package domain

type LogLine struct {
}

type LogRequest struct {
	Start string `json:"start"`
	Lines int64  `json:"lines"`
}
