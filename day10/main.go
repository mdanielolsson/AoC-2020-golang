package main

import (
	"fmt"
	"sort"
	"strings"

	utils "github.com/mdanielolsson/advent-of-code/utils"
)

func main() {
	parsedInput := parseInput(input)
	fmt.Println("Part1:", part1(parsedInput))
	fmt.Println("Part2:", part2(parsedInput))
}

func part1(input []int) int {
	oneJDiff := 0
	threeJDiff := 0

	for i := len(input) - 2; i >= 0; i-- {
		switch input[i+1] - input[i] {
		case 3:
			threeJDiff++
		case 1:
			oneJDiff++
		}
	}
	return threeJDiff * oneJDiff
}

func part2(input []int) int {
	cache := map[int]int{0: 1}
	for i := 1; i < len(input); i++ {
		for j := 0; j < i; j++ {
			diff := input[i] - input[j]
			if 1 <= diff && 3 >= diff {
				cache[i] += cache[j]
			}
		}
	}
	return cache[len(input)-1]
}

func parseInput(input string) []int {
	var result []int
	for _, v := range strings.Split(input, "\n") {
		result = append(result, utils.StringToInt(v))
	}
	sort.Ints(result)
	result = append(result, result[len(result)-1]+3)
	result = append([]int{0}, result...)
	return result
}

var exampleInput = `16
10
15
5
1
11
7
19
6
12
4`

var input = `83
69
170
56
43
111
117
135
136
176
154
65
107
169
141
151
158
134
108
143
114
104
49
55
72
73
144
13
35
152
98
133
31
44
150
70
118
64
39
137
142
28
130
167
101
100
120
79
153
157
89
163
177
3
1
38
16
128
18
62
76
78
17
63
160
59
175
168
54
34
22
174
53
25
129
90
42
119
92
173
4
166
19
2
121
7
71
99
66
46
124
86
127
159
12
91
140
52
80
45
33
9
8
77
147
32
95`
