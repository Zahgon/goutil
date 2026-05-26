package termenv

var (
	cmdList  = []string{"cmd", "cmd.exe"}
	pwshList = []string{"powershell", "powershell.exe", "pwsh", "pwsh.exe"}
)

// IsTerminal 检查是否为终端设备中
func IsTerminal() bool { _ = "STUB: not implemented"; return false }

// CurrentShell get current used shell env file.
//
// eg "/bin/zsh" "/bin/bash".
// if onlyName=true, will return "zsh", "bash"
func CurrentShell(onlyName bool, fallbackShell ...string) string {
	_ = "STUB: not implemented"
	return ""
}

// HasShellEnv has shell env check.
//
// Usage:
//
//	HasShellEnv("sh")
//	HasShellEnv("bash")
func HasShellEnv(shell string) bool {
	_ = "STUB: not implemented"
	// can also use: "echo $0"
	return false
}

// IsShellSpecialVar reports whether the character identifies a special
// shell variable such as $*.
func IsShellSpecialVar(c uint8) bool { _ = "STUB: not implemented"; return false }

func shellExec(expr, shell string) (string, error) {
	_ = "STUB: not implemented"
	// "-c" for bash,sh,zsh shell
	return "", nil
}

// special for Windows shell

// use cmd.exe, mark is "/c"

// "-Command" for powershell
