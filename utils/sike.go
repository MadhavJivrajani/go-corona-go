package utils

import (
	"fmt"
	"math/rand"
	"time"
)

func sike1() {
	sike := ` 

needs lots of boops and bewwy wubsss
	\
	 \
	/^-----^\
	V  o o  V
	 |  Y  |
	  \ u /
	  / - \
	  |    \
	  |     \     )
	  || (___\====
`
	fmt.Println(sike)
	fmt.Println("\n\nThis amazing ASCII art was taken from https://www.asciiart.eu/animals/dogs")
}

func sike2() {
	msg := `
	FOR THE LOVE OF DOG! WEAR A MASK!
	 /		
	/
	`
	sike := " ____|\\" + "\n`-/    \\" + "\n (\\_/)  \\" + "\n /_  _   |" + " \n\\/_/||) /" + "\n    '---'"
	fmt.Println(msg)
	fmt.Println(sike)
	fmt.Println("\n\nThis amazing ASCII art was taken from https://ascii.co.uk/art/dog")
}

func sike3() {
	sike := `
Meet ken the chicken. ken crossed the road 'cause it ken.
	
    \\
    (o>
 \\_//)
  \_/_)
   _|_
   `
	fmt.Println(sike)
	fmt.Println("\n\nThis amazing ASCII art was taken from http://www.ascii-art.de/ascii/c/chicken.txt")
}

// PullSike chooses a different type of sike to pull everytime called.
func PullSike() error {
	rand.Seed(time.Now().UnixNano())
	switch choice := rand.Intn(3) + 1; choice {
	case 1:
		sike1()
	case 2:
		sike2()
	case 3:
		sike3()
	}
	return nil
}
