package main;
import ("fmt");

func main(){
	const PI float64 = 3.14;
	var x1,y1,x2,y2 int = 1,1,4,4;


	/* To Print we Have
		Print, Println, Printf
	*/
	fmt.Println("Value of PI:",PI);
	fmt.Println("Distance between points:",x2-x1 + y2-y1)
	fmt.Println("GO Tutorials From W3 School")

	// Formatting Verbs For Strings
	fmt.Printf("\n\n%v\n","Hey, I'm Siddharth Choudhary")
	fmt.Printf("%#v\n","Hey, I'm Siddharth Choudhary")
	fmt.Printf("%T\n","Hey, I'm Siddharth Choudhary")
	fmt.Printf("%s\n","Hey, I'm Siddharth Choudhary")
	fmt.Printf("%q\n","Hey, I'm Siddharth Choudhary")
	fmt.Printf("%8s\n","Hey, I'm Siddharth Choudhary")
	fmt.Printf("%-8s\n","Hey, I'm Siddharth Choudhary")
	fmt.Printf("%x\n","Hey, I'm Siddharth Choudhary")
	// fmt.Printf("%t\n","Hey, I'm Siddharth Choudhary")

	// Arrays
	var zPoints = [5]int{4:5,1:2};
	var xPoints = [2]int{1,1};
	var yPoints = [...]int{4,5};

	zPoints[0] = 5;
	fmt.Println(zPoints,yPoints,xPoints)
	fmt.Println(len(xPoints))

	// Slices 
	myarray := [...]int{1,2,3,4,5,6,7,8,9,10,11};
	myslice := myarray[2:7] // A slice made from the array
	fmt.Printf("slice: %v, lenght:%v, capacity:%v\n",myslice,len(myslice),cap(myslice)) // empty slice with 3 capacity and size 3

	// Slice from make function
	myslice1 := make([]int,5,10)
	fmt.Println(myslice1)

	//Swtich Cases
	var day int = 4;

	switch day {
	case 1,3,5:
	 fmt.Println("Odd weekday")
	case 2,4:
	  fmt.Println("Even weekday")
	case 6,7:
	 fmt.Println("Weekend")
   default:
	 fmt.Println("Invalid day of day number")
   }
}