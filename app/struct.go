package app

// ErrorMessage Struct Response
type ErrorMessage struct {
	Code    int         `json:"code"`
	Source  SourceError `json:"source"`
	Title   string      `json:"title,omitempty"`
	Details string      `json:"details"`
}

// SourceError Struct Response
// Part from ErrorMessage
type SourceError struct {
	Pointer string `json:"pointer,omitempty"`
}

// Response Struct Response
type Response struct {
	Data     Data     `json:"data"`
	Included Included `json:"included"`
}

// Data Struct Response
// Part from Response
type Data struct {
	ID            int         `json:"id"`
	Type          string      `json:"type"`
	Attributes    interface{} `json:"attributes"`
	Relationships interface{} `json:"relationships"`
}

// Included Struct Response
// Part from Response
type Included struct {
	ID         int         `json:"id"`
	Type       string      `json:"type"`
	Attributes interface{} `json:"attributes"`
}
