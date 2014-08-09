> "Oh, my Glob, you guys, drama bomb!"

*[Lumpy Space Princess](http://adventuretime.wikia.com/wiki/Lumpy_Space_Princess) (LSP), Season 2, episode 1*




## lsp: list files in mildly human-frendlier manner





`lsp` is like [`ls`](http://en.wikipedia.org/wiki/Ls) command
but does not attempt to meet
that archaic POSIX specification, so instead of this:
```
(bash)$ ls -l

total 16
-rw-r--r--  1 peterborzov  staff  1079  9 Aug 00:22 LICENSE
-rw-r--r--  1 peterborzov  staff    60  9 Aug 00:22 README.md
```

you get this:

```
  $ lsp

  ... something fancy
  (actually I did not come up with what is it going to look like yet)
```
which is arguably more readable.

## Features

- shows reasonable summaries for subfolders
- timeout on running time so no hanging for remotely mounted devices and such
* knows what things like a git repo is
- so shows recent git diffs and things

## Installation

`lsp` is written in `go` and can be installed like a standard `go` program:

```
 $ go get github.com/dborzov/lsp
```

## Dedication

This tool is dedicated to the lifetime of achievement of my personal hero and muse, the one with the nicest LUMPS,
the [Lumpy Space Princess](http://adventuretime.wikia.com/wiki/Lumpy_Space_Princess) (LSP).

![can't handle these lumps](lumps.gif)
