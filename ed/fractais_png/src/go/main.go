package main

import (
	"fmt"
	"math/rand"
)

func tri(p *Pen, size float64) {

	if size < 3 {

		for i := 0; i < 3; i++ {
			p.Walk(size)
			p.Right(120)
		}
		return
	}

	for i := 0; i < 3; i++ {
		tri(p, size/2.0)
		p.Walk(size)
		p.Right(120)
	}
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	size := 900.0
	pen := NewPen(1000, 870) // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(50, 50)     // move o pincel para x 250, y 500
	pen.SetHeading(0)
	pen.SetLineWidth(1)
	pen.Down()
	tri(pen, size)

	pen.SavePNG("triangulo.png")
	fmt.Println("PNG file created successfully.")
}
