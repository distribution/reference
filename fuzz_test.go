package reference

import (
	"testing"
)

// fuzzParseNormalizedNamed implements a fuzzer
// that targets ParseNormalizedNamed
// nolint:deadcode
func FuzzParseNormalizedNamed(f *testing.F) {
	f.Fuzz(func(_ *testing.T, data string) {
		_, _ = ParseNormalizedNamed(data)
	})
}
