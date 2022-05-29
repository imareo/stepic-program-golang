/*Напишите "функцию голосования", возвращающую то значение (0 или 1), которое среди значений
ее аргументов x, y, z встречается чаще.

Sample Input:0 0 1
Sample Output:0
*/
package main

func vote(x int, y int, z int) int {
	if x+y+z > 1 {
		return 1
	}
	return 0
}

func vote1(x int, y int, z int) int {
	return (x + y + z) / 2
}
