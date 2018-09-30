# samp-terminal-launcher

[![Build status](https://ci.appveyor.com/api/projects/status/uyd6o81qe4l0nki9/branch/master?svg=true)](https://ci.appveyor.com/project/Bios-Marcel/samp-launcher-terminal/branch/master)

**This project is a WIP.**

This project mainly exists so I can test [rivo/tview](https://github.com/rivo/tview).

## How do I use it

You just start the executable in a commandline.

## Do I need anything else

No, it's a standalone executable file.

## What can it do

It uses [Southclaws/samp-servers-api](https://github.com/Southclaws/samp-servers-api) for
retrieving available servers. Connecting works via manual dll injection of the official SA-MP DLL.
Apparently this should work on linux, as I intend to add wine support. If your version is incompatible
with the servers version, it will automatically performa version change.