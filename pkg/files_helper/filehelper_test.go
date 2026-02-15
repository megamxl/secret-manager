package files_helper

import (
	"bytes"
	"os"
	"path/filepath"
	"secret-manager/internal/config"
	"testing"
)

func TestCreateFile_TableDriven(t *testing.T) {
	root := t.TempDir() // Isolated root for this test suite
	// Create a temporary config file inside the temp root

	configPath := "../../config-test.yml"
	config.Initialize(configPath)
	config.Get().Agent.TrustedRoots = []string{
		root,
		"/private" + root,
	}

	// Define the test schema
	tests := map[string]struct {
		filename string
		content  string
		wantErr  bool
	}{
		"Success: Valid JSON": {
			filename: "valid.json",
			content:  `{"status": "ok"}`,
			wantErr:  false,
		},
		"Success: Valid .env": {
			filename: "app.env",
			content:  "DB_PASSWORD=highly_secure",
			wantErr:  false,
		},
		"Failure: Malformed JSON": {
			filename: "broken.json",
			content:  `{"unclosed": "json"`, // Missing bracket
			wantErr:  true,
		},
		"Failure: Path Traversal": {
			filename: "../outside.txt",
			content:  "malicious",
			wantErr:  true,
		},
		"Failure: Invalid XML": {
			filename: "config.xml",
			content:  "<root><unclosed>",
			wantErr:  true,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			targetPath := filepath.Join(root, tc.filename)
			buf := bytes.NewBufferString(tc.content)

			err := CreateFileAtomic(targetPath, buf)

			if (err != nil) != tc.wantErr {
				t.Fatalf("CreateFileAtomic() error = %v, wantErr %v", err, tc.wantErr)
			}

			if !tc.wantErr {
				// Verify the file was actually created and has correct perms
				info, _ := os.Stat(targetPath)
				if info.Mode().Perm() != 0600 {
					t.Errorf("File perms are %v, want 0600", info.Mode().Perm())
				}
			}
		})
	}
}

func TestCreateFile_AtomicIntegrity(t *testing.T) {
	root := t.TempDir()
	target := filepath.Join(root, "config.json")

	initialContent := `{"version": 1}`
	os.WriteFile(target, []byte(initialContent), 0600)

	// 2. Attempt to overwrite with INVALID content
	invalidContent := bytes.NewBufferString(`{"version": 2, broken...`)
	err := CreateFileAtomic(target, invalidContent)

	if err == nil {
		t.Fatal("Expected error for invalid content")
	}

	// 3. Verify original file is UNTOUCHED (Reliability REL-07)
	currentContent, _ := os.ReadFile(target)
	if string(currentContent) != initialContent {
		t.Errorf("Atomic failure! File was corrupted despite validation error")
	}
}
