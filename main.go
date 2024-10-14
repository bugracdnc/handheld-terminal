package main

import (
	"fmt"
	"os"
	"os/exec"

	"atomicgo.dev/cursor"
	"github.com/fatih/color"
)

var fggr = color.New(color.FgGreen)
var bggr = color.New(color.BgGreen, color.FgBlack)

var main_menu_headers = [2]string{"LOGISTIK BIRSEYLER", "BIRSEY MENU"}
var main_menu_items = [12]string{"1.TESLIM ALMA", "2.TRANSFER", "3.GÖREVLER", "4.SAYIM MENU", "5.KONTROLLER", "6.SORTING", "7.SHIP EVRAK\n", "8.ETIKETLER", "9.SHIP MENU", "0.E-MENU", "SEÇİM: ", "________________"}

func main_menu() string {
	valid_inputs := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	for {
		//Print headers
		bggr.Println(main_menu_headers[0])
		fmt.Print("   ")
		bggr.Println(main_menu_headers[1])

		//Print items
		for _, item := range main_menu_items {
			fggr.Println(item)
		}
		cursor.Up(2)
		fggr.Print(main_menu_items[len(main_menu_items)-2])

		exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

		var b []byte = make([]byte, 1)
		os.Stdin.Read(b)
		for _, validator := range valid_inputs {
			if string(b) == validator {
				return validator
			}
		}
	}
}

func main() {
	//Clear the screen
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

	main_menu()
}
