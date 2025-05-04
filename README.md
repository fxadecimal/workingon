# Working On it

A go-lang command line app that logs what you're working on

```sh
wko log building the app
# writes to ~/workingon.log
# outputs:
> "building the app" [git:a023068] (/Users/neil/working_directory/coding/go-lang/go-workingonit)
```

Usage:

```sh
Usage: wko log [message]
       wko ls - List Messages
       wko last - Last Message
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


