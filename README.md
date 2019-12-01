# go-ghq-alfred

![](https://github.com/pddg/go-ghq-alfred/workflows/test/badge.svg?branch=master)

Search local repos with ghq in Alfred Workflow.

## Environment

* Alfred 3.4 (or later)
* [ghq](https://github.com/motemen/ghq)

## Usage

This tool is a CLI tool. Output JSON strings.

```bash
$ ./go-alfred-ghq '{query}' $(ghq list -p)
```

## In Alfred

This workflow start with `ghq {query}` in alfred.  

Filtering the result of `ghq list -p` with `{query}` and show them.

### Preparing

You should specify path to `ghq`. Open this workflow settings, and edit environment variables. Default is `/usr/local/bin/ghq`.  

And I recommend you to specify a editor and terminal app. Default is `Visual Studio Code.app` and `iTerm.app`

### Modifier key options

* **Enter**: Open repository in Finder.
* **Shift + Enter**: Open repository in your default browser.
* **Command + Enter**: Same as Enter only.
* **Option + Enter**: Search "user/repo" in google.
* **Fn + Enter**: Open repository in your terminal.
* **Control + Enter**: Open repository in your editor.

## Build

My environment is as follows.

* Go 1.9
* Glide 0.12.3

```bash
$ glide update
$ go test $(glide novendor) -v
$ go build .
```

## Attributes

Icons provided by www.flaticon.com.

### github and git logo

Icon made by Freepik from www.flaticon.com

### bitbucket logo

Icon made by Swifticons from www.flaticon.com

## Author

pudding

## License

MIT