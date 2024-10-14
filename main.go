package main

import (
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fggr := color.New(color.FgGreen)
	bggr := color.New(color.BgGreen)

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
}
