// package main

// import (
	// "fmt"
	// "errors"
	// "encoding/json"
	// "begin-with-go/exemplo"
	// "time"
// )

// type vehicle interface {
// 	start() string
// }

// type Car struct {
// 	CarName string `json:"name"`
// 	CarYear int    `json:"year"`
// }

// func (c Car) start() string {
// 	return "Comecou"
// }

// func (c Car) drive()  {
// // func (c *Car) drive()  {
// 	// c.CarName = "XPTWO"
// 	// Coloca essa estrela e ele aponta para a referencia na memoria
// 	fmt.Println(c.CarName, "Andou")
// }

// func exemplo(car vehicle) {
// }

// func main() {

	// for i := 0; i <= 100; i++ {
	// 	fmt.Println(i)
	// 	time.Sleep(time.Second * 3)
	// }

	// casa := exemplo.casa{
	// 	Cor: "Rosa",
	// 	Numero: 2680,
	// }

	// car := Car {
	// 	CarName: "Fusion",
	// 	CarYear: 2020,
	// }
	
	// exemplo(car)

	// j := []byte(`{"name":"BMW","year":2020}`)

	// var car Car

	// json.Unmarshal(j, &car)

	// fmt.Println(car.CarName)

	// car1 := Car{
	// 	CarName: "Fusion",
	// 	CarYear: 2020,
	// }
	// car1.drive()
	// fmt.Println(car1.CarName)

	// result, _ := json.Marshal(car1)
	// fmt.Println(string(result))

	// x := 10
	// res := abc(&x)
	// fmt.Println(res)

	// array com dez posições
	// var array [10]int
	// array[0] = 1

	// x := [10]int{0,1,2,3,4,5,6,7,8,9}
	// fmt.Println(x[0])

	// slice := make([]int, 5)
	// slice[0] = 1
	// slice[1] = 1
	// slice[2] = 1
	// slice[3] = 1
	// slice[4] = 1
	// fmt.Println(slice)
	// slice = append(slice, 0, 1, 2, 3)
	// fmt.Println(slice)

	// funcAnonima := func() int {
	// 	return 1
	// }
	// fmt.Println(funcAnonima())

	// a, b, err := nome("Pedro", 1)
	// if err != nil {
	// 	panic("panico meu deus")
	// }
	// fmt.Println(a, b)

	// m := make(map[string]int)
	// m["pedro"] = 10
	// m["larrisa"] = 11

	// fmt.Println(m["pedro"])


// }

// func nome(a string, b int) (string, int, error) {
// 	if b > 10 {
// 		return "", 0, errors.New("Deu ruim")
// 	}
// 	return a, b, nil
// }

// func abc(a *int) int {
// 	return *a
// }