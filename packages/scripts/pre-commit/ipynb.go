package main

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
