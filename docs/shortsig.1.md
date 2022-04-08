% HELLO(1) Version 1.0 | Frivolous "Hello World" Documentation

NAME
====
**shortsig** — starts the tcp server

SYNOPSIS
========
| **shortsig** \[**-p**|**--port** _number_] \[**-h**|**--help**|**-v**|**--version**]

DESCRIPTION
===========
Spawns a tcp server that executes one of many predefined commands depending on the incoming traffic.

Options
-------
-h, --help
:   Prints brief usage information.

-p, --port _number_
:   Sets the tcp server port.

<!-- -v, --version -->
<!-- :   Prints the current version number. -->

FILES
=====
config.toml
------
It's seached in these places in the following order:

1. /etc/shortsig
1. $XDG_CONFIG_HOME/shortsig
1. $HOME/.config/shortsig
1. /home/$USER/.config/shortsig

It can contain zero or more of the following toml keys/value pairs:

**Port**
: A uint16 number defines the port the TCP server will listen on.

**Routines**
: A routine array, Check Routine section for more detail.

<!-- **Whitelist** -->
<!-- : A string array that holds all the whitelisted mac addresses  -->

Routine
-------
A routine is a table that contained one or more of the following keys

linux
: A string that defines a command executed when the routine is called on a linux platform
darwin
: A string that defines a command executed when the routine is called on a darwin platform
windows
: A string that defines a command executed when the routine is called on a windows platform

Example
-------
port = 4321  

[routines.poweroff]
linux = "systemctl poweroff"  
[routines.reboot]
linux = "systemctl reboot"  
[routines.suspend]
linux = "systemctl suspend"  
[routines.lock]
linux = "loginctl lock-session $XDG_SESSION_ID"  
[routines.ls]
linux = "ls"
darwin = "ls"
windows = "dir"

<!-- ENVIRONMENT -->
<!-- =========== -->

BUGS
====
See GitHub Issues: <https://github.com/RaafatTurki/shortsig/issues>

AUTHOR
======
Raafat Turki <raafat.turki@pm.me>  
The project is located at https://github.com/RaafatTurki/shortsig

SEE ALSO
========
shortsig was inspired by unified remote server.

