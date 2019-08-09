package gotermux

// TValue used in TResult's struct
//
// Some functions returns multiple values with "Index" and "Text" fields
type TValue struct {
    Index uint   `json:"Index"`
    Text  string `json:"Text"`
}

// TResult returns some defaults fields that Termux API provides to us
type TResult struct {
    Code   int8     `json:"Code"` // -2, -1 or 0 in rare cases
    Text   string   `json:"Text"` // Usually just information about action
    Index  uint     `json:"Index"` // With new update returns index of himself
    Values []TValue `json:"Values"` // Some functions returns multiple results
    Error  string   `json:"Error"` // Usually if there is an error - it's single element of all struct
}

// TBattery is a structure for return values from TermuxBatteryStatus
type TBattery struct {
    Health      string  `json:"Health"` // Different statuses: COLD, DEAD, GOOD, OVERHEAT, OVER_VOLTAGE, UNKNOWN and UNSPECIFIED_FAILURE
    Percentage  uint    `json:"Percentage"` // How charged your battery is
    Plugged     string  `json:"Plugged"` // Different statuses: UNPLUGGED, PLUGGED_AC, PLUGGED_USB, PLUGGED_WIRELESS and PLUGGED_+int (0 means it is on battery other constants are different types of power sources)
    Status      string  `json:"Status"` // Different statuses: CHARGING, DISCHARGING, FULL, NOT_CHARGING, UNKNOWN
    Temperature float64 `json:"Temperature"` // Just temperature of your battery
}

// TClipboard used for TermuxClipboardSet function
type TClipboard struct {
    Text string // string that goes into clipboard
}

// TContact is a single contact from your phone, used in TContacts
type TContact struct {
    Name   string `json:"Name"` // Name of your contact
    Number string `json:"Number"` // Number of your contact
}

// TContacts used for TermuxContactList, list of contacts in your phone
type TContacts struct {
    Contact []TContact // Slice of your contacts
}

// TDialogTitle represents Title in TermuxDialog, used in every single TDialog function
type TDialogTitle struct {
    Title string // Name of the window
}

// TDialogConfirm used in TDialogConfirm function
type TDialogConfirm struct {
    Hint string // Like a placeholder - what your user see in input box
    TDialogTitle // Name of the window
}

// TDialogCheckbox used in TDialogCheckbox function
type TDialogCheckbox struct {
    Values []string // Values that your user can select
    TDialogTitle // Name of the window
}

// TDialogCounter used id TDialogCounter function
type TDialogCounter struct {
    Min   int // Minimum value
    Max   int // Maximum value
    Start int // The default value is where the counting starts
    TDialogTitle // Name of the window
}

// TDialogRadio have same struct as TDialogCheckbox
type TDialogRadio struct {
    TDialogCheckbox // Multiple values, but with radio buttons
}

// TDialogSheet have same struct as TDialogCheckbox
type TDialogSheet struct {
    TDialogCheckbox // Multiple values, but with sheet button
}

// TDialogSpinner have same struct as TDialogCheckbox
type TDialogSpinner struct {
    TDialogCheckbox // Multiple values, but with spinner button, where user can pick one
}

// TDialogSpeech have same struct as TDialogConfirm
type TDialogSpeech struct {
    TDialogConfirm // In that case there is no input box, just gray text in the middle
}

// TDialogText used in TDialogText function
type TDialogText struct {
    Hint          string // Like a placeholder - what your user see in input box
    MultipleLine  bool // Multiple lines instead of single
    NumberInput   bool // Enter input as numbers
    PasswordInput bool // Enter input as password
    TDialogTitle // Name of the window
}

// TDialogTime used in TDialogTime function, but it's just have the same struct as TDialogTitle
type TDialogTime struct {
    TDialogTitle // Name of the window
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

type TShare struct {
    Action      ShareAction
    ContentType string
    Default     bool
    TDialogTitle
}

type TVibrate struct {
    Duration uint
    SilentModeIgnore bool
}
