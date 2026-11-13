package generation

import (
	"io/fs"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Doc comments are Go source, not HTML: the generator must not escape backticks or quotes.
func TestGeneratedDocCommentsAreNotHTMLEscaped(t *testing.T) {
	entity := regexp.MustCompile(`&#x60;|&#39;|&quot;|&lt;|&gt;|&amp;`)
	var offenders []string
	err := filepath.WalkDir("../../src", func(path string, d fs.DirEntry, err error) error {
		if err != nil || d.IsDir() || !strings.HasSuffix(path, ".go") {
			return err
		}
		src, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		if entity.Match(src) {
			offenders = append(offenders, path)
		}
		return nil
	})
	require.NoError(t, err)
	assert.Empty(t, offenders)
}
