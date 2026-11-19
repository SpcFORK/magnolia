#!/usr/bin/env node
// ast-histogram.js -- node-type histogram of a JS file, as parsed by acorn.
//
// Used to scope the Oak-native parser (lib/js-parse.oak) to the constructs
// that real code (jQuery 1.0, three.js r58, ...) actually uses.
//
// Usage: node tools/ast-histogram.js [--ecmaVersion N] file.js [file.js ...]
'use strict';

const fs = require('fs');
const path = require('path');

let acorn;
try {
	acorn = require('acorn');
} catch (e) {
	acorn = require(path.join(__dirname, '..', 'build', 'node_modules', 'acorn'));
}

const argv = process.argv.slice(2);
let ecmaVersion = 'latest';
const files = [];
for (let i = 0; i < argv.length; i++) {
	if (argv[i] === '--ecmaVersion') { ecmaVersion = argv[++i]; continue; }
	files.push(argv[i]);
}
const version = ecmaVersion === 'latest' ? 'latest' : parseInt(ecmaVersion, 10);

const totals = new Map();
const perFile = [];
let grandTotal = 0;

for (const file of files) {
	const source = fs.readFileSync(file, 'utf8');
	const counts = new Map();
	let nodes = 0;
	try {
		const ast = acorn.parse(source, { ecmaVersion: version, allowHashBang: true });
		(function walk(node) {
			if (!node || typeof node !== 'object') return;
			if (typeof node.type === 'string') {
				nodes++;
				counts.set(node.type, (counts.get(node.type) || 0) + 1);
			}
			for (const key of Object.keys(node)) {
				if (key === 'leadingComments' || key === 'trailingComments') continue;
				const value = node[key];
				if (Array.isArray(value)) value.forEach(walk);
				else if (value && typeof value === 'object') walk(value);
			}
		})(ast);
	} catch (e) {
		console.log('PARSE FAILURE ' + file + ': ' + e.message);
		process.exitCode = 1;
		continue;
	}
	grandTotal += nodes;
	perFile.push({ file, nodes });
	for (const [type, count] of counts) {
		totals.set(type, (totals.get(type) || 0) + count);
	}
}

const sorted = [...totals.entries()].sort((a, b) => b[1] - a[1] || a[0].localeCompare(b[0]));
console.log('# ESTree node histogram (acorn, ecmaVersion=' + ecmaVersion + ')');
for (const { file, nodes } of perFile) {
	console.log('#   ' + file + ': ' + nodes + ' nodes');
}
console.log('# total: ' + grandTotal + ' nodes across ' + perFile.length + ' file(s)');
console.log('');
console.log('| Node type | Count |');
console.log('|---|---|');
for (const [type, count] of sorted) {
	console.log('| ' + type + ' | ' + count + ' |');
}
