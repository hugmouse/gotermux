package gotermuxwrapper

import (
	"testing"
)

func TestTermuxDialog(t *testing.T) {
	title := "test"
	result := TermuxDialog(title)
	if result.Code != -1 && len(result.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-1, \"\" \".", result.Code, result.Text)
	}
}