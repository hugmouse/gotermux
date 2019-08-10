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
	Code   int8     `json:"Code"`   // -2, -1 or 0 in rare cases
	Text   string   `json:"Text"`   // Usually just information about action
	Index  uint     `json:"Index"`  // With new update returns index of himself
	Values []TValue `json:"Values"` // Some functions returns multiple results
	Error  string   `json:"Error"`  // Usually if there is an error - it's single element of all struct
}

// TBattery is a structure for return values from TermuxBatteryStatus
type TBattery struct {
	Health      string  `json:"Health"`      // Different statuses: COLD, DEAD, GOOD, OVERHEAT, OVER_VOLTAGE, UNKNOWN and UNSPECIFIED_FAILURE
	Percentage  uint    `json:"Percentage"`  // How charged your battery is
	Plugged     string  `json:"Plugged"`     // Different statuses: UNPLUGGED, PLUGGED_AC, PLUGGED_USB, PLUGGED_WIRELESS and PLUGGED_+int (0 means it is on battery other constants are different types of power sources)
	Status      string  `json:"Status"`      // Different statuses: CHARGING, DISCHARGING, FULL, NOT_CHARGING, UNKNOWN
	Temperature float64 `json:"Temperature"` // Just temperature of your battery
}

// TClipboard used for TermuxClipboardSet function
type TClipboard struct {
	Text string // string that goes into clipboard
}

// TContact is a single contact from your phone, used in TContacts
type TContact struct {
	Name   string `json:"Name"`   // Name of your contact
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
	Hint         string // Like a placeholder - what your user see in input box
	TDialogTitle        // Name of the window
}

// TDialogCheckbox used in TDialogCheckbox function
type TDialogCheckbox struct {
	Values       []string // Values that your user can select
	TDialogTitle          // Name of the window
}

// TDialogCounter used id TDialogCounter function
type TDialogCounter struct {
	Min          int // Minimum value
	Max          int // Maximum value
	Start        int // The default value is where the counting starts
	TDialogTitle     // Name of the window
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
	MultipleLine  bool   // Multiple lines instead of single
	NumberInput   bool   // Enter input as numbers
	PasswordInput bool   // Enter input as password
	TDialogTitle         // Name of the window
}

// TDialogTime used in TDialogTime function, but it's just have the same struct as TDialogTitle
type TDialogTime struct {
	TDialogTitle // Name of the window
}

// TLocation used in TermuxLocation function
//
// Still in development
type TLocation struct {
	Provider string // location provider [gps/network/passive]
	Request  string // kind of request to make [once/last/updates]
}

// TLocationResult is a return of a TLocation function
type TLocationResult struct {
	Latitude  float64 `json:"Latitude"`  // The latitude, in degrees
	Longitude float64 `json:"Longitude"` // The longitude, in degrees.
	Altitude  float64 `json:"Altitude"`  // The altitude if available, in meters above the WGS 84 reference ellipsoid
	Accuracy  float64 `json:"Accuracy"`  // The estimated horizontal accuracy of this location, radial, in meters
	Bearing   float64 `json:"Bearing"`   // Bearing is the horizontal direction of travel of this device, and is not related to the device orientation
	Speed     float64 `json:"Speed"`     // Speed in meters/second over ground
	ElapsedMS uint64  `json:"ElapsedMs"` // Uncertainty of elapsed real-time of fix
	Provider  string  `json:"Provider"`  // The name of the provider that generated this fix
}

// TShare used in TermuxShare function
type TShare struct {
	Action       ShareAction // which action to performed on the shared content
	ContentType  string      // content-type to use (default: guessed from file extension, text/plain for stdin
	Default      bool        // share to the default receiver if one is selected instead of showing a chooser
	TDialogTitle             // Name of the window
}

// TVibrate used in TermuxVibrate function
type TVibrate struct {
	Duration         uint // the duration to vibrate in ms
	SilentModeIgnore bool // force vibration even in silent mode
}

// TAudioStream used in TermuxValue function
type TAudioStream struct {
	Stream    string `json:"Stream"` // Valid audio streams are: alarm, music, notification, ring, system, call
	Volume    uint   `json:"Volume"` // Current volume
	MaxVolume uint   `json:"Max_volume"`
}

// TConnection used in TermuxWifiConnectionInfo
//
// Note that NetworkID can be -1 if there is no currently connected network or if the caller has insufficient permissions to access the network ID
//
// If the SSID can be decoded as UTF-8, it will be returned surrounded by double quotation marks. Otherwise, it is returned as a string of hex digits
//
// The SSID may be <unknown ssid>, if there is no network currently connected or if the caller has insufficient permissions to access the SSID
//
// Supplicant states: ASSOCIATED, ASSOCIATING, AUTHENTICATING, COMPLETED, DISCONNECTED, DORMANT, FOUR_WAY_HANDSHAKE, GROUP_HANDSHAKE, INACTIVE, INTERFACE_DISABLED, INVALID, SCANNING and UNINITIALIZED
type TConnection struct {
	BSSID           string `json:"Bssid"`           // Basic service set identifier
	FrequencyMhz    int    `json:"Frequency_mhz"`   // The current frequency
	IP              string `json:"Ip"`              // IPV4 only
	LinkSpeedMbps   uint   `json:"Link_speed_mbps"` // The current link speed
	MACAddress      string `json:"Mac_address"`
	NetworkID       int    `json:"Network_id"`       // Each configured network has a unique small integer ID, used to identify the network
	RSSI            int    `json:"Rssi"`             // The received signal strength indicator of the current 802.11 network, in dBm
	SSID            string `json:"Ssid"`             // The service set identifier
	HiddenSSID      bool   `json:"Ssid_hidden"`      // True if this network does not broadcast its SSID, so an SSID-specific probe request must be used for scans
	SupplicantState string `json:"Supplicant_state"` // The detailed state of the supplicant's negotiation with an access point
}
