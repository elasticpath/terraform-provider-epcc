package epcc

type ApiErrors interface {
	Error() string
	ListOfErrors() *ErrorList
	HttpStatusCode() uint16
	HttpMethod() string
	HttpPath() string
}

type ApiErrorResult struct {
	errorString    string
	apiErrors      *ErrorList
	httpMethod     string
	httpPath       string
	httpStatusCode uint16
}

func (a *ApiErrorResult) Error() string {
	return a.errorString
}

func (a *ApiErrorResult) ListOfErrors() *ErrorList {
	return a.apiErrors
}

func (a *ApiErrorResult) HttpStatusCode() uint16 {
	return a.httpStatusCode
}

func (a *ApiErrorResult) HttpMethod() string {
	return a.httpMethod
}

func (a *ApiErrorResult) HttpPath() string {
	return a.httpPath
}

func FromError(error error) (errors ApiErrors) {
	return &ApiErrorResult{
		errorString: error.Error(),
		apiErrors: &ErrorList{
			APIErrors: []APIError{},
		},
		httpStatusCode: 0,
		httpMethod:     "n/a",
		httpPath:       "n/a",
	}
}
