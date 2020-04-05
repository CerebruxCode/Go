package main

/*
	Όταν χρειαζόμαστε περισσότερες λειτουργίες μπορούμε να εισάγουμε
	βιβλιοθήκες (έτοιμες ρουτίνες). Εδώ τις ομαδοποιούμε σε μια import
	χρησιμοποιώντας την παρένθεση
*/
import (
	"fmt"
	"math"
	"math/rand"
	"runtime"
	"time"
)

/*
	Οι δηλώσεις "if" της Go είναι όπως και οι βρόχοι "for".
	Η δήλωση (statement) δε χρειάζεται να περιβάλλεται από παρενθέσεις ()
	αλλά απαιτούνται τα άγκυστρα {}.
*/

func main() {

	fmt.Println("Καλώς όρισες στην Go!")
	fmt.Println("Η ημερομηνία και η ώρα είναι:", time.Now()) // εκτελούμε την υπορουτίνα Now της time
	fmt.Println("Ο αγαπημένος μου αριθμός είναι το:", rand.Intn(10))
	/* η Intn επιστρέφει έναν ακέραιο (integer) ψευδο-τυχαίο αριθμό απο 0 έως n*/

	fmt.Println(math.Pi)
	/*exported/un-exported: Όταν ξεκινάει με κεφαλαία τοτε η υπορουτίνα είναι
	προσβάσιμη από άλλα αρχεία της go*/

	fmt.Printf("Τώρα έχεις %g προβλήματα.\n", math.Sqrt(7))
	/*το %g είναι θέση που θα πάρει το αποτέλεσμα της εκτέλεσης Sqrt στο 7*/

	yourAge := 18

	if yourAge >= 16 {
		fmt.Println("You Can Drive by first if")
	} else {
		fmt.Println("You Can't Drive")
	}

	// Μπορείτε να χρησιμοποιήσετε την "else if" όπως παρακάτω για να
	// ελέγξετε και άλλες περιπτώσεις-συνθήκες, αλλά μόλις μία απο αυτές
	// είναι αληθής οι υπόλοιπες συνθήκες θα αγνοηθούν:

	if yourAge >= 16 {
		fmt.Println("You Can Drive by second if")
	} else if yourAge >= 18 {
		fmt.Println("You Can Vote")
	} else {
		fmt.Println("You Can Have Fun")
	}

	/*
		Μια εντολή switch είναι ένας συντομότερος τρόπος για να γράψετε μια ακολουθία
		if-else statements. Εκτελεί την πρώτη περίπτωση της οποίας η τιμή είναι
		ίση με την έκφραση της κατάστασης. Αν δεν ισχύει καμία τότε εκτελείται η
		default σηνθήκη.

		Ο διακόπτης(switch) της Go είναι σαν αυτόν που συναντάται στις C, C ++, Java, JavaScript και PHP,
		εκτός του το ότι η Go εκτελεί μόνο την επιλεγμένη περίπτωση,
		όχι όλες τις περιπτώσεις που ακολουθούν. Σε αντίθεση με τις γλώσσες αυτές, στη Go δεν είναι
		απαραίτητη η ύπαρξη εντολής διακοπής του βρόχου switch γιατί εφόσον πληρούται η συνθήκη ο βρόχος
		τερματίζει αυτόματα.

		Μια άλλη σημαντική διαφορά είναι ότι οι περιπτώσεις switch της Go
		δεν χρειάζεται να είναι σταθερές και οι τιμές που απαιτούνται
		δεν πρέπει να είναι ακέραιοι.
	*/

	switch yourAge {
	case 16:
		fmt.Println("Go Drive by switch")
	case 18:
		fmt.Println("Go Vote by switch")
	default:
		fmt.Println("Go Have Fun by switch")
	}

	// Όπως θα δείτε, εδω καλούμε την pow και την myOS() ενώ έχει
	// δημιουργηθεί μετά την func main. Στην Go δεν έχει σημασία
	// αν οι δικές σας συναρτήσεις θα μπούν πριν η μετά την "func main()"
	// σε αντίθεση με άλλες γλώσσες που πρέπει να δηλώνονται στην αρχή
	// του προγράμματος

	fmt.Println(
		"this pow:", pow(3, 2, 10),
		"\nand that pow:", pow(3, 3, 20),
		"\nthis pow2", pow2(3, 2, 10),
		"\nand that pow2", pow2(3, 3, 20),
	)

	myOS()

}

func pow(x, n, lim float64) float64 {
	/*
		Όπως και στην for, η if μπορεί να ξεκινήσει με μια σύντομη δήλωση
		μεταβλητής που θα εκτελεστεί πριν από την συνθήκη.
		Οι μεταβλητές που ορίζονται στην δήλωση ισχύουν μόνο μέχρι το
		τέλος του if.
		(Δοκιμάστε να χρησιμοποιήσετε το v στην τελευταία δήλωση επιστροφής.)
	*/
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim // Δοκίμασε να βάλεις "v"
}

func pow2(x, n, lim float64) float64 {
	/*
		Οι μεταβλητές που δηλώνονται μέσα σε μια σύντομη δήλωση είναι
		επίσης διαθέσιμες μέσα σε οποιοδήποτε άλλο μπλοκ.
		(Και οι δύο κλήσεις προς την "pow2" επιστρέφουν τα αποτελέσματά τους
		πριν την κλήση στο fmt.Println της main).
	*/
	if v := math.Pow(x, n); v < lim {
		return v
	} else {
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here, though
	return lim
}

func myOS() {
	fmt.Print("Your Go runs on ")
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
