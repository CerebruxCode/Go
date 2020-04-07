package main

import "fmt"

func main() {

	/*
		Maps - Χάρτες

		Τα "Map" είναι μία δομή δεδομένων η οποία αποθηκεύει
		ζευγάρια της μορφής "κλειδί,τιμή"

		Εναλλακτικά ονόματά του είναι
		 - πίνακας κατακερματισμού (hash table)
		 - ή λεξικό (dictionary)
		 - ή χάρτης/συσχετιστικός πίνακας (map/associative array)

		 είναι μία πιο ευέλικτη παραλλαγή της εγγραφής δεδομένων, στην οποία
		 τα ζευγάρια κλειδιού-τιμής μπορούν να εισαχθούν και να διαγραφούν ελεύθερα.

		οι χάρτες χρησιμοποιούνται για να αναζητήσετε μια τιμή από το αντίστοιχο κλειδί της (σαν μια μικρή βάση δεδομένων).

		Η μηδενική τιμή ενός χάρτη είναι "nil". Ένας "nil" χάρτης
		δεν έχει κλειδιά, ούτε μπορούν να προστεθούν κλειδιά.

	*/

	presAge := make(map[string]int) // Δημιουργήθηκε με varName := make(map[keyType] valueType)

	presAge["Molybdenum"] = 42

	fmt.Println(presAge["Molybdenum"])
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	fmt.Println(len(presAge)) // Λήψη αριθμού των στοιχείων στο χάρτη

	presAge["Technetium"] = 43 // Το μέγεθος αλλάζει όταν προστίθεται ένα νέο στοιχείο
	fmt.Println("Νεο μήκος = ", len(presAge))
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	delete(presAge, "Technetium") // Μπορούμε να διαγράψουμε χρησιμοποιώντας την function "delete()"
	fmt.Println("Νεο μήκος = ", len(presAge))
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	myCustomMap()
}

func myCustomMap() {

	m := make(map[string]int)                 // δημιουργούμε έναν nil ("άδειο") Map, έτοιμο για χρήση
	fmt.Println("Η τιμή είναι:", m["Answer"]) // άδειος άρα 0

	m["Answer"] = 42 // παίρνει τιμή 42
	fmt.Println("Η τιμή είναι:", m["Answer"])

	m["Answer"] = 48
	timi := m["Answer"]                           // αντιγράφω την τιμή σε μια μεταβλητή "timi"
	fmt.Println("Η τιμή είναι:", timi)            // αλλάζει σε 48
	fmt.Println("Το διπλάσιο είναι: ", timi+timi) // αλλάζει σε 96

	delete(m, "Answer")                       // η πρώτη παράμετρος πρέπει να είναι Map η δεύτερη το Κλειδί
	fmt.Println("Η τιμή είναι:", m["Answer"]) // διαγράφεται το 48, αρα 0

	/*
		Ελέγξτε ότι υπάρχει ένα κλειδί με μια εκχώρηση δύο τιμών:

				elem, ok = m[key]

		Αν το κλειδί είναι σε m, είναι αληθές.
		Αν όχι, το ok είναι ψευδές.

		Αν το κλειδί δε βρίσκεται στο map,
		τότε το elem είναι η μηδενική τιμή για τον τύπο του στοιχείου του χάρτη.
	*/
	elem, ok := m["Answer"]
	fmt.Println("Η τιμή:", elem, "Υπάρχει;", ok)
}
