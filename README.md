> "Oh, my Glob, you guys, drama bomb!"

*[Lumpy Space Princess](http://adventuretime.wikia.com/wiki/Lumpy_Space_Princess) (LSP), "Adventure Time with Finn and Jake", Season 2, episode 1*




## lsp: list files in a mildly human-frendlier manner

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
 ( I will eventually need to put a sublime-style gif showcasing major features)
```
which is arguably more readable.

## Why
No reason to cling to stuff that is there only for historical reasons. Nobody has time in life to learn how to read this: `-rw-r--r--`.
## Features
Marked ones are implemented:
- [ ] embraces and extends the original `ls` syntax (with flags and all)
whenever it does not actively interfere with common sense.
- [ ] fuzzy matching and fixes basic typos: `lsp ~/.bahs_profile` will still get you there.
- [ ] contents are shown classified in groups intelligently: directories, `*.js`-files, MIME/content types (text, binaries..)
- [ ] shows intelligent summaries for objects: things like sizes and types of subfolders, file's encoding and so on.
General running timeout threshold means no freezes because of things like suddenly unmounted devices, huge number of files and so on.
Each file is "investigated" asyncroneously and subdirectories are traversed in the async [BFS](http://en.wikipedia.org/wiki/Breadth-first_search).
- [ ] knows what things like a git repo is. shows recent git diffs and all the other things.

## Installation

`lsp` is written in `go` programming language.
For now it can be installed using `go get`:

```
 $ go get github.com/dborzov/lsp
```
Once it becomes more functional, `lsp` will be distributed in native binaries
for all platforms (Linux, MacOS, Windows). No dependancies or anything configurable by design, only binary.

## Soft stuff
![can't handle these lumps](lumps.jpg)

MIT license :)

Would love to be pull requested :)
