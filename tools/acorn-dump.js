#!/usr/bin/env node
// acorn-dump.js -- dump the ESTree AST of a JS file as JSON.
//
// This is the reference ("ground truth") front-end used to ablate the
// Oak-native JS suite (lib/js-tokenize.oak + lib/js-parse.oak). The JSON
// produced here is what lib/estree.oak (Oak) ingests, and what
// tools/ast-differ.js compares against the Oak-native output.
//
// Usage:
//   node tools/acorn-dump.js [--ecmaVersion 3|5|6|...|latest] [--loc] file.js
//
// Output: JSON AST on stdout. Locations are omitted unless --loc is passed.
'use strict';

const fs = require('fs');
const path = require('path');

let acorn;
try {
	acorn = require('acorn');
} catch (e) {
	try {
		// allow running with acorn resolved from the repo-local ablation dir
		acorn = require(path.join(__dirname, '..', 'build', 'node_modules', 'acorn'));
	} catch (e2) {
		console.error('acorn is not installed. Run: npm install --prefix build acorn@8.19.0');
		process.exit(2);
	}
}

const argv = process.argv.slice(2);
let ecmaVersion = 'latest';
let withLoc = false;
const files = [];
for (let i = 0; i < argv.length; i++) {
	if (argv[i] === '--ecmaVersion') { ecmaVersion = argv[++i]; continue; }
	if (argv[i] === '--loc') { withLoc = true; continue; }
	files.push(argv[i]);
}
if (files.length !== 1) {
	console.error('usage: acorn-dump.js [--ecmaVersion N|latest] [--loc] file.js');
	process.exit(2);
}

const source = fs.readFileSync(files[0], 'utf8');
const version = ecmaVersion === 'latest' ? 'latest' : parseInt(ecmaVersion, 10);
const ast = acorn.parse(source, {
	ecmaVersion: version,
	allowHashBang: true,
	locations: false,
	ranges: false,
	onComment: undefined,
	program: undefined,
});

// Normalize: strip volatile fields so snapshots are stable.
(function strip(node) {
	if (!node || typeof node !== 'object') return;
	delete node.start;
	delete node.end;
	delete node.loc;
	delete node.range;
	for (const key of Object.keys(node)) {
		const value = node[key];
		if (Array.isArray(value)) {
			value.forEach(strip);
		} else if (value && typeof value === 'object') {
			strip(value);
		}
	}
})(ast);

process.stdout.write(JSON.stringify(ast, null, 2) + '\n');
