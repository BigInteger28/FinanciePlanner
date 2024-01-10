package main

import (
    "fmt"
)

func main() {
    for {
        var monthlyIncome, otherAnnualIncome, fixedMonthlyCosts, plannedAnnualExpenses, annualSavings float64

        // User inputs
        fmt.Print("Maandelijks inkomen: ")
        fmt.Scanln(&monthlyIncome)

        fmt.Print("Extra jaarlijks inkomen: ")
        fmt.Scanln(&otherAnnualIncome)

        fmt.Print("Vaste verplichte maandelijkse kosten: ")
        fmt.Scanln(&fixedMonthlyCosts)

        fmt.Print("Geplande jaarlijkse uitgaven (b.v., reizen, geplande aankopen): ")
        fmt.Scanln(&plannedAnnualExpenses)

        fmt.Print("Jaarlijks doel om te sparen: ")
        fmt.Scanln(&annualSavings)

        // Aanpassingen
        annualSavingsRequired := annualSavings - otherAnnualIncome
        if annualSavingsRequired < 0 {
            plannedAnnualExpenses += -annualSavingsRequired
            annualSavingsRequired = 0
        }

        // Calculations
        totalAnnualIncome := (monthlyIncome * 12)
        totalPlannedAnnualExpenses := plannedAnnualExpenses
        totalAnnualSavings := annualSavingsRequired
        totalExpenses := (fixedMonthlyCosts * 12) + totalPlannedAnnualExpenses
        monthlyFunBudget := (totalAnnualIncome - totalExpenses - totalAnnualSavings) / 12

        // Output
        fmt.Printf("\nMaandelijks naar Fun Rekening: €%.2f\n", monthlyFunBudget)
        fmt.Printf("Maandelijks laten staan op Zichtrekening (Vaste verplichte kosten): €%.2f\n", fixedMonthlyCosts)
        fmt.Printf("Maandelijks naar Geplande Spaarrekening: €%.2f\n", plannedAnnualExpenses/12)
        fmt.Printf("Maandelijks naar Lange Termijn Spaarrekening: €%.2f\n", totalAnnualSavings/12)
        fmt.Printf("Jaarlijkse bedrag extra sparen op Lange Termijn Spaarrekening: €%.2f\n", totalAnnualSavings)
        fmt.Println("\n\n")
    }
}
