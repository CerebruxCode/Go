package main

import (
	"fmt"
	"strings"
)

/*
	Όπως είδαμε οι συναρτήσεις (func, functions) μας επιτρέπουν
	την επαναχρησιμοποίηση κώδικα και συμβάλουν στην δημιουργία
	μιας λειτουργικής δομής στον κώδικά μας.
	Μπορούμε να χρησιμοποιήσουμε τις συναρτήσεις ως παράμετροι,
	να τις αναθέσουμε σε variables και άλλες δομές δεδομένων κλπ.

	Σε γενικές γραμμές είναι της μορφής:

			func funcName (parametersPassed) returnType {
				....
			}

	οι variables που χρησιμοποιούνται εντός της συνάρτησης
	δε δύναται να χρησιμοποιηθούν εκτός αυτής

	Παρακάτω ξεκινάμε δημιουργώντας μερικές δικές μας
*/

func addThemUp(numbers []float64) float64 {
	/*
		H addThemUp λαμβάνει μια float64 Array και επιστρέφει float64
		Η επιστροφή θα μπορούσε να έχει και όνομα π.χ. rtrnNumbers float64
		αλλά επειδή δεν την χρειαζόμαστε δεν την ονομάζουμε.
		Οι παράμετροι - ορίσματα που λαμβάνει ως είσοδο και φυσικά επιστρέφει
		ως έξοδο μια function ονομάζονται function signature
	*/
	sum := 0.0 // Δηλώνουμε μια variable sum με τιμή εκκίνησης 0.0

	for _, val := range numbers {
		/*
			η range μας δίνει key:value
			αλλά εμάς δεν μας ενδιαφέρει η key οπότε δεν την αποθηκέυουμε
			σε κάποια variable αλλά την αγνοούμε με την "_"

		*/

		sum += val // Συντομογραφία για sum = sum + val

	}

	return sum

}

// Οι Go functions μπορούν να επιστρέψουν πολλές τιμές
// Εδώ οι επιστροφές είναι ανώνυμες
func next2Values(number int) (int, int) {

	return number + 1, number + 2

}

func converter(epaggelma, onomatEponimo string) (s1, s2 string) {
	/*
		Λέμε στην function να αποδεχτεί δύο strings ως είσοδο και να επιστρέψει
		δύο strings σε αυτό που την καλεί. Στη συνέχεια, μέσα στην function,
		αλλάζουμε την περίπτωση των δύο μεταβλητών που χρησιμοποιηθηκαν ως είσοδο,
		την μία σε titleCase (Κεφαλάιο αρχικό γράμμα) και την άλλη σε κεφαλαία
		τα οποία και επιστρέφουμε με τις τροποποιημένες τους συμβολοσειρές.

	*/
	s1 = strings.Title(epaggelma)
	s2 = strings.ToUpper(onomatEponimo)
	/*
	   Όταν θα καλούμε τη "converter" συνάρτηση μας, εδώ διαβιβάζουμε αντίγραφα της μεταβλητής
	   "epaggelma" και την variable "onomatEponimo", ενώ στην function signature,
	   αν θέλουμε, μπορούμε να τις αντιστοιχίσουμε με διαφορετικά ονόματα,
	   ας πούμε s1 και s2. Το μόνο που πρέπει να γνωρίζουμε εδώ
	   είναι η σειρά με την οποία πέρασαν.
	   Έτσι, στο s1 πρόκειται να  εκχωρηθεί η τιμή από το αλλαγμένο "epaggelma" και στην
	   s2 η τιμή από το αλλαγμένο "onomatEponimo".
	*/
	return s1, s2

	//return epaggelma, onomatEponimo
}

// Μπορείτε να λάβετε έναν απροσδιόριστο αριθμό τιμών με τις 3 τελείες "args ... int"
// Αυτά ονομάζονται variadic functions
func subtractThem(args ...int) int {

	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue

}

func mikroterosArithmos(kapoioiArithmoi ...int) int {
	mikroteros := kapoioiArithmoi[0]

	for _, i := range kapoioiArithmoi {
		if i < mikroteros {
			mikroteros = i
		}
	}

	return mikroteros
}

// Παράδειγμα αναδρομής (recursion): Όταν η function καλεί τον εαυτό της
// Εδώ βλέπουμε υπολογισμό παραγοντικού
// Παραγοντικό του 4: factorial(4)
// 4 * factorial(3) == 4 * 3 *2 * 1
// 3 * factorial(2) == 3 * 2 * 1
// 2 * factorial(1) == 2 * 1
// factorial(0) == 1

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num-1)
}

/*
	Μια defer αναβάλλει την εκτέλεση μιας λειτουργίας
	μέχρι να γίνει η επιστροφή της περιβάλλουσας function.

	Οι παράμετροι της αναβαλλόμενης κλήσης αξιολογούνται αμέσως,
	αλλά η κλήση της δεν εκτελείται μέχρι να γίνει η επιστροφή
	της περιβάλλουσας function.

	Παράδειγμα :
	func main() {
		defer fmt.Println("γεία σου")
		fmt.Println("κόσμε")
	}
	θα επιστρέψει "κόσμε γειά " αντί "γειά σου κόσμε"
	Τα παρακάτω printOne και printTwo θα χρησιμοποιηθούν στην main() ως defer
*/

func printOne() {
	fmt.Println(1)
}

func printTwo() {
	fmt.Println(2)
}

// Εάν εμφανιστεί ένα σφάλμα, μπορούμε να πιάσουμε το σφάλμα με την recover ώστε να επιτρέψουμε
// στον κώδικα μας να συνεχίσει να εκτελείται

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover())
	}()
	solution := num1 / num2
	return solution
}

// Εδώ δείχνουμε μια function που θα καλέσει την "panic" και
// θα τη διαχειριστεί με τη "recover"

func demPanic() {

	defer func() {

		// Εάν δεν κάναμε εκτύπωση του μηνύματος, τίποτα δε θα έδειχνε

		fmt.Println(recover())

	}()
	panic("PANIC")

}

func main() { // Ξεκινάμε το πρόγραμμά μας χρησιμοποιώντας εσωτερικά της main όλες τις function
	// που φτιάξαμε. Η main επίσης δεν παίρνει ορίσματα - παραμέτρους,
	// δεν μπορεί και δεν επιστρέφει καμία τιμή.
	// στην πραγματικότητα όταν η main τερματίσει, ολόκληρο το πρόγραμμα τερματίζεται και
	// αν υποθέσουμε τερματίσει χωρίς σφάλματα, τότε θα επιστρέψει τον κωδικό εξόδου 0 στο λειτουργικό σύστημα.

	listOfNums := []float64{1, 2, 3, 4, 5}

	fmt.Println("Άθροισμα των τιμών :", addThemUp(listOfNums))

	num1, num2 := next2Values(5) // Δίνουμε σε 2 variables, τιμές από μια function
	fmt.Println(num1, num2)

	epaggelma := "δάσκαλος"
	onomatEponimo := "αγνωστος αγνώστου"
	fmt.Println(converter(epaggelma, onomatEponimo))
	// Στείλτε έναν αόριστο αριθμό τιμών σε μια function
	// https://en.wikipedia.org/wiki/Variadic_function (Variadic Function)

	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// Μπορείτε να δημιουργήσετε μια func μέσα σε μια func. Έχει πρόσβαση στις
	// τοπικές variables της func στην οποία περιλαμβάνεται

	// Μια func όπως αυτή, χωρίς τοπικές variables είναι ένα "closure"

	num3 := 3

	doubleNum := func() int {

		num3 *= 2
		return num3

	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

	// Ψάχνουμε με το variadic function μας τον μικρότερο αριθμό
	kapoioiArithmoi := mikroterosArithmos(13, 5, 10, 13, 17, 14, 16)
	fmt.Println("Ο μικρότερος απο τους αριθμούς είναι το", kapoioiArithmoi)

	// Εδώ ένα έχουμε μια αναδρομική συνάρτηση (recursive function)

	fmt.Println(factorial(4))

	// Η Defer όπως είπαμε εκτελεί μια λειτουργία μετά την ολοκλήρωση της func στην οποία καλείται
	// Η Defer μπορεί να χρησιμοποιηθεί για να δημιουργήσουμε μια func με ένα λογικό τρόπο
	// αλλά ταυτόχρονα για να εκτελέσουμε μια τελευταία λειτουργία ως λειτουργία καθαρισμού
	// π.χ. Θα κάνουμε defer το κλείσιμο ενός αρχείου αφού πρώτα το ανοίξουμε και εκτελέσουμε εργασίες

	defer printTwo()
	printOne()

	// Καλούμε μια δικιά μας defer
	countDefer()

	// Εδω χρησιμοποιούμε τη recover() για να πιάσουμε το error που θα προκαλούσε η διαίρεση με 0
	// και να αποτρέψουμε το "κρασάριμα/σταμάτημα" της εφαρμογής μας

	fmt.Println(safeDiv(3, 0))
	fmt.Println(safeDiv(3, 2))

	// Μπορούμε να αποτρέψουμε τα δικά μας error και να επανέλθουμε με panic & recover

	demPanic()

}

/*
	Οι αναβαλλόμενες κλήσεις εντός της συνάρτησης μας στοιχίζονται σε μια στοίβα.
	Όταν η συνάρτηση επιστρέψει προκειμένου να τις εκτελέσει, οι αναβαλλόμενες κλήσεις της
	εκτελούνται με διάταξη last-in-first-out (η τελευταία αναβληθείσα κλήση εκτελείται πρώτη).

	Για να μάθετε περισσότερα σχετικά με τις αναβλητικές δηλώσεις,
	διαβάστε αυτήν την ανάρτηση :
	https://blog.golang.org/defer-panic-and-recover
*/
func countDefer() {
	// πρώτη προβολή
	fmt.Println("μετράω")

	for i := 0; i < 10; i++ {
		// /τέταρτη προβολή
		defer fmt.Println(i)
	}
	// τρίτη προβολή
	defer fmt.Println("τελείωσα δες παρακάτω:")
	// δεύτερη προβολή
	fmt.Println("οι αριθμοί που μέτρησα")
}
