package concept

import "fmt"

func anime() {
	var animeName = "One Piece" 
	var epNumber = 1008
	var protaganist = "Monkey D. Luffy"

	fmt.Printf("animeName is %v\n epNumber is %v\n protaganist is %v\n", animeName, epNumber, protaganist)
}
func Anime() {
	var animeName = "Attack On Titan"
	var epNumber = 80
	var protaganist = "Eren Jaegar"

	fmt.Printf("\nanimeName is %v\n epNumber is %v\n protaganist is %v\n", animeName, epNumber, protaganist)

	anime()
}
