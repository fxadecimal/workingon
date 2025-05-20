# Working On It 

Solution: A go-based personal log file

Problem: I've often have 2-6 projects on the go that I need to dip in and out of

## Usage:

```sh
Usage: wko log [message]
       wko ls [-n N] | list [-n N]
       wko last
       wko path
```


## Quickstart

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

## Build


```sh
make
```

## Test:


```sh
go test
```

## Installation


***nix**

```sh
cp bin/osx/wko /usr/local/bin/
```

## Todo:

- best testing
- config file
- built in grep for current directory
- better logging file (sqlite?)

