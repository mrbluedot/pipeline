package main

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

func main() {
	fmt.Println("Hello Go!!")

	data := [][]string{
		[]string{"cat", "Tom", "100", "13"},
		[]string{"dog", "Pluto", "22", "13"},
		[]string{"mouse", "Rat", "5", "13"},
	}

	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader([]string{"Type", "Name", "Power", "Age"})
	table.AppendBulk(data)
	table.Render()
}
