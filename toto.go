package main

import (
	"io"
	"os"
	"fmt"
	"bufio"
	"strconv"
	"strings"
	"sort"
)

type Combination []int

type Draw struct {
	number int
	draws [2]Combination
	max_hits [2]int
}

type Toto struct {
	numbers Combination
	draws []Draw
}

func NewToto() *Toto {
	return &Toto{}
}

func (t *Toto) ProcessNumbers(numbers string) {
	t.numbers = getCombination(numbers)
	t.checkNumbers();
}

func (t *Toto) checkNumbers() {
	for i, draw := range t.draws {
		for j, sdraw := range draw.draws {
			for _, num := range t.numbers {
				n := sort.SearchInts(sdraw, num) 
				if n < len(sdraw) && sdraw[n] == num {
					t.draws[i].max_hits[j]++
				}
			}
		}
	}
}

func (t *Toto) ProcessDraws (file_name string) {
	f, err := os.Open(file_name)
	if err != nil {
		fmt.Print(err)
		os.Exit(ERR_FILE_OPEN_ERROR)
	}

	defer f.Close()
	input := bufio.NewReader(f)
	for {
		line, err := input.ReadString('\n');
		if err == io.EOF {
			break
		}
		dash_ndx := strings.Index(line, "-")
		if dash_ndx < 0 {
			continue
		}
		var draw Draw
		sdraw_num := line[0:dash_ndx]
		idraw_num, _ := strconv.ParseInt(sdraw_num, 10, 8)
		draw.number = int(idraw_num)
		draw_separator :=  strings.Index(line, "\t\t")
		draw.draws[0] = getCombination(line[dash_ndx+1:draw_separator])
		draw.draws[1] = getCombination(line[draw_separator:])
		t.draws = append(t.draws, draw)
	}
}

func mytrim(s string) string {
	return strings.Trim(s, "\t \n")
}

func getCombination (numbers string) Combination {
	var c Combination
	// strip space
	numbers = strings.Replace(mytrim(numbers), " ", "", -1)
	anumbers := strings.Split(numbers, ",")
	sort.Sort(sort.StringSlice(anumbers))
	if len(anumbers) != 6 {
		fmt.Println("ERROR: Please provide six integers separated by (,)")
		os.Exit(ERR_INSUFFICIENT_NUMBERS)
	}
	// store last number to easy get duplicates
	last_number := int64(-1);
	for _, snumber := range anumbers {
		inumber, e := strconv.ParseInt(snumber, 10, 8)
		if e != nil {
			fmt.Printf("ERROR: Wrong number: %s\n", snumber)
			os.Exit(ERR_NOT_A_NUMBER)
		}
		if inumber < MIN_NUMBER || inumber > MAX_NUMBER {
			fmt.Printf("ERROR: Please provide number in the range [%s..%d]. Provided: %s", MIN_NUMBER, MAX_NUMBER, snumber)
			os.Exit(ERR_WRONG_NUMBER)
		}
		if (inumber == last_number) {
			fmt.Println("ERROR: Duplicate number: " + snumber)
			os.Exit(ERR_DUPLICATE_NUMBER)
		}
		last_number = inumber
		c = append(c, int(inumber))
	}
	sort.Sort(c)
	return c
}

func (t *Toto) Print () {
	for _, draw := range t.draws {
		draw.Print()
	}
}

func (d *Draw) Print () {
	fmt.Print(d.number)
	fmt.Print("\t")
	for i, draw := range d.draws {
		for j, number := range draw {
			fmt.Printf("%2s", strconv.Itoa(int(number)))
			if j < 5 {
				fmt.Print(", ")
			} else {
				fmt.Printf(" [%d]", d.max_hits[i])
			}
		}
		fmt.Print("\t\t")
	}
	fmt.Print("\n")
}

func (c Combination) Len() int           { return len(c) }
func (c Combination) Less(i, j int) bool { return c[i] < c[j] }
func (c Combination) Swap(i, j int)      { c[i], c[j] = c[j], c[i] }

func (t *Toto) FreqAnalysis() {
	var freq [MAX_NUMBER]int8
	var max int8 = 0;
	for _, draw := range t.draws {
		for _, sdraw := range draw.draws {
			for _, number := range sdraw {
				freq[number-1]++
				if freq[number-1] > max {
					max = freq[number-1]
				}
			}
		}
	}
	//Print
	for i, sum := range freq {
		fmt.Printf("%2d:\t%d |", i+1, sum)
		fmt.Print(strings.Repeat("=", int(68/int(max)*int(sum))));
		fmt.Println(">")
	}
}
