package types

import "fmt"

var ErrNone = fmt.Errorf("option is a none value")

type Option[T any] struct {
	value any
}

func None[T any]() Option[T] {
	return Option[T]{value: nil}
}

func Some[T any](value T) Option[T] {
	return Option[T]{value}
}

func (o *Option[T]) IsNone() bool {
	return o.value == nil
}

func (o *Option[T]) IsSome() bool {
	return !o.IsNone()
}

func (o *Option[T]) Some() (T, bool) {
	if o.IsNone() {
		var zero T
		return zero, false
	} else {
		return o.value.(T), true
	}
}

func (o *Option[T]) Unwrap() T {
	if value, ok := o.Some(); ok {
		return value
	}
	panic(ErrNone)
}

func (o *Option[T]) Take() (T, bool) {
	if o.IsNone() {
		var zero T
		return zero, false
	} else {
		value := o.value.(T)
		o.value = nil
		return value, true
	}
}

func (o *Option[T]) Ok() Result[T] {
	if value, ok := o.Some(); ok {
		return Ok[T](value)
	} else {
		return Err[T](ErrNone)
	}
}
