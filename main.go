package main

import (
	_ "embed"
	"flag"
	"os"
	"strings"

	"github.com/aprx/apeupres/config"
	"github.com/aprx/apeupres/iniconfigfile"
	"github.com/aprx/apeupres/utils"
)

//go:embed fixtures/fixtures.sh
var fixtures string

const (
	defaultConfigPath string = "$HOME/.config/apeupres"
	defaultOutputPath string = "$HOME/.apeupresrc"
)

var clOptions struct {
	configPath  string
	output      string
	showCurrent bool
}

func init() {
	flag.StringVar(&clOptions.configPath, "config", os.ExpandEnv(defaultConfigPath), "Path to configuration")
	flag.StringVar(&clOptions.output, "output", os.ExpandEnv(defaultOutputPath), "Output path")
	flag.BoolVar(&clOptions.showCurrent, "current", false, "Show active config if any")
}

func main() {
	data := []string{fixtures}
	sources := []string{}
	flag.Parse()

	if clOptions.showCurrent {
		_ = utils.ShowActiveConfig()
		return
	}

	configurationFiles := config.GatherConfiguration(clOptions.configPath)
	for _, v := range configurationFiles {
		content := iniconfigfile.ProcessIniFile(v)
		if len(content) > 0 {
			sources = append(sources, v)
			data = append(data, content)
		}
	}
	sources = utils.Map(sources, func(element string) string { return "# " + element })
	data = append([]string{strings.Join(sources, "")}, data...)
	out, err := utils.GetOutputFile(clOptions.output)
	if err != nil {
		panic(err)
	}
	defer utils.CloseFileWithLog(out)
	utils.WriteConfig(out, strings.Join(data, ""))
}
