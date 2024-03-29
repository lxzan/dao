package hashmap

type HashMap[K comparable, V any] map[K]V

// New instantiates a hashmap
// at most one param, means initial capacity
func New[K comparable, V any](capacity int) HashMap[K, V] {
	return make(map[K]V, capacity)
}

//Reset clear contents
//func (c HashMap[K, V]) Reset() {
//	keys := c.Keys()
//	for _, key := range keys {
//		delete(c, key)
//	}
//}

// Len get the length of hashmap
func (c HashMap[K, V]) Len() int {
	return len(c)
}

// Set insert a element into the hashmap
// if key exists, value will be replaced
func (c HashMap[K, V]) Set(key K, val V) {
	c[key] = val
}

// Get search if hashmap contains the key
func (c HashMap[K, V]) Get(key K) (val V, exist bool) {
	val, exist = c[key]
	return
}

// Exists if key exists, return true
func (c HashMap[K, V]) Exists(key K) bool {
	_, ok := c[key]
	return ok
}

// Delete delete a element if the key exists
func (c HashMap[K, V]) Delete(key K) {
	delete(c, key)
}

func (c HashMap[K, V]) Range(f func(key K, val V) bool) {
	for k, v := range c {
		if !f(k, v) {
			return
		}
	}
}

// Keys get all the keys of the hashmap, construct it as a dynamic array and return it
func (c HashMap[K, V]) Keys() []K {
	var keys = make([]K, 0, c.Len())
	c.Range(func(k K, v V) bool {
		keys = append(keys, k)
		return true
	})
	return keys
}

// Values get all the values of the hashmap, construct it as a dynamic array and return it
func (c HashMap[K, V]) Values() []V {
	var vals = make([]V, 0, c.Len())
	c.Range(func(k K, v V) bool {
		vals = append(vals, v)
		return true
	})
	return vals
}
