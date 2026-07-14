package commands

import (
	"errors"
	"mrx/internal/marks"
)

func AddMark(marks []marks.Mark, newmark marks.Mark, force bool) ([]marks.Mark, error) {
	for i := range marks {
		if marks[i].Label == newmark.Label {
			if force {
				marks[i].Path = newmark.Path
				return marks, nil
			}
			return nil, errors.New("mark with that name already exists!")
		}
	}

	return append(marks, newmark), nil
}

func RenameMark(marks []marks.Mark, oldName string, newName string) ([]marks.Mark, error) {
	for i := range marks {
		if marks[i].Label == oldName {
			marks[i].Label = newName
			return marks, nil
		}
	}

	return nil, errors.New("mark with that name doesn't exist!")
}

func RemoveMark(marks []marks.Mark, removeMark string) ([]marks.Mark, error) {
	for i := range marks {
		if marks[i].Label == removeMark {
			return append(marks[:i], marks[i+1:]...), nil
		}
	}

	return nil, errors.New("mark with that name doesn't exist!")
}

func ReasignMark(marks []marks.Mark, newmark marks.Mark) ([]marks.Mark, error) {
	for i := range marks {
		if marks[i].Label == newmark.Label {
			marks[i].Path = newmark.Path
			return marks, nil
		}
	}
	return nil, errors.New("mark with that name doesn't exist!")
}
