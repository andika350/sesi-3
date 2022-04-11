package main

import (
	"fmt"
	"math"
	"strings"
)

func main() {
	// 	greet, isTrue, err := hello("Hacktiv","Jakarta")

	// 	fmt.Println("Greet:", greet)
	// 	fmt.Println("Name is Hacktiv8?", isTrue)
	// 	fmt.Println("Location:", err)

	// }

	// func hello(name string, address string)  (string, bool, error){
	// 	msg := fmt.Sprintf("Hello, %s yang berlokasi di %s", name, address)

	// 	isTrue := false
	// 	var err error
	// 	if name == "Hacktiv8" {
	// 		isTrue = true
	// 	}

	// 	if address != "Jakarta" {
	// 		err = fmt.Errorf("lokasi %s bukan di jakarta", address)
	// 	}

	// 	return msg, isTrue, err

	// qualified("Dika",20,"male")

	greet("Dika", 25)

	personal("Andika", "Jakarta")

	var names = []string{"Ariell", "Dika", "Cindy"}
	var printMsg = hei("Hola", names)

	fmt.Println(printMsg)

	var diameter float64 = 15

	var luas, keliling = lingkaran(diameter)

	fmt.Printf("%.3f\n", luas)
	fmt.Printf("%.3f\n", keliling)

	var sisi1, sisi2, sisi3, tinggi float64
	sisi1 = 3
	sisi2 = 4
	sisi3 = 2
	tinggi = 5

	var luasS, kelilingS = calculate(sisi1, sisi2, sisi3, tinggi)

	fmt.Println(luasS)
	fmt.Println(kelilingS)

	studentlists := print("Dika", "Cindy", "Rehan", "Hanief")

	fmt.Printf("%v\n", studentlists)

	numberLists := []int{1, 2, 3, 4, 5, 6, 7, 8}

	result := sum(numberLists...)
	fmt.Println("Result: ", result)

}

//====================functions====================

func sum(numbers ...int) int {
	total := 0

	for _, v := range numbers {
		total += v
	}

	return total
}

func print(names ...string) []map[string]string {

	var result []map[string]string

	for i, v := range names {
		key := fmt.Sprintf("student%d", i+1)
		temp := map[string]string{
			key: v,
		}
		result = append(result, temp)
	}

	return result
}

func hei(msg string, names []string) string { //function with return value

	var joinStr = strings.Join(names, " ")

	var result string = fmt.Sprintf("%s %s", msg, joinStr)

	return result

}

func lingkaran(d float64) (float64, float64) { //function returning multiple values

	var luas float64 = math.Pi * math.Pow(d/2, 2)

	var keliling = math.Pi * d

	return luas, keliling

}

func calculate(sisi1, sisi2, sisi3, tinggi float64) (luas float64, keliling float64) { /*predefine return value, deklarasi variable pada returned value
	tidak perlu memanggil variable lagi pada return di akhir fungsi*/

	luas = sisi2 * tinggi / 2

	keliling = sisi1 + sisi2 + sisi3

	return
}

func greet(name string, age int8) {
	fmt.Printf("Hello, my name is %s and I'm %d years old!\n", name, age)
}

func personal(name, address string) {
	fmt.Printf("My name is %s. and I lived in %s\n", name, address)
}

//Latihan laki-laki umur 17+ lolos, perempuan 20+ lolos predefine

// func qualified(nama string, umur int, gender string) (notif string, string error) {

// 	msg := fmt.Sprintf("Hello, %s berjenis kelamin %s, dengan umur %d kamu dinyatakan lolos!", nama, gender, umur)

// 	var err error

// 	if gender == "male" && umur > 17{

// 	} else if gender == "female" && umur > 20 {

// 	} else {
// 		err = fmt.Errorf("Tidak memenuhi syarat")
// 	}

// 	return msg, err
// }
