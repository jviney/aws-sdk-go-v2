// +build codegen

package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
	"text/template"
)

const (
	sharedConfigType = "&SharedConfig{}"
	envConfigType    = "&EnvConfig{}"
)

var implAsserts = map[string][]string{
	"SharedConfigProfileProvider":     {envConfigType, `WithSharedConfigProfile("")`},
	"SharedConfigFilesProvider":       {envConfigType, "WithSharedConfigFiles([]string{})"},
	"CustomCABundleProvider":          {envConfigType, "WithCustomCABundle([]byte{})"},
	"RegionProvider":                  {envConfigType, sharedConfigType, `WithRegion("")`, "&WithEC2MetadataRegion{}"},
	"MFATokenFuncProvider":            {`WithMFATokenFunc(func() (string, error) { return "", nil })`},
	"EnableEndpointDiscoveryProvider": {envConfigType, sharedConfigType, "WithEnableEndpointDiscovery(true)"},
	"CredentialsProviderProvider":     {`WithCredentialsProvider{aws.NewStaticCredentialsProvider("", "", "")}`},
	"DefaultRegionProvider":           {`WithDefaultRegion("")`},
}

var tplProviderTests = template.Must(template.New("tplProviderTests").Funcs(map[string]interface{}{
	"SortKeys": func(m map[string][]string) []string {
		var keys []string
		for k := range m {
			keys = append(keys, k)
		}

		sort.Strings(keys)

		return keys
	},
}).Parse(`
// Code generated by aws/external/codegen. DO NOT EDIT.

package external

import (
	"github.com/jviney/aws-sdk-go-v2/aws"
)

{{ $sortedKeys := SortKeys . }}
{{- range $_, $provider := $sortedKeys }}
	{{- $implementors := index $ $provider -}}
	// {{ $provider }} implementor assertions
	var (
		{{- range $_, $impl := $implementors }}
		_ {{ $provider }} = {{ $impl }}
		{{- end }}
	)
{{ end -}}
`))

var config = struct {
	OutputPath string
}{}

func init() {
	flag.StringVar(&config.OutputPath, "output", "", "output file path")
	flag.Parse()
}

func main() {
	if len(config.OutputPath) == 0 {
		panic(fmt.Errorf("output path must be specified"))
	}

	file, err := os.OpenFile(config.OutputPath, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0744)
	if err != nil {
		panic(fmt.Errorf("failed to open file: %v", err))
	}
	defer file.Close()

	err = tplProviderTests.Execute(file, implAsserts)
	if err != nil {
		panic(fmt.Errorf("failed to execute template: %v", err))
	}
}
