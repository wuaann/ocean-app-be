package uploadmodel

import (
	"errors"
	"ocean-app-be/common"
)

const EntityName = "Upload"

type Upload struct {
	common.PSModel `json:",inline"`
	common.Image   `json:",inline"`
}

func (Upload) TableName() string {
	return "uploads"
}

func ErrFileIsNotImage(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"file is not image",
		"ErrFileISNotImage")
}

func ErrCannotSaveFile(err error) *common.AppError {
	return common.NewCustomError(
		err,
		"cannot save uploaded file",
		"ErrCannotSaveFile",
	)
}

var (
	ErrFileTooLarge = common.NewCustomError(
		errors.New("file too large"),
		"file too large",
		"ErrFileLarge",
	)
)
