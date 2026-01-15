# Goose Skills

Skills 是可重用的指令和资源集合，用于教 Goose 执行特定任务。Skill 可以是简单的检查清单，也可以是包含领域专业知识的详细工作流程，还可以包含脚本或模板等支持文件。典型用例包括部署流程、代码审查清单和 API 集成指南。

> [!NOTE]
> 此功能需要启用内置的 Skills 扩展（默认已启用）。

当会话开始时，Goose 会将发现的所有 Skills 添加到其指令中。在会话期间，Goose 会自动加载 Skill：

- 当你的请求明确匹配某个 Skill 的用途时
- 当你明确要求使用某个 Skill 时，例如：
  - "Use the code-review skill to review this PR"
  - "Follow the new-service skill to set up the auth service"
  - "Apply the deployment skill"

你也可以询问 Goose 有哪些可用的 Skills。

> [!TIP]
> Goose Skills 与 Claude Desktop 兼容，也与其他[支持 Agent Skills 的工具](https://agentskills.io/home#adoption)兼容。

## Skill 目录位置

Skills 可以全局存储和/或按项目存储。Goose 按顺序检查以下所有目录并合并发现的内容。如果相同的 Skill 名称存在于多个目录中，后面的目录优先：

1.  `~/.claude/skills/` — 全局，与 Claude Desktop 共享
2.  `~/.config/agents/skills/` — 全局，跨 AI 编码代理通用
3.  `~/.config/goose/skills/` — 全局，Goose 专用
4.  `./.claude/skills/` — 项目级，与 Claude Desktop 共享
5.  `./.goose/skills/` — 项目级，Goose 专用
6.  `./.agents/skills/` — 项目级，跨 AI 编码代理通用

全局 Skills 用于跨项目使用的工作流程。项目级 Skills 用于特定代码库的流程。

## 创建 Skill

当你有一个涉及多个步骤、专业知识或支持文件的可重复工作流程时，可以创建 Skill。

### Skill 文件结构

每个 Skill 位于自己的目录中，包含一个 `SKILL.md` 文件：

```text
~/.config/agents/skills/
└── code-review/
    └── SKILL.md
```

`SKILL.md` 文件需要包含 `name` 和 `description` 的 YAML frontmatter，然后是 Skill 内容：

```yaml
---
name: code-review
description: Comprehensive code review checklist for pull requests
---

# Code Review Checklist

When reviewing code, check each of these areas:

## Functionality
- [ ] Code does what the PR description claims
- [ ] Edge cases are handled
- [ ] Error handling is appropriate

## Code Quality
- [ ] Follows project style guide
- [ ] No hardcoded values that should be configurable
- [ ] Functions are focused and well-named

## Testing
- [ ] New functionality has tests
- [ ] Tests are meaningful, not just for coverage
- [ ] Existing tests still pass

## Security
- [ ] No credentials or secrets in code
- [ ] User input is validated
- [ ] SQL queries are parameterized
```

### 支持文件

Skills 可以包含脚本、模板或配置文件等支持文件。将它们放在 Skill 目录中：

```text
~/.config/agents/skills/
└── api-setup/
    ├── SKILL.md
    ├── setup.sh
    └── templates/
        └── config.template.json
```

当 Goose 加载 Skill 时，它会看到支持文件并可以使用文件工具访问它们。

### 示例：带支持文件的 Skill

**SKILL.md:**

```yaml
---
name: api-setup
description: Set up API integration with configuration and helper scripts
---

# API Setup

This skill helps you set up a new API integration with our standard configuration.

## Steps

1. Run `setup.sh <api-name>` to create the integration directory
2. Copy `templates/config.template.json` to your integration directory
3. Update the config with your API credentials
4. Test the connection

## Configuration

The config template includes:
- `api_key`: Your API key (get from the provider's dashboard)
- `endpoint`: API endpoint URL
- `timeout`: Request timeout in seconds (default: 30)

## Verification

After setup, verify:
- [ ] Config file is valid JSON
- [ ] API key is set and not a placeholder
- [ ] Test connection succeeds
```

**setup.sh:**

```bash
#!/bin/bash
API_NAME=$1
mkdir -p "integrations/$API_NAME"
cp templates/config.template.json "integrations/$API_NAME/config.json"
echo "Created integration directory for $API_NAME"
echo "Edit integrations/$API_NAME/config.json with your credentials"
```

**templates/config.template.json:**

```json
{
  "api_key": "YOUR_API_KEY_HERE",
  "endpoint": "https://api.example.com/v1",
  "timeout": 30,
  "retry_attempts": 3
}
```

## 常见用例示例

### 部署工作流

```yaml
---
name: production-deploy
description: Safe deployment procedure for production environment
---

# Production Deployment

## Pre-deployment
1. Ensure all tests pass
2. Get approval from at least 2 reviewers
3. Notify #deployments channel

## Deploy
1. Create release branch from main
2. Run `npm run build:prod`
3. Deploy to staging, verify, then production
4. Monitor error rates for 30 minutes

## Rollback

If error rate exceeds 1%:
1. Revert to previous deployment
2. Notify #incidents channel
3. Create incident report
```

### 测试策略

```yaml
---
name: testing-strategy
description: Guidelines for writing effective tests in this project
---

# Testing Guidelines

## Unit Tests
- Test one thing per test
- Use descriptive test names: `test_user_creation_fails_with_invalid_email`
- Mock external dependencies

## Integration Tests
- Test API endpoints with realistic data
- Verify database state changes
- Clean up test data after each test

## Running Tests
- `npm test` — Run all tests
- `npm test:unit` — Unit tests only
- `npm test:integration` — Integration tests (requires database)
```

### API 集成指南

```yaml
---
name: square-integration
description: How to integrate with our Square account
---

# Square Integration

## Authentication
- Test key: Use `SQUARE_TEST_KEY` from `.env.test`
- Production key: In 1Password under "Square Production"

## Common Operations

### Create a customer

```javascript
const customer = await squareup.customers.create({
  email: user.email,
  metadata: { userId: user.id }
});
```

### Handle webhooks

Always verify webhook signatures. See `src/webhooks/square.js` for our handler pattern.

## Error Handling
- `card_declined`: Show user-friendly message, suggest different payment method
- `rate_limit`: Implement exponential backoff
- `invalid_request`: Log full error, likely a bug in our code
```

## 其他支持复用的 Goose 功能

- [.goosehints](/goose/docs/guides/context-engineering/using-goosehints): 最适合一般偏好、项目上下文和重复指令如 "Always use TypeScript"
- [recipes](/goose/docs/guides/recipes/session-recipes): 可共享的配置，将指令、提示和设置打包在一起

## 最佳实践

- **保持 Skills 聚焦** — 每个 Skill 对应一个工作流程或领域。如果 Skill 太长，考虑拆分。
- **清晰编写** — Skills 是给 Goose 的指令。使用清晰、直接的语言和编号步骤。
- **包含验证步骤** — 帮助 Goose 确认工作流程成功完成。