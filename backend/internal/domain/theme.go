// Package domain 提供主题调色板相关类型。
package domain

// PalettePrimaryBackground 常量表示主要背景色。
// PaletteAccentBackground1 常量表示强调背景色之一。
// PaletteAccentBackground2 常量表示强调背景色之二。
const (
	PalettePrimaryBackground = "#fff8e7"
	PaletteAccentBackground1 = "#f8e7ff"
	PaletteAccentBackground2 = "#e7fff8"
)

// VisualThemePalette 定义视觉主题调色板的配色方案。
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

// DefaultPalette 返回默认的 Soft Study 主题调色板。
func DefaultPalette() VisualThemePalette {
	return VisualThemePalette{ID: "soft-study", Name: "Soft Study", PrimaryBackground: PalettePrimaryBackground, AccentBackground1: PaletteAccentBackground1, AccentBackground2: PaletteAccentBackground2, SemanticWarningColor: "#8a5a00", SemanticErrorColor: "#b42318", SemanticSuccessColor: "#027a48", ReadabilityNotes: "Text, controls, charts, focus indicators, and state labels must remain distinguishable on all palette backgrounds."}
}
