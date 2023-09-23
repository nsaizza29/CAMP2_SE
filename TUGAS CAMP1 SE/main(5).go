package main

import (
    "fmt"
)

func hitungEnergiPotensial(massa float64, gravitasi float64, ketinggian float64) float64 {
    energiPotensial := massa * gravitasi * ketinggian
    return energiPotensial
}

func hitungEnergiKinetik(massa float64, kecepatan float64) float64 {
    energiKinetik := 0.5 * massa * kecepatan * kecepatan
    return energiKinetik
}

func main() {
    massa := 10.0       // Massa benda (kg)
    gravitasi := 9.81   // Percepatan gravitasi di Bumi (m/s^2)
    ketinggian := 5.0   // Ketinggian dari permukaan tanah (m)
    kecepatan := 7.0    // Kecepatan benda (m/s)

    energiPotensial := hitungEnergiPotensial(massa, gravitasi, ketinggian)
    energiKinetik := hitungEnergiKinetik(massa, kecepatan)

    fmt.Printf("Energi Potensial: %.2f Joule\n", energiPotensial)
    fmt.Printf("Energi Kinetik: %.2f Joule\n", energiKinetik)
}
