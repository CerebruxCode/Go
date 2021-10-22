/*
Τι θα μάθουμε
- Θα δούμε τι είναι και πως φτιάχνουμε maps
- πως προσθέτουμε ή και αφαιρούμε στοιχεία
- να αποθηκεύουμε μια τιμή του map σε μεταβλητή
- να κάνουμε αναζήτηση αν υπάρχει κάποιο στοιχείο στο map
- Εντολές, make, delete,
*/

package main

import "fmt"

func main() {

	/*
		Maps - Χάρτες

		Τα "Map" είναι μία δομή δεδομένων η οποία αποθηκεύει
		ζευγάρια της μορφής "κλειδί,τιμή" όπως μια βάση δεδομένων

		Εναλλακτικά ονόματά του είναι
		 - πίνακας κατακερματισμού (hash table)
		 - ή λεξικό (dictionary)
		 - ή χάρτης/συσχετιστικός πίνακας (map/associative array)

		 είναι μία πιο ευέλικτη παραλλαγή της εγγραφής δεδομένων, στην οποία
		 τα ζευγάρια κλειδιού-τιμής μπορούν να εισαχθούν και να διαγραφούν ελεύθερα.

		οι χάρτες χρησιμοποιούνται για να αναζητήσετε μια
		τιμή από το αντίστοιχο κλειδί της (σαν μια μικρή βάση δεδομένων).

		Η μηδενική τιμή ενός χάρτη είναι "nil". Ένας "nil" χάρτης
		δεν έχει κλειδιά, ούτε μπορούν να προστεθούν κλειδιά.

			var name map[string]int

	*/

	/*
		Έστω ότι θέλουμε ένα map xartis

			var parking map[string]int
			fmt.Println(parking)

			xartis["Mercedes"] = 15
			fmt.Println(xartis)

		Ο κώδικας αυτός όμως θα σας εμφανίσει σφάλμα στην μεταγλώττιση
		διότι τα maps πρέπει να αρχικοποιηθούν πριν μας επιτραπεί
		να τα χρησιμοποιήσουμε
	*/
	// Άρα θα χρησιμοποιήσουμε την make για να αρχικοποιήσουμε ένα map μας
	presAge := make(map[string]int) // Δημιουργούμε ένα map με keyType string
	// και valueType intiger
	// varName := make(map[keyType] valueType)

	presAge["Molybdenum"] = 42 // Η εντολή αυτή είναι παρόμοια με αυτή που
	// είδαμε στα arrays, αλλά το κλειδί αντί να είναι ένας intiger (ακέραιος),
	// είναι ένα string επειδή ο τύπος του κλειδιού του map είναι string.

	fmt.Println(presAge["Molybdenum"])
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	fmt.Println(len(presAge)) // Λήψη αριθμού των στοιχείων στο χάρτη

	presAge["Technetium"] = 43 // Το μέγεθος του "presAge" map
	// αλλάζει όταν προστίθεται ένα νέο στοιχείο
	fmt.Println("Νεο μήκος = ", len(presAge))
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	delete(presAge, "Technetium") // Μπορούμε να διαγράψουμε
	// χρησιμοποιώντας την function "delete()"
	fmt.Println("Νεο μήκος = ", len(presAge))
	fmt.Println("Τι έχουμε εδώ: ", presAge)

	fmt.Println("\n== myCustomMap ==")
	myCustomMap()
}

func myCustomMap() {

	m := make(map[string]int) // δημιουργούμε έναν
	// nil ("άδειο") Map, έτοιμο για χρήση
	fmt.Println("Η αρχική τιμή είναι:", m["Answer"]) // άδειος άρα 0

	m["Answer"] = 42 // παίρνει τιμή 42
	fmt.Println("Η νέα τιμή είναι:", m["Answer"])

	m["Answer"] = 48 // Αλλάζω το 42 σε 48 χρησιμοποιώντας την ίδια θέση "Answer"
	fmt.Println("Η αλλαγμένη τιμή είναι:", m["Answer"])

	timi := m["Answer"] // αντιγράφω την τιμή σε
	// μια νεα μεταβλητή "timi"
	fmt.Println("Η τιμή της timi είναι:", timi)
	fmt.Println("Το διπλάσιο της timi είναι: ", timi+timi) // μπορούμε να  96

	delete(m, "Answer") // η πρώτη παράμετρος πρέπει
	// να είναι Map η δεύτερη το Κλειδί
	fmt.Println("Μετά την διαγραφή η τιμή του map είναι:", m["Answer"]) // διαγράφεται το 48, αρα 0

	fmt.Println("\n== Ένα map τύπου int ")
	intMap := make(map[int]int)
	fmt.Println("Αρχική τιμή του intMap:", intMap[0])
	intMap[0] = 5
	fmt.Println("Νεα τιμή του intMap:", intMap[0])

	/*
		Ελέγξτε ότι υπάρχει ένα κλειδί με μια εκχώρηση δύο τιμών:

				elem, ok = m[key]

		Αν το κλειδί είναι σε m, είναι αληθές.
		Αν όχι, το ok είναι ψευδές.

		Αν το κλειδί δε βρίσκεται στο map,
		τότε το elem είναι η μηδενική τιμή για τον τύπο του στοιχείου του χάρτη.
	*/
	fmt.Println("\n== Αναζήτηση στοιχείων ==")
	elem, ok := m["Answer"]
	fmt.Println("Η τιμή:", elem, "Υπάρχει;", ok)

	// Μπορούμε επίσης να αποθηκεύσουμε πολλαπλά στοιχεία σε ένα χάρτη
	fmt.Println("\n== Πολλαπλά στοιχεία ==")
	// Φτιάχνουμε μια μικρή "βάση δεδομένων" με όνομα :
	// 		superhero
	// Η δομή της θυμοίζει με πίνακα με σειρές και στήλες :

	/*
				__Superman_____Batman_____Wonder Woman__
		realname|Clark Kent |Bruce Wayne |Diana Prince |
		city	|Metropolis |Gotham City |Themyscira   |

	*/

	superhero := map[string]map[string]string{
		// Ο εξωτερικός χάρτης χρησιμοποιείται ως ένας πίνακας αναζήτησης
		// με βάση το όνομα του ήρωα, ενώ οι εσωτερικοί χάρτες χρησιμοποιούνται
		// για την αποθήκευση γενικών πληροφοριών σχετικά με τους ήρωες

		"Superman": {
			"realname": "Clark Kent",
			"city":     "Metropolis",
			"origin":   "Krypton", // τα maps δεν είναι	διαδοχικά
			// οπότε τα στοιχεία δεν αποθηκεύονται με αριθμητική σειρά.
			// Εδώ τα String θα προβληθούν με αλφαβητική
			// σειρά του ονόματος του κλειδιού
			// (πρώτα το city, μετά το origin, τέλος το realname)
		},

		"Batman": {
			"realname": "Bruce Wayne",
			"city":     "Gotham City",
		},

		"Wonder Woman": {
			"realname": "Diana Prince",
			"city":     "Themyscira",
		},
	}

	// Μπορούμε να εξάγουμε δεδομένα όπου το κλειδί ταιριάζει με το Superman
	fmt.Println("Η δομή του map Superhero είναι: ", superhero)

	// Αν ψάξουμε κάποιον ήρωα που δεν υπάρχει :
	fmt.Println(superhero["Flash"]) // θα εμφανίσει απλά οτι είναι ένα κενό map
	/*
		Εδώ το map θα επέστρεφε την τιμή μηδέν για τον
		τύπο της τιμής αλλά επειδή το map μας περιλαμβάνει strings
		τότε είναι το κενό string.

		Παρόλο που θα μπορούσαμε να ελέγξουμε τη μηδενική τιμή
		στην περίπτωση μας (superhero["Flash"] == "")
		η Go παρέχει έναν καλύτερο τρόπο:

		temp, ok := superhero["Flash"]
		fmt.Println(temp, ok)
	*/
	// Η πρόσβαση σε ένα στοιχείο ενός χάρτη μπορεί
	// να επιστρέψει δύο τιμές αντί για μια. Η πρώτη τιμή
	// είναι το αποτέλεσμα της αναζήτησης, η δεύτερη μας λέει
	// αν ήταν ή όχι η αναζήτηση επιτυχής

	if element, ok := superhero["Superman"]; ok { // πρώτα παίρνουμε την τιμή
		// απο το map, στη συνέχεια αν είναι επιτυχείς τρέχει ο παρακάτω κώδικας
		fmt.Println("ψάξατε για", element, "η αναζήτηση ήταν επιτυχής;", ok)
	}
	// Εναλλακτικά μπορούμε να εμφανίσουμε απευθείας τα στοιχεία
	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["realname"], "of", temp["city"])
	}

}
