package fp

type Range struct {
	start int
	end   int
}

func RangeOfou(args ...int) Range {
	if len(args) == 0 {
		return Range{0, 0}
	}

	if len(args) == 1 {
		return Range{0, args[0]}
	}

	return Range{args[0], args[1]}
}

func (r Range) Forward(action func(index int)) {
	for i := r.start; i < r.end; i++ {
		action(i)
	}
}

func (r Range) Iterate(stop func(index int) bool) int {
	i := r.start
	for i < r.end && !stop(i) {
		i++
	}

	return i
}

func (r Range) reduceToString(transform func(index int) string) string {
	result := ""
	r.Forward(func(index int) {
		result += transform(index)
	})
	return result
}

func (r Range) JoinToString(separator string, transform func(index int) string) string {
	return r.reduceToString(func(index int) string {
		accum := transform(index)
		if index != r.end-1 {
			accum += separator
		}
		return accum
	})
}

func (r Range) Contains(predicate func(index int) bool) bool {
	stop := func(index int) bool {
		return predicate(index)
	}
	return r.Iterate(stop) != r.end
}
