package target

import (
	"os"
	"path/filepath"
)

// droidProvider 实现 Droid (Factory AI) IDE 的 ToolProvider 接口
// Droid 是 Factory AI 的编码助手
// 使用以下目录结构：
// - 全局 Skills: ~/.factory/skills/
// - 项目级 Skills: .factory/skills/
type droidProvider struct {
	homeDir string
}

// NewDroidProvider 创建 Droid Provider 实例
func NewDroidProvider() ToolProvider {
	homeDir, _ := os.UserHomeDir()
	return &droidProvider{homeDir: homeDir}
}

// Type 返回工具类型枚举
func (d *droidProvider) Type() ToolType {
	return ToolDroid
}

// DisplayName 返回用户可见名称
func (d *droidProvider) DisplayName() string {
	return "Droid (Factory AI)"
}

// GlobalSkillsDir 返回全局 skills 扫描目录
// Droid 使用 ~/.factory/skills/
func (d *droidProvider) GlobalSkillsDir() (string, error) {
	return filepath.Join(d.homeDir, ".factory", "skills"), nil
}

// GlobalInstallDir 返回全局安装目录（与扫描目录相同）
func (d *droidProvider) GlobalInstallDir() (string, error) {
	return d.GlobalSkillsDir()
}

// LocalSkillsDir 返回项目级 skills 目录
// Droid 使用 .factory/skills/
func (d *droidProvider) LocalSkillsDir(projectRoot string) string {
	return filepath.Join(projectRoot, ".factory", "skills")
}

// Categories 返回分类子目录列表（Droid 无分类）
func (d *droidProvider) Categories() []string {
	return nil
}

// EnsureInstallDir 确保全局安装目录存在
func (d *droidProvider) EnsureInstallDir() (string, error) {
	dir, err := d.GlobalInstallDir()
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}

// EnsureLocalInstallDir 确保项目级安装目录存在
func (d *droidProvider) EnsureLocalInstallDir(projectRoot string) (string, error) {
	dir := d.LocalSkillsDir(projectRoot)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", err
	}
	return dir, nil
}
