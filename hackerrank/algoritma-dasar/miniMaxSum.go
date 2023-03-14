package main

import "fmt"

func miniMaxSum(arr []int32) {
    var minimum, maximum, sum int64
    minimum = int64(arr[0])
    maximum = int64(arr[0])

    for _, val := range arr {
        sum += int64(val)
        if int64(val) < minimum {
            minimum = int64(val)
        }
        if int64(val) > maximum {
            maximum = int64(val)
        }
    }
    fmt.Println(sum-maximum, sum-minimum)
}

func main() {
	miniMaxSum([]int32{1, 2, 3, 4, 5})
}
