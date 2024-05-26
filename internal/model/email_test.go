package model

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestEmailUnmarshalBadType(t *testing.T) {
	var e Email
	err := e.UnmarshalGQL(1234)
	require.Error(t, err)
}

func TestEmailUnmarshalBadValue(t *testing.T) {
	var e Email
	err := e.UnmarshalGQL("not an email")
	require.Error(t, err)
}

func TestEmailUnmarshal(t *testing.T) {
	var e Email
	err := e.UnmarshalGQL("bSbQp@example.com")
	require.NoError(t, err)
	require.Equal(t, "bSbQp@example.com", string(e))
}

func TestEmailMarshal(t *testing.T) {
	e := Email("bSbQp@example.com")
	var b bytes.Buffer
	e.MarshalGQL(&b)
	require.Equal(t, strconv.Quote("bSbQp@example.com"), b.String())
}
