package main

import (
	"fmt"
	"log"
	"os"
	"bufio"
	"strconv"
)

func main() {
	//fd, err := os.Create("output.txt")
	fd, err := os.OpenFile("output.txt", os.O_RDWR | os.O_APPEND | os.O_CREATE, 0660)
	if err != nil {
		log.Fatal(err)
	}
	//writer := bufio.NewWriter(fd)
	//fmt.Printf("err = %v\n", err)

	var N int
	fmt.Fscan(os.Stdin, &N)
	reader := bufio.NewReader(os.Stdin)
	arr := make([][]byte, N)
	for i := 0; i < N; i++ {
		line, err := reader.ReadSlice('\n')
		if err != nil {
			log.Fatal(err)
		}
		arr[i] = make([]byte, (len(line) - 1))
		for j := 0; j < (len(line) - 1); j++ {
			arr[i][j] = line[j]
		}
	}

	fmt.Fscan(os.Stdin, &N)
	request := struct {
		n int
		l_r string
		w_a string
	}{}
	var answer string = ""
	for i := 0; i < N; i++ {
		fmt.Fscanf(os.Stdin, "%d %s %s", &request.n, &request.l_r, &request.w_a)
		for idx, line := range arr {
			if request.l_r == "left" {
				answer = checkL(&line, request.n, request.w_a, strconv.Itoa(idx + 1))
			} else {
				answer = checkR(&line, request.n, request.w_a, strconv.Itoa(idx + 1))
			}
			if answer != "" {
				//os.Stdout.WriteString(answer + "\n")
				fd.WriteString(answer + "\n")
				viewMap(&arr, fd)
				break
			} else if idx == len(arr) - 1 {
				fd.WriteString("Cannot fulfill passengers requirements\n")
				//fmt.Println("Cannot fulfill passengers requirements")
			}
		}
	}
	//fd.Close()
}

func viewMap(arr *[][]byte, fd *os.File) {
	//fd, _ := os.Open("output.txt")
	//writer := bufio.NewWriter(fd)
	//defer fd.Close()
	for i, line := range *arr {
		//writer := bufio.NewWriter(os.Stdout)
		//writer.WriteString(string(line))
		//io.WriteString(os.Stdout, string(line) + "\n")
		fd.WriteString(string(line) + "\n")
		//fmt.Println(string(line))
		for j, val := range line {
			if val == 'X' {
				(*arr)[i][j] = '#'
			}
		}
	}

}

func checkL(line *[]byte, number int, preference string, num_line string) string {
	if number == 3 {
		for i := 0; i < 3; i++ {
			if (*line)[i] == '#' {
				return ""
			}
		}
		(*line)[0] = 'X'
		(*line)[1] = 'X'
		(*line)[2] = 'X'
		return "Passengers can take seats: " + num_line + "A "  + num_line + "B "  + num_line + "C"
	} else if number == 2 {
		if (*line)[1] == '#' || (preference == "window" && (*line)[0] == '#') || (preference == "aisle" && (*line)[2] == '#') {
			return ""
		}
		(*line)[1] = 'X'
		if preference == "window" {
			(*line)[0] = 'X'
			return "Passengers can take seats: " + num_line + "A " + num_line + "B"
		} else {
			(*line)[2] = 'X'
			return "Passengers can take seats: " + num_line + "B " + num_line + "C"
		}
	} else {
		if preference == "window" && (*line)[0] != '#' {
			(*line)[0] = 'X'
			return "Passengers can take seats: " + num_line + "A"
		} else if preference == "aisle" && (*line)[2] != '#' {
			(*line)[2] = 'X'
			return "Passengers can take seats: " + num_line + "C"
		} else {
			return ""
		}
	}
}

func checkR(line *[]byte, number int, preference string, num_line string) string {
	if number == 3 {
		for i := 4; i < 7; i++ {
			if (*line)[i] == '#' {
				return ""
			}
		}
		(*line)[4] = 'X'
		(*line)[5] = 'X'
		(*line)[6] = 'X'
		return "Passengers can take seats: " + num_line + "D "  + num_line + "E "  + num_line + "F"
	} else if number == 2 {
		if (*line)[5] == '#' || (preference == "window" && (*line)[6] == '#') || (preference == "aisle" && (*line)[4] == '#') {
			return ""
		}
		(*line)[5] = 'X'
		if preference == "window" {
			(*line)[6] = 'X'
			return "Passengers can take seats: " + num_line + "E " + num_line + "F"
		} else {
			(*line)[4] = 'X'
			return "Passengers can take seats: " + num_line + "D " + num_line + "E"
		}
	} else {
		if preference == "window" && (*line)[6] != '#' {
			(*line)[6] = 'X'
			return "Passengers can take seats: " + num_line + "F"
		} else if preference == "aisle" && (*line)[4] != '#' {
			(*line)[4] = 'X'
			return "Passengers can take seats: " + num_line + "D"
		} else {
			return ""
		}
	}
}
