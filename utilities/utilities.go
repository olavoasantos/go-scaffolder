package utilities

import (
	"errors"
	"strings"
	"unicode"

	"github.com/gertd/go-pluralize"
)

func Check(e error, params ...string) {
	if e != nil {
		err := e
		if len(params) >= 1 {
			err = errors.New(params[0])
		}
		panic(err)
	}
}

// Based on: https://gist.github.com/stoewer/fbe273b711e6a06315d19552dd4d33e6#gistcomment-3688390
func ToSnakeCase(value string) string {
	var result = make([]rune, 0, len(value))
	var p = '_'
	for i, r := range value {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			result = append(result, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				result = append(result, '_', unicode.ToLower(r))
			} else {
				result = append(result, unicode.ToLower(r))
			}
		} else {
			result = append(result, unicode.ToLower(r))
		}

		p = r
	}
	return string(result)
}

// Based on: https://github.com/iancoleman/strcase/blob/master/camel.go
func ToCamelCase(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return value
	}

	capNext := false
	result := strings.Builder{}
	result.Grow(len(value))
	for i, v := range []byte(value) {
		vIsNumber := v >= '0' && v <= '9'
		vIsUpperCase := v >= 'A' && v <= 'Z'
		vIsLowerCase := v >= 'a' && v <= 'z'

		if capNext {
			if vIsLowerCase {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsUpperCase {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsUpperCase || vIsLowerCase {
			result.WriteByte(v)
			capNext = false
		} else if vIsNumber {
			result.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}

	return result.String()
}

func ToPascalCase(value string) string {
	value = strings.TrimSpace(value)
	if value == "" {
		return value
	}

	capNext := true
	result := strings.Builder{}
	result.Grow(len(value))
	for i, v := range []byte(value) {
		vIsNumber := v >= '0' && v <= '9'
		vIsUpperCase := v >= 'A' && v <= 'Z'
		vIsLowerCase := v >= 'a' && v <= 'z'

		if capNext {
			if vIsLowerCase {
				v += 'A'
				v -= 'a'
			}
		} else if i == 0 {
			if vIsUpperCase {
				v += 'a'
				v -= 'A'
			}
		}
		if vIsUpperCase || vIsLowerCase {
			result.WriteByte(v)
			capNext = false
		} else if vIsNumber {
			result.WriteByte(v)
			capNext = true
		} else {
			capNext = v == '_' || v == ' ' || v == '-' || v == '.'
		}
	}

	return result.String()
}

func ToKebabCase(value string) string {
	var result = make([]rune, 0, len(value))
	var p = '-'
	for i, r := range value {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			result = append(result, '-')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				result = append(result, '-', unicode.ToLower(r))
			} else {
				result = append(result, unicode.ToLower(r))
			}
		} else {
			result = append(result, unicode.ToLower(r))
		}

		p = r
	}
	return string(result)
}

func ToMacroCase(value string) string {
	var result = make([]rune, 0, len(value))
	var p = '_'
	for i, r := range value {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			result = append(result, '_')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				result = append(result, '_', unicode.ToUpper(r))
			} else {
				result = append(result, unicode.ToUpper(r))
			}
		} else {
			result = append(result, unicode.ToUpper(r))
		}

		p = r
	}
	return string(result)
}

func ToUpperCase(value string) string {
	var result = make([]rune, 0, len(value))
	var p = ' '
	for i, r := range value {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			result = append(result, ' ')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				result = append(result, ' ', unicode.ToUpper(r))
			} else {
				result = append(result, unicode.ToUpper(r))
			}
		} else {
			result = append(result, unicode.ToUpper(r))
		}

		p = r
	}
	return string(result)
}

func ToLowerCase(value string) string {
	var result = make([]rune, 0, len(value))
	var p = ' '
	for i, r := range value {
		if !unicode.IsLetter(r) && !unicode.IsDigit(r) {
			result = append(result, ' ')
		} else if unicode.IsUpper(r) && i > 0 {
			if unicode.IsLetter(p) && !unicode.IsUpper(p) || unicode.IsDigit(p) {
				result = append(result, ' ', unicode.ToLower(r))
			} else {
				result = append(result, unicode.ToLower(r))
			}
		} else {
			result = append(result, unicode.ToLower(r))
		}

		p = r
	}
	return string(result)
}

var pluralizeClient = pluralize.NewClient()

func ToPlural(value string) string {
	return pluralizeClient.Plural(value)
}

func ToSingular(value string) string {
	return pluralizeClient.Singular(value)
}

type Variations struct {
	VALUE              string
	UPPERCASE          string
	LOWERCASE          string
	KEBAB              string
	CAMEL              string
	SNAKE              string
	PASCAL             string
	MACRO              string
	SINGULAR_VALUE     string
	SINGULAR_UPPERCASE string
	SINGULAR_LOWERCASE string
	SINGULAR_KEBAB     string
	SINGULAR_CAMEL     string
	SINGULAR_SNAKE     string
	SINGULAR_PASCAL    string
	SINGULAR_MACRO     string
	PLURAL_VALUE       string
	PLURAL_UPPERCASE   string
	PLURAL_LOWERCASE   string
	PLURAL_KEBAB       string
	PLURAL_CAMEL       string
	PLURAL_SNAKE       string
	PLURAL_PASCAL      string
	PLURAL_MACRO       string
}

func VariationsOf(value string) Variations {
	return Variations{
		VALUE:              value,
		UPPERCASE:          ToUpperCase(value),
		LOWERCASE:          ToLowerCase(value),
		KEBAB:              ToKebabCase(value),
		CAMEL:              ToCamelCase(value),
		SNAKE:              ToSnakeCase(value),
		PASCAL:             ToPascalCase(value),
		MACRO:              ToMacroCase(value),
		SINGULAR_VALUE:     ToSingular(value),
		SINGULAR_UPPERCASE: ToSingular(ToUpperCase(value)),
		SINGULAR_LOWERCASE: ToSingular(ToLowerCase(value)),
		SINGULAR_KEBAB:     ToSingular(ToKebabCase(value)),
		SINGULAR_CAMEL:     ToSingular(ToCamelCase(value)),
		SINGULAR_SNAKE:     ToSingular(ToSnakeCase(value)),
		SINGULAR_PASCAL:    ToSingular(ToPascalCase(value)),
		SINGULAR_MACRO:     ToSingular(ToMacroCase(value)),
		PLURAL_VALUE:       ToPlural(value),
		PLURAL_UPPERCASE:   ToPlural(ToUpperCase(value)),
		PLURAL_LOWERCASE:   ToPlural(ToLowerCase(value)),
		PLURAL_KEBAB:       ToPlural(ToKebabCase(value)),
		PLURAL_CAMEL:       ToPlural(ToCamelCase(value)),
		PLURAL_SNAKE:       ToPlural(ToSnakeCase(value)),
		PLURAL_PASCAL:      ToPlural(ToPascalCase(value)),
		PLURAL_MACRO:       ToPlural(ToMacroCase(value)),
	}
}
