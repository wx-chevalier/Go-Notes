
package main

func main() {
	mobydick := book{
		title: "moby dick",
		price: 10,
	}

	minecraft := game{
		title: "minecraft",
		price: 20,
	}

	tetris := game{
		title: "tetris",
		price: 5,
	}

	// #3
	mobydick.print()  // sends `mobydick` value to `book.print`
	minecraft.print() // sends `minecraft` value to `game.print`
	tetris.print()    // sends `tetris` value to `game.print`

	// #2
	// mobydick.printBook()
	// minecraft.printGame()

	// #1
	// printBook(mobydick)
	// printGame(minecraft)
}
