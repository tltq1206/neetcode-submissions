import "slices"

type MyHashSet struct {
	slice []int
	contain map[int]bool
}

func Constructor() MyHashSet {
	slice := []int{}
	contain := make(map[int]bool)
    return MyHashSet {
		slice: slice,
		contain: contain,
	}
}

func (this *MyHashSet) Add(key int) {
    if this.contain[key] == false {
		this.slice = append(this.slice, key)
		this.contain[key] = true
	}
}

func (this *MyHashSet) Remove(key int) {
    for i :=0; i<len(this.slice); i++ {
		if this.slice[i] == key {
			slices.Delete(this.slice, i, i+1)
			this.contain[key] = false
			break
		}
	}
}

func (this *MyHashSet) Contains(key int) bool {
    return this.contain[key]
}

/**
 * Your MyHashSet object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
 