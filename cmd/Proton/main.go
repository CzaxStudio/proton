package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		printHelp()
		return
	}

	switch os.Args[1] {
	case "set":
		if len(os.Args) < 5 || os.Args[3] != "for" {
			fmt.Println("Usage: Proton set <image_path> for <file1,file2,...>")
			return
		}
		imgPath := os.Args[2]
		files := strings.Split(os.Args[4], ",")
		for _, f := range files {
			targetFile := strings.TrimSpace(f)
			if targetFile == "" {
				continue
			}
			err := addImageToFile(imgPath, targetFile)
			if err != nil {
				fmt.Printf("Error setting image for %s: %v\n", targetFile, err)
			}
		}
	case "logo":
		if len(os.Args) < 3 {
			fmt.Println("usage: Proton logo <image_path>")
			return
		}
		err := setLogo(os.Args[2])
		if err != nil {
			fmt.Printf("error setting logo: %v\n", err)
		}
	default:
		printHelp()
	}
}

func printHelp() {
	fmt.Println("Proton GUI Framework CLI")
	fmt.Println()
	fmt.Println("usage:")
	fmt.Println("  Proton logo <image_path>          add a logo to your app")
	fmt.Println("  Proton set <img> for <files>      embed an image for specific files")
}

func setLogo(imgPath string) error {
	ext := filepath.Ext(imgPath)
	dstName := "logo" + ext

	err := copyFile(imgPath, dstName)
	if err != nil {
		return err
	}

	// find package name from first .go file
	pkgName := "main"
	matches, _ := filepath.Glob("*.go")
	for _, m := range matches {
		if strings.HasSuffix(m, "_gen.go") {
			continue
		}
		if p, err := getPackageName(m); err == nil {
			pkgName = p
			break
		}
	}

	genFileName := "logo_gen.go"
	content := fmt.Sprintf(`package %s

import _ "embed"

//go:embed %s
var Logo_data []byte
`, pkgName, dstName)

	err = os.WriteFile(genFileName, []byte(content), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("success: logo set and image copied to %s\n", dstName)
	fmt.Println()
	fmt.Println("add this to your code to display it:")
	fmt.Println("----------------------------------------")
	fmt.Println("a.SetLogo(Logo_data)")
	fmt.Println("...")
	fmt.Println("proton.Logo(win, 64, 64)")
	fmt.Println("----------------------------------------")
	return nil
}

func addImageToFile(imgPath, targetFile string) error {
	// 1. Copy image to target directory
	dir := filepath.Dir(targetFile)
	imgName := filepath.Base(imgPath)
	dstPath := filepath.Join(dir, imgName)

	if imgPath != dstPath {
		err := copyFile(imgPath, dstPath)
		if err != nil {
			return err
		}
	}

	// 2. Identify package name of target file
	pkgName, err := getPackageName(targetFile)
	if err != nil {
		return err
	}

	// 3. Create generated Go file for embedding
	baseName := strings.TrimSuffix(imgName, filepath.Ext(imgName))
	// Clean name for variable
	varName := strings.Map(func(r rune) rune {
		if (r >= 'a' && r <= 'z') || (r >= 'A' && r <= 'Z') || (r >= '0' && r <= '9') {
			return r
		}
		return '_'
	}, baseName)

	genFileName := filepath.Join(dir, varName+"_gen.go")
	content := fmt.Sprintf(`package %s

import _ "embed"

//go:embed %s
var %s_data []byte
`, pkgName, imgName, varName)

	err = os.WriteFile(genFileName, []byte(content), 0644)
	if err != nil {
		return err
	}

	fmt.Printf("Set %s for %s (package %s). Use variable: %s_data\n", imgName, targetFile, pkgName, varName)
	return nil
}

func copyFile(srcPath, dstPath string) error {
	srcAbs, _ := filepath.Abs(srcPath)
	dstAbs, _ := filepath.Abs(dstPath)
	if srcAbs == dstAbs {
		return nil // nothing to do
	}

	src, err := os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()

	dst, err := os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()

	_, err = io.Copy(dst, src)
	return err
}

func getPackageName(path string) (string, error) {
	f, err := os.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "package ") {
			parts := strings.Fields(line)
			if len(parts) >= 2 {
				// Remove semicolon if present (e.g. package main;)
				return strings.TrimSuffix(parts[1], ";"), nil
			}
		}
	}
	return "main", nil // fallback
}
