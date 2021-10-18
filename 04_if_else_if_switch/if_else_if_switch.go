/*
Τι θα μάθουμε:
- Ομαδοποίηση πολλών import
- Εσωτερικές, εξωτερικές μεταβλητές
- Πως παράγουμε τυχαίους αριθμούς
- Σύνταξη της if else (συνθήκες)
- Σύνταξη της Switch
- Πληροφορίες λειτουργικούς συστήματος
*/

package main

import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
	/*
		Όταν χρειαζόμαστε περισσότερες λειτουργίες μπορούμε να εισάγουμε
		βιβλιοθήκες (έτοιμες functions).

		Εδώ τις ομαδοποιούμε σε μια import χρησιμοποιώντας την παρένθεση
	*/)

func power(x, n, lim float64) float64 { // Φτιάχνουμε μια variable με όνομα power
	// που παίρνει 3 τιμές τύπου float64 και επιστρέφει αποτέλεσμα float64

	/*  Όπως σε άλλες γλώσσες, όταν θέλουμε να εξετάσουμε μια συνθήκη
		η να εκτελέσουμε υπο προυποθέσεις κάποια εργασία, τότε μπορούμε
		να το κάνουμε με την "if".

		Η βασική της χρήση είναι για να αξιολογούμε μια κατάσταση.
		Αν αυτή είναι αληθής, διακλαδίζουμε σε έναν τρόπο όπου
		εκτελούμε κάποιο συγκεκριμένο κώδικα. Αν όμως είναι ψευδής ή μη αληθής,
		διακλαδίζουμε σε έναν άλλο τρόπο.
	    Επομένως στην Go, η συνθήκη που αξιολογούμε πρέπει να αξιολογηθεί ως
		συνθήκη Boolean (true ή false). Δεν μπορούμε να χρησιμοποιήσουμε
		ακέραιους αριθμούς ή συμβολοσειρές για να αντιπροσωπεύσουμε
		μια αληθής ή ψευδής συνθήκη και στη συνέχεια να ελπίζουμε
		ότι η Go θα τα μετατρέψει σε Booleans.

		Οι δηλώσεις "if" της Go είναι όπως και οι βρόχοι "for".

		Η δήλωση (statement) δε χρειάζεται να περιβάλλεται από παρενθέσεις ()
		αλλά απαιτούνται τα αγκύλες {}.

		Όπως και στην for, το if μπορεί να ξεκινήσει με μια σύντομη δήλωση
		variable που θα εκτελεστεί πριν από την συνθήκη δηλαδή χωρίς να
		χρειαστεί να της δημιουργήσουμε εκ των προτέρων.
		Όμως οι variables που ορίζονται στην δήλωση ισχύουν μόνο μέχρι το
		τέλος του if και δεν μπορούν να χρησιμοποιηθούν εκτός.
		Στο παράδειγμά μας μπορείτε να δοκιμάσετε να χρησιμοποιήσετε
		το v στην τελευταία δήλωση επιστροφής (εκτός του βρόγχου if)
	*/

	if v := math.Pow(x, n); v < lim { // ο αριθμός v που θα δώσουμε ισούται
		// με τον αριθμό x εις την n αλλά και μικρότερο του αριθμού lim
		// επέστρεψε μου τον αριθμο v
		return v
	}
	return lim // Δοκίμασε να βάλεις "v"
	/* Σημείωση: Θα παρατηρήσατε οτι κάποια variable ξεκινάνε με κεφαλαίο το
	πρώτο γράμμα, ενώ άλλα με μικρά. Αυτά ονομάζονται exported/un-exported:
	Όταν ξεκινάει με κεφαλαία τότε η μέθοδος είναι προσβάσιμη από άλλα
	αρχεία της go. Αυτό θα το δούμε και σε επόμενα μαθήματα
	*/
}

func power2(x, n, lim float64) float64 {
	/*
		Οι μεταβλητές όμως που δηλώνονται μέσα σε μια σύντομη δήλωση
		if είναι διαθέσιμες μέσα σε οποιοδήποτε άλλο μπλοκ.
		(Και οι δύο κλήσεις προς την "power2" επιστρέφουν
		τα αποτελέσματά τους πριν την κλήση στο fmt.Println της main).
	*/
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	return lim
}

func main() {
	// Ας εκτελέσουμε τις power, power2 που φτιάξαμε παραπάνω, τροφοδοτόντας τις
	// με 3 αριθμούς
	fmt.Println( // μπορούμε να ομαδοποιήσουμε σε ένα fmt.Println με παρένθεση
		"this power:", power(3, 2, 10),
		"\nand that power:", power(3, 3, 20),
		"\nthis power2:", power2(3, 2, 10),
		"\nand that power2:", power2(3, 3, 20),
	)
	fmt.Println("Η ημερομηνία και η ώρα είναι:", time.Now())
	// εκτελούμε την υπορουτίνα Now της time
	fmt.Println("Ο τυχαίος αριθμός μου είναι το:", rand.Intn(10))
	// η Intn επιστρέφει έναν ακέραιο (integer) ψευδο-τυχαίο αριθμό απο 0 έως n
	// για να εμφανίζουμε πραγματικό τυχαίο αριθμό κάθε φορά
	// που τρέχουμε το πρόγραμμα μας, θα χρειαστεί να τροφοδοτήσουμε (rand.Seed)
	// την εφαρμογή με πραγματική τυχαιότητα π.χ. τον χρόνο (time.Now().UnixNano())
	// πηγή: https://pkg.go.dev/math/rand

	rand.Seed(time.Now().UnixNano()) // τροφοδοτούμε τυχαιότητα
	fmt.Println("Ο πραγματικά τυχαίος μου αριθμός είναι το:", rand.Intn(10))

	// το %g είναι θέση που θα πάρει το αποτέλεσμα της εκτέλεσης Sqrt στον αριθμό 7
	fmt.Printf("Τώρα έχεις %g προβλήματα.\n", math.Sqrt(7))
	fmt.Println("Ο αριθμός π =", math.Pi) // η math περιλαμβάνει και άλλες υπορουτίνες
	// προσβάσιμες με math.κάτι πηγή: https://pkg.go.dev/math#pkg-examples

	// Ας δούμε ένα παράδειγμα με if για να ελέγξουμε αν είσαι ενήλικας
	// και μπορείς να οδηγήσεις αυτοκίνητο:
	yourAge := 19 // αλλάξτε το σε ότι αριθμό ηλικίας θέλετε
	fmt.Println("\n === Πρώτη if ===")
	if yourAge >= 18 {
		fmt.Println("Μπορείς να οδηγήσεις (πρώτο if)")
	} else {
		fmt.Println("Δεν μπορείς να οδηγήσεις")
	}

	// Μπορείτε να χρησιμοποιήσετε την "else if" όπως παρακάτω για να
	// ελέγξετε και άλλες περιπτώσεις-συνθήκες, αλλά μόλις μία απο αυτές
	// συμβεί να είναι αληθής τότε οι υπόλοιπες συνθήκες θα αγνοηθούν:
	fmt.Println("\n === Δεύτερη if ===")
	if yourAge >= 18 {
		fmt.Println("Μπορείς να οδηγήσεις (δεύτερο if)")
	} else if yourAge >= 19 {
		fmt.Println("Μπορείς να ψηφίσεις")
	} else {
		fmt.Println("Μπορείς να διασκεδάσεις")
	}

	/*  Ας δούμε όμως πως μπορούμε να γράψουμε το παραπάνω με
	την Switch.

	Μια εντολή switch (διακόπτης) είναι ένας συντομότερος
	τρόπος για να γράψετε μια ακολουθία if-else statements.
	Εκτελεί την πρώτη περίπτωση της οποίας η τιμή είναι
	ίση με την έκφραση της κατάστασης.
	Αν δεν ισχύει καμία τότε εκτελείται η "default" συνθήκη.

	Η switch της Go είναι ίδια με αυτήν στην C, C ++, Java,
	JavaScript και PHP, με εξαίρεση ότι η Go εκτελεί μόνο την
	επιλεγμένη περίπτωση, όχι όλες τις περιπτώσεις που ακολουθούν.
	Επίσης σε αντίθεση με τις γλώσσες αυτές, στη Go δεν είναι
	απαραίτητη η ύπαρξη εντολής διακοπής (brake)
	του βρόχου switch γιατί εφόσον πληρείται η συνθήκη,
	ο βρόχος τερματίζει αυτόματα.

	Μια άλλη σημαντική διαφορά είναι ότι οι περιπτώσεις switch της Go
	δεν χρειάζεται να είναι σταθερές και οι τιμές που απαιτούνται
	δεν πρέπει να είναι ακέραιοι.
	*/
	fmt.Println("\n === Switch ===")
	switch yourAge {
	case 18:
		fmt.Println("Μπορείς να οδηγήσεις μέσω switch")
	case 19:
		fmt.Println("Μπορείς να ψηφίσεις μέσω switch")
	default:
		fmt.Println("Μπορείς να διασκεδάσεις μέσω switch")
	}

	// Όπως θα δείτε, εδω καλούμε την myOS() ενώ έχει
	// δημιουργηθεί μετά την "func main". Στην Go δεν έχει σημασία
	// αν οι δικές σας συναρτήσεις θα μπούνε πριν η μετά την "func main()"
	// σε αντίθεση με άλλες γλώσσες που πρέπει να δηλώνονται στην αρχή
	// του προγράμματος
	myOS()

}

func myOS() {
	fmt.Print("Η Go τρέχει σε ")
	arch := runtime.GOARCH
	fmt.Printf("%s / ", arch)
	switch os := runtime.GOOS; os {
	// για να δεις όλα τα διαθέσιμα λειτουργικά της GOOS τρέξε "go tool dist list"
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		// freebsd, openbsd,
		// plan9, windows...
		fmt.Printf("%s\n", os)
	}
}
