package aocfuncs

type Consumer[T any] func(item T)

func ForEach[T any](items []T, f Consumer[T]) {
	for _, item := range items {
		f(item)
	}
}

type Mapper[I, O any] func(item I) O

func Map[I, O any](items []I, f Mapper[I, O]) []O {
	result := make([]O, len(items))
	for _, item := range items {
		result = append(result, f(item))
	}
	return result
}

type Reducer[I, O any] func(accumulator O, value I) O

func Reduce[I, O any](items []I, f Reducer[I, O], initial O) O {
	result := initial
	for _, item := range items {
		result = f(result, item)
	}
	return result
}
