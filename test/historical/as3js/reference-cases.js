'use strict';
// Owned finite probe for the *external*, hash-pinned AS3JS source-to-Node
// comparator. Run only in a temporary, network-denied, bounded Node child.
// This is NOT AVM2, Flash behavior, real browser painting, or an AS3 oracle.
var assert = require('node:assert/strict');
var fs = require('node:fs');
var os = require('node:os');
var path = require('node:path');
var crypto = require('node:crypto');
var Compiler = require('./index.js');
// These exact two .as files are persisted under test/fixtures/ so a future
// Magnolia source runtime can consume the SAME source, not a copied example.
// trimEnd removes the one source-file newline, preserving the rawPackages
// input recorded by this reference gate.
var sourceItem = fs.readFileSync(path.join(__dirname, 'Item.as'), 'utf8').trimEnd();
var sourceMain = fs.readFileSync(path.join(__dirname, 'Main.as'), 'utf8').trimEnd();
function compile(s) {
  return new Compiler().compile({
    rawPackages: [sourceItem, s || sourceMain], entry: 'reference.Main',
    entryMode: 'instance', silent: true
  });
}
var generated = compile();
assert.ok(generated && typeof generated.compiledSource === 'string');
assert.ok(generated.compiledSource.length < 40000);
var mod = {exports: null};
new Function('AS3JS', 'module', generated.compiledSource)(Compiler.loader, mod);
var app = mod.exports;
var item = app.item();
var observed = {
  initial: app.initial, values: app.values(), regex: app.regex(),
  quotient: app.quotient(),
  castString: app.cast('hello'), castNonString: app.cast(42),
  integerTrue: app.isInt(3), integerFraction: app.isInt(3.5),
  itemTrue: app.isItem(item), itemFalse: app.isItem({})
};
assert.deepEqual(observed, {
  initial: 17, values: [1, 'two', 3], regex: true, quotient: 4,
  castString: 'hello', castNonString: null,
  integerTrue: true, integerFraction: false,
  itemTrue: true, itemFalse: false
});
var unsupported = sourceMain.replace('return value is Item;',
                                     'return value + 1 is Item;');
var unsupportedError;
try { compile(unsupported); } catch (error) { unsupportedError = error; }
assert.ok(unsupportedError instanceof Error);
assert.match(unsupportedError.message,
             /Unsupported AS3 'as' or 'is' type expression/);
var malformed = [
  'package diagnostics {',
  '  public class Broken {',
  '    public function broken():void {',
  '      var value:int = 1;'
].join('\n');
var rawError;
try { new Compiler().compile({rawPackages: [malformed], silent: true}); }
catch (error) { rawError = error; }
assert.ok(rawError instanceof Error);
assert.match(rawError.message, /line 4, column 25: Error, no closing '\}' found/);
var folder = fs.mkdtempSync(path.join(os.tmpdir(), 'as3js-reference-'));
var fileError;
try {
  fs.writeFileSync(path.join(folder, 'Broken.as'), malformed);
  try { new Compiler().compile({srcPaths: [folder], silent: true}); }
  catch (error) { fileError = error; }
  assert.ok(fileError instanceof Error);
  assert.match(fileError.message,
               /Broken\.as:4:25: Error, no closing '\}' found/);
} finally { fs.rmSync(folder, {recursive: true, force: true}); }
console.log(JSON.stringify({
  profile: 'AS3JS source-to-Node reference only', observed: observed,
  unsupported: unsupportedError.message,
  malformedRaw: rawError.message,
  malformedFile: fileError.message.replace(/^.*Broken\.as:/, 'Broken.as:'),
  compiledSourceBytes: Buffer.byteLength(generated.compiledSource, 'utf8'),
  compiledSourceSha256: crypto.createHash('sha256').update(generated.compiledSource).digest('hex'),
  AVM2Executed: false, MagnoliaAs3Executed: false
}));
