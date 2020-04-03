package main

import "fmt"

func main() {

	// Μπορούμε επίσης να αποθηκεύσουμε πολλαπλά στοιχεία σε ένα χάρτη

	superhero := map[string]map[string]string{
		"Superman": map[string]string{
			"realname": "Clark Kent",
			"city":     "Metropolis",
		},

		"Batman": map[string]string{
			"realname": "Bruce Wayne",
			"city":     "Gotham City",
		},
	}

	// Μπορούμε να εξάγουμε δεδομένα όπου το κλειδί ταιριάζει με το Superman

	if temp, hero := superhero["Superman"]; hero {
		fmt.Println(temp["realname"], temp["city"])
	}

}
