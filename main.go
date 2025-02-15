package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func createArrayloopInput(n int) []int{
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Fields(sc.Text())

	a := make([]int, n)
	for i := range line {
		a[i], _ = strconv.Atoi(line[i])
	}

	return a
}

func createArraySplitSpace(n int) []int{
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()
	line := strings.Split(sc.Text()," ")
	var a []int
	for _, input := range line{
		p, _ := strconv.Atoi(input)
		a = append(a, p)
	}
	return a
}

func bubbleSort(a []int) []int{
	n := len(a)
	
	for i := 0; i < n; i++{
		for j := 0; j < n-1-i;j++{
			if a[j] > a[j+1]{
				tmp := a[j]
				a[j]=a[j+1]
				a[j+1]=tmp
			}
		}
	}
	return a
}

func create2dimArraySplitSpace(n int) [][]int{
	sc := bufio.NewScanner(os.Stdin)
	
	var a [][]int

	for i := 0; i<n; i++{
		sc.Scan()
		line := strings.Split(sc.Text()," ")
		for j := 0; j < len(line); j++{
			p, _ := strconv.Atoi(line[j])
			a[i] = append(a[i], p)
		}
	}
	return a
}
func coktailSort(a []int) []int{
	n := len(a)
	swap := false
	init := 0
	end := n-1

	for swap{
		swap = false
		for i:=init; i < end; i++{
			if a[i] > a[i+1]{
				tmp := a[i]
				a[i] = a[i+1]
				a[i+1]=tmp
				swap = true
			}
		}

		if !swap{
			break
		}
		swap = false
		end -= 1

		for i:=end; i > init+1; i++{
			if a[i] < a[i-1]{
				tmp := a[i]
				a[i]=a[i-1]
				a[i-1]=tmp
				swap = true
			}
		}
		init += 1
	}
	return a
}

func main(){
	var a int
	var b, c int
	var s string

	fmt.Scanf("%d",&a)
	fmt.Scanf("%d %d",&b,&c)
	fmt.Scanf("%s",&s)

	fmt.Println(a+b+c, s)
}