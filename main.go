package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
	"golang.org/x/term"
)

func main() {
	//Clear the screen
	exec.Command("clear").Run()

	//Define colors
	fggr := color.New(color.FgGreen)
	bggr := color.New(color.BgGreen, color.FgBlack)

	bggr.Println("LOJISTIK BIRSEYLER")

	fmt.Print("   ")
	bggr.Println("BIRSEY MENU")

	fggr.Println("1.TESLIM ALMA")
	fggr.Println("2.TRANSFER")
	fggr.Println("3.GÖREVLER")
	fggr.Println("4.SAYIM MENU")
	fggr.Println("5.KONTROLLER")
	fggr.Println("6.SORTING")
	fggr.Println("7.SHIP EVRAK")
	fmt.Println()
	fggr.Println("8.ETIKETLER")
	fggr.Println("9.SHIP MENU")
	fggr.Println("0.E-MENU")
	fggr.Println("SEÇİM:")
	fggr.Println("________________")

	fggr.Println()

	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		fmt.Println(err)
		return
	}
	var b []byte
	os.Stdin.Read(b)
	defer term.Restore(int(os.Stdin.Fd()), oldState)
}
