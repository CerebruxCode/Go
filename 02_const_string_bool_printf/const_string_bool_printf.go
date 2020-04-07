package main

import "fmt"

/*
Νέοι Όροι που θα δούμε:
-- var: variable, μεταβλητή. Μια μεταβλητή είναι ένα στοιχείο με όνομα της αρεσκείας μας που αποθηκεύει μια τιμή (αριθμοί - γράμματα)
-- const: constant, σταθερά.
-- string: είναι ο τύπος που αφορά λέξεις - σύμβολα
-- escape symbols: Ειδικά σύμβολα που προσθέτουν η αγνοούν κάποια κατάσταση
-- boolean: bool, λογικοί τελεστές που εξετάζουν αν μια συνθήκη είναι αληθής (true) η ψευδής (false)
-- Printf: Σε αντίθεση με την Println, η Printf χρησιμοποιείται όταν θέλουμε να τυπώνουμε το αποτέλεσμα με συγκεκριμένο τρόπο
*/

func main() {

	// Ένα constant είναι ενα variable η τιμή του οποίου δεν μπορεί να αλλάξει

	const pi float64 = 3.14159265359

	// Μπορείτε να δηλώσετε πολλαπλά variable όπως παρακάτω

	var (
		varA = 2
		varB = 3
	)
	// Ας δούμε τι είναι αποθηκευμένο στις variable που δηλώσαμε:
	fmt.Println(varA, varB)

	var myName string = "Cerebrux"

	// Το μήκος ενός string το λαμβάνουμε με τη function len(string)

	fmt.Println(len(myName))

	// Μπορείτε να ενώσετε με +

	fmt.Println(myName + " : Ιστοσελίδα τεχνολογίας ")

	// Τα Strings εντός των " " μπορούν να περιλαμβάνουν escape symbols όπως το \n για να τυπώσουν μια νέα γραμμή

	fmt.Println("Μου αρέσουν \n")

	fmt.Println("Νέες γραμμές")

	// Τα Booleans είναι τύποι δεδομένων αλήθειας και μπορούν να πάρουν δύο τιμές, true ή false

	var isOver30 bool = true

	fmt.Println(isOver30)

	// Στην Printf η ένδειξη %f είναι για αριθμούς float
	// Ενώ με την %.xf (όπου χ αριθμός) μπορείτε να ορίσετε και την ακρίβεια του δεκαδικού ενός float

	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)

	// Η %T μας τυπώνει τον τύπο του variable για να ελέγξουμε π.χ. ως τι είναι αποθηκευμένος

	fmt.Printf("%T \n", pi)

	// Η %t τυπώνει τα booleans

	fmt.Printf("%t \n", isOver30)

	// %d χρησιμοποιείται για integers

	fmt.Printf("%d \n", 100)

	// %b τυπώνει σε binary μορφή

	fmt.Printf("%b \n", 100)

	// %c τυπώνει τον χαρακτήρα που σχετίζεται με το keycode

	fmt.Printf("%c \n", 44)

	// %x τυπώνει σε hexcode

	fmt.Printf("%x \n", 17)

	// %e τυπώνει με επιστημονικό συμβολισμό

	fmt.Printf("%e \n", pi)

}
