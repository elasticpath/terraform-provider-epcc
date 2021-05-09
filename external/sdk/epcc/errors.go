package epcc

import (
	"strings"
)

type APIError struct {
	Code   string
	Status string
	Title  string
	Detail string
}

type ErrorList struct {
	APIErrors []APIError `json:"errors"`
}

func (a *APIError) String() string {
	var sb strings.Builder

	if len(a.Title) > 0 {
		sb.WriteString("Title: ")
		sb.WriteString(a.Title)
		sb.WriteString("\n")
	}

	if len(a.Status) > 0 {
		sb.WriteString("Status: ")
		sb.WriteString(a.Status)
		sb.WriteString("\n")
	}

	if len(a.Code) > 0 {
		sb.WriteString("Code: ")
		sb.WriteString(a.Code)
		sb.WriteString("\n")
	}

	if len(a.Detail) > 0 {
		sb.WriteString("Detail: ")
		sb.WriteString(a.Detail)
		sb.WriteString("\n")
	}

	return sb.String()
}

func (e *ErrorList) String() string {
	var sb strings.Builder

	for _, apiError := range e.APIErrors {
		sb.WriteString(apiError.String())
	}

	return sb.String()
}
