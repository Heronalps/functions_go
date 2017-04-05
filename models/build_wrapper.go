package models

type BuildWrapper struct {
	Build *Build `json:"build"`
	Error *ErrorBody `json:"error,omitempty"`
}