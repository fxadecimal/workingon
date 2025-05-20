# Working On it

A go-lang command line app that logs what you're working on

```sh
wko log building the app
# writes to ~/workingon.log
# outputs:
> "building the app" [git:123456] (/some/path)
```

Usage:

```sh
Usage: wko log [message]
       wko ls - List Messages
       wko last - Last Message
       wko path - outputs log path
```


## Build

```sh
make
```


## Installation

Osx:

```sh
cp bin/osx/wko /usr/local/bin/
```


