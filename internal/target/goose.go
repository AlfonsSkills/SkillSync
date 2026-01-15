package target

import (
	"os"
	"path/filepath"
)

// gooseProvider 实现 Goose AI 的 ToolProvider 接口
// Goose 使用以下目录结构（按优先级）：
// 全局目录：
// - ~/.claude/skills/ — 与 Claude Desktop 共享
// - ~/.config/agents/skills/ — 跨 AI 编码代理的通用目录
// - ~/.config/goose/skills/ — Goose 专用
// 项目目录：
// - ./.claude/skills/ — 与 Claude Desktop 共享
// - ./.goose/skills/ — Goose 专用
// - ./.agents/skills/ — 跨 AI 编码代理的通用目录
// SkillSync 使用 Goose 专用目录：~/.config/goose/skills/ 和 .goose/skills/
type gooseProvider struct {
	homeDir string
}

// NewGooseProvider 创建 Goose AI Provider 实例
func NewGooseProvider() ToolProvider {
	homeDir, _ := os.UserHomeDir()
	return &gooseProvider{homeDir: homeDir}
}

// Type 返回工具类型枚举
func (g *gooseProvider) Type() ToolType {
	return ToolGoose
}

// DisplayName 返回用户可见名称
func (g *gooseProvider) DisplayName() string {
	return "Goose AI"
}

// GlobalSkillsDir 返回全局 skills 扫描目录
// Goose 使用 ~/.config/goose/skills/
func (g *gooseProvider) GlobalSkillsDir() (string, error) {
	return filepath.Join(g.homeDir, ".config", "goose", "skills"), nil
}

// GlobalInstallDir 返回全局安装目录（与扫描目录相同）
func (g *gooseProvider) GlobalInstallDir() (string, error) {
	return g.GlobalSkillsDir()
}

// LocalSkillsDir 返回项目级 skills 目录
// Goose 使用 .goose/skills/
func (g *gooseProvider) LocalSkillsDir(projectRoot string) string {
	return filepath.Join(projectRoot, ".goose", "skills")
}

// Categories 返回分类子目录列表（Goose 无分类）
func (g *gooseProvider) Categories() []string {
	return nil
}

// EnsureInstallDir 确保全局安装目录存在
func (g *gooseProvider) EnsureInstallDir() (string, error) {
	dir, err := g.GlobalInstallDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

// EnsureLocalInstallDir 确保项目级安装目录存在
func (g *gooseProvider) EnsureLocalInstallDir(projectRoot string) (string, error) {
	dir := g.LocalSkillsDir(projectRoot)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}
