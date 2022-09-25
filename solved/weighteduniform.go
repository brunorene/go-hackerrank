package main

import (
	"fmt"
)

func fragments(s string) (result []string) {
	result = append(result, "")

	for idx, c := range s {
		result[len(result)-1] += string(c)
		if idx < len(s)-1 && s[idx] != s[idx+1] {
			result = append(result, "")
		}
	}

	return
}

func weights(frags []string) (result map[int32]struct{}) {
	result = make(map[int32]struct{})

	for _, s := range frags {
		w := int32(s[0] - 'a' + 1)

		for idx := range s {
			result[w*int32(idx+1)] = struct{}{}
		}
	}

	return
}

func weightedUniformStrings(s string, queries []int32) (result []string) {
	ws := weights(fragments(s))

	for _, q := range queries {
		if _, exist := ws[q]; exist {
			result = append(result, "Yes")

			continue
		}

		result = append(result, "No")
	}

	return
}

func mainUniform() {
	fmt.Println(weightedUniformStrings("abccddde", []int32{1, 3, 12, 5, 9, 10}))
}
