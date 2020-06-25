// +build codegen

// Command aws-gen-gocli parses a JSON description of an AWS API and generates a
// Go file containing a client for the API.
//
//     aws-gen-gocli apis/s3/2006-03-03/api-2.json
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"strings"
	"sync"

	"github.com/jviney/aws-sdk-go-v2/private/model/api"
	"github.com/jviney/aws-sdk-go-v2/private/util"
)

func usage() {
	fmt.Fprintln(os.Stderr, `Usage: api-gen <options> [model path | file path]
Loads API models from file and generates SDK clients from the models.

The model path arguments can be globs, or paths to individual files. The
utiliity requires that the API model files follow the following pattern:

<root>/<servicename>/<api-version>/<model json files>

e.g:

./models/apis/s3/2006-03-01/*.json

Flags:`)
	flag.PrintDefaults()
}

// Generates service api, examples, and interface from api json definition files.
//
// Flags:
// -path: alternative service path to write generated files to for each service.
// -svc-import-path: The root Go import path to use for generated for service clients.
//
// Env:
//  SERVICES comma separated list of services to generate.
func main() {
	var svcPath, svcImportPath string
	flag.StringVar(&svcPath, "path", "service",
		"The `path` to generate service clients in to.",
	)
	flag.StringVar(&svcImportPath, "svc-import-path",
		api.SDKImportRoot+"/service",
		"The Go `import path` to generate client to be under.",
	)
	var keepUnsupportedAPIs bool
	flag.BoolVar(&keepUnsupportedAPIs, "keep-unsupported-apis",
		false,
		"If API models that use unsupported features will be ignored or return error.",
	)
	flag.Usage = usage
	flag.Parse()

	if len(os.Getenv("AWS_SDK_CODEGEN_DEBUG")) != 0 {
		api.LogDebug(os.Stdout)
	}

	// Make sure all paths are based on platform's pathing not Unix
	globs := flag.Args()
	for i, g := range globs {
		globs[i] = filepath.FromSlash(g)
	}
	svcPath = filepath.FromSlash(svcPath)

	modelPaths, err := api.ExpandModelGlobPath(globs...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to glob file pattern", err)
		os.Exit(1)
	}
	modelPaths, _ = api.TrimModelServiceVersions(modelPaths)

	loader := api.Loader{
		BaseImport:          svcImportPath,
		KeepUnsupportedAPIs: keepUnsupportedAPIs,
	}

	apis, err := loader.Load(modelPaths)
	if err != nil {
		fmt.Fprintln(os.Stderr, "failed to load API models", err)
		os.Exit(1)
	}
	if len(apis) == 0 {
		fmt.Fprintf(os.Stderr, "expected to load models, but found none")
		os.Exit(1)
	}

	if v := os.Getenv("SERVICES"); len(v) != 0 {
		svcs := strings.Split(v, ",")
		for pkgName, a := range apis {
			var found bool
			for _, include := range svcs {
				if a.PackageName() == include {
					found = true
					break
				}
			}
			if !found {
				delete(apis, pkgName)
			}
		}
	}

	var wg sync.WaitGroup
	servicePaths := map[string]struct{}{}
	for _, a := range apis {
		if _, ok := excludeServices[a.PackageName()]; ok {
			continue
		}

		// Create the output path for the model.
		pkgDir := filepath.Join(svcPath, a.PackageName())
		os.MkdirAll(filepath.Join(pkgDir, a.InterfacePackageName()), 0775)

		if _, ok := servicePaths[pkgDir]; ok {
			fmt.Fprintf(os.Stderr,
				"attempted to generate a client into %s twice. Second model package, %v\n",
				pkgDir, a.PackageName())
			os.Exit(1)
		}
		servicePaths[pkgDir] = struct{}{}

		g := &generateInfo{
			API:        a,
			PackageDir: pkgDir,
		}

		wg.Add(1)
		go func() {
			defer wg.Done()
			writeServiceFiles(g, pkgDir)
		}()
	}

	wg.Wait()
}

type generateInfo struct {
	*api.API
	PackageDir string
}

var excludeServices = map[string]struct{}{
	"importexport": {},
}

func writeServiceFiles(g *generateInfo, pkgDir string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Fprintf(os.Stderr, "Error generating %s\n%s\n%s\n",
				pkgDir, r, debug.Stack())
			os.Exit(1)
		}
	}()

	fmt.Printf("Generating %s (%s)...\n",
		g.API.PackageName(), g.API.Metadata.APIVersion)

	// write files for service client and API
	Must(writeServiceDocFile(g))
	Must(writeAPIFile(g))
	Must(writeClientFile(g))
	Must(writeClientConfig(g))
	Must(writeInterfaceFile(g))
	Must(writeWaitersFile(g))
	Must(writeAPIErrorsFile(g))
	Must(writeExamplesFile(g))

	if g.API.PackageName() == "s3" {
		Must(writeS3ManagerUploadInputFile(g))
	}

	if len(g.API.SmokeTests.TestCases) > 0 {
		Must(writeAPISmokeTestsFile(g))
	}
}

func writeClientConfig(g *generateInfo) error {
	if !g.HasExternalClientConfigFields() {
		return nil
	}

	err := writeGoFile(filepath.Join(g.PackageDir, "api_client_config.go"),
		codeLayout,
		"",
		g.API.PackageName(),
		g.API.ClientConfigGoCode(),
	)
	if err != nil {
		return err
	}

	err = writeGoFile(filepath.Join(g.PackageDir, "api_client_config_test.go"),
		codeLayout,
		"",
		g.API.PackageName(),
		g.API.ClientConfigTestsGoCode(),
	)
	if err != nil {
		return err
	}

	pkgName := g.InternalConfigResolverPackageName()
	pkgPath := filepath.Join(g.PackageDir, "internal", pkgName)
	err = os.MkdirAll(pkgPath, 0775)
	if err != nil {
		return fmt.Errorf("failed to create package directory: %v", err)
	}

	err = writeGoFile(
		filepath.Join(pkgPath, fmt.Sprintf("%s.go", pkgName)),
		codeLayout,
		"",
		pkgName,
		g.ClientConfigResolversGoCode(),
	)
	if err != nil {
		return err
	}

	err = writeGoFile(
		filepath.Join(pkgPath, fmt.Sprintf("%s_test.go", pkgName)),
		codeLayout,
		"",
		pkgName+"_test",
		g.ClientConfigResolversTestGoCode(),
	)
	if err != nil {
		return err
	}

	return nil
}

// Must will panic if the error passed in is not nil.
func Must(err error) {
	if err != nil {
		panic(err)
	}
}

const codeLayout = `// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

%s
package %s

%s
`

func writeGoFile(file string, layout string, args ...interface{}) error {
	return ioutil.WriteFile(file, []byte(util.GoFmt(fmt.Sprintf(layout, args...))), 0664)
}

// writeServiceDocFile generates the documentation for service package.
func writeServiceDocFile(g *generateInfo) error {
	return writeGoFile(filepath.Join(g.PackageDir, "api_doc.go"),
		codeLayout,
		strings.TrimSpace(g.API.ClientPackageDoc()),
		g.API.PackageName(),
		"",
	)
}

// writeExamplesFile writes out the service example file.
func writeExamplesFile(g *generateInfo) error {
	code := g.API.ExamplesGoCode()
	if len(code) > 0 {
		return writeGoFile(filepath.Join(g.PackageDir, "api_examples_test.go"),
			codeLayout,
			"",
			g.API.PackageName()+"_test",
			code,
		)
	}
	return nil
}

// writeClientFile writes out the service initialization file.
func writeClientFile(g *generateInfo) error {
	err := writeGoFile(filepath.Join(g.PackageDir, "api_client.go"),
		codeLayout,
		"",
		g.API.PackageName(),
		g.API.ClientGoCode(),
	)
	if err != nil {
		return err
	}

	return nil
}

// writeInterfaceFile writes out the service interface file.
func writeInterfaceFile(g *generateInfo) error {
	const pkgDoc = `
// Package %s provides an interface to enable mocking the %s service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.`
	return writeGoFile(filepath.Join(g.PackageDir, g.API.InterfacePackageName(), "interface.go"),
		codeLayout,
		fmt.Sprintf(pkgDoc, g.API.InterfacePackageName(), g.API.Metadata.ServiceFullName),
		g.API.InterfacePackageName(),
		g.API.InterfaceGoCode(),
	)
}

func writeWaitersFile(g *generateInfo) error {
	if len(g.API.Waiters) == 0 {
		return nil
	}

	return writeGoFile(filepath.Join(g.PackageDir, "api_waiters.go"),
		codeLayout,
		"",
		g.API.PackageName(),
		g.API.WaitersGoCode(),
	)
}

// writeAPIFile writes out the service API file.
func writeAPIFile(g *generateInfo) error {
	for opName, op := range g.API.Operations {
		err := writeGoFile(
			filepath.Join(g.PackageDir, "api_op_"+op.ExportedName+".go"),
			codeLayout,
			"",
			g.API.PackageName(),
			g.API.APIOperationGoCode(op),
		)

		if err != nil {
			return fmt.Errorf("failed to write %s operation file, %v", opName, err)
		}
	}

	if err := writeGoFile(filepath.Join(g.PackageDir, "api_enums.go"),
		codeLayout,
		"",
		g.API.PackageName(),
		g.API.APIEnumsGoCode(),
	); err != nil {
		return err
	}

	if err := writeGoFile(filepath.Join(g.PackageDir, "api_types.go"),
		codeLayout,
		"",
		g.API.PackageName(),
		g.API.APIParamShapesGoCode(),
	); err != nil {
		return err
	}

	return nil
}

// writeAPIErrorsFile writes out the service API errors file.
func writeAPIErrorsFile(g *generateInfo) error {
	return writeGoFile(filepath.Join(g.PackageDir, "api_errors.go"),
		codeLayout,
		"",
		g.API.PackageName(),
		g.API.APIErrorsGoCode(),
	)
}

func writeS3ManagerUploadInputFile(g *generateInfo) error {
	return writeGoFile(filepath.Join(g.PackageDir, "s3manager", "upload_input.go"),
		codeLayout,
		"",
		"s3manager",
		api.S3ManagerUploadInputGoCode(g.API),
	)
}

func writeAPISmokeTestsFile(g *generateInfo) error {
	return writeGoFile(filepath.Join(g.PackageDir, "api_integ_test.go"),
		codeLayout,
		"// +build integration\n",
		g.API.PackageName()+"_test",
		g.API.APISmokeTestsGoCode(),
	)
}
