package main

import (
	"root/home/is105sem03/mycrypt"
	"fmt"
)

func main() {

	kryptertRune := mycrypt.Krypter([]rune("Kjevik;SN39040;18.03.2022 01:50;6"), mycrypt.ALF_SEM03, 4)
	kryptertString := string(kryptertRune)
	fmt.Println(kryptertString)

	dekryptertRune := mycrypt.Krypter(kryptertRune, mycrypt.ALF_SEM03, -4)
	dekryptertString := string(dekryptertRune)
	fmt.Println(dekryptertString)

}
