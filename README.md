# pidgin-kwallet

Pidgin usually stores passwords as plaintext. This plugin instead saves all
passwords to kwallet, which some would argue is a more secure form
of password storage.

After the plugin is enabled, whenever an account with a pidgin-stored password
signs on, its password will automatically be saved to kwallet and removed
from the plaintext accounts.xml file.

## Building from source
You will need the libpurple development libraries, along with
pkg-config ang golang.

The go dependencies can be installed through:
```go get -v github.com/godbus/dbus```

Afterwards, use git clone to download the source (the version is set by git),
and run ```make``` to compile. For installation:
 - local ($HOME/.purple/plugins): ```make install_local```
 - global (/usr/lib/purple-2): ```sudo make install```

## Why go?
Because I don't like the dbus APIs for C

# Credits
This plugin is a heavily modified version of [pidgin-gnome-keyring](https://github.com/aebrahim/pidgin-gnome-keyring)

## Modifications
- 2018 Marcus Soll: Modified version to support kwallet instead of gnome-keyring
