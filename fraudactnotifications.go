package main

import (
	"fmt"
	"sort"
)

func median(data []int32, remove int32, add int32) (int32, []int32) {
	fmt.Printf("1 %v %d %d\n", data, remove, add)

	if add != remove {
		i := sort.Search(len(data), func(i int) bool { return data[i] >= remove })
		if i < len(data) && data[i] == remove {
			data = append(data[:i], data[i+1:]...)
		}

		j := sort.Search(len(data), func(i int) bool { return data[i] >= add })
		if j == len(data) {
			data = append(data, add)
		} else {
			mid := make([]int32, j)
			copy(mid, data)
			mid = append(mid[:j], add)
			data = append(mid, data[j:]...)
		}
	}

	fmt.Printf("2 %v\n\n", data)

	return data[len(data)/2], data
}

func activityNotifications(expenditure []int32, d int32) (notifCount int32) {
	size := int(d)

	window := expenditure[0:size]
	sort.Slice(window, func(i, j int) bool {
		return window[i] < window[j]
	})

	p50 := window[size/2]

	fmt.Printf("%d >= %d\n\n", expenditure[size], 2*p50)

	if expenditure[size] >= 2*p50 {
		notifCount++
	}

	for i := 1; i < len(expenditure)-size; i++ {
		p50, window = median(window, expenditure[i-1], expenditure[i+size-1])

		fmt.Printf("%d >= %d\n\n", expenditure[i+size], 2*p50)

		if expenditure[i+size] >= 2*p50 {
			notifCount++
		}
	}

	return
}

func main() {
	fmt.Println(activityNotifications([]int32{2, 3, 4, 2, 3, 6, 8, 4, 5}, 5))
}
