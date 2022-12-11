package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	// Read the input file.
	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("unable to open input file: %s", err)
	}
	defer f.Close()

	r := bufio.NewReader(f)

	// We create a variable to store the addition to X value for every cycle.
	var xAdditions = []int{0}
	for {
		l, err := r.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Fatalf("failed to read input file: %s", err)
		}

		if err == io.EOF {
			break
		}
		l = strings.TrimSpace(l)

		if strings.HasPrefix(l, "noop") {
			// This adds 1 cycle, but does nothing.
			xAdditions = append(xAdditions, 0)
		} else if strings.HasPrefix(l, "addx") {
			// This adds 2 cycle, and add the X value after 2 cycles.
			d := strings.Split(l, " ")
			c, _ := strconv.Atoi(d[1])
			xAdditions = append(xAdditions, 0)
			xAdditions = append(xAdditions, c)
		}
	}

	// Variables for Part 1
	x := 1
	strength := 0 // Initial variable to store the signal strength.

	// Variables for Part 2
	const w = 40    // CRT width
	spritePos := 1  // The current sprite position.
	spriteSize := 3 // This is the size of the sprite.
	row := 0        // The current row to draw on the CRT.

	for i := 1; i < len(xAdditions); i++ {
		// This part is for Part 1
		{
			x += xAdditions[i-1]

			switch i {
			case 20, 60, 100, 140, 180, 220:
				strength += i * x
			}
		}

		// This part is for Part 2
		{
			if i > 1 && (i-1)%w == 0 {
				row++
				fmt.Println()
			}

			if i >= (row*w)+spritePos && i < (row*w)+spritePos+spriteSize {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

			spritePos += xAdditions[i]
		}
	}

	fmt.Println()
	fmt.Println(strength) // Print the answer for Part 1
}
