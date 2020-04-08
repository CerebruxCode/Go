package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {

	// Συναρτήσεις με STRING

	// Ας δούμε μερικές μεθόδους με τις οποίες μπορούμε διαχειριστούμε
	// strings

	sampString := "Hello World"

	// Επιστρέφει αληθής εάν υπάρχει η φράση "lo" στην variable sampString
	fmt.Println(strings.Contains(sampString, "lo"))

	// Επιστρέφει το index (θέση) της φράσης
	fmt.Println(strings.Index(sampString, "lo"))

	// Επιστρέφει αριθμό που μας λέει το πόσες φορές εμφανίζεται το "l" στην variable sampString
	fmt.Println(strings.Count(sampString, "l"))

	// Αντικαθιστά το πρώτο γράμμα με το δεύτερο όσες φορές καθορίζετε (εδώ 3 φορές)
	fmt.Println(strings.Replace(sampString, "l", "x", 3))

	// Επιστρέψτε μια λίστα που χωρίζει με τον καθορισμένο διαχωριστή
	csvString := "1,2,3,4,5,6"
	fmt.Println(strings.Split(csvString, ","))

	listOfLetters := []string{"c", "a", "b"}
	sort.Strings(listOfLetters)
	fmt.Println("Γράμματα:", listOfLetters)

	// Επιστρέφει μια συμβολοσειρά χρησιμοποιώντας τις τιμές που πέρασαν χωρισμένες με διαχωριστικό
	listOfNums := strings.Join([]string{"3", "2", "1"}, ", ")

	fmt.Println(listOfNums)

	// FILE I/O
	/*
		Η πρότυπη βιβλιοθήκη "io" της Go περιλαμβάνει πολλές υλοποιήσεις των διεπαφών διαχείρισης
		αρχείων για εισαγωγή και εγγραφή δεδομένων, συνδέσεων δικτύου, συμπιεστών, κρυπτογράφησης και άλλων.
	*/

	// Δημιουργούμε ένα αρχείο με όνομα samp.txt στην διαδρομή απο όπου εκτελείται η εφαρμογή
	// π.χ. αν τρέξουμε την εφαρμογή <<go run enas-fakelos/kodikas.go>> τότε το αρχείο θα δημιουργηθεί δίπλα απο τον φάκελο
	// "enas-fakelos" και όχι μέσα στον φάκελο δίπλα απο το "kodikas.go"
	file, err := os.Create("samp.txt")

	// Πρόβαλε τυχόν σφάλματα
	if err != nil {
		log.Fatal(err)
	}

	// Γράψτε μια συμβολοσειρά στο αρχείο
	file.WriteString("Αυτό είναι ένα τυχαίο κείμενο")

	// Κλείστε το αρχείο
	file.Close()

	// Προσπάθησε να ανοίξεις το αρχείο. Αν δεν το καταφέρεις πρόβαλε το σφάλμα
	stream, err := ioutil.ReadFile("samp.txt")

	if err != nil {
		log.Fatal(err)
	}

	// Μετατρέψτε σε μια συμβολοσειρά
	readString := string(stream)

	fmt.Println(readString)

	// Αναμονή για είσοδο

	/*
		Παρακάτω, χρησιμοποιούμε την scan μέθοδο η οποία περιμένει απο τον χρήστη να εισάγει κάτι
	*/

	fmt.Println("Πως σε λένε; ")

	var name string

	fmt.Scan(&name)

	fmt.Println("Γειά σου", name, "!!")

	// CASTING
	/*
		Το type casting είναι ένας τρόπος για τη μετατροπή μιας variable από
		έναν τύπο δεδομένων σε έναν άλλο τύπο δεδομένων.
	*/

	randInt := 5
	randFloat := 10.5
	randString := "100"
	randString2 := "250.5"

	// Μετατροπή των τύπων των αριθμών
	fmt.Println(float64(randInt))
	fmt.Println(int(randFloat))

	// Μετατρέψτε μια συμβολοσειρά σε ένα int
	newInt, _ := strconv.ParseInt(randString, 0, 64)
	fmt.Println(newInt)

	// Μετατρέψτε μια συμβολοσειρά σε ένα float
	newFloat, _ := strconv.ParseFloat(randString2, 64)
	fmt.Println(newFloat)

}
