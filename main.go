package main

import (
	"fmt"
	"mrx/internal/commands"
	"mrx/internal/io"
	"mrx/internal/marks"
	"os"
	"path/filepath"
)

func handleErr(err error) {
	fmt.Println("Couldn't run command: ", err)
}

func main() {
	gargs := os.Args[1:]
	err := io.EnsureBookmarkdirExists()
	if err != nil {
		return
	}

	file, err := io.OpenMarksFile()
	if err != nil {
		return
	}
	defer file.Close()

	savedMarks := io.DecodeMarkFile(file)

	switch cmd := gargs[0]; cmd {

	case io.GO_CMD:
		args := io.ParseGoArgs(gargs)
		var mark marks.Mark
		for i := range savedMarks {
			if savedMarks[i].Label == args.Label {
				mark = savedMarks[i]
				break
			}
		}

		fmt.Println(mark.Path)
	case io.ADD_CMD:
		args := io.ParseAddArgs(gargs)
		absPath, err := filepath.Abs(args.Path)
		if err != nil {
			handleErr(err)
			return
		}

		_, err = os.Stat(absPath)
		if err != nil {
			handleErr(err)
			return
		}

		newmark := marks.NewMark(args.Label, absPath)
		savedMarks, err = commands.AddMark(savedMarks, newmark, false)
		if err != nil {
			handleErr(err)
			return
		}

		io.EncodeMarkFile(file, savedMarks)
	case io.REMOVE_CMD:
		args := io.ParseRemoveArgs(gargs)
		savedMarks, err = commands.RemoveMark(savedMarks, args.Label)
		if err != nil {
			handleErr(err)
			return
		}

		io.EncodeMarkFile(file, savedMarks)
	case io.RENAME_CMD:
		args := io.ParseRenameArgs(gargs)
		savedMarks, err = commands.RenameMark(savedMarks, args.OldLabel, args.NewLabel)
		if err != nil {
			handleErr(err)
			return
		}
		io.EncodeMarkFile(file, savedMarks)
	case io.LIST_CMD:
		for i := range savedMarks {
			fmt.Print("Label: ", savedMarks[i].Label, "\tPath:", savedMarks[i].Path, "\n")
		}
	}
}
