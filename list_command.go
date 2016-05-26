package main

import (
	"fmt"
	"github.com/cwchentw/libcurry"
	"sort"
)

func ListCurrencies() error {
	currencies, err := libcurry.ReadCurrencies()
	if err != nil {
		return err
	}

	keys := make([]string, 0)
	for k := range currencies {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		fmt.Println(k + " " + currencies[k])
	}

	return nil
}
