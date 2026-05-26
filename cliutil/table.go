package cliutil

// SimpleTable 使用标准库 text/tabwriter 实现简单的表格
func SimpleTable(cols []string, rows [][]any) { _ = "STUB: not implemented"; return }

// 参数说明：输出流, 最小单元宽度, 制表符宽度, 填充空格数, 填充字节, 标志位(Debug加"|")

// 渲染表头

// stdio.Fprintln(w, strings.Repeat("-", len(header) + len(cols)*2))

// 渲染行

// TableStyle 定义表格的样式
type TableStyle struct {
	// 边框字符
	BorderTop     string // 顶部边框
	BorderBottom  string // 底部边框
	BorderLeft    string // 左边框
	BorderRight   string // 右边框
	BorderJoin    string // 内部交叉点
	BorderRowSep  string // 行分隔符
	HeaderSep     string // 表头分隔符
	ColumnSep     string // 列分隔符
	ColumnPadding int    // 列内边距

	// 默认的对齐方式 0: 左对齐, 1: 居中, 2: 右对齐
	//  - 默认左对齐
	DefaultAlign int
	// 定义每列的对齐方式 (0: 左对齐, 1: 居中, 2: 右对齐)
	Align []int

	// 颜色设置 (需要终端支持 ANSI 颜色)
	HeaderColor string
	RowColor    string
	BorderColor string
}

// 默认的表格样式实例
var defaultStyle = DefaultStyle()

// MinimalStyle 极简风格 (无边框)
var MinimalStyle = &TableStyle{
	BorderTop:     "",
	BorderBottom:  "",
	BorderLeft:    "",
	BorderRight:   "",
	BorderJoin:    " ",
	HeaderSep:     "-",
	ColumnSep:     " ",
	BorderRowSep:  "",
	ColumnPadding: 1,
	HeaderColor:   "ylw", // 黄色
}

// DefaultStyle 返回默认的表格样式
func DefaultStyle() *TableStyle { _ = "STUB: not implemented"; return nil }

// 使用横线作为顶部边框
// 使用横线作为底部边框
// 使用竖线作为左边框
// 使用竖线作为右边框
// 使用加号作为交叉点
// 使用横线作为行分隔
// 使用等号作为表头分隔

// 默认1个空格的边距
// 粗体
// 默认无颜色
// 默认无颜色

// ShowTable CLI渲染显示显示表格
func ShowTable(cols []string, rows [][]any, options ...*TableStyle) {
	_ = "STUB: not implemented"
	return
}

// FormatTable CLI渲染显示显示表格
func FormatTable(cols []string, rows [][]any, options ...*TableStyle) string {
	_ = "STUB: not implemented"
	return ""
}

type TableBuilder struct {
	*TableStyle
	// sb strings.Builder
	// context data
	colWidths  []int
	totalWidth int
}

func NewTableBuilder(options ...*TableStyle) *TableBuilder { _ = "STUB: not implemented"; return nil }

// 处理选项

func (t *TableBuilder) prepare(cols []string, rows [][]any) {
	_ = "STUB: not implemented"
	// 计算每列的最大宽度 (包含padding)
	return
}

// 计算总宽度

// 内容 + 左右padding + 分隔符

// 最后的右边框

// Format 格式化构建CLI表格
func (t *TableBuilder) Format(cols []string, rows [][]any) string {
	_ = "STUB: not implemented"
	return ""
}

// 1. 顶部边框

// 2. 表头

// 3. 表头分隔符

// 4. 数据行

// 底部边框

// 构建辅助函数：构建分隔线
func (t *TableBuilder) buildSep(char string) string { _ = "STUB: not implemented"; return "" }

// 构建辅助函数：构建一行
func (t *TableBuilder) buildRow(cells []string, isHeader bool) string {
	_ = "STUB: not implemented"
	return ""
}

// sb.WriteString(t.BorderColor)

// 应用对齐

// 默认左对齐

// 居中

// 右对齐

// 左对齐

// sb.WriteString(cell)
