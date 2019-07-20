package tritonmedia

import (
	proto "github.com/tritonmedia/tritonmedia.go/pkg/proto"
)

// V1MediaObject is a media object returned by the API, it is
// the same format as the proto
type V1MediaObject proto.Media

// V1Metadata is the error object used in cases of error
type V1Metadata struct {
	// Success, did this op successfully occur or not
	Success bool `json:"success"`

	// ErrorMessage is an error message if set
	ErrorMessage string `json:"error_message"`

	// Host this was processed on
	Host string `json:"host"`
}

// V1 is the bare shared fields across all responses
type V1 struct {
	Metadata V1Metadata `json:"metadata"`
}

// V1Media is a common response that includes metadata and just
// a single media object
type V1Media struct {
	Metadata V1Metadata    `json:"metadata"`
	Media    V1MediaObject `json:"data"`
}

// V1MediaList is a common response that includes metadata and
// a list of media objects
type V1MediaList struct {
	Metadata V1Metadata      `json:"metadata"`
	Media    []V1MediaObject `json:"data"`
}

// V1Token is a token created response
type V1Token struct {
	Metadata V1Metadata `json:"metadata"`
	Token    string     `json:"data"`
}
