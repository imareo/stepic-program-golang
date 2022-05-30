/*
Вам необходимо реализовать структуру со свойствами-полями On, Ammo и Power, с типами bool,
int, int соответственно. У этой структуры должны быть методы: Shoot и RideBike, которые не
принимают аргументов, но возвращают значение bool.

Если значение On == false, то оба метода вернут false.

Делать Shoot можно только при наличии Ammo (тогда Ammo уменьшается на единицу, а метод
возвращает true), если его нет, то метод вернет false. Метод RideBike работает также, но
только зависит от свойства Power.

Чтобы проверить, что вы все сделали правильно, вы должны создать указатель на экземпляр этой
структуры с именем testStruct в функции main, в дальнейшем программа проверит результат.
*/
package main

type gun struct {
	On    bool
	Ammo  int
	Power int
}

func (g *gun) Shoot() bool {
	canShot := g.On && g.Ammo > 0
	if canShot {
		g.Ammo--
	}
	return canShot
}

func (g *gun) RideBike() bool {
	canShot := g.On && g.Power > 0
	if canShot {
		g.Power--
	}
	return canShot
}

func main() {

	testStruct := new(gun)

	// solution above this comment
	_ = testStruct
}
