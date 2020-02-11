package main

import "fmt"

// 배열이나 슬라이스는 순서 인덱스라면
// 맵은 말 그대로 key-value 표현

func main() {
	var a map[string]int // string => int
	var b map[string]int = make(map[string]int)
	var c = make(map[string]int)
	d := make(map[string]int)
	_, _, _, _ = a, b, c, d

	map1 := map[string]int{
		"abc": 3,
		"def": 5,
	}
	fmt.Println(map1["sds"]) // 없는 key 접근하면 기본값
	for k, v := range map1 {
		fmt.Println(k, v)
	} // 그런데, c++ 에서는 없는 key 터치만 해도 데이터가 생겼는데 요건 다르네

	value, ok := map1["aaaa"] // 없는 키 접근시 기본값, bool
	fmt.Println(value, ok)

	fmt.Println(map1)
	delete(map1, "abc")
	delete(map1, "nothing_key")
	fmt.Println(map1)

	assocMap := map[string]map[string]int{
		"a": map[string]int{
			"a_a": 1,
			"a_b": 2,
			"a_c": 3,
		},
		"b": { // 생략 가능
			"b_a": 4,
			"b_b": 5,
		},
	}
	for _, v := range assocMap {
		for innerK, innerV := range v {
			fmt.Println(innerK, innerV)
		}
	}
}
