package model

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"strconv"
	"testing"
)

func TestURLBadType(t *testing.T) {
	var u URL
	err := u.UnmarshalGQL(1234)
	require.Error(t, err)
}

func TestURLBadValue(t *testing.T) {
	var u URL
	err := u.UnmarshalGQL("not a url")
	require.Error(t, err)
}

func TestURL(t *testing.T) {
	var u URL
	err := u.UnmarshalGQL("https://github.com/darleet/blog-graphql")
	require.NoError(t, err)
	require.Equal(t, "https://github.com/darleet/blog-graphql", string(u))
}

func TestURLMarshal(t *testing.T) {
	u := URL("https://github.com/darleet/blog-graphql")
	var b bytes.Buffer
	u.MarshalGQL(&b)
	require.Equal(t, strconv.Quote("https://github.com/darleet/blog-graphql"), b.String())
}
