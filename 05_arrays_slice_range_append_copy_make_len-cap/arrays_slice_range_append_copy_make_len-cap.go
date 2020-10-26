package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		Arrays - Πίνακες

		Οι πίνακες έχουν την μορφή [n]T όπου "n" είναι το σύνολο τιμών και Τ είναι ο τύπος(type) δεδομένων που
		θα έχουν τα στοιχεία του πίνακα (integer,float,string κλπ)

		Για να αποθηκεύσουμε έναν πίνακα σε ένα variable, χρησιμοποιούμε την έκφραση:

			var a [10]int

		η οποία δηλώνει μια variable με όνομα "a" ως μια σειρά από δέκα ακέραιους (int) αριθμούς.
		Μπορεί να είναι οποιουδήποτε τύπου αρκεί όλα τα στοιχεία που περιέχονται
		στον πίνακα να είναι του του ίδιου τύπου.
		Το σύνολο ή μήκος [n] ενός πίνακα είναι μέρος του τύπου του,
		έτσι δεν είναι δυνατή η αλλαγή μεγέθους των συστοιχιών.
		Αυτό φαίνεται περιοριστικό, αλλά μην ανησυχείτε.
		Η Go παρέχει έναν βολικό τρόπο εργασίας με συστοιχίες.
	*/
	// Έστω ένας πίνακας με όνομα favNums2  5 στοιχείων:
	var favNums2 [5]float64
	// Έστω στις θέσεις απο 0 έως 4 έχει τους παρακάτω αριθμούς (θυμηθείτε οι αρίθμηση ξεκινάει απο το 0 και όχι απο το 1)
	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618

	fmt.Println("\n== Πρώτο παράδειγμα ==")
	// Μπορείτε να αποκτήσετε πρόσβαση σε κάποια τιμή του πίνακα παρέχοντας
	// τον αριθμό ευρετηρίου (θέση)
	fmt.Println("favNums2 στην θέση 3 (4ος αριθμός) έχει:", favNums2[3])

	// Ένας άλλος τρόπος αρχικοποίησης ενός πίνακα είναι η απευθείας δήλωση των αριθμών

	favNums3 := [5]float64{1, 2, 3, 4, 5} // τιμές των θέσεων: {0,1,2,3,4}

	// Πώς να αξιολογήσετε τα στοιχεία ενός πίνακα:
	// Η range επιστρέφει key:value ζευγάρια δεδομένων
	// Αν δεν μας ενδιαφέρει να κρατήσουμε / επιστρέψουμε
	// κάποιο απο τα δύο χρησιμοποιούμε "_"
	fmt.Println("\n favNums3 έχει:")
	for i, value := range favNums3 {

		fmt.Println(value, "στη θέση", i)

	}

	/*
		Δεύτερο παράδειγμα:
	*/
	fmt.Println("\n== Δεύτερο παράδειγμα ==")
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{2, 3, 5, 7, 11, 13}                       // τιμές των θέσεων: {0,1,2,3,4,5}
	fmt.Println(primes)                                        // το ίδιο με fmt.Println(primes[:])
	fmt.Println("Η τιμή prime στην 3η θέση είναι:", primes[3]) // για να καλέσουμε το στοιχείο στην 3η θέση μόνο.
	fmt.Println(primes[3:])                                    // δώσε την 3η τιμή μέχρι και την τελευταία
	fmt.Println(primes[:])                                     // το ίδιο με fmt.Println(primes)
	fmt.Println(primes[:3])                                    // δώσε τα στοιχεία πρίν την 3η θέση
	fmt.Println(primes[0:3])                                   // ίδιο με fmt.Println(primes[:3])
	fmt.Println(primes[3:3])                                   // τι θα επιστρέψει ; :)
	fmt.Println(primes[1:6])
	fmt.Printf("Σύνολο των primes: %v\n", len(primes)) // για να δούμε το σύνολο/μήκος του πίνακα

	/*
		Τρίτο παράδειγμα

	*/
	fmt.Println("\n== Τρίτο παράδειγμα ==")
	// Γραμμική άλγεβρα : Πίνακας 3x3
	var identityMatrixA [3][3]int = [3][3]int{[3]int{1, 0, 0}, [3]int{0, 1, 0}, [3]int{0, 0, 1}}
	fmt.Println("identityMatrixA: ", identityMatrixA)
	// Το παραπάνω μπορεί να γραφτεί και ως εξής:
	var identityMatrixB [3][3]int
	identityMatrixB[0] = [3]int{1, 0, 0}
	identityMatrixB[1] = [3]int{0, 1, 0}
	identityMatrixB[2] = [3]int{0, 0, 1}
	fmt.Println("identityMatrixB: ", identityMatrixB)

	/*
		Slices


		Ένα Array είπαμε ότι έχει σταθερό μέγεθος.
		Ένα Slice από την άλλη, είναι μια δυναμικού μεγέθους μορφή
		απεικόνισης των στοιχείων μιας array.
		Στην πράξη, τα Slices είναι πολύ πιο συχνά από τα Arrays.

		H μορφή []T{} είναι ένα Slice με στοιχεία τύπου Τ
	*/
	fmt.Println("\n== SLICES ==")
	// η σύνταξή τους μοιάζει με arrays αλλά χωρίς να ορίζουμε το μέγεθος

	numSlice := []int{5, 4, 3, 2, 1}

	// Μπορείτε να δημιουργήσετε μια νέα slice (φέτα) από ένα άλλο slice
	// χρησιμοποιώντας της θέσεις που θέλετε να αντιγράψει

	// Παρακάτω βλέπουμε ένα παράδειγμα όπου η Slice σχηματίζεται
	// με τον καθορισμό δύο δεικτών, ενός χαμηλού και ενός υψηλού,
	// που χωρίζονται από μια άνω-κάτω τελεία:

	//		a[low : high]

	// Αυτό επιλέγει μια περιοχή της Slice που περιλαμβάνει το πρώτο στοιχείο,
	// αλλά αποκλείει το τελευταίο.
	// Η ακόλουθη έκφραση δημιουργεί μια Slice που περιλαμβάνει τα στοιχεία 1 έως 3 μιας a:
	// a[1:4]
	numSlice2 := numSlice[3:5] // δηλαδή numSlice3 με τιμές [2,1]

	fmt.Println("numSlice2[0] =", numSlice2[0])

	// Αν δεν δηλώσετε το πρώτο τμήμα του ευρετηρίου (θέση), τότε το προεπιλεγμένο είναι 0
	// Αν δεν δηλώσετε το τελευταίο τμήμα του ευρετηρίου (θέση), τότε θα λάβει το τελευταίο

	fmt.Println("numSlice[:2] =", numSlice[:2])
	fmt.Println("numSlice[2:] =", numSlice[2:])

	// Μπορείτε επίσης να δημιουργήσετε ένα κενό slice και να καθορίσετε τον τύπο δεδομένων,
	// μήκος (λαμβάνει 0), χωρητικότητα (μέγιστο μέγεθος)

	numSlice3 := make([]int, 5, 10)

	// Μπορείτε να αντιγράψετε με την μέθοδο "copy()" ένα τεμάχιο σε άλλο
	// Το παρακάτω αντιγράφει το numSlice στο numSlice3.
	// Φανταστείτε το ως "copy(αντιγραμμένο, απο)"

	copy(numSlice3, numSlice)

	fmt.Println("numSlice3[0] = ", numSlice3[0])

	/*
		Αρκετές φορές θα χρειαστεί να προσθέσετε νέα στοιχεία σε ένα slice
		και έτσι η Go παρέχει μια ενσωματωμένη μέθοδο "append".
		Η τεκμηρίωση της είναι εδώ https://golang.org/pkg/builtin/#append.

				func append(s []T, vs ...T) []T

		Η πρώτη παράμετρος "s" της append αναφέρεται σε ένα slice "s" τύπου Τ
		και τα υπόλοιπα είναι τιμές τύπου Τ που θα προσαρτηθούν (append) στο slice.

		Η προκύπτουσα τιμή της append είναι ένα slice που περιέχει
		όλα τα στοιχεία της αρχικής slice συν τις παρεχόμενες τιμές.

		Εάν το array "s" είναι πολύ μικρό για να χωρέσει όλες τις τιμές
		που γίνονται Append, τότε θα διατεθεί μεγαλύτερο array.
		Η επιστρεφόμενη slice θα είναι αναφορά (θα δείχνει) στον νέο array που θα δημιουργηθεί.

	*/
	fmt.Println("\n== append SLICES ==")
	numSlice3 = append(numSlice3, 0, -1)

	fmt.Println("numSlice3[6] = ", numSlice3[6])

	/*

		Τα Slices είναι σαν αναφορές σε Arrays
		Ένα Slice δεν αποθηκεύει δεδομένα, απλά περιγράφει
		ένα τμήμα μιας array.

		Αν αλλάξουμε τα στοιχεία μιας slice, τότε τροποποιούμε
		τα αντίστοιχα στοιχεία της array στην οποία ανήκουν.
		Άλλες slices που μοιράζονται τον ίδιο array θα δουν επίσης αυτές τις αλλαγές.
	*/

	names := [4]string{
		"Γιάννης",
		"Παύλος",
		"Μαρία",
		"Άννα",
	} // Μπορείτε να το γράψετε και σε μια γραμμή "names := [4]string{"Γιάννης", "Παύλος", "Μαρία", "Άννα"}"
	fmt.Println("names = ", names)
	// Ας δημιουργήσουμε 2 νέα slices:
	alfa := names[0:2]
	beta := names[1:3]
	fmt.Println("alfa=", alfa, "και", "beta=", beta)
	// Ας αλλάξουμε την τιμή της θέσης 0
	beta[0] = "XXX"
	fmt.Println("Μετά την αλλαγή έχουμε το alfa=", alfa, "και", "beta=", beta)
	fmt.Println("Οπότε τα names = ", names)

	/*
		Συμπέρασμα:

		Ένα Slice (τεμάχιο) έχει μήκος αλλά και χωρητικότητα.

		Το μήκος ενός τεμαχίου είναι ο συνολικός αριθμός των στοιχείων που περιέχει.
		Η χωρητικότητα ενός τεμαχίου είναι ο αριθμός των στοιχείων της μετρώντας
		από το πρώτο στοιχείο της.

		Το μήκος και η χωρητικότητα ενός τεμαχίου "s" μπορεί να ληφθεί χρησιμοποιώντας
		τις εκφράσεις len(s) και cap(s).

		Μπορείτε να επεκτείνετε το μήκος του τεμαχίου, τεμαχίζοντας το ξανά αρκεί
		να έχει επαρκή χωρητικότητα.

		Δοκιμάστε να αλλάξετε μια από τις functions τεμαχισμού στο πρόγραμμα
		παραδειγμάτων για να το επεκτείνετε πέρα από την χωρητικότητά του και να δείτε τι συμβαίνει.
	*/
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Τεμαχίστε ένα slice για να του δώσετε μηδενικό μήκος.
	s = s[:0]
	printSlice(s)

	// 	Μεγαλώστε το μήκος του.
	s = s[:4]
	printSlice(s)

	// Αφαιρέστε τις πρώτες δυο τιμές του
	s = s[2:]
	printSlice(s)

	/*
		Η μηδενική τιμή μιας slice είναι η "nil".
		Μια μηδενική slice έχει μήκος και χωρητικότητα "0" και
		δεν περιέχει υποκείμενη array.
	*/
	var s0 []int
	// τύπωσε τον πίνακα, το μήκος και την χωρητικότητα του
	fmt.Println(s, len(s0), cap(s0))
	if s0 == nil {
		fmt.Println("nil! ρε !")
	}

	/*
		Slices μπορούν να δημιουργηθούν με την function "make".
		Με αυτόν τον τρόπο δημιουργείτε arrays δυναμικού μεγέθους.

		Η function make φτιάχνει έναν μηδενικό πίνακα (array) και επιστρέφει
		ένα slice που αναφέρεται σε αυτήν:
		a := make([]int, 5)  // που σημαίνει len(a)=5

		Για να καθορίσετε χωρητικότητα, δηλώστε μία τρίτη παράμετρο όπως παρακάτω:
		b := make([]int, 0, 5) // που σημαίνει len(b)=0 και cap(b)=5

		Επίσης μπορούμε να κάνουμε τα παρακάτω
		b = b[:cap(b)] // len(b)=5, cap(b)=5
		b = b[1:]      // len(b)=4, cap(b)=4
	*/
	fmt.Println("\n== Make SLICES ==")
	ant := make([]int, 5)
	printSliceB("ant =", ant)

	bant := make([]int, 0, 5)
	printSliceB("bant =", bant)

	cant := bant[:2]
	printSliceB("cant =", cant)

	dant := cant[2:5]
	printSliceB("dant= ", dant)

	// Τα Slices μπορούν να περιέχουν οποιοδήποτε τύπο,
	// συμπεριλαμβανομένων άλλων slices.
	// Φτιάξε ένα board τρίλιζας (tic-tac-toe).

	fmt.Println("\n== tic-tac-toe ==")
	board := [][]string{
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
		[]string{"_", "_", "_"},
	}

	// Οι παίχτες πέρνουν την σειρά τους.
	board[0][0] = "X"
	board[2][2] = "O"
	board[1][2] = "X"
	board[1][0] = "O"
	board[0][2] = "X"

	for kouti := 0; kouti < len(board); kouti++ {
		fmt.Printf("%s\n", strings.Join(board[kouti], "|")) // εδώ βλέπουμε την μέθοδο .Join
		// που ενώνει στοιχεία με ένα διαχωριστικό που θέτουμε εμείς (μπορεί να είναι ένα κενό)
	}

	/*
		Range
	*/
	myRange()

}

func printSlice(s []int) { // φτιάξαμε μια δική μας προβολή που χρησιμοποιούμε πιο πάνω
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

func printSliceB(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s, len(x), cap(x), x)
}

func myRange() {

	/*
		Η "range" ενός βρόχου for, σαρώνει μια "slice" ή "map" (τα map θα τα δούμε σε άλλο μάθημα).
		Όταν τρέχουμε την "range" σε ένα slice, επιστρέφονται δύο τιμές
		για κάθε επανάληψη.
		Ο πρώτος είναι ο δείκτης (index) και ο δεύτερος είναι ένα αντίγραφο του
		στοιχείου σε αυτόν τον δείκτη.
	*/
	fmt.Println("\n== Range Slice ==")
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}

	// Άλλο παράδειγμα:
	powMake := make([]int, 10)

	//	Εάν θέλετε μόνο το index,
	//	μπορείτε να παραλείψετε τη δεύτερη variable.

	//	for i := range pow

	for i := range powMake {
		powMake[i] = 1 << uint(i) // == 2**i
	}
	/*
		Μπορείτε να παραλείψετε τον δείκτη ή την τιμή
		αναθέτοντας τα σε "_" (κάτω παύλα).

		for i, _ := range pow
		for _, value := range pow

	*/
	for _, value := range powMake {
		fmt.Printf("%d\n", value)
	}
}

/*

	Για να μάθετε περισσότερα για τα Slices διαβάστε:
	https://blog.golang.org/go-slices-usage-and-internals

*/
