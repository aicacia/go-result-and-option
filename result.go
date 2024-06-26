package result_and_option

type Result[T any] struct {
	value any
	err   error
}

func Err[T any](err error) Result[T] {
	return Result[T]{value: nil, err: err}
}

func Ok[T any](value T) Result[T] {
	return Result[T]{value: value, err: nil}
}

func (o *Result[T]) IsErr() bool {
	return o.err != nil
}

func (o *Result[T]) IsOk() bool {
	return !o.IsErr()
}

func (o *Result[T]) Ok() (T, bool) {
	if o.IsErr() {
		var zero T
		return zero, false
	} else {
		return o.value.(T), true
	}
}

func (o *Result[T]) Err() (error, bool) {
	if o.IsOk() {
		return nil, false
	} else {
		return o.err, true
	}
}

func (o *Result[T]) Unwrap() T {
	if value, ok := o.Ok(); ok {
		return value
	}
	panic(o.err)
}

func (o *Result[T]) UnwrapErr() error {
	if err, ok := o.Err(); ok {
		return err
	}
	panic(o.value)
}

func (o *Result[T]) Option() Option[T] {
	if value, ok := o.Ok(); ok {
		return Some(value)
	} else {
		return None[T]()
	}
}
