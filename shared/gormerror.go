package shared

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

func ConvertGormErrorCode(err error) error {
	fmt.Println(err.Error())
	if errors.Is(err, gorm.ErrDuplicatedKey) {
		return NewError(AlreadyExists, "already exists")
	} else if errors.Is(err, gorm.ErrRecordNotFound) {
		return NewError(NotFound, "not found")
	}
	return err
}
