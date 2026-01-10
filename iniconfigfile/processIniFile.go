// Package iniconfigfile handle ini format file and transforms it to shell functions.
package iniconfigfile

import (
	"bytes"
	"fmt"
	"strings"

	"gopkg.in/ini.v1"
)

func sectionToFn(section string, params *map[string]string, keys *[]string) string {
	var buffer bytes.Buffer

	fmt.Fprintf(&buffer, "conf_%s() {\n", section)
	fmt.Fprintf(&buffer, "  apeupres_set_env %s\n", section)
	fmt.Fprintf(&buffer, "  apeupres_set_clean_env %s\n", strings.Join(*keys, ":"))
	for k, v := range *params {
		fmt.Fprintf(&buffer, "  export %s=%s\n", k, v)
	}
	fmt.Fprintf(&buffer, "}\n")
	return buffer.String()
}

func ProcessIniFile(iniFileName string) string {
	iniFile, err := ini.Load(iniFileName)
	if err != nil {
		panic(err)
	}

	var sectionsData []string

	for _, section := range iniFile.Sections() {
		var keysData []string
		fnData := map[string]string{}
		if section.Name() == "unset" {
			panic("You cannot have as section named unset\n")
		}
		if section.Name() == "DEFAULT" {
			continue
		}
		if len(section.Keys()) > 0 {
			for _, field := range section.Keys() {
				keysData = append(keysData, field.Name())
				fnData[field.Name()] = field.Value()
			}
		}
		sectionsData = append(sectionsData, sectionToFn(section.Name(), &fnData, &keysData))
	}
	return strings.Join(sectionsData, "\n")
}
