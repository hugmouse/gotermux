<p align="center">
    <img src="https://raw.githubusercontent.com/hugmouse/gotermux/Development/icon/logo.webp.webp" height="256">
</p>

# GoTermux

[![Codacy Badge](https://api.codacy.com/project/badge/Grade/380f19e0a1bc4fb19d3eeafa914fc1ad)](https://www.codacy.com/app/mysh/gotermux?utm_source=github.com&utm_medium=referral&utm_content=hugmouse/gotermux&utm_campaign=Badge_Grade)
[![Go Report Card](https://goreportcard.com/badge/github.com/hugmouse/gotermux)](https://goreportcard.com/report/github.com/hugmouse/gotermux)
[![GoDoc](https://godoc.org/github.com/hugmouse/gotermux?status.svg)](https://godoc.org/github.com/hugmouse/gotermux)

`Обертка вокруг Termux API`

В настоящее время в активной разработке!

## Codestyle

Вызов команды termux из golang выглядит так же, как в termux cmd, но без тире (-).

Пример:

```shell
// Termux CMD:
termux-battery-status
// Golang: 
TermuxBatteryStatus()
```

## Примеры работы (WIP)

Примеры кода можно найти на: [gotermux.mysh.dev](https://gotermux.mysh.dev/)!

Или здесь: [now.sh](https://gotermux.mysh.now.sh/)

## Установка Termux API

Используйте [F-droid](https://f-droid.org/packages/com.termux.api/) или [Google Play](https://play.google.com/store/apps/details?id=com.termux.api) для установки Termux API.

**Не миксуйте установки Termux и аддонов для него из Google Play и F-Droid**.

Чтобы **использовать** Termux API вам так же нужно установить пакет [termux-api](https://github.com/termux/termux-api-package):

```shell
pkg install termux-api
```

## Установка пакета

Сперва нужно установить пакет в Golang в Termux:

```shell
pkg install golang
```

Затем:

```shell
go get -u github.com/hugmouse/gotermux
```

## TODO

### По проекту

- [ ] Добавить примеры
- [ ] Добавить более подробное API

### По Termux API

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

- [x] termux-infrared-frequencies

- [x] termux-infrared-transmit

- [ ] termux-job-scheduler

- [x] termux-location

- [x] termux-media-player

  - [x] termux-media-player play
  - [x] termux-media-player play (resume)
  - [x] termux-media-player info
  - [x] termux-media-player pause
  - [x] termux-media-player stop

- [x] termux-media-scan

- [ ] termux-microphone-record

- [ ] termux-sensor

- [ ] termux-messages-remove

- [ ] termux-sensor

- [x] termux-share

- [x] termux-telephony-call

- [x] termux-telephony-cellinfo

- [x] termux-telephony-deviceinfo

- [ ] termux-toast

- [x] termux-torch

- [x] termux-tts-engines

- [ ] termux-telephony-call

- [x] termux-vibrate

- [x] termux-volume

- [x] termux-wallpaper

- [x] termux-wifi-connectioninfo

- [x] termux-wifi-enable

- [x] termux-wifi-scaninfo

## Тесты функций Termux API

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

## Штуки, которые Google больше не поддерживает

**И я скорее всего не буду добавлять такие функции**

- termux-sms-list
- termux-sms-send

## Не смог протестировать или использовать

**Штуки, которые не работают на моем телефоне и я не могу их протестировать**

Так что по этой причине (невозможности протестировать) в моем пакете нет следующих функций:

- termux-storage-get

- termux-fingerprint (не работает по неизвестным мне причинам)
