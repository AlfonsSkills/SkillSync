package target

import (
	"os"
	"path/filepath"
)

// clineProvider 实现 Cline IDE 的 ToolProvider 接口
// Cline 使用以下目录结构：
// - 全局 Skills: ~/.cline/skills/
// - 项目级 Skills: .cline/skills/
// 注意：Cline 也支持 .clinerules/skills/ 和 .claude/skills/（兼容模式），
// 但 SkillSync 仅使用主目录 .cline/skills/
type clineProvider struct {
	homeDir string
}

// NewClineProvider 创建 Cline IDE Provider 实例
func NewClineProvider() ToolProvider {
	homeDir, _ := os.UserHomeDir()
	return &clineProvider{homeDir: homeDir}
}

// Type 返回工具类型枚举
func (c *clineProvider) Type() ToolType {
	return ToolCline
}

// DisplayName 返回用户可见名称
func (c *clineProvider) DisplayName() string {
	return "Cline IDE"
}

// GlobalSkillsDir 返回全局 skills 扫描目录
// Cline 使用 ~/.cline/skills/
func (c *clineProvider) GlobalSkillsDir() (string, error) {
	return filepath.Join(c.homeDir, ".cline", "skills"), nil
}

// GlobalInstallDir 返回全局安装目录（与扫描目录相同）
func (c *clineProvider) GlobalInstallDir() (string, error) {
	return c.GlobalSkillsDir()
}

// LocalSkillsDir 返回项目级 skills 目录
// Cline 使用 .cline/skills/
func (c *clineProvider) LocalSkillsDir(projectRoot string) string {
	return filepath.Join(projectRoot, ".cline", "skills")
}

// Categories 返回分类子目录列表（Cline 无分类）
func (c *clineProvider) Categories() []string {
	return nil
}

// EnsureInstallDir 确保全局安装目录存在
func (c *clineProvider) EnsureInstallDir() (string, error) {
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
func (c *clineProvider) EnsureLocalInstallDir(projectRoot string) (string, error) {
	dir := c.LocalSkillsDir(projectRoot)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}
