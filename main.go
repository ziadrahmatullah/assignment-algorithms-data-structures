package main

import "fmt"

// If you want to test your code, you may test it here...
func main() {
	// Task 1.a
	
	var criminalss = []Person{
		{Name:"Melissa Garcia", CriminalScore:51},
		{Name:"Derek Perez", CriminalScore:24},
		{Name:"Laura Hall", CriminalScore:42},
		{Name:"Laura Matthews", CriminalScore:52},
		{Name:"Debra Norris", CriminalScore:94},
		{Name:"Amanda Griffin", CriminalScore:69},
		{Name:"David Mendoza", CriminalScore:51},
		{Name:"Emily Smith", CriminalScore:63},
		{Name:"Sara Miranda", CriminalScore:16},
		{Name:"Christopher Gallagher", CriminalScore:54},
		{Name:"Stephanie Mason", CriminalScore:4},
		{Name:"Heather Tucker", CriminalScore:82},
		{Name:"Erin Brown", CriminalScore:79},
		{Name:"Jessica Greer", CriminalScore:73},
		{Name:"Joseph Garcia", CriminalScore:86},
		{Name:"Gavin Daniel", CriminalScore:91},
		{Name:"Gloria Hill", CriminalScore:53},
		{Name:"Cindy Perez", CriminalScore:78},
		{Name:"Mallory Snow", CriminalScore:31},
		{Name:"Ellen Gallegos", CriminalScore:70},
		{Name:"Anthony Johnson", CriminalScore:15},
}
	vehicle, waiting := LastDayInJail(criminalss, "Ellen Gallegos")
	fmt.Print(vehicle, waiting)
	// Q := queue{}
	// Q.enqueue(criminals[0])
	// Q.enqueue(criminals[1])
	// Q.enqueue(criminals[2])
	// Q.enqueue(criminals[3])
	// Q.enqueue(criminals[4])
	// Q.enqueue(criminals[5])
	// Q.enqueue(criminals[6])
	// Q.enqueue(criminals[7])
	// Q.enqueue(criminals[8])
	// Q.enqueue(criminals[9])
	// count := 0
	// for Q.head != nil{
	// 	fmt.Println(Q.dequeue())
	// }
	// fmt.Print(count)


	// curr := Q.head.next.next
	// fmt.Print(curr.val.Name)
	// hasil := Q.dequeue()
	// fmt.Print(hasil)
	// Q.getPerson("Laura Hall")
	// fmt.Print(Q.head.next)
	// fmt.Print(Q.len)
	// fmt.Print(Q.head.val.Name)
	// fmt.Print(Q.head.next.val.Name)
	// fmt.Print(Q.head.next.next.val.Name)


	// Task 2.a
	// result := RotateImage([][]int{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{7, 8, 9},
	// })
	// fmt.Print(result)

	// Task 2.b
	// RunRotateActualImage()

	// Task 3.a
	// string := RobotTranslatorV2("LLRRX XX")
	// fmt.Printf(string)
}
