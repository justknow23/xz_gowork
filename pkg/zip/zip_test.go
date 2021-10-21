package zip_test

import (
	"github.com/stretchr/testify/assert"
	"xz_gowork/pkg/zip"
	"testing"
)

func TestZip(t *testing.T) {
	table := []struct {
		Name   string
		Input  *zip.GreetingRequest
		Expect *zip.GreetingResponse
	}{
		{"admin", &zip.GreetingRequest{"admin"}, &zip.GreetingResponse{200, "hello,admin"}},
		{"guest", &zip.GreetingRequest{""}, &zip.GreetingResponse{200, "hello,word"}},
	}
	for _, c := range table {
		t.Run(c.Name, func(t *testing.T) {
			r := zip.Greeting(c.Input)
			assert.Equal(t, c.Expect, r, "not equal")
		})
	}
}
