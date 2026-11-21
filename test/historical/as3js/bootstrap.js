'use strict';

// Owned, bounded source-only bootstrap driver for the pinned AS3JS fork.
// argv: input-runtime.js output-runtime.js; both files stay in the test temp.
var fs = require('node:fs');
var path = require('node:path');
var crypto = require('node:crypto');
if (process.argv.length !== 4) throw new Error('expected two local filenames');
var input = path.basename(process.argv[2]);
var output = path.basename(process.argv[3]);
if (!/^runtime\.js$|^bootstrap-[12]\.js$/.test(input) ||
    !/^bootstrap-[12]\.js$/.test(output) || input === output) {
  throw new Error('only finite local bootstrap generations are admitted');
}
var Compiler = require('./' + input);
if (typeof Compiler !== 'function') throw new Error('compiler export missing');
var result = new Compiler().compile({
  srcPaths: [path.join(__dirname, 'src')],
  entry: 'com.mcleodgaming.as3js.Main', entryMode: 'static',
  silent: true, safeRequire: true
});
if (!result || typeof result.compiledSource !== 'string' ||
    Object.keys(result.packageSources).length < 1) {
  throw new Error('the original AS3 source did not produce a module');
}
var text = "var AS3JS = require('./lib/as3.js');\n" + result.compiledSource;
if (Buffer.byteLength(text, 'utf8') > 200000) throw new Error('generated text too large');
fs.writeFileSync(path.join(__dirname, output), text, {encoding: 'utf8', flag: 'wx'});
console.log(JSON.stringify({
  bytes: Buffer.byteLength(text, 'utf8'),
  sha256: crypto.createHash('sha256').update(text).digest('hex'),
  packages: Object.keys(result.packageSources).length,
  exportType: typeof require('./' + output)
}));
