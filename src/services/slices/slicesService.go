package slices

func Merge(s1 []string, s2 []string) []string {
	s1s2 := append(s1, s2...)
	var tmp []string
	for _, sv := range s1s2 {
		if len(tmp) == 0 {
			tmp = append(tmp, sv)
		} else {
			for k, v := range tmp {
				if sv == v {
					break
				}
				if k == len(tmp)-1 {
					tmp = append(tmp, sv)
				}
			}
		}
	}

	return tmp
}
