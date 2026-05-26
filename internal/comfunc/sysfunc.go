package comfunc

// Workdir get
func Workdir() string { _ = "STUB: not implemented"; return "" }

// ExpandHome will parse first `~` as user home dir path.
func ExpandHome(pathStr string) string { _ = "STUB: not implemented"; return "" }

// ExecCmd an command and return output.
//
// Usage:
//
//	ExecCmd("ls", []string{"-al"})
func ExecCmd(binName string, args []string, workDir ...string) (string, error) {
	_ = "STUB: not implemented"
	// create a new Cmd instance
	return "", nil
}

var (
	cmdList  = []string{"cmd", "cmd.exe"}
	pwshList = []string{"powershell", "powershell.exe", "pwsh", "pwsh.exe"}
)

// ShellExec exec command by shell
// cmdLine e.g. "ls -al"
func ShellExec(cmdLine string, shells ...string) (string, error) {
	_ = "STUB: not implemented"
	// shell := "/bin/sh"
	return "", nil
}

// curShellCache value
var curShellCache string

// CurrentShell get current used shell env file.
//
// return like: "/bin/zsh" "/bin/bash". if onlyName=true, will return "zsh", "bash"
func CurrentShell(onlyName bool, fallbackShell ...string) (binPath string) {
	_ = "STUB: not implemented"
	return ""
}

// 检查父进程名称

// 适用于 Unix-like 系统

// TODO check on Windows git bash

// fix: 去除 .exe 后缀

// cache result

func checkWinCurrentShell() string {
	_ = "STUB: not implemented"
	// 在 Windows 上，可以检查 COMSPEC 环境变量
	return ""
}

// 没法检查 pwsh, 返回的还是 cmd

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
