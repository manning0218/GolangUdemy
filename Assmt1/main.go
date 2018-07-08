package main

import "fmt"

func main() {
    nums := createInts(10)

    if len(nums) > 0 {
        for _, num := range nums {
            if isEven(num) {
                fmt.Println(num, " is even")
            } else {
                fmt.Println(num, " is odd")
            }
        }
    }
}

func createInts(n int) []int {
    if n > 0 {
        nums := make([]int, n + 1) // Need to account for 0 as well

        for i := 0; i < len(nums); i++ {
            nums[i] = i
        }

        return nums
    }

    return []int{}
}

func isEven(n int) bool {
    if n % 2 == 0 {
        return true
    }

    return false
}
