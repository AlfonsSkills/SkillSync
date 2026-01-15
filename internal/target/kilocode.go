package target

import (
	"os"
	"path/filepath"
)

// kiloCodeProvider 实现 Kilo Code VSCode 扩展的 ToolProvider 接口
// Kilo Code 是一个支持 Agent Skills 的 VSCode AI 扩展
// 使用以下目录结构：
// - 全局 Skills: ~/.kilocode/skills/
// - 项目级 Skills: .kilocode/skills/
// 注意：Kilo Code 还支持 mode-specific skills（如 skills-code/, skills-architect/），
// 但 SkillSync 目前仅使用主目录 skills/
type kiloCodeProvider struct {
	homeDir string
}

// NewKiloCodeProvider 创建 Kilo Code Provider 实例
func NewKiloCodeProvider() ToolProvider {
	homeDir, _ := os.UserHomeDir()
	return &kiloCodeProvider{homeDir: homeDir}
}

// Type 返回工具类型枚举
func (k *kiloCodeProvider) Type() ToolType {
	return ToolKiloCode
}

// DisplayName 返回用户可见名称
func (k *kiloCodeProvider) DisplayName() string {
	return "Kilo Code"
}

// GlobalSkillsDir 返回全局 skills 扫描目录
// Kilo Code 使用 ~/.kilocode/skills/
func (k *kiloCodeProvider) GlobalSkillsDir() (string, error) {
	return filepath.Join(k.homeDir, ".kilocode", "skills"), nil
}

// GlobalInstallDir 返回全局安装目录（与扫描目录相同）
func (k *kiloCodeProvider) GlobalInstallDir() (string, error) {
	return k.GlobalSkillsDir()
}

// LocalSkillsDir 返回项目级 skills 目录
// Kilo Code 使用 .kilocode/skills/
func (k *kiloCodeProvider) LocalSkillsDir(projectRoot string) string {
	return filepath.Join(projectRoot, ".kilocode", "skills")
}

// Categories 返回分类子目录列表（暂不支持 mode-specific）
func (k *kiloCodeProvider) Categories() []string {
	return nil
}

// EnsureInstallDir 确保全局安装目录存在
func (k *kiloCodeProvider) EnsureInstallDir() (string, error) {
	dir, err := k.GlobalInstallDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

// EnsureLocalInstallDir 确保项目级安装目录存在
func (k *kiloCodeProvider) EnsureLocalInstallDir(projectRoot string) (string, error) {
	dir := k.LocalSkillsDir(projectRoot)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}
