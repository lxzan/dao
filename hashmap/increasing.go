package hashmap

func (this *HashMap[K, V]) increase() {
	if float64(this.storage.Length)/float64(this.size) > this.load_factor {
		var m = New[K, V](this.size * 2)
		var n = len(this.storage.Buckets)
		for i := 1; i < n; i++ {
			var dst = &this.storage.Buckets[i]
			if dst.Ptr != 0 {
				var idx = dst.Data.hashCode & (m.size - 1)
				var entrypoint = &m.indexes[idx]
				if entrypoint.Head == 0 {
					var ptr = m.storage.NextID()
					entrypoint.Head = ptr
					entrypoint.Tail = ptr
				}
				m.storage.Push(entrypoint, &dst.Data)
			}
		}
		*this = *m
	}
}

//func (c *HashMap) migrate() {
//	if c.old.storage == nil {
//		return
//	}
//
//	var i = c.old.offset
//	for ; i < c.old.cap; i++ {
//		var entrypoint = &c.old.indexes[i]
//		if entrypoint.Head != 0 {
//			var arr = make([]rapid.Pointer, 0)
//			for j := c.old.storage.Begin(*entrypoint); !c.old.storage.End(j); j = c.old.storage.Next(j) {
//				c.incrSet(&j.Data)
//				arr = append(arr, j.Ptr)
//			}
//			for _, ptr := range arr {
//				c.old.storage.Buckets[ptr].Reset()
//			}
//			entrypoint.Head = 0
//			entrypoint.Tail = 0
//			c.old.offset = i
//			break
//		}
//	}
//	c.old.offset++
//	if i >= c.old.cap {
//		c.resetOld()
//	}
//}
