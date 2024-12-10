package main

import (
	"fmt"
	"plugin"
)

func main() {
	p, err := plugin.Open("statisticsPlugin.so")
	if err != nil {
		fmt.Println("Error Loading Plugin")
	}

	statisticsSymbol, err := p.Lookup("GetDiskStatisitcs")
	if err != nil {
		fmt.Println("Error Looking Up For Symbol")
	}

	statisticsFunction := statisticsSymbol.(func() string)

	diskStatisitcs := statisticsFunction()
	fmt.Println("Statistics From Plugin", diskStatisitcs)
}
