package main

import (
	"encoding/json"
	"io"
	"io/fs"
	"os"
	"path/filepath"

	"github.com/fatih/color"
)

type ipynbData struct {
	Cells []struct {
		CellType       string `json:"cell_type"`
		ExecutionCount int    `json:"execution_count"`
		Metadata       struct {
		} `json:"metadata"`
		Outputs []struct {
			Name       string   `json:"name"`
			OutputType string   `json:"output_type"`
			Text       []string `json:"text"`
		} `json:"outputs"`
		Source []string `json:"source"`
	} `json:"cells"`
	Metadata struct {
		Kernelspec struct {
			DisplayName string `json:"display_name"`
			Language    string `json:"language"`
			Name        string `json:"name"`
		} `json:"kernelspec"`
		LanguageInfo struct {
			CodemirrorMode struct {
				Name    string `json:"name"`
				Version int    `json:"version"`
			} `json:"codemirror_mode"`
			FileExtension     string `json:"file_extension"`
			Mimetype          string `json:"mimetype"`
			Name              string `json:"name"`
			NbconvertExporter string `json:"nbconvert_exporter"`
			PygmentsLexer     string `json:"pygments_lexer"`
			Version           string `json:"version"`
		} `json:"language_info"`
		OrigNbformat int `json:"orig_nbformat"`
		Vscode       struct {
			Interpreter struct {
				Hash string `json:"hash"`
			} `json:"interpreter"`
		} `json:"vscode"`
	} `json:"metadata"`
	Nbformat      int `json:"nbformat"`
	NbformatMinor int `json:"nbformat_minor"`
}

func (i ipynbData) HasOutputs() bool {
	if len(i.Cells) == 0 {
		return false
	}

	for _, cell := range i.Cells {
		if len(cell.Outputs) > 0 {
			return true
		}
	}

	return false
}

// walks filepath and finds filenames ending with .ipynb. pass glob patterns to
// ignore files/directories
func findIpynbPaths(
	root string,
	ignorePatterns ...string,
) (
	[]string,
	error,
) {
	matcher := newGitPatternMatcher(ignorePatterns...)
	paths := []string{}

	if err := filepath.WalkDir(
		root,
		func(path string, d fs.DirEntry, err error) error {
			if matcher.Match([]string{path}, d.IsDir()) {
				return filepath.SkipDir
			}

			if filepath.Ext(path) == ".ipynb" {
				paths = append(paths, path)
			}

			return nil
		},
	); err != nil {
		return nil, err
	}

	return paths, nil
}

func readIpynbFile(path string) (ipynbData, error) {
	file, err := os.Open(path)
	if err != nil {
		return ipynbData{}, err
	}
	defer file.Close()

	data, err := io.ReadAll(file)
	if err != nil {
		return ipynbData{}, err
	}

	ipynb := ipynbData{}

	if err := json.Unmarshal(data, &ipynb); err != nil {
		return ipynbData{}, err
	}

	return ipynb, nil
}

func scrubOutputData(ipynb ipynbData) (scrubbed ipynbData, outputs int) {
	outputs = 0

	for i := range ipynb.Cells {
		outputs += len(ipynb.Cells[i].Outputs)

		ipynb.Cells[i].ExecutionCount = 0
		ipynb.Cells[i].Outputs = nil
	}

	return ipynb, outputs
}

func writeIpynbFile(path string, ipynb ipynbData) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	data, err := json.MarshalIndent(ipynb, "", "  ")
	if err != nil {
		return err
	}

	if _, err := file.Write(data); err != nil {
		return err
	}

	return nil
}

func scrubIpynbFile(path string) (outputs int, err error) {
	ipynb, err := readIpynbFile(path)
	if err != nil {
		return 0, err
	}

	if !ipynb.HasOutputs() {
		return 0, nil
	}

	scrubbed, outputs := scrubOutputData(ipynb)

	if err := writeIpynbFile(path, scrubbed); err != nil {
		return 0, err
	}

	color.White("  %s [%d outputs]\n", path, outputs)

	return 0, nil
}

func scrubIpynbFiles(ignorePatterns ...string) error {
	color.Cyan("scrubbing .ipynb files...")

	paths, err := findIpynbPaths(".", ignorePatterns...)
	if err != nil {
		return err
	}

	if len(paths) == 0 {
		color.Cyan("no .ipynb files found")
		return nil
	}

	scrubbed := []string{}

	for _, path := range paths {
		outputs, err := scrubIpynbFile(path)
		if err != nil {
			return err
		}

		if outputs > 0 {
			scrubbed = append(scrubbed, path)
		}
	}

	if err := gitAddPaths(scrubbed); err != nil {
		return err
	}

	c := color.New(color.FgRed)
	if len(scrubbed) == 0 {
		c = color.New(color.FgGreen)
	}

	c.Println("scrubbed %d files", len(scrubbed))

	return nil
}
