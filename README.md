# SHORTSIG
a tool that executes routines from a TCP connection

## Installation
###### From Source
clone and `go build`

## Config
```toml
# ~/.config/shortsig/config.toml

port = 3003

[routines.poweroff]
linux = "poweroff"

[routines.reboot]
linux = "reboot"

[routines.suspend]
linux = "systemctl suspend"

[routines.lock]
linux = "loginctl lock-session $XDG_SESSION_ID"

[routines.sleep]
linux = "sleep 2"

[routines.ls]
linux = "ls"
darwin = "ls"
windows = "dir"
```

## Usage
launch the server  
`shortsig`

send it something through tcp  
`nc localhost 3003`  
`exec lock`  

## WIP
this project is a work in progress, there's a lot left to be desired such as
- auth
- android client
- aur package (it's on my [personal arch repo](https://github.com/RaafatTurki/pkgs) for now)
