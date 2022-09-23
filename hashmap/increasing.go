package hashmap

func (c *HashMap[K, V]) increase() {
	if float64(c.storage.Length)/float64(c.size) > c.load_factor {
		var m = New[K, V](c.size * 2)
		var n = len(c.storage.Buckets)
		for i := 1; i < n; i++ {
			var dst = &c.storage.Buckets[i]
			if dst.Ptr != 0 {
				var idx = dst.Data.hashCode & (m.size - 1)
				var entrypoint = &m.indexes[idx]
				m.storage.Push(entrypoint, &dst.Data)
			}
		}
		*c = *m
	}
}
