package main

import (
	"fmt"
	"reflect" // έχει την TypeOf() function για να προσδιορίσουμε τον τύπο δεδομένων που αποθηκεύει μια μεταβλητή
)

/*
Νέοι Όροι που θα δούμε:
-- var: variable, μεταβλητή. Μια μεταβλητή είναι ένα στοιχείο με όνομα της αρεσκείας μας που αποθηκεύει μια τιμή (αριθμοί - γράμματα)
-- const: constant, σταθερά.
-- string: είναι ο τύπος που αφορά λέξεις - σύμβολα
-- escape symbols: Ειδικά σύμβολα που προσθέτουν η αγνοούν κάποια κατάσταση
-- boolean: bool, λογικοί τελεστές που εξετάζουν αν μια συνθήκη είναι αληθής (true) η ψευδής (false)
-- Printf: Σε αντίθεση με την Println, η Printf χρησιμοποιείται όταν θέλουμε να τυπώνουμε το αποτέλεσμα με συγκεκριμένο τρόπο
*/

// Ένα constant είναι ενα variable η τιμή του οποίου δεν μπορεί να αλλάξει

const pi float64 = 3.14159265359

/*
Τα παρακάτω variables είναι τα λεγόμενα "package level variables". Αυτά τα variables βρίσκονται εξω απο functions
με αποτέλεσμα να μπορούν να τα χρησιμοποιήσουν όλα τα functions που φτιάχνουμε. Αντίθετα όσα variables δημιουργούμε
μέσα σε μια function χρησιμοποιούνται μόνο μέσα σε αυτήν.

Σε package level (global) variables πάντα ξεκινάμε με την λέξη-κλειδί var. Η Go δεν θα μας αφήσει να χρησιμοποιήσουμε μια δήλωση
τύπου ":=". Στη συνέχεια, πρέπει να ονομάσουμε τις μεταβλητές μας και να καθορίσουμε τον τύπο τους.
Στο θέμα της ονομασίας, αυτά πρέπει να ξεκινήσουν με ένα γράμμα ή μια κάτω παύλα "_" και φυσικά, βεβαιωθείτε ότι δεν
προσπαθείτε να τα ονομάσετε με μια λέξη-κλειδί δεσμευμένη απο την Go. Άρα μην προσπαθήσετε να ονομάσετε μια var μεταβλητή
με το όνομα var

Τέλος σε αντίθεση με τα variables που δηλώνουμε εντός μιας function, εδω μπορούμε να δημιουργήσουμε variable και να μην
τα χρησιμοποιήσουμε άμεσα. Αν όμως τα variables δημιουργούνται εντός της function, η Go θα μας παραπονεθεί οτι
"αν δεν την χρησιμοποιείς πουθενά, διέγραψε την"
*/
var domainName string = "Cerebrux"

// Μπορείτε να δηλώσετε πολλαπλά variable όπως παρακάτω

var (
	dedomenaA = 2 // Δηλώνουμε μια var dedeomenaA και την αρχηκοποιούμε (initializing) με την τιμή 2
	// και άρα η GO καταλαβαίνει οτι είναι int
	dedomenaB = 3.1 // Δηλώνουμε μια var dedeomenaΒ και την αρχηκοποιούμε (initializing) με την τιμή 3.1
	//και άρα η GO καταλαβαίνει οτι είναι float
	dedomenaC string // Δηλώνουμε μια var dedeomenaC και ΔΕΝ την αρχηκοποιούμε (no initialization value).
	// Όταν δεν αρχηκοποιούμε πρέπει να δηλώσουμε τον τύπο της var.
	//Έπειτα η Go την "αρχηκοποιεί" με τιμή "empty" για χαρακτήρες (string types)
	dedomenaD int // ενώ για αριθμούς (int και float) δίνει την αρχική τιμή "0"
	dedomenaE float64
	// Δεδομένα ίδιου τύπου μπορούν γραφτούν σε μια γραμμή όπως παρακάτω
	dedomenaF, dedomenaG, dedomenaH int
)

func main() {
	// Τα Strings εντός των " " μπορούν να περιλαμβάνουν escape symbols όπως το \n για να τυπώσουν μια νέα γραμμή
	fmt.Printf("Μου αρέσουν οι \n\n") // το πρώτο \n είναι για να πάει στην επόμενη γραμμή ενώ το δεύτερο \n για να πάει στην αμέσως επόμενη
	// με αυτό τον τρόπο δημιουργούμε μια κενή γραμμή
	fmt.Println("Νέες κενές γραμμές όπως η απο πάνω μου ^^^")
	// Ας δούμε τι είναι αποθηκευμένο στις variable που δηλώσαμε στο package level:
	fmt.Println(
		"dedomenaA:", dedomenaA, "και είναι τύπου:", reflect.TypeOf(dedomenaA),
		"\ndedomenaB:", dedomenaB, "και είναι τύπου:", reflect.TypeOf(dedomenaB),
		"\ndedomenaC:", dedomenaC, "και είναι τύπου:", reflect.TypeOf(dedomenaC),
		"\ndedomenaD:", dedomenaD, "και είναι τύπου:", reflect.TypeOf(dedomenaD),
		"\ndedomenaE:", dedomenaE, "και είναι τύπου:", reflect.TypeOf(dedomenaE),
		"\ndedomenaF:", dedomenaF, "και είναι τύπου:", reflect.TypeOf(dedomenaF),
		"\ndedomenaG:", dedomenaG, "και είναι τύπου:", reflect.TypeOf(dedomenaG),
		"\ndedomenaH:", dedomenaH, "και είναι τύπου:", reflect.TypeOf(dedomenaH),
		// Όταν γράφουμε μια μεγάλη Println, είναι καλύτερα να την γράφουμε σε μορφή λίστας προς τα κάτω.
		// Σε αυτή την περίπτωση πάντα η τελευταία γραμμή λήγει με κόμμα
	)
	// Το μήκος ενός string το λαμβάνουμε με τη function len(string)
	fmt.Println("Το μήκος του string Cerebrux είναι :", len(domainName))
	// Μπορείτε να ενώσετε με +
	fmt.Println(domainName + " : Ιστοσελίδα τεχνολογίας ")
	// Τα Booleans είναι τύποι δεδομένων αλήθειας και μπορούν να πάρουν δύο τιμές, true ή false
	var isOver30 bool = true
	fmt.Println(isOver30)
	// Στην Printf η ένδειξη %f είναι για αριθμούς float
	// Ενώ με την %.xf (όπου χ αριθμός) μπορείτε να ορίσετε και την ακρίβεια του δεκαδικού ενός float
	fmt.Printf("%f \n", pi)
	fmt.Printf("%.3f \n", pi)
	// Αντί της reflect.TypeOf μπορούμε να χρησιμοποιήσουμε και την %T.
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
