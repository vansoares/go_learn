package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
    value := fmt.Sprintf("%d", x)

    queue1 := make([]string, 0)
    queue2 := make([]string, 0)

    j:= len(value)-1
    for i:=0; i<len(value) ; i++ {
        queue1 = append(queue1, string(value[i]))
        queue2 = append(queue2, string(value[j]))
        j--
    }

    for k:=0; k<len(queue1); k++ {
        fmt.Println( queue1[k], queue2[k])
        if queue1[k] != queue2[k]{
            return false
        }
    }
    return true

}

func main() {
    fmt.Println(isPalindrome(-10))
}