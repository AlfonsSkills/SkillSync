# SkillSync

[![CI](https://github.com/AlfonsSkills/SkillSync/actions/workflows/ci.yml/badge.svg)](https://github.com/AlfonsSkills/SkillSync/actions/workflows/ci.yml)
[![Release](https://github.com/AlfonsSkills/SkillSync/actions/workflows/release.yml/badge.svg)](https://github.com/AlfonsSkills/SkillSync/releases)
[![GitHub release](https://img.shields.io/github/v/release/AlfonsSkills/SkillSync)](https://github.com/AlfonsSkills/SkillSync/releases/latest)
[![Go Report Card](https://goreportcard.com/badge/github.com/AlfonsSkills/SkillSync)](https://goreportcard.com/report/github.com/AlfonsSkills/SkillSync)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[中文文档](README_CN.md)

**Sync skills from Git repositories to 14+ AI coding tools with one command.**

## Preview

```
📦 Installed Skills:

  Gemini CLI (1):
  📁 ~/.gemini/skills
    ✓ devops

  Claude Code (2):
  📁 ~/.claude/skills
    ✓ devops
    [project:MyProject]
      ✓ docx

  Codex CLI (5):
  📁 ~/.codex/skills
    ✓ gh-address-comments
    [public]
      ✓ devops
    [.system]
      ✓ skill-creator
```

## Installation

```bash
curl -fsSL https://raw.githubusercontent.com/AlfonsSkills/SkillSync/main/install.sh | bash
```

### From Release

Download the latest binary from [Releases](https://github.com/AlfonsSkills/SkillSync/releases).

### From Source

```bash
git clone https://github.com/AlfonsSkills/SkillSync.git
cd SkillSync
make build
# Binary will be at ./build/skillsync
```

## Quick Start

```bash
# Install skills from a repository
skillsync install anthropics/skills

# Install to specific tool
skillsync install anthropics/skills -t gemini

# Install to project-local directory
skillsync install anthropics/skills --local

# List installed skills
skillsync list

# Remove a skill
skillsync remove skill-name

# Install to multiple tools
skillsync install AlfonsSkills/skills -t claude,codex,gemini

# Install from GitLab or other platforms
skillsync install https://gitlab.com/user/skill-repo.git

# List skills for specific tool
skillsync list --target gemini

# Remove from specific tool
skillsync remove skill-name --target claude

# Remove from project directories only
skillsync remove skill-name --local
```

## Supported Tools

SkillSync supports **14 AI coding tools** across terminal and IDE environments.

### Terminal Tools

| Tool | Skills Directory | Flag |
|------|-----------------|------|
| Gemini CLI | `~/.gemini/skills/` | `-t gemini` |
| Claude Code | `~/.claude/skills/` | `-t claude` |
| Codex CLI | `~/.codex/skills/public/` | `-t codex` |
| OpenCode | `~/.config/opencode/skill/` | `-t opencode` |
| Goose AI | `~/.config/goose/skills/` | `-t goose` |
| Crush | `~/.config/crush/skills/` | `-t crush` |

### IDE Tools

| Tool | Skills Directory | Flag |
|------|-----------------|------|
| Antigravity IDE | `~/.gemini/antigravity/skills/` | `-t antigravity` |
| Copilot | `~/.copilot/skills/` | `-t copilot` |
| Cursor | `~/.cursor/skills/` | `-t cursor` |
| Cline IDE | `~/.cline/skills/` | `-t cline` |
| Droid (Factory AI) | `~/.factory/skills/` | `-t droid` |
| Kilo Code | `~/.kilocode/skills/` | `-t kilocode` |
| Roo Code | `~/.roo/skills/` | `-t roocode` |
| VSCode (Copilot) | `~/.copilot/skills/` | `-t vscode` |

## Skill Format

A valid skill repository must contain a `SKILL.md` file:

```
my-skill/
├── SKILL.md          # Required: Skill definition
├── references/       # Optional: Reference docs
└── scripts/          # Optional: Scripts
```

## License

MIT License - see [LICENSE](LICENSE) for details.
