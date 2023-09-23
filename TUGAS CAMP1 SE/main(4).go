package main

import "fmt"

func celsiusToFahrenheit(celsius float64) float64 {
    return (celsius * 9/5) + 32
}

func fahrenheitToCelsius(fahrenheit float64) float64 {
    return (fahrenheit - 32) * 5/9
}

func celsiusToKelvin(celsius float64) float64 {
    return celsius + 273.15
}

func kelvinToCelsius(kelvin float64) float64 {
    return kelvin - 273.15
}

func fahrenheitToKelvin(fahrenheit float64) float64 {
    celsius := fahrenheitToCelsius(fahrenheit)
    return celsiusToKelvin(celsius)
}

func kelvinToFahrenheit(kelvin float64) float64 {
    celsius := kelvinToCelsius(kelvin)
    return celsiusToFahrenheit(celsius)
}

func main() {
    // Contoh penggunaan fungsi konversi
    celsius := 25.0
    fmt.Printf("%.2f°C = %.2f°F\n", celsius, celsiusToFahrenheit(celsius))

    fahrenheit := 77.0
    fmt.Printf("%.2f°F = %.2f°C\n", fahrenheit, fahrenheitToCelsius(fahrenheit))

    kelvin := 298.15
    fmt.Printf("%.2fK = %.2f°C\n", kelvin, kelvinToCelsius(kelvin))

    fmt.Printf("%.2f°F = %.2fK\n", fahrenheit, fahrenheitToKelvin(fahrenheit))

    fmt.Printf("%.2fK = %.2f°F\n", kelvin, kelvinToFahrenheit(kelvin))
}
