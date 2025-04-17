package option

// Option 是一个泛型结构体，用于表示一个可能存在或不存在的值。
type Option[T any] struct {
	v  T    // 值
	ok bool // 是否存在值
}

// Some 创建一个包含值的 Option。
// 参数：
//   v T - 要包含的值
// 返回：
//   Option[T] - 包含值的 Option
func Some[T any](v T) Option[T] {
	return Option[T]{v: v, ok: true}
}

// None 创建一个不包含值的 Option。
// 返回：
//   Option[T] - 不包含值的 Option
func None[T any]() Option[T] {
	var zero T
	return Option[T]{v: zero, ok: false}
}

// Wrap 根据给定的值和状态创建一个 Option。
// 参数：
//   v T - 值
//   ok bool - 是否存在值
// 返回：
//   Option[T] - 包含值或不包含值的 Option
func Wrap[T any](v T, ok bool) Option[T] {
	return Option[T]{v: v, ok: ok}
}

// Value 返回 Option 中的值和状态。
// 返回：
//   T - 值
//   bool - 是否存在值
func (o *Option[T]) Value() (T, bool) {
	return o.v, o.ok
}

// ValueOrError 返回 Option 中的值或错误。
// 参数：
//   err error - 当值不存在时返回的错误
// 返回：
//   T - 值
//   error - 错误
func (o *Option[T]) ValueOrError(err error) (T, error) {
	if !o.ok {
		var zero T
		return zero, err
	}
	return o.v, nil
}

// Match 根据 Option 的状态执行不同的回调函数。
// 参数：
//   someFn func(T) - 当值存在时执行的回调函数
//   noneFn func() - 当值不存在时执行的回调函数
func (o *Option[T]) Match(someFn func(T), noneFn func()) {
	if o.ok {
		someFn(o.v)
	} else {
		noneFn()
	}
}

// IfSome 当值存在时执行回调函数。
// 参数：
//   someFn func(T) - 回调函数
func (o *Option[T]) IfSome(someFn func(T)) {
	if o.ok {
		someFn(o.v)
	}
}

// IfNone 当值不存在时执行回调函数。
// 参数：
//   noneFn func() - 回调函数
func (o *Option[T]) IfNone(noneFn func()) {
	if !o.ok {
		noneFn()
	}
}

// Unwrap_unsafe 返回 Option 中的值，如果值不存在则触发 panic。
// 返回：
//   T - 值
func (o *Option[T]) Unwrap_unsafe() T {
	if !o.ok {
		panic("Value is nil")
	}
	return o.v
}

// Expect_unsafe 返回 Option 中的值，如果值不存在则触发 panic 并显示错误信息。
// 参数：
//   err error - 错误信息
// 返回：
//   T - 值
func (o *Option[T]) Expect_unsafe(err error) T {
	if !o.ok {
		panic(err)
	}
	return o.v
}
