package main

import "fmt"

/*
	Οι συναρτήσεις (func, functions) μας επιτρέπουν την επαναχρησιμοποίηση κώδικα και
	συμβάλουν στην δημιουργία μιας δομής στον κώδικά μας

	Είναι της μορφής:
	func funcName (parametersPassed) returnType {
		....
	}

	οι μεταβλητες που χρησιμοποιούνται εντος της συνάρτησης
	δε δύναται να χρησιμοποιηθούν εκτός αυτής

	Παρακάτω ξεκινάμε δημιουργώντας μερικές δικές μας
*/

func addThemUp(numbers []float64) float64 {

	sum := 0.0

	for _, val := range numbers {

		// Συντομογραφία για sum = sum + val
		sum += val

	}

	return sum

}

// Οι Go functions μπορούν να επιστρέψουν πολλές τιμές

func next2Values(number int) (int, int) {

	return number + 1, number + 2

}

// Μπορείτε να λάβετε έναν απροσδιόριστο αριθμό τιμών με args ... int

func subtractThem(args ...int) int {

	finalValue := 0

	for _, value := range args {
		finalValue -= value
	}

	return finalValue

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
	αλλά η κλήση της δεν εκτελείται μέχρι
	να γίνει η επιστροφή της περιβάλλουσας function.
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

func main() {

	listOfNums := []float64{1, 2, 3, 4, 5}

	fmt.Println("Sum :", addThemUp(listOfNums))

	// Λάβετε 2 τιμές από μια function

	num1, num2 := next2Values(5)

	fmt.Println(num1, num2)

	// Στείλτε έναν αόριστο αριθμό τιμών σε μια function
	// https://en.wikipedia.org/wiki/Variadic_function (Variadic Function)

	fmt.Println(subtractThem(1, 2, 3, 4, 5))

	// Μπορείτε να δημιουργήσετε μια συνάρτηση func μέσα σε μια func. Έχει πρόσβαση στις
	// τοπικές μεταβλητές της func στην οποια περιλαμβάνεται
	// Μια func όπως αυτή, χωρίς τοπικές μεταβλητές είναι ένα "closure"

	num3 := 3

	doubleNum := func() int {

		num3 *= 2
		return num3

	}

	fmt.Println(doubleNum())
	fmt.Println(doubleNum())

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
	εκτελούνται με διαταξη last-in-first-out (η τελευταία αναβληθείσα κλήση εκτελείται πρώτη).

	Για να μάθετε περισσότερα σχετικά με τις αναβλητικές δηλώσεις,
	διαβάστε αυτήν την ανάρτηση ιστολογίου :
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
