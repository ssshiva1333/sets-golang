package set

type set []int

func (set *set) add(element int) {
	*set = append(*set, element)
}

func (set *set) remove(element_index int) {
	*set = append((*set)[:element_index], (*set)[:element_index+1]...)
}

func (set *set) find(element int) (wanted_element int) {
	for i := 0; i <= len(*set); i++ {
		if (*set)[i] == element {
			wanted_element = i
		}
	}

	return
}

func (set *set) combine(slice []int) {
	for i := 0; i < len(slice); i++ {
		indicator := false

		for j := 0; j < len(*set); j++ {
			if slice[i] == (*set)[j] {
				indicator = true
			}
		}

		if indicator == false {
			*set = append(*set, slice[i])
		}
	}
}

func (set *set) intersection(slice []int) (intersection_set []int) {
	for i := 0; i < len(*set); i++ {
		for j := 0; j < len(slice); j++ {
			if (*set)[i] == slice[j] {
				intersection_set = append(intersection_set, slice[j])
			}
		}
	}

	return
}

func (set *set) difference(slice []int) (difference_set []int) {
	for i := 0; i < len(slice); i++ {
		indicator := false

		for j := 0; j < len(*set); j++ {
			if slice[i] == (*set)[j] {
				indicator = true

				break
			}
		}

		if indicator == false {
			difference_set = append(difference_set, slice[i])
		}
	}

	for i := 0; i < len(*set); i++ {
		indicator := false

		for j := 0; j < len(slice); j++ {
			if (*set)[i] == slice[j] {
				indicator = true

				break
			}
		}

		if indicator == false {
			for k := 0; k < len(difference_set); k++ {
				if difference_set[k] == (*set)[i] {
					continue
				} else {
					difference_set = append(difference_set, (*set)[i])
				}
			}
		}
	}

	return
}

func (set *set) cartesian_set(slice []int) (cartesian_set map[int]int) {
	for i := 0; i < len(*set); i++ {
		for j := 0; j < len(slice); j++ {
			cartesian_set[(*set)[i]] = slice[j]
		}
	}

	return
}
