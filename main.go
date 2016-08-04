package main

import (
	"fmt"
	"os/exec"
	"strconv"
	"time"
)

func Run(n int) {
	var ok = make(chan int)
	var count int

	start := time.Now()
	for i := 0; i < n; i++ {
		go func(i int, ok chan int) {
			cmd := exec.Command("hypercli", "run", "-d", "--name", "name"+strconv.Itoa(i), "busybox")
			//cmd := exec.Command("hypercli", "rm", "-f", "name"+strconv.Itoa(i))
			out, err := cmd.CombinedOutput()
			if err != nil {
				ok <- 2
				fmt.Println("container name is ", i)
				fmt.Println("hyper out is:", string(out))
				fmt.Println("err is ", err.Error())
			} else {
				ok <- 1
			}
		}(i, ok)
	}

	for i := 0; i < n; i++ {
		if <-ok == 1 {
			count++
			fmt.Println(count)
		}
	}
	last := time.Now().Sub(start).String()
	fmt.Println("start or rm ", count, "containers!")
	fmt.Println("last --> ", last)
}

func main() {
	utils.Run(100)
	fmt.Println("pressure test finish!")
}
