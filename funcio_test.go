package gotermuxwrapper

import (
	"log"
	"testing"
)

func TestTermuxDialog(t *testing.T) {
	log.Println("Press \"NO\" on first dialog and \"YES\" on second dialog")
	title := "test"

	resultNO := TermuxDialog(title)
	if resultNO.Code != -2 || len(resultNO.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-2, \"\" \".", resultNO.Code, resultNO.Text)
	}

	resultYES := TermuxDialog(title)
	if resultYES.Code == -1 || len(resultYES.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-1, \"\" \".", resultYES.Code, resultYES.Text)
	}
}