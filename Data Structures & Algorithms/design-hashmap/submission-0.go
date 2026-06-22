type MyHashMap struct {
	data map[int]int
}

func Constructor() MyHashMap {
    data := make(map[int]int)
	return MyHashMap {
		data: data,
	}
}

func (this *MyHashMap) Put(key int, value int) {
    this.data[key] = value
}

func (this *MyHashMap) Get(key int) int {
    value, ok := this.data[key]
	if ok {
		return value
	} else {
		return -1
	}
}

func (this *MyHashMap) Remove(key int) {
    this.data[key] = -1
}

/**
 * Your MyHashMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Put(key,value);
 * param_2 := obj.Get(key);
 * obj.Remove(key);
 */