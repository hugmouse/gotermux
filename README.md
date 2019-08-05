# Go Termux Wrapper
### Simple wrapper around Termux API

# Work In Progress
This package is still in development. 

# Codestyle
Calling termux command from golang looks just like in termux cmd, but without dash (`-`). 

Example: 
```shell
// Termux CMD:
termux-battery-status
// Golang: 
TermuxBatteryStatus()
```

# TODO
### More project related:
- [ ] Add examples
- [ ] Add better description of API, in contrast to the official in Termux
- [ ] Make test coverage >50%
### Termux API:
- [x] termux-battery-status

- [x] termux-brightness

- [x] termux-call-log

- [ ] termux-camera-info

- [ ] termux-camera-photo

- [x] termux-clipboard-get

- [x] termux-clipboard-set

- [x] termux-contact-list

- [x] termux-dialog
  - [x] dialog-confirm
  - [x] dialog-checkbox
  - [x] dialog-counter
  - [ ] dialog-date
  - [x] dialog-sheet
  - [x] dialog-spinner
  - [x] dialog-speech
  - [x] dialog-text
  - [x] dialog-time

- [x] termux-download

- [ ] termux-fingerprint

- [x] termux-infrared-frequencies

- [x] termux-infrared-transmit

- [ ] termux-job-scheduler

- [x] termux-location

- [x] termux-media-player
  - [x] termux-media-player play <file>
  - [x] termux-media-player play (resume)
  - [x] termux-media-player info
  - [x] termux-media-player pause
  - [x] termux-media-player stop

- [x] termux-media-scan

- [ ] termux-microphone-record

- [ ] termux-notification

- [ ] termux-notification-remove

- [ ] termux-sensor

- [ ] termux-share

- [ ] termux-sms-list

- [ ] termux-sms-send

- [ ] termux-storage-get

- [ ] termux-telephony-call

- [ ] termux-telephony-cellinfo

- [ ] termux-telephony-deviceinfo

- [ ] termux-toast

- [ ] termux-torch

- [ ] termux-tts-engines

- [ ] termux-tts-speak

- [ ] termux-vibrate

- [ ] termux-volume

- [ ] termux-wallpaper

- [ ] termux-wifi-connectioninfo

- [ ] termux-wifi-enable

- [ ] termux-wifi-scaninfo

### Termux API tests:

- [x] termux-dialog
  - [x] dialog-confirm
  - [x] dialog-checkbox
  - [x] dialog-counter
  - [ ] dialog-date
  - [x] dialog-sheet
  - [x] dialog-spinner
  - [x] dialog-speech
  - [ ] dialog-text
  - [ ] dialog-time
