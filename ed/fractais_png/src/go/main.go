package main

import (
	"fmt"
	"math/rand"
)

func circle(p *Pen, size float64) {
	p.DrawCircle(size)

}

func randInt(min, max int) int {
	return min + rand.Intn(max-min+1)
}

func main() {
	pen := NewPen(500, 500)   // cria um canvas de 500 de largura por 500 de altura
	pen.SetRGB(255, 0, 0)     // muda a cor do pincel para vermelho
	pen.SetPosition(250, 250) // move o pincel para x 250, y 500
	circle(pen, 100.0)
	pen.SavePNG("circle.png")
	fmt.Println("PNG file created successfully.")
}
