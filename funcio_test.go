package gotermuxwrapper

import (
	"testing"
)

var (
	Title       = "test"
	Hint        = "Press \"Yes\" then \"No\""
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

	resultYes := TermuxDialogConfirm(TDialogConfirm{
		Hint,
		TDialog{Title},
	})
	if resultYes.Code != 0 || resultYes.Text != "yes" {
		t.Errorf("TermuxDialogConfirm() was incorrect, got: \"%d, %s\" want: \"0, \"yes\" \".", resultYes.Code, resultYes.Text)
	}

	resultNo := TermuxDialogConfirm(TDialogConfirm{
		Hint,
		TDialog{Title},
	})
	if resultNo.Code != 0 || resultNo.Text != "no" {
		t.Errorf("TermuxDialogConfirm() was incorrect, got: \"%d, %s\" want: \"0, \"no\" \".", resultNo.Code, resultNo.Text)
	}

}

func TestTermuxDialogCheckbox(t *testing.T) {

	resultYes := TermuxDialogCheckbox(TDialogCheckbox{
		[]string{"value"},
		TDialog{Title},
	})
	if resultYes.Values[0].Index != 0 || resultYes.Values[0].Text != "value" {
		t.Errorf("TermuxDialogCheckbox() was incorrect, got: \"%d, %s\" want: \"0, \"value\" \".", resultYes.Values[0].Index, resultYes.Values[0].Text)
	}

}

func TestTermuxDialogCounter(t *testing.T) {
	tests := []struct {
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

	for _, test := range tests {
		result := TermuxDialogCounter(TDialogCounter{
			test.Min,
			test.Max,
			test.Start,
			TDialog{"Just press \"ok\". On the 4th test press \"Cancel\""},
		})
		if result.Code != test.WantedCode || result.Text != test.WantedString {
			t.Errorf("TermuxDialogCounter() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \".", result.Code, result.Text, test.WantedCode, test.WantedString)
		}
	}
	
}

func TestTermuxDialogDate(t *testing.T) {
	tests := []struct{
		Day uint
		Month uint
		Year uint
		Khour uint
		Minutes uint
		Seconds uint
		WantedText string
		WantedCode int8
	} {
		{01,01,2000,12,00,00, "01-01-2000 12:00:00", -1 },
		{1,1,2000,12,00,00, "1-1-2000 12:00:00", -1 },
	}

	for _, test := range tests {
		result := TermuxDialogDate(TDialogDate{
			TDialogDatePattern{
				test.Day,
				test.Month,
				test.Year,
				test.Khour,
				test.Minutes,
				test.Seconds,
			},
			TDialog{"Just press \"OK\""},
		})
		if result.Code != test.WantedCode || result.Text != test.WantedText {
			t.Errorf("TermuxDialogDate() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \"", result.Code, result.Text, test.WantedCode, test.WantedText)
		}
	}
}

// TODO: do time.Now() and make format like in Termux output.
//func TestTermuxDialogDateWithoutDate(t *testing.T) {
//	tests := []struct{
//		WantedText string
//		WantedCode int8
//	} {
//
//	}
//
//	for _, test := range tests {
//		result := TermuxDialogWithoutDate(TDialog{"Just press \"OK\""})
//		if result.Code != test.WantedCode || result.Text != test.WantedText {
//			t.Errorf("TermuxDialogDate() was incorrect, got: \"%d, %s\". Want: \"%d, \"%s\" \"", result.Code, result.Text, test.WantedCode, test.WantedText)
//		}
//	}
//}
