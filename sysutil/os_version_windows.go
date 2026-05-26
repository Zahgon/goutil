package sysutil

// OSVersionInfo 结构体用于存储操作系统版本信息
//
// NOTE: Windows 10 和 Windows 11 在主版本号和次版本号上是相同的，因此需要通过构建号（win11: DwBuildNumber>=22000）来进一步区分
type OSVersionInfo struct {
	// 主版本号
	MajorVersion uint16
	// 次版本号
	MinorVersion uint16
	// 构建号
	BuildNumber uint32
	// 修订号
	RevisionNumber uint32
}

// cache: 全局变量
var stdOv = VersionInfoBySys()

// OsVersion Get operating system version information
func OsVersion() *OSVersionInfo {
	_ = "STUB: not implemented"

	// VersionInfoBySys Get Windows system version information by sys/windows
	return nil
}

func VersionInfoBySys() *OSVersionInfo {
	_ = "STUB: not implemented"
	// Get the Windows Version and Build Number
	return nil
}

// OsVersionByParse Get Windows system version information by parse string
//
//	cmdOut eg: "Microsoft Windows [Version 10.0.22631.4391]"
func OsVersionByParse(cmdOut string) (*OSVersionInfo, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// OsVersionByVerCmd Get Windows system version information
//
// 还可用使用dll获取：
//
//	通过 GetVersion, GetVersionEx 函数获取的信息不准确. win11获取到 6.2.9200, 实际是 10.0.22631
func OsVersionByVerCmd() (*OSVersionInfo, error) {
	_ = "STUB: not implemented"
	// Windows cmd 执行 ver 命令
	return nil, nil
}

// IsLtWindows7 判断是否小于 Windows 7
func (ov *OSVersionInfo) IsLtWindows7() bool { _ = "STUB: not implemented"; return false }

// IsWindows7 判断是否为 Windows 7
func (ov *OSVersionInfo) IsWindows7() bool { _ = "STUB: not implemented"; return false }

// IsWindows8 判断是否为 Windows 8/8.1
func (ov *OSVersionInfo) IsWindows8() bool { _ = "STUB: not implemented"; return false }

// IsWindows10 判断是否为 Windows 10
func (ov *OSVersionInfo) IsWindows10() bool { _ = "STUB: not implemented"; return false }

// IsWindows11 判断是否为 Windows 11
func (ov *OSVersionInfo) IsWindows11() bool { _ = "STUB: not implemented"; return false }

// Name 获取 Windows 通用的版本名称. eg: xp, win7, win8, win10, win11, unknown
func (ov *OSVersionInfo) Name() string { _ = "STUB: not implemented"; return "" }

// win server 2003

// String format
func (ov *OSVersionInfo) String() string { _ = "STUB: not implemented"; return "" }

func parseOsVersionString(out string) (*OSVersionInfo, error) {
	_ = "STUB: not implemented"
	// out eg: Microsoft Windows [Version 10.0.22631.4391] => 10.0.22631.4391
	// 部分系统会输出中文 eg: Microsoft Windows [版本 10.0.22631.4391]
	return nil, nil
}

// get like: 10.0.22631.4391 or 10.0.22631
