package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func contains(s []int32, e int32) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {

	bribesQuantity := 0
	bribedValues := make([]int32, len(q))

	for index, value := range q {
		hasBribedMoreThanTwoTimes := value > int32(index+3)
		if hasBribedMoreThanTwoTimes {
			fmt.Println("Too chaotic")
			return
		}

		hasBribed := value != int32(index+1)
		wasBribed := contains(bribedValues, int32(index+1))

		if hasBribed && wasBribed {
			fmt.Println(index + 1)
			bribedValues = append(bribedValues, int32(index+1))
			bribesQuantity++
		}

	}

	fmt.Println(bribesQuantity)

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		var q []int32

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			qItem := int32(qItemTemp)
			q = append(q, qItem)
		}

		minimumBribes(q)
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
