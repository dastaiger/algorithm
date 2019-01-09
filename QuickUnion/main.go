//This Programm is an implementation of the Quick-Union algorithm as presented in https://www.coursera.org/learn/algorithms-part1 Week1. (PDF Page 26)
//

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	//define the number of Elements in the DataStruct
	fmt.Println("Enter the number of Elements:")
	scanner.Scan()
	numOfElements, err := strconv.Atoi(scanner.Text())
	if err != nil {
		panic(err)
	}
	ufDataStruct := make([]int, numOfElements)
	//initialize the ufDataStruct
	for i := 0; i < numOfElements; i++ {
		ufDataStruct[i] = i
	}

	//ask user for input for 6 connections
	fmt.Println("Please enter the connection you want. Like 2,3 to connect node 2 and 3")
	for i := 0; i < 6; i++ {
		scanner.Scan()
		pid, qid := convertStringToPQinput(scanner.Text())
		union(pid, qid, &ufDataStruct)

	}
	fmt.Println("This is the final Data Structure:")
	fmt.Println(ufDataStruct)

	//ask user for input to check if nodes are connected
	//endless loop
	fmt.Println("You can now ask if nodes are connected. Like 2,3 to check if node 2 and 3 are connected. Cancel with CTRL+C")
	for scanner.Scan() {
		pid, qid := convertStringToPQinput(scanner.Text())
		fmt.Println(connected(pid, qid, &ufDataStruct))

	}

}

//union will take in two nodes and the *datestrcutre and connect these two nodes with each other
//It does so by connecting the roots with each other
func union(p int, q int, uf *[]int) {
	i, j := root(p, uf), root(q, uf)
	(*uf)[i] = j
}

//connected will take in two nodes and the *datastrucutre and check if these two nodes are connected to each other.
//Returnes true or false
func connected(p int, q int, uf *[]int) bool {
	return root(p, uf) == root(q, uf)
}

//root is a helper function that will find the root of provided node (int) in the provided *datastructure.
//returns the root
func root(i int, uf *[]int) int {
	for i != (*uf)[i] {
		i = (*uf)[i]
	}
	return i
}

//convertStringToPQinput is a helper function which takes in a string in the format "x,y", i.E "int,int", "2,3"
//it will seperate these two ints inside the string with the seperator ","
//returns the two ints
//
//convertStringToPQinput is intendet to be used with scanner.Text() i.E. : p,q := convertStringToPQinput(scanner.Text())
func convertStringToPQinput(str string) (int, int) {
	connection := strings.Split(str, ",")
	p, err := strconv.Atoi(connection[0])
	if err != nil {
		fmt.Println("wrong entry!")
	}
	q, errr := strconv.Atoi(connection[1])
	if errr != nil {
		fmt.Println("wrong entry!")
	}

	return p, q
}
