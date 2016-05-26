package main

import (
	"fmt"
	"github.com/cwchentw/libcurry"
	"log"
	"os"
	"strconv"
	"time"
)

func ParseFromCmd(args []string) error {
	if len(args) == 4 {
		if args[1] != "from" {
			return libcurry.WrongArguments()
		}
		return fromToBase(args[2], args[3])
	} else if len(args) == 6 {
		if args[1] != "from" && args[3] != "to" {
			return libcurry.WrongArguments()
		}
		return fromTo(args[2], args[4], args[5])
	} else {
		return libcurry.WrongArguments()
	}
}

func fromToBase(cur string, amountStr string) error {
	configTree, err := libcurry.LoadConfigFile()
	if err != nil {
		return err
	}

	base := configTree.Get("base")
	if base == nil {
		return libcurry.NoBaseCurrency()
	}

	return fromTo(cur, base.(string), amountStr)
}

func fromTo(cur1 string, cur2 string, amountStr string) error {
	// Update data if modification time > 24 hours
	ratesPath, err := libcurry.GetCurrencyRatesPath()
	if err != nil {
		return err
	}

	info, err := os.Stat(ratesPath)
	if err != nil {
		return err
	}

	modTime := info.ModTime()
	delta := time.Now().Sub(modTime)
	if delta.Hours() > 24 {
		err := UpdateData()
		if err != nil {
			log.Println("Data older than 24 hours")
		}
	}

	rates, err := libcurry.ReadCurrencyRates()

	curFrom, ok := rates[cur1]
	if ok != true {
		return libcurry.UnknownCurrency()
	}

	curTo, ok := rates[cur2]
	if ok != true {
		return libcurry.UnknownCurrency()
	}

	amount, err := strconv.ParseFloat(amountStr, 64)
	if err != nil {
		return err
	}

	fmt.Printf("%.2f\n", amount*(curTo/curFrom))

	return nil
}
