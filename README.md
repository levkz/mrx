# mrx

`mrx` is a simple command-line bookmark manager for directories. It lets you assign labels to frequently used locations and quickly jump back to them.

Bookmarks are stored in:

```text
~/.bookmarkz/marks.json
```

## Features

- Add bookmarks
- Remove bookmarks
- Rename bookmarks
- List all bookmarks
- Jump to bookmarked directories

## Installation

Clone the repository and install the binary:

```bash
git clone <repository-url>
cd mrx
go install .
```

Ensure your Go binary directory is on your `PATH`:

```bash
export PATH="$HOME/go/bin:$PATH"
```

You can verify the installation:

```bash
which mrx
```

## Commands

### Add a bookmark

```bash
mrx add work ~/Projects/work
```

or from inside a directory:

```bash
mrx add work .
```

Relative paths are resolved to their absolute paths before being stored.

### List bookmarks

```bash
mrx ls
```

Example output:

```text
Label: work    Path: /home/alice/Projects/work
Label: notes   Path: /home/alice/Documents/notes
```

### Rename a bookmark

```bash
mrx rename work office
```

### Remove a bookmark

```bash
mrx remove office
```

## Jumping to a bookmark

A CLI application cannot directly change the working directory of the shell that launched it. To make `mrx go` behave like a normal `cd`, add the following wrapper to your `~/.zshrc`:

```zsh
mrx() {
    if [[ "$1" == "go" ]]; then
        shift
        cd "$(command mrx go "$@")"
    else
        command mrx "$@"
    fi
}
```

Reload your shell:

```bash
source ~/.zshrc
```

Now you can simply run:

```bash
mrx go work
```

and your shell will change into the bookmarked directory.

## Example

```bash
mrx add project .
mrx add downloads ~/Downloads

mrx ls

mrx go project
```

## Notes

The project was created not to be used, but for practice.
Look into other projects if you were looking for a better way to traverse through your file system:

- [zsh-z](https://github.com/agkozak/zsh-z) Jump quickly to directories that you have visited "frecently.".
- [zoxide](https://github.com/ajeetdsouza/zoxide) A smarter cd command.

## License

MIT
