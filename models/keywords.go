package models

import "fmt"

type Keywords struct {
	ID   int
	Word string
}

func (p *Keywords) Insert() {
	fmt.Println("未実装")
}
