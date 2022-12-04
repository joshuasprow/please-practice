package scripts

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
)

// validates file paths are correct and copies file from src to dst, creating
// directories if necessary
func CopyFile(src, dst string) error {
	srcAbs, err := filepath.Abs(
		src,
	)
	if err != nil {
		return err
	}

	dstAbs, err := filepath.Abs(dst)
	if err != nil {
		return err
	}

	fmt.Printf("copying %s to %s\n", srcAbs, dstAbs)

	dstDir := filepath.Dir(dstAbs)

	if err := os.MkdirAll(dstDir, 0755); err != nil {
		log.Printf("error creating directory %s: %v", dstDir, err)
	}

	dstFile, err := os.Create(dstAbs)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	srcFile, err := os.Open(srcAbs)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	if _, err := io.Copy(dstFile, srcFile); err != nil {
		return err
	}

	return nil
}
