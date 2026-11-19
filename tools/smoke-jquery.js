#!/usr/bin/env node
// smoke-jquery.js -- browser-independent smoke test for jQuery 1.0.
//
// Loads a jQuery-compatible source file in a minimal `window`/`document`
// sandbox and asserts the behaviour of the *browser-independent* utilities.
// DOM-dependent paths are intentionally NOT exercised here.
//
// Used two ways:
//   1. As the reference baseline on the original jQuery 1.0 source
//      (`node tools/smoke-jquery.js build/corpus/jquery-1.0/jquery-1.0.js`)
//   2. As the pass/fail gate on a transpiled build of the same source.
//
// Exits non-zero and prints diagnostics on the first failed assertion group.
'use strict';

const fs = require('fs');
const vm = require('vm');

const file = process.argv[2];
if (!file) {
	console.error('usage: smoke-jquery.js jquery.js');
	process.exit(2);
}
const source = fs.readFileSync(file, 'utf8');

// Minimal host environment. Anything jQuery touches during load time must
// exist; utilities asserted below must not depend on the DOM.
const documentStub = {
	getElementById: () => null,
	getElementsByTagName: () => [],
	createElement: () => ({ style: {}, setAttribute: () => {}, getElementsByTagName: () => [] }),
	body: { appendChild: () => {}, removeChild: () => {} },
	documentElement: { style: {} },
};
const windowStub = {
	document: documentStub,
	navigator: { userAgent: 'smoke-jquery/1.0' },
	location: { href: 'about:blank', protocol: 'about:' },
	setTimeout: (fn) => 0,
	clearTimeout: () => {},
	setInterval: () => 0,
	clearInterval: () => {},
};
windowStub.window = windowStub;
windowStub.self = windowStub;

const sandbox = vm.createContext(Object.assign(Object.create(null), windowStub));

try {
	vm.runInContext(source, sandbox, { filename: file, timeout: 5000 });
} catch (e) {
	console.error('LOAD FAILURE: ' + e.stack);
	process.exit(1);
}

const jQuery = sandbox.jQuery || (sandbox.window && sandbox.window.jQuery);
if (typeof jQuery !== 'function') {
	console.error('LOAD FAILURE: jQuery global was not defined');
	process.exit(1);
}
const $ = jQuery;

const failures = [];
function check(name, actual, expected) {
	const pass = Object.is(actual, expected)
		|| (Array.isArray(actual) && Array.isArray(expected) && actual.length === expected.length && actual.every((v, i) => Object.is(v, expected[i])));
	if (pass) {
		console.log('ok   ' + name);
	} else {
		failures.push(name + ': expected ' + JSON.stringify(expected) + ', got ' + JSON.stringify(actual));
		console.log('FAIL ' + name + ' (expected ' + JSON.stringify(expected) + ', got ' + JSON.stringify(actual) + ')');
	}
}

check('jQuery.fn.jquery is 1.0', $.fn && $.fn.jquery, '1.0');
check('$ is jQuery', sandbox.$ === jQuery || sandbox.window.$ === jQuery, true);

// Browser-independent API surface actually shipped in jQuery 1.0
// (verified against the source; there is no $.isArray/$.type/$.inArray/
// $.makeArray/$.isFunction/$.isEmptyObject in this release).
check('$.trim', $.trim('  hello  '), 'hello');
check('$.trim tabs/newlines', $.trim('\t\n x \r\n'), 'x');
check('$.trim empty', $.trim(''), '');

// jQuery.merge is a *unique* merge in 1.0: items from the second list that
// already exist in the first are dropped.
check('$.merge unique', $.merge([1, 2], [2, 3]), [1, 2, 3]);
check('$.merge keeps first order', $.merge(['a', 'b'], ['c', 'd']), ['a', 'b', 'c', 'd']);

check('$.grep evens', $.grep([1, 2, 3, 4], (n) => n % 2 === 0), [2, 4]);
check('$.grep inverted', $.grep([1, 2, 3, 4], (n) => n % 2 === 0, true), [1, 3]);

check('$.map double', $.map([1, 2, 3], (n) => n * 2), [2, 4, 6]);
check('$.map flattens arrays', vm.runInContext('jQuery.map([1, 2], function (n) { return [n, n]; })', sandbox), [1, 1, 2, 2]);

check('$.each over array', (() => {
	const out = [];
	$.each(['a', 'b', 'c'], function (i, v) { out.push(i + v); });
	return out.join(',');
})(), '0a,1b,2c');

check('$.extend shallow', $.extend({ a: 1, b: 2 }, { b: 3, c: 4 }).b, 3);
check('$.extend keeps lhs keys', (() => { const out = $.extend({ a: 1 }, { b: 2 }); return Object.keys(out).sort().join(','); })(), 'a,b');

check('$.isFunction absent in 1.0 (sanity)', typeof $.isFunction, 'undefined');

if (failures.length) {
	console.error('\nSMOKE FAILED: ' + failures.length + ' assertion(s)');
	process.exit(1);
}
console.log('\nsmoke OK: ' + file);
