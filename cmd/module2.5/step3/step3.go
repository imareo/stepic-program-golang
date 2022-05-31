package main

import "fmt"

func main() {
	exampleByteSlice()
}

func exampleByteSlice() {
	// Используем тип string, который конвертируем в []byte,
	// чтобы не использовать абстрактные цифры
	bs := []byte("Байтовый срез")
	bsL := []byte("Byte slice")

	// 2 байта на нелатиницу
	fmt.Printf("%s %v\n", bs, bs)
	fmt.Printf("%s %v\n", bs[:1], bs)
	fmt.Printf("%s %v\n", bs[:2], bs)
	fmt.Printf("%s %v\n", bs[:3], bs)
	fmt.Printf("%s %v\n", bs[:4], bs)

	fmt.Printf("%s %v\n", bsL, bsL)
	fmt.Printf("%s %v\n", bsL[:1], bsL)
	fmt.Printf("%s %v\n", bsL[:2], bsL)
	fmt.Printf("%s %v\n", bsL[:3], bsL)
	fmt.Printf("%s %v\n", bsL[:4], bsL)

	bsL[0] = bsL[1]
	fmt.Printf("%s", bsL)
}
