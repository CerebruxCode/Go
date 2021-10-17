/*
Καλώς ήρθες!

Για να βρίσκεσαι εδώ σημαίνει ότι έχεις εγκαταστήσει με επιτυχία την Go στον
υπολογιστή σου σύμφωνα με τις οδηγίες στο https://wp.me/pq2ce-hNG . 
Έχεις επίσης βρει το "go workspace" το οποίο είναι
ο βασικός χώρος εργασίας της Go που βρίσκεται στη διαδρομή "~/go/src".
Τέλος έχεις μπει και στο συγκεκριμένο φάκελο του πρώτου μαθήματος ο οποίος
ονομάζεται "00_package_main" και περιλαμβάνει το αρχείο "00_package_main.go"
που αυτή τη στιγμή διαβάζεις.

Το μάθημα ξεκινάει !
Κάθε πρόγραμμα/αρχείο go αποτελείται σχεδόν πάντα απο 4 στοιχεία:
(1) Package declaration
(2) Imports
(3) Function (main)
(4) Code

(1) Το πρώτο πράγμα που χρειάζεται κάθε πρόγραμμα Go είναι ένα "package declaration".
Το βασικό μας πρόγραμμα θα είναι πάντα το "main", με το οποίο λέμε στην ουσία στον
compiler (μεταγλωττιστή) να το μεταγλωττίσει ως αυτόνομο εκτελέσιμο πρόγραμμα.
Όλα τα υπόλοιπα Go προγράμματα που θα γράψουμε και δε θα τους δώσουμε
το "main" declaration, χρησιμοποιούνται ως κοινόχρηστες βιβλιοθήκες, τις οποίες
μπορούμε να καλέσουμε στο πρόγραμμά μας με το "import".
Επομένως, κάθε εκτελέσιμο που θέλουμε να γράψουμε χρειάζεται το "package main"
ως την πρώτη γραμμή κώδικα.

(2) Το επόμενο πράγμα που χρειαζόμαστε είναι το import (εισαγωγή) με το οποίο μπορούμε
να καλέσουμε τρίτες βιβλιοθήκες στο πρόγραμμά μας, Είτε τις έχουμε γράψει εμείς,
είτε είτε έρχονται από τις έτοιμες βιβλιοθήκες της Go (Go Standard Libraries).
Στην περίπτωσή μας έχουμε μόνο την "fmt". Τα περισσότερα IDE εισάγουν/αφαιρούν
αυτόματα τις βιβλιοθήκες με βάση τον κώδικα που έχουμε γράψει.

(3) Το τρίτο πράγμα που χρειαζόμαστε είναι μια main function (κύρια συνάρτηση).
Μέσα στη συνάρτηση main (func main) βρίσκεται η ουσία της εφαρμογής μας και εκεί
βρίσκεται ο βασικός κώδικας που εκτελείται όταν την τρέχουμε. Έτσι, αυτός ο συνδυασμός
"main package" και "func main()" είναι αυτό που ενεργοποιεί, δίνει ζωή και λειτουργεί
το πρόγραμμά μας.

Άρα ότι έχουμε μέσα στο main package αρχείο μας είναι το σημείο εισόδου του εκτελέσιμου
προγράμματος μας. Φυσικά το προγραμμά μας μπορεί να μην αποτελείται μόνο απο αυτό το
αρχείο αλλά και πάρα πολλά άλλα αρχεία, βιβλιοθήκες κλπ. Σε αυτή την περίπτωση όμως
το πρόγραμμα "package_main" αποτελείται μόνο απο το "package_main.go" αρχείο.

(4) Όταν εκτελούμε απο το τερματικό το πρόγραμμά μας (θα δείτε παρακάτω πώς), αυτό
μπαίνει στο main package και βρίσκει την main function για να εκτελέσει το
Code (κώδικα) που περιλαμβάνει. Για τις functions θα μιλήσουμε εκτενώς
αργότερα οπότε εδώ το πρόγραμμά μας το μόνο που κάνει είναι να λέει το
κλασσικό "Γειά σου κόσμε".

*/
package main // (1)

import "fmt" // (2)

func main() { //(3) η κύρια συνάρτηση που εκτελεί την εφαρμογή

	fmt.Println("Γειά σου κόσμε") // (4) εκτελούμε την Println function που βρίσκεται
	// μέσα στην fmt βιβλιοθήκη που εισάγαμε στο (2).
}

/*
Ένα πρόγραμμα όπως αυτό, μπορούμε να το τρέξουμε σε τερματικό γράφοντας:

	$ go run package_main.go

Η απάντηση που θα πάρουμε από το τερματικό είναι:

	Γειά σου κόσμε

Για να φτιάξουμε ένα εκτελέσιμο αρχείο απο το πρόγραμμά μας και να το
μοιράσουμε θα πρέπει να το κάνουμε compile (μεταγλώττιση) ως εξής:

	$ go build package_main.go

Δίνουμε την ακόλουθη εντολή για να δούμε τα αρχεία που δημιουργήσαμε:

	$ ls

Το τερματικό θα μας απαντήσει το εξής:

	package_main    package_main.go

Το εκτελέσιμο αρχείο μας είναι το "package_main" και μπορούμε να το
τρέξουμε με την εντολή:

	$ ./package_main

Η απάντηση που θα πάρουμε είναι:

	Γειά σου κόσμε

Η Go υποστηρίζει seamless cross compilation για Windows, macOS, Linux και BSD.
Επομένως, δεν έχει σημασία σε ποια πλατφόρμα αναπτύσσουμε και μεταγλωττίζουμε τον κώδικα.
Ακολουθεί ένα παράδειγμα της σύνταξης που χρησιμοποιούμε για να δημιουργήσουμε
εκτελέσιμα αρχεία για Windows και Linux σε macOS περιβάλλον:

Για Windows 32bit μπορούμε να δημιουργήσουμε εκτελέσιμα ως εξής:

	$ GOOS=windows GOARCH=386 go build -o appname.exe

Για Linux 64bit μπορούμε να δημιουργήσουμε εκτελέσιμα ως εξής:

	$ GOOS=linux GOARCH=amd64 go build -o appname.linux

Περισσότερα σχετικά με τα λειτουργικά συστήματα και τις αρχιτεκτονικές για τις οποίες μπορούμε να δημιουργήσουμε εκτελέσιμα
προγράμματα μπορείτε να διαβάστε εδώ: https://golang.org/doc/install/source#environment

*/
