package tools

import "testing"

// 由于相对路径的问题，测试时需要修改源代码中的相对路径
// 具体为将"./"修改为"../../"
func TestExtractCover(t *testing.T) {
	if err := ExtractCover(125794582365478165, 597814563214596347); err != nil {
		t.Error(err)
	}
}
