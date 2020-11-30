package map

func findIndexByBinarySearch(s []int, k int) (int, bool) {
	lo, hi := 0, len(s)
	var m int
	max := len(s)
	if max == 0 {
		return 0, false
	}
	res := false
	for lo <= hi {
		m = (lo + hi) >> 1
		if m == 0 && s[0] > k {
			return 0, res
		}
		if m == max-1 && s[max-1] < k {
			return m + 1, res
		}
		if s[m] < k && s[m+1] > k {
			return m + 1, res
		}
		if s[m] > k && s[m-1] < k {
			return m, res
		}
		if s[m] < k {
			lo = m + 1
		} else if s[m] > k {
			hi = m - 1
		} else {
			return m, true
		}
	}
	return -1, false
}

type Int_Map struct {
	dataMap  map[int]interface{}
	keyArray []int
}

func NewIntMap(cap int) *Int_Map {
	return &Int_Map{
		dataMap:  make(map[int]interface{}),
		keyArray: make([]int, 0, cap),
	}
}

func (m *Int_Map) Exists(key int) bool {
	_, exists := m.dataMap[key]
	return exists
}

func (m *Int_Map) Insert(key int, data interface{}) bool {
	m.dataMap[key] = data
	index, res := findIndexByBinarySearch(m.keyArray, key)
	if index == -1 {
		return false
	}
	if res == true { //存在则直接返回
		return true
	}
	if len(m.keyArray) == 0 {
		m.keyArray = append(m.keyArray, key)
		return true
	}
	//追加末尾
	if index >= len(m.keyArray) {
		m.keyArray = append(m.keyArray[0:], []int{key}...)
	} else if index == 0 { //追加头部
		m.keyArray = append([]int{key}, m.keyArray[:len(m.keyArray)]...)
	} else { //插入
		rear := append([]int{}, m.keyArray[index:]...)
		m.keyArray = append(m.keyArray[0:index], key)
		m.keyArray = append(m.keyArray, rear...)
	}
	return true
}

func (m *Int_Map) Erase(key int) {
	if !m.Exists(key) {
		return
	}
	index, res := findIndexByBinarySearch(m.keyArray, key)
	if res == false {
		return
	}
	delete(m.dataMap, key)
	if index == 0 {
		m.keyArray = m.keyArray[1:]
	} else if index == len(m.keyArray) {
		m.keyArray = m.keyArray[:len(m.keyArray)-2]
	} else {
		m.keyArray = append(m.keyArray[:index], m.keyArray[index+1:]...)
	}

}

func (m *Int_Map) Size() int {
	return len(m.keyArray)
}

func (m *Int_Map) GetByOrderIndex(index int) (int, interface{}, bool) {
	if index < 0 || index >= len(m.keyArray) {
		return -1, nil, false
	}
	key := m.keyArray[index]
	return key, m.dataMap[key], true
}