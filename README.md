# Building

## All Systems
This application uses the awesome [robotgo](https://github.com/go-vgo/robotgo) library, which requires `gcc` to be
installed and available on the path.

## Windows
This application uses the [systray](https://github.com/getlantern/systray) library, which requires specific linker flags
when targeting Windows in order to prevent a console application window from opening. For a windows build, include
`-H=windowsgui` in your `ldflags` options.

```
go build -ldflags -H=windowsgui -o presence.exe cmd/presence/main.go
```
