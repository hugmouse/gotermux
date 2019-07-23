package gotermuxwrapper

type TResult struct {
	Code  int8   `json:"Code"`
	Text  string `json:"Text"`
	Error string `json:"Error"`
}

type TBattery struct {
	Health      string  `json:"Health"`
	Percentage  uint    `json:"Percentage"`
	Plugged     string  `json:"Plugged"`
	Status      string  `json:"Status"`
	Temperature float64 `json:"Temperature"`
}

type TPhysicalSize struct {
	Width  float64
	Height float64
}

type TCamera struct {
	ID           int
	Facing       string
	FocalLengths float64
	TPhysicalSize
}

type TClipboard struct {
	Text string
}

type TContact struct {
	Name   string `json:"Name"`
	Number string `json:"Number"`
}

type TContacts struct {
	Contact []TContact
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
	Minutes uint
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

type TCall struct {
	Name        string `json:"Name"`
	PhoneNumber string `json:"Phonenumber"`
	Type        string `json:"Type"`
	Date        string `json:"Date"`
	Duration    string `json:"Duration"`
}

type TCalls struct {
	Calls []TCall
}

type TLocation struct {
	Provider string
	Request  string
}

type TLocationResult struct {
	Latitude  float64 `json:"Latitude"`
	Longitude float64 `json:"Longitude"`
	Altitude  float64 `json:"Altitude"`
	Accuracy  float64 `json:"Accuracy"`
	Bearing   float64 `json:"Bearing"`
	Speed     float64 `json:"Speed"`
	ElapsedMS uint64  `json:"ElapsedMs"`
	Provider  string  `json:"Provider"`
}
