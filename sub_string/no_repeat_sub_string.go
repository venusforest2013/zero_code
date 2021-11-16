package sub_string

func GetLongestNoRepeatSubString(str string) string {
	//边界判断
	n := len(str)
	if n == 0 {
		return ""
	}

	//定义中间变量，map，index，max，firstIndex
	mp := make(map[byte]int, 0)
	max := 0
	index := 0
	firstIndex := 0
	//循环求解
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(mp, str[i])
		}
		for index < n && mp[str[index]] == 0 {
			mp[str[index]]++
			index++
		}
		diff := index - i
		if diff > max {
			max = diff
			firstIndex = i
		}
	}
	return str[firstIndex:max]

}
