package embed

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseEmbed(t *testing.T) {
	cfg, err := Parse([]string{"tools/please_go/embed/test_data/test.go"})
	require.NoError(t, err)
	assert.Equal(t, map[string][]string{
		"hello.txt":   {"hello.txt"},
		"files/*.txt": {"files/test.txt"},
		"files":       {"files/test.txt"},
	}, cfg.Patterns)
	assert.Equal(t, map[string]string{
		"hello.txt":      "tools/please_go/embed/test_data/hello.txt",
		"files/test.txt": "tools/please_go/embed/test_data/files/test.txt",
	}, cfg.Files)
}
