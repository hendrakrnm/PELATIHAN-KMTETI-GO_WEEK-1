package main
import "fmt"

func main() {
    var celcius int
    var pilihan int
    
    fmt.Println("Masukan Suhu dalam celcius: ")
    fmt.Scan( &celcius)
    
    fmt.Println("Pilihan konsversi suhu: ")
    fmt.Println("1. Kelvin  ")
    fmt.Println("2. Reamur  ")
    fmt.Println("3. Fahrenheit ")
    
    fmt.Print("Masukan pilihan[1/2/3]: ")
    fmt.Scan( &pilihan)
    
    
    switch pilihan {
    case 1: 
        kelvin := celcius + 273
        fmt.Printf("Suhu dalam Kelvin: %d", kelvin)
    case 2: 
        reamur := celcius * 4 / 5
        fmt.Printf("Suhu dalam Reamur: %d", reamur)
    case 3: 
        fahrenheit := celcius*9/5 + 32
        fmt.Printf("Suhu dalam Fahrenheit: %d", fahrenheit)
    default:
        fmt.Println("Pilihan yang Anda masukkan salah.")
    }
}