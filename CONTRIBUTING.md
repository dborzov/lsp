Welcome! Here be a general repo overview to serve as an intro to a young lady or gentleman who is considering reading up on they things are set up around here and possibly contributing.

Conventions
-----
Everything is `go fmt` and `go lint` compliant.
All the major editors support plugins that enable doing this on the fly, they are very helpful.

Another useful tool that I found helpful is `goimports`.

At the moment I want to move to maximum modularity, so that each separate functionality domain is moved into its own separate package.
I have heard of people arguing that it is not sustainable but not yet convinced. Maybe this will fail and that would be my way to grok their argument.

Concepts
-----
###Mode

`mode *Mode`, defined in [arguments.go](arguments.go), stores all things involving the settings of the running instance:
```go
  type Mode struct {
    inputPath  string // target path literal from the argument parsing (e.g. "~/hi")
    absolutePath string // absolute path to the target directory (e.g. "/home/dima/hi")
    // ...
    // and ... boolean flags for specific modes
  }
```
We have target directory's path here and all the mode flags.

`ParseArguments()` parses the arguments passed on to `lsp` into mode flags and target dir.  There are many excellent libraries that parse flag arguments, I ended up rewriting my own to enable `ls`-style single letter triggers within one flag (so that `lsp -al` is equivalent to `lsp -a -l` or `lsp -la`).
