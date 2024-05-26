package model

import (
	"fmt"
	"io"
	"net/mail"
	"strconv"
)

type Email string

func (e *Email) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("emails must be strings")
	}
	_, err := mail.ParseAddress(str)
	if err != nil {
		return fmt.Errorf("email is not valid: %w", err)
	}
	*e = Email(str)
	return nil
}

func (e Email) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(string(e)))
}
