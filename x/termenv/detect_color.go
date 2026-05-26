package termenv

// ColorLevel is the color level supported by a terminal.
type ColorLevel uint8

const (
	TermColorNone ColorLevel = iota // not support color
	TermColor16                     // 16(4bit) ANSI color supported
	TermColor256                    // 256(8bit) color supported
	TermColorTrue                   // support TRUE(RGB) color
)

// String returns the string name of the color level.
func (l ColorLevel) String() string { _ = "STUB: not implemented"; return "" }

// NoColor returns true if the NO_COLOR environment variable is set.
func NoColor() bool {
	_ = "STUB: not implemented"

	// TermColorLevel returns the color support level for the current terminal.
	return false
}

func TermColorLevel() ColorLevel {
	_ = "STUB: not implemented"

	// IsSupportColor returns true if the terminal supports color.
	return *new(ColorLevel)
}

func IsSupportColor() bool { _ = "STUB: not implemented"; return false }

// IsSupport256Color returns true if the terminal supports 256 colors.
func IsSupport256Color() bool { _ = "STUB: not implemented"; return false }

// IsSupportTrueColor returns true if the terminal supports true color.
func IsSupportTrueColor() bool { _ = "STUB: not implemented"; return false }

//
// ---------------- Force set color support ----------------
//

var backLevel ColorLevel

// SetColorLevel value force.
func SetColorLevel(level ColorLevel) {
	_ = "STUB: not implemented"
	// backup old value
	return
}

// force set color level

// DisableColor in the current terminal
func DisableColor() {
	_ = "STUB: not implemented"
	// backup old value
	return
}

// force disable color

// ForceEnableColor flags value. TIP: use for unit testing.
//
// Usage:
//
//	ccolor.ForceEnableColor()
//	defer ccolor.RevertColorSupport()
func ForceEnableColor() {
	_ = "STUB: not implemented"
	// backup old value
	return
}

// force enables color

// return colorLevel

// RevertColorSupport flags to init value.
func RevertColorSupport() {
	_ = "STUB: not implemented"
	// revert color flags var
	return
}

/*************************************************************
 * helper methods for detect color supports
 *************************************************************/

// DetectColorLevel for current env
//
// NOTICE: The method will detect terminal info each time.
//
//	if only want to get current color level, please direct call IsSupportColor() or TermColorLevel()
func DetectColorLevel() ColorLevel { _ = "STUB: not implemented"; return *new(ColorLevel) }

// on TERM=screen: not support true-color
const noTrueColorTerm = "screen"

// detect terminal color support level
//
// refer https://github.com/Delta456/box-cli-maker
func detectTermColorLevel() (level ColorLevel, needVTP bool) {
	_ = "STUB: not implemented"
	return *new(ColorLevel), false
}

// On JetBrains Terminal
// - TERM value not set, but support true-color
// env:
// 	TERMINAL_EMULATOR=JetBrains-JediTerm

// fallback: simple detect by TERM value string.

// detectColorFromEnv returns the color level COLORTERM, FORCE_COLOR,
// TERM_PROGRAM, or determined from the TERM environment variable.
//
// refer the github.com/xo/terminfo.ColorLevelFromEnv()
// https://en.wikipedia.org/wiki/Terminfo
func detectColorLevelFromEnv(termVal string, isWin bool) ColorLevel {
	_ = "STUB: not implemented"
	return *new(ColorLevel)
}

// on TERM=screen: not support true-color

// check for overriding environment variables

// check iTerm version

// return TermColorNone

// otherwise determine from TERM's max_colors capability
// if !isWin && termVal != "" {
// 	debugf("TERM=%s - TODO check color level by load terminfo file", termVal)
// 	return TermColor16
// }

// no TERM env value. default return none level
