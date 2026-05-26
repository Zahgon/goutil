package reflects

import (
	"errors"
	"reflect"
)

// FuncX wrap a go func. represent a function
type FuncX struct {
	CallOpt
	// Name of func. eg: "MyFunc"
	Name string
	// rv is the `reflect.Value` of func
	rv reflect.Value
	rt reflect.Type
}

// NewFunc instance. param fn support func and reflect.Value
func NewFunc(fn any) *FuncX { _ = "STUB: not implemented"; return nil }

// NumIn get the number of func input args
func (f *FuncX) NumIn() int { _ = "STUB: not implemented"; return 0 }

// NumOut get the number of func output args
func (f *FuncX) NumOut() int { _ = "STUB: not implemented"; return 0 }

// Call the function with given arguments.
//
// Usage:
//
//	func main() {
//		fn := func(a, b int) int {
//			return a + b
//		}
//
//		fx := NewFunc(fn)
//		ret, err := fx.Call(1, 2)
//		fmt.Println(ret[0], err) // Output: 3 <nil>
//	}
func (f *FuncX) Call(args ...any) ([]any, error) {
	_ = "STUB: not implemented"
	// convert args to []reflect.Value
	return nil, nil
}

// convert ret to []any

// Call2 returns the result of evaluating the first argument as a function.
// The function must return 1 result, or 2 results, the second of which is an error.
//
//   - Only support func with 1 or 2 return values: (val) OR (val, err)
//   - Will check args and try to convert input args to func args type.
func (f *FuncX) Call2(args ...any) (any, error) {
	_ = "STUB: not implemented"
	// convert args to []reflect.Value
	return *new(any), nil
}

// do call func

// func return like: (val, err)

// CallRV call the function with given reflect.Value arguments.
func (f *FuncX) CallRV(args []reflect.Value) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// WithTypeChecker set type checker
func (f *FuncX) WithTypeChecker(checker TypeCheckerFn) *FuncX {
	_ = "STUB: not implemented"
	return nil
}

// WithEnhanceConv set enhance convert
func (f *FuncX) WithEnhanceConv() *FuncX { _ = "STUB: not implemented"; return nil }

// String of func
func (f *FuncX) String() string { _ = "STUB: not implemented"; return "" }

// TypeCheckerFn type checker func
type TypeCheckerFn func(typ reflect.Type) error

// CallOpt call options
type CallOpt struct {
	// TypeChecker check func type before call func. eg: check return values
	TypeChecker TypeCheckerFn
	// EnhanceConv try to enhance auto convert args to func args type
	// 	- support more type: string, int, uint, float, bool
	EnhanceConv bool
}

// OneOrTwoOutChecker check func type. only allow 1 or 2 return values
//
// Allow func returns:
//   - 1 return: (value)
//   - 2 return: (value, error)
var OneOrTwoOutChecker = func(typ reflect.Type) error {
	if !good1or2outFunc(typ) {
		return errors.New("func allow with 1 result or 2 results where the second is an error")
	}
	return nil
}

//
// TIP:
// 	flow func refer from text/template package.
//
//

// reports whether the function or method has the right result signature.
func good1or2outFunc(typ reflect.Type) bool {
	_ = "STUB: not implemented"
	// We allow functions with 1 result or 2 results where the second is an error.
	return false
}

// Call2 returns the result of evaluating the first argument as a function.
// The function must return 1 result, or 2 results, the second of which is an error.
//
//   - Only support func with 1 or 2 return values: (val) OR (val, err)
//   - Will check args and try convert input args to func args type.
func Call2(fn reflect.Value, args []reflect.Value) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// func return like: (val, err)

// Call returns the result of evaluating the first argument as a function.
//
//   - Will check args and try convert input args to func args type.
//
// from text/template/funcs.go#call
func Call(fn reflect.Value, args []reflect.Value, opt *CallOpt) ([]reflect.Value, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Convert each arg to the type of the function's arg.

// Compute the expected type. Clumsy because of variadic.

// SafeCall2 runs fun.Call(args), and returns the resulting value and error, if
// any. If the call panics, the panic value is returned as an error.
//
// NOTE: Only support func with 1 or 2 return values: (val) OR (val, err)
//
// from text/template/funcs.go#safeCall
func SafeCall2(fun reflect.Value, args []reflect.Value) (val reflect.Value, err error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// func return like: (val, err)

// SafeCall runs fun.Call(args), and returns the resulting values, or an error.
// If the call panics, the panic value is returned as an error.
func SafeCall(fun reflect.Value, args []reflect.Value) (ret []reflect.Value, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// prepareArg checks if value can be used as an argument of type argType, and
// converts an invalid value to appropriate zero if possible.
func prepareArg(value reflect.Value, argType reflect.Type, enhanced bool) (reflect.Value, error) {
	_ = "STUB: not implemented"
	return *new(reflect.Value), nil
}

// If the argument is an int-like type, and the value is an int-like type, auto-convert.

// enhance convert value to argType, support more type: string, int, uint, float, bool
