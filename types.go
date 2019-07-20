package tritonmedia

import (
	proto "github.com/tritonmedia/tritonmedia.go/pkg/proto"
)

// V1Media is a media object returned by the API
type V1Media proto.Media

// V1Error is the error object used in cases of error
type V1Error struct {
	// Success is the status of this request, did it succeed?
	Success bool `json:"success"`

	// Message is the error message if set
	Message string `json:"message"`
}
