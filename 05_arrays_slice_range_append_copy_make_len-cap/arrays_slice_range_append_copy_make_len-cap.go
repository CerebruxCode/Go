/*
Τι θα μάθουμε:
- Λίστα "πραγμάτων": Arrayes (πίνακες) και Slices (φέτες)
- Πως αλλάζουμε στοιχεία
- Πως δημιουργούμε νέα
- Πως αντιγράφουμε
- Θα φτιάξουμε υπολογιστής μέσου όρου
- Μεθόδους : append, copy, strings.Join, len, cap, make, range
- Να φτιάχνουμε νεα slices ενώνοντας 2 άλλα slices
- Να φτιάχνουμε πίνακες 3x3
- Θα μεγαλώνουμε και θα μικρύνουμε slices

Όταν μιλάμε για πίνακες και φέτες, ουσιαστικά μιλάμε
για αριθμημένες λίστες και αυτές οι λίστες πρέπει να
περιέχουν στοιχεία ενός μόνο τύπου (type)
Αφού λοιπόν μιλάμε για λίστες, μπορείτε να το
φανταστείτε όπως μια λίστα αγορών ή λίστα supermarket.
Κάθε στοιχείο σε αυτήν τη λίστα είναι αριθμημένο ή
όπως τείνουμε να λέμε, ευρετηριασμένο (indexed).
Και τότε κάθε στοιχείο στη λίστα πρέπει να είναι του
ίδιου τύπου, οπότε ίσως όλες συμβολοσειρές (strings)
ή όλοι ακέραιοι αριθμοί (intigers) ή όλα κινητής υποδιαστολής (float).

Ναι αλλά μεταξύ των array και slice, ποια να χρησιμοποιήσουμε ;

Γενικά μπορούμε να πούμε ότι οι slices, όπως θα πούμε και παρακάτω,
επειδή είναι πιο δυναμικές, έχουν καταλήξει χρησιμοποιούνται περισσότερο και άρα
να είναι πολύ σπάνιο να δείτε arrays που χρησιμοποιούνται σε κώδικα Go.
Οι φέτες είναι εντελώς εκεί που βρίσκονται
*/

package main

import (
	"fmt"
	"strings"
)

func main() {
	/*
		Arrays - Πίνακες
		Ένας πίνακας είναι μια αριθμημένη ακολουθία (λίστα) στοιχείων
		ίδιου τύπου με σταθερό μήκος.
		Οι πίνακες έχουν την μορφή [n]T όπου "n" είναι το σύνολο
		τιμών και Τ είναι ο τύπος(type) δεδομένων που
		θα έχουν τα στοιχεία του πίνακα (integer,float,string κλπ)

		Για να αποθηκεύσουμε έναν πίνακα σε ένα variable,
		χρησιμοποιούμε την έκφραση:

			var a [10]int

		η οποία δηλώνει μια variable με όνομα "a" ως μια σειρά
		από δέκα int (ακέραιους) αριθμούς.
		Μπορεί να είναι οποιουδήποτε τύπου αρκεί όλα τα
		στοιχεία που περιέχοντα στον πίνακα να είναι του ίδιου τύπου.
		Το σύνολο ή μήκος [n] ενός πίνακα είναι μέρος του τύπου του,
		έτσι δεν είναι δυνατή η αλλαγή μεγέθους των συστοιχιών.
		Αυτό φαίνεται περιοριστικό, αλλά μην ανησυχείτε.
		Η Go παρέχει έναν βολικό τρόπο εργασίας με πίνακες.
	*/
	// Έστω ένας πίνακας με όνομα favNums2, 5 float64 στοιχείων:
	var favNums2 [5]float64
	// Έστω στις θέσεις απο 0 έως 4 έχει τους παρακάτω αριθμούς
	// (θυμηθείτε οι αρίθμηση ξεκινάει απο το 0 και όχι απο το 1)
	favNums2[0] = 163
	favNums2[1] = 78557
	favNums2[2] = 691
	favNums2[3] = 3.141
	favNums2[4] = 1.618

	fmt.Println("\n== Πρώτο παράδειγμα ==")
	// Μπορείτε να αποκτήσετε πρόσβαση σε κάποια τιμή του
	// πίνακα παρέχοντας τον αριθμό ευρετηρίου (θέση)
	fmt.Println("favNums2 στην θέση 3 (4ος αριθμός) έχει:", favNums2[3])

	// Ένας άλλος τρόπος αρχικοποίησης ενός πίνακα είναι
	// η απευθείας δήλωση των αριθμών:

	favNums3 := [5]float64{1, 2, 3, 4, 5} // όπου οι θέσεις νοούνται αυτόματα (τιμές των θέσεων: {0,1,2,3,4})

	// Πώς να αξιολογήσετε τα στοιχεία ενός πίνακα:
	// Η range επιστρέφει key:value ζευγάρια δεδομένων
	// Αν δεν μας ενδιαφέρει να κρατήσουμε / επιστρέψουμε
	// κάποιο απο τα δύο χρησιμοποιούμε "_"
	fmt.Println("\n == favNums3 έχει ==")
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
	fmt.Println("Ενωμένα τα στοιχεία: ", a[0], a[1])
	fmt.Println("Αυτούσιος ο πίνακας: ", a)

	primes := [6]int{2, 3, 5, 7, 11, 13}                        // τιμές των θέσεων: {0,1,2,3,4,5}
	fmt.Println("Αρχική κατάσταση primes:", primes)             // το ίδιο με fmt.Println(primes[:])
	fmt.Println("Η τιμή primes στην 3η θέση είναι:", primes[3]) // για να καλέσουμε το στοιχείο στην 3η θέση μόνο.
	fmt.Println("primes[3:] ->", primes[3:])                    // δώσε την 3η τιμή μέχρι και την τελευταία
	fmt.Println("primes[:] ->", primes[:])                      // το ίδιο με fmt.Println(primes)
	fmt.Println("primes[:3] ->", primes[:3])                    // δώσε τα στοιχεία πρίν την 3η θέση
	fmt.Println("primes[0:3] ->", primes[0:3])                  // ίδιο με fmt.Println(primes[:3])
	fmt.Println("primes[3:3] ->", primes[3:3])                  // τι θα επιστρέψει ; :)
	fmt.Println("primes[1:6] ->", primes[1:6])
	fmt.Printf("Σύνολο (μήκος) των primes: %v\n", len(primes)) // για να δούμε το σύνολο/μήκος του πίνακα

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
		Υπολογιστής μέσου όρου:

		Ας φτιάξουμε ενα πρόγραμμα υπολογισμού μέσου όρου.
		Έστω οι βαθμοί 5 Μαθημάτων είναι 18, 17, 19, 16, 20.
	*/
	fmt.Println("\n== Υπολογισμός μέσου όρου μαθημάτων ==")
	// Πρώτα θα δημιουργήσουμε ένα πίνακα μήκους 5 στοιχείων
	// για να κρατήσει τις βαθμολογίες μας
	var vathmologia [5]int
	//μετά θα γεμίσουμε κάθε στοιχείο απο ένα βαθμό.
	vathmologia[0] = 18
	vathmologia[1] = 17
	vathmologia[2] = 19
	vathmologia[3] = 16
	vathmologia[4] = 20
	// Αρχικοποιούμε μια μεταβλητή που θα κρατάει την συνολική
	// βαθμολογία
	var synoloVatmologias int = 0
	// Στη συνέχεια θα δημιουργήσουμε ένα βρόχο for
	// για να υπολογιστεί η συνολική βαθμολογία
	for i := 0; i < len(vathmologia); i++ { // όσο το i (σειρά) είναι μικρότερο
		// του συνολικού μήκους του πίνακα "vathmologia"
		synoloVatmologias += vathmologia[i] // πρόσθεσε το έναν αριθμό
		// με αυτόν του επόμενου
	}
	// Τέλος θα διαιρέσουμε το συνολικό σκορ με
	// τον αριθμό των στοιχείων για να βρούμε το μέσο όρο
	fmt.Println("Ο μέσος όρος βαθμολογίας μαθημάτων είναι :", synoloVatmologias/len(vathmologia))

	/*
		Τον υπολογιστή μέσου όρου μπορούμε να τον γράψουμε και με την range.
		Συγκρίνετε την παρακάτω μέθοδο με την επάνω, όπου αντί για μεταβλητή σειράς (i)
		χρησιμοποιούμε _
	*/
	var synoloVatmologias2 int = 0
	for _, mikos := range vathmologia { // Εδω αν χρησιμοποιήσουμε i
		// αντί της κενής/προσωρινής var η Go θα παραπονεθεί
		// για μη χρησιμοποιούμενη variable.
		// Οπότε με μια μόνο _ (κάτω παύλα) χρησιμοποιούμε για να πούμε στην Go ότι δεν θα χρειαστούμε
		// την μεταβλητή αυτή.
		synoloVatmologias2 += mikos
	}
	fmt.Println("Ο μέσος όρος βαθμολογίας μαθημάτων με range είναι :",
		synoloVatmologias2/len(vathmologia))
	/*
		Κάποιες φορές όμως μπορεί οι αριθμοί να είναι πολύ και
		να μην χωράνε σε μια γραμμή οπότε μπορούμε να τα γράψουμε ως :

			faveNums3a := [10]float64{
				15,
				16,
				25,
				34,
				85,
				34,
				49,
				32,
				42,
				// προσέξτε το επιπλέον κόμμα , μετά το 42.
				// Αυτό απαιτείται από τη Go και μας επιτρέπει
				// να αφαιρέσουμε εύκολα ένα στοιχείο από ένα πίνακα
				// μετατρέποντάς το σε σχόλιο:
					// 12,
			}
			Για να αφαιρέσετε όμως το τελευταίο αυτό στοιχείο, θα πρέπει
			να αλλάξετε και το μέγεθος του πίνακα. Σε διαφορετική περίπτωση,
			όπως πάνω, θα γεμίσει τις κενές θέσεις με την default τιμή του
			τύπου (στο int είναι 0)
			Η λύση της Go σε αυτό το πρόβλημα είναι να
			χρησιμοποιήσουμε έναν διαφορετικό τύπο: Τα slices.

	*/
	/*
		Τρίτο παράδειγμα
	*/
	fmt.Println("\n== Τέταρτο παράδειγμα ==")
	// Σε περίπτωση που θέλουμε να φτιάξουμε array
	// που δεν ξέρουμε το μέγεθός του εξαρχής,
	// τότε μπορούμε να βάλουμε τρεις τελίτσες
	// στην θέση του καθορισμού μεγέθους
	glosses := [...]string{"Go", "Python", "C"}
	fmt.Println("Οι αγαπημένες γλώσσες προγραμματισμού είναι: ", glosses)
	fmt.Printf("%T \n", glosses)

	/*
		Slices

		H μορφή []T{} είναι ένα Slice με στοιχεία τύπου Τ

		Ένα Array είπαμε ότι έχει σταθερό μέγεθος.
		Ένα Slice από την άλλη, είναι μια δυναμικού μεγέθους μορφή
		απεικόνισης των στοιχείων μιας array. Ένα slice με άλλα λόγια
		είναι ένα κομμάτι ενός πίνακα. Όπως στους πίνακες έτσι και στα
		slices έχουμε δείκτες θέσης και ένα μήκος.
		Βεβαία, σε αντίθεση με τους πίνακες, αυτό το μήκος
		επιτρέπεται να αλλάζει.

		Στην πράξη, τα Slices είναι πολύ πιο συχνά στην χρήση
		από ότι τα Arrays.

	*/
	fmt.Println("\n== SLICES ==")

	// η σύνταξή τους μοιάζει με arrays αλλά χωρίς να ορίζουμε το μέγεθος
	numSlice := []int{5, 4, 3, 2, 1} // Η μόνη διαφορά μεταξύ αυτού και
	// ενός πίνακα είναι ότι λείπει το μήκος μεταξύ των αγκύλων.
	// Στην περίπτωση μας, το numSlice έχει δημιουργηθεί με ένα μήκος 0.

	/*
		Μπορείτε να δημιουργήσετε μια νέα slice (φέτα) από ένα άλλο slice
		χρησιμοποιώντας της θέσεις που θέλετε να αντιγράψει

		Παρακάτω βλέπουμε ένα παράδειγμα όπου η Slice σχηματίζεται
		με τον καθορισμό δύο δεικτών, ενός χαμηλού και ενός υψηλού,
		που χωρίζονται από μια άνω-κάτω τελεία:

				a[low : high]

		Αυτό επιλέγει μια περιοχή της Slice που περιλαμβάνει το πρώτο στοιχείο,
		αλλά αποκλείει το τελευταίο. Άρα το low είναι ο δείκτης για το
		πού να ξεκινήσει το slice και το high είναι ο δείκτης που να τελειώσει
		(αλλά δεν συμπεριλαμβάνεται ο ίδιος ο δείκτης).

		Η ακόλουθη έκφραση δημιουργεί μια Slice που περιλαμβάνει
		τα στοιχεία 1 έως 3 μιας a:
		 a[1:4]
	*/
	numSlice2 := numSlice[3:5] // δηλαδή numSlice2 με τιμές [2,1]

	fmt.Println("numSlice2[0] =", numSlice2[0])
	/*
		Αν δεν δηλώσετε το πρώτο τμήμα του ευρετηρίου (θέση),
		τότε το προεπιλεγμένο είναι 0
		Αν δεν δηλώσετε το τελευταίο τμήμα του ευρετηρίου (θέση),
		τότε θα λάβει το τελευταίο
		Επομένως μπορείτε να παραλείψετε το low, high ή ακόμη
		και τα δύο.

		pinakas[0] είναι το ίδιο με pinakas[0:len(pinakas)],
		pinakas[:8] είναι το ίδιο με pinakas[0:8] και
		pinakas[:] είναι το ίδιο με pinakas[0:len(pinakas)]

	*/

	fmt.Println("numSlice[:2] =", numSlice[:2])
	fmt.Println("numSlice[2:] =", numSlice[2:])
	/*
		Μπορείτε επίσης να δημιουργήσετε ένα κενό slice
		και να καθορίσετε τον τύπο δεδομένων,
		Ως μήκος (λαμβάνει 0), χωρητικότητα (μέγιστο μέγεθος)
		Το 10 εδω αντιπροσωπεύει την χωρητικότητα του
		υπο-πίνακα στον οποίο εμπεριέχεται το numSlice3

	*/

	numSlice3 := make([]int, 5, 10)

	/*
	   Η Go παρέχει δύο ενσωματωμένες συναρτήσεις για να δουλέψουμε με τα slices
	    - copy
	    - append

	   Με την μέθοδο "copy()" μπορείτε να αντιγράψετε ένα τεμάχιο σε ένα άλλο
	   Με την "append()" δημιουργείται ένα νέο slice λαμβάνοντας ένα υπάρχον slice
	   ως το πρώτο όρισμα και προσαρτήσει όλα τα επακόλουθα
	   ορίσματα σε αυτό (όπως τα βαγόνια ενός τραίνου).

	*/

	// Το παρακάτω αντιγράφει το numSlice στο numSlice3.
	// Φανταστείτε το ως "copy(αντιγραμμένο, απο)"
	copy(numSlice3, numSlice)

	fmt.Println("numSlice3[0] = ", numSlice3[0])
	// Αν όμως αντί για 5 που δώσαμε ως χώρο στο
	// numSlice3 := make([]int, 5, 10)
	// δίναμε π.χ. 2, τότε κάποια στοιχεία δεν θα χωρούσαν να αντιγραφούν

	/*
		Αρκετές φορές θα χρειαστεί να προσθέσετε νέα στοιχεία σε ένα slice
		οπότε θα χρησιμοποιήσουμε την μέθοδο "append".
		Η τεκμηρίωση της είναι εδώ https://golang.org/pkg/builtin/#append.

				func append(s []T, vs ...T) []T

		Η πρώτη παράμετρος "s" της append αναφέρεται σε ένα slice "s" τύπου Τ
		και τα υπόλοιπα είναι τιμές τύπου Τ που θα προσαρτηθούν (append) στο slice.

		Η προκύπτουσα τιμή της append είναι ένα slice που περιέχει
		όλα τα στοιχεία της αρχικής slice συν τις παρεχόμενες τιμές.

		Εάν το array "s" είναι πολύ μικρό για να χωρέσει όλες τις τιμές
		που γίνονται Append, τότε θα διατεθεί μεγαλύτερο array.

		Η επιστρεφόμενη slice θα είναι αναφορά (θα δείχνει)
		στον νέο array που θα δημιουργηθεί.

	*/
	fmt.Println("\n== append SLICES ==")
	// Θυμίζουμε οτι numSlice3 := make([]int, 5, 10)
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

	names := []string{
		"Γιάννης",
		1: "Παύλος",
		2: "Μαρία",
		3: "Άννα",
	}
	// Μπορείτε να το γράψετε και σε μια γραμμή
	// καθορίζοντας και τις θέσεις των τιμών
	// ξεκινώντας πάντα απο το 0
	nameOneline := []string{0: "Γιάννης", 1: "Παύλος", 2: "Μαρία", 3: "Άννα"}
	_ = nameOneline // muted για να μην παραπονιέται η Go οτι δεν το χρησιμοποιούμε
	fmt.Println("\n== Ονόματα ==")
	fmt.Println("Ονόματα =", names)
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

		Το μήκος ενός slice είναι ο συνολικός αριθμός
		των στοιχείων που περιέχει.
		Η χωρητικότητα ενός slice είναι ο αριθμός των
		στοιχείων της μετρώντας από το πρώτο [0] στοιχείο της.

		Το μήκος και η χωρητικότητα ενός slice "s" μπορεί
		να ληφθεί χρησιμοποιώντας τις μεθόδους len(s) και cap(s).

		Μπορείτε να επεκτείνετε το μήκος του slice,
		τεμαχίζοντας το ξανά αρκεί να έχει επαρκή χωρητικότητα.

		Δοκιμάστε να αλλάξετε μια από τις functions slicing στο πρόγραμμα
		παραδειγμάτων για να το επεκτείνετε πέρα από
		την χωρητικότητά του και να δείτε τι συμβαίνει.
	*/
	fmt.Println("\n== Δυναμικά Slices ==")
	dinamikoSlice := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("Αρχική μορφή")
	printSlice(dinamikoSlice) // την func printSlice την φτιάξαμε παρακάτω
	// για να μας εμφανίζει το μήκος 'len'
	// και την χωρητικότητα 'cap'

	// Τεμαχίστε ένα slice για να του δώσετε μηδενικό μήκος.
	fmt.Println("Μετά απο μεταβολή σε 0 μήκος")
	dinamikoSlice = dinamikoSlice[:0]
	printSlice(dinamikoSlice)

	// 	Μεγαλώστε το μήκος του.
	fmt.Println("Μετά απο επιμήκυνση του")
	dinamikoSlice = dinamikoSlice[:4]
	printSlice(dinamikoSlice)

	// Αφαιρέστε τις πρώτες δυο τιμές του
	fmt.Println("Μετά απο αφαίρεση των 2 πρώτων τιμών")
	dinamikoSlice = dinamikoSlice[2:]
	printSlice(dinamikoSlice)

	/*
		Η μηδενική τιμή μιας slice είναι η "nil".
		Μια μηδενική slice έχει μήκος και χωρητικότητα "0" και
		δεν περιέχει υποκείμενη array.
	*/
	var s0 []int
	// τύπωσε τον πίνακα, το μήκος και την χωρητικότητα του
	fmt.Println(dinamikoSlice, len(s0), cap(s0))
	if s0 == nil {
		fmt.Println("Μηδενική τιμή μιας Slice")
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
		Αρα η make φτιάχνει a := make([]type, len, cap)

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

	// Οι παίχτες παίρνουν την σειρά τους.
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

func printSlice(s []int) { // φτιάξαμε μια δική μας προβολή
	// για να μας εμφανίζει το μήκος 'len' και την χωρητικότητα
	// 'cap' που χρησιμοποιούμε πιο πάνω
	fmt.Printf("Με len=%d και cap=%d το αποτέλεσμα είναι :%v\n", len(s), cap(s), s)
}

func printSliceB(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n",
		s,
		len(x),
		cap(x),
		x)
}

func myRange() {

	/*
		Η "range" σε έναν βρόγχο for, είχαμε πει οτι σαρώνει
		μια "slice" ή "map" (τα map θα τα δούμε σε άλλο μάθημα).
		Όταν τρέχουμε την "range" σε ένα slice,
		επιστρέφονται δύο τιμές για κάθε επανάληψη.
		- ο δείκτης (index - θέση)
		- ένα αντίγραφο του στοιχείου σε αυτόν τον δείκτη.
	*/
	fmt.Println("\n== Range Slice ==")
	// Ας φτιάξουμε την slice pow με ένα σύνολο αριθμών που είναι
	// το αποτέλεσμα της δύναμης του 2 εις την 0, 2 εις την  1,
	// 2 εις την 2, 2 εις την 3, 2 εις την 4, 2 εις την 5, κλπ.
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	// θα σαρώσουμε όλες τις τιμές της pow, και θα αντιστοιχίσουμε την θέση
	// στο i και την τιμή στο v
	for i, v := range pow {
		fmt.Printf("το αποτέλεσμα του 2**%d = %d\n", i, v)
	}

	// Άλλο παράδειγμα δημιουργόντας ένα slice με την make:
	fmt.Println("\n== powMake ==")

	powMake := make([]int, 10)
	fmt.Println("\npowMake - Αρχική κατάσταση :", powMake)

	for i := range powMake {
		powMake[i] = 1 << uint(i) // == 2**i
	}
	fmt.Println("powMake - Μετά το γέμισμα: ", powMake)
	/*
		Μπορείτε να παραλείψετε τον δείκτη ή την τιμή
		αναθέτοντας τα σε "_" (κάτω παύλα).

		for i, _ := range pow
		for _, value := range pow

	*/
	fmt.Println("powMake - λίστα :")
	for _, value := range powMake {
		fmt.Printf("%d\n", value)
	}
}

/*

	Για να μάθετε περισσότερα για τα Slices διαβάστε:
	https://blog.golang.org/go-slices-usage-and-internals

*/
