# GoLang

#### What is Go?
&nbsp;&nbsp;&nbsp;&nbsp; It is a cross-platform, open source, high-performance, fast, statiscally typed, complied langage that feels like dynamically typed, inerpreted language. It was developed at Google by Robert Griesemer, Rob Pike, and Ken Thompson in 2007.

#### What is Go used for?
* Web development (server-side)
* Developing network-based programs
* Developing cross-platform enterprise applications
* Cloud-native development 
> Go Features
>* Supports concurrency through goroutines and channel
>* Has automatic garbage collection
>* Does not support classes and objects
>* Does not support inheritance
>&nbsp;


#### Comments 
##### &nbsp;&nbsp;&nbsp;&nbsp; Single Line Comment
```
package main;
import ("fmt");

func main(){
    // This is one line Commment
    fmt.Println("GO Tutorials From W3 School");
}
```
##### &nbsp;&nbsp;&nbsp;&nbsp; Multi-Line Comment
```
package main;
import ("fmt");

func main(){
    /*
        This
        is 
        MutliLine Comment
    */
    fmt.Println("GO Tutorials From W3 School");
}
```

##### Formatting verbs for Printf
**%v**	: Prints the value in the default format
**%#v**	: Prints the value in Go-syntax format
**%T**	: Prints the type of the value
**%%**	: Prints the % sign


##### Integer Formatting verbs
**%b**	=> Base 2
**%d**	=> Base 10
**%+d**	=> Base 10 and always show sign
**%o**	=> Base 8
**%O**	=> Base 8, with leading 0o
**%x**	=> Base 16, lowercase
**%X**	=> Base 16, uppercase
**%#x**	=> Base 16, with leading 0x
**%4d**	=> Pad with spaces (width 4, right justified)
**%-4d** => Pad with spaces (width 4, left justified)
**%04d** => Pad with zeroes (width 4

##### String Formatting verbs
**%s**	Prints the value as plain string
**%q**	Prints the value as a double-quoted string
**%8s**	Prints the value as plain string (width 8, right justified)
**%-8s**	Prints the value as plain string (width 8, left justified)
**%x**	Prints the value as hex dump of byte values
**% x**	Prints the value as hex dump with spaces

##### Float Formatting verbs
**%e**	Scientific notation with 'e' as exponent
**%f**	Decimal point, no exponent
**%.2f**	Default width, precision 2
**%6.2f**	Width 6, precision 2
**%g**	Exponent as needed, only necessary digits


#### Arrays
```
var array_name = [length]datatype{values} // here length is defined
or
var array_name = [...]datatype{values} // here length is inferred
```
d .
#### Slices
Slices are similar to arrays, but are more powerful and flexible.
Like arrays, slices are also used to store multiple values of the same type in a single variable.
However, unlike arrays, the length of a slice can grow and shrink as you see fit.
In Go, there are several ways to create a slice:
&nbsp;&nbsp;&nbsp;&nbsp;Using the []datatype{values} format
&nbsp;&nbsp;&nbsp;&nbsp;Create a slice from an array
&nbsp;&nbsp;&nbsp;&nbsp;Using the make() function

##### &nbsp;&nbsp;Creating a Slice
```
slice_name := []datatype{values}
myslice := myarray[start:end] // A slice made from the array
slice_name := make([]type, length, capacity) // A slice from make function
slice_name = append(slice_name, element1, element2, ...) // append element to a slice
slice3 = append(slice1, slice2...) // append a slice to another slice  
'...' after slice2 is necessary 
copy(dest, src)
```

#### Loops(The only loop for loop)
Loops are handy if you want to run the same code over and over again, each time with a different value.
Each execution of a loop is called an iteration. 
```
for statement1; statement2; statement3 {
   // code to be executed for each iteration
}
```
##### &nbsp;&nbsp;&nbsp; Range
&nbsp;&nbsp;&nbsp;&nbsp;_The range keyword is used to more easily iterate over an array, slice or map. It returns both the index and the value._
```
for index, value := array|slice|map {
   // code to be executed for each iteration
}
```


#### Functions 
```
func FunctionName(param1 type, param2 type) type {
  // code to be executed
  return output
}
```
&nbsp;&nbsp;&nbsp; ***return the value with a naked return***
```
func FunctionName(param1 type, param2 type) (output type) {
  // code to be executed
  return 
}
```

#### Struct i.e. Structure
A struct (short for structure) is used to create a collection of members of different data types, into a single variable.
While arrays are used to store multiple values of the same data type into a single variable, structs are used to store multiple values of different data types into a single variable.
A struct can be useful for grouping data together to create records.
```
type Person struct {
  name string
  age int
  job string
  salary int
}
```
