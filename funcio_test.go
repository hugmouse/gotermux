package gotermuxwrapper

import (
	"testing"
)

var (
	Title       = "test"
	Hint        = "Press \"Yes\" then \"No\""
	TestConfirm = TDialogConfirm{
		Hint,
		TDialog{Title},
	}
	TestCheckbox = TDialogCheckbox{
		[]string{"Value"},
		TDialog{Title},
	}
	TestDialog = TDialogCounter{
		0,
		2,
		1,
		TDialog{Title},
	}
)

func TestTermuxDialog(t *testing.T) {

	resultYes := TermuxDialog(Title)
	if resultYes.Code != -1 || len(resultYes.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-1, \"\" \".", resultYes.Code, resultYes.Text)
	}

	resultNo := TermuxDialog(Title)
	if resultNo.Code != -2 || len(resultNo.Text) != 0 {
		t.Errorf("TermuxDialog() was incorrect, got: \"%d, %s\" want: \"-2, \"\" \".", resultNo.Code, resultNo.Text)
	}

}

func TestTermuxDialogConfirm(t *testing.T) {

	resultYes := TermuxDialogConfirm(TestConfirm)
	if resultYes.Code != 0 || resultYes.Text != "yes" {
		t.Errorf("TermuxDialogConfirm() was incorrect, got: \"%d, %s\" want: \"0, \"yes\" \".", resultYes.Code, resultYes.Text)
	}

	resultNo := TermuxDialogConfirm(TestConfirm)
	if resultNo.Code != 0 || resultNo.Text != "no" {
		t.Errorf("TermuxDialogConfirm() was incorrect, got: \"%d, %s\" want: \"0, \"no\" \".", resultNo.Code, resultNo.Text)
	}

}

func TestTermuxDialogCheckbox(t *testing.T) {

	resultYes := TermuxDialogCheckbox(TestCheckbox)
	if resultYes.Code != -1 || resultYes.Text != "[value]" {
		t.Errorf("TermuxDialogCheckbox() was incorrect, got: \"%d, %s\" want: \"-1, \"[value]\" \".", resultYes.Code, resultYes.Text)
	}
	if resultYes.Values[0].Index != 0 || resultYes.Values[0].Text != "value" {
		t.Errorf("TermuxDialogCheckbox() was incorrect, got: \"%d, %s\" want: \"0, \"value\" \".", resultYes.Values[0].Index, resultYes.Values[0].Text)
	}

	resultNo := TermuxDialogCheckbox(TestCheckbox)
	if resultNo.Code != -2 || len(resultNo.Text) != 0 {
		t.Errorf("TermuxDialogCheckbox() was incorrect, got: \"%d, %s\" want: \"-2, \"\" \".", resultNo.Code, resultNo.Text)
	}
}

func TestTermuxDialogCounter(t *testing.T) {
	ughs := []struct {
		Min int
		Max int
		Start int
		WantedCode int8
		WantedString string
	}{
		{0, 2, 1, -1, "1"},
		{0, 2, 2, -1, "2"},
		{0, 2, 0,-1, "0"},
		{0,2,0,-2,""},
	}

	for _, ugh := range ughs {
		result := TermuxDialogCounter(TDialogCounter{
			ugh.Min,
			ugh.Max,
			ugh.Start,
			TDialog{"Just press \"ok\""},
		})
		if result.Code != ugh.WantedCode || result.Text != ugh.WantedString {
			t.Errorf("TermuxDialogCounter() was incorrect, got: \"%d, %s\" want: \"%d, \"%s\" \".", result.Code, result.Text, ugh.WantedCode, ugh.WantedString)
		}
	}
	
}
