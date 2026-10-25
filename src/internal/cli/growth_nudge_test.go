// SPDX-License-Identifier: Apache-2.0

package cli

import (
	"bytes"
	"strings"
	"testing"
)

func TestPrintSignupNudge_OnceThenSilent(t *testing.T) {
	t.Setenv("SNOWOPS_HOME", t.TempDir())

	var first bytes.Buffer
	printSignupNudge(&first)
	if !strings.Contains(first.String(), "labs.snowops.net") {
		t.Fatalf("first call: expected the signup URL, got %q", first.String())
	}

	var second bytes.Buffer
	printSignupNudge(&second)
	if second.Len() != 0 {
		t.Fatalf("second call: expected silence once the marker exists, got %q", second.String())
	}
}

func TestPrintSignupNudge_MissingHomeDoesNotPanic(t *testing.T) {
	t.Setenv("SNOWOPS_HOME", "")
	t.Setenv("HOME", "")

	var out bytes.Buffer
	printSignupNudge(&out) // must not panic even if config.Home() errors
}
