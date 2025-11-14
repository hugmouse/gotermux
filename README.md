<p align="center">
    <img src="https://raw.githubusercontent.com/hugmouse/gotermux/Development/icon/logo.webp">
</p>

# GoTermux

> [!WARNING]
> This library is no longer supported and will not receive any updates. Feel free to browse available methods and documentation, but nothing new will be added.

GoTermux is a wrapper library around the termux-api. 
It allows you to call `termux-*` scripts and read the their output in a convenient way.

## Example

```go
package main

import (
    "fmt"
    t "github.com/hugmouse/gotermux"
)

func main() {
    battery := t.TermuxBatteryStatus()
    fmt.Println(battery.Percentage) // This will print "38.0"
}
```

## Documentation

<picture>
  <source srcset="https://github.com/user-attachments/assets/38edb1d2-0067-43fb-a155-f7275d879b31" media="(prefers-color-scheme: light)">
  <source srcset="https://github.com/user-attachments/assets/fd4c8793-7f78-42f3-b296-4cd8fa3aad10" media="(prefers-color-scheme: dark)">
  <img src="https://github.com/user-attachments/assets/fd4c8793-7f78-42f3-b296-4cd8fa3aad10" alt="Documentation">
</picture>


You can find a quickstart guide here: https://hugmouse.github.io/gotermux-docs/0.1/overview/

Termux API reference: https://hugmouse.github.io/gotermux-docs/0.1/termux-api/

GoTermux reference: https://hugmouse.github.io/gotermux-docs/0.1/gotermux

API documentation can also be found on pkg.go.dev: https://pkg.go.dev/github.com/hugmouse/gotermux
