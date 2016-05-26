package main

import (
	"errors"
	"github.com/cwchentw/libcurry"
)

func SetFavor(cur string) error {
	currencies, err := libcurry.ReadCurrencies()
	if err != nil {
		return err
	}

	_, ok := currencies[cur]
	if ok != true {
		return errors.New("Unknown currency in OpenExchangeRates")
	}

	configTree, err := libcurry.LoadConfigFile()
	if err != nil {
		return err
	}

	var _cur interface{} = cur
	configTree.Set("base", _cur)

	err = libcurry.WriteConfigFile(configTree)
	if err != nil {
		return err
	}

	return nil
}
