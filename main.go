package main

import "fmt"

// Shape adalah interface untuk semua bentuk geometri
type Shape interface {
	Draw()
}

// Circle adalah struct untuk lingkaran
type Circle struct{}

// Draw menampilkan informasi tentang lingkaran
func (c *Circle) Draw() {
	fmt.Println("Menggambar lingkaran")
}

// Square adalah struct untuk persegi
type Square struct{}

// Draw menampilkan informasi tentang persegi
func (s *Square) Draw() {
	fmt.Println("Menggambar persegi")
}

// ShapeFactory adalah factory untuk membuat objek bentuk
type ShapeFactory struct{}

// CreateShape membuat objek bentuk berdasarkan jenis yang diberikan
func (sf *ShapeFactory) CreateShape(kind string) Shape {
	switch kind {
	case "circle":
		return &Circle{}
	case "square":
		return &Square{}
	default:
		return nil
	}
}

func main() {
	// Membuat factory bentuk
	factory := &ShapeFactory{}

	// Membuat beberapa jenis bentuk
	circle := factory.CreateShape("circle")
	square := factory.CreateShape("square")

	// Menggambar setiap bentuk
	circle.Draw()
	square.Draw()
}
