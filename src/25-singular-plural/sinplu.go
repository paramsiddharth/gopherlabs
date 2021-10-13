package main

import (
	"golang.org/x/text/feature/plural"
	"golang.org/x/text/language"
	"golang.org/x/text/message"
)

func main() {
	message.Set(language.Hindi, "%d वस्तुएँ बाक़ी हैं ।",
		plural.Selectf(1, "%d",
			"=0", "कुछ भी नहीं बचा है ।",
			plural.One, "एक वस्तु बाक़ी हैं ।",
			"<10", "%[1]d वस्तुएँ बाक़ी हैं ।",
			plural.Other, "बहुत कुछ बाक़ी है ।",
		),
	)

	prt := message.NewPrinter(language.Hindi)

	prt.Printf("%d वस्तुएँ बाक़ी हैं ।", 11)
	prt.Println()
	prt.Printf("%d वस्तुएँ बाक़ी हैं ।", 3)
	prt.Println()
	prt.Printf("%d वस्तुएँ बाक़ी हैं ।", 2)
	prt.Println()
	prt.Printf("%d वस्तुएँ बाक़ी हैं ।", 1)
	prt.Println()
	prt.Printf("%d वस्तुएँ बाक़ी हैं ।", 0)
	prt.Println()
}
