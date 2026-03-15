package valueobjects

import (
	"regexp"
	"strings"
)

type Slug struct {
	value string
}

func NewSlug(text string) Slug {
	// 1. Convertir a minúsculas
	str := strings.ToLower(text)

	// 2. Reemplazo manual de caracteres con acentos/tildes comunes
	// Esto cubre los casos más habituales en español y otros idiomas latinos.
	replacer := strings.NewReplacer(
		"á", "a", "é", "e", "í", "i", "ó", "o", "ú", "u",
		"ñ", "n", "ü", "u",
		"ä", "a", "ë", "e", "ï", "i", "ö", "o",
	)
	str = replacer.Replace(str)

	// 3. Regex para sustituir cualquier cosa que NO sea letra minúscula o número por un guion
	reg := regexp.MustCompile(`[^a-z0-9]+`)
	str = reg.ReplaceAllString(str, "-")

	// 4. Limpieza final: quitar guiones de los extremos y evitar guiones dobles
	// Trim quita los de los bordes. El regex anterior ya colapsa secuencias como "---" en "-"
	str = strings.Trim(str, "-")

	return Slug{
		value: str,
	}
}

func ReconstituteSlug(value string) Slug {
	return Slug{
		value: value,
	}
}

func (s Slug) Value() string {
	return s.value
}
