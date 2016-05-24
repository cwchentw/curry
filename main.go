package main

import (
	"fmt"
	"github.com/cwchentw/libcurry"
	"log"
	"os"
)

func usage(cmd string) {
	fmt.Println("Usage: " + cmd + " <command> [<args>]")
	fmt.Println("Available commands:")
	fmt.Println("  init      Initialize setttings")
	fmt.Println("  update    Update data")
	fmt.Println("  list      List available currencies")
	fmt.Println("  clear     Clear both config and data")
	fmt.Println(`  favor CURRENCY
            Set base currency to CURRENCY`)
	fmt.Println(`  from CUR AMOUNT
            Convert AMOUNT from CUR to base currency`)
	fmt.Println(`  from CUR1 to CUR2 AMOUNT
            Convert AMOUNT from CUR1 to CUR2`)
}

func main() {
	if len(os.Args) == 1 {
		usage(os.Args[0])
		os.Exit(0)
	}

	initFlag := false
	listFlag := false
	updateFlag := false
	clearFlag := false
	favorFlag := false
	fromToFlag := false
	switch os.Args[1] {
	case "init":
		initFlag = true
	case "update":
		updateFlag = true
	case "list":
		listFlag = true
	case "clear":
		clearFlag = true
	case "favor":
		favorFlag = true
	case "from":
		fromToFlag = true
	default:
		log.Fatalln(os.Args[1] + " is not a valid command.\n")
	}

	if initFlag {
		err := libcurry.InitSetting()
		if err != nil {
			log.Fatalln(err)
		}
	}

	if updateFlag {
		err := libcurry.UpdateData()
		if err != nil {
			log.Fatalln(err)
		}
	}

	if listFlag {
		err := libcurry.ListCurrencies()
		if err != nil {
			log.Fatalln(err)
		}
	}

	if clearFlag {
		err := libcurry.Clear()
		if err != nil {
			log.Fatalln(err)
		}
	}

	if favorFlag {
		if len(os.Args[1:]) != 2 {
			log.Fatalln("Wrong parameters")
		}

		err := libcurry.SetFavor(os.Args[2])
		if err != nil {
			log.Fatalln(err)
		}
	}

	if fromToFlag {
		err := libcurry.ParseFromCmd(os.Args)
		if err != nil {
			log.Fatalln(err)
		}
	}
}
