RUN = go run -race .
LDFLAGS = -ldflags="-s -w"
INCLUDES = runners:test/runners,std.test:test/std.test,build:test/build,bitwise.test:test/bitwise.test.oak,bmp.test:test/bmp.test.oak,ico.test:test/ico.test.oak,ppm.test:test/ppm.test.oak,tga.test:test/tga.test.oak,qoi.test:test/qoi.test.oak,build.test:test/build.test.oak,class.test:test/class.test.oak,cli.test:test/cli.test.oak,compression.test:test/compression.test.oak,crypto.test:test/crypto.test.oak,dataprot.test:test/dataprot.test.oak,datetime.test:test/datetime.test.oak,debug.test:test/debug.test.oak,fmt.test:test/fmt.test.oak,go.test:test/go.test.oak,gpu.test:test/gpu.test.oak,gui.test:test/gui.test.oak,http.test:test/http.test.oak,import-ext.test:test/import-ext.test.oak,json.test:test/json.test.oak,mg.test:test/mg.test.oak,msgpack.test:test/msgpack.test.oak,macro.test:test/macro.test.oak,math.test:test/math.test.oak,md.test:test/md.test.oak,pack.test:test/pack.test.oak,path.test:test/path.test.oak,random.test:test/random.test.oak,shell.test:test/shell.test.oak,sort.test:test/sort.test.oak,std.test:test/std.test.oak,str.test:test/str.test.oak,syntax.test:test/syntax.test.oak,syscall.test:test/syscall.test.oak,thread.test:test/thread.test.oak,transpile.test:test/transpile.test.oak,vfs-bundle.test:test/vfs-bundle.test,video.test:test/video.test.oak,virtual.test:test/virtual.test,VirtualToken.test:test/VirtualToken.test,websocket.test:test/websocket.test,windows.test:test/windows.test.oak,Linux.test:test/Linux.test.oak,writes.test:test/writes.test.oak,audio.test:test/audio.test.oak

all: ci

# run the interpreter
run:
	${RUN}

# run the autoformatter (from system Oak)
fmt:
	magnolia fmt --changes --fix
f: fmt

# run Go tests
tests:
	go test -race .
t: tests

# run Oak tests
test-oak:
	${RUN} test/main.oak
tk: test-oak

# run oak build tests
test-bundle:
	${RUN} build --entry test/main.oak --output /tmp/oak-test.oak --include ${INCLUDES}
	${RUN} /tmp/oak-test.oak

# run oak pack tests
test-pack:
	${RUN} pack --entry test/main.oak --output /tmp/oak-pack --include ${INCLUDES}
	/tmp/oak-pack

# run oak build --web tests
test-js:
	${RUN} build --entry test/main.oak --output /tmp/oak-test.js --web --include ${INCLUDES}
	node /tmp/oak-test.js

# run oak build --wasm tests
test-wasm:
	${RUN} build --entry test/main.oak --output /tmp/oak-test.wat --wasm --include ${INCLUDES}
	wasm-interp /tmp/oak-test.wat

# build for a specific GOOS target
build-%:
	GOOS=$* go build ${LDFLAGS} -o magnolia-$* .

# build for all OS targets
build: build-linux build-darwin build-windows build-openbsd

# build Oak sources for the website
site:
	magnolia build --entry www/src/app.js.oak --output www/static/js/bundle.js --web
	magnolia build --entry www/src/highlight.js.oak --output www/static/js/highlight.js --web

# build Oak source for the website on file change, using entr
site-w:
	ls www/src/app.js.oak | entr -cr make site

# generate static site pages
site-gen:
	magnolia www/src/gen.oak

# install as "magnolia" binary
install:
	cp tools/oak.vim ~/.vim/syntax/oak.vim
	go build ${LDFLAGS} -o ${GOPATH}/bin/magnolia

# ci in travis
ci: tests test-oak test-bundle test-pack
