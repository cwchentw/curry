package main

import (
	"github.com/cwchentw/libcurry"
)

func UpdateData() error {
	if libcurry.IsInit() {
		err := InitSetting()
		if err != nil {
			return err
		}
	}

	err := libcurry.UpdateData()
	if err != nil {
		return err
	}

	return nil
}
