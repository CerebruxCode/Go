package main

import (
	"fmt"
	"math"
)

// STRUCTS AND INTERFACES

/*
Η Go δεν έχει κλάσεις (classes). Ωστόσο, μπορείτε να ορίσετε μεθόδους σε τύπους.
Μια μέθοδος είναι μια συνάρτηση με μία ειδική παράμετρο που είναι παραλήπτης.
Ο παραλήπτης εμφανίζεται στη δική του λίστα παραμέτρων μεταξύ της λέξης "func"
και του ονόματος της μεθόδου.
*/

// Vertex : κορυφή τρίγωνου
type Vertex struct {
	X, Y float64
}

// Shape : Σχήμα
type Shape interface {
	area() float64
}

// Rectangle : Ορθογώνιο
type Rectangle struct {
	height float64
	width  float64
}

// Circle : Κύκλος
type Circle struct {
	radius float64
}

// INTERFACES

// Μια διεπαφή (interface) ορίζει μια λίστα με τις μεθόδους που πρέπει να εφαρμόσει ένας τύπος (type)
// Εάν αυτός ο τύπος εφαρμόζει αυτές τις μεθόδους, η κατάλληλη μέθοδος εκτελείται
// ακόμα και το πρωτότυπο στο οποίο αναφέρεται είναι με το όνομα διεπαφής

// Abs : με την μέθοδο αυτή, ορίζουμε τι θα κάνουμε τα Χ, Υ που έχει ο τύπος Vertex που φτιάξαμε
// Abs : (Vertex.Abs) η μέθοδος Abs έχει έναν δέκτη τύπου Vertex που ονομάζεται v, άρα δέχεται
// την func Vertex με τις 2 τιμές που λαμβάνει
func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (r Rectangle) area() float64 {
	return r.height * r.width
}

func (c Circle) area() float64 {

	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64 {

	return shape.area()

}

func main() {

	v := Vertex{3, 4}                // εδώ αρχικοποιούμε την v που είχαμε δημιουργήσει επάνω
	fmt.Println("v.Abs()=", v.Abs()) // εδώ εφαρμόζουμε την Abs στις τιμές που έλαβε το v

	rect := Rectangle{20, 50}
	circ := Circle{4}

	fmt.Println("Rectangle Area =", getArea(rect))
	fmt.Println("Circle Area =", getArea(circ))

}
