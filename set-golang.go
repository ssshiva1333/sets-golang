package setgolang

type set []int

func (set *set) add(element int) {
	*set = append(*set, element)
}

func (set *set) remove(element_index int) {
	*set = append((*set)[:element_index], (*set)[:element_index+1]...)
}

func (set *set) find(element int) (wanted_element int) {
	wanted_element = -1
	for i := 0; i < len(*set); i++ {
		if (*set)[i] == element {
			wanted_element = i
		}
	}

	return
}

func (set *set) combine(slice_set []int) {
	for i := 0; i < len(slice_set); i++ {
		indicator := false

		for j := 0; j < len(*set); j++ {
			if slice_set[i] == (*set)[j] {
				indicator = true
			}
		}

		if indicator == false {
			*set = append(*set, slice_set[i])
		}
	}
}

func (set *set) intersection(slice_set []int) (intersection_set []int) {
	for i := 0; i < len(*set); i++ {
		for j := 0; j < len(slice_set); j++ {
			if (*set)[i] == slice_set[j] {
				intersection_set = append(intersection_set, slice_set[j])
			}
		}
	}

	return
}

func (set *set) difference(slice_set []int) (difference_set []int) {
	for i := 0; i < len(*set); i++ {
		indicator := false
		for j := 0; j < len(slice_set); j++ {
			if (*set)[i] == slice_set[j] {
				indicator = true
			}
		}

		if indicator != true {
			difference_set = append(difference_set, (*set)[i])
		}
	}

	for i := 0; i < len(slice_set); i++ {
		indicator := false
		for j := 0; j < len(*set); j++ {
			if slice_set[i] == (*set)[j] {
				indicator = true
			}
		}

		if indicator == false {
			inner_indicator := false
			for m := 0; m < len(difference_set); m++ {
				if slice_set[i] == difference_set[m] {
					inner_indicator = true
				}
			}

			if inner_indicator == false {
				difference_set = append(difference_set, slice_set[i])
			}
		}
	}

	return
}

func (set *set) cartesian(slice_set []int) (cartesian_set [][]int) {
	for i := 0; i < len(*set); i++ {
		for j := 0; j < len(slice_set); j++ {
			set_pair := []int{(*set)[i], slice_set[j]}

			cartesian_set = append(cartesian_set, set_pair)
		}
	}

	return cartesian_set
}
