package main

import "fmt"
import "time"

// If you want to test your code, you may test it here...
func main() {
	// Task 1.a
	
	// var criminalss = []Person{
	// 	{Name:"Budi", CriminalScore:51},
	// 	{Name:"budi", CriminalScore:51},
	// 	{Name:"Catur", CriminalScore:51},
		// {Name:"Laura Matthews", CriminalScore:52},
		// {Name:"Debra Norris", CriminalScore:94},
		// {Name:"Amanda Griffin", CriminalScore:69},
		// {Name:"David Mendoza", CriminalScore:51},
		// {Name:"Emily Smith", CriminalScore:63},
		// {Name:"Sara Miranda", CriminalScore:16},
		// {Name:"Christopher Gallagher", CriminalScore:54},
		// {Name:"Stephanie Mason", CriminalScore:4},
		// {Name:"Heather Tucker", CriminalScore:82},
		// {Name:"Erin Brown", CriminalScore:79},
		// {Name:"Jessica Greer", CriminalScore:73},
		// {Name:"Joseph Garcia", CriminalScore:86},
		// {Name:"Gavin Daniel", CriminalScore:91},
		// {Name:"Gloria Hill", CriminalScore:53},
		// {Name:"Cindy Perez", CriminalScore:78},
		// {Name:"Mallory Snow", CriminalScore:31},
		// {Name:"Ellen Gallegos", CriminalScore:70},
		// {Name:"Anthony Johnson", CriminalScore:15},
// }
	start:= time.Now()
	vehicle, waiting := LastDayInJail(crim, "Ellen Gallegos")
	fmt.Print(vehicle, waiting)
	elapsed:= time.Since(start)
	fmt.Print(elapsed)

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
