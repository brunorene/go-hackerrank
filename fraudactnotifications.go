package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type valueWithFreq struct {
	value int32
	freq  int
}

type listOfValuesWithFreq []valueWithFreq

func (l *listOfValuesWithFreq) orderedInsert(v valueWithFreq) {

}

func get(vals []valueWithFreq, index int) int32 {
	var pivot int

	for _, vf := range vals {
		pivot += vf.freq

		if pivot > index {
			return vf.value
		}
	}

	return vals[len(vals)-1].value
}

func median(data []valueWithFreq, remove int32, add int32) (int32, []valueWithFreq) {
	// fmt.Printf("1 %v %d %d\n", data, remove, add)

	if add != remove {
		i := sort.Search(len(data), func(i int) bool { return data[i].value >= remove })
		if i < len(data) && data[i].value == remove {
			data[i].freq--
			if data[i].freq == 0 {
				ret := make([]valueWithFreq, 0)
				ret = append(ret, data[:i]...)
				data = append(ret, data[i+1:]...)
			}
			fmt.Println(data)
			os.Exit(1)
		}

		j := sort.Search(len(data), func(i int) bool { return data[i].value >= add })
		if j == len(data) {
			data = append(data, valueWithFreq{add, 1})
		} else {
			data[j].freq++
		}
	}

	// fmt.Printf("2 %v\n\n", data)

	cut := len(data) / 2
	if len(data)%2 == 0 && get(data, cut) != get(data, cut-1) {
		return (get(data, cut)+get(data, cut-1))/2 + +((get(data, cut) + get(data, cut-1)) % 2), data
	}

	return get(data, cut), data
}

func activityNotifications(expenditure []int32, d int32) (notifCount int32) {
	size := int(d)

	sort.Slice(expenditure, func(i, j int) bool {
		return expenditure[i] < expenditure[j]
	})

	p50 := expenditure[size/2]

	// fmt.Printf("%d >= %d\n\n", expenditure[size], 2*p50)

	if expenditure[size] >= 2*p50 {
		notifCount++
	}

	for i := 1; i < len(expenditure)-size; i++ {
		p50, window = median(window, expenditure[i-1], expenditure[i+size-1])

		if expenditure[i+size] >= 2*p50 {
			notifCount++
		}
	}

	return
}

func main() {
	f, err := os.Open("/home/brsantos/Projetos/hackerrank/fraudactnotifications_input.txt")
	if err != nil {
		fmt.Println("error opening file ", err)
		os.Exit(1)
	}
	defer f.Close()
	rd := bufio.NewScanner(f)

	rd.Split(bufio.ScanWords)

	rd.Scan()

	rd.Scan()

	window, err := strconv.Atoi(rd.Text())

	var nums []int32

	for rd.Scan() {
		n, err := strconv.Atoi(rd.Text())
		if err != nil {
			fmt.Println("error reading value ", err)
			os.Exit(1)
		}
		nums = append(nums, int32(n))
	}

	fmt.Println(activityNotifications(nums, int32(window)))
}
