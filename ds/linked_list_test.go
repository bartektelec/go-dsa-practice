package ds

import "testing"

func TestGetNth(t *testing.T) {
	list := &Node[int]{
		value: 5,
		next: &Node[int]{
			value: 8,
			next: &Node[int]{
				value: 10,
			},
		},
	}

	if res, err := GetNth[int](list, 0); err && res.value != 5 {
		t.Fatal()
	}
	if res, err := GetNth[int](list, 1); err && res.value != 8 {
		println("Res was ", res.value)
		t.Fatal()
	}
	if res, err := GetNth[int](list, 2); err && res.value != 10 {
		t.Fatal()
	}
}

func TestInsert(t *testing.T) {
	list := &Node[int]{
		value: 5,
		next: &Node[int]{
			value: 8,
			next: &Node[int]{
				value: 10,
			},
		},
	}

	InsertNth(list, 1, 14)

	if res, err := GetNth[int](list, 1); err && res.value != 8 {
		println("Res was ", res.value)
		t.Fatal()
	}
	if res, err := GetNth[int](list, 3); err && res.value != 14 {
		println("Res was ", res.value)
		t.Fatal()
	}
	// if res, err := GetNth[int](list, 3); err && res.value != 10 {
	// 	println("Res was ", res.value)
	// 	t.Fatal()
	// }
}
