package target

import (
	"os"
	"path/filepath"
)

// crushProvider 实现 Crush 终端工具的 ToolProvider 接口
// Crush 是 Charmbracelet 出品的终端 AI 编程助手
// 使用以下目录结构：
// - 全局 Skills: ~/.config/crush/skills/
// - 项目级 Skills: .crush/skills/
type crushProvider struct {
	homeDir string
}

// NewCrushProvider 创建 Crush Provider 实例
func NewCrushProvider() ToolProvider {
	homeDir, _ := os.UserHomeDir()
	return &crushProvider{homeDir: homeDir}
}

// Type 返回工具类型枚举
func (c *crushProvider) Type() ToolType {
	return ToolCrush
}

// DisplayName 返回用户可见名称
func (c *crushProvider) DisplayName() string {
	return "Crush"
}

// GlobalSkillsDir 返回全局 skills 扫描目录
// Crush 使用 ~/.config/crush/skills/
func (c *crushProvider) GlobalSkillsDir() (string, error) {
	return filepath.Join(c.homeDir, ".config", "crush", "skills"), nil
}

// GlobalInstallDir 返回全局安装目录（与扫描目录相同）
func (c *crushProvider) GlobalInstallDir() (string, error) {
	return c.GlobalSkillsDir()
}

// LocalSkillsDir 返回项目级 skills 目录
// Crush 使用 .crush/skills/
func (c *crushProvider) LocalSkillsDir(projectRoot string) string {
	return filepath.Join(projectRoot, ".crush", "skills")
}

// Categories 返回分类子目录列表（Crush 无分类）
func (c *crushProvider) Categories() []string {
	return nil
}

// EnsureInstallDir 确保全局安装目录存在
func (c *crushProvider) EnsureInstallDir() (string, error) {
	dir, err := c.GlobalInstallDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

// EnsureLocalInstallDir 确保项目级安装目录存在
func (c *crushProvider) EnsureLocalInstallDir(projectRoot string) (string, error) {
	dir := c.LocalSkillsDir(projectRoot)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}
