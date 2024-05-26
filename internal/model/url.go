package model

import (
	"fmt"
	"io"
	"net/url"
	"strconv"
)

type URL string

func (u *URL) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("urls must be strings")
	}
	_, err := url.ParseRequestURI(str)
	if err != nil {
		return fmt.Errorf("url is not valid: %w", err)
	}
	*u = URL(str)
	return nil
}

func (u *URL) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(string(*u)))
}
