package main

import (
	"fmt"
	mrxcli "mrx/internal/cli"
	"mrx/internal/commands"
	"mrx/internal/io"
	"mrx/internal/marks"
	"os"
	"path/filepath"
	"strings"

	"github.com/alecthomas/kong"
	kongcompletion "github.com/jotaen/kong-completion"
)

var cli mrxcli.CLI

func handleErr(err error) {
	fmt.Println("Couldn't run command: ", err)
}

func fpAbsExisting(path string) (string, error) {
	absPath, err := filepath.Abs(path)
	if err != nil {
		return "", err
	}

	_, err = os.Stat(absPath)
	if err != nil {
		return "", err
	}
	return absPath, nil
}

func main() {
	app := kong.Must(&cli)

	kongcompletion.Register(app, kongcompletion.WithPredictor("labels", mrxcli.LabelPredictor))

	ctx, err := app.Parse(os.Args[1:])
	if err != nil {
		return
	}

	err = io.EnsureBookmarkdirExists()
	if err != nil {
		return
	}

	file, err := io.OpenMarksFile()
	if err != nil {
		return
	}
	defer file.Close()

	savedMarks := io.DecodeMarkFile(file)

	switch ctx.Command() {
	case "path <label>":
		var mark marks.Mark
		for i := range savedMarks {
			if savedMarks[i].Label == cli.Path.Label {
				mark = savedMarks[i]
				break
			}
		}

		fmt.Println(mark.Path)
	case "show":
		path, err := os.Getwd()
		if err != nil {
			handleErr(err)
			return
		}

		absPath, err := fpAbsExisting(path)
		if err != nil {
			handleErr(err)
			return
		}

		for i := range savedMarks {
			if savedMarks[i].Path == absPath {
				savedMarks[i].Print()
			}
		}
	case "show <path>":
		absPath, err := fpAbsExisting(cli.Show.Path)
		if err != nil {
			handleErr(err)
			return
		}

		for i := range savedMarks {
			if savedMarks[i].Path == absPath {
				savedMarks[i].Print()
			}
		}

	case "go <label>":
		zshScript := `mrx() {
    if [ "$1" = "go" ]; then
        shift
        cd "$(command mrx path "$@")"
    else
        command mrx "$@"
    fi
}`
		fmt.Printf("To use the `mrx go <label>` command, add the following code to your .zshrc, or whatever shell you're using:\n%s", zshScript)
	case "ls":
		for i := range savedMarks {
			savedMarks[i].Print()
		}
	case "add <label> <path>":

		absPath, err := fpAbsExisting(cli.Add.Path)
		if err != nil {
			handleErr(err)
			return
		}

		newmark := marks.NewMark(cli.Add.Label, absPath)
		savedMarks, err = commands.AddMark(savedMarks, newmark, cli.Add.Force)
		if err != nil {
			handleErr(err)
			return
		}

		io.EncodeMarkFile(file, savedMarks)

	case "add <label>":
		path, err := os.Getwd()
		if err != nil {
			handleErr(err)
			return
		}
		absPath, err := fpAbsExisting(path)

		if err != nil {
			handleErr(err)
			return
		}

		newmark := marks.NewMark(cli.Add.Label, absPath)
		savedMarks, err = commands.AddMark(savedMarks, newmark, cli.Add.Force)
		if err != nil {
			handleErr(err)
			return
		}

		io.EncodeMarkFile(file, savedMarks)
	case "rm <label>":
		savedMarks, err = commands.RemoveMark(savedMarks, cli.Rm.Label)
		if err != nil {
			handleErr(err)
			return
		}

		io.EncodeMarkFile(file, savedMarks)
	case "rn <old_label> <new_label>":
		savedMarks, err = commands.RenameMark(savedMarks, cli.Rn.OldLabel, cli.Rn.NewLabel)
		if err != nil {
			handleErr(err)
			return
		}
		io.EncodeMarkFile(file, savedMarks)

	default:
		if strings.HasPrefix(ctx.Command(), "completion") {
			err := cli.Completion.Run(ctx)
			if err != nil {
				panic(err)
			}
		}
		panic(ctx.Command())
	}
}
