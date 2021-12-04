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

}