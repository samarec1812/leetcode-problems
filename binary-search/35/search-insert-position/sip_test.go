package search_insert_position

import "testing"

type TestCase struct {
	in  TestCaseIn
	out int
}

type TestCaseIn struct {
	nums   []int
	target int
}

func TestSearchInsert(t *testing.T) {
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
			out: 2,
		},
		{
			in: TestCaseIn{
				nums:   []int{1, 3, 5, 6},
				target: 5,
			},
			out: 2,
		},
		{
			in: TestCaseIn{
				nums:   []int{1, 3, 5, 6},
				target: 2,
			},
			out: 1,
		},
		{
			in: TestCaseIn{
				nums:   []int{1, 3, 5, 6},
				target: 7,
			},
			out: 4,
		},
	}

	for idx, cur := range cases {
		res := SearchInsert(cur.in.nums, cur.in.target)

		if res != cur.out {
			t.Errorf("error case: [%d], expected - %d, got - %d", idx, cur.out, res)
		}
	}
}
