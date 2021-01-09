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

	styleMappingSetFile             = "./definitions/utility-mapping-set.gen.go"
	styleGeneratedFile              = "./definitions/utility-styles.gen.go"
	unitGeneratedFile               = "./definitions/style-units.gen.go"
	compositionGeneratedFile        = "./definitions/composer-styles.gen.go"
	styleMappingGeneratedFile       = "./definitions/mapping-styles.gen.go"
	styleGroupingsGeneratedFile     = "./definitions/groupings-styles.gen.go"
	styleMappingToRootGeneratedFile = "./definitions/root-mapping-styles.gen.go"
)

type StyleUnit struct {
	Status string   `json:"status"`
	Groups []string `json:"groups"`
}

type StyleUnits map[string]*StyleUnit

type StyleRule struct {
	Name           string   `json:"-"`
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
	unitContent.WriteString(`package definitions

		type Unit string
	`)

	for name, unit := range unitStyles {
		createUnit(&unitContent, name, unit)
	}

	var formattedUnitsContent, unitsFormatErr = formatCode(nunsafe.String2Bytes(unitContent.String()))
	if unitsFormatErr != nil {
		noError(unitsFormatErr, "Format file failure")
	}

	if err := writeFile(unitGeneratedFile, bytes.NewReader(formattedUnitsContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	var stylesUtilitiesToParentContent strings.Builder
	stylesUtilitiesToParentContent.WriteString(`package definitions

		var UtilitiesToParentStyles = map[string]string {
	`)

	var stylesUtilitiesContent strings.Builder
	stylesUtilitiesContent.WriteString(`package definitions

		var UtilitiesToCSS = map[string]string {
	`)

	var stylesToMapSetContent strings.Builder
	stylesToMapSetContent.WriteString(`package definitions

		var UtilityToMappingSet = map[string]map[string]string {
	`)

	var stylesContent strings.Builder
	stylesContent.WriteString(`package definitions
		
	`)

	var groupings = map[string][]string{}

	for name, styleUnit := range baseStyles {
		styleUnit.Name = name

		var webkitName = fmt.Sprintf("-%s-%s", "webkit", name)
		var mozName = fmt.Sprintf("-%s-%s", "moz", name)
		var msName = fmt.Sprintf("-%s-%s", "ms", name)

		var webkitStyle = chromeStyles[webkitName]
		if webkitStyle != nil {
			webkitStyle.Name = webkitName
		}

		var mozStyle = mozillaStyles[mozName]
		if mozStyle != nil {
			mozStyle.Name = mozName
		}

		var msStyle = microsoftStyles[msName]
		if msStyle != nil {
			msStyle.Name = msName
		}

		var utilitiesMapping, utilitiesToRoot, mappingSet, utilitiesNames = createStyling(&stylesContent, name, styleUnit, webkitStyle, mozStyle, msStyle)

		stylesToMapSetContent.WriteString(mappingSet.String())
		stylesUtilitiesContent.WriteString(utilitiesMapping.String())
		stylesUtilitiesToParentContent.WriteString(utilitiesToRoot.String())

		for _, group := range styleUnit.Groups {
			groupings[group] = append(groupings[group], name)
			groupings[group] = append(groupings[group], msName)
			groupings[group] = append(groupings[group], mozName)
			groupings[group] = append(groupings[group], webkitName)
			groupings[group] = append(groupings[group], utilitiesNames...)
		}
	}

	var formattedStylesContent, stylesFormatErr = formatCode(nunsafe.String2Bytes(stylesContent.String()))
	if stylesFormatErr != nil {
		noError(stylesFormatErr, "Format file failure")
	}

	if err := writeFile(styleGeneratedFile, bytes.NewReader(formattedStylesContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	var compositionContent strings.Builder
	compositionContent.WriteString(`package definitions

		// These are styles which define how a giving parent style will be compose from it's
        // utilities which break down a style into individual bits.
		var ParentCompositionStyles = map[string][]string {
	`)

	for name, styleUnit := range compositionStyles {
		var webkitName = fmt.Sprintf("-%s-%s", "webkit", name)
		var mozName = fmt.Sprintf("-%s-%s", "moz", name)
		var msName = fmt.Sprintf("-%s-%s", "ms", name)

		var webkitStyle = chromeStyles[webkitName]
		if webkitStyle != nil {
			webkitStyle.Name = webkitName
		}

		var mozStyle = mozillaStyles[mozName]
		if mozStyle != nil {
			mozStyle.Name = mozName
		}

		var msStyle = microsoftStyles[msName]
		if msStyle != nil {
			msStyle.Name = msName
		}

		createCompositionStyle(&compositionContent, name, styleUnit, webkitStyle, mozStyle, msStyle)
	}

	compositionContent.WriteString(`
		}
	`)

	var formattedCompositionContent, compositionFormatErr = formatCode(nunsafe.String2Bytes(compositionContent.String()))
	if compositionFormatErr != nil {
		noError(compositionFormatErr, "Format file failure")
	}

	if err := writeFile(compositionGeneratedFile, bytes.NewReader(formattedCompositionContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	stylesToMapSetContent.WriteString(`
		}
	`)

	stylesUtilitiesContent.WriteString(`
		}
	`)

	stylesUtilitiesToParentContent.WriteString(`
		}
	`)

	var formattedMapSet, mapSetFormatErr = formatCode(nunsafe.String2Bytes(stylesToMapSetContent.String()))
	if mapSetFormatErr != nil {
		noError(mapSetFormatErr, "Format file failure")
	}

	if err := writeFile(styleMappingSetFile, bytes.NewReader(formattedMapSet)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	var formattedUtilitiesContent, utilitiesFormatErr = formatCode(nunsafe.String2Bytes(stylesUtilitiesContent.String()))
	if utilitiesFormatErr != nil {
		noError(utilitiesFormatErr, "Format file failure")
	}

	if err := writeFile(styleMappingGeneratedFile, bytes.NewReader(formattedUtilitiesContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	var formattedUtilitiesToRootContent, utilitiesToRootFormatErr = formatCode(nunsafe.String2Bytes(stylesUtilitiesToParentContent.String()))
	if utilitiesToRootFormatErr != nil {
		noError(utilitiesToRootFormatErr, "Format file failure")
	}

	if err := writeFile(styleMappingToRootGeneratedFile, bytes.NewReader(formattedUtilitiesToRootContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	var stylesGroup strings.Builder
	stylesGroup.WriteString(`package definitions
		
		var Groupings = map[string][]string {
	`)
	for groupName, group := range groupings {
		var listGroup strings.Builder
		toList(&listGroup, group)
		stylesGroup.WriteString(fmt.Sprintf(`
			%q: {
				%s
			},
		`, groupName, listGroup.String()))
	}
	stylesGroup.WriteString(`
		}
	`)

	var formattedGroupingsContent, groupingFormatErr = formatCode(nunsafe.String2Bytes(stylesGroup.String()))
	if groupingFormatErr != nil {
		noError(groupingFormatErr, "Format file failure")
	}

	if err := writeFile(styleGroupingsGeneratedFile, bytes.NewReader(formattedGroupingsContent)); err != nil {
		noError(err, "Write file failure %q", chromePropertiesDataFile)
	}

	log.Println("Finished generating css style properties")
}

func createCompositionStyle(builder *strings.Builder, name string, unit, webkit, moz, ms *StyleRule) {
	var listGroup strings.Builder
	toList(&listGroup, unit.Composes)

	builder.WriteString(fmt.Sprintf(`
		%q: {
			%s
		},
	`, name, listGroup.String()))

	if ms != nil {
		builder.WriteString(fmt.Sprintf(`
			%q: {
				%s
			},
		`, ms.Name, listGroup.String()))
	}

	if moz != nil {
		builder.WriteString(fmt.Sprintf(`
			%q: {
				%s
			},
		`, moz.Name, listGroup.String()))
	}

	if webkit != nil {
		builder.WriteString(fmt.Sprintf(`
			%q: {
				%s
			},
		`, webkit.Name, listGroup.String()))
	}
}

func createStyling(builder *strings.Builder, name string, unit, webkit, moz, ms *StyleRule) (strings.Builder, strings.Builder, strings.Builder, []string) {
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

	var browserVariantMap strings.Builder
	if webkit != nil {
		browserVariantMap.WriteString(fmt.Sprintf("%q,\n", webkit.Name))
	}
	if moz != nil {
		browserVariantMap.WriteString(fmt.Sprintf("%q,\n", moz.Name))
	}
	if ms != nil {
		browserVariantMap.WriteString(fmt.Sprintf("%q,\n", ms.Name))
	}

	var syntaxToParentMap strings.Builder
	var styleSelectionMap strings.Builder
	styleSelectionMap.WriteString(fmt.Sprintf("%q: \"%s: %s\",\n", name, name, unit.Initial))

	var utilities = make([]string, 0, 10)
	if len(unit.Syntax) > 0 {
		var valueTypes = toSyntax(unit.Syntax)
		for _, valueType := range valueTypes {
			valueType = strings.TrimSpace(valueType)
			if len(valueType) == 0 {
				continue
			}

			var utilityTypeName = name + "-" + valueType
			utilities = append(utilities, utilityTypeName)

			styleSelectionMap.WriteString(fmt.Sprintf("%q: \"%s: %s\",\n", utilityTypeName, name, valueType))

			syntaxToParentMap.WriteString(fmt.Sprintf("%q: %q,\n", utilityTypeName, name))

			var valueTypeName = toCap(valueType)
			builder.WriteString(fmt.Sprintf(`
		const %sStyle%s %sStyle = %q
		`, typeName, valueTypeName, typeName, valueType))
		}
	}

	builder.WriteString(fmt.Sprintf(`
		var %sStyleBrowserVariantsList = []string{
			%s
		}
		func (t %sStyle) BrowserVariants() []string {
			return %sStyleBrowserVariantsList
		}
	`, typeName, browserVariantMap.String(), typeName, typeName))

	builder.WriteString(fmt.Sprintf(`
		var %sStyleUtilitiesMapping = map[string]string{
			%s
		}
		func (t %sStyle) Utilities() map[string]string {
			return %sStyleUtilitiesMapping
		}
	`, typeName, styleSelectionMap.String(), typeName, typeName))

	builder.WriteString(fmt.Sprintf(`
		func (t %sStyle) UtilityFor(tu string) (string, bool) {
			var value, hasValue = %sStyleUtilitiesMapping[tu]
			return value, hasValue
		}
	`, typeName, typeName))

	var utilityToMapset strings.Builder
	utilityToMapset.WriteString(fmt.Sprintf(`
		%q: %sStyleUtilitiesMapping,
	`, name, typeName))

	return styleSelectionMap, syntaxToParentMap, utilityToMapset, utilities
}

func toList(builder *strings.Builder, groups []string) {
	for _, group := range groups {
		builder.WriteString(fmt.Sprintf("%q,\n", group))
	}
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
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		parts[index] = strings.TrimSpace(mapSpacedToDashed(part))
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

func mapSpacedToDashed(s string) string {
	s = strings.TrimSpace(s)
	var partSlice = strings.Split(s, " ")
	for index, part := range partSlice {
		part = strings.TrimSpace(part)
		if part == "" {
			continue
		}
		partSlice[index] = part
	}
	return strings.Join(partSlice, "-")
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

func appendFile(targetFile string, content io.Reader) error {
	var endFile, err = os.OpenFile(targetFile, os.O_APPEND, 0)
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
