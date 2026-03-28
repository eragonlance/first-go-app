package main

import (
	"io"
	"net/http"
	"testing"

	"github.com/go-jose/go-jose/v4/testutils/assert"
)

func TestGetVersion(t *testing.T) {
	tests := []struct {
		deployed, ref, sha string
		expected           string
	}{
		{"1", "refs/heads/main", "abc123", "refs/heads/main\nabc123"},
		{"", "", "", "Applicable for deployed only."},
	}

	app := setup()
	for _, test := range tests {
		t.Setenv("GIT_REF", test.ref)
		t.Setenv("GIT_SHA", test.sha)
		t.Setenv("DEPLOYED", test.deployed)

		req, _ := http.NewRequest("GET", "/version", nil)
		res, _ := app.Test(req)
		body, _ := io.ReadAll(res.Body)

		assert.Equal(t, string(body), test.expected)
	}
}
