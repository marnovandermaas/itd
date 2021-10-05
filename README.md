# ITD
## InfiniTime Daemon

`itd` is a daemon that uses my infinitime [library](https://go.arsenm.dev/infinitime) to interact with the [PineTime](https://www.pine64.org/pinetime/) running [InfiniTime](https://infinitime.io).

[![Build status](https://ci.appveyor.com/api/projects/status/xgj5sobw76ndqaod?svg=true)](https://ci.appveyor.com/project/moussaelianarsen/itd)
[![Binary downloads](https://img.shields.io/badge/download-binary-orange)](https://minio.arsenm.dev/minio/itd/)
[![AUR package](https://img.shields.io/aur/version/itd-git?label=itd-git&logo=archlinux)](https://aur.archlinux.org/packages/itd-git/)

---

### Features

- Notification relay
- Notificstion transliteration
- Music control
- Get info from watch (HRM, Battery level, Firmware version)
- Set current time
- Control socket
- Firmware upgrades

---

### Socket

This daemon creates a UNIX socket at `/tmp/itd/socket`. It allows you to directly control the daemon and, by extension, the connected watch.

The socket accepts JSON requests. For example, sending a notification looks like this:

```json
{"type": "notify", "data": {"title": "title1", "body": "body1"}}
```

It will return a JSON response. A response can have 3 fields: `error`, `msg`, and `value`. Error is a boolean that signals whether an error was returned. If error is true, the msg field will contain the error. Value can contain any data and depends on what the request was.

The various request types and their data requirements can be seen in `internal/types`. I can make separate docs for it if I get enough requests.

---

### Transliteration

Since the PineTime does not have enough space to store all unicode glyphs, it only stores the ASCII space and Cyrillic. Therefore, this daemon can transliterate unsupported characters into supported ones. Since some languages have different transliterations, the transliterators to be used must be specified in the config. Here are the available transliterators:

- eASCII
- Scandinavian
- German
- Hebrew
- Greek
- Russian
- Ukranian
- Arabic
- Farsi
- Polish
- Lithuanian
- Estonian
- Icelandic
- Czeck
- French
- Armenian
- Korean
- Emoji

Place the desired map names in an array as `notifs.translit.use`. They will be evaluated in order. You can also put custom transliterations in `notifs.translit.custom`. These take priority over any other maps. The `notifs.translit` config section should look like this:

```toml
[notifs.translit]
    use = ["eASCII", "Russian", "Emoji"]
    custom = [
        "test", "replaced"
    ]
```

---

### `itctl`

This daemon comes with a binary called `itctl` which uses the socket to control the daemon from the command line. As such, it can be scripted using bash.

This is the `itctl` usage screen:
```
Control the itd daemon for InfiniTime smartwatches

Usage:
  itctl [command]

Available Commands:
  completion  generate the autocompletion script for the specified shell
  firmware    Manage InfiniTime firmware
  get         Get information from InfiniTime
  help        Help about any command
  notify      Send notification to InfiniTime
  set         Set information on InfiniTime

Flags:
  -h, --help   help for itctl

Use "itctl [command] --help" for more information about a command.
```

---

### `itgui`

In `cmd/itgui`, there is a gui frontend to the socket of `itd`. It uses the [fyne library](https://fyne.io/) for Go. It can be compiled by running:

```shell
go build ./cmd/itgui
```

#### Screenshots

![Info tab](https://i.imgur.com/okxG9EI.png)

![Notify tab](https://i.imgur.com/DrVhOAq.png)

![Set time tab](https://i.imgur.com/j9civeY.png)

![Upgrade tab](https://i.imgur.com/1KY6fG4.png)

![Upgrade in progress](https://i.imgur.com/w5qbWAw.png)

---

#### Interactive mode

Running `itctl` by itself will open interactive mode. It's essentially a shell where you can enter commands. For example:

```
$ itctl                        
itctl> fw ver
1.3.0
itctl> get batt
81%
itctl> get heart
92 BPM
itctl> set time 2021-08-22T00:06:18-07:00
itctl> set time now
itctl> exit
```

---

### Installation

To install, install the go compiler and make. Usually, go is provided by a package either named `go` or `golang`, and make is usually provided by `make`. The go compiler must be version 1.16 or newer for the `io/fs` module.

To install, run
```shell
make && sudo make install
```

---

### Starting

To start the daemon, run the following **without root**:

```shell
systemctl --user start itd
```

To autostart on login, run:
```shell
systemctl --user enable itd
```

---

### Cross compiling

To cross compile, simply set the go environment variables. For example, for PinePhone, use:

```shell
make GOOS=linux GOARCH=arm64
```

This will compile `itd` and `itctl` for Linux aarch64 which is what runs on the PinePhone. This daemon only runs on Linux due to the library's dependencies (`dbus`, `bluez`, and `playerctl` specifically).

---

### Configuration

This daemon places a config file at `/etc/itd.toml`. This is the global config. `itd` will also look for a config at `~/.config/itd.toml`.

Most of the time, the daemon does not need to be restarted for config changes to take effect.