package cerrors

import (
	"encoding/json"
	"errors"
	"log"
)

type Severity int

const (
	SeverityCritical = Severity(40)
	SeverityError    = Severity(30)
	SeverityWarning  = Severity(20)
	SeverityInfo     = Severity(10)
)

const DefaultSeverity = SeverityInfo

type Error struct {
	Code           string            `json:"code" bson:"code"`
	Message        string            `json:"message" bson:"message"`
	Parent         error             `json:"parent" bson:"parent"`
	Severity       Severity          `json:"severity" bson:"severity"`
	Classification string            `json:"classification" bson:"classification"`
	Meta           map[string]string `json:"meta" bson:"meta"`
}

func (e *Error) SeverityValue() Severity {
	if e.Severity == 0 {
		return DefaultSeverity
	}

	return e.Severity
}

func (e Error) Error() string {
	b, err := json.MarshalIndent(e, "<prefix>", "<indent>")
	if err != nil {
		log.Fatal(err)
	}

	return string(b)
}

func (e Error) clone() Error {
	return Error{
		Code:           e.Code,
		Message:        e.Message,
		Parent:         e.Parent,
		Severity:       e.Severity,
		Classification: e.Classification,
		Meta:           e.Meta,
	}
}

func (e Error) Cause(err error) Error {
	cloned := e.clone()
	cloned.Parent = err

	return cloned
}

func (e Error) WithMeta(meta map[string]string) Error {
	cloned := e.clone()
	cloned.Meta = meta

	return cloned
}

func FilterErrorsBySeverity(errorList []error, severity Severity) []error {
	var filteredErrors []error

	for _, err := range errorList {
		v := Error{}
		if errors.As(err, &v) {
			if v.SeverityValue() >= severity {
				filteredErrors = append(filteredErrors, err)
			}
		} else {
			filteredErrors = append(filteredErrors, err)
		}
	}

	return filteredErrors
}
