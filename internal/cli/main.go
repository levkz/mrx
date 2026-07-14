package cli

import (
	"mrx/internal/io"
	"sort"

	kongcompletion "github.com/jotaen/kong-completion"
	"github.com/posener/complete"
)

type CLI struct {
	Path struct {
		Label string `arg:"" help:"Label for a predefined location. see 'mrx ls' and 'mrx add'." completion-predictor:"labels"`
	} `cmd:"" help:"return the path to the labeled location."`
	Show struct {
		Path string `arg:"" optional:"" help:"Path of a location you want to see shortcuts to."`
	} `cmd:"" help:"return the shortcuts created for specified path."`
	Go struct {
		Label string `arg:"" help:"Label for a predefined location. see 'mrx ls' and 'mrx add'." completion-predictor:"labels"`
	} `cmd:"" help:"cd into the labeled location."`
	Ls  struct{} `cmd:"" help:"List all of the shortcuts (labels and their paths."`
	Add struct {
		Label string `arg:"" help:"Label for a predefined location. see 'mrx ls' and 'mrx go'."`
		Path  string `arg:"" optional:"" help:"Path to a location you want to create a shortcut to."`
		Force bool   `name:"force" short:"f" help:"replace the labeled location if it already exists (act as rename command)."`
	} `cmd:"" help:"Create a new labeled location."`
	Rm struct {
		Label string `arg:"" help:"Label for a predefined location. see 'mrx ls' and 'mrx add'." completion-predictor:"labels"`
	} `cmd:"" help:"Remove labeled location from memory."`
	Rn struct {
		OldLabel string `arg:"" help:"Old label for a predefined location. see 'mrx ls' and 'mrx add'." completion-predictor:"labels"`
		NewLabel string `arg:"" help:"New label for a predefined location. see 'mrx ls' and 'mrx add'."`
	} `cmd:"" help:"Rename shortcut with a <old_label> to a <new_label>"`
	Completion kongcompletion.Completion `cmd:"" help:"Generate shell completions."`
}

var LabelPredictor = complete.PredictFunc(func(a complete.Args) []string {
	file, err := io.OpenMarksFile()
	if err != nil {
		return nil
	}
	defer file.Close()

	marks := io.DecodeMarkFile(file)

	labels := make([]string, 0, len(marks))
	for _, mark := range marks {
		labels = append(labels, mark.Label)
	}

	sort.Strings(labels)
	return labels
})
