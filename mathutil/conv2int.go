package mathutil

/*************************************************************
 * region convert to int
 *************************************************************/

// Int convert value to int
func Int(in any) (int, error) {
	_ = "STUB: not implemented"

	// SafeInt convert value to int, will ignore error
	return 0, nil
}

func SafeInt(in any) int { _ = "STUB: not implemented"; return 0 }

// QuietInt convert value to int, will ignore error
func QuietInt(in any) int {
	_ = "STUB: not implemented"

	// IntOrPanic convert value to int, will panic on error
	return 0
}

func IntOrPanic(in any) int { _ = "STUB: not implemented"; return 0 }

// MustInt convert value to int, will panic on error
func MustInt(in any) int { _ = "STUB: not implemented"; return 0 }

// IntOrDefault convert value to int, return defaultVal on failed
func IntOrDefault(in any, defVal int) int { _ = "STUB: not implemented"; return 0 }

// IntOr convert value to int, return defaultVal on failed
func IntOr(in any, defVal int) int { _ = "STUB: not implemented"; return 0 }

// IntOrErr convert value to int, return error on failed
func IntOrErr(in any) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// ToInt convert value to int, return error on failed
		nil
}

func ToInt(in any) (int, error) {
	_ = "STUB: not implemented"
	return 0,

		// ToIntWith convert value to int, can with some option func.
		//
		// Example:
		//
		//	ToIntWithFunc(val, mathutil.WithNilAsFail, mathutil.WithUserConvFn(func(in any) (int, error) {
		//	})
		nil
}

func ToIntWith(in any, optFns ...ConvOptionFn[int]) (iVal int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// in strict mode, cannot convert string to int

// try convert to int

// default support int ptr type

// eg: json.Number

/*************************************************************
 * region convert to int64
 *************************************************************/

// Int64 convert value to int64, return error on failed
func Int64(in any) (int64, error) {
	_ = "STUB: not implemented"

	// SafeInt64 convert value to int64, will ignore error
	return 0, nil
}

func SafeInt64(in any) int64 { _ = "STUB: not implemented"; return 0 }

// QuietInt64 convert value to int64, will ignore error
func QuietInt64(in any) int64 { _ = "STUB: not implemented"; return 0 }

// MustInt64 convert value to int64, will panic on error
func MustInt64(in any) int64 { _ = "STUB: not implemented"; return 0 }

// Int64OrDefault convert value to int64, return default val on failed
func Int64OrDefault(in any, defVal int64) int64 { _ = "STUB: not implemented"; return 0 }

// Int64Or convert value to int64, return default val on failed
func Int64Or(in any, defVal int64) int64 { _ = "STUB: not implemented"; return 0 }

// ToInt64 convert value to int64, return error on failed
func ToInt64(in any) (int64, error) {
	_ = "STUB: not implemented"
	return 0,

		// Int64OrErr convert value to int64, return error on failed
		nil
}

func Int64OrErr(in any) (int64, error) {
	_ = "STUB: not implemented"
	return 0,

		// ToInt64With try to convert value to int64. can with some option func, more see ConvOption.
		nil
}

func ToInt64With(in any, optFns ...ConvOptionFn[int64]) (i64 int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// in strict mode, cannot convert string to int

// try convert to int64

// default support int64 ptr type

// in strict mode, cannot convert float to int

// in strict mode, cannot convert float to int

// eg: json.Number

/*************************************************************
 * region convert to uint
 *************************************************************/

// Uint convert any to uint, return error on failed
func Uint(in any) (uint, error) {
	_ = "STUB: not implemented"

	// SafeUint convert any to uint, will ignore error
	return 0, nil
}

func SafeUint(in any) uint { _ = "STUB: not implemented"; return 0 }

// QuietUint convert any to uint, will ignore error
func QuietUint(in any) uint {
	_ = "STUB: not implemented"

	// MustUint convert any to uint, will panic on error
	return 0
}

func MustUint(in any) uint { _ = "STUB: not implemented"; return 0 }

// UintOrDefault convert any to uint, return default val on failed
func UintOrDefault(in any, defVal uint) uint { _ = "STUB: not implemented"; return 0 }

// UintOr convert any to uint, return default val on failed
func UintOr(in any, defVal uint) uint { _ = "STUB: not implemented"; return 0 }

// UintOrErr convert value to uint, return error on failed
func UintOrErr(in any) (uint, error) {
	_ = "STUB: not implemented"
	return 0,

		// ToUint convert value to uint, return error on failed
		nil
}

func ToUint(in any) (u64 uint, err error) {
	_ = "STUB: not implemented"
	return 0,

		// ToUintWith try to convert value to uint. can with some option func, more see ConvOption.
		nil
}

func ToUintWith(in any, optFns ...ConvOptionFn[uint]) (uVal uint, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// in strict mode, cannot convert string to int

// try convert to uint64

// default support uint ptr type

// eg: json.Number

/*************************************************************
 * region convert to uint64
 *************************************************************/

// Uint64 convert any to uint64, return error on failed
func Uint64(in any) (uint64, error) {
	_ = "STUB: not implemented"

	// QuietUint64 convert any to uint64, will ignore error
	return 0, nil
}

func QuietUint64(in any) uint64 { _ = "STUB: not implemented"; return 0 }

// SafeUint64 convert any to uint64, will ignore error
func SafeUint64(in any) uint64 { _ = "STUB: not implemented"; return 0 }

// MustUint64 convert any to uint64, will panic on error
func MustUint64(in any) uint64 { _ = "STUB: not implemented"; return 0 }

// Uint64OrDefault convert any to uint64, return default val on failed
func Uint64OrDefault(in any, defVal uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// Uint64Or convert any to uint64, return default val on failed
func Uint64Or(in any, defVal uint64) uint64 { _ = "STUB: not implemented"; return 0 }

// Uint64OrErr convert value to uint64, return error on failed
func Uint64OrErr(in any) (uint64, error) {
	_ = "STUB: not implemented"
	return 0,

		// ToUint64 convert value to uint64, return error on failed
		nil
}

func ToUint64(in any) (uint64, error) {
	_ = "STUB: not implemented"
	return 0,

		// ToUint64With try to convert value to uint64. can with some option func, more see ConvOption.
		nil
}

func ToUint64With(in any, optFns ...ConvOptionFn[uint64]) (u64 uint64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// in strict mode, cannot convert string to int

// try convert to uint64

// default support uint64 ptr type

// eg: json.Number

/*************************************************************
 * region string to intX/uintX
 *************************************************************/

// StrInt convert string to int, ignore error
func StrInt(s string) int { _ = "STUB: not implemented"; return 0 }

// StrIntOr convert string to int, return default val on failed
func StrIntOr(s string, defVal int) int { _ = "STUB: not implemented"; return 0 }

// TryStrInt convert string to int, return error on failed.
//
//   - empty string will return 0.
//   - allow float string.
func TryStrInt(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

// try convert to int

// handle the case where the string might be a float

// TryStrInt64 convert string to int64, return error on failed.
//
//   - empty string will return 0.
//   - allow float string.
func TryStrInt64(s string) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

// handle the case where the string might be a float

// TryStrUint64 try to convert string to uint64, return error on failed
//
//   - empty string will return 0.
//   - allow float string.
func TryStrUint64(s string) (uint64, error) { _ = "STUB: not implemented"; return 0, nil }

// try convert to int64

// handle the case where the string might be a float
