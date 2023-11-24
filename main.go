package main

import(
	"fmt"
	"math/rand"
	"time"
	"strings"
	
	pal "github.com/abusomani/go-palette/palette"
)

// p := [][]int {
	// 	{1},
	// 	{1, 2},
	// 	{1, 2, 3},
	// 	{1, 2, 3, 4},
	// 	{1, 2, 3, 4, 5},
	// }

func main() {

	ers := []string{"good jump", "jump out of line", "space occupied", "need to jump a piece"}

	p := [][]string {
		{"    *"},
		{"   *", " *"},
		{"  *", " *", " *"},
		{" *", " *", " *", " *"},
		{"*", " *", " *", " *", " *"},
	}

	var fromI, fromJ, toI, toJ int

	ps := setPuzzle(p)

	printPuzzle(ps)

	eternalLoop:
		for {
			if checkState(p) == 0 {
				fmt.Println("The Game Is Over")
				break eternalLoop
			}

			fmt.Println("Enter a move:")
			fmt.Scanf("%d%d%d%d", &fromI, &fromJ, &toI, &toJ)
			if fromI == -1 {
				break eternalLoop
			}
			ps2, msg := jumpPiece(ps, fromI, fromJ, toI, toJ)
			if msg != 0 {
				fmt.Println(ers[msg])
			}

			printPuzzle(ps2)
		}

		_ = checkState(p)
}

func printPuzzle(p[][]string) {
	foreground := pal.Color(0)
	pl := pal.New()

	for i := 0; i < len(p); i++ {
		switch i {
		case 0:
			foreground = pal.Color(214)
		case 1:
			foreground = pal.Color(40)
		case 2:
			foreground = pal.Color(202)
		case 3:
			foreground = pal.Color(196)
		case 4:
			foreground = pal.Color(118)
		}
		for j := 0; j < len(p[i]); j++ {
			if strings.Contains(p[i][j], "^") {
				pl.SetOptions(pal.WithForeground(foreground))
				pl.Print(p[i][j])
			} else {
				pl.SetOptions(pal.WithForeground(50))
				pl.Print(p[i][j])
			}
		}
		fmt.Print("\n")
	}
}

func setPuzzle(p[][]string) [][]string{

	x1 := rand.NewSource(time.Now().UnixNano())
	y1 := rand.New(x1)
	num1 := y1.Intn(5)

	var num2 int = 0

	if num1 > 0 {
		switch num1 {
		case 1:
			num2 = y1.Intn(2)
		case 2:
			num2 = y1.Intn(3)
		case 3:
			num2 = y1.Intn(4)
		case 4:
			num2 = y1.Intn(5)
		}

		for i := 0; i < len(p); i++ {
			for j := 0; j < len(p[i]); j++ {
				if i == num1 && j == num2 {
					continue
				} else {
					p[i][j] = strings.Replace (p[i][j], "*", "^", 1)
				}
			}
		}
	} else {
		for i := 0; i < len(p); i++ {
			for j := 0; j < len(p[i]); j++ {
				if i == 0 {
					continue
				} else {
					p[i][j] = strings.Replace (p[i][j], "*", "^", 1)
				}
			}
		}
	}

	return p
}

func jumpPiece(p[][]string, fromI int, fromJ int, toI int, toJ int) ([][]string, int) {
	var jumpI int = fromI - toI
	var jumpJ int = fromJ - toJ
	var chkI, chkJ int

	if fromI < toI {
		chkI = toI - 1
		if jumpJ == 0 {
			// chkJ = jumpJ
			chkJ = fromJ
		} else {
			chkJ = fromJ + 1
		}
	} else if fromI > toI {
		chkI = fromI - 1
		if jumpJ == 0 {
			chkJ = jumpJ
		} else {
			chkJ = fromJ - 1
		}
	} else {
		chkI = fromI
		if fromJ < toJ {
			chkJ = fromJ + 1
		} else {
			chkJ = fromJ - 1
		}
	}

	if jumpI < 0 {
		jumpI = jumpI * -1
	}

	if jumpJ < 0 {
		jumpJ = jumpJ * -1
	}

	if jumpI == 0 || jumpI == 2 {
		if jumpJ == 0 || jumpJ == 2 {
			if strings.Contains(p[chkI][chkJ], "^") {
				if strings.Contains(p[toI][toJ], "*") {
					//fmt.Println("chkI and J: ", chkI, chkJ)
					p[fromI][fromJ] = strings.Replace (p[fromI][fromJ], "^", "*", 1)
					p[chkI][chkJ] = strings.Replace (p[chkI][chkJ], "^", "*", 1)
					p[toI][toJ] = strings.Replace (p[toI][toJ], "*", "^", 1)
					return p, 0
				} else {
					return p, 2
				}
			} else {
				fmt.Println("jaguar: ", p[chkI][chkJ])
				fmt.Println("chkI et chkJ: ", chkI, chkJ)
				return p, 3
			}
		} else {
			return p, 1
		}
	} else {
		return p, 1
	}
}

func checkState(p[][]string) int {
	state := 0

	//     0      0 [0]
	//    0 1     1 [0 1]
	//   0 1 2    2 [0 1 2]
	//  0 1 2 3   3 [0 1 2 3]
	// 0 1 2 3 4  4 [0 1 2 3 4]
	//strings.Contains(p[i][j], "*") && strings.Contains(p[i+1][j+1], "^") && strings.Contains(p[i+2][j+2], "^"

	if (strings.Contains(p[0][0], "*") &&  strings.Contains(p[1][0], "^") && strings.Contains(p[2][0], "^")) || (strings.Contains(p[0][0], "^") &&  strings.Contains(p[1][0], "^") && strings.Contains(p[2][0], "*")) {
		//fmt.Println("hyena")
		state = 1
	}

	if (strings.Contains(p[0][0], "*") &&  strings.Contains(p[1][1], "^") && strings.Contains(p[2][2], "^")) || (strings.Contains(p[0][0], "^") &&  strings.Contains(p[1][1], "^") && strings.Contains(p[2][2], "*")) {
		//fmt.Println("lion")
		state = 1
	}

	
	gameLoop1:
	for i := 0; i < len(p)-2; i++ {
		for j := 0; j < len(p[i]); j++ {
			if (strings.Contains(p[i][j], "^") && strings.Contains(p[i+1][j], "^") && strings.Contains(p[i+2][j], "*")) || (strings.Contains(p[i][j], "^") && strings.Contains(p[i+1][j+1], "^") && strings.Contains(p[i+2][j+2], "*")) {
				state = 1
				//fmt.Println("rhino")
				break gameLoop1
			} 
		}
	}

	gameLoop2:
	for i := 0; i < len(p)-2; i++ {
		for j := 0; j < len(p[i]); j++ {
			if (strings.Contains(p[i][j], "*") && strings.Contains(p[i+1][j], "^") && strings.Contains(p[i+2][j], "^")) || (strings.Contains(p[i][j], "*") && strings.Contains(p[i+1][j+1], "^") && strings.Contains(p[i+2][j+2], "^")) {
				state = 1
				//fmt.Println("gazelle")
				break gameLoop2
			}
		}
	}

	gameLoop3:
	for i := 2; i < len(p); i++ {
		for j := 0; j < len(p[i])-2; j++ {
			if (strings.Contains(p[i][j], "^") && strings.Contains(p[i][j+1], "^") && strings.Contains(p[i][j+2], "*")) || (strings.Contains(p[i][j], "*") && strings.Contains(p[i][j+1], "^") && strings.Contains(p[i][j+2], "^")) {
				state = 1
				//fmt.Println("elephant")
				break gameLoop3
			} 
		}
	}

	return state
}