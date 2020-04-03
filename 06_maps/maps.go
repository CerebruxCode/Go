package main

import "fmt"

func main() {

	/*
		Τα "Map" είναι μία δομή δεδομένων η οποία αποθηκεύει
		ζευγάρια της μορφής "κλειδί,τιμή"
		Επίσης γνωστό ως συσχετιστικός πίνακας,
		πίνακας κατατεμαχισμού ή λεξικό, οι χάρτες χρησιμοποιούνται
		για να αναζητήσετε μια τιμή από το αντίστοιχο κλειδί του (σαν μια μικρή βάση δεδομένων).

		Η μηδενική τιμή ενός χάρτη είναι "nil". Ένας "nil" χάρτης
		δεν έχει κλειδιά, ούτε μπορούν να προστεθούν κλειδιά.

	*/

	// Δημιουργήθηκε με varName := make(map[keyType] valueType)

	presAge := make(map[string]int)

	presAge["Molybdenum"] = 42

	fmt.Println(presAge["Molybdenum"])
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	// Get the number of items in the Map

	fmt.Println(len(presAge))

	// The size changes when a new item is added

	presAge["Technetium"] = 43
	fmt.Println("Νεο μήκος = ", len(presAge))
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	// We can delete by passing the key to delete

	delete(presAge, "Technetium")
	fmt.Println("Νεο μήκος = ", len(presAge))
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	myCustomMap()
}

func myCustomMap() {
	// δημιουργούμε εναν nil ("άδειο"), έτοιμο για χρήση Map
	m := make(map[string]int)
	fmt.Println("Η τιμή είναι:", m["Answer"]) // άδειος άρα 0

	m["Answer"] = 42
	fmt.Println("Η τιμή είναι:", m["Answer"]) // παίρνει τιμή 42

	m["Answer"] = 48
	timi := m["Answer"]                           // αντιγράφω την τιμή σε μια μεταβλητή "timi"
	fmt.Println("Η τιμή είναι:", timi)            // αλλάζει σε 48
	fmt.Println("Το διπλάσιο είναι: ", timi+timi) // αλλάζει σε 48

	delete(m, "Answer")                       // η πρώτη παράμετρος πρέπει να είναι Map η δεύτερη το Κλειδί
	fmt.Println("Η τιμή είναι:", m["Answer"]) // διαγράφεται το 48, αρα 0

	/*
		Ελέγξτε ότι υπάρχει ένα κλειδί με μια εκχώρηση δύο τιμών:
		elem, ok = m[key]
		Αν το κλειδί είναι σε m, είναι αληθές.
		Αν όχι, το ok είναι ψευδές.

		Αν το κλειδί δεν βρίσκεται στον χάρτη,
		τότε το elem είναι η μηδενική τιμή για τον τύπο του στοιχείου του χάρτη.
	*/
	v, ok := m["Answer"]
	fmt.Println("Η τιμή:", v, "Present?", ok)
}
