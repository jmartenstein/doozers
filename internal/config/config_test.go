package config

import (
        "os"
        "testing"
)

func TestLoad(t *testing.T) {
        content := `
personas:
  - name: Cary
    aliases: ["cary", "coder"]
    script: "script content"
    agent: "gemini"
    use_git_worktree: true
  - name: Archie
    aliases: ["archie", "architect"]
    script: "planning script"
    use_git_worktree: false
default_agent: "gemini"
`
        tmpFile, err := os.CreateTemp("", "config_test.yaml")
        if err != nil {
                t.Fatal(err)
        }
        defer os.Remove(tmpFile.Name())

        if err := os.WriteFile(tmpFile.Name(), []byte(content), 0644); err != nil {
                t.Fatal(err)
        }

        cfg, err := Load(tmpFile.Name())
        if err != nil {
                t.Fatal(err)
        }
        if cfg == nil {
                t.Fatal("config is nil")
        }
        if len(cfg.Personas) != 2 {
                t.Fatalf("expected 2 personas, got %d", len(cfg.Personas))
        }

        cary, ok := cfg.GetPersona("cary")
        if !ok {
                t.Fatal("cary persona not found")
        }
        if cary.Name != "Cary" {
                t.Fatalf("expected Cary, got %s", cary.Name)
        }
        if !cary.UseGitWorktree {
                t.Fatal("expected UseGitWorktree true for cary")
        }

        archie, ok := cfg.GetPersona("archie")
        if !ok {
                t.Fatal("archie persona not found")
        }
        if archie.Name != "Archie" {
                t.Fatalf("expected Archie, got %s", archie.Name)
        }
        if archie.UseGitWorktree {
                t.Fatal("expected UseGitWorktree false for archie")
        }
}
