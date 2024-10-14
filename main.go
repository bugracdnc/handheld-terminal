package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/fatih/color"
)

func main() {
	//Clear the screen
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()

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

	fggr.Print("\033[2ASEÇİM: ")

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()

	var b []byte = make([]byte, 1)
	os.Stdin.Read(b)

	fmt.Println()
	fmt.Println()
}
