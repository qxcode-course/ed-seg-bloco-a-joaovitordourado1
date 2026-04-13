package main

import (
	"fmt"
	"math/rand"
)

func square(p *Pen, size float64) {
	if size < 4 {
		return
	}
	p.Up()
	p.Walk(size / 2)
	p.Right(90)
	p.Walk(size / 2)
	p.Down()
	for range 4 {
		p.Right(90)
		square(p, size*0.45)
		p.Walk(size)
	}
	p.Up()

	p.Walk(-size / 2)
	p.Left(90)
	p.Walk(-size / 2)

	p.Down()
}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	size := 500.0
	pen := NewPen(1000, 1000) // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(500, 500) // move o pincel para x 250, y 500
	pen.SetHeading(90)
	pen.SetLineWidth(1)
	pen.Down()
	square(pen, size)

	pen.SavePNG("square.png")
	fmt.Println("PNG file created successfully.")
}
