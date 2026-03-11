package main

import (
	_ "embed"
	"fmt"
	"os"
	"strings"
)

//go:embed lib/std.oak
var libstd string

//go:embed lib/str.oak
var libstr string

//go:embed lib/math.oak
var libmath string

//go:embed lib/sort.oak
var libsort string

//go:embed lib/random.oak
var librandom string

//go:embed lib/fs.oak
var libfs string

//go:embed lib/fmt.oak
var libfmt string

//go:embed lib/json.oak
var libjson string

//go:embed lib/datetime.oak
var libdatetime string

//go:embed lib/path.oak
var libpath string

//go:embed lib/http.oak
var libhttp string

//go:embed lib/test.oak
var libtest string

//go:embed lib/thread.oak
var libthread string

//go:embed lib/debug.oak
var libdebug string

//go:embed lib/cli.oak
var libcli string

//go:embed lib/md.oak
var libmd string

//go:embed lib/crypto.oak
var libcrypto string

//go:embed lib/bitwise.oak
var libbitwise string

//go:embed lib/bmp.oak
var libbmp string

//go:embed lib/gpu.oak
var libgpu string

//go:embed lib/gpus.oak
var libgpus string

//go:embed lib/syntax.oak
var libsyntax string

//go:embed lib/syntaxfmt.oak
var libsyntaxfmt string

//go:embed lib/Virtual.oak
var libvirtual string

//go:embed lib/VirtualToken.oak
var libvirtualtoken string

//go:embed lib/pack-utils.oak
var libpackutils string

//go:embed lib/bundle-utils.oak
var libbundleutils string

//go:embed lib/ast-transform.oak
var libasttransform string

//go:embed lib/ast-analyze.oak
var libastanalyze string

//go:embed lib/bundle-ast.oak
var libbundleast string

//go:embed lib/shell.oak
var libshell string

//go:embed lib/transpile.oak
var libtranspile string

//go:embed lib/pack.oak
var libpack string

//go:embed lib/build.oak
var libbuild string

//go:embed lib/build-includes.oak
var libbuildincludes string

//go:embed lib/build-ident.oak
var libbuildident string

//go:embed lib/build-ast.oak
var libbuildast string

//go:embed lib/build-analyze.oak
var libbuildanalyze string

//go:embed lib/build-render.oak
var libbuildrender string

//go:embed lib/build-config.oak
var libbuildconfig string

//go:embed lib/build-imports.oak
var libbuildimports string

//go:embed lib/runtime-native.oak
var libruntimenative string

//go:embed lib/runtime-js.oak
var libruntimejs string

//go:embed lib/runtime-codegen.oak
var libruntimecodegen string

//go:embed lib/codegen-common.oak
var libcodegencommon string

//go:embed lib/sys.oak
var libsys string

//go:embed lib/writes.oak
var libwrites string

//go:embed lib/windows.oak
var libwindows string

//go:embed lib/Linux.oak
var liblinux string

//go:embed lib/websocket.oak
var libwebsocket string

var stdlibs = map[string]string{
	"std":             libstd,
	"str":             libstr,
	"math":            libmath,
	"sort":            libsort,
	"random":          librandom,
	"fs":              libfs,
	"fmt":             libfmt,
	"json":            libjson,
	"datetime":        libdatetime,
	"path":            libpath,
	"http":            libhttp,
	"test":            libtest,
	"debug":           libdebug,
	"cli":             libcli,
	"md":              libmd,
	"crypto":          libcrypto,
	"bitwise":         libbitwise,
	"bmp":             libbmp,
	"gpu":             libgpu,
	"gpus":            libgpus,
	"syntax":          libsyntax,
	"syntaxfmt":       libsyntaxfmt,
	"Virtual":         libvirtual,
	"VirtualToken":    libvirtualtoken,
	"pack-utils":      libpackutils,
	"bundle-ast":      libbundleast,
	"bundle-utils":    libbundleutils,
	"ast-analyze":     libastanalyze,
	"ast-transform":   libasttransform,
	"shell":           libshell,
	"thread":          libthread,
	"transpile":       libtranspile,
	"pack":            libpack,
	"build":           libbuild,
	"build-includes":  libbuildincludes,
	"build-ident":     libbuildident,
	"build-ast":       libbuildast,
	"build-analyze":   libbuildanalyze,
	"build-render":    libbuildrender,
	"build-config":    libbuildconfig,
	"build-imports":   libbuildimports,
	"runtime-native":  libruntimenative,
	"runtime-js":      libruntimejs,
	"runtime-codegen": libruntimecodegen,
	"codegen-common":  libcodegencommon,
	"sys":             libsys,
	"writes":          libwrites,
	"windows":         libwindows,
	"linux":           liblinux,
	"Linux":           liblinux,
	"websocket":       libwebsocket,
}

var stdlibAutoLoadOrder = []string{
	"std",
	"str",
	"math",
	"sort",
	"random",
	"fs",
	"fmt",
	"json",
	"datetime",
	"path",
	"http",
	"test",
	"debug",
	"cli",
	"md",
	"crypto",
	"gpu",
	"gpus",
	"syntax",
	"Virtual",
	"shell",
}

func isStdLib(name string) bool {
	_, ok := stdlibs[name]
	return ok
}

func (c *Context) LoadLib(name string) (Value, *runtimeError) {
	program, ok := stdlibs[name]
	if !ok {
		return nil, &runtimeError{
			reason: fmt.Sprintf("%s is not a valid standard library; could not import", name),
		}
	}

	if imported, ok := c.eng.importMap[name]; ok {
		return ObjectValue(imported.vars), nil
	}

	ctx := c.ChildContext(c.rootPath)
	ctx.currentFile = "(lib:" + name + ")"
	ctx.LoadBuiltins()

	ctx.Unlock()
	_, err := ctx.evalGo(strings.NewReader(program))
	ctx.Lock()
	if err != nil {
		if runtimeErr, ok := err.(*runtimeError); ok {
			return nil, runtimeErr
		} else {
			return nil, &runtimeError{
				reason: fmt.Sprintf("Error loading %s: %s", name, err.Error()),
			}
		}
	}

	c.eng.importMap[name] = ctx.scope
	return ObjectValue(ctx.scope.vars), nil
}

func (c *Context) loadAllLibs() error {
	for _, libname := range stdlibAutoLoadOrder {
		_, err := c.Eval(strings.NewReader(fmt.Sprintf("import('%s')", libname)))
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Context) mustLoadAllLibs() {
	if err := c.loadAllLibs(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
