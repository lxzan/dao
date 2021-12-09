package trie

import "testing"

func TestNewTrie(t *testing.T) {
	var d = NewTrie()
	d.Set("caster", 1)
	d.Set("teemo", 2)
	d.Set("hasaki", 3)
	d.Set("tesla", 4)

	{
		v, ok := d.Get("caster")
		if !ok || v != 1 {
			t.Error("error!")
		}
	}

	{
		v, ok := d.Get("teemo")
		if !ok || v != 2 {
			t.Error("error!")
		}
	}

	{
		v, ok := d.Get("hasaki")
		if !ok || v != 3 {
			t.Error("error!")
		}
	}

	{
		v, ok := d.Get("xxx")
		if ok || v != 0 {
			t.Error("error!")
		}
	}

	res := d.Match("tesla")
	println(&res)
}
