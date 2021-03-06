/*
Επανάληψη :
Κάθε πρόγραμμα Go ξεκινά με μια "Package Declaration" (δήλωση πακέτου) που παρέχει έναν τρόπο για
επαναχρησιμοποίηση του κώδικα. Η 'main' είναι ο "πυρήνας" του προγράμματος απο όπου
ξεκινάει η εφαρμογή μας. Έτσι αφού οι εφαρμογές αποτελούνται απο αρχεία κώδικα
που ανήκουν σε πακέτα, όλες οι εφαρμογές έχουν μια μόνο main. Είναι η πύλη εισόδου
της εφαρμογής απο όπου ξεκινάει η εκτέλεση της.
Όλα τα υπόλοιπα αρχεία έχουν και αυτά δήλωση για το σε ποιό package ανήκουν
*/
package main

// Το import επιτρέπει την εισαγωγή βιβλιοθηκών και την εισαγωγή κώδικα από
// άλλα πακέτα (π.χ. απο github θα γράφαμε << import "github.com/CerebruxCode/Go" >>)
// Εδώ η βιβλιοθήκη fmt που περιλαμβάνεται στην standrad library και που εισάγουμε
// μας παρέχει εργαλεία μορφοποίησης - προβολής δεδομένων. Επειδή είναι στην standrad library
// την καλούμε με το όνομά της και όχι με την online διεύθυνσή της. Το ίδιο ισχύει και για την runtime

import (
	"fmt"     // προφέρεται ως "φαμτ" ή formatter
	"runtime" // την χρειαζόμαστε για να προβάλουμε πληροφορίες του συστήματος στο οποίο τρέχει η εφαρμογή μας.
)

// ============== Σχόλια στον κώδικα ========================
//Ξεκινώντας μια γραμμή με τα "//" μπορούμε να προσθέσουμε ένα προσωπικό σχόλιο σε μια γραμμή
/*
 Αν θέλουμε πολλαπλές γραμμές όπως αυτό που διαβάζετε τώρα,
 τότε βάζουμε το σχόλιο μας ανάμεσα σε αυτά τα
*/

// Όπως και σε άλλες γλώσσες, έτσι και στην Go μπορούμε να συνδυάσουμε ενσωματωμένες εντολές ώστε να
// δημιουργήσουμε νέες μεθόδους και λειτουργίες. Αυτές ονομάζονται functions

func add(x int, y int) int { // δημιουργούμε μια δικιά μας function (συνάρτηση) ξεκινώντας την εντολή με την
	// λέξη func και με ένα όνομα (π.χ. add) η οποία εδώ θα την κάνουμε να κάνει μια απλή αριθμητική
	// πράξη προσθήκης δύο αριθμών. Η func add παίρνει 2 ορίσματα x και y
	// τύπου integer (ακέραιος αριθμός) και δίνουν έξοδο (αποτέλεσμα πράξης) πάλι int
	// πηγή: https://blog.golang.org/gos-declaration-syntax

	return x + y // Η έξοδος: προσθέτει δυο αριθμούς τύπου int που θα δοθούν στο add(x,y)

	/*  TIP: εφόσον τα ορίσματα είναι καί τα δύο int θα μπορούσαμε επίσης να τα γράψουμε με ένα int
	όπως π.χ.

	func add(x, y int) int {
		return x + y
	}

	*/

}

// Οι functions του προγράμματός μας ξεκινούν όπως είπαμε με τη λέξη κλειδί "func" και περιβάλλουν τον κώδικα μέσα σε { }
// Η main είναι η function που εκτελείται όταν τρέχετε το πρόγραμμά σας.
// Θα εκτελεστεί αυτόματα το πρόγραμμα, χωρίς να χρειάζεται να το ξεκινήσετε γράφοντας π.χ. main() στο τέλος του κώδικα
// όπως απαιτείται σε άλλες γλώσσες.

func main() { // η κύρια συνάρτηση που εκτελεί την εφαρμογή

	/*	Println

		Η Println είναι και αυτή μια function μέσα απο το πακέτο fmt που "φορμάρει"- διαμορφώνει - εξάγει - τυπώνει
		μια συμβολοσειρά (strings), η οποία περιβάλλεται από διπλά εισαγωγικά
		π.χ. "Γειά σου κόσμε", σε μια νέα γραμμή στην οθόνη
	*/
	fmt.Printf("Το Λειτουργικό σου σύστημα: %s\nΗ Αρχιτεκτονική του: %s\n", runtime.GOOS, runtime.GOARCH)
	// εκτελούμε την Printf που βρίσκεται μέσα στην fmt βιβλιοθήκη που εισάγαμε και αντικαθιστά τα %s %s με
	// τα αποτελέσματα των εντολών GOOS και GOARCH της βιβλιοθήκης runtime.
	// η εντολή \n λέει στην Printf να τυπώσει το αμέσως επόμενο στοιχείο σε νέα γραμμή

	fmt.Println("Γειά σου κόσμε") // εκτελούμε την Println function που βρίσκεται
	// μέσα στην fmt βιβλιοθήκη που εισάγαμε.

	fmt.Println(add(42, 13))
	/*
		εδώ καλούμε την δικιά μας function (add) που είχαμε
		δημιουργήσει και την χρησιμοποιούμε για να προσθέσουμε τους δύο int αριθμούς x,y
		και με την βοήθεια της fmt.Println να προβάλουμε το αποτέλεσμα στην οθόνη

		Μπορείτε να μάθετε για την Println αλλά και άλλες "built-in functions of standard library"
		πληκτρολογώντας σε ένα τερματικό:

		go doc fmt Println

	*/

	/*	VARIABLES

		Οι μεταβλητές (Variables) γράφονται στατικά (statically typed), πράγμα που σημαίνει ότι ο τύπος
		(type - είδος - π.χ αριθμός/γράμμα/δεκαδικό κλπ) τους δεν μπορεί να αλλάξει (static).
		Τα ονόματα μεταβλητών πρέπει να ξεκινούν με ένα γράμμα και μπορεί να περιέχουν γράμματα,
		αριθμούς ή την κάτω παύλα _

		Για τους αριθμούς έχουμε τους παρακάτω τύπους:
		Μια int (integer - ακέραιος) είναι ένας θετικός ή αρνητικός αριθμός χωρίς δεκαδικά ψηφία

		uint8  : ονομάζεται unsigned  8-bit integers (0 έως 255)
		uint16 : ονομάζεται unsigned 16-bit integers (0 έως 65535)
		uint32 : ονομάζεται unsigned 32-bit integers (0 έως 4294967295)
		uint64 : ονομάζεται unsigned 64-bit integers (0 έως 18446744073709551615)
		uint, uintptr σε 32bit επεξεργαστές είναι uint32 και uint64 σε 64bit

		int8  : ονομάζεται signed  8-bit integers (-128 έως 127)
		int16 : ονομάζεται signed 16-bit integers (-32768 έως 32767)
		int32 : ονομάζεται signed 32-bit integers (-2147483648 έως 2147483647)
		int64 : ονομάζεται signed 64-bit integers (-9223372036854775808 έως 9223372036854775807)
		το int σε 32bit επεξεργαστές είναι int32 και int64 σε 64bit
	*/

	var age int64 = 40

	// Αντίθετα ένας float είναι ένας αριθμός με δεκαδικά ψηφία:
	// float32, float64

	var favNum float64 = 1.61803398875

	// Στην πραγματικότητα δε χρειάζεται να καθορίσετε τον τύπο δεδομένων, ούτε χρειάζεστε
	// ερωτηματικό στο τέλος όπως ίσως έχετε δει σε άλλες γλώσσες προγραμματισμού:

	randNum := 1 // Η Go καταλαβαίνει οτι πρόκειται για var τύπου int
	fmt.Println(randNum)

	/*
		Δεν μπορείτε, ωστόσο αργότερα, να εκχωρήσετε έναν μη συμβατό τύπο στην μεταβλητή
		π.χ. να αλλάξετε το υφιστάμενο randNum που είναι int να το κάνετε

			randNum = "Hello"

		Στην Println εκτός απο strings μπορείτε να χρησιμοποιήσετε
		μεταβλητές (το κενό διάστημα προστίθεται αυτόματα οπότε δεν χρειάζεται)
	*/

	fmt.Println(age, " ", favNum)

	var numOne = 1.000 // Η Go καταλαβαίνει οτι πρόκειται για float
	var num99 = .999   // Η Go καταλαβαίνει οτι πρόκειται για float

	/*
		ΑΡΙΘΜΗΤΙΚΕΣ ΠΡΑΞΕΙΣ

		Μπορείτε επίσης να κάνετε αριθμητικές πράξεις μέσα σε Println (Σημειώστε ότι οι float αριθμοί δεν είναι ακριβείς)

		Διαθέσιμες αριθμητικές πράξεις : +, -, *, /, %
		το % ή αλλιώς modulus - modulo ή remainder είναι το υπόλοιπο μιας ακέραιας διαίρεσης
	*/
	fmt.Println(numOne - num99)

	a := 6
	b := 4.0 // float
	fmt.Println("6 + 4 =", 6+4)
	fmt.Println("6 - 4 =", 6-4)
	fmt.Println("6 * 4 =", 6*4)
	fmt.Println("6 / 4 =", 6/4)
	fmt.Println("6 % 4 =", 6%4)
	// το % operator (ακέραιας διαίρεσης) είναι διαθέσιμο μόνο για integers. Άρα θα πρέπει να μετατρέψουμε το
	// float σε int όπως παρακάτω.

	fmt.Println("6 % int(4.0) =", a%int(b)) // χρησιμοποιώ την function int() για να μετατρέψω το 4.0 (float) σε 4 (int)

}
