package main

import (
	"log"
	"net/http"
)

// Δημιουργία HTTP Server
/*
Στο παρόν πρόγραμμά μας φτιάχνουμε έναν εξυπηρετητή ο οποίος θα "σερβίρει" στους επισκέπτες μας,
μια έτοιμη στατική ιστοσελίδα που βρίσκεται στον φάκελο ./static (η τελεία ./ είναι για να πούμε οτι
ο φάκελος βρίσκεται εδώ που τρέχει και ο κώδικάς μας.)

*/

func main() {
	// Εδώ χρησιμοποιούμε ένα handle στο οποίο λέμε κάλεσε την http.FileServer για να
	// προβάλει ότι βρίσκεται στον φάκελο ./static
	// Σημείωση: Η ιστοσελίδα που παρέχουμε διαθέτει και βίντεο που εκτελείται αυτόματα. Θα χρειαστεί
	// να το επιτρέψετε στον browser σας για να το δείτε.
	http.Handle("/", http.FileServer(http.Dir("./static")))
	// Καλούμε την log.Fatal για να καταγράψει σφάλματα αφού ξεκινήσει ο server μας να περιμένει τους
	// επισκέπτες του
	log.Fatal(http.ListenAndServe(":8081", nil))
}

/* Σχετικά με την log.Fatal και γιατί την χρησιμοποιούμε

Σύμφωνα με τις UNIX συμβάσεις, μια διαδικασία που αντιμετωπίζει ένα σφάλμα θα πρέπει
"να αποτύχει όσο το δυνατόν νωρίτερα, με μη μηδενικό κωδικό εξόδου".
Αυτό μας οδηγεί στις ακόλουθες οδηγίες για να χρησιμοποιήσω το log.Fatal όταν:

1) ... ένα σφάλμα συμβαίνει σε οποιοδήποτε από τα func, καθώς αυτά συμβαίνουν
όταν επεξεργάζονται οι import ή πριν την func main() αντίστοιχα.
Για παράδειγμα, έχω ρυθμίσει την καταγραφή σφαλμάτων για να ελέγξω ότι έχουμε
ένα φυσιολογικό περιβάλλον και παραμέτρους. Οπότε δεν έχουμε σωστές παραμέτρους
τότε δεν χρειάζεται να τρέχουμε το func main() οπότε θα πρέπει να το πούμε νωρίς
στο χρήστη οτι δεν έχει δώσει σωστές παραμέτρους.

2)... συμβαίνει ένα σφάλμα το οποίο γνωρίζω ότι δεν μπορεί λόγο εξωγενών παραγόντων
να παραμερίσει. Ας υποθέσουμε ότι έχουμε ένα πρόγραμμα που δημιουργεί μια μικρογραφία
ενός μεγάλου αρχείου εικόνας που δίνεται η διαδρομή του αρχείου αυτού στη γραμμή εντολών.
Εάν αυτό το αρχείο δεν υπάρχει ή είναι ακατάλληλο για ανάγνωση λόγω π.χ. ανεπαρκών δικαιωμάτων,
δεν υπάρχει λόγος να συνεχιστεί το πρόγραμμα. Συνεπώς, τηρούμε τις UNIX συμβάσεις
και αποτυγχάνουμε.

3)... εμφανίζεται ένα σφάλμα κατά τη διάρκεια μιας διαδικασίας που μπορεί να
μην είναι αναστρέψιμη. Ας υποθέσουμε ότι έχουμε μια "μη διαδραστική/αυτόματη"
εφαρμογή αντιγραφής φακέλων σε κάποιον προορισμό. Τώρα, ας υποθέσουμε
ότι το πρόγραμμά μας συναντάει ένα αρχείο στον φάκελο προορισμού που έχει το ίδιο όνομα
(αλλά διαφορετικό περιεχόμενο) με το αρχείο που θα αντιγραφεί εκεί.
Δεδομένου ότι δεν μπορούμε να ζητήσουμε από τον χρήστη να αποφασίσει τι
να κάνει το πρόγραμμά μας με αυτό το αρχείο (μη διαδραστική/αυτόματη) και δεν μπορούμε
να αντιγράψουμε αυτό το αρχείο, έχουμε ένα πρόβλημα.
Επειδή όταν τελειώσει το πρόγραμμά μας θα γίνει μηδενικό κωδικό εξόδου (επιτυχώς),
ο χρήστης θα υποθέσει ότι οι πηγές και οι φάκελοι προορισμού είναι ακριβή αντίγραφα
μεταξύ τους, αφού δεν μπορούμε απλά να παραλείψουμε το εν λόγω αρχείο.
Επίσης, δεν μπορούμε απλά να το αντικαταστήσουμε, δεδομένου ότι αυτό
θα μπορούσε ενδεχομένως να καταστρέψει τις πληροφορίες.
Αυτή είναι μια κατάσταση που είναι μη αναστρέψιμη από κάθε ρητό αίτημα του χρήστη
και γι 'αυτό θα χρησιμοποιούσα το log.Fatal για να εξηγήσω την κατάσταση,
υπακούοντας έτσι στην UNIX αρχή να αποτύχει όσο το δυνατόν νωρίτερα.
πηγή: https://stackoverflow.com/questions/33885235/should-a-go-package-ever-use-log-fatal-and-when
*/
