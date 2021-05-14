package epcc

import (
	"encoding/json"
	"fmt"
	"strings"
)

type APIError struct {
	Code   string
	Status string
	Title  string
	Detail string
}

func (e *APIError) UnmarshalJSON(data []byte) error {
	var objmap map[string]json.RawMessage
	err := json.Unmarshal(data, &objmap)

	if err != nil {
		return err
	}

	if val, ok := objmap["code"]; ok {
		e.Code = fmt.Sprintf("%s", val)
	}

	if val, ok := objmap["status"]; ok {
		e.Status = fmt.Sprintf("%s", val)
	}

	if val, ok := objmap["title"]; ok {
		e.Title = fmt.Sprintf("%s", val)
	}

	if val, ok := objmap["detail"]; ok {
		e.Detail = fmt.Sprintf("%s", val)
	}

	return nil
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
