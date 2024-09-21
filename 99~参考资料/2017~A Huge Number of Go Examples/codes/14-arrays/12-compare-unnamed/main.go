
package main

import "fmt"

// STORY:
// You want to compare two bookcases,
// whether they're equal or not.

func main() {
	type (
		// integer int

		bookcase [5]int
		cabinet  [5]int
		//          ^- try changing this to: integer
		//             but first: uncomment the integer type above
	)

	blue := bookcase{6, 9, 3, 2, 1}
	red := cabinet{6, 9, 3, 2, 1}

	fmt.Print("Are they equal? ")

	if cabinet(blue) == red {
		fmt.Println("✅")
	} else {
		fmt.Println("❌")
	}

	fmt.Printf("blue: %#v\n", blue)
	fmt.Printf("red : %#v\n", bookcase(red))

	// ------------------------------------------------
	// The underlying type of an unnamed type is itself.
	//
	//   [5]integer's underlying type: [5]integer
	//   [5]int's underlying type    : [5]int
	//
	//   > [5]integer and [5]int are different types.
	//   > Their memory layout is not important.
	//   > Their types are not the same.

	// _ = [5]integer{} == [5]int{}

	// ------------------------------------------------
	// An unnamed and a named type can be compared,
	// if they've identical underlying types.
	//
	//   [5]integer's underlying type: [5]integer
	//   cabinet's underlying type   : [5]integer
	//
	// Note: Assuming the cabinet's type definition is like so:
	//       type cabinet [5]integer

	// _ = [5]integer{} == cabinet{}
}
