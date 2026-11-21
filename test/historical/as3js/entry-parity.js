'use strict';

// Owned comparator for the pinned MIT AS3JS source-only bootstrap. The oracle
// expects exactly one JSON line; it does NOT establish Flash/AVM semantics.
var Compiler = require('./index.js');
var loader = Compiler.loader;
var source = [
  'package api { public class Main {',
  '  public function Main() {}',
  '  public static function label():String { return "static-entry"; }',
  '  public function value():Number { return 12 / 3; }',
  '} }'
].join('\n');

function result(mode) {
  var generated = new Compiler().compile({
    rawPackages: [source], entry: 'api.Main', entryMode: mode, silent: true
  });
  var mod = {exports: null};
  new Function('AS3JS', 'module', generated.compiledSource)(loader, mod);
  if (mode === 'static') return {label: mod.exports.label(), type: typeof mod.exports};
  return {value: mod.exports.value(), label: mod.exports.constructor.label(),
    type: typeof mod.exports};
}

console.log(JSON.stringify({
  version: require('./package.json').version,
  compilerType: typeof Compiler,
  loaderType: typeof loader.load,
  globalPolluted: Object.prototype.hasOwnProperty.call(global, 'AS3JS'),
  compilerKeys: Object.getOwnPropertyNames(Compiler.prototype).sort(),
  instance: result('instance'), static: result('static')
}));
