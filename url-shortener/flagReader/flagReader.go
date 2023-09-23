package flagReader

import "flag"

type FileType int

const (
	None FileType = iota
	Yaml
	Json
)

const (
	DefaultFileName = ""
)

func Parse() (string, FileType) {
	fileYaml := flag.String("yaml", DefaultFileName, "YAML file with short urls to redirect")
	fileJson := flag.String("json", DefaultFileName, "JSON file with short urls to redirect")

	flag.Parse()

	areBothDefault := *fileYaml == DefaultFileName && *fileJson == DefaultFileName
	if areBothDefault {
		return DefaultFileName, None
	}

	areBothSpecified := *fileYaml != DefaultFileName && *fileJson != DefaultFileName
	if areBothSpecified {
		panic("You must specify only one file with short urls to redirect")
	}

	if *fileYaml != DefaultFileName {
		var fileType FileType = Yaml
		return *fileYaml, fileType
	}

	var fileType FileType = Json
	return *fileJson, fileType
}
