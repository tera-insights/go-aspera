package aspera

import "errors"

// Error represents an error response from the aspera API.
type Error struct {
	Code            int
	InternalMessage string
	UserMessage     string
}

// SDK ERROR CODES
var ErrEndpointNotFound = errors.New("endpoint not found")
