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

// TToast used in TermuxToast function
type TToast struct {
	BackgroundColor Color    // BLACK, BLUE, CYAN, DKGRAY, GRAY (default), GREEN, LTGRAY, MAGENTA, RED, TRANSPARENT, WHITE, YELLOW
	TextColor       Color    // BLACK, BLUE, CYAN, DKGRAY, GRAY (default), GREEN, LTGRAY, MAGENTA, RED, TRANSPARENT, WHITE, YELLOW
	ToastPosition   Position // TOP, MIDDLE (default), BOTTOM
	Short           bool     // Only show the toast for a short while?
}

// TRecording used for TermuxMicrophone related functions
type TRecording struct {
	Filename   string  // Specific file
	TimeLimit  int     // Specific time limit (0 for unlimited)
	Encoder    Encoder // Specific encoder (AAC, AACELD, HEAAC, AMRWB, AMRNB, OPUS, Vorbis)
	BitRate    int     // The number of bits that are conveyed or processed per unit of time
	SampleRate int     // A commonly seen unit of sampling rate is Hz, which stands for Hertz and means "samples per second".
	Channels   Channel // 1 for Mono, 2 for Stereo
}

// TRecordingInfo used in TermuxMicrophoneRecordInfo
type TRecordingInfo struct {
	IsRecording bool   `json:"isRecording"` // true of false
	OutputFile  string `json:"outputFile"`  // Absolute path
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

// TConnectionScan used in TermuxWifiScanInfo
type TConnectionScan struct {
	BSSID               string `json:"Bssid"`         // Basic service set identifier
	FrequencyMhz        int    `json:"Frequency_mhz"` // The current frequency
	RSSI                int    `json:"Rssi"`          // The received signal strength indicator of the current 802.11 network, in dBm
	SSID                string `json:"Ssid"`          // The service set identifier
	Timestamp           int64  `json:"Timestamp"`
	ChannelBandwidthMhz string `json:"Channel_bandwidth_mhz"` // 20, 40, 80, 80+80, 160 or unknown
	CenterFrequencyMhz  int    `json:"Center_frequency_mhz"`
}

// TWallpaper used in TermuxWallpaper function
type TWallpaper struct {
	Path       string // Absolute path
	URL        string
	Lockscreen bool
}

// TTTSEngine used in TermuxTTSEngines function
type TTSEngine struct {
	Name    string `json:"name"` // Something like "com.google.android.tts"
	Label   string `json:"label"`
	Default bool   `json:"default"`
}

// TCellInfo used in TermuxTelephonyCellInfo function
//
// You can read more about each thing here at https://developer.android.com/reference/android/telephony/package-summary
type TCellInfo struct {
	Type          string  `json:"type"`           // LTE, GSM, CDMA, NR, TDSCDMA, WCDMA
	Registered    bool    `json:"registered"`     // True if the phone is registered to a mobile network that provides service on this cell and this cell is being used or would be used for network signaling
	ASU           int     `json:"asu"`            // The RSSI in ASU. Asu is calculated based on 3GPP RSSI. Refer to 3GPP 27.007 (Ver 10.3.0) Sec 8.69
	DBM           int     `json:"dbm"`            // The signal strength as dBm
	Level         int     `json:"level"`          // An abstract level value for the overall signal quality
	TimingAdvance int     `json:"timing_advance"` // The GSM timing advance between 0..219 symbols (normally 0..63)
	CID           int     `json:"cid"`            // 16-bit GSM Cell Identity described in TS 27.007
	LAC           int     `json:"lac"`            // 16-bit Location Area Code, GSM only
	TAC           int     `json:"tac"`            // 16-bit Tracking Area Code, LTE only
	MCC           int     `json:"mcc"`            // Mobile Country Code
	MNC           int     `json:"mnc"`            // 2 or 3-digit Mobile Network Code
	CDMADBM       int     `json:"cdma_dbm"`       // The CDMA RSSI value in dBm
	CDMAECIO      int     `json:"cdma_ecio"`      // The CDMA Ec/Io value in dB*10
	CDMALevel     int     `json:"cdma_level"`     // CDMA as level 0..4
	CDMAdBm       int     `json:"cdm_dbm"`        // The signal strength as dBm
	EVDOdBm       int     `json:"evdo_dbm"`       // The EVDO RSSI value in dBm
	EVDOECIO      int     `json:"evdo_ecio"`      // The EVDO Ec/Io value in dB*10
	EVDOLevel     int     `json:"evdo_level"`     // EVDO as level 0..4
	EVDOSNR       int     `json:"evdo_snr"`       // The signal to noise ratio. Valid values are 0-8. 8 is the highest
	BasestationID int     `json:"basestation"`    // Base Station Id 0..65535
	Latitude      float64 `json:"latitude"`       // Base station latitude, which is a decimal number as specified in 3GPP2 C.S0005-A v6.0
	Longitude     float64 `json:"longitude"`      // Base station longitude, which is a decimal number as specified in 3GPP2 C.S0005-A v6.0
	NetworkID     int     `json:"network"`        // Network Id 0..65535
	SystemID      int     `json:"system"`         // System Id 0..32767
}

// TDevice used in TermuxTelephonyDeviceInfo function
type TDevice struct {
	DataEnabled           string `json:"data_enabled"`            // Whether mobile data is enabled or not per user setting
	DataActivity          string `json:"data_activity"`           // A constant indicating the current data connection state (cellular) [disconnected/connected/connecting/suspended]
	DataState             string `json:"data_state"`              // States: DISCONNECTED, CONNECTING, CONNECTED, SUSPENDED
	DeviceID              string `json:"device_id"`               // The unique device ID, for example, the IMEI for GSM and the MEID or ESN for CDMA phones
	DeviceSoftwareVersion string `json:"device_software_version"` // The software version number for the device, for example, the IMEI/SV for GSM phones
	PhoneCount            int    `json:"phone_count"`             // 0 - if none of voice, sms, data is not supported / 1 - Single standby mode / 2 - for Dual standby mode / 3 - for Tri standby mode
	PhoneType             string `json:"phone_type"`              //  This indicates the type of radio used to transmit voice calls [CDMA/GSM/NONE/SIP]
	NetworkOperator       string `json:"network_operator"`        // The numeric name (MCC+MNC) of current registered operator
	NetworkOperatorName   string `json:"network_operator_name"`   // The alphabetic name of current registered operator
	NetworkCountryISO     string `json:"network_country_iso"`     // The ISO country code equivalent of the MCC (Mobile Country Code) of the current registered operator or the cell nearby, if available
	NetworkType           string `json:"network_type"`            // Codes: 1xRTT, CDMA, EDGE, EHRPD, EVDO_0, EVDO_A, EVDO_B, GPRS, HSDPA, HSPA, HSPAP, HSUPA, IDEN, LTE, UMTS, UNKNOWN
	NetworkRoaming        bool   `json:"network_roaming"`         // True if the device is considered roaming on the current network, for GSM purposes
	SimCountryIso         string `json:"sim_country_iso"`         // The ISO country code equivalent for the SIM provider's country code
	SimOperator           string `json:"sim_operator"`            // The MCC+MNC (mobile country code + mobile network code) of the provider of the SIM. 5 or 6 decimal digits
	SimOperatorName       string `json:"sim_operator_name"`       // The Service Provider Name (SPN)
	SimSerialNumber       string `json:"sim_serial_number"`       // The serial number of the SIM, if applicable
	SimSubscriberID       string `json:"sim_subscriber_id"`       // The unique subscriber ID, for example, the IMSI for a GSM phone
	SimState              string `json:"sim_state"`               // States: ABSENT, NETWORK_LOCKED, PIN_REQUIRED, PUK_REQUIRED, READY, UNKNOWN
}
