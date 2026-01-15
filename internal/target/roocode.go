package target

import (
	"os"
	"path/filepath"
)

// rooCodeProvider 实现 Roo Code VSCode 扩展的 ToolProvider 接口
// Roo Code 是一个支持 Agent Skills 的 VSCode AI 扩展
// 使用以下目录结构：
// - 全局 Skills: ~/.roo/skills/
// - 项目级 Skills: .roo/skills/
// 注意：Roo Code 还支持 mode-specific skills（如 skills-code/, skills-architect/），
// 但 SkillSync 目前仅使用主目录 skills/
type rooCodeProvider struct {
	homeDir string
}

// NewRooCodeProvider 创建 Roo Code Provider 实例
func NewRooCodeProvider() ToolProvider {
	homeDir, _ := os.UserHomeDir()
	return &rooCodeProvider{homeDir: homeDir}
}

// Type 返回工具类型枚举
func (r *rooCodeProvider) Type() ToolType {
	return ToolRooCode
}

// DisplayName 返回用户可见名称
func (r *rooCodeProvider) DisplayName() string {
	return "Roo Code"
}

// GlobalSkillsDir 返回全局 skills 扫描目录
// Roo Code 使用 ~/.roo/skills/
func (r *rooCodeProvider) GlobalSkillsDir() (string, error) {
	return filepath.Join(r.homeDir, ".roo", "skills"), nil
}

// GlobalInstallDir 返回全局安装目录（与扫描目录相同）
func (r *rooCodeProvider) GlobalInstallDir() (string, error) {
	return r.GlobalSkillsDir()
}

// LocalSkillsDir 返回项目级 skills 目录
// Roo Code 使用 .roo/skills/
func (r *rooCodeProvider) LocalSkillsDir(projectRoot string) string {
	return filepath.Join(projectRoot, ".roo", "skills")
}

// Categories 返回分类子目录列表（暂不支持 mode-specific）
func (r *rooCodeProvider) Categories() []string {
	return nil
}

// EnsureInstallDir 确保全局安装目录存在
func (r *rooCodeProvider) EnsureInstallDir() (string, error) {
	dir, err := r.GlobalInstallDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

// EnsureLocalInstallDir 确保项目级安装目录存在
func (r *rooCodeProvider) EnsureLocalInstallDir(projectRoot string) (string, error) {
	dir := r.LocalSkillsDir(projectRoot)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}
