package model

func (p Limit) GetFirst() *int {
	if p.First != nil {
		if *p.First > 100 {
			first := 100
			return &first
		}
	}
	return p.First
}

func (p Limit) GetLast() *int {
	if p.Last != nil {
		if *p.Last > 100 {
			last := 100
			return &last
		}
	}
	return p.First
}
