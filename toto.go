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

type Combination [DRAW_NUMBERS]int8

type Draw struct {
	number int8
	draws [2]Combination
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
		draw.number = int8(idraw_num)
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
	for i, snumber := range anumbers {
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
		c[i] = int8(inumber)
	}
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
	for _, draw := range d.draws {
		for i, number := range draw {
			fmt.Print(strconv.Itoa(int(number)))
			if i < 5 {
				fmt.Print(",")
			}
		}
		fmt.Print("\t\t")
	}
	fmt.Print("\n")
}
