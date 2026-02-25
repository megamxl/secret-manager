package files_helper

import (
	"bytes"
	"encoding/xml"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"secret-manager/internal/config"
	"strings"

	"sigs.k8s.io/yaml"
)

func CreateFileAtomic(path string, evaluatedTemplate *bytes.Buffer) error {
	// 1. Security Check (CWE-35 Mitigation)
	cleanedPath, err := verifyPath(path)
	if err != nil {
		return fmt.Errorf("path verification failed: %w", err)
	}

	// 2. Pre-flight Validation (Check if content is valid JSON/YAML/etc before writing)
	if err := validateContentFormat(cleanedPath, evaluatedTemplate.Bytes()); err != nil {
		return fmt.Errorf("pre-flight validation failed: %w", err)
	}

	// 3. Atomic Write Strategy: Create a temporary file in the same directory
	dir := filepath.Dir(cleanedPath)
	tmpFile, err := os.CreateTemp(dir, ".secret-tmp-*")
	if err != nil {
		return fmt.Errorf("failed to create temp file: %w", err)
	}
	tmpPath := tmpFile.Name()
	defer os.Remove(tmpPath)

	// 4. Secure Permissions: Set to 0600 (Owner Read/Write Only) MA thesis SEC-03
	if err := tmpFile.Chmod(0600); err != nil {
		tmpFile.Close()
		return err
	}

	// 5. Write and Sync (REL-06: Ensure bits are physically on disk)
	if _, err := tmpFile.Write(evaluatedTemplate.Bytes()); err != nil {
		tmpFile.Close()
		return fmt.Errorf("failed writing to temp file: %w", err)
	}

	// Sync ensures the OS flushes buffers to disk before we swap
	if err := tmpFile.Sync(); err != nil {
		tmpFile.Close()
		return err
	}
	tmpFile.Close()

	// 6. Atomic Swap (FC-08): os.Rename is an atomic operation in Linux
	if err := os.Rename(tmpPath, cleanedPath); err != nil {
		return fmt.Errorf("atomic swap failed: %w", err)
	}

	return nil
}

// validateContentFormat provides the "Pre-flight" check logic
func validateContentFormat(path string, data []byte) error {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".xml":
		return xml.Unmarshal(data, &struct{}{})
	case ".yaml", ".yml", ".json":
		return yaml.Unmarshal(data, &struct{}{})
	case ".env":
		if !bytes.Contains(data, []byte("=")) {
			return errors.New("invalid .env format: missing key-value pair")
		}
	}
	return nil
}

func DeleteFile(path string) error {
	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}

// For master thesis https://cwe.mitre.org/data/definitions/35.html
func verifyPath(path string) (string, error) {

	c := filepath.Clean(path)
	dir := filepath.Dir(c)
	filename := filepath.Base(c)

	// Evaluate the Directory File can be new ...
	resolvedDir, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return "", fmt.Errorf("directory does not exist or is unreachable: %w", err)
	}

	finalPath := filepath.Join(resolvedDir, filename)

	err = inTrustedRoot(finalPath)
	if err != nil {
		return "", errors.New("security violation: path is outside of trusted root")
	}

	return finalPath, nil
}

func inTrustedRoot(path string) error {
	path = filepath.Clean(path)

	for _, root := range config.Get().Agent.TrustedRoots {
		root = filepath.Clean(root)

		// 2. It must START with the root (not just contain it)
		if strings.HasPrefix(path, root) {

			// 3. Boundary check:
			// Either the path IS the root, or it's followed by a separator
			// This prevents the "Sibling Attack" (/var/lib/tc matching /var/lib/tc-extra)
			if len(path) == len(root) || path[len(root)] == filepath.Separator {
				return nil
			}
		}
	}
	return errors.New("security violation: path is outside of trusted root")
}
