# SkillSync

[![CI](https://github.com/AlfonsSkills/SkillSync/actions/workflows/ci.yml/badge.svg)](https://github.com/AlfonsSkills/SkillSync/actions/workflows/ci.yml)
[![Release](https://github.com/AlfonsSkills/SkillSync/actions/workflows/release.yml/badge.svg)](https://github.com/AlfonsSkills/SkillSync/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/AlfonsSkills/SkillSync)](https://goreportcard.com/report/github.com/AlfonsSkills/SkillSync)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

[中文文档](README_CN.md)

Sync skills from Git repositories to local AI coding tools (Gemini CLI / Claude Code / Codex CLI).

## Features

- 📦 **Install skills** from any Git repository (GitHub by default)
- 📋 **List skills** installed locally across all tools
- 🗑️ **Remove skills** from specific or all tools
- 🎯 **Target selection** - choose which tools to sync to

## Installation

### Quick Install (Recommended)

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

## Usage

```bash
# Install skills from monorepo (interactive selection)
skillsync install anthropics/skills

# Install to specific tool
skillsync install anthropics/skills --target gemini
skillsync install AlfonsSkills/skills -t claude,codex

# Install to project-local directories (.gemini/skills, .claude/skills, .codex/skills)
skillsync install anthropics/skills --local

# Install from other Git platforms
skillsync install https://gitlab.com/user/skill-repo.git

# List installed skills (global + project-local)
skillsync list
skillsync list --target gemini

# Example output:
#   📦 Installed Skills:
#
#     Gemini CLI (1):
#     📁 ~/.gemini/skills
#       ✓ devops
#
#     Claude Code (2):
#     📁 ~/.claude/skills
#       ✓ devops
#       [project:MyProject]
#         ✓ docx
#
#     Codex CLI (5):
#     📁 ~/.codex/skills
#       ✓ gh-address-comments
#       [public]
#         ✓ devops
#       [.system]
#         ✓ skill-creator

# Remove skill
skillsync remove skill-name
skillsync remove skill-name --target claude
skillsync remove skill-name --local  # Remove from project directories
```

## Supported Tools

| Tool | Skills Directory | Flag |
|------|-----------------|------|
| Gemini CLI | `~/.gemini/skills/` | `-t gemini` |
| Claude Code | `~/.claude/skills/` | `-t claude` |
| Codex CLI | `~/.codex/skills/public/` | `-t codex` |
| Antigravity IDE | `~/.gemini/antigravity/skills/` | `-t antigravity` |
| Copilot / VSCode | `~/.copilot/skills/` | `-t copilot` |
| Cursor | `~/.cursor/skills/` | `-t cursor` |
| OpenCode | `~/.config/opencode/skill/` | `-t opencode` |

## Skill Repository Requirements

A valid skill repository must contain a `SKILL.md` file:

```
my-skill/
├── SKILL.md          # Required: Skill definition
├── references/       # Optional: Reference docs
│   └── *.md
└── scripts/          # Optional: Scripts
    └── *.sh
```

## License

MIT License - see [LICENSE](LICENSE) for details.
