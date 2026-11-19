#!/usr/bin/env node
/* Offline DOM/control smoke test for the exported Fliplay-decoder frame HTML.
 * No browser, network, original class files or external Node modules needed.
 * The Python oracle separately checks all PNG pixels against JDK/Pillow.
 * Usage: node tools/fliplay-viewer-oracle.js /path/to/exported.html
 */
'use strict';
const fs = require('fs');
const vm = require('vm');
const assert = require('assert');
if (process.argv.length !== 3) throw new Error('pass one exported HTML path');
const html = fs.readFileSync(process.argv[2], 'utf8');
const frameBlock = html.match(/<script type="application\/json" id="frame-data">([\s\S]*?)<\/script>/);
const codeBlock = html.match(/<script>\s*([\s\S]*?)<\/script>/);
assert(frameBlock && codeBlock, 'offline frame JSON and inline controls script required');
const data = JSON.parse(frameBlock[1]);
assert.strictEqual(data.frames.length, 60);
assert.strictEqual(data.hashes.length, 60);
assert(data.frames.every(url => url.startsWith('data:image/png;base64,')));

const elements = new Map();
for (const id of ['screen', 'seek', 'play', 'loop', 'speed', 'counter',
                  'elapsed', 'frame-hash', 'back', 'next']) {
  const events = new Map();
  elements.set('#' + id, {
    value: id === 'speed' ? '1' : '0', src: '', textContent: '',
    attrs: {}, style: {setProperty() {}},
    setAttribute(key, value) { this.attrs[key] = value; },
    addEventListener(name, handler) { events.set(name, handler); },
    fire(name) { assert(events.has(name), name); events.get(name)({}); }
  });
}
elements.set('#frame-data', {textContent: JSON.stringify(data)});
const documentEvents = new Map();
const document = {
  querySelector(id) { assert(elements.has(id), id); return elements.get(id); },
  addEventListener(name, handler) { documentEvents.set(name, handler); },
  hidden: false
};
let tick, token = 0;
vm.runInNewContext(codeBlock[1], {
  document,
  requestAnimationFrame(handler) { tick = handler; return ++token; },
  cancelAnimationFrame() { tick = null; }
}, {filename: 'fliplay-exported-viewer-inline.js', timeout: 2000});
const el = id => elements.get('#' + id);
assert.strictEqual(el('screen').src, data.frames[0]);
assert.strictEqual(el('frame-hash').textContent, data.hashes[0]);
el('play').fire('click');
assert.strictEqual(el('play').attrs['aria-label'], 'Pause');
tick(1000); tick(1071);
assert.strictEqual(el('screen').src, data.frames[1]);
el('next').fire('click');
assert.strictEqual(el('screen').src, data.frames[2]);
el('back').fire('click');
assert.strictEqual(el('screen').src, data.frames[1]);
el('seek').value = '59'; el('seek').fire('input');
assert.strictEqual(el('screen').src, data.frames[59]);
el('loop').fire('click');
assert.strictEqual(el('loop').attrs['aria-pressed'], 'false');
tick(2000); tick(2071);
assert.strictEqual(el('play').attrs['aria-label'], 'Play');
el('play').fire('click');
assert.strictEqual(el('screen').src, data.frames[0]);
documentEvents.get('keydown')({code: 'ArrowRight', target: null,
                                 preventDefault() {}});
assert.strictEqual(el('screen').src, data.frames[1]);
console.log('PASS offline viewer: 60 embedded data frames, playback, steps, scrub, end, restart, keyboard');
