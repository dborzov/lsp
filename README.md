## lsp: list files in a mildly human-frendlier manner

`lsp` lists files, like [`ls`](http://en.wikipedia.org/wiki/Ls) command,
but it does not attempt to meet
that archaic POSIX specification, so instead of this:
```
(bash)$ ls -l

total 16
-rw-r--r--  1 peterborzov  staff  1079  9 Aug 00:22 LICENSE
-rw-r--r--  1 peterborzov  staff    60  9 Aug 00:22 README.md
```

you get this:
![screenshot](https://raw.githubusercontent.com/dborzov/lsp/screenshots/symlinks.png)

## Features
#### File Groups
Files grouped by type (with `-l` key or in modes when file type not shown). `lsp` distinguishes binary, text and executable files, symlinks and is aware of weird types like devices and unix socket thingy:
![lsp can show files grouped by type](https://raw.githubusercontent.com/dborzov/lsp/screenshots/grouped.png)
#### Modification time in human-friendly format
`-t` key for when you are interested in modification time. It turns to the mode that makes most sense to me when I want to look up modtimes, sorted within file groups from recent to latest:
![](https://raw.githubusercontent.com/dborzov/lsp/screenshots/modtime.png)
Sometimes relative times are  not very readible as well (like when you are interested in a specific date), use two flags `-sl` to show the full UTC timestamp in properties.
#### Size in human-friendly format
`-s` key, similarly to modtime key, shows file sizes and sorts within file groups from largest to smallest:
![](https://raw.githubusercontent.com/dborzov/lsp/screenshots/size.png)

#### Align by left
I have been playing with aligning files and descriptions by center, and I like that you can see files with the same extension right away, but there are deifinitely cases when it gets weird.
For now, there is `-p` key to render the file table in the left-aligned columns:
![](https://raw.githubusercontent.com/dborzov/lsp/screenshots/table.png)


## Todo before v1.0
- Think about how to represent file rights and ownership
- too long filenames break things
- Show executable files
- Add tests once I settle on functionality
- Think of TODO list points


## Installation

`lsp` is written in `go` programming language.
For now it can be installed using `go get`:

```
 $ go get github.com/dborzov/lsp
```
Once it becomes more functional, `lsp` will be distributed in native binaries
for all platforms (Linux, MacOS, Windows). No dependancies or anything configurable by design, just one binary.

## Misc
MIT license.

Github Issues and pull requests are very welcome, feel free to [message me](tihoutrom@gmail.com) if you are considering contributing.

This tool is named after Lumpy Space Princess(LSP), a very quotable character from the TV show "Adventure Time with Finn and Jake".

![can't handle these lumps](beans.gif)
