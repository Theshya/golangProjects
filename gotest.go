package main

import "fmt"

func GetData() (int, int) {
	return 100, 200
}

func main() {
	a, b := 5, 8
	var attack = 40
	var defence = 20
	var damageRate float32 = 0.17
	var damage = float32(attack-defence) * damageRate
	var m int = 100
	var n int = 200
	n, m = m, n
	fmt.Println(m, n)
	fmt.Println(damage)
	fmt.Println("helloworld")
	fmt.Println(a, b)
	o, _ := GetData() // _ 为匿名变量
	_, p := GetData()
	fmt.Println(o, p)

}
