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

//go:embed lib/math-geo.oak
var libmathgeo string

//go:embed lib/math-stats.oak
var libmathstats string

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

//go:embed lib/msgpack.oak
var libmsgpack string

//go:embed lib/datetime.oak
var libdatetime string

//go:embed lib/path.oak
var libpath string

//go:embed lib/compression.oak
var libcompression string

//go:embed lib/compression-rle.oak
var libcompressionrle string

//go:embed lib/compression-huffman.oak
var libcompressionhuffman string

//go:embed lib/compression-lzw.oak
var libcompressionlzw string

//go:embed lib/dataprot.oak
var libdataprot string

//go:embed lib/http.oak
var libhttp string

//go:embed lib/socket.oak
var libsocket string

//go:embed lib/smtp.oak
var libsmtp string

//go:embed lib/pop.oak
var libpop string

//go:embed lib/imap.oak
var libimap string

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

//go:embed lib/ico.oak
var libico string

//go:embed lib/gpu.oak
var libgpu string

//go:embed lib/gpus.oak
var libgpus string

//go:embed lib/syntax.oak
var libsyntax string

//go:embed lib/syntax-tokenize.oak
var libsyntaxtokenize string

//go:embed lib/syntax-parse.oak
var libsyntaxparse string

//go:embed lib/syntax-macros.oak
var libsyntaxmacros string

//go:embed lib/syntax-print.oak
var libsyntaxprint string

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

//go:embed lib/build-render-node.oak
var libbuildrendernode string

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

//go:embed lib/wasm-vm.oak
var libwasmvm string

//go:embed lib/wasm-vm-runtime.oak
var libwasmvmruntime string

//go:embed lib/codegen-common.oak
var libcodegencommon string

//go:embed lib/sys.oak
var libsys string

//go:embed lib/writes.oak
var libwrites string

//go:embed lib/windows.oak
var libwindows string

//go:embed lib/windows-constants.oak
var libwindowsconstants string

//go:embed lib/windows-flags.oak
var libwindowsflags string

//go:embed lib/windows-registry.oak
var libwindowsregistry string

//go:embed lib/windows-net.oak
var libwindowsnet string

//go:embed lib/windows-windowing.oak
var libwindowswindowing string

//go:embed lib/windows-gdi.oak
var libwindowsgdi string

//go:embed lib/windows-kernel.oak
var libwindowskernel string

//go:embed lib/windows-loader.oak
var libwindowsloader string

//go:embed lib/windows-core.oak
var libwindowscore string

//go:embed lib/Linux.oak
var liblinux string

//go:embed lib/linux-constants.oak
var liblinuxconstants string

//go:embed lib/linux-core.oak
var liblinuxcore string

//go:embed lib/linux-loader.oak
var liblinuxloader string

//go:embed lib/linux-windowing.oak
var liblinuxwindowing string

//go:embed lib/linux-libc.oak
var liblinuxlibc string

//go:embed lib/websocket.oak
var libwebsocket string

//go:embed lib/p2p.oak
var libp2p string

//go:embed lib/gui-common.oak
var libguicommon string

//go:embed lib/win-common.oak
var libwincommon string

//go:embed lib/GUI.oak
var libgui string

//go:embed lib/gui-mesh.oak
var libguimesh string

//go:embed lib/gui-render.oak
var libguirender string

//go:embed lib/gui-3dmath.oak
var libgui3dmath string

//go:embed lib/gui-raster.oak
var libguiraster string

//go:embed lib/gui-web.oak
var libguiweb string

//go:embed lib/gui-native-win.oak
var libguinativewin string

//go:embed lib/gui-native-win-present.oak
var libguinativewinpresent string

//go:embed lib/gui-native-win-icons.oak
var libguinativewinicons string

//go:embed lib/gui-native-win-frame.oak
var libguinativewinframe string

//go:embed lib/gui-native-win-poll.oak
var libguinativewinpoll string

//go:embed lib/gui-native-win-close.oak
var libguinativewinclose string

//go:embed lib/gui-native-win-ddraw.oak
var libguinativewinddraw string

//go:embed lib/gui-native-win-vulkan.oak
var libguinativewinvulkan string

//go:embed lib/gui-native-win-opengl.oak
var libguinativewinopengl string

//go:embed lib/gui-native-win-probe.oak
var libguinativewinprobe string

//go:embed lib/gui-native-linux.oak
var libguinativelinux string

//go:embed lib/gui-shader.oak
var libguishader string

//go:embed lib/gui-shader-math.oak
var libguishadermath string

//go:embed lib/gui-shader-color.oak
var libguishadercolor string

//go:embed lib/gui-shader-noise.oak
var libguishadernoise string

//go:embed lib/gui-shader-sdf.oak
var libguishadersdf string

//go:embed lib/gui-shader-codegen.oak
var libguishadercodegen string

//go:embed lib/gui-draw.oak
var libguidraw string

//go:embed lib/gui-2d.oak
var libgui2d string

//go:embed lib/gui-color.oak
var libguicolor string

//go:embed lib/gui-events.oak
var libguievents string

//go:embed lib/gui-graph.oak
var libguigraph string

//go:embed lib/gui-form.oak
var libguiform string

//go:embed lib/gui-loop.oak
var libguiloop string

//go:embed lib/ai.oak
var libai string

//go:embed lib/audio.oak
var libaudio string

//go:embed lib/audio-complex.oak
var libaudiocomplex string

//go:embed lib/audio-fft.oak
var libaudiofft string

//go:embed lib/audio-dsp.oak
var libaudiodsp string

//go:embed lib/audio-wav.oak
var libaudiowav string

//go:embed lib/video.oak
var libvideo string

//go:embed lib/async-event-bus.oak
var libasynceventbus string

//go:embed lib/WLAN.oak
var libwlan string

var stdlibs = map[string]string{
	"std":                    libstd,
	"str":                    libstr,
	"math":                   libmath,
	"math-geo":               libmathgeo,
	"math-stats":             libmathstats,
	"sort":                   libsort,
	"random":                 librandom,
	"fs":                     libfs,
	"linux-constants":        liblinuxconstants,
	"linux-core":             liblinuxcore,
	"linux-loader":           liblinuxloader,
	"linux-windowing":        liblinuxwindowing,
	"linux-libc":             liblinuxlibc,
	"fmt":                    libfmt,
	"json":                   libjson,
	"msgpack":                libmsgpack,
	"datetime":               libdatetime,
	"path":                   libpath,
	"compression":            libcompression,
	"compression-rle":        libcompressionrle,
	"compression-huffman":    libcompressionhuffman,
	"compression-lzw":        libcompressionlzw,
	"dataprot":               libdataprot,
	"http":                   libhttp,
	"socket":                 libsocket,
	"smtp":                   libsmtp,
	"pop":                    libpop,
	"imap":                   libimap,
	"test":                   libtest,
	"debug":                  libdebug,
	"cli":                    libcli,
	"md":                     libmd,
	"crypto":                 libcrypto,
	"bitwise":                libbitwise,
	"bmp":                    libbmp,
	"ico":                    libico,
	"gpu":                    libgpu,
	"gpus":                   libgpus,
	"syntax":                 libsyntax,
	"syntax-tokenize":        libsyntaxtokenize,
	"syntax-parse":           libsyntaxparse,
	"syntax-macros":          libsyntaxmacros,
	"syntax-print":           libsyntaxprint,
	"syntaxfmt":              libsyntaxfmt,
	"Virtual":                libvirtual,
	"VirtualToken":           libvirtualtoken,
	"pack-utils":             libpackutils,
	"bundle-ast":             libbundleast,
	"bundle-utils":           libbundleutils,
	"ast-analyze":            libastanalyze,
	"ast-transform":          libasttransform,
	"shell":                  libshell,
	"thread":                 libthread,
	"transpile":              libtranspile,
	"pack":                   libpack,
	"build":                  libbuild,
	"build-includes":         libbuildincludes,
	"build-ident":            libbuildident,
	"build-ast":              libbuildast,
	"build-analyze":          libbuildanalyze,
	"build-render":           libbuildrender,
	"build-render-node":      libbuildrendernode,
	"build-config":           libbuildconfig,
	"build-imports":          libbuildimports,
	"runtime-native":         libruntimenative,
	"runtime-js":             libruntimejs,
	"runtime-codegen":        libruntimecodegen,
	"wasm-vm":                libwasmvm,
	"wasm-vm-runtime":        libwasmvmruntime,
	"codegen-common":         libcodegencommon,
	"sys":                    libsys,
	"writes":                 libwrites,
	"windows-constants":      libwindowsconstants,
	"windows-flags":          libwindowsflags,
	"windows-registry":       libwindowsregistry,
	"windows-net":            libwindowsnet,
	"windows-windowing":      libwindowswindowing,
	"windows-gdi":            libwindowsgdi,
	"windows-kernel":         libwindowskernel,
	"windows-loader":         libwindowsloader,
	"windows-core":           libwindowscore,
	"windows":                libwindows,
	"linux":                  liblinux,
	"Linux":                  liblinux,
	"websocket":              libwebsocket,
	"p2p":                    libp2p,
	"gui-mesh":               libguimesh,
	"gui-common":             libguicommon,
	"win-common":             libwincommon,
	"gui-render":             libguirender,
	"gui-3dmath":             libgui3dmath,
	"gui-raster":             libguiraster,
	"gui-web":                libguiweb,
	"gui-native-win":         libguinativewin,
	"gui-native-win-present": libguinativewinpresent,
	"gui-native-win-icons":   libguinativewinicons,
	"gui-native-win-frame":   libguinativewinframe,
	"gui-native-win-poll":    libguinativewinpoll,
	"gui-native-win-close":   libguinativewinclose,
	"gui-native-win-ddraw":   libguinativewinddraw,
	"gui-native-win-vulkan":  libguinativewinvulkan,
	"gui-native-win-opengl":  libguinativewinopengl,
	"gui-native-win-probe":   libguinativewinprobe,
	"gui-native-linux":       libguinativelinux,
	"gui-draw":               libguidraw,
	"gui-2d":                 libgui2d,
	"gui-color":              libguicolor,
	"gui-events":             libguievents,
	"gui-graph":              libguigraph,
	"gui-form":               libguiform,
	"gui-loop":               libguiloop,
	"GUI":                    libgui,
	"gui":                    libgui,
	"ai":                     libai,
	"audio":                  libaudio,
	"audio-complex":          libaudiocomplex,
	"audio-fft":              libaudiofft,
	"audio-dsp":              libaudiodsp,
	"audio-wav":              libaudiowav,
	"video":                  libvideo,
	"async/event-bus":        libasynceventbus,
	"WLAN":                   libwlan,
	"gui-shader":             libguishader,
	"gui-shader-math":        libguishadermath,
	"gui-shader-color":       libguishadercolor,
	"gui-shader-noise":       libguishadernoise,
	"gui-shader-sdf":         libguishadersdf,
	"gui-shader-codegen":     libguishadercodegen,
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
	"compression",
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

	c.eng.importLock.RLock()
	imported, ok := c.eng.importMap[name]
	c.eng.importLock.RUnlock()
	if ok {
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

	c.eng.importLock.Lock()
	c.eng.importMap[name] = ctx.scope
	c.eng.importLock.Unlock()
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
