package binary_search_test

import (
	bs "binary-search/binary-search"
	"testing"
)

type TestCase struct {
	in  TestCaseIn
	out int
}

type TestCaseIn struct {
	nums   []int
	target int
}

func TestSearch(t *testing.T) {
	cases := []TestCase{
		{
			in: TestCaseIn{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 9,
			},
			out: 4,
		},
		{
			in: TestCaseIn{
				nums:   []int{-1, 0, 3, 5, 9, 12},
				target: 2,
			},
			out: -1,
		},
	}

	for idx, cur := range cases {
		res := bs.Search(cur.in.nums, cur.in.target)

		if res != cur.out {
			t.Errorf("error case: [%d], expected - %d, got - %d", idx, cur.out, res)
		}
	}
}
