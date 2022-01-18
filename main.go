package main
import (
    "fmt"
    "os"
)

func zbrajanje() {
	fmt.Println(`
		░█▀▀▀█ ░█▀▀█ ░█▀▀█ ─█▀▀█ ───░█ ─█▀▀█ ░█▄─░█ ───░█ ░█▀▀▀ 
		─▄▄▄▀▀ ░█▀▀▄ ░█▄▄▀ ░█▄▄█ ─▄─░█ ░█▄▄█ ░█░█░█ ─▄─░█ ░█▀▀▀ 
		░█▄▄▄█ ░█▄▄█ ░█─░█ ░█─░█ ░█▄▄█ ░█─░█ ░█──▀█ ░█▄▄█ ░█▄▄▄


	`)

	var x int
	var y int

	fmt.Println("Unesi prvi broj: ")
	fmt.Scan(&x)
	fmt.Println("Unesi drugi broj: ")
	fmt.Scan(&y)

	var r int = x + y 
	fmt.Println("Rezultat ->  ", r)

}

func oduzimanje() {
	fmt.Println(`
		░█▀▀▀█ ░█▀▀▄ ░█─░█ ░█▀▀▀█ ▀█▀ ░█▀▄▀█ ─█▀▀█ ░█▄─░█ ───░█ ░█▀▀▀ 
		░█──░█ ░█─░█ ░█─░█ ─▄▄▄▀▀ ░█─ ░█░█░█ ░█▄▄█ ░█░█░█ ─▄─░█ ░█▀▀▀ 
		░█▄▄▄█ ░█▄▄▀ ─▀▄▄▀ ░█▄▄▄█ ▄█▄ ░█──░█ ░█─░█ ░█──▀█ ░█▄▄█ ░█▄▄▄


	`)

	var x int
	var y int

	fmt.Println("Unesi prvi broj: ")
	fmt.Scan(&x)
	fmt.Println("Unesi drugi broj: ")
	fmt.Scan(&y)

	var r int = x - y 
	fmt.Println("Rezultat ->  ", r)
}

func dijeljenje() {
	fmt.Println(`
		░█▀▀▄ ▀█▀ ───░█ ░█▀▀▀ ░█─── ───░█ ░█▀▀▀ ░█▄─░█ ───░█ ░█▀▀▀ 
		░█─░█ ░█─ ─▄─░█ ░█▀▀▀ ░█─── ─▄─░█ ░█▀▀▀ ░█░█░█ ─▄─░█ ░█▀▀▀ 
		░█▄▄▀ ▄█▄ ░█▄▄█ ░█▄▄▄ ░█▄▄█ ░█▄▄█ ░█▄▄▄ ░█──▀█ ░█▄▄█ ░█▄▄▄


	`)

	var x int
	var y int

	fmt.Println("Unesi prvi broj: ")
	fmt.Scan(&x)
	fmt.Println("Unesi drugi broj: ")
	fmt.Scan(&y)

	var r int = x / y 
	fmt.Println("Rezultat ->  ", r)
}

func mnozenje() {
	fmt.Println(`
		░█▀▄▀█ ░█▄─░█ ░█▀▀▀█ ░█▀▀▀█ ░█▀▀▀ ░█▄─░█ ───░█ ░█▀▀▀ 
		░█░█░█ ░█░█░█ ░█──░█ ─▄▄▄▀▀ ░█▀▀▀ ░█░█░█ ─▄─░█ ░█▀▀▀ 
		░█──░█ ░█──▀█ ░█▄▄▄█ ░█▄▄▄█ ░█▄▄▄ ░█──▀█ ░█▄▄█ ░█▄▄▄


	`)

	var x int
	var y int

	fmt.Println("Unesi prvi broj: ")
	fmt.Scan(&x)
	fmt.Println("Unesi drugi broj: ")
	fmt.Scan(&y)

	var r int = x * y 
	fmt.Println("Rezultat ->  ", r)
}

func main() {
	fmt.Println(`
            
	    ________________________
            
	        [0] Exit
			[1] Zbrajanje
			[2] Oduzimanje
			[3] Dijeljenje
			[4] Množenje
		________________________
		

	`)

	var odgovor int

	fmt.Println("[!] Unesi odgovor ->  ")
	fmt.Scan(&odgovor)

	if(odgovor == 0) {
		os.Exit(1)
	} else if (odgovor == 1) {
		zbrajanje()
	} else if (odgovor == 2) {
		oduzimanje()
	} else if (odgovor == 3) {
		dijeljenje()
	} else if (odgovor == 4) {
		mnozenje()
	}
}
