//Code has a control flow in which it executes such as sequence which means it executes from top to bottom.

//Types of control flow
//(1) sequence - top to bottom
//(2) loop; iterative
//(3) conditional


//Sequence example
package main

import "fmt"
func main() {
	fmt.Println("Hello, World!")
}
//and
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	foo()
}

func foo() {
	fmt.Println("I'm in foo")
}

//////////////
//if you don't want your code to end and you want to make more functions you can do this

package main
import "fmt"

func main () {
	fmt.Printline("Main function here")
	stuff() //new function called stuff
}

func stuff() {
	fmt.Println("Second function")
}

//you can also add extra functions in additional functions

package main
import "fmt"

func main () {
	fmt.Printline("Main function here")
	stuff() //new function called stuff
}

func stuff() {
	fmt.Println("Second function")
	libby()
}

func libby() {
	fmt.Println("Third function")
}
//all these extra functions get added to the main function and will play in sequence from top to bottom


//Iterative example 
package main

import "fmt"

func main() {
	fmt.Println("Hello, World!")
	foo()

	for i := 0; i < 100; i++ {//iterative example
		if i%2 == 0 {//conditional example
			fmt.Println(i)
		}
	}
}

func foo() {
	fmt.Println("I'm in foo")
}
//---------------------------------------------------------------------------

//Declaration operators 

package main

import "fmt"

func main() {
	x := 42
	fmt.Println(x)
}

//if i want to change the value of a declared var you would do this

package main

import "fmt"

func main() {
	x := 42 // x : is i declare this var while = 42 is giving it a value. := declare var and add value
	fmt.Println(x)
	x = 99
	fmt.Println(x)
}

//You cannot use reserved words/keywords from the golang as a var
//for example i cant declare a var name as break because break is a Golang keyword


package main

import "fmt"

func main() {
	//Declare a variable and assign value
	//short declaration operator
	x := 42
	fmt.Println(x)// fmt is the package while Println is the print line command we imported from the fmt package

}

//If you want to declare a variable outside of a function body you must use var. 
//exmample

package main

import "fmt"

var y = 500

func main() {
	x := 150 //always try to use a short declaration operator := as much as you can
	fmt.Println(x)
	fmt.Println(y)
}

//another example

package main

import "fmt"

var y = 500

func main() {
	x := 150
	fmt.Println(x)
	fmt.Println(y)

	foo()
}

func foo() {
	fmt.Println("again:", y)
}

//x only exists in the function body while y is outside of the function body by using var

package main//You cannot start a code/script without starting this way first!

import "fmt"

var y = 500

var z int //declares a var with name z but no value so it uses its default value is 0 since the default value of var is 0. 

func main() {// Var z was given type int
	x := 150
	fmt.Println(x)
	fmt.Println(y)
	foo()
	fmt.Println(z)
}

func foo() {
	fmt.Println("again:", y)
}

//Default values are given when you don't assign a value to a var
//Default values are:
//Boolean: false
//Int or integers: 0
//Floats: 0.0
//Strings: ""
//and for all the following
//Pointers, Functions, Interfaces, Slices, Channels and maps: Nil

//This initialization is done recursively(repeated application: repetition/recurrence), so for instance each element of an array of struct
//will have all of its fields zeroed if no value is specified.

//Go is a static programming language not a dynamic one. Java is an example of dynamic.
//VARIABLE is DECLARED to hold a VALUE of a certain type ie: STRING, INT, BOOL, FLOAT and so on.
package main

import "fmt"

var y = 42//DECLARED type INT
var z = "shaken, not stirred"//DECLARED type STRING

func main(){
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	//z = 43 ----this wont work because z was already DECLARED as a STRING
	fmt.Println(z)
	fmt.Printf("%T\n", z)
}

//There are also raw sting literals which allow you to print `words "and more words"`
//It will print with "" if you include them
//EXAMPLE
package main

import "fmt"

var a = `Whats up, "I hope you are having a good day"`

func main() {
	fmt.Println(a)
}
//OUTPUT = Whats up, "I hope you are having a good day"

//You can also make it print many lines
//Example
package main

import "fmt"

var a = `1
2
skip a few
99
100`

func main() {
	fmt.Println(a)
}
//OUTPUT is multiline just like how its typed into the code

package main

import "fmt"

var y string
//DECLARE a VARIABLE to be of a certain TYPE
//and then ASSIGN a VALUE of that TYPE to the variable
func main() {
	fmt.Println("printing the value of y", y, "ending")
	fmt.Printf("%T", y)
}

//OUTPUT = printing the value of y  ending because y holds no value yet

package main

import "fmt"

var y string

func main() {
	fmt.Println("printing the value of y", y, "endings")
	fmt.Printf("%T\n", y)
	y = "because i love happy"
	fmt.Println("printing the value of y", y, "endings")
	fmt.Printf("%T", y)
}
//fmt.Printf("%T") is printing the type of var. In this case its STRING
//fmt.Println("%T\n") \n allows to add a new line.

package main

import "fmt"

var y string
var z int
func main() {
	fmt.Println("printing the value of y", y, "endings")
	fmt.Printf("%T\n", y)
	
	y = "because i love happy"
	
	fmt.Println("y")
	fmt.Printf("%T", y)
	
	fmt.Println(z)
	fmt.Printf("%T", z)
}
//this will print our string, the TYPE of y "STRING" the value of z "0" and the TYPE of z INT
//This is useful to declare the TYPE of var being used

package main

import "fmt"

var z = 26

func main() {
	fmt.Println(z)
	fmt.Printf("%#x\n", z)//this outputs 26 in hex 0x1a
	//you can also do 
	fmt.Printf("%#x\t%b\t%x",z ,z, z) //which will print out %#x - hex, %b - binary and %x - regular hexadecimal
}// \t is how you separate each command and you must give 3 variable's since we are requesting 3 outputs
//OUTPUT 0x1a	11010	1a which all == 26


package main

import "fmt"

var z = 26
var y = "Hex | Binary | Hexadecimal"
func main() {
	fmt.Println(z)
	fmt.Println(y)
	fmt.Printf("%#x\t%b\t%x\n",z ,z, z)

	s := fmt.Sprintf("%#x\t%b\t%x",z ,z, z)//adding sprintf which will output to a string that can be assigned a var
// s in this case is equal to the output of fmt.Sprint
	fmt.Println(s)
}

//Println - print line
//Printf - print type (type of var like int, bool, and float for example)

//DECLARING my own TYPE on a VAR
package main

import "fmt"

var a int
type hotdog int
//declaring hotdog as an integer and naming it b
var b hotdog//here we declare our integer(hotdog) as b

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)
}

//EXAMPLE of type declaring WITH CONVERSION added
package main

import "fmt"

var a int
type hotdog int
//declaring hotdog as an integer and naming it b
var b hotdog//here we declare our integer(hotdog) as b

func main() {
	a = 42
	b = 43
	fmt.Println(a)
	fmt.Printf("%T\n", a)
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	a = int(b)//this is how we would convert a to the value of b. a = type(var)
	fmt.Println(a)
	fmt.Printf("%T\n", a)
}


//Hands on exercise #1
/*Using the short declaration operator, ASSIGN these VALUES to VARIABLES with the IDENTIFIERS “x” and “y” and “z”
42
“James Bond”
true
Now print the values stored in those variables using 
a single print statement
multiple print statements*/


package main

import "fmt"

func main() {

	x := 42
	y := "James Bond"
	z := true

	fmt.Println(x, y, z)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}

//Hands on exercise #2
/*Use var to DECLARE three VARIABLES. The variables should have package level scope. Do not assign VALUES to the variables. Use the following IDENTIFIERS for the variables and make sure the variables are of the following TYPE (meaning they can store VALUES of that TYPE).
identifier “x” type int
identifier “y” type string
identifier “z” type bool
in func main
print out the values for each identifier
The compiler assigned values to the variables. What are these values called?
*/

package main

import "fmt"

var x int
var y string
var z bool

func main(){

	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Printf("%T\n", z)

}

//Default value will be 0, nil, false

//Hand on Exercise #3
/*Using the code from the previous exercise,
At the package level scope, assign the following values to the three variables
for x assign 42
for y assign “James Bond”
for z assign true
in func main
use fmt.Sprintf to print all of the VALUES to one single string. ASSIGN the returned value of TYPE string using the short declaration operator to a VARIABLE with the IDENTIFIER “s”
print out the value stored by variable “s”*/

package main

import "fmt"

var	x int = 42
var	y string = "James Bond"
var	z bool = true

func main(){
	s := fmt.Sprintf("%v%v%v", x, y, z)
	fmt.Println(s)
}

//Hands on Exercise #4
/*For this exercise
Create your own type. Have the underlying type be an int.
create a VARIABLE of your new TYPE with the IDENTIFIER “x” using the “VAR” keyword
in func main
print out the value of the variable “x”
print out the type of the variable “x”
assign 42 to the VARIABLE “x” using the “=” OPERATOR
print out the value of the variable “x”*/

package main

import "fmt"

	type naruto int
	var x naruto

func main() {
	
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}

//Hands on Exercise #5
/*Building on the code from the previous example
at the package level scope, using the “var” keyword, create a VARIABLE with the IDENTIFIER “y”. The variable should be of the UNDERLYING TYPE of your custom TYPE “x”
eg:
----type hotdog int----
----var x hotdog-------
----var y int----------

in func main
this should already be done
print out the value of the variable “x”
print out the type of the variable “x”
assign your own VALUE to the VARIABLE “x” using the “=” OPERATOR
print out the value of the variable “x”
now do this
now use CONVERSION to convert the TYPE of the VALUE stored in “x” to the UNDERLYING TYPE
then use the “=” operator to ASSIGN that value to “y”
print out the value stored in “y”
print out the type of “y”*/

package main

import "fmt"

type naruto int
var x naruto
var y int

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)

	y = int(x)
	fmt.Println(y)
	fmt.Printf("%T", y)
}

//Hands on Exercise #6
//this was a test. I scored 28/36 points. What i got wrong

//(1.) The smallest standalone element of a program that expresses some action to be carried out is a statement

//(2.) A combination of one or more explicit values, constants, variables, operators, and functions
//that the programming language interprets and computes to produce another value is an expression. 

//(3.) A composite data type allows you to compose together values of other data types.

//(4.) In "2+2" the "2"s are the OPERANDS while + is the operator

//(5.) Package is a keyword in Golang

//(6.) The three words to describe good package names in the "effective go" document are short, concise, and evocative.

//(7.) Using - allows you to throw away returns/send them into the void. This character tells
//the compiler that you aren't going to use a value returned by a function

//Go Language Types: Bool, int, floats, strings, arrays, slices, maps, channels. structs, pointers etc

//Boolean - a type that holds a true and false value.
package main

import "fmt"

var x bool //unassigned value of bool is false. 

func main() {
	fmt.Println(x)
}

//the output of this will be false since no value was given to x

package main

import "fmt"

func main() {
a := 7
b := 42
	fmt.Println(a == b)//you can use these operators called equality comparison operators
}//Here are a list of ECO's: ==(does it equal) <=(less than or equal to) >=(greater than or equal to) !=(not equal) <(less than) >(greater than)
//there is no tripple equal or === in go. only single and double == and =

//== is seeing if something is equal to another while = is assignment or giving a value to something

//Computers work in a binary way ie 0 and 1 where 0 is off and 1 is on.  
// 1 = 2
// 2 = 4
// 3 = 8
// 4 = 16
// 5 = 32
// 6 = 64
// 7 = 128
// 8 = 256
//These are the basic numbers you typically see when dealing with pcs. x86-x64 = 32 bit and 64 bit for example

//example
//if 2 = 4 then we use 2 numbers and have 4 possible outcomes
// 11
// 10
// 01
// 00
//and if we increase it to 4 = 16 you get 16 outcomes
// 0000
// 0001
// 0010
// 0011
// 0100
// 0101
// 0110
// 0111
// 1111
// 1000
// 1001
// 1010
// 1011
// 1100
// 1101
// 1110
// these are called coding schemes.
// 0s and 1s make up something called binary code/digit(s)
// binary digits = bits (bi)nary digi(ts)
// 


// ascii was one of the most popular forms of coding
// ascii is connected to binary code for example the following code
// 100 0001 = A and 100 0010 = B 


// The most popular coding scheme today is UTF-8 which has taken over the roll of ascii
// The first part of UTF-8 contains ascii and incorporates it and the guys who made UTF-8 also made the Go programming language
// 

//Measurement of bits
// 1 bit - or single binary digit
// 8 bits = byte
// 1000 bytes = kilobyte - really 1024 bytes
// 1000 kb = megabyte
// 1000 mb = gigabyte
// 1000 gb = terabyte

// When we fix bugs in programs/computers its called debugging
// We call it debugging because the first computers used hot vaccum tubes that attracted moths
// Grace Hopper taped the first dead moth ever found in a pc in her journal and she was one of the most famous computer engineers
// The first "real" bug found in a computer lead to the term debugging or removing the bugs.


//INTEGERS 
//Integers and floats both have numbers so what is the difference?
//Floats or floating points use what we call real numbers because they can be a whole number or a fraction of a number like 5 and 5.5
//Integers are only whole numbers like 1,2,3,4,5,6,...10, 20 and so on.

package main

import "fmt"
	
func main() {
	x := 42
	y := 42.34534
	fmt.Println(x, y)//this will just print the value of x and y
	fmt.Printf("%T\n", x, y)//this will print out telling us x is an integer while y is a float64
}

//if we wanted to declare their types we could write

package main

import "fmt"

var x int
var y float64//writing var y float64 this way declares its type so later we can use = instead of := but it makes code longer
	
func main() {
	x = 42
	y = 42.34534
	fmt.Println(x, y)
	fmt.Printf("%T\n", x, y)
}

//if you write var x int then you declare it as an int which means you cant assign it a decimal number later like x = 2.3857394
//if you try to assign a decimal number to an int it will not run. the code will fail.

//In go we have integers and uintegers.
// int8's go from a neg to a pos like -128 to 127
// uint8's go from 0 to 225
// any number used in go is called a constant
// the u in uint stands for unsigned
// if it is signed it would go from neg to pos
// if its unsigned it always stays positive
// 8 bits allow you to store 256 things like 0 - 255 and 0 counts as the 256th.

package main

import "fmt"

var x int8 = -128

func main() {
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

//Output is -128 and int8 
//if we put -129 then the code wouldn't run because -129 doesn't exist withing the 256 numbers stored by int8
//-128 through 127 are supported by int8

// There are also a set of predeclared numeric types with implementation-specific sizes:
//uint is either 32 or 64 bits
//int is the same size as uint(either 32 or 64 bits)
//uintptr is an unsigned integer large enough to store the uninterpreted bits of a pointer value.

//All numeric types are distinct except
//byte which is an alias for uint8 (this is because a byte is 8 bits)
//rune which is an alias for int32 (rune is another way to say character which needs to be 32 possible bits)

package main

import (//I don't usually use parens but parens are needed if you are importing multiple packages
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
//the output for this is windows and amd64
//GOOS grabbed the OS im running which is windows
//GOARCH grabbed the architecture of my pc which is amd64 or 64bits

//STRING TYPES
//this code will print a basic string and will also tell us the type of var which is string
package main

import (
	"fmt"
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
}

//We can convert strings to bytes by doing this

package main

import (
	"fmt"
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	fmt.Println([]byte(s))
}

//fmt.Println([]byte(s)) lets us convert type string to type of byte but we can also write it like

package main

import (
	"fmt"
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)
}
//OUTPUT
/*
Hello, World
string
[72 101 108 108 111 44 32 87 111 114 108 100]
[]uint8*/

//So its printing our string
//its printing our type which is string
//Its converting our string to byte information and storing it in a slice or []
//its printing our type from the conversion which is now []uint8
//[72 101 108 108 111 44 32 87 111 114 108 100] is our byte information which can be converted to a string again by using an ascii table to decode it
//[72 101 108 108 111 44 32 87 111 114 108 100]
// H | e | l | l | o | , | | W | o | r | l | d |
// Thats the ascii conversion
//in utf-8 a code point is 1 to 4 code bytes and each point is represented by character. like capital H is represented by the two bytes 72

package main

import (
	"fmt"
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {//for i(index) equals 0(declaring the value of i as 0). as long as index is less than the length of s ad 1 to the index until its length matches s
		fmt.Printf("%#U", s[i])//this is going to print out the UTF-8 conversion and also give us a letter for each code point. 
	}
	
	fmt.Println("")//this will make the below code appear on a new line when printed to the output terminal

	for i, v := range s {//for i(index), v(value) equals the range of s which adds 0-11 since there are 12 characters and will put them in a line going down with their assigned byte data
		fmt.Println(i, v)

	}
}
//"%#U" should return the UTF-8 code point
//below is the output for for i := 0; i < len(s); i++ { fmt.Printf("%#U", s[i]) }
//Output U+0048 'H'U+0065 'e'U+006C 'l'U+006C 'l'U+006F 'o'U+002C ','U+0020 ' 'U+0057 'W'U+006F 'o'U+0072 'r'U+006C 'l'U+0064 'd'
//below is the output for for i, v := range s { fmt.Println(i, v) }
//0 72
//1 101
//2 108
//3 108
//4 111
//5 44
//6 32
//7 87
//8 111
//9 114
//10 108
//11 100

//If we wanted to print the bottom part with a decimal we could use
package main

import (
	"fmt"
)

func main() {
	s := "Hello, World"
	fmt.Println(s)
	fmt.Printf("%T\n", s)
	
	bs := []byte(s)
	fmt.Println(bs)
	fmt.Printf("%T\n", bs)

	for i := 0; i < len(s); i++ {//for i(index) equals 0(declaring the value of i as 0). as long as index is less than the length of s ad 1 to the index until its length matches s
		fmt.Printf("%#U", s[i])//this is going to print out the UTF-8 conversion and also give us a letter for each code point. 
	}
	
	fmt.Println("")//this will make the below code appear on a new line when printed to the output terminal

	for i, v := range s {//for i(index), v(value) equals the range of s which adds 0-11 since there are 12 characters and will put them in a line going down with their assigned byte data
		fmt.Printf("at the index position %d we have hex %#x\n", i, v)

	}
}
//Changing from fmt.Println(i, v) to fmt.Printf("at the index position %d we have hex %#x\n", i, v) gives this output
//Output
//at the index position 0 we have hex 0x48
//at the index position 1 we have hex 0x65
//at the index position 2 we have hex 0x6c
//at the index position 3 we have hex 0x6c
//at the index position 4 we have hex 0x6f
//at the index position 5 we have hex 0x2c
//at the index position 6 we have hex 0x20
//at the index position 7 we have hex 0x57
//at the index position 8 we have hex 0x6f
//at the index position 9 we have hex 0x72
//at the index position 10 we have hex 0x6c
//at the index position 11 we have hex 0x64


//if we look at the previous UTF-8 output we see a code point that looks like like U+0048 'H'
//these UTF-8 code points are called runes. so U+0048 'H' is a rune
//Hello, World written as runes would be (U+0048 'H'U+0065 'e'U+006C 'l'U+006C 'l'U+006F 'o'U+002C ','U+0020 ' 'U+0057 'W'U+006F 'o'U+0072 'r'U+006C 'l'U+0064 'd')



//NUMERAL SYSTEMS
//numeral means number 
//as humans we use the base-10 decimal 
//for example when we are counting number places(like commas in a large dollar amount) we go 

//Decimal example table below aka base-10
//one | ten | hundred | thousand | ten thousand | hundred thousand | million | ten million | place
//0-9 |10-99| 100-999 | 1000-9999| 10000-99999  | 100000-999999    | 1m-9m   | 10m-99m     | supported values
//10^0|10^1 |  10^2   |   10^3   |     10^4     |       10^5       |   10^6  |     10^7    | power
//10^0 = 1| 10^1=10| 10^2=hundred and so on
//If i were to have $12 in base-10 id have a $10 bill and a $2 bill or a 1 in the ten spot and a 2 in the one spot
//if i had 42420 it would be written on our chart like
//one | ten | hundred | thousand | ten thousand |
// 0  |  20 |   400   |   2000   |     40000    |
//this is because you have 0 ones, 2 tens, 4 hundreds, 2 thousands, and 4 ten thousands.

//Binary example table aka base-2
//ones | twos | fours | eights | sixteens | thirty twos | sixty fours | one twenty eights | place
//2^0  |  2^1 |  2^2  |   2^3  |   2^4    |     2^5     |     2^6     |        2^7        | power
// 0   |   0  |   0   |    0   |    0     |      0      |      0      |         0         | value
//so if i wanted to write the number 1 in binary id put a 1 in the ones spot
//if i wanted the number 2 id put the number 1 in the two spot
//if i wanted the number 3 id put the number 1 in both the one and two spot
//if i wanted the number 4 id put the number 1 in the four spot
//so 4 in binary is 100
//42 in binary is 010101

//Hexadecimal example table below aka base-16
//ones  |  16s  |  256s  | 4,096s  | 65,536s |      |      |      |     |   |   |   |   |   |   |place
//16^0  |  16^1 |  16^2  |   16^3  |  16^4   | 16^5 | 16^6 | 16^7 |     |   |   |   |   |   |   |power
//  a   |   b   |   c    |    d    |    e    |   f  |      |      |     |   |   |   |   |   |   |alphabet value
//  10  |   11  |   12   |    13   |    14   |  15  |      |      |     |   |   |   |   |   |   |numeric value for alphabet
//  0   |   1   |   2    |    3    |    4    |  5   |   6  |   7  |  8  | 9 | a | b | c | d | e | f |0 through 15 in hexadecimal
//  0   |   0   |   0    |    0    |    0    |  0   |   0  |   0  |	 0  | 0 | 0 | 0 | 0 | 0 | 0 | 0 |value
//so if i want to write 16 in base-16 id do 10
//if i wanted to do 13 id put 0 in the 16s spot and d in the ones spot
//if i want the number 8 id put 8 in the one spot and its just written as 8
//if i wanted 10 id put an a in the ones spot
//if i wanted 911 id do 38f
//to get 911 you can divide 911/256=3.5 so we will use 3 256's, 3*256=768, 911-768=143, 143/16=8.9 so we'll use 8 16s,8*16=128, 143-128=15, so we will use f in the one spot since its value is 15
//
//here is a good way to get byte data from a specific number or letter
package main
import "fmt"

func main() {
	s := "H"//number or letter you want the byte data of
	fmt.Println(s)//print s which == H

	bs := []byte(s)//byte s is converted into a byteString
	fmt.Println(bs)//print the byteString which will output 72
	fmt.Println(bs[0])//we can use this command to access our byteString on line 0 of the slice to print it without the []
}

//we can also assign the bs[0] to a letter like this

package main
import "fmt"

func main() {
	s := "H"
	fmt.Println(s)

	bs := []byte(s)
	fmt.Println(bs)
	
	n := bs[0]
	fmt.Println(n)
	fmt.Printf("%T\n", n)
}

//then we can also convert it to hexadecimal 

package main
import "fmt"

func main() {
	s := "H"
	fmt.Println(s)//H

	bs := []byte(s)
	fmt.Println(bs)//[72]
	
	n := bs[0]
	fmt.Println(n)//72
	fmt.Printf("%T\n", n)//uint8
	fmt.Printf("%b\n", n)//1001000
	fmt.Printf("%#x\n", n)//0x48
}


//Constants 
//constants allow ease of coding. 
//there are typed and untyped constants
//typed constants are constants given a type like const a int = 43
//untyped constants are written like const = 43
package main
import "fmt"

const a = 42
const b = 42.78
const c = "James Bond"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

//to make this code shorter using const we can type const once like this

package main
import "fmt"

const (
	a = 42
	b = 42.78
	c = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

//we can also give vars in a const a type so that the const doesn't get to be flexable and choose for itself

package main
import "fmt"

const (
	a int = 42
	b float64 = 42.78
	c string = "James Bond"
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

//IOTA
//Iotas are used to construct a set of related constants 
//Iota represents successive untyped integer constants
//Iota gets reset to 0 whenever the reserve word cont appears in the source and increments after each ConstSpec

package main
import "fmt"

const (
	a = iota
	b = iota
	c = iota
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}

//we can make this even easier by writing 
package main
import "fmt"

const (
	a = iota
	b
	c
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
}
//we didn't assign values to b and c but we still got values outputted for them because iota incrememnted the code

package main
import "fmt"

const (
	a = iota
	b
	c
)

const (
	d = iota
	e
	f	
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)

	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}

//this lets you run two separate constants with the same 0,1,2 output but if we write it like this

package main
import "fmt"

const (
	a = iota
	b
	c
	d = iota
	e
	f	
)

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b)
	fmt.Printf("%T\n", c)
	fmt.Printf("%T\n", d)
	fmt.Printf("%T\n", e)
	fmt.Printf("%T\n", f)
}
//the output will be 0,1,2,3,4,5
//Iota is good if you need code to autoincrement by one
//if you want to specify a value and have it autoincrement up then us + iota like this
package main
import "fmt"

const (

	a = 2021 + iota 
	b 
	c
	d
	e
)

func main() {

	fmt.Println(a, b, c, d, e)
	
}


//BITSHIFTING 
//<< >> are used in Go to shift bits

package main
import "fmt"

func main() {
	x := 4
	fmt.Printf("%d\t\t%b\n", x, x)
	
	y := x << 1
	fmt.Printf("%d\t\t%b", y, y)
}

//Output
//4		100
//8		1000
//as we can see the bit 1 was shifted with << from 100 to 1000 adding an extra 0 into the binary outcome 

//using iota in a creative way

package main
import "fmt"

func main() {
	kb := 1024
	mb := 1024 * kb
	gb := 1024 * mb

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

//OUTPUT
//1024				10000000000		<--kb
//1048576			100000000000000000000		<--mb
//1073741824		1000000000000000000000000000000		<--gb
//this is telling us that our binary digit shifted by ten spaces each time

//here is how we add cont and iota to the code to make it easier on ourselves
package main
import "fmt"

const (
	_ = iota//this is the 0 iota and we use _ to throw the value returned away since we don't want it
	kb = 1 << (iota * 10)//this is the 1 iota and its shifting 1 over 10 times
	mb = 1 << (iota * 10)//this is the 2 iota and its shifting 1 over 20 times
	gb = 1 << (iota * 10)//this is the 3 iota and its shifting 1 over 30 times
)

func main() {

	fmt.Printf("%d\t\t\t%b\n", kb, kb)
	fmt.Printf("%d\t\t\t%b\n", mb, mb)
	fmt.Printf("%d\t\t%b\n", gb, gb)
}

//Ninja level 2
//Hands on exercise #1
//Write a program that prints a number in decimal, binary, and hex

package main
import "fmt"

func main() {
	x := 36
	fmt.Printf("%#x\t%b\t%d\n", x, x, x)
}

//Hands on exercise #2
/*Using the following operators, write expressions and assign their values to variables:
==
<=
>=
!=
<
>
Now print each of the variables.
*/


package main
import "fmt"


func main() {
	a := 42
	b := 25
	c := 25
	d := 0
	e := 100
	f := 1

	if a == 42 {
		fmt.Println("Congrats champ, ")
	}
	if b <= a {
		fmt.Println("You are ")
	}
	if c >= b {
		fmt.Println("on your ")
	}
	if d != b {
		fmt.Println("way to ")
	}
	if e > a {
		fmt.Println("becoming a true ")
	}

	if f < e {
		println("coding ninja! ")
	}
}

//note if then statements in go don't need you to write out the word then. 
//{} are your then statement

//Hands on exercise #3
//Create TYPED and UNTYPED constants. Print the values of the constants.

package main
import "fmt"

const (
	a = 1
	b int = 2
	c = 3
)

const d = "Cereal"
const e = "PB & J"
const f = "are both breakfast foods"

func main() {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
}

//Hands on exercise #4
/*Write a program that 
assigns an int to a variable
prints that int in decimal, binary, and hex
shifts the bits of that int over 1 position to the left, and assigns that to a variable
prints that variable in decimal, binary, and hex*/


package main
import "fmt"

const (
	a int = 100	
)

func main() {
	fmt.Printf("%d\t%b\t%#x\n", a, a, a)
	y := a << 1
	fmt.Printf("%d\t%b\t%#x", y, y, y)
}


//Hands on exercise #5
//Create a variable of type string using a raw string literal. Print it.

package main
import "fmt"

func main() {
	a := `Hello Tod McLeod, 
	I am taking your course on golang`
	fmt.Println(a)
}

//Hands on exercise #6
//Using iota, create 4 constants for the NEXT 4 years. Print the constant values.

package main
import "fmt"

const (

	a = 2021 + iota 
	b 
	c
	d
	e
)

func main() {

	fmt.Println(b, c, d, e)
	
}

//Hands on exercise #7
//was a test and i scored 19 out of 21 and got a 90%
//questions i got wrong 

//rune is an alias for int32 which is true
//a string is a sequence of bytes that represent unicode code points called runes which is true

//LOOPS
//for loop

package main
import "fmt"

func main() {

	for init; condition; post {
	}
	fmt.Println()
	
}

//the way a for statement works is (for init; condition; post {})
//so for init apply this condition and post 
//example

package main
import "fmt"

func main() {

	for i := 0; i <= 100; i++ {
	fmt.Println(i)
}
	
}

//in this example we see that is i := 0 - init
//then my condition is while i is less than or equal to 100 then - condition
//add 1 while i is less than or equal to 100 - post
//so our for loop contains init, condition, post and {} curly braces with code in them

//LOOP - LOOP NESTING
//this will do a loop inside a loop
//inside loop and outside loop which will run the outside loop will run once and trigger the inside loop to run however many times and then the outside loop will loop again
//this can be done as much as you want
//inside and outside loops can both run 10 times for example

//nesting loops means you put a loop in a loop

package main
import "fmt"

func main() {
	//for init; condition; post {}
	for i := 0; i <= 10; i++ {//this is the outer loop. When this runs once it passes to the inside loop
		for j := 0; j < 3; j++ {//this is the inner . When it runs 3 times it passes back to the outer loop and repeats until the outer loop has run 10 times
			fmt.Printf("The outer loop: %d\t The inner loop: %d\n", i, j)//separate both loops into columns with their numbers printed next to them in decimal
		}
	}
}
//this code shows us the outer loop running 10 times while the inner loop runs 3
//the way it does this is running outer first and then inner 3 times then passing back to the outer to repeat the proccess until the condition is met
//a better visual would be if i write the code this way and the output will show you the way the loop is working

package main
import "fmt"

func main() {
	//for init; condition; post {}
	for i := 0; i <= 10; i++ {
		fmt.Printf("The outer loop: %d\n", i)
		for j := 0; j < 3; j++ {
			fmt.Printf("\t\t The inner loop: %d\n", j)
		}
	}
}
//output
//The outer loop: 0
//		The inner loop: 0
//		The inner loop: 1
//		The inner loop: 2
//The outer loop: 1
//		The inner loop: 0
//		The inner loop: 1
//		The inner loop: 2
//and so on until the outer loop hits 10

//you can also write it like this

package main
import"fmt"

func main() {
	x := 1//init
	for x < 10 {//for loop statement and condition
		fmt.Println(x)//print
		x++//adds one if the loop statement gets printed and x isn't greater than 10
	}
	fmt.Println("done.")//prints "done." when the x value hits 10
}
//note this will not print the number 10 because when its at ten its no longer less than 10 so there isn't a reason to print.
//if we change < to <= it would print 10

//for statements with a ForClause is also controlled by its condition, but additionally it may specify an int and a post statement
//such as an assignment, increment or decrement statement
//The init statement may be a short variable declaration but the post statement must not be
//Variables declared by the init statement are re-used in each iteration
//for = [ InitStmt ] ";" [ Condition ] ";" [ PostStmt ] .
//InitStmt = SimpleStmt .
//PostStmt = SimpleStmt .
//Example

//	for i := 0; i < 10;  i++ {
//		f(i)

//If non-empty the init statement is excecuted once before the condition for the first iteration
//this means if for i := 0 is not empty it will be executed before i < 10 is even evaluated 

//the post statement is executed after each execution of the block
//this means the post or i++ in this case gets executed every time i := 0; i <10 gets executed

//Using If in a loop

package main
import "fmt"

func main() {
	x := 1
	for {//eternally loop until i tell you to stop
		if x > 9 {//if x is greater than 9 (so when it hits 10) break the loop
			break//breaks the loop
		}
		fmt.Println(x)//otherwise print the value of x (until its value is higher than 9)
		x++//increment x value +1 until it reaches 10
	}
	fmt.Println("done.")//when x == 10 print done
}

//for loops in go are similar to C loops but they are not the same
//Go unifies for and while and there is no do-while
//there are only three forms of for statements

//Like a C for loop
//for init; condition; post {}

//Like a C while loop
//for condition {}

//Like a C for(;;) loop
//for {}

//short declaration := makes it easy to declare the index value right in the loop
//sum := 0
//for i := 0; i < 10; i++ {
//		sum += i
//}

//LOOP Breaks & Continues

package main
import "fmt"

func main() {
	x := 41 / 40
	fmt.Println(x)
}

//output is one because we are dividing 41 by 40

package main
import "fmt"

func main() {
	x := 43 % 40
	fmt.Println(x)
}

//if you change / division to % it will tell us the remainder of numbers left after removing 40 from 43

//if we use division and remainder we can see the output
package main
import "fmt"

func main() {
	x := 83 / 40
	y := 83 % 40
	fmt.Println(x)
	fmt.Println(y)
}
//Output is 2 and 3. 83 / 40 = 2 and the remainder of 83 % 40 is 3
//% can be used to get the remainder on dividing two integer numbers.
//you can also find remainders using math.Remainder like this but you need to import the math package
//Example

package main
import (
"fmt"
"math"
)

func main() {
	res := math.Remainder(83, 40)
	fmt.Println(res)
}

//this is also called modulus
//if we look at the way this works and we use some logic we can print 0-100 but print only even or odd numbers rather than all the numbers between 0 and 100

package main
import "fmt"

func main() {
	x := 1
	for {
		x++
		if x > 100 {
			break
		}
		if x%2 != 0{
			continue	
		}
		fmt.Println(x)
	}
}

//

// this code lets us print particular things like ascii
package main
import "fmt"

func main() {
	for i := 33; i <= 122; i++ {//here we are telling the code to print 33 through 122 of the ascii characters
			fmt.Printf("%v\t%#U\t%#x\n", i, i, i)//%v returns the decimal number for the ascii and %#U returns the code point(rune) and the ascii character
	}//%#x will return the hexadecimal value
}


















































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































































