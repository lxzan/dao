package multi_queue

type Pointer uint32

type interator[T any] interface {
	Empty(ptr Pointer) bool
	End(ele *T) bool
	Get(ptr Pointer) *T
	Next(ele *T) *T
}

func for_each[T](iter interator[T], start Pointer, fn func(ele *T) (next bool)) {
	if iter.Empty(start) {
		return
	}

	var ele = iter.Get(start)
	var flag = true
	for {
		flag = fn(ele)
		if !iter.End(ele) && flag {
			ele = iter.Next(ele)
		} else {
			break
		}
	}
}

