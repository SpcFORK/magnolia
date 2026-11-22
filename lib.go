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

//go:embed lib/math-base.oak
var libmathbase string

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

//go:embed lib/compression-ans.oak
var libcompressionans string

//go:embed lib/dataprot.oak
var libdataprot string

//go:embed lib/http.oak
var libhttp string

//go:embed lib/socket.oak
var libsocket string

//go:embed lib/email-smtp.oak
var libemailsmtp string

//go:embed lib/email-pop.oak
var libemailpop string

//go:embed lib/email-imap.oak
var libemailimap string

//go:embed lib/email.oak
var libemail string

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

//go:embed lib/mermaid.oak
var libmermaid string

//go:embed lib/crypto.oak
var libcrypto string

//go:embed lib/bitwise.oak
var libbitwise string

//go:embed lib/rmqr.oak
var librmqr string

//go:embed lib/bytes.oak
var libbytes string

//go:embed lib/sample.oak
var libsample string

//go:embed lib/beam-vm.oak
var libbeamvm string

//go:embed lib/hex.oak
var libhex string

//go:embed lib/ieee754.oak
var libieee754 string

//go:embed lib/int64.oak
var libint64 string

//go:embed lib/unicode-units.oak
var libunicodeunits string

//go:embed lib/image-bmp.oak
var libimagebmp string

//go:embed lib/image-ico.oak
var libimageico string

//go:embed lib/image-ppm.oak
var libimageppm string

//go:embed lib/image-tga.oak
var libimagetga string

//go:embed lib/image-qoi.oak
var libimageqoi string

//go:embed lib/image.oak
var libimage string

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

//go:embed lib/Virtual-Bytecode.oak
var libvirtualbytecode string

//go:embed lib/pack-utils.oak
var libpackutils string

//go:embed lib/bundle-utils.oak
var libbundleutils string

//go:embed lib/ast-transform.oak
var libasttransform string

//go:embed lib/ast-analyze.oak
var libastanalyze string

//go:embed lib/ast-js.oak
var libastjs string

//go:embed lib/js-runtime.oak
var libjsrtn string

//go:embed lib/estree.oak
var libestree string

//go:embed lib/js-tokenize.oak
var libjstok string

//go:embed lib/js-parse.oak
var libjsparse string

//go:embed lib/ui-surface.oak
var libuisurface string

//go:embed lib/ui-core.oak
var libuicore string

//go:embed lib/ui-sdk.oak
var libuisdk string

//go:embed lib/ui-sim.oak
var libuisim string

//go:embed lib/style-bake.oak
var libstylebake string

//go:embed lib/font-bake.oak
var libfontbake string

//go:embed lib/transpile-ui.oak
var libtranspileui string

//go:embed lib/gui-drawlist.oak
var libguidrawlist string

//go:embed lib/ast-ts.oak
var libastts string

//go:embed lib/ast-lua.oak
var libastlua string

//go:embed lib/ast-java.oak
var libastjava string

//go:embed lib/ast-go.oak
var libastgo string

//go:embed lib/ast-ir.oak
var libastir string

//go:embed lib/applet.oak
var libapplet string

//go:embed lib/applet-api.oak
var libappletapi string

//go:embed lib/applet-bytecode.oak
var libappletbytecode string

//go:embed lib/applet-class.oak
var libappletclass string

//go:embed lib/applet-host.oak
var libapplethost string

//go:embed lib/applet-java.oak
var libappletjava string

//go:embed lib/applet-js-host.oak
var libappletjshost string

//go:embed lib/applet-manifest.oak
var libappletmanifest string

//go:embed lib/applet-oak.oak
var libappletoak string

//go:embed lib/applet-recompile.oak
var libappletrecompile string

//go:embed lib/applet-runtime.oak
var libappletruntime string

//go:embed lib/applet-zip.oak
var libappletzip string

//go:embed lib/build-applet.oak
var libbuildapplet string

//go:embed lib/jvm-bridge.oak
var libjvmbridge string

//go:embed lib/jvm-jar-manifest.oak
var libjvmjarmanifest string

//go:embed lib/jvm-jar-write.oak
var libjvmjarwrite string

//go:embed lib/jvm-java-input.oak
var libjvmjavainput string

//go:embed lib/jvm-java-subset.oak
var libjvmjavasubset string

//go:embed lib/jvm-fp.oak
var libjvmfp string

//go:embed lib/jvm-vm-libs.oak
var libjvmvmlibs string

//go:embed lib/codegen-go.oak
var libcodegengo string

//go:embed lib/codegen-native.oak
var libcodegennative string

//go:embed lib/assembler.oak
var libassembler string

//go:embed lib/linker.oak
var liblinker string

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

//go:embed lib/gui-dialogs.oak
var libguidialogs string

//go:embed lib/win-common.oak
var libwincommon string

//go:embed lib/GUI.oak
var libgui string

//go:embed lib/gui-mesh.oak
var libguimesh string

//go:embed lib/gui-render.oak
var libguirender string

//go:embed lib/gui-resolution.oak
var libguiresolution string

//go:embed lib/gui-canvas.oak
var libguicanvas string

//go:embed lib/gui-accessibility.oak
var libguiaccessibility string

//go:embed lib/gui-clipboard.oak
var libguiclipboard string

//go:embed lib/gui-filedrop.oak
var libguifiledrop string

//go:embed lib/gui-audio.oak
var libguiaudio string

//go:embed lib/gui-gamepad.oak
var libguigamepad string

//go:embed lib/gui-aa.oak
var libguiaa string

//go:embed lib/gui-draw-ops.oak
var libguidrawops string

//go:embed lib/gui-gpu-info.oak
var libguigpuinfo string

//go:embed lib/gui-leak-detect.oak
var libguileakdetect string

//go:embed lib/gui-menus.oak
var libguimenus string

//go:embed lib/gui-print.oak
var libguiprint string

//go:embed lib/gui-test.oak
var libguitest string

//go:embed lib/gui-theme.oak
var libguitheme string

//go:embed lib/gui-systray.oak
var libguisystray string

//go:embed lib/gui-3dmath.oak
var libgui3dmath string

//go:embed lib/gui-raster.oak
var libguiraster string

//go:embed lib/gui-lighting.oak
var libguilighting string

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

//go:embed lib/gui-native-win-d3d11.oak
var libguinatived3d11 string

//go:embed lib/gui-native-win-vulkan.oak
var libguinativewinvulkan string

//go:embed lib/gui-native-win-vulkan-constants.oak
var libguinativewinvulkanconstants string

//go:embed lib/gui-native-win-vulkan-init.oak
var libguinativewinvulkaninit string

//go:embed lib/gui-native-win-vulkan-swapchain.oak
var libguinativewinvulkanswapchain string

//go:embed lib/gui-native-win-vulkan-present.oak
var libguinativewinvulkanpresent string

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

//go:embed lib/gui-input.oak
var libguiinput string

//go:embed lib/gui-graph.oak
var libguigraph string

//go:embed lib/gui-form.oak
var libguiform string

//go:embed lib/gui-loop.oak
var libguiloop string

//go:embed lib/gui-thread.oak
var libguithread string

//go:embed lib/ai.oak
var libai string

//go:embed lib/ai-vec.oak
var libaivec string

//go:embed lib/ai-linalg.oak
var libailinalg string

//go:embed lib/ai-nn.oak
var libainn string

//go:embed lib/ai-optim.oak
var libaioptim string

//go:embed lib/ai-data.oak
var libaidata string

//go:embed lib/ai-text.oak
var libaitext string

//go:embed lib/ai-ml.oak
var libaiml string

//go:embed lib/ai-decode.oak
var libaidecode string

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

//go:embed lib/audio-aiff.oak
var libaudioaiff string

//go:embed lib/audio-au.oak
var libaudioau string

//go:embed lib/audio-raw.oak
var libaudioraw string

//go:embed lib/audio-ogg.oak
var libaudioogg string

//go:embed lib/video.oak
var libvideo string

//go:embed lib/gui-video.oak
var libguivideo string

//go:embed lib/async-event-bus.oak
var libasynceventbus string

//go:embed lib/WLAN.oak
var libwlan string

//go:embed lib/gui-fonts.oak
var libguifonts string

//go:embed lib/windows-fonts.oak
var libwindowsfonts string

//go:embed lib/linux-fonts.oak
var liblinuxfonts string

//go:embed lib/codecols.oak
var libcodecols string

//go:embed lib/data-csv.oak
var libdatacsv string

//go:embed lib/data-toml.oak
var libdatatoml string

//go:embed lib/data-xml.oak
var libdataxml string

//go:embed lib/data-ini.oak
var libdataini string

//go:embed lib/data-yaml.oak
var libdatayaml string

//go:embed lib/data.oak
var libdata string

//go:embed lib/web-fetch.oak
var libwebfetch string

//go:embed lib/web-runtime.oak
var libwebruntime string

//go:embed lib/b64.oak
var libb64 string

//go:embed lib/b64-hex.oak
var libb64hex string

// Opt-in, separately preflighted JVM/MIDP and media helpers. The auto-load list is unchanged.
//
//go:embed lib/applet.oak
var libapplet string

//go:embed lib/applet-api.oak
var libappletapi string

//go:embed lib/applet-bytecode.oak
var libappletbytecode string

//go:embed lib/applet-class.oak
var libappletclass string

//go:embed lib/applet-gui.oak
var libappletgui string

//go:embed lib/applet-host.oak
var libapplethost string

//go:embed lib/applet-java.oak
var libappletjava string

//go:embed lib/applet-js-host.oak
var libappletjshost string

//go:embed lib/applet-manifest.oak
var libappletmanifest string

//go:embed lib/applet-oak.oak
var libappletoak string

//go:embed lib/applet-recompile.oak
var libappletrecompile string

//go:embed lib/applet-runtime.oak
var libappletruntime string

//go:embed lib/applet-zip.oak
var libappletzip string

//go:embed lib/build-applet.oak
var libbuildapplet string

//go:embed lib/fliplay-control.oak
var libfliplaycontrol string

//go:embed lib/font-bake.oak
var libfontbake string

//go:embed lib/gui-bounded-input.oak
var libguiboundedinput string

//go:embed lib/gui-drawlist.oak
var libguidrawlist string

//go:embed lib/gui-image.oak
var libguiimage string

//go:embed lib/gui-pixels.oak
var libguipixels string

//go:embed lib/jvm-bridge.oak
var libjvmbridge string

//go:embed lib/jvm-fp.oak
var libjvmfp string

//go:embed lib/jvm-jar-manifest.oak
var libjvmjarmanifest string

//go:embed lib/jvm-jar-write.oak
var libjvmjarwrite string

//go:embed lib/jvm-java-input.oak
var libjvmjavainput string

//go:embed lib/jvm-java-subset.oak
var libjvmjavasubset string

//go:embed lib/jvm-vm-libs.oak
var libjvmvmlibs string

//go:embed lib/midlet.oak
var libmidlet string

//go:embed lib/midp-device-profile.oak
var libmidpdeviceprofile string

//go:embed lib/midp-indexed-png.oak
var libmidpindexedpng string

//go:embed lib/style-bake.oak
var libstylebake string

//go:embed lib/swf-as3-cws.oak
var libswfas3cws string

//go:embed lib/swf-as3-abc-pool.oak
var libswfas3abcpool string

//go:embed lib/swf-as3-abc-signatures.oak
var libswfas3abcsignatures string

//go:embed lib/swf-as3-abc-structure.oak
var libswfas3abcstructure string

//go:embed lib/swf-tag-scan.oak
var libswftagscan string

//go:embed lib/avm1-preloader.oak
var libavm1preloader string

//go:embed lib/avm1-swf-timeline-prefix.oak
var libavm1swftimelineprefix string

//go:embed lib/avm1-swf-button-prefix.oak
var libavm1swfbuttonprefix string

//go:embed lib/transpile-ui.oak
var libtranspileui string

//go:embed lib/ui-core.oak
var libuicore string

//go:embed lib/ui-sdk.oak
var libuisdk string

//go:embed lib/ui-sim.oak
var libuisim string

//go:embed lib/ui-surface.oak
var libuisurface string

var stdlibs = map[string]string{
	"std":                             libstd,
	"str":                             libstr,
	"math":                            libmath,
	"math-base":                       libmathbase,
	"math-geo":                        libmathgeo,
	"math-stats":                      libmathstats,
	"sort":                            libsort,
	"random":                          librandom,
	"fs":                              libfs,
	"linux-constants":                 liblinuxconstants,
	"linux-core":                      liblinuxcore,
	"linux-loader":                    liblinuxloader,
	"linux-windowing":                 liblinuxwindowing,
	"linux-libc":                      liblinuxlibc,
	"fmt":                             libfmt,
	"json":                            libjson,
	"msgpack":                         libmsgpack,
	"datetime":                        libdatetime,
	"path":                            libpath,
	"compression":                     libcompression,
	"compression-rle":                 libcompressionrle,
	"compression-huffman":             libcompressionhuffman,
	"compression-lzw":                 libcompressionlzw,
	"compression-ans":                 libcompressionans,
	"dataprot":                        libdataprot,
	"http":                            libhttp,
	"socket":                          libsocket,
	"email-smtp":                      libemailsmtp,
	"email-pop":                       libemailpop,
	"email-imap":                      libemailimap,
	"email":                           libemail,
	"smtp":                            libsmtp,
	"pop":                             libpop,
	"imap":                            libimap,
	"test":                            libtest,
	"debug":                           libdebug,
	"cli":                             libcli,
	"md":                              libmd,
	"mermaid":                         libmermaid,
	"crypto":                          libcrypto,
	"bitwise":                         libbitwise,
	"bytes":                           libbytes,
	"sample":                          libsample,
	"beam-vm":                         libbeamvm,
	"hex":                             libhex,
	"ieee754":                         libieee754,
	"int64":                           libint64,
	"unicode-units":                   libunicodeunits,
	"bmp":                             libimagebmp,
	"ico":                             libimageico,
	"ppm":                             libimageppm,
	"tga":                             libimagetga,
	"qoi":                             libimageqoi,
	"rmqr":                            librmqr,
	"image-bmp":                       libimagebmp,
	"image-ico":                       libimageico,
	"image-ppm":                       libimageppm,
	"image-tga":                       libimagetga,
	"image-qoi":                       libimageqoi,
	"image":                           libimage,
	"gpu":                             libgpu,
	"gpus":                            libgpus,
	"syntax":                          libsyntax,
	"syntax-tokenize":                 libsyntaxtokenize,
	"syntax-parse":                    libsyntaxparse,
	"syntax-macros":                   libsyntaxmacros,
	"syntax-print":                    libsyntaxprint,
	"syntaxfmt":                       libsyntaxfmt,
	"Virtual":                         libvirtual,
	"VirtualToken":                    libvirtualtoken,
	"Virtual-Bytecode":                libvirtualbytecode,
	"pack-utils":                      libpackutils,
	"bundle-ast":                      libbundleast,
	"bundle-utils":                    libbundleutils,
	"ast-analyze":                     libastanalyze,
	"ast-transform":                   libasttransform,
	"ast-js":                          libastjs,
	"js-runtime":                       libjsrtn,
	"estree":                           libestree,
	"js-tokenize":                      libjstok,
	"js-parse":                         libjsparse,
	"ui-surface":                       libuisurface,
	"ui-core":                          libuicore,
	"ui-sdk":                           libuisdk,
	"ui-sim":                           libuisim,
	"style-bake":                       libstylebake,
	"font-bake":                        libfontbake,
	"transpile-ui":                     libtranspileui,
	"gui-drawlist":                     libguidrawlist,
	"ast-ts":                          libastts,
	"ast-lua":                         libastlua,
	"ast-java":                        libastjava,
	"ast-go":                          libastgo,
	"ast-ir":                          libastir,
	"applet":                          libapplet,
	"applet-api":                          libappletapi,
	"applet-bytecode":                          libappletbytecode,
	"applet-class":                          libappletclass,
	"applet-host":                          libapplethost,
	"applet-java":                          libappletjava,
	"applet-js-host":                          libappletjshost,
	"applet-manifest":                          libappletmanifest,
	"applet-oak":                          libappletoak,
	"applet-recompile":                          libappletrecompile,
	"applet-runtime":                          libappletruntime,
	"applet-zip":                          libappletzip,
	"build-applet":                          libbuildapplet,
	"jvm-bridge":                          libjvmbridge,
	"jvm-jar-manifest":                          libjvmjarmanifest,
	"jvm-jar-write":                          libjvmjarwrite,
	"jvm-java-input":                          libjvmjavainput,
	"jvm-java-subset":                          libjvmjavasubset,
	"jvm-fp":                           libjvmfp,
	"jvm-vm-libs":                          libjvmvmlibs,
	"codegen-go":                      libcodegengo,
	"codegen-native":                  libcodegennative,
	"assembler":                       libassembler,
	"linker":                          liblinker,
	"shell":                           libshell,
	"thread":                          libthread,
	"transpile":                       libtranspile,
	"pack":                            libpack,
	"build":                           libbuild,
	"build-includes":                  libbuildincludes,
	"build-ident":                     libbuildident,
	"build-ast":                       libbuildast,
	"build-analyze":                   libbuildanalyze,
	"build-render":                    libbuildrender,
	"build-render-node":               libbuildrendernode,
	"build-config":                    libbuildconfig,
	"build-imports":                   libbuildimports,
	"runtime-native":                  libruntimenative,
	"runtime-js":                      libruntimejs,
	"runtime-codegen":                 libruntimecodegen,
	"wasm-vm":                         libwasmvm,
	"wasm-vm-runtime":                 libwasmvmruntime,
	"codegen-common":                  libcodegencommon,
	"sys":                             libsys,
	"writes":                          libwrites,
	"windows-constants":               libwindowsconstants,
	"windows-flags":                   libwindowsflags,
	"windows-registry":                libwindowsregistry,
	"windows-net":                     libwindowsnet,
	"windows-windowing":               libwindowswindowing,
	"windows-gdi":                     libwindowsgdi,
	"windows-kernel":                  libwindowskernel,
	"windows-loader":                  libwindowsloader,
	"windows-core":                    libwindowscore,
	"windows":                         libwindows,
	"linux":                           liblinux,
	"Linux":                           liblinux,
	"websocket":                       libwebsocket,
	"p2p":                             libp2p,
	"gui-mesh":                        libguimesh,
	"gui-common":                      libguicommon,
	"win-common":                      libwincommon,
	"gui-render":                      libguirender,
	"gui-resolution":                  libguiresolution,
	"gui-3dmath":                      libgui3dmath,
	"gui-raster":                      libguiraster,
	"gui-lighting":                    libguilighting,
	"gui-web":                         libguiweb,
	"gui-native-win":                  libguinativewin,
	"gui-native-win-present":          libguinativewinpresent,
	"gui-native-win-icons":            libguinativewinicons,
	"gui-native-win-frame":            libguinativewinframe,
	"gui-native-win-poll":             libguinativewinpoll,
	"gui-native-win-close":            libguinativewinclose,
	"gui-native-win-ddraw":            libguinativewinddraw,
	"gui-native-win-d3d11":            libguinatived3d11,
	"gui-native-win-vulkan":           libguinativewinvulkan,
	"gui-native-win-vulkan-constants": libguinativewinvulkanconstants,
	"gui-native-win-vulkan-init":      libguinativewinvulkaninit,
	"gui-native-win-vulkan-swapchain": libguinativewinvulkanswapchain,
	"gui-native-win-vulkan-present":   libguinativewinvulkanpresent,
	"gui-native-win-opengl":           libguinativewinopengl,
	"gui-native-win-probe":            libguinativewinprobe,
	"gui-native-linux":                libguinativelinux,
	"gui-canvas":                      libguicanvas,
	"gui-accessibility":               libguiaccessibility,
	"gui-audio":                       libguiaudio,
	"gui-clipboard":                   libguiclipboard,
	"gui-filedrop":                    libguifiledrop,
	"gui-gamepad":                     libguigamepad,
	"gui-aa":                          libguiaa,
	"gui-draw-ops":                    libguidrawops,
	"gui-gpu-info":                    libguigpuinfo,
	"gui-leak-detect":                 libguileakdetect,
	"gui-menus":                       libguimenus,
	"gui-print":                       libguiprint,
	"gui-test":                        libguitest,
	"gui-theme":                       libguitheme,
	"gui-systray":                     libguisystray,
	"gui-draw":                        libguidraw,
	"gui-dialogs":                     libguidialogs,
	"gui-2d":                          libgui2d,
	"gui-color":                       libguicolor,
	"gui-events":                      libguievents,
	"gui-input":                       libguiinput,
	"gui-graph":                       libguigraph,
	"gui-form":                        libguiform,
	"gui-loop":                        libguiloop,
	"gui-thread":                      libguithread,
	"GUI":                             libgui,
	"gui":                             libgui,
	"ai":                              libai,
	"ai-vec":                          libaivec,
	"ai-linalg":                       libailinalg,
	"ai-nn":                           libainn,
	"ai-optim":                        libaioptim,
	"ai-data":                         libaidata,
	"ai-text":                         libaitext,
	"ai-ml":                           libaiml,
	"ai-decode":                       libaidecode,
	"audio":                           libaudio,
	"audio-complex":                   libaudiocomplex,
	"audio-fft":                       libaudiofft,
	"audio-dsp":                       libaudiodsp,
	"audio-wav":                       libaudiowav,
	"audio-aiff":                      libaudioaiff,
	"audio-au":                        libaudioau,
	"audio-raw":                       libaudioraw,
	"audio-ogg":                       libaudioogg,
	"video":                           libvideo,
	"gui-video":                       libguivideo,
	"async/event-bus":                 libasynceventbus,
	"WLAN":                            libwlan,
	"gui-shader":                      libguishader,
	"gui-shader-math":                 libguishadermath,
	"gui-shader-color":                libguishadercolor,
	"gui-shader-noise":                libguishadernoise,
	"gui-shader-sdf":                  libguishadersdf,
	"gui-shader-codegen":              libguishadercodegen,
	"gui-fonts":                       libguifonts,
	"windows-fonts":                   libwindowsfonts,
	"linux-fonts":                     liblinuxfonts,
	"codecols":                        libcodecols,
	"data-csv":                        libdatacsv,
	"data-toml":                       libdatatoml,
	"data-xml":                        libdataxml,
	"data-ini":                        libdataini,
	"data-yaml":                       libdatayaml,
	"data":                            libdata,
	"web-fetch":                       libwebfetch,
	"web-runtime":                     libwebruntime,
	"b64":                             libb64,
	"b64-hex":                         libb64hex,
	"applet":                          libapplet,
	"applet-api":                      libappletapi,
	"applet-bytecode":                 libappletbytecode,
	"applet-class":                    libappletclass,
	"applet-gui":                      libappletgui,
	"applet-host":                     libapplethost,
	"applet-java":                     libappletjava,
	"applet-js-host":                  libappletjshost,
	"applet-manifest":                 libappletmanifest,
	"applet-oak":                      libappletoak,
	"applet-recompile":                libappletrecompile,
	"applet-runtime":                  libappletruntime,
	"applet-zip":                      libappletzip,
	"build-applet":                    libbuildapplet,
	"fliplay-control":                 libfliplaycontrol,
	"font-bake":                       libfontbake,
	"gui-bounded-input":               libguiboundedinput,
	"gui-drawlist":                    libguidrawlist,
	"gui-image":                       libguiimage,
	"gui-pixels":                      libguipixels,
	"jvm-bridge":                      libjvmbridge,
	"jvm-fp":                          libjvmfp,
	"jvm-jar-manifest":                libjvmjarmanifest,
	"jvm-jar-write":                   libjvmjarwrite,
	"jvm-java-input":                  libjvmjavainput,
	"jvm-java-subset":                 libjvmjavasubset,
	"jvm-vm-libs":                     libjvmvmlibs,
	"midlet":                          libmidlet,
	"midp-device-profile":             libmidpdeviceprofile,
	"midp-indexed-png":                libmidpindexedpng,
	"style-bake":                      libstylebake,
	"swf-as3-cws":                     libswfas3cws,
	"swf-as3-abc-pool":                libswfas3abcpool,
	"swf-as3-abc-signatures":          libswfas3abcsignatures,
	"swf-as3-abc-structure":           libswfas3abcstructure,
	"swf-tag-scan":                    libswftagscan,
	"avm1-preloader":                  libavm1preloader,
	"avm1-swf-timeline-prefix":        libavm1swftimelineprefix,
	"avm1-swf-button-prefix":          libavm1swfbuttonprefix,
	"transpile-ui":                    libtranspileui,
	"ui-core":                         libuicore,
	"ui-sdk":                          libuisdk,
	"ui-sim":                          libuisim,
	"ui-surface":                      libuisurface,
}

var stdlibAutoLoadOrder = []string{
	"std",
	"str",
	"math",
	"math-base",
	"math-geo",
	"math-stats",
	"sort",
	"random",
	"fs",
	"fmt",
	"json",
	"msgpack",
	"datetime",
	"path",
	"compression",
	"compression-rle",
	"compression-huffman",
	"compression-lzw",
	"compression-ans",
	"dataprot",
	"http",
	"socket",
	"email",
	"email-smtp",
	"email-pop",
	"email-imap",
	"test",
	"thread",
	"debug",
	"cli",
	"md",
	"crypto",
	"bitwise",
	"rmqr",
	"hex",
	"image",
	"image-bmp",
	"image-ico",
	"image-ppm",
	"image-tga",
	"image-qoi",
	"gpu",
	"gpus",
	"syntax",
	"syntax-tokenize",
	"syntax-parse",
	"syntax-macros",
	"syntax-print",
	"syntaxfmt",
	"Virtual",
	"VirtualToken",
	"Virtual-Bytecode",
	"pack-utils",
	"bundle-utils",
	"ast-transform",
	"ast-analyze",
	"ast-js",
	"js-runtime",
	"estree",
	"ast-ts",
	"ast-lua",
	"ast-java",
	"ast-go",
	"ast-ir",
	"codegen-go",
	"codegen-native",
	"assembler",
	"linker",
	"bundle-ast",
	"shell",
	"transpile",
	"pack",
	"build",
	"build-includes",
	"build-ident",
	"build-ast",
	"build-analyze",
	"build-render",
	"build-render-node",
	"build-config",
	"build-imports",
	"runtime-native",
	"runtime-js",
	"runtime-codegen",
	"wasm-vm",
	"wasm-vm-runtime",
	"codegen-common",
	"sys",
	"writes",
	// "windows",
	// "windows-constants",
	// "windows-flags",
	// "windows-registry",
	// "windows-net",
	// "windows-windowing",
	// "windows-gdi",
	// "windows-kernel",
	// "windows-loader",
	// "windows-core",
	// "windows-fonts",
	// "linux",
	// "linux-constants",
	// "linux-core",
	// "linux-loader",
	// "linux-windowing",
	// "linux-libc",
	// "linux-fonts",
	"websocket",
	"p2p",
	"gui",
	// "gui-common",
	// "gui-dialogs",
	// "gui-mesh",
	// "gui-render",
	// "gui-resolution",
	// "gui-canvas",
	// "gui-accessibility",
	// "gui-clipboard",
	// "gui-filedrop",
	// "gui-audio",
	// "gui-gamepad",
	// "gui-aa",
	// "gui-draw-ops",
	// "gui-gpu-info",
	// "gui-leak-detect",
	// "gui-menus",
	// "gui-print",
	// "gui-test",
	// "gui-theme",
	// "gui-systray",
	// "gui-3dmath",
	// "gui-raster",
	// "gui-lighting",
	// "gui-web",
	// "gui-native-win",
	// "gui-native-win-present",
	// "gui-native-win-icons",
	// "gui-native-win-frame",
	// "gui-native-win-poll",
	// "gui-native-win-close",
	// "gui-native-win-ddraw",
	// "gui-native-win-d3d11",
	// "gui-native-win-vulkan",
	// "gui-native-win-vulkan-constants",
	// "gui-native-win-vulkan-init",
	// "gui-native-win-vulkan-swapchain",
	// "gui-native-win-vulkan-present",
	// "gui-native-win-opengl",
	// "gui-native-win-probe",
	// "gui-native-linux",
	// "gui-shader",
	// "gui-shader-math",
	// "gui-shader-color",
	// "gui-shader-noise",
	// "gui-shader-sdf",
	// "gui-shader-codegen",
	// "gui-draw",
	// "gui-2d",
	// "gui-color",
	// "gui-events",
	// "gui-input",
	// "gui-graph",
	// "gui-form",
	// "gui-loop",
	// "gui-thread",
	// "gui-video",
	// "gui-fonts",
	"ai",
	"audio",
	"audio-complex",
	"audio-fft",
	"audio-dsp",
	"audio-wav",
	"audio-aiff",
	"audio-au",
	"audio-raw",
	"audio-ogg",
	"video",
	"async/event-bus",
	"WLAN",
	"codecols",
	"data",
	"data-csv",
	"data-toml",
	"data-xml",
	"data-ini",
	"data-yaml",
	"web-runtime",
	"web-fetch",
	"b64-hex",
	"b64",
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
