package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go/format"
	"io"
	"log"
	"os"
	"path"
	"strings"

	"github.com/influx6/npkg/nerror"
	"github.com/influx6/npkg/nunsafe"
)

var (
	rulesDirectory              = "./data"
	unitsDataFile               = path.Join(rulesDirectory, "units.json")
	propertiesDataFile          = path.Join(rulesDirectory, "properties.json")
	compositionDataFile         = path.Join(rulesDirectory, "composition.json")
	chromePropertiesDataFile    = path.Join(rulesDirectory, "chrome.json")
	mozillaPropertiesDataFile   = path.Join(rulesDirectory, "mozilla.json")
	microsoftPropertiesDataFile = path.Join(rulesDirectory, "microsoft.json")

	styleGeneratedFile       = "utility_styles.gen.go"
	unitGeneratedFile        = "style_units.gen.go"
	compositionGeneratedFile = "composer_styles.gen.go"
)

type StyleUnit struct {
	Status string   `json:"status"`
	Groups []string `json:"groups"`
}

type StyleUnits map[string]*StyleUnit

type StyleRule struct {
	Inherited      bool     `json:"inherited"`
	Type           string   `json:"type"`
	Syntax         string   `json:"syntax"`
	Initial        string   `json:"initial"`
	Media          string   `json:"media"`
	Order          string   `json:"order"`
	Status         string   `json:"status"`
	AnimationType  string   `json:"animationType"`
	Percentage     string   `json:"percentages"`
	MDNUrl         string   `json:"mdn_url"`
	AppliesTo      string   `json:"appliesTo"`
	ComputedFor    string   `json:"computed_for"`
	Groups         []string `json:"groups"`
	MediaList      []string `json:"medias"`
	InitialList    []string `json:"initials"`
	PercentageList []string `json:"percentagesSet"`
	AlsoAppliesTo  []string `json:"alsoAppliesTo"`
	Computed       []string `json:"computed"`
	Composes       []string `json:"compose"`
	AnimationTypes []string `json:"animationTypes"`
}

type Styles map[string]*StyleRule

func main() {
	var unitStyles, unitErr = parseUnitFile(unitsDataFile)
	noError(unitErr, "Unit file %q", unitsDataFile)

	var baseStyles, baseReaderErr = parseStyleFile(propertiesDataFile)
	noError(baseReaderErr, "Style file %q", propertiesDataFile)

	var compositionStyles, compositionReaderErr = parseStyleFile(compositionDataFile)
	noError(compositionReaderErr, "Style file %q", compositionDataFile)

	var microsoftStyles, microsoftReaderErr = parseStyleFile(microsoftPropertiesDataFile)
	noError(microsoftReaderErr, "Style file %q", unitsDataFile)

	var mozillaStyles, mozillaReaderErr = parseStyleFile(mozillaPropertiesDataFile)
	noError(mozillaReaderErr, "Style file %q", mozillaPropertiesDataFile)

	var chromeStyles, chromeReaderErr = parseStyleFile(chromePropertiesDataFile)
	noError(chromeReaderErr, "Style file %q", chromePropertiesDataFile)

	var unitContent strings.Builder
	unitContent.WriteString(`package styled

		type Unit string
	`)

	for name, unit := range unitStyles {
		createUnit(&unitContent, name, unit)
	}

	var formattedUnitsContent, unitsformatErr = formatCode(nunsafe.String2Bytes(unitContent.String()))
	if unitsformatErr != nil {
		noError(unitsformatErr, "Format file failure")
	}

	if err := writeFile(unitGeneratedFile, bytes.NewReader(formattedUnitsContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	var stylesContent strings.Builder
	stylesContent.WriteString(`package styled
		
	`)

	for name, styleUnit := range baseStyles {
		var webkitName = fmt.Sprintf("-%s-%s", "webkit", name)
		var mozName = fmt.Sprintf("-%s-%s", "moz", name)
		var msName = fmt.Sprintf("-%s-%s", "ms", name)

		var webkitStyle = chromeStyles[webkitName]
		var mozStyle = mozillaStyles[mozName]
		var msStyle = microsoftStyles[msName]
		createStyling(&stylesContent, name, styleUnit, webkitStyle, mozStyle, msStyle)
	}

	var formattedStylesContent, stylesformatErr = formatCode(nunsafe.String2Bytes(stylesContent.String()))
	if stylesformatErr != nil {
		noError(stylesformatErr, "Format file failure")
	}

	if err := writeFile(styleGeneratedFile, bytes.NewReader(formattedStylesContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	var compositionContent strings.Builder
	compositionContent.WriteString(`package styled

	`)

	for name, styleUnit := range compositionStyles {
		var webkitName = fmt.Sprintf("-%s-%s", "webkit", name)
		var mozName = fmt.Sprintf("-%s-%s", "moz", name)
		var msName = fmt.Sprintf("-%s-%s", "ms", name)

		var webkitStyle = chromeStyles[webkitName]
		var mozStyle = mozillaStyles[mozName]
		var msStyle = microsoftStyles[msName]
		createCompositionStyle(&compositionContent, name, styleUnit, webkitStyle, mozStyle, msStyle)
	}

	var formattedCompositionContent, compositionformatErr = formatCode(nunsafe.String2Bytes(compositionContent.String()))
	if compositionformatErr != nil {
		noError(compositionformatErr, "Format file failure")
	}

	if err := writeFile(compositionGeneratedFile, bytes.NewReader(formattedCompositionContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

}

func createCompositionStyle(builder *strings.Builder, name string, unit, webkit, moz, ms *StyleRule) {
	// var styleType = "string"
	// if len(unit.Type) != 0 {
	// 	styleType = unit.Type
	// }
	//
	// var typeName = toCap(name)
	// builder.WriteString(fmt.Sprintf(`
	// 	// %s represent the CSS style %q with value %q
	// 	// See %s
	// 	type %sStyle %s
	// `, typeName, name, unit.Syntax, unit.MDNUrl, typeName, styleType))
}

func createStyling(builder *strings.Builder, name string, unit, webkit, moz, ms *StyleRule) {
	var styleType = "string"
	if len(unit.Type) != 0 {
		styleType = unit.Type
	}

	var typeName = toCap(name)
	builder.WriteString(fmt.Sprintf(`
		// %s represent the CSS style %q with value %q
		// See %s
		type %sStyle %s
	`, typeName, name, unit.Syntax, unit.MDNUrl, typeName, styleType))

	builder.WriteString(fmt.Sprintf(`
		func (t %sStyle) Name() string {
			return %q
		}
	`, typeName, name))

	var styleSelectionMap strings.Builder
	styleSelectionMap.WriteString(fmt.Sprintf("%s: true,", name))

	if len(unit.Syntax) > 0 {
		var valueTypes = toSyntax(unit.Syntax)
		for _, valueType := range valueTypes {
			valueType = strings.TrimSpace(valueType)
			if len(valueType) == 0 {
				continue
			}

			var utilityTypeName = name + "-" + valueType
			styleSelectionMap.WriteString(fmt.Sprintf("%s: true,", utilityTypeName))

			var valueTypeName = toCap(valueType)
			builder.WriteString(fmt.Sprintf(`
		const %sStyle%s = %q
		`, typeName, valueTypeName, valueType))
		}
	}

	builder.WriteString(fmt.Sprintf(`
		func (t %sStyle) Utilities() map[string]bool {
			return map[string]bool{
			%s
			}
		}
	`, typeName, styleSelectionMap.String()))
}

func createUnit(builder *strings.Builder, name string, unit *StyleUnit) {
	builder.WriteString(fmt.Sprintf(`
		// %s is a %s unit
		const %s Unit = %q
	`, name, unit.Status, strings.ToUpper(name), name))
}

func formatCode(content []byte) ([]byte, error) {
	var formatted, formattedErr = format.Source(content)
	if formattedErr != nil {
		return nil, nerror.Wrap(formattedErr, nunsafe.Bytes2String(content))
	}
	return formatted, nil
}

func toSyntax(s string) []string {
	s = strings.TrimSpace(s)
	var parts = strings.Split(s, "|")
	for index, part := range parts {
		parts[index] = strings.TrimSpace(part)
	}
	return parts
}

func toCap(s string) string {
	var parts = strings.Split(s, "-")
	for index, part := range parts {
		parts[index] = toCapitalized(mapSpaced(part))
	}
	return strings.Join(parts, "")
}

func mapSpaced(s string) string {
	var partSlice = strings.Split(s, " ")
	for index, part := range partSlice {
		partSlice[index] = toCapitalized(part)
	}
	return strings.Join(partSlice, "")
}

func toCapitalized(s string) string {
	if len(s) == 0 {
		return s
	}
	if len(s) == 1 {
		return strings.ToUpper(s)
	}
	return strings.ToUpper(s[:1]) + s[1:]
}

func writeFile(targetFile string, content io.Reader) error {
	var endFile, err = os.Create(targetFile)
	if err != nil {
		return nerror.WrapOnly(err)
	}

	defer endFile.Close()

	var written, writeErr = io.Copy(endFile, content)
	if writeErr != nil {
		return nerror.WrapOnly(writeErr)
	}
	if written == 0 {
		return nerror.New("failed to write content to file")
	}
	return nil
}

func noError(err error, message string, items ...interface{}) {
	if err != nil {
		var msg = fmt.Sprintf(message, items...)
		log.Fatalf("Failed: %s\nError: %s", msg, err)
	}
}

func parseUnitFile(targetFile string) (StyleUnits, error) {
	var styleReader, readerErr = os.Open(targetFile)
	if readerErr != nil {
		return nil, nerror.WrapOnly(readerErr)
	}

	defer styleReader.Close()

	var styles StyleUnits
	if parseErr := json.NewDecoder(styleReader).Decode(&styles); parseErr != nil {
		return nil, nerror.WrapOnly(parseErr)
	}
	return styles, nil
}

func parseStyleFile(targetFile string) (Styles, error) {
	var styleReader, readerErr = os.Open(targetFile)
	if readerErr != nil {
		return nil, nerror.WrapOnly(readerErr)
	}

	defer styleReader.Close()

	var styles Styles
	if parseErr := json.NewDecoder(styleReader).Decode(&styles); parseErr != nil {
		return nil, nerror.WrapOnly(parseErr)
	}
	return styles, nil
}
