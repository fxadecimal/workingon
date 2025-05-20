# Working On it

A go-lang command line app that logs what you're working on

```sh
wko log building the app
# writes to ~/workingon.log
# outputs: 
> "building the app" [git:123456] (/some/path)

# show the last thing you logged
wko last 
# > "building the app" [git:123456] (/some/path)
who list # lists whole log
who list 3 # last 3 lines

```

You can pipe stuff into the log:

```sh
echo "howdy" | wko log "there"
# more useful: .git/hooks/post-commit
git log -1 --oneline | wko log

```

Usage:

```sh
Usage: wko log [message]
       wko ls | list - List Messages
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


