package domain

const (
	PalettePrimaryBackground = "#fff8e7"
	PaletteAccentBackground1 = "#f8e7ff"
	PaletteAccentBackground2 = "#e7fff8"
)

type VisualThemePalette struct {
	ID                   ID     `json:"id"`
	Name                 string `json:"name"`
	PrimaryBackground    string `json:"primaryBackground"`
	AccentBackground1    string `json:"accentBackground1"`
	AccentBackground2    string `json:"accentBackground2"`
	SemanticWarningColor string `json:"semanticWarningColor"`
	SemanticErrorColor   string `json:"semanticErrorColor"`
	SemanticSuccessColor string `json:"semanticSuccessColor"`
	ReadabilityNotes     string `json:"readabilityNotes"`
}

func DefaultPalette() VisualThemePalette {
	return VisualThemePalette{ID: "soft-study", Name: "Soft Study", PrimaryBackground: PalettePrimaryBackground, AccentBackground1: PaletteAccentBackground1, AccentBackground2: PaletteAccentBackground2, SemanticWarningColor: "#8a5a00", SemanticErrorColor: "#b42318", SemanticSuccessColor: "#027a48", ReadabilityNotes: "Text, controls, charts, focus indicators, and state labels must remain distinguishable on all palette backgrounds."}
}
