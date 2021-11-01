package hangman

import "fmt"

func DrawWelcome() {
	fmt.Println(`
	_   _   ___   _   _ _____ ___  ___  ___   _   _ 
	| | | | / _ \ | \ | |  __ \|  \/  | / _ \ | \ | |
	| |_| |/ /_\ \|  \| | |  \/| .  . |/ /_\ \|  \| |
	|  _  ||  _  || . " | | __ | |\/| ||  _  || . " |
	| | | || | | || |\  | |_\ \| |  | || | | || |\  |
	\_| |_/\_| |_/\_| \_/\____/\_|  |_/\_| |_/\_| \_/
													 
	`)
}

func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
	fmt.Println()
}

func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
		 _______
		 |/      |
		 |      (_)
		 |      \|/
		 |       |
		 |      / \
		 |
		_|___
		`
	case 1:
		draw = `
		 _______
		 |/      |
		 |      (_)
		 |      \|/
		 |       |
		 |      
		 |
		_|___
		`
	case 2:
		draw = `
		 _______
		 |/      |
		 |      (_)
		 |       |
		 |       |
		 |     
		 |
		_|___
		`
	case 3:
		draw = `
		 _______
		 |/      |
		 |      (_)
		 |      
		 |       
		 |     
		 |
		_|___
		`
	case 4:
		draw = `
		 _______
		 |/      |
		 |      
		 |      
		 |       
		 |      
		 |
		_|___
		`
	case 5:
		draw = `
		 _______
		 |/     
		 |      
		 |      
		 |      
		 |      
		 |
		_|___
		`
	case 6:
		draw = `
 		 |      
 		 |      
 		 |
		_|___
		`
	case 7:
		draw = `
	   _|___
		`
	case 8:
		draw = ``
	}

	fmt.Println(draw)

}

func drawState(g *Game, guess string) {
	fmt.Print("Guessed:")
	drawLetters(g.FoundLetters)
	fmt.Print("Used: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess!")
	case "alreadyGuessed":
		fmt.Printf("Letter, %s, was already used \n", guess)
	case "badGuess":
		fmt.Printf("Bad guess, %s is not in the word \n", guess)
	case "lost":
		fmt.Printf("You lost :(")
	case "won":
		fmt.Print("YOU WON The word was : ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string) {
	for _, c := range l {
		fmt.Printf("%v", c)
	}
	fmt.Println()
}
