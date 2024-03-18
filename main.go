package main

import (
	"fmt"

	"github.com/dzrise/go-bitrix-faset/element"
)

func main() {
	fmt.Println("Start Bitrix Reindex Faset Service")
	element.GetAllElements()
}
