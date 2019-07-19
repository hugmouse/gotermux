package gotermuxwrapper

type TBattery struct {
	Health      string
	Percentage  uint
	Plugged     string
	Status      string
	Temperature float64
}

//type AutoExposureModes struct {
//
//}

type TPhysicalSize struct {
	Width  float64
	Height float64
}

type TCamera struct {
	ID     int
	Facing string
	// Struct of struct, jpeg bluya
	FocalLengths float64
	TPhysicalSize
}

type TClipboard struct {
	Text string
}

type TContact struct {
	Name   string
	Number string
}

type TDialog struct {
	Title string
}

type TDialogConfirm struct {
	Hint string
	TDialog
}

type TDialogCheckbox struct {
	Values []string
	TDialog
}

type TDialogCounter struct {
	Min   int
	Max   int
	Start int
	TDialog
}

type TDialogDatePattern struct {
	Day     uint
	Month   uint
	Year    uint
	KHours  uint // %k - range 0-23
	Minuts  uint
	Seconds uint
}
type TDialogDate struct {
	TDialogDatePattern
	TDialog
}

type TDialogRadio struct {
	TDialogCheckbox
}

type TDialogSheet struct {
	TDialogCheckbox
}

type TDialogSpinner struct {
	TDialogCheckbox
}

type TDialogSpeech struct {
	TDialogConfirm
}

type TDialogText struct {
	Hint          string
	MultipleLine  bool
	NumberInput   bool
	PasswordInput bool
	TDialog
}

type TDialogTime struct {
	TDialog
}
