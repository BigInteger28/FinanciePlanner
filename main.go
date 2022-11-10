package main

import (
	"fmt"
)

func main() {
	for {
		var spaargeld float64
		var gepland float64
		var leefgeld float64
		var maandloon float64
		var rechtstreeks float64
		var jaar float64
		var maand float64
		var maandsparen float64
		fmt.Print("\n\nHoeveel spaargeld wil je (+++): €")
		fmt.Scanln(&spaargeld)
		fmt.Print("Geplande aankopen (++): €")
		fmt.Scanln(&gepland)
		fmt.Print("Hoeveel leefgeld per maand heb je nodig: €")
		fmt.Scanln(&leefgeld)
		fmt.Print("Bedrag dat je rechtstreeks op je spaarrekening zet: €")
		fmt.Scanln(&rechtstreeks)
		fmt.Print("Maandelijkse verdienste: €")
		fmt.Scanln(&maandloon)
		fmt.Print("Aantal jaren: ")
		fmt.Scanln(&jaar)
		fmt.Print("Aantal maanden: ")
		fmt.Scanln(&maand)
		maand += (12 * jaar)
		maandsparen = ((spaargeld + gepland - rechtstreeks) / maand)
		fmt.Print("\nMaandelijks moet je €")
		fmt.Printf("%.2f", maandsparen)
		fmt.Print(" sparen")
		if spaargeld > gepland {
			fmt.Print("\nWaarvan €")
			fmt.Printf("%.2f", gepland/maand)
			fmt.Print(" voor geplande aankopen")
			fmt.Print("\nWaarvan €")
			fmt.Printf("%.2f", (spaargeld-rechtstreeks)/maand)
			fmt.Print(" voor spaargeld")
		} else {
			fmt.Print("\nWaarvan €")
			fmt.Printf("%.2f", ((gepland - rechtstreeks) / maand))
			fmt.Print(" voor geplande aankopen")
			fmt.Print("\nWaarvan €")
			fmt.Printf("%.2f", (spaargeld / maand))
			fmt.Print(" voor spaargeld")
		}
		fmt.Print("\nJe hebt €")
		fmt.Printf("%.2f", maandloon-leefgeld-maandsparen)
		fmt.Print(" over voor maandelijkse fun")
	}
}
