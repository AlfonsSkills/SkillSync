package target

// vscodeProvider 实现 VSCode (GitHub Copilot) 的 ToolProvider 接口
// VSCode 使用与 GitHub Copilot 相同的目录结构，作为 copilot 的别名
// - 全局 Skills: ~/.copilot/skills/
// - 项目级 Skills: .github/skills/
type vscodeProvider struct {
	// 内嵌 copilotProvider，复用相同的目录配置
	*copilotProvider
}

// NewVSCodeProvider 创建 VSCode Provider 实例
// VSCode 与 Copilot 共享相同的目录配置
func NewVSCodeProvider() ToolProvider {
	base := NewCopilotProvider().(*copilotProvider)
	return &vscodeProvider{copilotProvider: base}
}

// Type 返回工具类型枚举（覆盖基类方法）
func (v *vscodeProvider) Type() ToolType {
	return ToolVSCode
}

// DisplayName 返回用户可见名称（覆盖基类方法）
func (v *vscodeProvider) DisplayName() string {
	return "VSCode (Copilot)"
}
