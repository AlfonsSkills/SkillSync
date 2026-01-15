# Skills


Skills package task-specific instructions that Roo loads on-demand when your request matches the skill's purpose. Unlike custom instructions that apply to everything, skills activate only when needed—making Roo more effective at specialized tasks without cluttering the base prompt.





## Why It Matters​


Custom Instructions apply broadly across all your work. They're great for general coding standards or style preferences, but not ideal for specific workflows like "process PDF files" or "generate API documentation."


Skills solve this: Create a skill for PDF processing, and Roo only loads those instructions when you actually ask to work with PDFs. This keeps the system prompt focused and gives Roo deep expertise in specific domains without affecting unrelated tasks.


You can't package bundled assets (scripts, templates, references) with custom instructions. Skills let you store related files alongside the instructions, creating self-contained workflow packages.


## What Skills Let You Do​


- Task-Specific Expertise: Package detailed instructions for specialized workflows (data processing, documentation generation, code migration patterns)
- Bundled Resources: Include helper scripts, templates, or reference files alongside instructions
- Mode Targeting: Create skills that only activate in specific modes (e.g., code refactoring skills only in Code mode)
- Team Sharing: Version-control project skills in .roo/skills/ for consistent team workflows
- Personal Library: Build a global skills library in ~/.roo/skills/ that works across all projects
- Override Control: Project skills override global skills, mode-specific override generic


## How Skills Work​


Skills use progressive disclosure to efficiently load content only when needed:


Level 1: Discovery - Roo reads each SKILL.md file and parses its frontmatter to extract name and description. Only this metadata is stored for matching—the full content isn't held in memory until needed.


Level 2: Instructions - When your request matches a skill's description, Roo uses read_file to load the full SKILL.md instructions into context.


Level 3: Resources - The prompt tells Roo it may access bundled files (scripts, templates, references) alongside the skill. There's no separate resource manifest—Roo discovers these files on-demand when the instructions reference them.


This architecture means skills remain dormant until activated—they don't bloat your base prompt. You can install many skills, and Roo loads only what's relevant for each task.



## Creating Your First Skill​


#### 1. Choose a location​


Global skills (available in all projects):



Project skills (specific to current workspace):



#### 2. Create the skill directory and file​



#### 3. Write the SKILL.md file​


The file requires frontmatter with name and description:



Naming rules:


- The name field must exactly match the directory name (or symlink name)
- Names must be 1–64 characters, lowercase letters/numbers/hyphens only
- No leading/trailing hyphens, no consecutive hyphens (e.g., my--skill is invalid)
- Both name and description are required
- Descriptions must be 1–1024 characters (trimmed)
- The description tells Roo when to use this skill—be specific


#### 4. Test the skill​


Ask Roo something matching the description:



Roo should recognize the request matches your skill description, load the SKILL.md file, and follow its instructions.



## Directory Structure​


#### Basic structure​



#### Mode-specific skills​


Create skills that only activate in specific modes:



When to use mode-specific skills:


- Code refactoring patterns (Code mode only)
- System design templates (Architect mode only)
- Documentation standards (specific to a doc writing mode)



## Override Priority​


When skills with the same name exist in multiple locations, this priority applies (project vs global is evaluated first, then mode-specific vs generic within each source):


1. Project mode-specific (.roo/skills-code/my-skill/)
2. Project generic (.roo/skills/my-skill/)
3. Global mode-specific (~/.roo/skills-code/my-skill/)
4. Global generic (~/.roo/skills/my-skill/)


This means a project generic skill overrides a global mode-specific skill—project location takes precedence over mode specificity.


This lets you:


- Set global standards that work everywhere
- Override them per-project when needed (even with generic skills)
- Specialize skills for specific modes within each location



## Skill Discovery​


Roo automatically discovers skills:


- At startup: All skills are indexed by reading and parsing each SKILL.md
- During development: File watchers detect changes to SKILL.md files
- Mode filtering: Only skills relevant to the current mode are available


You don't need to register or configure skills—just create the directory structure.


Custom System Prompts Override SkillsIf you have a file-based custom system prompt (.roo/system-prompt-{mode-slug}), it replaces the standard system prompt entirely—including the skills section. Skills won't be available when a custom system prompt is active.


#### Symlink support​


Skills support symbolic links for sharing skill libraries across projects:



The skill name comes from the symlink name (or directory name if not symlinked). The frontmatter name field must match this name exactly—you can't create aliases with different names pointing to the same skill.



## Troubleshooting​


#### Skill isn't loading​


Symptom: Roo doesn't use your skill even when you request something matching the description.


Causes & fixes:


1. 
Name mismatch: The frontmatter name field must exactly match the directory name
# ✗ Wrong - directory is "pdf-processing"---name: pdf_processing---# ✓ Correct---name: pdf-processing---

2. 
Missing required fields: Both name and description are required in frontmatter

3. 
Wrong mode: If the skill is in skills-code/ but you're in Architect mode, it won't load. Move to skills/ for all modes or create mode-specific variants.

4. 
Description too vague: Make descriptions specific so Roo can match them to requests
# ✗ Vaguedescription: Handle files# ✓ Specificdescription: Extract text and tables from PDF files using Python libraries



#### Skill loads but doesn't help​


Symptom: Roo reads the skill but doesn't follow instructions.


Cause: Instructions may be too general or missing critical details.


Fix: Make instructions actionable:


- Include specific function names or library choices
- Provide code templates
- List common edge cases and how to handle them
- Add troubleshooting guidance for the specific task


#### Multiple skills conflict​


Symptom: Unclear which skill Roo will use when multiple might match.


Cause: Overlapping descriptions or mode configurations.


Prevention:


- Make descriptions distinct and specific
- Use mode-specific directories to separate concerns
- Rely on override priority—project skills override global


#### Can't share skills with team​


Symptom: Want team members to use the same skills.


Solution: Place skills in .roo/skills/ within your project and commit to version control. Each team member gets the same skills automatically.



## Skills vs Custom Instructions vs Slash Commands​


FeatureSkillsCustom InstructionsSlash Commands**When loaded**On-demand (when request matches)Always (part of base prompt)On-demand (when invoked)**Best for**Task-specific workflowsGeneral coding standardsRetrieving pre-written content**Can bundle files**YesNoNo**Mode targeting**Yes (`skills-{mode}` directories)Yes (`rules-{mode}` directories)No**Override priority**Project > Global, Mode > GenericProject > GlobalProject > Global**Format**SKILL.md with frontmatterAny text fileJSON metadata + content**Discovery**Automatic (directory scan)Automatic (directory scan)Automatic (directory scan)
When to use each:


- Skills: "Generate API docs following OpenAPI spec" → Detailed OpenAPI processing instructions load only when needed
- Custom Instructions: "Always use TypeScript strict mode" → Applies to all TypeScript work
- Slash Commands: /init → Returns standardized project setup instructions



## Skill Specification​


Roo Code skills follow the Agent Skills format for skill packaging and metadata. Skills are instruction packages with optional bundled files—they don't register new executable tools.


Required conventions:


- The frontmatter name must exactly match the directory (or symlink) name
- Both name and description fields are required in frontmatter
- Names: 1–64 chars, lowercase alphanumeric + hyphens, no leading/trailing/consecutive hyphens
- Descriptions: 1–1024 chars (trimmed)


#### Roo-specific enhancements​


Roo Code adds mode-specific targeting beyond the base format:


- Standard locations: .roo/skills/ and ~/.roo/skills/
- Mode-specific directories: skills-{mode}/ (e.g., skills-code/, skills-architect/) enable mode targeting



## See Also​


- Custom Instructions - Set general rules that apply to all work
- Slash Commands - Execute commands that return content
- Custom Modes - Create specialized modes with specific tool access