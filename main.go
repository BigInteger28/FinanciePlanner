package main

import (
	"fmt"
)

func sparen(kosten float64, spaargeld float64, tijd float64) float64 {
	return (kosten + spaargeld) / tijd
}

func main() {
	for {
		var vasteKosten float64
		var kosten float64
		var spaargeld float64
		var tijd float64
		var verdien float64
		var anderverdien float64
		fmt.Print("\nHoeveel verdien je maandelijks: €")
		fmt.Scanln(&verdien)
		fmt.Print("Hoeveel verdien je extra wat niet maandelijks is: €")
		fmt.Scanln(&anderverdien)
		fmt.Print("Leefkosten per maand: €")
		fmt.Scanln(&vasteKosten)
		fmt.Print("Niet maandelijkse verplichte kosten: €")
		fmt.Scanln(&kosten)
		temp := kosten
		fmt.Print("Geplande kosten van dingen die je wil kopen: €")
		fmt.Scanln(&kosten)
		kosten += temp
		fmt.Print("Geld dat je wil sparen: €")
		fmt.Scanln(&spaargeld)
		fmt.Print("Hoeveel jaar: ")
		fmt.Scanln(&tijd)
		jaar := 12 * tijd
		fmt.Print("Hoeveel maand: ")
		fmt.Scanln(&tijd)
		tijd += jaar
		verdien += (anderverdien / tijd)
		maandelijks := sparen(kosten, spaargeld, tijd)
		fmt.Print("Maandelijks heb je €", vasteKosten, " aan leefkosten")
		fmt.Print("\nMaandelijks moet je ", maandelijks, " opzij zetten voor geplande aankopen en spaargeld")
		fmt.Print("\nMaandelijks heb je ", verdien-maandelijks-vasteKosten, " over")
	}
}
