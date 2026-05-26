package assert

import "reflect"

// Nil asserts that the given is a nil value
func (as *Assertions) Nil(give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotNil asserts that the given is a not nil value
func (as *Assertions) NotNil(val any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// True check, please see True()
func (as *Assertions) True(give bool, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// False check, please see False()
func (as *Assertions) False(give bool, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Empty check, please see Empty()
func (as *Assertions) Empty(give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotEmpty check, please see NotEmpty()
func (as *Assertions) NotEmpty(give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Zero check, please see Zero()
func (as *Assertions) Zero(give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotZero check, please see NotZero()
func (as *Assertions) NotZero(give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Panics check, please see Panics()
func (as *Assertions) Panics(fn PanicRunFunc, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotPanics check, please see NotPanics()
func (as *Assertions) NotPanics(fn PanicRunFunc, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// PanicsMsg check, please see PanicsMsg()
func (as *Assertions) PanicsMsg(fn PanicRunFunc, wantVal any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// PanicsErrMsg check, please see PanicsErrMsg()
func (as *Assertions) PanicsErrMsg(fn PanicRunFunc, errMsg string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Contains asserts that the given data(string,slice,map) should contain element
func (as *Assertions) Contains(src, elem any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotContains asserts that the given data(string,slice,map) should not contain element
func (as *Assertions) NotContains(src, elem any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// ContainsKey asserts that the given map is contains key
func (as *Assertions) ContainsKey(mp, key any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotContainsKey asserts that the given map is not contains key
func (as *Assertions) NotContainsKey(mp, key any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// ContainsKeys asserts that the map is contains all given keys
func (as *Assertions) ContainsKeys(mp any, keys any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotContainsKeys asserts that the map is not contains all given keys
func (as *Assertions) NotContainsKeys(mp any, keys any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// ContainsElems asserts that the given list should contain sub elements
func (as *Assertions) ContainsElems(list any, sub any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// StrContains asserts that the given strings is contains sub-string
func (as *Assertions) StrContains(s, sub string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// StrNotContains asserts that the given strings is not contains sub-string
func (as *Assertions) StrNotContains(s, sub string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// StrContainsAll asserts that the given strings is contains all sub-strings
func (as *Assertions) StrContainsAll(s string, subs []string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// StrCount asserts that the given strings contains substring count
func (as *Assertions) StrCount(s, sub string, count int, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NoErr asserts that the given is a nil error
func (as *Assertions) NoErr(err error, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NoError asserts that the given is a nil error
func (as *Assertions) NoError(err error, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Err asserts that the given is a not nil error
func (as *Assertions) Err(err error, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Error asserts that the given is a not nil error
func (as *Assertions) Error(err error, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// ErrIs asserts that the given error is equals wantErr
func (as *Assertions) ErrIs(err, wantErr error, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// ErrMsg asserts that the given is a not nil error and error message equals wantMsg
func (as *Assertions) ErrMsg(err error, errMsg string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// ErrSubMsg asserts that the given is a not nil error and the error message contains subMsg
func (as *Assertions) ErrSubMsg(err error, subMsg string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// ErrMsgContains asserts that the given is a not nil error and error message contains subMsg
func (as *Assertions) ErrMsgContains(err error, subMsg string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// ErrHasMsg asserts that the given is a not nil error and the error message contains subMsg
func (as *Assertions) ErrHasMsg(err error, subMsg string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// FileExists asserts that the given file exists
func (as *Assertions) FileExists(filePath string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// FileNotExists asserts that the given file not exists
func (as *Assertions) FileNotExists(filePath string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// DirExists asserts that the given dir exists
func (as *Assertions) DirExists(dirPath string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// DirNotExists asserts that the given dir not exists
func (as *Assertions) DirNotExists(dirPath string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Len assert given length is equals to wantLn
func (as *Assertions) Len(give any, wantLn int, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// LenGt assert given length is greater than to minLn
func (as *Assertions) LenGt(give any, minLn int, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Eq asserts that the want should equal to the given
func (as *Assertions) Eq(want, give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Equal asserts that the want should equal to the given
//
// Alias of Eq()
func (as *Assertions) Equal(want, give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Neq asserts that the want should not be equal to the given.
// alias of NotEq()
func (as *Assertions) Neq(want, give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotEq asserts that the want should not be equal to the given
func (as *Assertions) NotEq(want, give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotEqual asserts that the want should not be equal to the given
//
// Alias of NotEq()
func (as *Assertions) NotEqual(want, give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Lt asserts that the give(intX) should not be less than max
func (as *Assertions) Lt(give, max any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Lte asserts that the give(intX) should not be less than or equal to max
func (as *Assertions) Lte(give, max any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Gt asserts that the give(intX) should not be greater than min
func (as *Assertions) Gt(give, min any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Gte asserts that the give(intX) should not be greater than or equal to min
func (as *Assertions) Gte(give, min any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// EqInt asserts that the want should equal to the given intX
func (as *Assertions) EqInt(want, give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// EqFloat asserts that the want should equal to the given float with delta
func (as *Assertions) EqFloat(want, give any, delta float64, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// InDelta asserts that two floating-point values differ within a certain range
func (as *Assertions) InDelta(want, give any, delta float64, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// IsType type equals assert
func (as *Assertions) IsType(wantType, give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// IsKind asserts that the given data reflect.Kind equals
func (as *Assertions) IsKind(wantKind reflect.Kind, give any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Same asserts that two pointers reference the same object
func (as *Assertions) Same(wanted, actual any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// NotSame asserts that two pointers do not reference the same object
func (as *Assertions) NotSame(want, actual any, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// Fail reports a failure through
func (as *Assertions) Fail(failMsg string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}

// FailNow fails test
func (as *Assertions) FailNow(failMsg string, fmtAndArgs ...any) *Assertions {
	_ = "STUB: not implemented"
	return nil
}
