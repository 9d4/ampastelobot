### /matchers

`package matchers`

This package is like router. It catchs text and check the type of the text. Condition:
- text is a command --> `command.go`
- text is callback --> `callback.go`
- text is general text --> `text.go`

Note: Command has been implemented. In `command.go` there is function named Command, it routes the command to the `package commands` in `/commands`
