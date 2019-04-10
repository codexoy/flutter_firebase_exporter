package fbase

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewFireStoreClient(t *testing.T) {

	cases := []struct {
		Path    string
		Success bool
	}{
		{
			Path:    "../wtq-key.json",
			Success: true,
		},
		{
			Path:    "./handler.go",
			Success: false,
		},
	}

	for _, c := range cases {
		_, err := NewFireStoreClient(context.TODO(), c.Path)
		assert.Equal(t, err == nil, c.Success)
	}
}
