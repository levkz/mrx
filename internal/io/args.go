package io

const (
	GO_CMD     = "go"
	ADD_CMD    = "add"
	REMOVE_CMD = "rm"
	RENAME_CMD = "rn"
	LIST_CMD   = "ls"
)

type LabelArgs struct {
	Command string
	Label   string
}

type MarkArgs struct {
	Command string
	Label   string
	Path    string
}

type RenameArgs struct {
	Command  string
	OldLabel string
	NewLabel string
}

type ListArgs struct {
	Command string
}

type GoArgs LabelArgs
type RemoveArgs LabelArgs

type AddArgs MarkArgs

func ParseGoArgs(args []string) GoArgs {
	return GoArgs{args[0], args[1]}
}

func ParseRemoveArgs(args []string) RemoveArgs {
	return RemoveArgs{args[0], args[1]}
}

func ParseAddArgs(args []string) AddArgs {
	return AddArgs{args[0], args[1], args[2]}
}

func ParseListArgs(args []string) ListArgs {
	return ListArgs{args[0]}
}

func ParseRenameArgs(args []string) RenameArgs {
	return RenameArgs{args[0], args[1], args[2]}
}
