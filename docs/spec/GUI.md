# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `async/event-bus`

- `_eventKeyCache` · `{}`
### `_eventKey(name)`

### `cs EventBus()`

> returns `:object`

### `create()`

## Module: `gui-2d`

- `threadLib` · `import(...)`
- `PI` · `3.141592653589793`
### `_degToRad(deg)`

### `Vec2(x, y)`

> returns `:object`

### `Vec4(x, y, w, h)`

> returns `:object`

### `Rect2(x, y, width, height)`

> returns `:object`

### `vec2Add(a, b)`

### `vec2Sub(a, b)`

### `vec2Scale(v, s)`

### `vec2Dot(a, b)`

### `vec2Len(v)`

### `vec2Normalize(v)`

### `rectTranslate(rect, dx, dy)`

### `rectContains(rect, point)`

> returns `:bool`

### `rectIntersects(a, b)`

> returns `:bool`

### `Circle2D(cx, cy, r)`

> returns `:object`

### `pointInRect2D(point, rect)`

> returns `:bool`

### `pointInCircle2D(point, circle)`

### `circleIntersects(a, b)`

### `circleRectIntersects(circle, rect)`

### `vec2Distance2(a, b)`

### `vec2Distance(a, b)`

### `vec2Lerp(a, b, t)`

### `rectUnion(a, b)`

### `rectClampPoint(rect, point)`

### `lineSegmentIntersect2D(a1, a2, b1, b2)`

> returns `:object`

### `rayRectIntersect2D(origin, dir, rect)`

> returns `:object`

### `sweptRectIntersect2D(movingRect, vel, targetRect)`

### `rectOverlapDepth2D(a, b)`

### `_rectCollidesAnySub(rect, colliders, idx)`

> returns `:bool`

### `rectCollidesAny(rect, colliders)`

### `_circleCollidesAnySub(circle, colliders, idx)`

> returns `:bool`

### `circleCollidesAny(circle, colliders)`

### `resolveRectMove(player, nextPos, colliders)`

### `resolveCircleMove(player, nextPos, colliders)`

### `Transform2D(options)`

> returns `:object`

### `applyTransform2D(point, transform)`

### `Camera2D(options)`

> returns `:object`

### `worldToScreen2D(point, camera, window)`

### `screenToWorld2D(point, camera, window)`

### `_drawPolylineSub(deps, window, points, i, color, closed)`

> returns `:object`

### `drawPolyline2D(deps, window, points, color, closed)`

> returns `:object`

### `drawRect2D(deps, window, x, y, width, height, color, filled, borderColor)`

### `_drawCircleOutlineSub(deps, window, cx, cy, x, y, cr, sr, segs, i, color)`

> returns `:object`

### `_drawCircleFilledSub(deps, window, cx, cy, r, y, color)`

> returns `:object`

### `drawCircle2D(deps, window, cx, cy, radius, color, filled, borderColor)`

> returns `:object`

### `drawPolygon2D(deps, window, points, color, filled, borderColor)`

> returns `:object`

### `Element(deps, bounds)`

> returns `:object`

### `drawGrid2D(deps, window, spacing, color, originX, originY)`

> returns `:object`

### `_drawEllipseOutlineSub(deps, window, cx, cy, px, py, cr, sr, segs, i, rx, ry, color)`

> returns `:object`

### `_drawEllipseFilledSub(deps, window, cx, cy, rx, ry, y, color)`

> returns `:object`

### `drawEllipse2D(deps, window, cx, cy, rx, ry, color, filled, borderColor)`

> returns `:object`

### `_drawArcSub(deps, window, cx, cy, radius, curAngle, endAngle, stepAngle, color)`

### `drawArc2D(deps, window, cx, cy, radius, startAngle, endAngle, color)`

> returns `:object`

### `drawTriangle2D(deps, window, x1, y1, x2, y2, x3, y3, color, filled, borderColor)`

### `_drawRoundedRectCorner(deps, window, cx, cy, r, startDeg, endDeg, color)`

### `_drawRoundedRectFilledRows(deps, window, x, y, width, height, r, row, color)`

> returns `:object`

### `drawRoundedRect2D(deps, window, x, y, width, height, radius, color, filled, borderColor)`

### `_starVertices(cx, cy, outerR, innerR, points, i, out)`

### `drawStar2D(deps, window, cx, cy, outerR, innerR, points, color, filled, borderColor)`

### `_bezierLerp(a, b, t)`

### `_bezierQuadPoint(p0, p1, p2, t)`

### `_bezierCubicPoint(p0, p1, p2, p3, t)`

### `_drawBezierSub(deps, window, evalFn, prev, i, steps, color)`

### `drawBezier2D(deps, window, points, color, steps)`

### `_drawRingRow(deps, window, cx, cy, outerR, innerR, y, color)`

### `drawRing2D(deps, window, cx, cy, outerR, innerR, color)`

> returns `:object`

### `drawCross2D(deps, window, cx, cy, size, thickness, color, filled)`

### `drawDiamond2D(deps, window, cx, cy, width, height, color, filled)`

### `drawArrow2D(deps, window, x1, y1, x2, y2, color, headSize)`

> returns `:object`

### `drawCapsule2D(deps, window, cx, cy, width, height, color, filled)`

### `drawSector2D(deps, window, cx, cy, radius, startAngle, endAngle, color, filled)`

> returns `:object`

### `_regularPolyVerts(cx, cy, radius, sides, i, out)`

### `drawRegularPolygon2D(deps, window, cx, cy, radius, sides, color, filled)`

### `_drawSpiralSub(deps, window, cx, cy, r, growth, angle, endAngle, step, prevX, prevY, color)`

### `drawSpiral2D(deps, window, cx, cy, startRadius, growth, turns, color)`

### `drawThickLine2D(deps, window, x1, y1, x2, y2, thickness, color)`

### `_drawDashedLineSub(deps, window, x1, y1, ux, uy, totalLen, pos, dashLen, gapLen, color)`

> returns `:object`

### `drawDashedLine2D(deps, window, x1, y1, x2, y2, color, dashLen, gapLen)`

## Module: `gui-3dmath`

- `threadLib` · `import(...)`
- `PI` · `3.141592653589793`
### `degToRad(deg)`

### `Vec3(x, y, z)`

> returns `:object`

### `_rotateX(v, radians)`

> returns `:object`

### `_rotateY(v, radians)`

> returns `:object`

### `_rotateZ(v, radians)`

> returns `:object`

### `transformPoint(v, transform)`

> returns `:object`

### `projectPoint(window, p, camera)`

> returns `:object`

### `transformVertices(vertices, transform, i, out)`

### `_transformVertex(v, scale, tx, ty, tz, cx, sx, cy, sy, cz, sz)`

> returns `:object`

### `transformVerticesBatch(meshes, transforms, i, out)`

### `transformAndProjectVertices(vertices, transform, projParams)`

### `Mat4Identity()`

> returns `:list`

### `Mat4Translate(x, y, z)`

> returns `:list`

### `Mat4Scale(x, y, z)`

> returns `:list`

### `Mat4RotateX(radians)`

> returns `:list`

### `Mat4RotateY(radians)`

> returns `:list`

### `Mat4RotateZ(radians)`

> returns `:list`

### `Mat4Multiply(a, b)`

### `Mat4TransformPoint(m, v)`

### `Quat(x, y, z, w)`

> returns `:object`

### `QuatFromAxisAngle(axis, radians)`

### `QuatMultiply(a, b)`

### `QuatNormalize(q)`

### `QuatRotateVector(q, v)`

### `QuatToMat4(q)`

> returns `:list`

### `QuatSlerp(a, b, t)`

### `Vec3Add(a, b)`

### `Vec3Sub(a, b)`

### `Vec3Scale(v, s)`

### `Vec3Dot(a, b)`

### `Vec3Cross(a, b)`

### `Vec3Length(v)`

### `Vec3Normalize(v)`

### `Vec3Distance(a, b)`

### `Vec3Lerp(a, b, t)`

### `Vec3Negate(v)`

### `Vec3Reflect(v, n)`

### `AABB3(minX, minY, minZ, maxX, maxY, maxZ)`

> returns `:object`

### `Sphere3D(cx, cy, cz, r)`

> returns `:object`

### `Plane3D(nx, ny, nz, d)`

> returns `:object`

### `pointInAABB3(point, box)`

> returns `:bool`

### `pointInSphere3D(point, sphere)`

### `aabb3Intersects(a, b)`

> returns `:bool`

### `sphere3DIntersects(a, b)`

### `sphereAABB3Intersects(sphere, box)`

### `planePointDistance(plane, point)`

### `planeClassifyPoint(plane, point)`

> returns `:atom`

### `raySphere3DIntersect(origin, dir, sphere)`

> returns `:object`

### `rayAABB3Intersect(origin, dir, box)`

> returns `:object`

### `rayPlane3DIntersect(origin, dir, plane)`

> returns `:object`

### `aabb3ClosestPoint(box, point)`

### `aabb3Union(a, b)`

### `aabb3Center(box)`

### `aabb3HalfExtents(box)`

### `parallelTransformVertices(vertices, transform, numWorkers)`

## Module: `gui-aa`

- `guiColor` · `import(...)`
- `threadLib` · `import(...)`
- `_OK` · `{1 entries}`
### `_abs(x)`

> returns `:int`

### `_floor(x)`

### `_fpart(x)`

### `_rfpart(x)`

> returns `:int`

### `_ipart(x)`

### `_round(x)`

### `_swap(a, b)`

> returns `:list`

### `_min(a, b)`

### `_max(a, b)`

### `_blendColor(fg, alpha, bg)`

### `_smoothstep(edge0, edge1, x)`

### `drawLineAA(deps, window, x0, y0, x1, y1, color, bgColor)`

### `drawCircleFilledAA(deps, window, cx, cy, radius, color, bgColor)`

### `drawCircleOutlineAA(deps, window, cx, cy, radius, color, bgColor)`

### `drawEllipseFilledAA(deps, window, cx, cy, rx, ry, color, bgColor)`

### `drawRoundedRectFilledAA(deps, window, x, y, width, height, radius, color, bgColor)`

### `_edgeDistSigned(px, py, ax, ay, bx, by)`

> returns `:int`

### `_triMin3(a, b, c)`

### `_triMax3(a, b, c)`

### `drawTriangleFilledAA(deps, window, p0, p1, p2, color, bgColor)`

### `drawCircleFilledAAParallel(deps, window, cx, cy, radius, color, bgColor, numWorkers)`

## Module: `gui-accessibility`

- `sys` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `ROLE_SYSTEM_TITLEBAR` · `1`
- `ROLE_SYSTEM_MENUBAR` · `2`
- `ROLE_SYSTEM_SCROLLBAR` · `3`
- `ROLE_SYSTEM_GRIP` · `4`
- `ROLE_SYSTEM_SOUND` · `5`
- `ROLE_SYSTEM_CURSOR` · `6`
- `ROLE_SYSTEM_CARET` · `7`
- `ROLE_SYSTEM_ALERT` · `8`
- `ROLE_SYSTEM_WINDOW` · `9`
- `ROLE_SYSTEM_CLIENT` · `10`
- `ROLE_SYSTEM_MENUPOPUP` · `11`
- `ROLE_SYSTEM_MENUITEM` · `12`
- `ROLE_SYSTEM_TOOLTIP` · `13`
- `ROLE_SYSTEM_APPLICATION` · `14`
- `ROLE_SYSTEM_DOCUMENT` · `15`
- `ROLE_SYSTEM_PANE` · `16`
- `ROLE_SYSTEM_CHART` · `17`
- `ROLE_SYSTEM_DIALOG` · `18`
- `ROLE_SYSTEM_BORDER` · `19`
- `ROLE_SYSTEM_GROUPING` · `20`
- `ROLE_SYSTEM_SEPARATOR` · `21`
- `ROLE_SYSTEM_TOOLBAR` · `22`
- `ROLE_SYSTEM_STATUSBAR` · `23`
- `ROLE_SYSTEM_TABLE` · `24`
- `ROLE_SYSTEM_COLUMNHEADER` · `25`
- `ROLE_SYSTEM_ROWHEADER` · `26`
- `ROLE_SYSTEM_COLUMN` · `27`
- `ROLE_SYSTEM_ROW` · `28`
- `ROLE_SYSTEM_CELL` · `29`
- `ROLE_SYSTEM_LINK` · `30`
- `ROLE_SYSTEM_HELPBALLOON` · `31`
- `ROLE_SYSTEM_CHARACTER` · `32`
- `ROLE_SYSTEM_LIST` · `33`
- `ROLE_SYSTEM_LISTITEM` · `34`
- `ROLE_SYSTEM_OUTLINE` · `35`
- `ROLE_SYSTEM_OUTLINEITEM` · `36`
- `ROLE_SYSTEM_PAGETAB` · `37`
- `ROLE_SYSTEM_PROPERTYPAGE` · `38`
- `ROLE_SYSTEM_INDICATOR` · `39`
- `ROLE_SYSTEM_GRAPHIC` · `40`
- `ROLE_SYSTEM_STATICTEXT` · `41`
- `ROLE_SYSTEM_TEXT` · `42`
- `ROLE_SYSTEM_PUSHBUTTON` · `43`
- `ROLE_SYSTEM_CHECKBUTTON` · `44`
- `ROLE_SYSTEM_RADIOBUTTON` · `45`
- `ROLE_SYSTEM_COMBOBOX` · `46`
- `ROLE_SYSTEM_DROPLIST` · `47`
- `ROLE_SYSTEM_PROGRESSBAR` · `48`
- `ROLE_SYSTEM_DIAL` · `49`
- `ROLE_SYSTEM_HOTKEYFIELD` · `50`
- `ROLE_SYSTEM_SLIDER` · `51`
- `ROLE_SYSTEM_SPINBUTTON` · `52`
- `ROLE_SYSTEM_DIAGRAM` · `53`
- `ROLE_SYSTEM_ANIMATION` · `54`
- `ROLE_SYSTEM_EQUATION` · `55`
- `ROLE_SYSTEM_BUTTONDROPDOWN` · `56`
- `ROLE_SYSTEM_BUTTONMENU` · `57`
- `ROLE_SYSTEM_BUTTONDROPDOWNGRID` · `58`
- `ROLE_SYSTEM_WHITESPACE` · `59`
- `ROLE_SYSTEM_PAGETABLIST` · `60`
- `ROLE_SYSTEM_CLOCK` · `61`
- `ROLE_SYSTEM_SPLITBUTTON` · `62`
- `STATE_SYSTEM_NORMAL` · `0`
- `STATE_SYSTEM_UNAVAILABLE` · `1`
- `STATE_SYSTEM_SELECTED` · `2`
- `STATE_SYSTEM_FOCUSED` · `4`
- `STATE_SYSTEM_PRESSED` · `8`
- `STATE_SYSTEM_CHECKED` · `16`
- `STATE_SYSTEM_MIXED` · `32`
- `STATE_SYSTEM_READONLY` · `64`
- `STATE_SYSTEM_HOTTRACKED` · `128`
- `STATE_SYSTEM_DEFAULT` · `256`
- `STATE_SYSTEM_EXPANDED` · `512`
- `STATE_SYSTEM_COLLAPSED` · `1024`
- `STATE_SYSTEM_BUSY` · `2048`
- `STATE_SYSTEM_INVISIBLE` · `32768`
- `STATE_SYSTEM_OFFSCREEN` · `65536`
- `STATE_SYSTEM_SIZEABLE` · `131072`
- `STATE_SYSTEM_MOVEABLE` · `262144`
- `STATE_SYSTEM_FOCUSABLE` · `1048576`
- `STATE_SYSTEM_SELECTABLE` · `2097152`
- `STATE_SYSTEM_LINKED` · `4194304`
- `STATE_SYSTEM_TRAVERSED` · `8388608`
- `STATE_SYSTEM_MULTISELECTABLE` · `16777216`
- `STATE_SYSTEM_HASPOPUP` · `1073741824`
- `EVENT_SYSTEM_SOUND` · `1`
- `EVENT_SYSTEM_ALERT` · `2`
- `EVENT_SYSTEM_FOREGROUND` · `3`
- `EVENT_SYSTEM_MENUSTART` · `4`
- `EVENT_SYSTEM_MENUEND` · `5`
- `EVENT_SYSTEM_MENUPOPUPSTART` · `6`
- `EVENT_SYSTEM_MENUPOPUPEND` · `7`
- `EVENT_SYSTEM_CAPTURESTART` · `8`
- `EVENT_SYSTEM_CAPTUREEND` · `9`
- `EVENT_SYSTEM_MOVESIZESTART` · `10`
- `EVENT_SYSTEM_MOVESIZEEND` · `11`
- `EVENT_SYSTEM_CONTEXTHELPSTART` · `12`
- `EVENT_SYSTEM_CONTEXTHELPEND` · `13`
- `EVENT_SYSTEM_DRAGDROPSTART` · `14`
- `EVENT_SYSTEM_DRAGDROPEND` · `15`
- `EVENT_SYSTEM_DIALOGSTART` · `16`
- `EVENT_SYSTEM_DIALOGEND` · `17`
- `EVENT_SYSTEM_SCROLLINGSTART` · `18`
- `EVENT_SYSTEM_SCROLLINGEND` · `19`
- `EVENT_SYSTEM_SWITCHSTART` · `20`
- `EVENT_SYSTEM_SWITCHEND` · `21`
- `EVENT_OBJECT_CREATE` · `32768`
- `EVENT_OBJECT_DESTROY` · `32769`
- `EVENT_OBJECT_SHOW` · `32770`
- `EVENT_OBJECT_HIDE` · `32771`
- `EVENT_OBJECT_REORDER` · `32772`
- `EVENT_OBJECT_FOCUS` · `32773`
- `EVENT_OBJECT_SELECTION` · `32774`
- `EVENT_OBJECT_SELECTIONADD` · `32775`
- `EVENT_OBJECT_SELECTIONREMOVE` · `32776`
- `EVENT_OBJECT_SELECTIONWITHIN` · `32777`
- `EVENT_OBJECT_STATECHANGE` · `32778`
- `EVENT_OBJECT_LOCATIONCHANGE` · `32779`
- `EVENT_OBJECT_NAMECHANGE` · `32780`
- `EVENT_OBJECT_DESCRIPTIONCHANGE` · `32781`
- `EVENT_OBJECT_VALUECHANGE` · `32782`
- `EVENT_OBJECT_PARENTCHANGE` · `32783`
- `EVENT_OBJECT_HELPCHANGE` · `32784`
- `EVENT_OBJECT_DEFACTIONCHANGE` · `32785`
- `EVENT_OBJECT_ACCELERATORCHANGE` · `32786`
- `EVENT_OBJECT_INVOKED` · `32787`
- `EVENT_OBJECT_TEXTSELECTIONCHANGED` · `32788`
- `EVENT_OBJECT_CONTENTSCROLLED` · `32789`
- `EVENT_OBJECT_LIVEREGIONCHANGED` · `32793`
- `WM_GETOBJECT` · `61`
- `OBJID_SELF` · `0`
- `OBJID_WINDOW` · `0`
- `OBJID_CLIENT` — constant
- `OBJID_NATIVEOM` — constant
- `_UIA_UIAROOT_OBJECTID` — constant
- `_oleaccProcCache` · `{}`
- `Oleacc` · `'oleacc.dll'`
### `oleacc(symbol, args...)`

### `accNode(id, name, role, bounds, state, children)`

> returns `:object`

### `accTree(window)`

### `accRegister(window, id, name, role, bounds, state)`

### `accUnregister(window, id)`

### `accSetName(window, id, name)`

### `accSetState(window, id, state)`

### `accSetValue(window, id, value)`

### `accSetBounds(window, id, bounds)`

### `accSetDescription(window, id, desc)`

### `accSetDefaultAction(window, id, action)`

### `accFocus(window, id)`

### `accSelection(window, id)`

### `notifyAccEvent(window, event, childId)`

### `enableAccessibility(window)`

> returns `:bool`

### `accRegisterButton(window, id, label, x, y, w, h)`

### `accRegisterCheckbox(window, id, label, x, y, w, h, checked?)`

### `accRegisterRadio(window, id, label, x, y, w, h, selected?)`

### `accRegisterTextField(window, id, label, x, y, w, h, value)`

### `accRegisterSlider(window, id, label, x, y, w, h, value)`

### `accRegisterProgressBar(window, id, label, x, y, w, h, value)`

### `accRegisterTab(window, id, label, x, y, w, h, selected?)`

### `accRegisterListItem(window, id, label, x, y, w, h, selected?)`

### `accRegisterLink(window, id, label, x, y, w, h)`

### `accRegisterStaticText(window, id, label, x, y, w, h)`

### `accRegisterGroup(window, id, label, x, y, w, h)`

### `accRegisterTreeItem(window, id, label, x, y, w, h, expanded?)`

### `accRegisterTable(window, id, label, x, y, w, h)`

### `accRegisterTableCell(window, id, label, x, y, w, h)`

### `accRegisterDropdown(window, id, label, x, y, w, h, expanded?)`

### `accAnnounce(window, id, text)`

### `accDump(window)`

### `accVerify(window)`

> returns `:object`

### `accRoleName(role)`

> returns `:string`

### `accStateName(state)`

> returns `:string`

## Module: `gui-audio`

- `guiThread` · `import(...)`
### `spatialAudioSource(id, x, y, volume)`

> returns `:object`

### `spatialAudioUpdate(source, listenerX, listenerY, viewW, viewH, maxDist)`

### `spatialApplyToSamples(samples, pan, gain)`

### `spatialMixSources(sources, listenerX, listenerY, viewW, viewH, maxDist, bufLen)`

### `_escapePs(s)`

### `setAppVolumeName(displayName)`

### `getSystemVolume()`

### `setSystemVolume(level)`

### `enableMediaTransportControls(options)`

### `setMediaPlaybackStatus(status)`

### `updateMediaInfo(title, artist, albumTitle)`

## Module: `gui-canvas`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `linux` · `import(...)`
- `_OK` · `{1 entries}`
- `SRCCOPY` · `13369376`
- `_parentMap` · `{}`
- `_nextId` · `{1 entries}`
### `_getParent(canvas)`

### `create(parentWindow, options)`

### `_ensureCanvasSurface(canvas)`

### `beginCanvas(canvas)`

### `endCanvas(canvas)`

### `_sortedVisibleCanvases(window)`

### `_compositeCanvasWindows(window, canvas)`

### `_compositeCanvasLinux(window, canvas)`

### `_compositeCanvasWeb(window, canvas)`

### `compositeAll(window)`

### `move(canvas, x, y)`

### `resize(canvas, w, h)`

### `setVisible(canvas, vis)`

### `setZIndex(canvas, z)`

### `setOpacity(canvas, alpha)`

### `setTransparentColor(canvas, color)`

### `_releaseCanvasSurface(canvas)`

### `destroy(canvas)`

### `isCanvas?(obj)`

### `canvases(window)`

### `canvasCount(window)`

### `hitTest?(canvas, px, py)`

> returns `:bool`

### `canvasAt(window, px, py)`

### `_findTopmost(sorted, i, px, py)`

> returns `?`

### `toLocal(canvas, px, py)`

> returns `:object`

### `destroyAll(window)`

## Module: `gui-clipboard`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `CF_UNICODETEXT` · `13`
- `GMEM_MOVEABLE` · `2`
### `_readUtf16Str(addr)`

### `clipboardGetText()`

### `clipboardSetText(text)`

> returns `:bool`

### `clipboardHasText()`

> returns `:bool`

## Module: `gui-color`

### `_clampByte(value)`

### `_clampOpacity(value)`

### `rgb(r, g, b)`

> returns `:bool`

### `colorR(color)`

### `colorG(color)`

### `colorB(color)`

### `opacity(color, amount, background)`

### `rgba(r, g, b, a, background)`

## Module: `gui-common`

### `_default(value, fallback)`

### `_err(message, detail)`

> returns `:object`

### `_clamp(v, minV, maxV)`

## Module: `gui-dialogs`

- `std` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `_ptrSize` — constant
### `_writePtr(address, value)`

### `_readUtf16Str(addr)`

### `_zeros(n)`

- `OFN_SIZE` — constant
- `OFN_FILEBUFSIZE` · `520`
- `OFN_PATHMUSTEXIST` · `2048`
- `OFN_FILEMUSTEXIST` · `4096`
- `OFN_OVERWRITEPROMPT` · `2`
- `OFN_NOCHANGEDIR` · `8`
- `OFN_EXPLORER` · `524288`
### `_buildFilter(filter)`

### `_fileDialog(hwnd, title, filter, flags, apiName)`

> returns `:object`

### `openFileDialog(options)`

### `saveFileDialog(options)`

- `CC_SIZE` — constant
- `CC_RGBINIT` · `1`
- `CC_FULLOPEN` · `2`
### `chooseColor(options)`

> returns `:object`

- `BI_SIZE` — constant
- `BIF_RETURNONLYFSDIRS` · `1`
- `BIF_NEWDIALOGSTYLE` · `64`
### `pickFolder(options)`

> returns `:object`

- `CF_SIZE` — constant
- `LOGFONT_SIZE` · `92`
- `CF_SCREENFONTS` · `1`
- `CF_EFFECTS` · `256`
- `CF_INITTOLOGFONTSTRUCT` · `64`
### `chooseFont(options)`

> returns `:object`

## Module: `gui-draw`

- `windows` · `import(...)`
- `linux` · `import(...)`
- `guiFonts` · `import(...)`
- `guiThread` · `import(...)`
- `threadLib` · `import(...)`
- `_OK` · `{1 entries}`
### `_asBool(value)`

> returns `:bool`

### `_fontKey(fontSpec)`

### `_windowsDeleteCachedFont(window)`

> returns `?`

### `_ensureWindowsFont(window)`

> returns `:int`

### `_ensureLinuxFont(window, gcHandle)`

> returns `:int`

### `_webFontString(fontSpec)`

### `_windowsDeleteCachedPens(window)`

> returns `:int`

### `_windowsDeleteCachedBrushes(window)`

> returns `:int`

### `releaseResources(window)`

> returns `:int`

### `invalidateDrawCaches(window)`

> returns `:int`

### `drawText(window, x, y, text, color, defaultColor)`

### `textWidth(window, text)`

### `_estimateTextWidth(window, text)`

### `fillRect(window, x, y, width, height, color, defaultColor, borderColor)`

- `_maxCacheSize` · `256`
### `_evictPenCache(window)`

> returns `?`

### `_evictBrushCache(window)`

> returns `?`

### `_getCachedPen(window, hdcValue, useColor)`

### `_getCachedBrush(window, useColor)`

- `_nullPenHandle` · `0`
### `_getNullPen()`

### `pushMask(window, x, y, w, h)`

### `popMask(window)`

### `drawLine(window, x1, y1, x2, y2, color, defaultColor)`

### `_setupFillGDI(window, hdcValue, fillColor, borderColor)`

> returns `:bool`

### `fillEllipse(window, cx, cy, rx, ry, fillColor, borderColor)`

### `fillRoundedRect(window, x, y, width, height, radius, fillColor, borderColor)`

### `_writeI32(address, value)`

### `_ensurePointBuf(window, n)`

### `fillPolygon(window, pts, fillColor, borderColor)`

### `drawPolylineNative(window, pts, color)`

### `drawLinesBatch(window, segs, color, defaultColor)`

### `DrawQueue()`

> returns `:object`

### `parallelFillRects(window, coords, computeFn, numWorkers)`

## Module: `gui-draw-ops`

- `guiThread` · `import(...)`
- `threadLib` · `import(...)`
- `_OK` · `{1 entries}`
### `_get(op, key, fallback)`

### `_postOrExec(window, asyncFlag, execFn)`

### `draw(deps, window, op)`

### `drawThreaded(deps, window, op)`

### `drawBatch(deps, window, ops)`

### `drawBatchParallel(deps, window, ops, numWorkers)`

### `drawBatchThreaded(deps, window, ops)`

### `drawBatchParallelThreaded(deps, window, ops, numWorkers)`

## Module: `gui-events`

- `windows` · `import(...)`
- `eventBusLib` · `import(...)`
- `guiInput` · `import(...)`
- `guiThread` · `import(...)`
### `emitThreadSafe(window, event, payload)`

### `emitFromWorker(window, event, payload)`

- `_mapResult` · `{2 entries}`
### `_mapMouseXY(window, physX, physY)`

### `_ensureEventBus(window)`

### `eventBus(window)`

### `on(window, event, handler)`

> returns `?`

### `once(window, event, handler)`

> returns `?`

### `off(window, event, tokenOrHandler)`

> returns `:int`

### `emit(window, event, payload, onDone)`

> returns `:int`

### `listenerCount(window, event)`

> returns `:int`

### `clearListeners(window, event)`

> returns `:bool`

### `publish(window, event, payload)`

> returns `:int`

### `onDispatch(window, handler)`

### `onceDispatch(window, handler)`

### `onRunStart(window, handler)`

### `onceRunStart(window, handler)`

### `onIdle(window, handler)`

### `onceIdle(window, handler)`

### `onFrame(window, handler)`

### `onceFrame(window, handler)`

### `onClosing(window, handler)`

### `onceClosing(window, handler)`

### `onClosed(window, handler)`

### `onceClosed(window, handler)`

- `MSG_OFF_TYPE` · `8`
- `MSG_OFF_WPARAM` · `16`
- `MSG_OFF_LPARAM` · `24`
- `FORM_WM_MOUSEMOVE` · `512`
- `FORM_WM_LBUTTONDOWN` · `513`
- `FORM_WM_LBUTTONUP` · `514`
- `FORM_WM_RBUTTONDOWN` · `516`
- `FORM_WM_RBUTTONUP` · `517`
- `FORM_WM_MBUTTONDOWN` · `519`
- `FORM_WM_MBUTTONUP` · `520`
- `FORM_WM_MOUSEWHEEL` · `522`
- `FORM_WM_LBUTTONDBLCLK` · `515`
- `FORM_WM_KEYDOWN` · `256`
- `FORM_WM_KEYUP` · `257`
- `FORM_WM_CHAR` · `258`
- `GUI_WM_SIZE` · `5`
- `GUI_WM_PAINT` · `15`
- `GUI_WM_ERASEBKGND` · `20`
- `GUI_WM_SIZING` · `532`
- `GUI_WM_WINDOWPOSCHANGED` · `71`
- `GUI_WM_ENTERSIZEMOVE` · `561`
- `GUI_WM_EXITSIZEMOVE` · `562`
- `GUI_WM_MOUSEHOVER` · `673`
- `GUI_WM_MOUSELEAVE` · `675`
- `GUI_WM_DPICHANGED` · `736`
- `WM_IME_STARTCOMPOSITION` · `269`
- `WM_IME_ENDCOMPOSITION` · `270`
- `WM_IME_COMPOSITION` · `271`
- `WM_IME_SETCONTEXT` · `641`
- `WM_IME_NOTIFY` · `642`
- `WM_IME_CHAR` · `646`
- `GCS_COMPSTR` · `8`
- `GCS_RESULTSTR` · `2048`
- `GCS_COMPATTR` · `16`
- `GCS_CURSORPOS` · `128`
- `TME_HOVER` · `1`
- `TME_LEAVE` · `2`
- `TME_SIZEOF` · `24`
- `FORM_VK_BACK` — constant
- `FORM_VK_TAB` — constant
- `FORM_VK_RETURN` — constant
- `FORM_VK_SHIFT` — constant
- `FORM_VK_CONTROL` — constant
- `FORM_VK_ALT` — constant
- `FORM_VK_ESCAPE` — constant
- `MK_LBUTTON` · `1`
- `MK_RBUTTON` · `2`
- `MK_SHIFT` · `4`
- `MK_CONTROL` · `8`
- `MK_MBUTTON` · `16`
### `modShift?(wp)`

> returns `:bool`

### `modCtrl?(wp)`

> returns `:bool`

### `modAlt?(window)`

> returns `:bool`

### `keyShiftDown?(window)`

> returns `:bool`

### `keyCtrlDown?(window)`

> returns `:bool`

### `keyAltDown?(window)`

### `formMsgType(window)`

### `formMsgWParam(window)`

### `formMsgLParam(window)`

### `formLoWord(v)`

> returns `:bool`

### `formHiWord(v)`

> returns `:bool`

### `_cacheDispatchContext(window)`

### `_clearDispatchContext(window)`

> returns `?`

### `formEventContext(window)`

### `formInRect?(mx, my, rx, ry, rw, rh)`

> returns `:bool`

### `_vkKeyName(vk)`

- `_noKeyMatch` · `{1 entries}`
### `_extractKeyEvent(window, evt, expectedMsgType, expectedEvtType)`

> returns `:object`

### `onKeyDownEvent(window, handler)`

### `onceKeyDownEvent(window, handler)`

### `onKeyUpEvent(window, handler)`

### `onceKeyUpEvent(window, handler)`

### `onMouseMove(window, handler)`

### `onLButtonDown(window, handler)`

### `onLButtonUp(window, handler)`

### `onRButtonDown(window, handler)`

### `onRButtonUp(window, handler)`

### `onMButtonDown(window, handler)`

### `onMButtonUp(window, handler)`

### `onMouseWheel(window, handler)`

### `onLButtonDblClk(window, handler)`

### `_trackMouseEvent(window, flags)`

### `enableMouseTracking(window)`

### `onMouseHover(window, handler)`

### `onMouseLeave(window, handler)`

### `onKeyDown(window, handler)`

### `onKeyUp(window, handler)`

### `onChar(window, handler)`

### `onResize(window, handler)`

### `onDpiChanged(window, handler)`

### `_imeGetString(hwnd, gcsFlag)`

### `_imeGetCursorPos(hwnd)`

### `onImeStartComposition(window, handler)`

### `onImeEndComposition(window, handler)`

### `onImeComposition(window, handler)`

### `setImePosition(window, x, y)`

- `WM_TOUCH` · `576`
- `_TOUCHINPUT_SIZE` · `40`
- `TOUCHEVENTF_MOVE` · `1`
- `TOUCHEVENTF_DOWN` · `2`
- `TOUCHEVENTF_UP` · `4`
- `TOUCHEVENTF_INRANGE` · `8`
- `TOUCHEVENTF_PRIMARY` · `16`
- `TOUCHEVENTF_NOCOALESCE` · `32`
- `TOUCHEVENTF_PEN` · `64`
- `TOUCHEVENTF_PALM` · `128`
- `TOUCHINPUTMASKF_CONTACTAREA` · `4`
- `TOUCHINPUTMASKF_EXTRAINFO` · `2`
- `TOUCHINPUTMASKF_TIMEFROMSYSTEM` · `1`
- `TWF_FINETOUCH` · `1`
- `TWF_WANTPALM` · `2`
- `WM_POINTERDOWN` · `582`
- `WM_POINTERUP` · `583`
- `WM_POINTERUPDATE` · `581`
- `WM_POINTERENTER` · `585`
- `WM_POINTERLEAVE` · `586`
- `WM_POINTERCAPTURECHANGED` · `588`
- `WM_POINTERWHEEL` · `590`
- `WM_POINTERHWHEEL` · `591`
- `PT_POINTER` · `1`
- `PT_TOUCH` · `2`
- `PT_PEN` · `3`
- `PT_MOUSE` · `4`
- `PT_TOUCHPAD` · `5`
- `POINTER_FLAG_NONE` · `0`
- `POINTER_FLAG_NEW` · `1`
- `POINTER_FLAG_INRANGE` · `2`
- `POINTER_FLAG_INCONTACT` · `4`
- `POINTER_FLAG_FIRSTBUTTON` · `16`
- `POINTER_FLAG_SECONDBUTTON` · `32`
- `POINTER_FLAG_PRIMARY` · `8192`
- `POINTER_FLAG_DOWN` · `65536`
- `POINTER_FLAG_UPDATE` · `131072`
- `POINTER_FLAG_UP` · `262144`
- `PEN_FLAG_BARREL` · `1`
- `PEN_FLAG_INVERTED` · `2`
- `PEN_FLAG_ERASER` · `4`
- `PEN_MASK_PRESSURE` · `1`
- `PEN_MASK_ROTATION` · `2`
- `PEN_MASK_TILT_X` · `4`
- `PEN_MASK_TILT_Y` · `8`
### `registerTouchWindow(window, flags)`

### `unregisterTouchWindow(window)`

### `_parseTouchInputs(window, count)`

### `onTouch(window, handler)`

### `enableTouchInput(window, options)`

### `_pointerIdFromWParam(wp)`

> returns `:bool`

### `_getPointerType(pointerId)`

- `_POINTER_INFO_SIZE` · `96`
### `_getPointerInfo(pointerId)`

> returns `:object`

- `_POINTER_PEN_INFO_SIZE` · `120`
### `_getPointerPenInfo(pointerId)`

> returns `:object`

### `_makePointerEvent(window, pointerId)`

### `onPointerDown(window, handler)`

### `onPointerUp(window, handler)`

### `onPointerUpdate(window, handler)`

### `onPointerEnter(window, handler)`

### `onPointerLeave(window, handler)`

### `enablePointerInput(window)`

- `WM_INPUT` · `255`
- `HID_USAGE_PAGE_GENERIC` · `1`
- `HID_USAGE_PAGE_GAME` · `5`
- `HID_USAGE_PAGE_LED` · `8`
- `HID_USAGE_PAGE_BUTTON` · `9`
- `HID_USAGE_GENERIC_POINTER` · `1`
- `HID_USAGE_GENERIC_MOUSE` · `2`
- `HID_USAGE_GENERIC_JOYSTICK` · `4`
- `HID_USAGE_GENERIC_GAMEPAD` · `5`
- `HID_USAGE_GENERIC_KEYBOARD` · `6`
- `HID_USAGE_GENERIC_KEYPAD` · `7`
- `HID_USAGE_GENERIC_MULTI_AXIS` · `8`
- `RIDEV_REMOVE` · `1`
- `RIDEV_INPUTSINK` · `256`
- `RIDEV_NOLEGACY` · `48`
- `RIDEV_DEVNOTIFY` · `8192`
- `RID_INPUT` · `268435459`
- `RID_HEADER` · `268435461`
- `RIM_TYPEMOUSE` · `0`
- `RIM_TYPEKEYBOARD` · `1`
- `RIM_TYPEHID` · `2`
- `_RAWINPUTDEVICE_SIZE` — constant
### `registerRawInputDevice(window, usagePage, usage, flags)`

### `unregisterRawInputDevice(usagePage, usage)`

- `_RAWINPUTHEADER_SIZE` — constant
### `_getRawInputData(lParam)`

> returns `:object`

### `onRawInput(window, handler)`

### `enableRawMouse(window)`

### `enableRawKeyboard(window)`

### `enableRawGamepad(window)`

### `enableRawJoystick(window)`

### `isResizeDispatch?(window, step)`

> returns `:bool`

## Module: `gui-filedrop`

- `std` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `WM_DROPFILES` · `563`
### `enableFileDrop(window)`

### `disableFileDrop(window)`

### `_readUtf16Str(addr)`

### `onFileDrop(window, handler)`

- `DROPEFFECT_NONE` · `0`
- `DROPEFFECT_COPY` · `1`
- `DROPEFFECT_MOVE` · `2`
- `DROPEFFECT_LINK` · `4`
### `enableOleDrop(window)`

> returns `:bool`

### `disableOleDrop(window)`

> returns `:bool`

### `onOleDrop(window, handler)`

### `dragDropState()`

> returns `:object`

### `onDragOver(window, state, handler)`

## Module: `gui-fonts`

- `guiThread` · `import(...)`
### `_asBool(v)`

> returns `:bool`

- `_platformCache` — constant
### `_platformId()`

### `isWindows?()`

### `isLinux?()`

- `_winMod` · `?`
- `_linuxMod` · `?`
### `_win()`

### `_lnx()`

- `FW_THIN` · `100`
- `FW_EXTRALIGHT` · `200`
- `FW_LIGHT` · `300`
- `FW_NORMAL` · `400`
- `FW_MEDIUM` · `500`
- `FW_SEMIBOLD` · `600`
- `FW_BOLD` · `700`
- `FW_EXTRABOLD` · `800`
- `FW_HEAVY` · `900`
### `defaultFontSpec()`

> returns `:object`

### `createFont(spec)`

### `deleteFont(fontResult)`

> returns `:int`

### `_webFontString(spec)`

> returns `?`

### `fontKey(spec)`

> returns `?`

### `cachedFont(windowOrDisplay, spec)`

### `releaseCachedFonts(windowOrDisplay)`

### `measureText(windowOrDisplay, spec, text)`

### `selectFont(args)`

### `getTextMetrics(hdc)`

### `fontLineHeight(hdcOrFontStruct)`

### `buildXLFD(spec)`

### `webFontString(spec)`

## Module: `gui-form`

- `std` · `import(...)`
- `guiInput` · `import(...)`
- `guiThread` · `import(...)`
### `formSetStatus(state, message, ok)`

### `formPopChar(s)`

### `formClamp(v, minVal, maxVal)`

### `formInRect?(mx, my, rx, ry, rw, rh)`

> returns `:bool`

### `formHitListIndex(mx, my, x, y, itemW, itemH, count)`

### `formHitRectKey(mx, my, rects)`

### `formToggleIfHit(value, mx, my, x, y, w, h)`

> returns `:bool`

### `formSelectByHit(current, mx, my, rects)`

### `formSetByAssignments(state, assignments)`

### `formResetFlags(state, keys, value)`

### `formSetHoverFromRects(state, mx, my, rects)`

### `formToggleKeysByHit(state, mx, my, rects)`

### `formSetKeyByHit(state, targetKey, mx, my, rects)`

### `formTruncateText(s, maxChars)`

### `formIsPrintableChar?(code)`

> returns `:bool`

- `_charWidth` · `7`
### `formSelectionState()`

> returns `:object`

### `formSelSetCursor(sel, pos, shifting)`

### `formSelMoveCursor(sel, text, dir, shifting)`

### `formSelHome(sel, shifting)`

### `formSelEnd(sel, text, shifting)`

### `formSelAll(sel, text)`

### `formSelRange(sel)`

> returns `:object`

### `formSelHasSelection?(sel)`

### `formSelSelectedText(sel, text)`

### `formSelDeleteSelection(sel, text)`

### `formSelInsertAtCursor(sel, text, insert)`

### `formSelBackspace(sel, text)`

### `formSelClickPos(mx, fieldX, fieldPadding, text)`

### `formDrawFieldWithSel(ctx, window, x, y, w, h, text, placeholder, focused, sel, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor, selColor)`

### `formUndoState(maxHistory)`

> returns `:object`

### `formUndoPush(hist, text, sel)`

### `formUndo(hist, text, sel)`

### `formRedo(hist, text, sel)`

### `formAppendByFocus(state, focus, c, fieldSpecs, notesSpec)`

### `formBackspaceByFocus(state, focus, fieldKeys, notesSpec)`

- `_clipboard` · `?`
### `_cb()`

### `formCopyByFocus(state, focus, fieldKeys)`

### `formPasteByFocus(state, focus, fieldSpecs)`

### `formCutByFocus(state, focus, fieldKeys)`

### `formNotesAppendChar(lines, c)`

### `formNotesBackspace(lines)`

### `formNotesNewLine(lines, maxLines)`

### `formSliderValue(mx, sliderX, sliderW, handleW, minVal, maxVal)`

### `formApplySliderDrag(state, dragKey, mx, sliderX, sliderW, handleW, minVal, maxVal, bindings)`

### `formNextField(current, order)`

### `formPrevField(current, order)`

### `formFocusState(tabOrder)`

> returns `:object`

### `formFocusNext(fs)`

### `formFocusPrev(fs)`

### `formFocusSet(fs, key)`

### `formFocusIs?(fs, key)`

### `formDrawFocusRing(ctx, window, x, y, w, h, color, options)`

### `formHandleTabKey(fs, shiftDown)`

- `_VK_RETURN` — constant
- `_VK_SPACE` — constant
- `_VK_ESCAPE` — constant
- `_VK_UP` — constant
- `_VK_DOWN` — constant
- `_VK_LEFT` — constant
- `_VK_RIGHT` — constant
- `_VK_HOME` — constant
- `_VK_END` — constant
- `_VK_PGUP` — constant
- `_VK_PGDN` — constant
### `formIsActivateKey?(vk)`

> returns `:bool`

### `formHandleKeyNav(fs, vk, shiftDown, widgetHandlers, options)`

> returns `:atom`

### `formCheckboxKeyToggle(checked, vk)`

> returns `:bool`

### `formRadioKeySelect(current, vk, choices)`

### `formSliderKeyAdjust(value, vk, minVal, maxVal, step)`

### `formDropdownKeyNav(open?, selectedIdx, vk, itemCount)`

> returns `:object`

### `formTreeKeyNav(flatList, selectedIdx, vk)`

> returns `:object`

### `formTableKeyNav(selectedRow, vk, rowCount, pageSize)`

### `formMaskText(s)`

### `formDrawBorder(ctx, window, x, y, w, h, color)`

### `formDrawField(ctx, window, x, y, w, h, text, placeholder, focused, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawPasswordField(ctx, window, x, y, w, h, text, placeholder, focused, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawCheckbox(ctx, window, x, y, checked, label, fieldColor, borderColor, checkColor)`

### `formDrawRadio(ctx, window, x, y, selected, label, fieldColor, borderColor, accentColor)`

### `formDrawPrimaryButton(ctx, window, x, y, w, h, label, hover, press, baseColor, hoverColor, pressColor, bottomLineColor)`

### `formDrawSecondaryButton(ctx, window, x, y, w, h, label, hover, press, baseColor, hoverColor, pressColor, borderColor)`

### `formDrawSlider(ctx, window, x, y, w, value, maxVal, trackColor, fillColor, thumbColor, thumbBorderColor, options)`

### `formDrawLabeledPercentSlider(ctx, window, x, labelY, sliderY, w, label, value, maxVal, trackColor, fillColor, thumbColor, thumbBorderColor, percentRectColor)`

### `formDrawNotes(ctx, window, x, y, w, h, lines, focused, placeholder, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawStatusBanner(ctx, window, x, y, w, h, message, ok, fieldColor, borderColor, successColor, errorColor)`

### `formDrawProgressBar(ctx, window, x, y, w, value, maxVal, trackColor, fillColor, borderColor, options)`

### `formDrawTabStrip(ctx, window, x, y, tabs, activeIdx, bgColor, activeBgColor, borderColor, activeTextColor, inactiveTextColor, options)`

### `formHitTab(mx, my, x, y, tabs, options)`

### `formDrawDropdown(ctx, window, x, y, w, h, selectedLabel, open, fieldColor, fieldFocusColor, borderColor, arrowColor)`

### `formDrawDropdownList(ctx, window, x, y, w, items, hoverIdx, itemHeight, listBgColor, hoverBgColor, borderColor)`

### `formHitDropdownItem(mx, my, x, y, w, items, itemHeight)`

### `formTooltipState()`

> returns `:object`

### `formTooltipUpdate(state, mx, my, regionKey, text, delay)`

### `formTooltipHide(state)`

### `formDrawTooltip(ctx, window, state, bgColor, borderColor, textColor, options)`

### `formDrawSpinner(ctx, window, x, y, w, h, value, focused, fieldColor, fieldFocusColor, borderColor, arrowColor, options)`

### `formSpinnerHit(mx, my, x, y, w, h, options)`

> returns `:atom`

### `formSpinnerAdjust(value, direction, options)`

### `formDrawScrollbar(ctx, window, x, y, w, h, scrollPos, contentSize, viewSize, trackColor, thumbColor, borderColor, options)`

### `formScrollbarHit(mx, my, x, y, w, h, scrollPos, contentSize, viewSize, options)`

### `_flattenTree(nodes, depth, out)`

### `formDrawTreeView(ctx, window, x, y, w, h, nodes, selectedIdx, scrollPos, bgColor, selectedBgColor, textColor, selectedTextColor, borderColor, options)`

### `formTreeHitRow(mx, my, x, y, w, flat, scrollPos, options)`

### `formTreeToggle(flat, idx)`

> returns `:bool`

### `formTreeContentHeight(flat, options)`

### `_isLeapYear?(y)`

> returns `:bool`

### `_daysInMonth(year, month)`

> returns `:int`

### `_dayOfWeek(y, m, d)`

### `formDatePickerState(year, month, day)`

> returns `:object`

### `formDatePickerPrevMonth(state)`

### `formDatePickerNextMonth(state)`

### `_padZero(n)`

> returns `:string`

- `_monthNames` · `[12 items]`
### `formDateLabel(state)`

### `formDrawDateField(ctx, window, x, y, w, h, state, fieldColor, fieldFocusColor, borderColor, arrowColor)`

### `formDrawDateCalendar(ctx, window, x, y, state, bgColor, headerBgColor, selectedBgColor, todayBorderColor, textColor, selectedTextColor, headerTextColor, borderColor, options)`

### `formDateCalendarHit(mx, my, x, y, state, options)`

> returns `:object`

### `formDrawTable(ctx, window, x, y, w, h, columns, rows, scrollPos, selectedRow, headerBgColor, rowBgColor, altRowBgColor, selectedBgColor, headerTextColor, textColor, selectedTextColor, borderColor, options)`

### `formTableHitRow(mx, my, x, y, w, columns, rows, scrollPos, options)`

### `formTableContentHeight(rows, options)`

### `formTableHitColumn(mx, x, columns)`

### `formFrameTimerState(maxSamples)`

> returns `:object`

### `formFrameTimerTick(state)`

### `formDrawFrameTimingOverlay(ctx, window, state, x, y, w, h, options)`

### `formRichTextState(lines)`

> returns `:object`

### `formRichTextSetLines(state, lines)`

### `formRichTextAppendLine(state, spans)`

### `formRichTextInsertSpan(state, lineIdx, spanIdx, span)`

### `_spanWidth(text, fontSize)`

### `formDrawRichText(ctx, window, x, y, w, h, state, bgColor, borderColor)`

### `formRichTextScroll(state, delta)`

### `formRichTextTotalHeight(state)`

### `richSpan(text)`

> returns `:object`

### `richBold(text)`

> returns `:object`

### `richItalic(text)`

> returns `:object`

### `richUnderline(text)`

> returns `:object`

### `richColored(text, r, g, b)`

> returns `:object`

### `richStyled(text, options)`

## Module: `gui-gamepad`

- `sys` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `_xinputProcCache` · `{}`
- `_xinputDll` · `?`
### `_resolveXInputDll()`

### `xinput(symbol, args...)`

- `XUSER_MAX_COUNT` · `4`
- `ERROR_SUCCESS` · `0`
- `ERROR_DEVICE_NOT_CONNECTED` · `1167`
- `XINPUT_GAMEPAD_DPAD_UP` · `1`
- `XINPUT_GAMEPAD_DPAD_DOWN` · `2`
- `XINPUT_GAMEPAD_DPAD_LEFT` · `4`
- `XINPUT_GAMEPAD_DPAD_RIGHT` · `8`
- `XINPUT_GAMEPAD_START` · `16`
- `XINPUT_GAMEPAD_BACK` · `32`
- `XINPUT_GAMEPAD_LEFT_THUMB` · `64`
- `XINPUT_GAMEPAD_RIGHT_THUMB` · `128`
- `XINPUT_GAMEPAD_LEFT_SHOULDER` · `256`
- `XINPUT_GAMEPAD_RIGHT_SHOULDER` · `512`
- `XINPUT_GAMEPAD_A` · `4096`
- `XINPUT_GAMEPAD_B` · `8192`
- `XINPUT_GAMEPAD_X` · `16384`
- `XINPUT_GAMEPAD_Y` · `32768`
- `XINPUT_GAMEPAD_LEFT_THUMB_DEADZONE` · `7849`
- `XINPUT_GAMEPAD_RIGHT_THUMB_DEADZONE` · `8689`
- `XINPUT_GAMEPAD_TRIGGER_THRESHOLD` · `30`
- `XINPUT_VIBRATION_MAX` · `65535`
- `_XINPUT_STATE_SIZE` · `16`
- `_XINPUT_VIBRATION_SIZE` · `4`
### `_readI16(base, off)`

### `getGamepadState(playerIndex)`

> returns `:object`

### `gamepadButtonDown?(state, button)`

> returns `:bool`

### `gamepadDpadUp?(state)`

### `gamepadDpadDown?(state)`

### `gamepadDpadLeft?(state)`

### `gamepadDpadRight?(state)`

### `gamepadStart?(state)`

### `gamepadBack?(state)`

### `gamepadA?(state)`

### `gamepadB?(state)`

### `gamepadX?(state)`

### `gamepadY?(state)`

### `gamepadLeftShoulder?(state)`

### `gamepadRightShoulder?(state)`

### `gamepadLeftThumb?(state)`

### `gamepadRightThumb?(state)`

### `applyDeadzone(value, deadzone)`

> returns `:int`

### `applyThumbDeadzone(state)`

> returns `:object`

### `setGamepadVibration(playerIndex, leftMotor, rightMotor)`

### `stopGamepadVibration(playerIndex)`

### `gamepadPollState()`

> returns `:object`

### `gamepadPoll(ps)`

### `gamepadButtonPressed?(ps, idx, button)`

> returns `:bool`

### `gamepadButtonReleased?(ps, idx, button)`

> returns `:bool`

### `gamepadConnected?(ps, idx)`

### `gamepadJustConnected?(ps, idx)`

> returns `:bool`

### `gamepadJustDisconnected?(ps, idx)`

> returns `:bool`

## Module: `gui-gpu-info`

- `guiThread` · `import(...)`
### `getGPUAdapters()`

### `getGPUAdaptersParsed()`

### `getDXGIAdapters()`

### `getD3DFeatureLevel()`

### `getDisplayModes()`

### `getMonitorInfo()`

### `gpuCapabilityDump()`

### `gpuCapabilityDumpParallel()`

## Module: `gui-graph`

- `threadLib` · `import(...)`
### `graphRange(values, options)`

> returns `:object`

### `graphMapX(index, count, x, width)`

### `graphMapY(value, y, height, range)`

### `_graphDrawBorder(ctx, window, x, y, width, height, color)`

### `drawGraphAxes(ctx, window, x, y, width, height, options)`

### `drawLineGraph(ctx, window, x, y, width, height, values, options)`

### `drawBarGraph(ctx, window, x, y, width, height, values, options)`

### `drawSparkline(ctx, window, x, y, width, height, values, options)`

## Module: `gui-input`

- `VK_BACK` · `8`
- `VK_TAB` · `9`
- `VK_CLEAR` · `12`
- `VK_RETURN` · `13`
- `VK_SHIFT` · `16`
- `VK_CONTROL` · `17`
- `VK_ALT` · `18`
- `VK_PAUSE` · `19`
- `VK_CAPSLOCK` · `20`
- `VK_ESCAPE` · `27`
- `VK_SPACE` · `32`
- `VK_PAGEUP` · `33`
- `VK_PAGEDOWN` · `34`
- `VK_END` · `35`
- `VK_HOME` · `36`
- `VK_LEFT` · `37`
- `VK_UP` · `38`
- `VK_RIGHT` · `39`
- `VK_DOWN` · `40`
- `VK_INSERT` · `45`
- `VK_DELETE` · `46`
- `VK_0` · `48`
- `VK_1` · `49`
- `VK_2` · `50`
- `VK_3` · `51`
- `VK_4` · `52`
- `VK_5` · `53`
- `VK_6` · `54`
- `VK_7` · `55`
- `VK_8` · `56`
- `VK_9` · `57`
- `VK_A` · `65`
- `VK_B` · `66`
- `VK_C` · `67`
- `VK_D` · `68`
- `VK_E` · `69`
- `VK_F` · `70`
- `VK_G` · `71`
- `VK_H` · `72`
- `VK_I` · `73`
- `VK_J` · `74`
- `VK_K` · `75`
- `VK_L` · `76`
- `VK_M` · `77`
- `VK_N` · `78`
- `VK_O` · `79`
- `VK_P` · `80`
- `VK_Q` · `81`
- `VK_R` · `82`
- `VK_S` · `83`
- `VK_T` · `84`
- `VK_U` · `85`
- `VK_V` · `86`
- `VK_W` · `87`
- `VK_X` · `88`
- `VK_Y` · `89`
- `VK_Z` · `90`
- `VK_NUMPAD0` · `96`
- `VK_NUMPAD1` · `97`
- `VK_NUMPAD2` · `98`
- `VK_NUMPAD3` · `99`
- `VK_NUMPAD4` · `100`
- `VK_NUMPAD5` · `101`
- `VK_NUMPAD6` · `102`
- `VK_NUMPAD7` · `103`
- `VK_NUMPAD8` · `104`
- `VK_NUMPAD9` · `105`
- `VK_MULTIPLY` · `106`
- `VK_ADD` · `107`
- `VK_SEPARATOR` · `108`
- `VK_SUBTRACT` · `109`
- `VK_DECIMAL` · `110`
- `VK_DIVIDE` · `111`
- `VK_F1` · `112`
- `VK_F2` · `113`
- `VK_F3` · `114`
- `VK_F4` · `115`
- `VK_F5` · `116`
- `VK_F6` · `117`
- `VK_F7` · `118`
- `VK_F8` · `119`
- `VK_F9` · `120`
- `VK_F10` · `121`
- `VK_F11` · `122`
- `VK_F12` · `123`
- `VK_NUMLOCK` · `144`
- `VK_SCROLLLOCK` · `145`
- `VK_OEM_SEMICOLON` · `186`
- `VK_OEM_PLUS` · `187`
- `VK_OEM_COMMA` · `188`
- `VK_OEM_MINUS` · `189`
- `VK_OEM_PERIOD` · `190`
- `VK_OEM_SLASH` · `191`
- `VK_OEM_TILDE` · `192`
- `VK_OEM_LBRACKET` · `219`
- `VK_OEM_BACKSLASH` · `220`
- `VK_OEM_RBRACKET` · `221`
- `VK_OEM_QUOTE` · `222`
- `VK_LSHIFT` · `160`
- `VK_RSHIFT` · `161`
- `VK_LCONTROL` · `162`
- `VK_RCONTROL` · `163`
- `VK_LALT` · `164`
- `VK_RALT` · `165`
### `isLetterKey?(vk)`

> returns `:bool`

### `isDigitKey?(vk)`

> returns `:bool`

### `isNumpadKey?(vk)`

> returns `:bool`

### `isFunctionKey?(vk)`

> returns `:bool`

### `isArrowKey?(vk)`

> returns `:bool`

### `isModifierKey?(vk)`

> returns `:bool`

## Module: `gui-leak-detect`

- `guiThread` · `import(...)`
- `windows` · `import(...)`
- `GR_GDIOBJECTS` · `0`
- `GR_USEROBJECTS` · `1`
- `GR_GDIOBJECTS_PEAK` · `2`
- `GR_USEROBJECTS_PEAK` · `4`
### `_getCurrentProcess()`

### `getGDIObjectCount()`

### `getUserObjectCount()`

### `getGDIObjectPeak()`

### `getUserObjectPeak()`

### `getHandleCount()`

- `_PMC_SIZE` · `72`
### `getWorkingSetSize()`

### `getPeakWorkingSetSize()`

### `leakDetectorState()`

> returns `:object`

### `leakSnapshot(state)`

### `leakSnapshotParallel(state)`

### `leakCheck(state)`

### `leakReport(state)`

### `leakSetThresholds(state, thresholds)`

### `leakResetBaseline(state)`

### `leakTrend(state)`

## Module: `gui-lighting`

- `threadLib` · `import(...)`
- `PI` · `3.141592653589793`
### `_len3(x, y, z)`

### `_norm3(x, y, z)`

> returns `:list`

### `_dot(ax, ay, az, bx, by, bz)`

### `DirectionalLight(options)`

> returns `:object`

### `PointLight(options)`

> returns `:object`

### `SpotLight(options)`

> returns `:object`

### `AmbientLight(options)`

> returns `:object`

### `Material(options)`

> returns `:object`

### `LightScene(options)`

> returns `:object`

### `addLight(scene, light)`

### `removeLight(scene, index)`

### `clearLights(scene)`

### `lightCount(scene)`

### `faceNormal(pa, pb, pc)`

### `faceCenter(pa, pb, pc)`

> returns `:object`

### `_shadeDirectional(light, mat, n0, n1, n2, vdx, vdy, vdz, acc)`

### `_shadePoint(light, mat, n0, n1, n2, cx, cy, cz, vdx, vdy, vdz, acc)`

### `_shadeSpot(light, mat, n0, n1, n2, cx, cy, cz, vdx, vdy, vdz, acc)`

### `_shadeAmbient(light, mat, acc)`

### `_accumulate(lights, mat, n0, n1, n2, cx, cy, cz, vdx, vdy, vdz, i, acc)`

### `shadeFaceColor(deps, baseColor, scene, material, pa3, pb3, pc3, camPos)`

### `shadeFaceIntensity(scene, material, pa3, pb3, pc3, camPos)`

### `prepareScene(scene, material, camPos)`

> returns `:object`

### `_prepareLights(lights, i, acc)`

### `shadeFaceColorFast(rgbFn, baseColor, prep, pa3, pb3, pc3)`

### `shadeFacesBatchParallel(rgbFn, faces, prepared, numWorkers)`

## Module: `gui-loop`

- `windows` · `import(...)`
- `guiEvents` · `import(...)`
- `guiNativeLinux` · `import(...)`
- `guiThread` · `import(...)`
### `frameIntervalNt(window)`

### `frameMaxDtNt(window)`

### `markUrgentFrameIfResizeDispatch(window, step)`

> returns `:bool`

### `maybeRunFrame(window, onFrame, force, publish)`

> returns `?`

### `sleepUntilNextFrame(window)`

> returns `:int`

## Module: `gui-menus`

- `std` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `MF_STRING` · `0`
- `MF_SEPARATOR` · `2048`
- `MF_POPUP` · `16`
- `MF_CHECKED` · `8`
- `MF_UNCHECKED` · `0`
- `MF_GRAYED` · `1`
- `MF_ENABLED` · `0`
- `MF_BYCOMMAND` · `0`
- `MF_BYPOSITION` · `1024`
- `TPM_LEFTBUTTON` · `0`
- `TPM_RIGHTBUTTON` · `2`
- `TPM_RETURNCMD` · `256`
- `WM_COMMAND` · `273`
### `createMenu()`

### `createPopupMenu()`

### `appendItem(hMenu, id, label)`

### `appendSeparator(hMenu)`

### `appendSubmenu(hMenu, hSubMenu, label)`

### `setMenuBar(window, hMenu)`

> returns `:bool`

### `removeMenuBar(window)`

> returns `:bool`

### `destroyMenu(hMenu)`

> returns `:bool`

### `showPopupMenu(window, hMenu, x, y)`

### `onMenuCommand(window, handler)`

### `buildMenu(spec)`

- `FVIRTKEY` · `1`
- `FSHIFT` · `4`
- `FCONTROL` · `8`
- `FALT` · `16`
### `createAcceleratorTable(entries)`

### `destroyAcceleratorTable(hAccel)`

> returns `:bool`

### `translateAccelerator(window, hAccel)`

### `installAccelerators(window, entries)`

> returns `:object`

- `WM_SYSCOMMAND` · `274`
### `getSystemMenu(window, reset)`

### `appendSystemMenuItem(window, id, label)`

### `appendSystemMenuSeparator(window)`

### `resetSystemMenu(window)`

### `onSysCommand(window, handler)`

## Module: `gui-mesh`

- `threadLib` · `import(...)`
### `Vec3(x, y, z)`

> returns `:object`

### `Mesh(vertices, edges)`

> returns `:object`

### `GridMesh(size, step)`

### `AxesMesh(length)`

### `_vecKey(v)`

### `_edgeKey(a, b)`

### `_addVertex(vertices, indexByKey, v)`

> returns `:object`

### `_addEdge(edges, edgeSet, a, b)`

> returns `:object`

### `_voxelMeshSub(voxels, voxelSize, i, vertices, indexByKey, edges, edgeSet)`

### `VoxelMesh(voxels, voxelSize)`

### `PlaneMesh(width, depth, subdivisionsW, subdivisionsD)`

> returns `:object`

### `PyramidMesh(base, height)`

> returns `:object`

### `_cylinderRingVerts(cx, cy, cz, radius, segments, i, out)`

### `CylinderMesh(radius, height, segments)`

> returns `:object`

### `ConeMesh(radius, height, segments)`

> returns `:object`

### `SphereMesh(radius, segments, rings)`

> returns `:object`

### `TorusMesh(majorRadius, minorRadius, majorSegments, minorSegments)`

> returns `:object`

### `VoxelGrid(options)`

> returns `:object`

### `HemisphereMesh(radius, segments, rings)`

> returns `:object`

### `WedgeMesh(width, height, depth)`

> returns `:object`

### `TubeMesh(outerRadius, innerRadius, height, segments)`

> returns `:object`

### `ArrowMesh(shaftRadius, shaftHeight, headRadius, headHeight, segments)`

> returns `:object`

### `PrismMesh(radius, height, sides)`

> returns `:object`

### `StairsMesh(steps, width, stepHeight, stepDepth)`

### `IcosphereMesh(radius)`

> returns `:object`

### `parallelMeshGenerate(specs)`

## Module: `gui-native-linux`

- `linux` · `import(...)`
- `guiThread` · `import(...)`
### `createWindowState(title, width, height, frameMs, updateOnDispatch)`

> returns `:object`

### `showWindow(window)`

### `hideWindow(window)`

### `moveWindow(window, x, y)`

### `resizeWindow(window, width, height)`

### `_displaySize(window)`

> returns `:object`

### `setFullscreen(window, enabled)`

### `setTitle(window, title)`

### `lockResize(window, locked)`

### `poll(window)`

> returns `:object`

### `close(window)`

> returns `:int`

### `sleepFrame(frameMs)`

## Module: `gui-native-win`

- `_nwImportT0` · `nanotime(...)`
- `windows` · `import(...)`
- `_nwImportT1` · `nanotime(...)`
- `guiNativeWinPresent` · `import(...)`
- `guiNativeWinIcons` · `import(...)`
- `guiNativeWinFrame` · `import(...)`
- `guiNativeWinPoll` · `import(...)`
- `guiNativeWinClose` · `import(...)`
- `guiNativeWinDdraw` · `import(...)`
- `guiNativeWinD3d11` · `import(...)`
- `guiNativeWinVulkan` · `import(...)`
- `guiNativeWinOpenGL` · `import(...)`
- `guiNativeWinProbe` · `import(...)`
- `guiCanvas` · `import(...)`
- `_nwImportT2` · `nanotime(...)`
- `guiThread` · `import(...)`
- `_nwImportT3` · `nanotime(...)`
- `_nwImportTimings` · `{4 entries}`
### `getNativeWinImportTimings()`

- `D3D_SDK_VERSION` · `32`
- `COM_RELEASE` · `2`
- `_windowThreadLockCount` · `0`
### `_acquireWindowThreadLock()`

### `_releaseWindowThreadLock()`

### `_startDllProbeAsync(state)`

### `_startD3d9ProbeAsync(state)`

### `_applyDllProbeResult(state)`

### `_init2DLayer(window)`

### `_init3DLayer(window)`

> returns `:object`

- `SRCCOPY` · `13369376`
- `SWP_NOSIZE` · `1`
- `SWP_NOMOVE` · `2`
- `SWP_NOZORDER` · `4`
- `SWP_NOACTIVATE` · `16`
- `SWP_FRAMECHANGED` · `32`
- `SWP_NOOWNERZORDER` · `512`
### `_fallbackPresenterAfterVulkanFailure(state)`

> returns `:atom`

### `_handleVulkan2DInit(state)`

### `_handleOpenGL2DInit(state)`

> returns `:object`

### `_handleD3d11Init(state)`

> returns `:object`

### `_finalize2DInit(state)`

### `_finalize3DInit(state)`

> returns `:object`

### `_runLayerInit(state)`

### `_ensureLayerInit(window)`

> returns `:object`

### `createWindowState(title, width, height, options, className, frameMs, updateOnDispatch)`

### `createWindowAsync(title, width, height, options, className, frameMs, updateOnDispatch)`

### `awaitWindow(future)`

- `_showTimings` · `?`
### `showWindow(window)`

### `getShowTimings()`

### `hideWindow(window)`

### `moveWindow(window, x, y)`

### `resizeWindow(window, width, height)`

### `setFullscreen(window, enabled)`

### `lockResize(window, locked)`

### `setAlwaysOnTop(window, enabled)`

- `WS_EX_LAYERED` · `524288`
- `LWA_ALPHA` · `2`
- `LWA_COLORKEY` · `1`
### `setWindowOpacity(window, alpha)`

### `setWindowColorKey(window, colorKey)`

### `removeLayeredStyle(window)`

### `beginDrag(window)`

### `updateDrag(window)`

### `endDrag(window)`

> returns `?`

### `setTitle(window, title)`

### `setIcon(window, iconSpec)`

### `_ensureLayerInitForFrame(window)`

### `beginFrame(window)`

### `endFrame(window)`

> returns `:int`

### `poll(window)`

> returns `:object`

### `close(window)`

> returns `:int`

- `_windowRegistry` · `[]`
### `registerWindow(window)`

### `unregisterWindow(window)`

### `allWindows()`

### `pollAllWindows()`

### `closeAllWindows()`

> returns `:list`

### `anyWindowOpen?()`

### `saveWindowState(window)`

> returns `:object`

### `restoreWindowState(window, state)`

- `_ERROR_ALREADY_EXISTS` · `183`
### `acquireSingleInstance(name)`

> returns `:object`

### `releaseSingleInstance(inst)`

- `_ptrSize` — constant
- `TBPF_NOPROGRESS` · `0`
- `TBPF_INDETERMINATE` · `1`
- `TBPF_NORMAL` · `2`
- `TBPF_ERROR` · `4`
- `TBPF_PAUSED` · `8`
### `_createTaskbarList3()`

### `_comCall(pInterface, vtableIdx, args...)`

- `_taskbarList3` · `?`
- `_taskbarInited` · `false`
### `_getTaskbar()`

### `setTaskbarProgress(window, completed, total)`

### `setTaskbarProgressState(window, flags)`

### `createOwnedWindow(parent, title, width, height, options)`

### `showModalDialog(parent, title, width, height, options, setupFn)`

### `closeModalDialog(dialog)`

### `setWindowOwner(child, parent)`

### `getWindowMonitor(window)`

### `centerOnMonitor(window)`

### `moveToMonitor(window, hMonitor)`

### `getWindowDpi(window)`

### `extendFrameIntoClientArea(window, margins)`

### `enableGlassSheet(window)`

- `HTCLIENT` · `1`
- `HTCAPTION` · `2`
- `HTSYSMENU` · `3`
- `HTMINBUTTON` · `8`
- `HTMAXBUTTON` · `9`
- `HTLEFT` · `10`
- `HTRIGHT` · `11`
- `HTTOP` · `12`
- `HTTOPLEFT` · `13`
- `HTTOPRIGHT` · `14`
- `HTBOTTOM` · `15`
- `HTBOTTOMLEFT` · `16`
- `HTBOTTOMRIGHT` · `17`
- `HTCLOSE` · `20`
- `_WM_NCHITTEST` · `132`
### `onNcHitTest(window, handler)`

### `customChromeHitTest(mx, my, width, height, borderSize, captionHeight)`

### `setWindowRgn(window, hRgn, redraw)`

### `createRoundRectRgn(left, top, right, bottom, rx, ry)`

### `_escapeXml(s)`

### `_escapePsString(s)`

### `showToastNotification(title, message, options)`

### `showToastWithFallback(title, message, options)`

### `addJumpListTask(title, path, arguments, iconPath, iconIndex, description)`

### `clearJumpList()`

### `addJumpListRecentFile(filePath)`

### `registerFileAssociation(extension, progId, description, command, iconPath)`

### `unregisterFileAssociation(extension, progId)`

### `refreshShellAssociations()`

### `addSearchFolder(folderPath, scope)`

### `searchFiles(query, maxResults)`

### `searchFilesWithProperty(query, property, maxResults)`

## Module: `gui-native-win-close`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
### `cleanupFrameBuffers(window)`

> returns `:int`

### `cleanupOpenGL(window)`

> returns `:int`

### `cleanupVulkan(window)`

> returns `:int`

## Module: `gui-native-win-d3d11`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `COM_RELEASE` · `2`
- `IDXGISWAPCHAIN_PRESENT` · `8`
- `IDXGISWAPCHAIN_GETBUFFER` · `9`
- `IDXGISWAPCHAIN_RESIZEBUFFERS` · `13`
- `ID3D11DEVICECONTEXT_MAP` · `14`
- `ID3D11DEVICECONTEXT_UNMAP` · `15`
- `ID3D11DEVICECONTEXT_UPDATESUBRESOURCE` · `48`
- `ID3D11DEVICE_CREATETEXTURE2D` · `5`
- `DXGI_FORMAT_B8G8R8A8_UNORM` · `87`
- `DXGI_USAGE_RENDER_TARGET_OUTPUT` · `32`
- `DXGI_SWAP_EFFECT_DISCARD` · `0`
- `D3D11_SDK_VERSION` · `7`
- `D3D_DRIVER_TYPE_HARDWARE` · `1`
- `D3D_DRIVER_TYPE_WARP` · `5`
- `D3D11_CREATE_DEVICE_BGRA_SUPPORT` · `32`
- `DIB_RGB_COLORS` · `0`
- `BI_RGB` · `0`
### `_iidD3d11Texture2D()`

### `initD3d112DLayer(window)`

> returns `:object`

### `presentFrameViaD3d11(window)`

> returns `:object`

### `releaseD3d11(window)`

> returns `:int`

## Module: `gui-native-win-ddraw`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `COM_RELEASE` · `2`
- `IDirectDraw7_CreateSurface` · `6`
- `IDirectDraw7_SetCooperativeLevel` · `20`
- `IDirectDrawSurface7_GetDC` · `17`
- `IDirectDrawSurface7_ReleaseDC` · `26`
- `DDSCL_NORMAL` · `8`
- `DDSD_CAPS` · `1`
- `DDSCAPS_PRIMARYSURFACE` · `512`
- `SRCCOPY` · `13369376`
### `_iidDirectDraw7()`

### `_ddrawCreatePrimarySurface(ddrawObj)`

> returns `:object`

### `_ddrawPresentViaPrimarySurface(window)`

> returns `:object`

### `initDdraw2DLayer(window)`

> returns `:object`

### `presentFrameViaDdraw(window)`

> returns `:object`

### `releasePrimarySurface(window)`

> returns `:int`

## Module: `gui-native-win-frame`

- `windows` · `import(...)`
- `guiResolution` · `import(...)`
- `guiThread` · `import(...)`
### `_windowClientSize(window)`

### `_ensureBackbuffer(window, targetHdc)`

### `_retryEnsureBackbuffer(window, targetHdc)`

### `prepareBeginFrameTarget(window, targetHdc)`

> returns `:object`

## Module: `gui-native-win-icons`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `WM_SETICON` · `128`
- `ICON_SMALL` · `0`
- `ICON_BIG` · `1`
- `ICON_SMALL2` · `2`
- `IMAGE_ICON` · `1`
- `LR_DEFAULTSIZE` · `64`
- `LR_LOADFROMFILE` · `16`
- `GCLP_HICON` — constant
- `GCLP_HICONSM` — constant
- `SWP_NOMOVE` · `2`
- `SWP_NOSIZE` · `1`
- `SWP_NOZORDER` · `4`
- `SWP_NOACTIVATE` · `16`
- `SWP_FRAMECHANGED` · `32`
### `_loadWindowsIconSized(iconSpec, width, height)`

> returns `:int`

### `resolveWindowIcons(opts)`

> returns `:object`

### `applyWindowIcons(hwnd, iconState)`

> returns `:int`

## Module: `gui-native-win-opengl`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `PFD_DRAW_TO_WINDOW` · `4`
- `PFD_SUPPORT_OPENGL` · `32`
- `PFD_DOUBLEBUFFER` · `1`
- `GL_COLOR_BUFFER_BIT` · `16384`
- `GL_BGRA_EXT` · `32993`
- `GL_UNSIGNED_BYTE` · `5121`
- `GL_UNPACK_ALIGNMENT` · `3317`
- `GL_RGBA8` · `32856`
- `GL_RENDERBUFFER` · `36161`
- `GL_READ_FRAMEBUFFER` · `36008`
- `GL_DRAW_FRAMEBUFFER` · `36009`
- `GL_COLOR_ATTACHMENT0` · `36064`
- `GL_NEAREST` · `9728`
- `GL_LINEAR` · `9729`
- `DIB_RGB_COLORS` · `0`
- `BI_RGB` · `0`
- `_glProcs` · `{}`
### `_glGetProc(name)`

### `_glCall(proc, args...)`

### `_initOpenGLWithContext(window, hdcHandle, chosen)`

> returns `:object`

### `initOpenGL2DLayer(window)`

### `_ensureGlPixelBuffer(window)`

### `_ensureGlBmiHeader(window)`

### `_ensureGlFbo(window, dw, dh)`

> returns `:bool`

### `_presentGlScaled(window, hdc, ctx, r)`

> returns `:object`

### `presentFrameOpenGL(window)`

> returns `:object`

### `releaseGlPixelBuffer(window)`

> returns `:int`

## Module: `gui-native-win-poll`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
### `pollWindowMessages(window)`

> returns `:object`

## Module: `gui-native-win-present`

- `windows` · `import(...)`
- `guiNativeWinOpenGL` · `import(...)`
- `guiThread` · `import(...)`
- `SRCCOPY` · `13369376`
- `COLORONCOLOR` · `3`
- `HALFTONE` · `4`
### `presentFrameViaOpenGL(window)`

### `presentFrameViaGdi(window)`

## Module: `gui-native-win-probe`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `_probeCacheReady` · `false`
- `_probeCache` · `?`
- `_probeDdrawTiming` · `0`
- `_probeD3d9Timing` · `0`
- `_probeD3d11Timing` · `0`
- `_probeOpenGLTiming` · `0`
- `_probeVulkanTiming` · `0`
### `_probeDdraw()`

### `_probeD3d9()`

### `_probeD3d11()`

### `_probeOpenGL()`

### `_probeVulkan()`

### `_pendingProbe(library, backend)`

> returns `:object`

### `_pendingDdrawProbe()`

> returns `:object`

### `pendingProbeSet()`

> returns `:object`

### `probeAllDlls()`

### `probe2DGraphicsStack()`

> returns `:object`

### `probe2DGdiOnly()`

> returns `:object`

### `probeNoGpu()`

> returns `:object`

### `select2DLayer(opts, opengl, ddraw, vulkan, d3d11)`

> returns `:atom`

### `getProbeTimings()`

> returns `:object`

### `select3DLayer(opts, d3d9)`

> returns `:atom`

## Module: `gui-native-win-vulkan`

- `guiThread` · `import(...)`
## Module: `gui-native-win-vulkan-constants`

- `windows` · `import(...)`
- `VK_STRUCTURE_TYPE_APPLICATION_INFO` · `0`
- `VK_STRUCTURE_TYPE_INSTANCE_CREATE_INFO` · `1`
- `VK_STRUCTURE_TYPE_WIN32_SURFACE_CREATE_INFO_KHR` · `1000009000`
- `VK_STRUCTURE_TYPE_DEVICE_QUEUE_CREATE_INFO` · `2`
- `VK_STRUCTURE_TYPE_DEVICE_CREATE_INFO` · `3`
- `VK_STRUCTURE_TYPE_SWAPCHAIN_CREATE_INFO_KHR` · `1000001000`
- `VK_STRUCTURE_TYPE_PRESENT_INFO_KHR` · `1000001001`
- `VK_STRUCTURE_TYPE_COMMAND_POOL_CREATE_INFO` · `39`
- `VK_STRUCTURE_TYPE_COMMAND_BUFFER_ALLOCATE_INFO` · `40`
- `VK_STRUCTURE_TYPE_COMMAND_BUFFER_BEGIN_INFO` · `42`
- `VK_STRUCTURE_TYPE_SUBMIT_INFO` · `4`
- `VK_STRUCTURE_TYPE_FENCE_CREATE_INFO` · `8`
- `VK_STRUCTURE_TYPE_SEMAPHORE_CREATE_INFO` · `9`
- `VK_STRUCTURE_TYPE_IMAGE_MEMORY_BARRIER` · `45`
- `VK_STRUCTURE_TYPE_BUFFER_CREATE_INFO` · `12`
- `VK_STRUCTURE_TYPE_MEMORY_ALLOCATE_INFO` · `5`
- `VK_SUCCESS` · `0`
- `VK_API_VERSION_1_0` · `4194304`
- `VK_QUEUE_GRAPHICS_BIT` · `1`
- `VK_KHR_SURFACE_EXTENSION_NAME` · `'VK_KHR_surface'`
- `VK_KHR_WIN32_SURFACE_EXTENSION_NAME` · `'VK_KHR_win32_surface'`
- `VK_KHR_SWAPCHAIN_EXTENSION_NAME` · `'VK_KHR_swapchain'`
- `VK_IMAGE_LAYOUT_UNDEFINED` · `0`
- `VK_IMAGE_LAYOUT_TRANSFER_DST_OPTIMAL` · `7`
- `VK_IMAGE_LAYOUT_PRESENT_SRC_KHR` · `1000001002`
- `VK_PIPELINE_STAGE_TOP_OF_PIPE_BIT` · `1`
- `VK_PIPELINE_STAGE_TRANSFER_BIT` · `4096`
- `VK_PIPELINE_STAGE_BOTTOM_OF_PIPE_BIT` · `8192`
- `VK_ACCESS_TRANSFER_WRITE_BIT` · `2048`
- `VK_COMMAND_POOL_CREATE_RESET_COMMAND_BUFFER_BIT` · `2`
- `VK_COMMAND_BUFFER_LEVEL_PRIMARY` · `0`
- `VK_COMMAND_BUFFER_USAGE_ONE_TIME_SUBMIT_BIT` · `1`
- `VK_BUFFER_USAGE_TRANSFER_SRC_BIT` · `1`
- `VK_MEMORY_PROPERTY_HOST_VISIBLE_BIT` · `2`
- `VK_MEMORY_PROPERTY_HOST_COHERENT_BIT` · `4`
- `VK_PRESENT_MODE_FIFO_KHR` · `2`
- `VK_PRESENT_MODE_MAILBOX_KHR` · `1`
- `VK_PRESENT_MODE_IMMEDIATE_KHR` · `0`
- `VK_COMPOSITE_ALPHA_OPAQUE_BIT_KHR` · `1`
- `VK_IMAGE_USAGE_COLOR_ATTACHMENT_BIT` · `16`
- `VK_IMAGE_USAGE_TRANSFER_DST_BIT` · `8`
- `VK_IMAGE_USAGE_TRANSFER_SRC_BIT` · `4`
- `VK_SHARING_MODE_EXCLUSIVE` · `0`
- `VK_SURFACE_TRANSFORM_IDENTITY_BIT_KHR` · `1`
- `VK_FENCE_CREATE_SIGNALED_BIT` · `1`
- `VK_IMAGE_ASPECT_COLOR_BIT` · `1`
- `VK_QUEUE_FAMILY_IGNORED` · `4294967295`
- `VK_FORMAT_B8G8R8A8_UNORM` · `44`
- `VK_FORMAT_B8G8R8A8_SRGB` · `50`
- `VK_COLOR_SPACE_SRGB_NONLINEAR_KHR` · `0`
- `VK_IMAGE_TYPE_2D` · `1`
- `VK_IMAGE_TILING_OPTIMAL` · `0`
- `VK_SAMPLE_COUNT_1_BIT` · `1`
- `VK_STRUCTURE_TYPE_IMAGE_CREATE_INFO` · `14`
- `VK_MEMORY_PROPERTY_DEVICE_LOCAL_BIT` · `1`
- `VK_IMAGE_LAYOUT_TRANSFER_SRC_OPTIMAL` · `6`
- `VK_ACCESS_TRANSFER_READ_BIT` · `2048`
- `VK_FILTER_NEAREST` · `0`
- `VK_FILTER_LINEAR` · `1`
- `DIB_RGB_COLORS` · `0`
- `BI_RGB` · `0`
### `_vkZeros(n)`

### `_vkWritePtr(address, value)`

### `_default(v, d)`

### `_vkGetProc(instance, name)`

> returns `:object`

### `_vkCall(proc, args...)`

### `_vkCallOk?(r)`

> returns `:bool`

## Module: `gui-native-win-vulkan-init`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
### `_vkFindPresentQueue(physicalDevice, surface, propsPtr, queueCount, idx)`

> returns `:object`

### `_vkFindPresentQueueForDevice(physicalDevice, surface)`

> returns `:object`

### `_vkFindPresentQueueAcrossDevices(devicesPtr, deviceCount, surface, idx)`

### `initVulkan2DLayer(window)`

> returns `:object`

### `createVulkanDevice(instance, physicalDevice, queueFamily)`

> returns `:object`

## Module: `gui-native-win-vulkan-present`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
### `presentFrameVulkan(window)`

> returns `:object`

## Module: `gui-native-win-vulkan-swapchain`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
### `_querySurfaceCaps(instance, physicalDevice, surface)`

### `_querySurfaceFormat(instance, physicalDevice, surface)`

### `createSwapchain(window)`

### `createVulkanCommandResources(window)`

> returns `:object`

### `_findMemoryType(physicalDevice, typeBits, properties)`

### `createStagingBuffer(window)`

> returns `:object`

### `createScaleImage(window)`

> returns `:object`

### `destroyVulkanSwapchain(window)`

> returns `:int`

### `initVulkanSwapchain(window)`

## Module: `gui-print`

- `sys` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `PD_ALLPAGES` · `0`
- `PD_SELECTION` · `1`
- `PD_PAGENUMS` · `2`
- `PD_NOSELECTION` · `4`
- `PD_NOPAGENUMS` · `8`
- `PD_COLLATE` · `16`
- `PD_PRINTTOFILE` · `32`
- `PD_PRINTSETUP` · `64`
- `PD_NOWARNING` · `128`
- `PD_RETURNDC` · `256`
- `PD_RETURNIC` · `512`
- `PD_RETURNDEFAULT` · `1024`
- `PD_SHOWHELP` · `2048`
- `PD_USEDEVMODECOPIES` · `262144`
- `PD_DISABLEPRINTTOFILE` · `524288`
- `PD_HIDEPRINTTOFILE` · `1048576`
- `PD_CURRENTPAGE` · `4194304`
- `DI_APPBANDING` · `1`
- `_PRINTDLGW_SIZE` — constant
- `_PD_OFF_HWNDOWNER` · `8`
- `_PD_OFF_HDEVMODE` · `16`
- `_PD_OFF_HDEVNAMES` · `24`
- `_PD_OFF_HDC` · `32`
- `_PD_OFF_FLAGS` · `40`
- `_PD_OFF_FROMPAGE` · `44`
- `_PD_OFF_TOPAGE` · `46`
- `_PD_OFF_MINPAGE` · `48`
- `_PD_OFF_MAXPAGE` · `50`
- `_PD_OFF_NCOPIES` · `52`
- `_DOCINFOW_SIZE` — constant
- `_DI_OFF_DOCNAME` — constant
- `_DI_OFF_OUTPUT` — constant
- `_DI_OFF_DATATYPE` — constant
- `_DI_OFF_FWTYPE` — constant
### `_writeU16(base, value)`

### `_readU16(base)`

### `showPrintDialog(options)`

> returns `:object`

### `startDoc(hDC, docName, outputFile)`

> returns `:object`

### `startPage(hDC)`

> returns `:bool`

### `endPage(hDC)`

> returns `:bool`

### `endDoc(hDC)`

> returns `:bool`

### `abortDoc(hDC)`

### `deleteDC(hDC)`

### `printTextOut(hDC, x, y, text)`

### `printMoveTo(hDC, x, y)`

### `printLineTo(hDC, x, y)`

### `printRectangle(hDC, left, top, right, bottom)`

### `printEllipse(hDC, left, top, right, bottom)`

### `printSetFont(hDC, height, weight, italic, fontName)`

### `printSetTextColor(hDC, r, g, b)`

### `printSetBkMode(hDC, mode)`

### `printSetPen(hDC, style, width, r, g, b)`

### `printDeleteObject(hObj)`

- `DEVCAP_HORZRES` · `8`
- `DEVCAP_VERTRES` · `10`
- `DEVCAP_LOGPIXELSX` · `88`
- `DEVCAP_LOGPIXELSY` · `90`
- `DEVCAP_PHYSICALWIDTH` · `110`
- `DEVCAP_PHYSICALHEIGHT` · `111`
- `DEVCAP_PHYSICALOFFSETX` · `112`
- `DEVCAP_PHYSICALOFFSETY` · `113`
### `getDeviceCaps(hDC, index)`

### `getPrinterPageSize(hDC)`

> returns `:object`

### `createPreviewDC(width, height)`

> returns `:object`

### `destroyPreviewDC(preview)`

### `printToFile(outputPath, docName, renderFn)`

> returns `:object`

### `printDocument(options, renderPageFn)`

> returns `:object`

## Module: `gui-raster`

- `guiLighting` · `import(...)`
- `gui3dmath` · `import(...)`
- `guiThread` · `import(...)`
- `threadLib` · `import(...)`
- `_OK` · `{1 entries}`
### `_pointOutCode(window, p)`

> returns `:bool`

### `_lineVisible?(window, a, b)`

### `_min2(a, b)`

### `_max2(a, b)`

### `_triVisible?(window, a, b, c)`

> returns `:bool`

### `_triArea2x(pa, pb, pc)`

> returns `:int`

### `_clipPolyEdge(verts, n, getVal, limit, isMin)`

> returns `:list`

### `_clipPolyToViewport(window, verts, n)`

> returns `:list`

### `_clipAndFillTriangle(deps, window, p0, p1, p2, color)`

> returns `:bool`

### `_lerpX(pa, pb, y)`

### `_buildEdgeTable(p0, p1, p2)`

> returns `:object`

### `_drawScanline(deps, window, y, xa, xb, color)`

> returns `?`

### `_edgeSlope(pa, pb)`

> returns `:int`

### `_fillScanStepped(deps, window, y, splitY, maxY, xaTop, xaBot, xb, slopeATop, slopeABot, slopeB, color)`

### `_fillScan(deps, window, y, maxY, p0, p1, p2, color)`

### `_sortTriByY(p0, p1, p2)`

> returns `:list`

### `drawTriangleFilled(deps, window, p0, p1, p2, color, borderColor)`

### `drawTriangleFilledAA(deps, window, p0, p1, p2, color, borderColor, bgColor)`

### `_concatLists(left, right, i, out)`

### `_compactTrisInPlace(tris)`

### `_insertionSortRange(arr, lo, hi)`

### `_partition(arr, lo, hi)`

### `_sortDepthRange(arr, lo, hi)`

### `_slice(arr, start, end, acc)`

### `_sortDepthInPlace(tris, count)`

### `_sortDepth(tris)`

### `_vecSub(deps, a, b)`

### `_vecCross(deps, a, b)`

### `_vecDot(a, b)`

### `_vecLen(v)`

### `_vecNormalize(deps, v)`

- `_projParamsCache` · `?`
- `_projCacheW` · `?`
- `_projCacheH` · `?`
- `_projCacheFov` · `?`
- `_projCacheZ` · `?`
- `_projCacheMode` · `?`
- `_projCacheOrtho` · `?`
### `_projectionParams(window, camera)`

### `_projectPointFast(p, params)`

> returns `:object`

### `projectVertices(deps, window, verts, camera)`

### `_computeMeshBounds(verts, i, mnX, mxX, mnY, mxY, mnZ, mxZ)`

> returns `:object`

### `_meshBounds(verts)`

> returns `:object`

### `_ensureMeshBounds(mesh)`

### `_transformBoundsQuick(localBounds, transform)`

> returns `:object`

### `_sphereInFrustum?(bounds, params, farPlane)`

> returns `:bool`

### `computeLightParams(deps, light)`

### `faceShadeGeneric(deps, lp, pa3, pb3, pc3)`

### `_colorR(c)`

> returns `:bool`

### `_colorG(c)`

> returns `:bool`

### `_colorB(c)`

> returns `:bool`

### `_shadeColor(deps, color, intensity)`

### `_frontFacing?(pa, pb, pc)`

> returns `:bool`

### `_computeDepth(pa3, pb3, pc3, pd3)`

### `_buildTriangle(pa, pb, pc, depth, color)`

> returns `:object`

### `_shouldCullTriangle?(backfaceCulling, pa, pb, pc, window)`

> returns `:bool`

### `_processFace(deps, window, i, faces, faceColors, verts, projected, defaultFaceColor, faceShade, backfaceCulling, acc)`

### `_collectFaces(deps, window, faces, faceColors, verts, projected, defaultFaceColor, faceShade, backfaceCulling, i, acc)`

### `_collectFacesRange(deps, window, faces, faceColors, verts, projected, defaultFaceColor, faceShade, backfaceCulling, start, end, acc)`

### `_drawTriangles(deps, window, tris, i, count, drawn)`

### `drawMeshSolid(deps, window, mesh, transform, camera, color, light, borderColor)`

> returns `:object`

### `drawMeshWireframe(deps, window, mesh, transform, camera, color)`

> returns `:object`

### `_processFaceLit(deps, window, i, faces, faceColors, verts, projected, defaultFaceColor, colorFn, backfaceCulling, acc)`

### `_collectFacesLit(deps, window, faces, faceColors, verts, projected, defaultFaceColor, colorFn, backfaceCulling, i, acc)`

### `_collectFacesLitRange(deps, window, faces, faceColors, verts, projected, defaultFaceColor, colorFn, backfaceCulling, start, end, acc)`

### `drawMeshLit(deps, window, mesh, transform, camera, color, scene, material, borderColor)`

> returns `:object`

## Module: `gui-render`

- `guiThread` · `import(...)`
- `threadLib` · `import(...)`
### `Renderer3D(deps, window, options)`

> returns `:object`

## Module: `gui-resolution`

- `guiThread` · `import(...)`
### `_res(window)`

### `_hasRes?(window)`

### `setDesignResolution(window, logicalWidth, logicalHeight, options)`

### `clearDesignResolution(window)`

### `hasDesignResolution?(window)`

### `designWidth(window)`

### `designHeight(window)`

### `physicalWidth(window)`

### `physicalHeight(window)`

### `scaleX(window)`

### `scaleY(window)`

### `offsetX(window)`

### `offsetY(window)`

### `physicalToLogical(window, px, py)`

> returns `:object`

### `logicalToPhysical(window, lx, ly)`

> returns `:object`

### `updatePhysicalSize(window, physW, physH)`

### `_recomputeScale(window)`

> returns `:int`

## Module: `gui-shader`

- `_draw` · `import(...)`
- `threadLib` · `import(...)`
- `m` · `import(...)`
- `col` · `import(...)`
- `noise` · `import(...)`
- `sdf` · `import(...)`
- `codegen` · `import(...)`
- `PI` — constant
- `TAU` — constant
- `HALF_PI` — constant
- `E` — constant
- `DEG2RAD` — constant
- `RAD2DEG` — constant
- `SQRT2` — constant
### `fract(x)`

### `mod(x, y)`

### `sign(x)`

### `abs2(x)`

### `clamp(x, lo, hi)`

### `saturate(x)`

### `lerpFloat(a, b, t)`

### `inverseLerp(a, b, x)`

### `remap(x, inLo, inHi, outLo, outHi)`

### `step(edge, x)`

### `smoothstep(edge0, edge1, x)`

### `smootherstep(edge0, edge1, x)`

### `min2(a, b)`

### `max2(a, b)`

### `sqr(x)`

### `sqrt(x)`

### `lerp(a, b, t)`

### `atan2(y, x)`

### `pingpong(t, length)`

### `degToRad(d)`

### `radToDeg(r)`

### `easeInQuad(t)`

### `easeOutQuad(t)`

### `easeInOutQuad(t)`

### `easeInCubic(t)`

### `easeOutCubic(t)`

### `easeInOutCubic(t)`

### `easeInSine(t)`

### `easeOutSine(t)`

### `easeInOutSine(t)`

### `easeInExpo(t)`

### `easeOutExpo(t)`

### `easeOutElastic(t)`

### `easeOutBounce(t)`

### `vec2(x, y)`

### `dot2(a, b)`

### `length2(v)`

### `distance2(a, b)`

### `normalize2(v)`

### `rotate2(v, angle)`

### `scale2(v, s)`

### `add2(a, b)`

### `sub2(a, b)`

### `lerp2(a, b, t)`

### `negate2(v)`

### `abs2v(v)`

### `min2v(a, b)`

### `max2v(a, b)`

### `floor2(v)`

### `fract2(v)`

### `reflect2(v, n)`

### `toPolar(v)`

### `fromPolar(r, theta)`

### `vec3(x, y, z)`

### `add3(a, b)`

### `sub3(a, b)`

### `scale3(v, s)`

### `dot3(a, b)`

### `length3(v)`

### `distance3(a, b)`

### `normalize3(v)`

### `cross3(a, b)`

### `lerp3(a, b, t)`

### `negate3(v)`

### `reflect3(v, n)`

### `packRGB(r, g, b)`

### `unpackRGB(c)`

### `colorR(c)`

### `colorG(c)`

### `colorB(c)`

### `mix(a, b, t)`

### `mix3(a, b, c, t)`

### `brighten(c, amount)`

### `darken(c, amount)`

### `invert(c)`

### `grayscale(c)`

### `overlay(fg, bg, alpha)`

### `hsl2rgb(h, s, l)`

### `rgb2hsl(c)`

### `hsv2rgb(h, s, v)`

### `rgb2hsv(c)`

### `floatStr(c)`

### `cosinePalette(t, a, b, c, d)`

### `contrast(c, amount)`

### `sepia(c)`

### `blendMultiply(a, b)`

### `blendScreen(a, b)`

### `blendAdd(a, b)`

### `hash(seed)`

### `hash2(a, b)`

### `hash3(a, b, c)`

### `noise2D(x, y)`

### `fbm(x, y, octaves?)`

### `noiseGrid2DParallel(w, h, scaleFn, numWorkers)`

### `fbmGrid2DParallel(w, h, scaleFn, octaves, numWorkers)`

### `sdCircle(px, py, cx, cy, r)`

### `sdBox(px, py, cx, cy, hw, hh)`

### `sdLine(px, py, ax, ay, bx, by)`

### `sdRoundedBox(px, py, cx, cy, hw, hh, r)`

### `sdfFill(d, color)`

### `sdfSmoothFill(d, color, bg, edge)`

### `sdfStroke(d, thickness, color)`

### `sdfGlow(d, color, intensity, radius)`

### `sdUnion(d1, d2)`

### `sdIntersect(d1, d2)`

### `sdSubtract(d1, d2)`

### `sdSmoothUnion(d1, d2, k)`

### `sdSmoothIntersect(d1, d2, k)`

### `sdSmoothSubtract(d1, d2, k)`

### `sdAnnular(d, r)`

### `sdRepeat2(px, py, cx, cy)`

### `checkerboard(x, y, size)`

### `stripes(x, y, angle, width)`

### `grid(x, y, size, thickness)`

### `dots(x, y, spacing, radius)`

### `voronoi(x, y, scale_)`

### `glslVersion(ver?)`

### `glslPrecision(prec?, type?)`

### `glslStdUniforms()`

### `glslUniform(type, name)`

### `glslUniforms(uniforms)`

### `glslIn(type, name)`

### `glslOut(type, name)`

### `glslQuadVertex()`

### `glslQuadVertexCompat()`

### `glslFragmentWrap(body, version?)`

### `glslMathLib()`

### `hlslStdCBuffer()`

### `hlslCBuffer(name, uniforms)`

### `hlslQuadVertex()`

### `hlslFragmentWrap(body)`

### `hlslMathLib()`

### `submitWebGL(window, fragSource, vertSource?)`

### `drawWebGL(window, clearR?, clearG?, clearB?)`

### `renderWebGL(window, fragSource)`

### `compileGLSL(source, stage?, outputPath?)`

### `compileHLSL(source, profile?, entry?, outputPath?)`

### `compileDXC(source, profile?, entry?, outputPath?, spirv?)`

### `assembleGLSL(opts)`

### `assembleHLSL(opts)`

- `_registry` · `[]`
### `_registerShader(shader)`

> returns `?`

### `unregisterShader(shader)`

### `clearAll()`

### `destroyAll()`

> returns `:list`

### `registeredCount()`

### `cs Shader(fragment?, opts?)`

### `elapsed(shader)`

### `pause(shader)`

### `resume(shader)`

> returns `?`

### `reset(shader)`

> returns `?`

### `setUniform(shader, key, value)`

### `getUniform(shader, key)`

### `setResolution(shader, res)`

### `beginFrame(shader)`

### `endFrame(shader)`

### `dt(shader)`

### `isRunning(shader)`

### `frameCount(shader)`

### `render(window, shader, x, y, w, h)`

### `cs ShaderPass(shader, x, y, w, h)`

### `composePasses(window, passes)`

### `renderShader(window, shader, x, y, w, h)`

### `renderShaderLines(window, shader, x, y, w, h)`

### `renderGradientBands(window, gradientFn, x, y, w, h, time, bands?)`

### `renderGradient(window, gradientFn, x, y, w, h, time)`

### `renderHorizontalBands(window, gradientFn, x, y, w, h, time, bands?)`

### `renderColumns(window, columns, ox, oy, h)`

### `updateColumns(columns, h, t, rate?)`

### `createBuffer(w, h)`

> returns `:object`

### `clearBuffer(buf, color?)`

### `setPixel(buf, x, y, color)`

> returns `?`

### `getPixel(buf, x, y)`

> returns `:int`

### `renderBuffer(window, buf, ox, oy)`

### `renderShaderToBuffer(buf, shader)`

### `renderShaderToBufferParallel(buf, shader, numWorkers)`

### `renderParallel(window, shader, x, y, w, h, numWorkers)`

## Module: `gui-shader-codegen`

- `threadLib` · `import(...)`
### `glslVersion(ver?)`

> returns `:string`

### `glslPrecision(prec?, type?)`

> returns `:string`

### `glslStdUniforms()`

> returns `:string`

### `glslUniform(type, name)`

> returns `:string`

### `glslUniforms(uniforms)`

### `glslIn(type, name)`

> returns `:string`

### `glslOut(type, name)`

> returns `:string`

### `glslQuadVertex()`

### `glslQuadVertexCompat()`

> returns `:string`

### `glslFragmentWrap(body, version?)`

### `glslMathLib()`

> returns `:string`

### `hlslStdCBuffer()`

> returns `:string`

### `hlslCBuffer(name, uniforms)`

### `hlslQuadVertex()`

> returns `:string`

### `hlslFragmentWrap(body)`

### `hlslMathLib()`

> returns `:string`

### `submitWebGL(window, fragSource, vertSource?)`

> returns `:object`

### `drawWebGL(window, clearR?, clearG?, clearB?)`

### `renderWebGL(window, fragSource)`

### `compileGLSL(source, stage?, outputPath?)`

### `compileHLSL(source, profile?, entry?, outputPath?)`

### `compileDXC(source, profile?, entry?, outputPath?, spirv?)`

### `assembleGLSL(opts)`

### `assembleHLSL(opts)`

## Module: `gui-shader-color`

- `guiColor` · `import(...)`
- `m` · `import(...)`
- `threadLib` · `import(...)`
### `packRGB(r, g, b)`

### `unpackRGB(c)`

> returns `:object`

### `colorR(c)`

> returns `:bool`

### `colorG(c)`

> returns `:bool`

### `colorB(c)`

> returns `:bool`

### `mix(a, b, t)`

### `mix3(a, b, c, t)`

### `brighten(c, amount)`

### `darken(c, amount)`

### `invert(c)`

### `grayscale(c)`

### `overlay(fg, bg, alpha)`

### `hsl2rgb(h, s, l)`

### `rgb2hsl(c)`

> returns `:object`

### `hsv2rgb(h, s, v)`

### `rgb2hsv(c)`

> returns `:object`

### `floatStr(v)`

### `cosinePalette(t, a, b, c, d)`

### `contrast(c, amount)`

### `sepia(c)`

### `blendMultiply(a, b)`

### `blendScreen(a, b)`

### `blendAdd(a, b)`

## Module: `gui-shader-math`

- `threadLib` · `import(...)`
- `PI` · `3.14159265358979`
- `TAU` · `6.28318530717959`
- `HALF_PI` · `1.5707963267949`
- `E` · `2.71828182845905`
- `DEG2RAD` — constant
- `RAD2DEG` — constant
- `SQRT2` · `1.4142135623731`
### `fract(x)`

### `mod(x, y)`

### `sign(x)`

> returns `:int`

### `abs2(x)`

> returns `:int`

### `clamp(x, lo, hi)`

### `saturate(x)`

### `lerpFloat(a, b, t)`

### `inverseLerp(a, b, x)`

> returns `:float`

### `remap(x, inLo, inHi, outLo, outHi)`

### `step(edge, x)`

> returns `:int`

### `smoothstep(edge0, edge1, x)`

### `smootherstep(edge0, edge1, x)`

### `min2(a, b)`

### `max2(a, b)`

### `sqr(x)`

### `sqrt(x)`

### `lerp(a, b, t)`

### `atan2(y, x)`

### `pingpong(t, length)`

### `degToRad(d)`

### `radToDeg(r)`

### `easeInQuad(t)`

### `easeOutQuad(t)`

### `easeInOutQuad(t)`

> returns `:float`

### `easeInCubic(t)`

### `easeOutCubic(t)`

### `easeInOutCubic(t)`

> returns `:float`

### `easeInSine(t)`

> returns `:float`

### `easeOutSine(t)`

### `easeInOutSine(t)`

> returns `:int`

### `easeInExpo(t)`

> returns `:float`

### `easeOutExpo(t)`

> returns `:float`

### `easeOutElastic(t)`

> returns `:float`

### `easeOutBounce(t)`

> returns `:float`

### `vec2(x, y)`

> returns `:object`

### `dot2(a, b)`

### `length2(v)`

### `distance2(a, b)`

### `normalize2(v)`

### `rotate2(v, angle)`

### `scale2(v, s)`

### `add2(a, b)`

### `sub2(a, b)`

### `lerp2(a, b, t)`

### `negate2(v)`

### `abs2v(v)`

### `min2v(a, b)`

### `max2v(a, b)`

### `floor2(v)`

### `fract2(v)`

### `reflect2(v, n)`

### `toPolar(v)`

> returns `:object`

### `fromPolar(r, theta)`

### `vec3(x, y, z)`

> returns `:object`

### `add3(a, b)`

### `sub3(a, b)`

### `scale3(v, s)`

### `dot3(a, b)`

### `length3(v)`

### `distance3(a, b)`

### `normalize3(v)`

### `cross3(a, b)`

### `lerp3(a, b, t)`

### `negate3(v)`

### `reflect3(v, n)`

## Module: `gui-shader-noise`

- `m` · `import(...)`
- `threadLib` · `import(...)`
### `hash(seed)`

### `hash2(a, b)`

### `hash3(a, b, c)`

### `noise2D(x, y)`

### `fbm(x, y, octaves?)`

### `noiseGrid2DParallel(w, h, scaleFn, numWorkers)`

### `fbmGrid2DParallel(w, h, scaleFn, octaves, numWorkers)`

## Module: `gui-shader-sdf`

- `m` · `import(...)`
- `col` · `import(...)`
- `noise` · `import(...)`
- `threadLib` · `import(...)`
### `sdCircle(px, py, cx, cy, r)`

### `sdBox(px, py, cx, cy, hw, hh)`

### `sdLine(px, py, ax, ay, bx, by)`

### `sdRoundedBox(px, py, cx, cy, hw, hh, r)`

### `sdfFill(d, color)`

### `sdfSmoothFill(d, color, bg, edge)`

### `sdfStroke(d, thickness, color)`

### `sdfGlow(d, color, intensity, radius)`

### `sdUnion(d1, d2)`

### `sdIntersect(d1, d2)`

### `sdSubtract(d1, d2)`

### `sdSmoothUnion(d1, d2, k)`

### `sdSmoothIntersect(d1, d2, k)`

### `sdSmoothSubtract(d1, d2, k)`

### `sdAnnular(d, r)`

### `sdRepeat2(px, py, cx, cy)`

### `checkerboard(x, y, size)`

### `stripes(x, y, angle, width)`

### `grid(x, y, size, thickness)`

> returns `:int`

### `dots(x, y, spacing, radius)`

> returns `:int`

### `voronoi(x, y, scale_)`

> returns `:object`

## Module: `gui-systray`

- `std` · `import(...)`
- `windows` · `import(...)`
- `guiThread` · `import(...)`
- `NIM_ADD` · `0`
- `NIM_MODIFY` · `1`
- `NIM_DELETE` · `2`
- `NIM_SETVERSION` · `4`
- `NIF_MESSAGE` · `1`
- `NIF_ICON` · `2`
- `NIF_TIP` · `4`
- `NIF_INFO` · `16`
- `_NOTIFYICONDATA_SIZE` · `956`
- `WM_TRAYICON` · `32769`
### `_zeros(n)`

### `_writeUtf16At(baseAddr, text, maxBytes)`

- `_NID_OFF_HWND` · `8`
- `_NID_OFF_UID` · `16`
- `_NID_OFF_FLAGS` · `20`
- `_NID_OFF_CALLBACKMSG` · `24`
- `_NID_OFF_HICON` · `32`
- `_NID_OFF_TIP` · `40`
- `_NID_OFF_INFO` · `304`
- `_NID_OFF_INFOTITLE` · `820`
- `_NID_OFF_INFOFLAGS` · `948`
### `addTrayIcon(window, id, tooltip, hIcon)`

> returns `:bool`

### `removeTrayIcon(window, id)`

> returns `:bool`

### `updateTrayTooltip(window, id, tooltip)`

> returns `:bool`

### `showBalloonTip(window, id, title, message, iconFlag)`

> returns `:bool`

- `TRAY_WM_LBUTTONDOWN` · `513`
- `TRAY_WM_LBUTTONUP` · `514`
- `TRAY_WM_RBUTTONDOWN` · `516`
- `TRAY_WM_RBUTTONUP` · `517`
- `TRAY_WM_LBUTTONDBLCLK` · `515`
### `onTrayEvent(window, handler)`

### `minimizeToTray(window)`

### `restoreFromTray(window)`

## Module: `gui-test`

- `threadLib` · `import(...)`
- `windows` · `import(...)`
- `guiInput` · `import(...)`
- `INPUT_MOUSE` · `0`
- `INPUT_KEYBOARD` · `1`
- `INPUT_HARDWARE` · `2`
- `MOUSEEVENTF_MOVE` · `1`
- `MOUSEEVENTF_LEFTDOWN` · `2`
- `MOUSEEVENTF_LEFTUP` · `4`
- `MOUSEEVENTF_RIGHTDOWN` · `8`
- `MOUSEEVENTF_RIGHTUP` · `16`
- `MOUSEEVENTF_MIDDLEDOWN` · `32`
- `MOUSEEVENTF_MIDDLEUP` · `64`
- `MOUSEEVENTF_WHEEL` · `2048`
- `MOUSEEVENTF_ABSOLUTE` · `32768`
- `MOUSEEVENTF_VIRTUALDESK` · `16384`
- `KEYEVENTF_KEYDOWN` · `0`
- `KEYEVENTF_KEYUP` · `2`
- `KEYEVENTF_UNICODE` · `4`
- `KEYEVENTF_SCANCODE` · `8`
- `VK_RETURN` — constant
- `VK_ESCAPE` — constant
- `VK_TAB` — constant
- `VK_BACK` — constant
- `VK_SPACE` — constant
- `VK_LEFT` — constant
- `VK_UP` — constant
- `VK_RIGHT` — constant
- `VK_DOWN` — constant
- `VK_SHIFT` — constant
- `VK_CONTROL` — constant
- `VK_MENU` — constant
- `VK_DELETE` — constant
- `VK_HOME` — constant
- `VK_END` — constant
- `_INPUT_SIZE` · `40`
### `_buildMouseInput(dx, dy, flags, data)`

### `_buildKeyboardInput(vk, scan, flags)`

### `_sendInputs(inputs)`

### `testMouseMove(x, y)`

### `testMouseClick(x, y)`

### `testMouseRightClick(x, y)`

### `testMouseDoubleClick(x, y)`

### `testMouseDrag(x1, y1, x2, y2)`

### `testMouseWheel(delta)`

### `testKeyDown(vk)`

### `testKeyUp(vk)`

### `testKeyPress(vk)`

### `testTypeText(text)`

### `testKeyCombo(keys)`

### `testSuite(name)`

> returns `:object`

### `testCase(suite, name, testFn)`

### `testRun(suite)`

### `testAssert(result, condition, message)`

### `testAssertEqual(result, actual, expected, message)`

### `testReport(suite)`

### `testGetWindowRect(hwnd)`

> returns `:object`

### `testGetWindowText(hwnd)`

### `testIsWindowVisible?(hwnd)`

### `testGetForegroundWindow()`

### `testFindWindow(className, windowName)`

### `testScreenshot(hwnd)`

### `testFreeScreenshot(ss)`

## Module: `gui-theme`

- `windows` · `import(...)`
- `guiThread` · `import(...)`
### `_zeros(n)`

### `isDarkMode?()`

- `SPI_GETHIGHCONTRAST` · `66`
- `HCF_HIGHCONTRASTON` · `1`
### `isHighContrast?()`

> returns `:bool`

### `accentColor()`

> returns `:object`

## Module: `gui-thread`

- `threadLib` · `import(...)`
### `CommandQueue()`

> returns `:object`

### `FrameFence(workerCount)`

> returns `:object`

### `WorkerPool(numWorkers)`

> returns `:object`

### `StateGuard()`

> returns `:object`

### `parallelTransformVertices(vertices, transformFn, numWorkers)`

### `AsyncLoader(cmdQueue)`

> returns `:object`

### `FrameScheduler(pool, cmdQueue)`

> returns `:object`

### `initWindowThreading(window, options)`

### `threadingEnabled?(window)`

### `commandQueue(window)`

### `workerPool(window)`

### `scheduler(window)`

### `stateGuard(window)`

### `flushThreadedCommands(window)`

### `destroyWindowThreading(window)`

> returns `?`

## Module: `gui-video`

- `video` · `import(...)`
- `guiColor` · `import(...)`
- `guiShader` · `import(...)`
- `threadLib` · `import(...)`
### `frameToBuffer(frame)`

### `bufferToFrame(buf)`

### `renderFrame(drawCtx, window, frame, ox, oy)`

### `renderFrameScaled(drawCtx, window, frame, ox, oy, scale)`

### `bufferGrayscale(buf)`

### `bufferInvert(buf)`

### `bufferThreshold(buf, t)`

### `bufferBlend(bufA, bufB, alpha)`

### `bufferDiff(bufA, bufB)`

### `captureBuffer(buf)`

### `frameToBmpPixels(frame)`

## Module: `gui-web`

- `guiThread` · `import(...)`
### `_isFn?(value)`

### `_nextResourceId(webgl)`

### `_webRecord(window, op)`

### `createWindowState(title, width, height, frameMs, updateOnDispatch, options)`

> returns `:object`

### `showWindow(window)`

### `hideWindow(window)`

### `moveWindow(window, x, y)`

### `resizeWindow(window, width, height)`

### `setFullscreen(window, enabled)`

### `lockResize(window, locked)`

### `createCanvas(window, id, options)`

### `initWebGL(window, contextName, attrs)`

> returns `:object`

### `webglCreateShader(window, shaderType, source)`

> returns `:object`

### `webglCreateProgram(window, vertexShader, fragmentShader)`

> returns `:object`

### `webglUseProgram(window, program)`

> returns `:object`

### `webglClearColor(window, r, g, b, a)`

> returns `:object`

### `webglViewport(window, x, y, width, height)`

> returns `:object`

### `webglClear(window, mask, colorBufferBit)`

> returns `:object`

### `webglDrawArrays(window, mode, first, count, trianglesMode)`

> returns `:object`

### `webglFlush(window)`

> returns `:object`

### `colorToCss(packed)`

> returns `:string`

### `beginFrame(window)`

### `endFrame(window)`

### `runFrameLoop(window, onFrame)`

### `stopFrameLoop(window)`

### `dispatchEvent(window, evt)`

### `pollEvents(window)`

### `setAlpha(window, alpha)`

### `drawImageData(window, x, y, w, h, pixelBytes)`

## Module: `lib\GUI.oak`

- `_guiImportT0` · `nanotime(...)`
- `std` · `import(...)`
- `windows` · `import(...)`
- `linux` · `import(...)`
- `_guiImportT1` · `nanotime(...)`
- `guiMesh` · `import(...)`
- `guiRender` · `import(...)`
- `gui3dmath` · `import(...)`
- `gui2d` · `import(...)`
- `guiRaster` · `import(...)`
- `guiLighting` · `import(...)`
- `_guiImportT2` · `nanotime(...)`
- `guiWeb` · `import(...)`
- `guiNativeWin` · `import(...)`
- `guiNativeLinux` · `import(...)`
- `_guiImportT3` · `nanotime(...)`
- `guiDraw` · `import(...)`
- `guiColor` · `import(...)`
- `guiEvents` · `import(...)`
- `guiInput` · `import(...)`
- `guiGraph` · `import(...)`
- `guiForm` · `import(...)`
- `guiLoop` · `import(...)`
- `guiShader` · `import(...)`
- `guiFonts` · `import(...)`
- `guiVideo` · `import(...)`
- `guiResolution` · `import(...)`
- `guiCanvas` · `import(...)`
- `guiAcc` · `import(...)`
- `guiClipboard` · `import(...)`
- `guiFiledrop` · `import(...)`
- `guiAudio` · `import(...)`
- `guiGamepad` · `import(...)`
- `guiAA` · `import(...)`
- `guiDrawOps` · `import(...)`
- `guiGpuInfo` · `import(...)`
- `guiLeakDetect` · `import(...)`
- `guiTest` · `import(...)`
- `guiDialogs` · `import(...)`
- `guiMenus` · `import(...)`
- `guiPrint` · `import(...)`
- `guiTheme` · `import(...)`
- `guiSystray` · `import(...)`
- `guiThread` · `import(...)`
- `_guiImportT4` · `nanotime(...)`
- `_importTimings` · `{5 entries}`
### `getImportTimings()`

### `backend()`

> returns `:atom`

### `isWindows?()`

### `isLinux?()`

### `isWeb?()`

### `_ensureEventBus(window)`

### `eventBus(window)`

### `on(window, event, handler)`

### `once(window, event, handler)`

### `off(window, event, tokenOrHandler)`

### `emit(window, event, payload, onDone)`

### `listenerCount(window, event)`

### `clearListeners(window, event)`

### `onDispatch(window, handler)`

### `onceDispatch(window, handler)`

### `onRunStart(window, handler)`

### `onceRunStart(window, handler)`

### `onIdle(window, handler)`

### `onceIdle(window, handler)`

### `onFrame(window, handler)`

### `onceFrame(window, handler)`

### `onClosing(window, handler)`

### `onceClosing(window, handler)`

### `onClosed(window, handler)`

### `onceClosed(window, handler)`

### `onKeyDownEvent(window, handler)`

### `onceKeyDownEvent(window, handler)`

### `onKeyUpEvent(window, handler)`

### `onceKeyUpEvent(window, handler)`

### `_publish(window, event, payload)`

### `_clampByte(value)`

### `_clampOpacity(value)`

### `rgb(r, g, b)`

### `rgba(r, g, b, a, background)`

### `colorR(color)`

### `colorG(color)`

### `colorB(color)`

### `opacity(color, amount, background)`

- `GL_COLOR_BUFFER_BIT` · `16384`
- `GL_DEPTH_BUFFER_BIT` · `256`
- `GL_TRIANGLES` · `4`
### `createWindow(title, width, height, options)`

### `createWindowAsync(title, width, height, options)`

### `awaitWindow(future)`

### `show(window)`

### `hide(window)`

### `move(window, x, y)`

### `resize(window, width, height)`

### `scale(window, scaleX, scaleY)`

### `fullscreen(window, enabled)`

### `lockResize(window, locked)`

### `setAlwaysOnTop(window, enabled)`

### `setWindowOpacity(window, alpha)`

### `setWindowColorKey(window, colorKey)`

### `removeLayeredStyle(window)`

### `registerWindow(window)`

### `unregisterWindow(window)`

### `allWindows()`

### `pollAllWindows()`

### `closeAllWindows()`

### `anyWindowOpen?()`

### `saveWindowState(window)`

### `restoreWindowState(window, state)`

### `acquireSingleInstance(name)`

### `releaseSingleInstance(inst)`

### `setTaskbarProgress(window, completed, total)`

### `setTaskbarProgressState(window, flags)`

- `TBPF_NOPROGRESS` — constant
- `TBPF_INDETERMINATE` — constant
- `TBPF_NORMAL` — constant
- `TBPF_ERROR` — constant
- `TBPF_PAUSED` — constant
### `createOwnedWindow(parent, title, width, height, options)`

### `showModalDialog(parent, title, width, height, options, setupFn)`

### `closeModalDialog(dialog)`

### `setWindowOwner(child, parent)`

### `getWindowMonitor(window)`

### `centerOnMonitor(window)`

### `moveToMonitor(window, hMonitor)`

### `getWindowDpi(window)`

### `getMonitorInfo(hMonitor)`

### `monitorFromPoint(x, y, flags)`

### `monitorFromRect(left, top, right, bottom, flags)`

### `extendFrameIntoClientArea(window, margins)`

### `enableGlassSheet(window)`

### `onNcHitTest(window, handler)`

### `customChromeHitTest(mx, my, width, height, borderSize, captionHeight)`

### `setWindowRgn(window, hRgn, redraw)`

### `createRoundRectRgn(left, top, right, bottom, rx, ry)`

- `HTCLIENT` — constant
- `HTCAPTION` — constant
- `HTMINBUTTON` — constant
- `HTMAXBUTTON` — constant
- `HTCLOSE` — constant
### `showToastNotification(title, message, options)`

### `showToastWithFallback(title, message, options)`

### `enableMica(window)`

### `enableAcrylic(window)`

### `enableTabbedMica(window)`

### `disableBackdrop(window)`

### `setDwmDarkMode(window, dark)`

### `setDwmAttribute(window, attribute, value)`

### `beginDrag(window)`

### `updateDrag(window)`

### `endDrag(window)`

### `setDesignResolution(window, logicalWidth, logicalHeight, options)`

### `clearDesignResolution(window)`

### `hasDesignResolution?(window)`

### `designWidth(window)`

### `designHeight(window)`

### `physicalWidth(window)`

### `physicalHeight(window)`

### `resolutionScaleX(window)`

### `resolutionScaleY(window)`

### `resolutionOffsetX(window)`

### `resolutionOffsetY(window)`

### `physicalToLogical(window, px, py)`

### `logicalToPhysical(window, lx, ly)`

### `createCanvas(window, id, options)`

### `initWebGL(window, contextName, attrs)`

### `webglCreateShader(window, shaderType, source)`

### `webglCreateProgram(window, vertexShader, fragmentShader)`

### `webglUseProgram(window, program)`

### `webglClearColor(window, r, g, b, a)`

### `webglViewport(window, x, y, width, height)`

### `webglClear(window, mask)`

### `webglDrawArrays(window, mode, first, count)`

### `webglFlush(window)`

### `setTitle(window, title)`

### `setIcon(window, iconSpec)`

### `beginFrame(window)`

### `endFrame(window)`

### `drawText(window, x, y, text, color)`

### `textWidth(text, window)`

### `setFont(window, fontSpec)`

### `clearFont(window)`

### `fillRect(window, x, y, width, height, color, borderColor)`

### `pushMask(window, x, y, w, h)`

### `popMask(window)`

### `degToRad(deg)`

### `Vec3(x, y, z)`

### `_transformPoint(v, transform)`

### `_projectPoint(window, p, camera)`

### `_transformVertices(vertices, transform, i, out)`

### `drawLine(window, x1, y1, x2, y2, color)`

### `drawLinesBatch(window, segs, color)`

- `_graphCtxCached` · `{5 entries}`
### `graphRange(values, options)`

### `graphMapX(index, count, x, width)`

### `graphMapY(value, y, height, range)`

### `drawGraphAxes(window, x, y, width, height, options)`

### `drawLineGraph(window, x, y, width, height, values, options)`

### `drawBarGraph(window, x, y, width, height, values, options)`

### `drawSparkline(window, x, y, width, height, values, options)`

- `MSG_OFF_TYPE` — constant
- `MSG_OFF_WPARAM` — constant
- `MSG_OFF_LPARAM` — constant
- `FORM_WM_MOUSEMOVE` — constant
- `FORM_WM_LBUTTONDOWN` — constant
- `FORM_WM_LBUTTONUP` — constant
- `FORM_WM_RBUTTONDOWN` — constant
- `FORM_WM_RBUTTONUP` — constant
- `FORM_WM_MBUTTONDOWN` — constant
- `FORM_WM_MBUTTONUP` — constant
- `FORM_WM_MOUSEWHEEL` — constant
- `FORM_WM_LBUTTONDBLCLK` — constant
- `FORM_WM_KEYDOWN` — constant
- `FORM_WM_KEYUP` — constant
- `FORM_WM_CHAR` — constant
- `GUI_WM_SIZE` — constant
- `GUI_WM_PAINT` — constant
- `GUI_WM_ERASEBKGND` — constant
- `GUI_WM_SIZING` — constant
- `GUI_WM_WINDOWPOSCHANGED` — constant
- `GUI_WM_ENTERSIZEMOVE` — constant
- `GUI_WM_EXITSIZEMOVE` — constant
- `FORM_VK_BACK` — constant
- `FORM_VK_TAB` — constant
- `FORM_VK_RETURN` — constant
- `FORM_VK_SHIFT` — constant
- `FORM_VK_CONTROL` — constant
- `FORM_VK_ALT` — constant
- `FORM_VK_ESCAPE` — constant
- `MK_LBUTTON` — constant
- `VK_BACK` — constant
- `VK_TAB` — constant
- `VK_CLEAR` — constant
- `VK_RETURN` — constant
- `VK_SHIFT` — constant
- `VK_CONTROL` — constant
- `VK_ALT` — constant
- `VK_PAUSE` — constant
- `VK_CAPSLOCK` — constant
- `VK_ESCAPE` — constant
- `VK_SPACE` — constant
- `VK_PAGEUP` — constant
- `VK_PAGEDOWN` — constant
- `VK_END` — constant
- `VK_HOME` — constant
- `VK_LEFT` — constant
- `VK_UP` — constant
- `VK_RIGHT` — constant
- `VK_DOWN` — constant
- `VK_INSERT` — constant
- `VK_DELETE` — constant
- `VK_0` — constant
- `VK_1` — constant
- `VK_2` — constant
- `VK_3` — constant
- `VK_4` — constant
- `VK_5` — constant
- `VK_6` — constant
- `VK_7` — constant
- `VK_8` — constant
- `VK_9` — constant
- `VK_A` — constant
- `VK_B` — constant
- `VK_C` — constant
- `VK_D` — constant
- `VK_E` — constant
- `VK_F` — constant
- `VK_G` — constant
- `VK_H` — constant
- `VK_I` — constant
- `VK_J` — constant
- `VK_K` — constant
- `VK_L` — constant
- `VK_M` — constant
- `VK_N` — constant
- `VK_O` — constant
- `VK_P` — constant
- `VK_Q` — constant
- `VK_R` — constant
- `VK_S` — constant
- `VK_T` — constant
- `VK_U` — constant
- `VK_V` — constant
- `VK_W` — constant
- `VK_X` — constant
- `VK_Y` — constant
- `VK_Z` — constant
- `VK_NUMPAD0` — constant
- `VK_NUMPAD1` — constant
- `VK_NUMPAD2` — constant
- `VK_NUMPAD3` — constant
- `VK_NUMPAD4` — constant
- `VK_NUMPAD5` — constant
- `VK_NUMPAD6` — constant
- `VK_NUMPAD7` — constant
- `VK_NUMPAD8` — constant
- `VK_NUMPAD9` — constant
- `VK_MULTIPLY` — constant
- `VK_ADD` — constant
- `VK_SEPARATOR` — constant
- `VK_SUBTRACT` — constant
- `VK_DECIMAL` — constant
- `VK_DIVIDE` — constant
- `VK_F1` — constant
- `VK_F2` — constant
- `VK_F3` — constant
- `VK_F4` — constant
- `VK_F5` — constant
- `VK_F6` — constant
- `VK_F7` — constant
- `VK_F8` — constant
- `VK_F9` — constant
- `VK_F10` — constant
- `VK_F11` — constant
- `VK_F12` — constant
- `VK_NUMLOCK` — constant
- `VK_SCROLLLOCK` — constant
- `VK_OEM_SEMICOLON` — constant
- `VK_OEM_PLUS` — constant
- `VK_OEM_COMMA` — constant
- `VK_OEM_MINUS` — constant
- `VK_OEM_PERIOD` — constant
- `VK_OEM_SLASH` — constant
- `VK_OEM_TILDE` — constant
- `VK_OEM_LBRACKET` — constant
- `VK_OEM_BACKSLASH` — constant
- `VK_OEM_RBRACKET` — constant
- `VK_OEM_QUOTE` — constant
- `VK_LSHIFT` — constant
- `VK_RSHIFT` — constant
- `VK_LCONTROL` — constant
- `VK_RCONTROL` — constant
- `VK_LALT` — constant
- `VK_RALT` — constant
### `isLetterKey?(vk)`

### `isDigitKey?(vk)`

### `isNumpadKey?(vk)`

### `isFunctionKey?(vk)`

### `isArrowKey?(vk)`

### `isModifierKey?(vk)`

- `MK_RBUTTON` — constant
- `MK_SHIFT` — constant
- `MK_CONTROL` — constant
- `MK_MBUTTON` — constant
### `modShift?(wp)`

### `modCtrl?(wp)`

### `modAlt?(window)`

### `keyShiftDown?(window)`

### `keyCtrlDown?(window)`

### `keyAltDown?(window)`

### `formInRect?(mx, my, rx, ry, rw, rh)`

### `formMsgType(window)`

### `formMsgWParam(window)`

### `formMsgLParam(window)`

### `formLoWord(v)`

### `formHiWord(v)`

### `formEventContext(window)`

### `onMouseMove(window, handler)`

### `onLButtonDown(window, handler)`

### `onLButtonUp(window, handler)`

### `onRButtonDown(window, handler)`

### `onRButtonUp(window, handler)`

### `onMButtonDown(window, handler)`

### `onMButtonUp(window, handler)`

### `onMouseWheel(window, handler)`

### `onLButtonDblClk(window, handler)`

### `enableMouseTracking(window)`

### `onMouseHover(window, handler)`

### `onMouseLeave(window, handler)`

### `onKeyDown(window, handler)`

### `onKeyUp(window, handler)`

### `onChar(window, handler)`

### `onResize(window, handler)`

### `onDpiChanged(window, handler)`

### `onImeStartComposition(window, handler)`

### `onImeEndComposition(window, handler)`

### `onImeComposition(window, handler)`

### `setImePosition(window, x, y)`

### `registerTouchWindow(window, flags)`

### `unregisterTouchWindow(window)`

### `onTouch(window, handler)`

### `enableTouchInput(window, options)`

- `TOUCHEVENTF_MOVE` — constant
- `TOUCHEVENTF_DOWN` — constant
- `TOUCHEVENTF_UP` — constant
- `TOUCHEVENTF_PRIMARY` — constant
- `TOUCHEVENTF_PEN` — constant
- `TOUCHEVENTF_PALM` — constant
- `TWF_FINETOUCH` — constant
- `TWF_WANTPALM` — constant
### `onPointerDown(window, handler)`

### `onPointerUp(window, handler)`

### `onPointerUpdate(window, handler)`

### `onPointerEnter(window, handler)`

### `onPointerLeave(window, handler)`

### `enablePointerInput(window)`

- `PT_POINTER` — constant
- `PT_TOUCH` — constant
- `PT_PEN` — constant
- `PT_MOUSE` — constant
- `PT_TOUCHPAD` — constant
- `POINTER_FLAG_INCONTACT` — constant
- `POINTER_FLAG_PRIMARY` — constant
- `POINTER_FLAG_DOWN` — constant
- `POINTER_FLAG_UPDATE` — constant
- `POINTER_FLAG_UP` — constant
- `PEN_FLAG_BARREL` — constant
- `PEN_FLAG_ERASER` — constant
- `PEN_FLAG_INVERTED` — constant
### `getGamepadState(playerIndex)`

### `setGamepadVibration(playerIndex, left, right)`

### `stopGamepadVibration(playerIndex)`

### `applyThumbDeadzone(state)`

### `applyDeadzone(value, deadzone)`

### `gamepadPollState()`

### `gamepadPoll(ps)`

### `gamepadButtonDown?(state, button)`

### `gamepadButtonPressed?(ps, idx, button)`

### `gamepadButtonReleased?(ps, idx, button)`

### `gamepadConnected?(ps, idx)`

### `gamepadJustConnected?(ps, idx)`

### `gamepadJustDisconnected?(ps, idx)`

### `gamepadDpadUp?(state)`

### `gamepadDpadDown?(state)`

### `gamepadDpadLeft?(state)`

### `gamepadDpadRight?(state)`

### `gamepadStart?(state)`

### `gamepadBack?(state)`

### `gamepadA?(state)`

### `gamepadB?(state)`

### `gamepadX?(state)`

### `gamepadY?(state)`

### `gamepadLeftShoulder?(state)`

### `gamepadRightShoulder?(state)`

- `XUSER_MAX_COUNT` — constant
- `XINPUT_GAMEPAD_DPAD_UP` — constant
- `XINPUT_GAMEPAD_DPAD_DOWN` — constant
- `XINPUT_GAMEPAD_DPAD_LEFT` — constant
- `XINPUT_GAMEPAD_DPAD_RIGHT` — constant
- `XINPUT_GAMEPAD_START` — constant
- `XINPUT_GAMEPAD_BACK` — constant
- `XINPUT_GAMEPAD_LEFT_THUMB` — constant
- `XINPUT_GAMEPAD_RIGHT_THUMB` — constant
- `XINPUT_GAMEPAD_LEFT_SHOULDER` — constant
- `XINPUT_GAMEPAD_RIGHT_SHOULDER` — constant
- `XINPUT_GAMEPAD_A` — constant
- `XINPUT_GAMEPAD_B` — constant
- `XINPUT_GAMEPAD_X` — constant
- `XINPUT_GAMEPAD_Y` — constant
### `registerRawInputDevice(window, usagePage, usage, flags)`

### `unregisterRawInputDevice(usagePage, usage)`

### `onRawInput(window, handler)`

### `enableRawMouse(window)`

### `enableRawKeyboard(window)`

### `enableRawGamepad(window)`

### `enableRawJoystick(window)`

- `HID_USAGE_PAGE_GENERIC` — constant
- `HID_USAGE_GENERIC_POINTER` — constant
- `HID_USAGE_GENERIC_MOUSE` — constant
- `HID_USAGE_GENERIC_JOYSTICK` — constant
- `HID_USAGE_GENERIC_GAMEPAD` — constant
- `HID_USAGE_GENERIC_KEYBOARD` — constant
- `RIDEV_INPUTSINK` — constant
- `RIDEV_NOLEGACY` — constant
- `RIDEV_DEVNOTIFY` — constant
- `RIM_TYPEMOUSE` — constant
- `RIM_TYPEKEYBOARD` — constant
- `RIM_TYPEHID` — constant
### `setProcessDpiAwarenessContext(context)`

### `getDpiForWindow(hwnd)`

### `getDpiForSystem()`

### `getDpiForMonitor(hMonitor, dpiType)`

### `monitorFromWindow(hwnd, flags)`

### `adjustWindowRectExForDpi(x, y, w, h, style, exStyle, dpi)`

### `enableNonClientDpiScaling(hwnd)`

### `dpiScale(value, dpi)`

### `dpiUnscale(value, dpi)`

- `DPI_AWARENESS_CONTEXT_UNAWARE` — constant
- `DPI_AWARENESS_CONTEXT_SYSTEM_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2` — constant
### `clipboardGetText()`

### `clipboardSetText(text)`

### `clipboardHasText()`

### `enableFileDrop(window)`

### `disableFileDrop(window)`

### `onFileDrop(window, handler)`

### `enableOleDrop(window)`

### `disableOleDrop(window)`

### `onOleDrop(window, handler)`

### `dragDropState()`

### `onDragOver(window, state, handler)`

- `DROPEFFECT_NONE` — constant
- `DROPEFFECT_COPY` — constant
- `DROPEFFECT_MOVE` — constant
- `DROPEFFECT_LINK` — constant
### `openFileDialog(options)`

### `saveFileDialog(options)`

### `chooseColor(options)`

### `chooseFont(options)`

### `pickFolder(options)`

### `showPrintDialog(options)`

### `startDoc(hDC, docName, outputFile)`

### `startPage(hDC)`

### `endPage(hDC)`

### `endDoc(hDC)`

### `abortDoc(hDC)`

### `deletePrintDC(hDC)`

### `printTextOut(hDC, x, y, text)`

### `printMoveTo(hDC, x, y)`

### `printLineTo(hDC, x, y)`

### `printRectangle(hDC, l, t, r, b)`

### `printEllipse(hDC, l, t, r, b)`

### `printSetFont(hDC, height, weight, italic, fontName)`

### `printSetTextColor(hDC, r, g, b)`

### `printSetBkMode(hDC, mode)`

### `printSetPen(hDC, style, width, r, g, b)`

### `printDeleteObject(hObj)`

### `getDeviceCaps(hDC, index)`

### `getPrinterPageSize(hDC)`

### `createPreviewDC(width, height)`

### `destroyPreviewDC(preview)`

### `printToFile(outputPath, docName, renderFn)`

### `printDocument(options, renderPageFn)`

- `PD_ALLPAGES` — constant
- `PD_SELECTION` — constant
- `PD_PAGENUMS` — constant
- `PD_COLLATE` — constant
- `PD_PRINTTOFILE` — constant
- `PD_CURRENTPAGE` — constant
- `DEVCAP_HORZRES` — constant
- `DEVCAP_VERTRES` — constant
- `DEVCAP_LOGPIXELSX` — constant
- `DEVCAP_LOGPIXELSY` — constant
### `createMenu()`

### `createPopupMenu()`

### `menuAppendItem(hMenu, id, label)`

### `menuAppendSeparator(hMenu)`

### `menuAppendSubmenu(hMenu, hSub, label)`

### `setMenuBar(window, hMenu)`

### `removeMenuBar(window)`

### `destroyMenu(hMenu)`

### `showPopupMenu(window, hMenu, x, y)`

### `onMenuCommand(window, handler)`

### `buildMenu(spec)`

### `createAcceleratorTable(entries)`

### `destroyAcceleratorTable(hAccel)`

### `installAccelerators(window, entries)`

### `getSystemMenu(window, reset)`

### `appendSystemMenuItem(window, id, label)`

### `appendSystemMenuSeparator(window)`

### `resetSystemMenu(window)`

### `onSysCommand(window, handler)`

### `isDarkMode?()`

### `isHighContrast?()`

### `accentColor()`

### `addTrayIcon(window, id, tooltip, hIcon)`

### `removeTrayIcon(window, id)`

### `updateTrayTooltip(window, id, tooltip)`

### `showBalloonTip(window, id, title, message, iconFlag)`

### `onTrayEvent(window, handler)`

### `minimizeToTray(window)`

### `restoreFromTray(window)`

### `formSetStatus(state, message, ok)`

### `formPopChar(s)`

### `formClamp(v, minVal, maxVal)`

### `formHitListIndex(mx, my, x, y, itemW, itemH, count)`

### `formHitRectKey(mx, my, rects)`

### `formToggleIfHit(value, mx, my, x, y, w, h)`

### `formSelectByHit(current, mx, my, rects)`

### `formSetByAssignments(state, assignments)`

### `formResetFlags(state, keys, value)`

### `formSetHoverFromRects(state, mx, my, rects)`

### `formToggleKeysByHit(state, mx, my, rects)`

### `formSetKeyByHit(state, targetKey, mx, my, rects)`

### `formTruncateText(s, maxChars)`

### `formIsPrintableChar?(code)`

### `formSelectionState()`

### `formSelSetCursor(sel, pos, shifting)`

### `formSelMoveCursor(sel, text, dir, shifting)`

### `formSelHome(sel, shifting)`

### `formSelEnd(sel, text, shifting)`

### `formSelAll(sel, text)`

### `formSelRange(sel)`

### `formSelHasSelection?(sel)`

### `formSelSelectedText(sel, text)`

### `formSelDeleteSelection(sel, text)`

### `formSelInsertAtCursor(sel, text, insert)`

### `formSelBackspace(sel, text)`

### `formSelClickPos(mx, fieldX, fieldPadding, text)`

### `formDrawFieldWithSel(window, x, y, w, h, text, placeholder, focused, sel, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor, selColor)`

### `formUndoState(maxHistory)`

### `formUndoPush(hist, text, sel)`

### `formUndo(hist, text, sel)`

### `formRedo(hist, text, sel)`

### `formAppendByFocus(state, focus, c, fieldSpecs, notesSpec)`

### `formBackspaceByFocus(state, focus, fieldKeys, notesSpec)`

### `formCopyByFocus(state, focus, fieldKeys)`

### `formPasteByFocus(state, focus, fieldSpecs)`

### `formCutByFocus(state, focus, fieldKeys)`

### `formNotesAppendChar(lines, c)`

### `formNotesBackspace(lines)`

### `formNotesNewLine(lines, maxLines)`

### `formSliderValue(mx, sliderX, sliderW, handleW, minVal, maxVal)`

### `formApplySliderDrag(state, dragKey, mx, sliderX, sliderW, handleW, minVal, maxVal, bindings)`

### `formNextField(current, order)`

### `formPrevField(current, order)`

### `formFocusState(tabOrder)`

### `formFocusNext(fs)`

### `formFocusPrev(fs)`

### `formFocusSet(fs, key)`

### `formFocusIs?(fs, key)`

### `formHandleTabKey(fs, shiftDown)`

### `formIsActivateKey?(vk)`

### `formHandleKeyNav(fs, vk, shiftDown, widgetHandlers, options)`

### `formCheckboxKeyToggle(checked, vk)`

### `formRadioKeySelect(current, vk, choices)`

### `formSliderKeyAdjust(value, vk, minVal, maxVal, step)`

### `formDropdownKeyNav(isOpen, selectedIdx, vk, itemCount)`

### `formTreeKeyNav(flatList, selectedIdx, vk)`

### `formTableKeyNav(selectedRow, vk, rowCount, pageSize)`

### `formMaskText(s)`

- `_formCtxCached` · `{4 entries}`
### `formDrawBorder(window, x, y, w, h, color)`

### `formDrawFocusRing(window, x, y, w, h, color, options)`

### `formDrawField(window, x, y, w, h, text, placeholder, focused, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawPasswordField(window, x, y, w, h, text, placeholder, focused, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawCheckbox(window, x, y, checked, label, fieldColor, borderColor, checkColor)`

### `formDrawRadio(window, x, y, selected, label, fieldColor, borderColor, accentColor)`

### `formDrawPrimaryButton(window, x, y, w, h, label, hover, press, baseColor, hoverColor, pressColor, bottomLineColor)`

### `formDrawSecondaryButton(window, x, y, w, h, label, hover, press, baseColor, hoverColor, pressColor, borderColor)`

### `formDrawSlider(window, x, y, w, value, maxVal, trackColor, fillColor, thumbColor, thumbBorderColor, options)`

### `formDrawLabeledPercentSlider(window, x, labelY, sliderY, w, label, value, maxVal, trackColor, fillColor, thumbColor, thumbBorderColor, percentRectColor)`

### `formDrawNotes(window, x, y, w, h, lines, focused, placeholder, fieldColor, fieldFocusColor, borderColor, borderFocusColor, cursorColor)`

### `formDrawStatusBanner(window, x, y, w, h, message, ok, fieldColor, borderColor, successColor, errorColor)`

### `formDrawProgressBar(window, x, y, w, value, maxVal, trackColor, fillColor, borderColor, options)`

### `formDrawTabStrip(window, x, y, tabs, activeIdx, bgColor, activeBgColor, borderColor, activeTextColor, inactiveTextColor, options)`

### `formHitTab(mx, my, x, y, tabs, options)`

### `formDrawDropdown(window, x, y, w, h, selectedLabel, open, fieldColor, fieldFocusColor, borderColor, arrowColor)`

### `formDrawDropdownList(window, x, y, w, items, hoverIdx, itemHeight, listBgColor, hoverBgColor, borderColor)`

### `formHitDropdownItem(mx, my, x, y, w, items, itemHeight)`

### `formTooltipState()`

### `formTooltipUpdate(state, mx, my, regionKey, text, delay)`

### `formTooltipHide(state)`

### `formDrawTooltip(window, state, bgColor, borderColor, textColor, options)`

### `formDrawSpinner(window, x, y, w, h, value, focused, fieldColor, fieldFocusColor, borderColor, arrowColor, options)`

### `formSpinnerHit(mx, my, x, y, w, h, options)`

### `formSpinnerAdjust(value, direction, options)`

### `formDrawScrollbar(window, x, y, w, h, scrollPos, contentSize, viewSize, trackColor, thumbColor, borderColor, options)`

### `formScrollbarHit(mx, my, x, y, w, h, scrollPos, contentSize, viewSize, options)`

### `formDrawTreeView(window, x, y, w, h, nodes, selectedIdx, scrollPos, bgColor, selectedBgColor, textColor, selectedTextColor, borderColor, options)`

### `formTreeHitRow(mx, my, x, y, w, flat, scrollPos, options)`

### `formTreeToggle(flat, idx)`

### `formTreeContentHeight(flat, options)`

### `formDatePickerState(year, month, day)`

### `formDatePickerPrevMonth(state)`

### `formDatePickerNextMonth(state)`

### `formDateLabel(state)`

### `formDrawDateField(window, x, y, w, h, state, fieldColor, fieldFocusColor, borderColor, arrowColor)`

### `formDrawDateCalendar(window, x, y, state, bgColor, headerBgColor, selectedBgColor, todayBorderColor, textColor, selectedTextColor, headerTextColor, borderColor, options)`

### `formDateCalendarHit(mx, my, x, y, state, options)`

### `formDrawTable(window, x, y, w, h, columns, rows, scrollPos, selectedRow, headerBgColor, rowBgColor, altRowBgColor, selectedBgColor, headerTextColor, textColor, selectedTextColor, borderColor, options)`

### `formTableHitRow(mx, my, x, y, w, columns, rows, scrollPos, options)`

### `formTableContentHeight(rows, options)`

### `formTableHitColumn(mx, x, columns)`

### `formFrameTimerState(maxSamples)`

### `formFrameTimerTick(state)`

### `formDrawFrameTimingOverlay(window, state, x, y, w, h, options)`

### `formRichTextState(lines)`

### `formRichTextSetLines(state, lines)`

### `formRichTextAppendLine(state, spans)`

### `formRichTextInsertSpan(state, lineIdx, spanIdx, span)`

### `formDrawRichText(window, x, y, w, h, state, bgColor, borderColor)`

### `formRichTextScroll(state, delta)`

### `formRichTextTotalHeight(state)`

### `richSpan(text)`

### `richBold(text)`

### `richItalic(text)`

### `richUnderline(text)`

### `richColored(text, r, g, b)`

### `richStyled(text, options)`

### `enableAccessibility(window)`

### `accTree(window)`

### `accNode(id, name, role, bounds, state, children)`

### `accRegister(window, id, name, role, bounds, state)`

### `accUnregister(window, id)`

### `accSetName(window, id, name)`

### `accSetState(window, id, state)`

### `accSetValue(window, id, value)`

### `accSetBounds(window, id, bounds)`

### `accSetDescription(window, id, desc)`

### `accSetDefaultAction(window, id, action)`

### `accFocus(window, id)`

### `accSelection(window, id)`

### `notifyAccEvent(window, event, childId)`

### `accRegisterButton(window, id, label, x, y, w, h)`

### `accRegisterCheckbox(window, id, label, x, y, w, h, checked?)`

### `accRegisterRadio(window, id, label, x, y, w, h, selected?)`

### `accRegisterTextField(window, id, label, x, y, w, h, value)`

### `accRegisterSlider(window, id, label, x, y, w, h, value)`

### `accRegisterProgressBar(window, id, label, x, y, w, h, value)`

### `accRegisterTab(window, id, label, x, y, w, h, selected?)`

### `accRegisterListItem(window, id, label, x, y, w, h, selected?)`

### `accRegisterLink(window, id, label, x, y, w, h)`

### `accRegisterStaticText(window, id, label, x, y, w, h)`

### `accRegisterGroup(window, id, label, x, y, w, h)`

### `accRegisterTreeItem(window, id, label, x, y, w, h, expanded?)`

### `accRegisterTable(window, id, label, x, y, w, h)`

### `accRegisterTableCell(window, id, label, x, y, w, h)`

### `accRegisterDropdown(window, id, label, x, y, w, h, expanded?)`

### `accAnnounce(window, id, text)`

### `accDump(window)`

### `accVerify(window)`

### `accRoleName(role)`

### `accStateName(state)`

- `ROLE_SYSTEM_PUSHBUTTON` — constant
- `ROLE_SYSTEM_CHECKBUTTON` — constant
- `ROLE_SYSTEM_RADIOBUTTON` — constant
- `ROLE_SYSTEM_TEXT` — constant
- `ROLE_SYSTEM_STATICTEXT` — constant
- `ROLE_SYSTEM_SLIDER` — constant
- `ROLE_SYSTEM_PROGRESSBAR` — constant
- `ROLE_SYSTEM_PAGETAB` — constant
- `ROLE_SYSTEM_LISTITEM` — constant
- `ROLE_SYSTEM_LINK` — constant
- `ROLE_SYSTEM_OUTLINEITEM` — constant
- `ROLE_SYSTEM_TABLE` — constant
- `ROLE_SYSTEM_CELL` — constant
- `ROLE_SYSTEM_COMBOBOX` — constant
- `ROLE_SYSTEM_APPLICATION` — constant
- `ROLE_SYSTEM_CLIENT` — constant
- `ROLE_SYSTEM_GROUPING` — constant
- `STATE_SYSTEM_NORMAL` — constant
- `STATE_SYSTEM_FOCUSED` — constant
- `STATE_SYSTEM_SELECTED` — constant
- `STATE_SYSTEM_CHECKED` — constant
- `STATE_SYSTEM_PRESSED` — constant
- `STATE_SYSTEM_EXPANDED` — constant
- `STATE_SYSTEM_COLLAPSED` — constant
- `STATE_SYSTEM_READONLY` — constant
- `STATE_SYSTEM_FOCUSABLE` — constant
- `STATE_SYSTEM_UNAVAILABLE` — constant
- `EVENT_OBJECT_FOCUS` — constant
- `EVENT_OBJECT_STATECHANGE` — constant
- `EVENT_OBJECT_VALUECHANGE` — constant
- `EVENT_OBJECT_NAMECHANGE` — constant
- `EVENT_OBJECT_LIVEREGIONCHANGED` — constant
### `Vec2(x, y)`

### `Vec4(x, y, w, h)`

### `Rect2(x, y, width, height)`

### `vec2Add(a, b)`

### `vec2Sub(a, b)`

### `vec2Scale(v, s)`

### `vec2Dot(a, b)`

### `vec2Len(v)`

### `vec2Normalize(v)`

### `rectTranslate(rect, dx, dy)`

### `rectContains(rect, point)`

### `rectIntersects(a, b)`

### `Transform2D(options)`

### `applyTransform2D(point, transform)`

### `Camera2D(options)`

### `worldToScreen2D(point, camera, window)`

### `screenToWorld2D(point, camera, window)`

- `_depsLine` · `{1 entries}`
- `_depsLineBatch` · `{2 entries}`
- `_depsLineRect` · `{2 entries}`
- `_depsLineEllipse` · `{2 entries}`
- `_depsLinePoly` · `{2 entries}`
- `_depsLineBatchPoly` · `{3 entries}`
- `_depsLineRectRounded` · `{3 entries}`
- `_depsLineRectEllipse` · `{3 entries}`
### `drawRect2D(window, x, y, width, height, color, filled, borderColor)`

### `drawCircle2D(window, cx, cy, radius, color, filled, borderColor)`

### `drawPolyline2D(window, points, color, closed)`

### `drawPolygon2D(window, points, color, filled, borderColor)`

### `drawGrid2D(window, spacing, color, originX, originY)`

### `drawEllipse2D(window, cx, cy, rx, ry, color, filled, borderColor)`

### `drawArc2D(window, cx, cy, radius, startAngle, endAngle, color)`

### `drawTriangle2D(window, x1, y1, x2, y2, x3, y3, color, filled, borderColor)`

### `drawRoundedRect2D(window, x, y, width, height, radius, color, filled, borderColor)`

### `drawStar2D(window, cx, cy, outerR, innerR, points, color, filled, borderColor)`

### `drawBezier2D(window, points, color, steps)`

### `drawRing2D(window, cx, cy, outerR, innerR, color)`

### `drawCross2D(window, cx, cy, size, thickness, color, filled)`

### `drawDiamond2D(window, cx, cy, width, height, color, filled)`

### `drawArrow2D(window, x1, y1, x2, y2, color, headSize)`

### `drawCapsule2D(window, cx, cy, width, height, color, filled)`

### `drawSector2D(window, cx, cy, radius, startAngle, endAngle, color, filled)`

### `drawRegularPolygon2D(window, cx, cy, radius, sides, color, filled)`

### `drawSpiral2D(window, cx, cy, startRadius, growth, turns, color)`

### `drawThickLine2D(window, x1, y1, x2, y2, thickness, color)`

### `drawDashedLine2D(window, x1, y1, x2, y2, color, dashLen, gapLen)`

### `Element(bounds)`

### `drawLineAA(window, x1, y1, x2, y2, color, bgColor)`

### `drawCircleFilledAA(window, cx, cy, radius, color, bgColor)`

### `drawCircleOutlineAA(window, cx, cy, radius, color, bgColor)`

### `drawEllipseFilledAA(window, cx, cy, rx, ry, color, bgColor)`

### `drawRoundedRectFilledAA(window, x, y, w, h, radius, color, bgColor)`

### `drawTriangleFilledAA(window, p0, p1, p2, color, bgColor)`

- `_drawOpsDeps` · `?`
### `_initDrawOpsDeps()`

> returns `:object`

### `draw(window, op)`

### `drawBatch(window, ops)`

### `drawBatchParallel(window, ops, numWorkers)`

### `drawThreaded(window, op)`

### `drawBatchThreaded(window, ops)`

### `drawBatchParallelThreaded(window, ops, numWorkers)`

### `CubeMesh(size)`

> returns `:object`

### `drawTriangleFilled(window, p0, p1, p2, color)`

- `_depsMesh` · `{5 entries}`
- `_depsMeshWire` · `{3 entries}`
- `_depsRgb` · `{1 entries}`
### `drawMeshSolid(window, mesh, transform, camera, color, light, borderColor)`

### `drawMeshLit(window, mesh, transform, camera, color, scene, material, borderColor)`

### `DirectionalLight(options)`

### `PointLight(options)`

### `SpotLight(options)`

### `AmbientLight(options)`

### `Material(options)`

### `LightScene(options)`

### `addLight(scene, light)`

### `removeLight(scene, index)`

### `clearLights(scene)`

### `lightCount(scene)`

### `faceNormal(pa, pb, pc)`

### `faceCenter(pa, pb, pc)`

### `shadeFaceColor(baseColor, scene, material, pa3, pb3, pc3, camPos)`

### `shadeFaceIntensity(scene, material, pa3, pb3, pc3, camPos)`

### `prepareScene(scene, material, camPos)`

### `shadeFaceColorFast(baseColor, prep, pa3, pb3, pc3)`

### `Mesh(vertices, edges)`

### `GridMesh(size, step)`

### `AxesMesh(length)`

### `VoxelMesh(voxels, voxelSize)`

### `VoxelGrid(options)`

### `SphereMesh(radius, segments, rings)`

### `PyramidMesh(base, height)`

### `CylinderMesh(radius, height, segments)`

### `ConeMesh(radius, height, segments)`

### `TorusMesh(majorRadius, minorRadius, majorSegments, minorSegments)`

### `PlaneMesh(width, depth, subdivisionsW, subdivisionsD)`

### `HemisphereMesh(radius, segments, rings)`

### `WedgeMesh(width, height, depth)`

### `TubeMesh(outerRadius, innerRadius, height, segments)`

### `ArrowMesh(shaftRadius, shaftHeight, headRadius, headHeight, segments)`

### `PrismMesh(radius, height, sides)`

### `StairsMesh(steps, width, stepHeight, stepDepth)`

### `IcosphereMesh(radius)`

### `drawMeshWireframe(window, mesh, transform, camera, color)`

### `Renderer3D(window, options)`

### `poll(window)`

### `_frameIntervalNt(window)`

### `_frameMaxDtNt(window)`

### `_markUrgentFrameIfResizeDispatch(window, step)`

### `_maybeRunFrame(window, onFrame, force)`

### `_sleepUntilNextFrame(window)`

### `run(window, onEvent, onFrame)`

### `close(window)`

> returns `:int`

### `initThreading(window, options)`

### `destroyThreading(window)`

### `threadingEnabled?(window)`

### `threadCommandQueue(window)`

### `threadWorkerPool(window)`

### `threadScheduler(window)`

### `threadStateGuard(window)`

### `flushThreadedCommands(window)`

### `emitThreadSafe(window, event, payload)`

### `emitFromWorker(window, event, payload)`

### `DrawQueue()`

### `parallelFillRects(window, coords, computeFn, numWorkers)`

### `CommandQueue()`

### `WorkerPool(numWorkers)`

### `FrameFence(workerCount)`

### `FrameScheduler(pool, cmdQueue)`

### `StateGuard()`

### `AsyncLoader(cmdQueue)`

### `parallelTransformVertices(vertices, transformFn, numWorkers)`

- `shaderPI` — constant
- `shaderTAU` — constant
- `shaderHALF_PI` — constant
- `shaderE` — constant
- `shaderDEG2RAD` — constant
- `shaderRAD2DEG` — constant
- `shaderSQRT2` — constant
### `shaderFract(x)`

### `shaderMod(x, y)`

### `shaderSign(x)`

### `shaderAbs2(x)`

### `shaderClamp(x, lo, hi)`

### `shaderSaturate(x)`

### `shaderLerpFloat(a, b, t)`

### `shaderInverseLerp(a, b, x)`

### `shaderRemap(x, inLo, inHi, outLo, outHi)`

### `shaderStep(edge, x)`

### `shaderSmoothstep(edge0, edge1, x)`

### `shaderSmootherstep(edge0, edge1, x)`

### `shaderMin2(a, b)`

### `shaderMax2(a, b)`

### `shaderSqr(x)`

### `shaderSqrt(x)`

### `shaderLerp(a, b, t)`

### `shaderAtan2(y, x)`

### `shaderPingpong(t, length)`

### `shaderDegToRad(d)`

### `shaderRadToDeg(r)`

### `shaderEaseInQuad(t)`

### `shaderEaseOutQuad(t)`

### `shaderEaseInOutQuad(t)`

### `shaderEaseInCubic(t)`

### `shaderEaseOutCubic(t)`

### `shaderEaseInOutCubic(t)`

### `shaderEaseInSine(t)`

### `shaderEaseOutSine(t)`

### `shaderEaseInOutSine(t)`

### `shaderEaseInExpo(t)`

### `shaderEaseOutExpo(t)`

### `shaderEaseOutElastic(t)`

### `shaderEaseOutBounce(t)`

### `shaderVec2(x, y)`

### `shaderDot2(a, b)`

### `shaderLength2(v)`

### `shaderDistance2(a, b)`

### `shaderNormalize2(v)`

### `shaderRotate2(v, angle)`

### `shaderScale2(v, s)`

### `shaderAdd2(a, b)`

### `shaderSub2(a, b)`

### `shaderLerp2(a, b, t)`

### `shaderNegate2(v)`

### `shaderAbs2v(v)`

### `shaderMin2v(a, b)`

### `shaderMax2v(a, b)`

### `shaderFloor2(v)`

### `shaderFract2(v)`

### `shaderReflect2(v, n)`

### `shaderToPolar(v)`

### `shaderFromPolar(r, theta)`

### `shaderVec3(x, y, z)`

### `shaderAdd3(a, b)`

### `shaderSub3(a, b)`

### `shaderScale3(v, s)`

### `shaderDot3(a, b)`

### `shaderLength3(v)`

### `shaderDistance3(a, b)`

### `shaderNormalize3(v)`

### `shaderCross3(a, b)`

### `shaderLerp3(a, b, t)`

### `shaderNegate3(v)`

### `shaderReflect3(v, n)`

### `shaderPackRGB(r, g, b)`

### `shaderUnpackRGB(c)`

### `shaderColorR(c)`

### `shaderColorG(c)`

### `shaderColorB(c)`

### `shaderMix(a, b, t)`

### `shaderMix3(a, b, c, t)`

### `shaderBrighten(c, amount)`

### `shaderDarken(c, amount)`

### `shaderInvert(c)`

### `shaderGrayscale(c)`

### `shaderOverlay(fg, bg, alpha)`

### `shaderHsl2rgb(h, s, l)`

### `shaderRgb2hsl(c)`

### `shaderHsv2rgb(h, s, v)`

### `shaderRgb2hsv(c)`

### `shaderFloatStr(c)`

### `shaderCosinePalette(t, a, b, c, d)`

### `shaderContrast(c, amount)`

### `shaderSepia(c)`

### `shaderBlendMultiply(a, b)`

### `shaderBlendScreen(a, b)`

### `shaderBlendAdd(a, b)`

### `shaderHash(seed)`

### `shaderHash2(a, b)`

### `shaderHash3(a, b, c)`

### `shaderNoise2D(x, y)`

### `shaderFbm(x, y, octaves?)`

### `shaderNoiseGrid2DParallel(w, h, scaleFn, numWorkers)`

### `shaderFbmGrid2DParallel(w, h, scaleFn, octaves, numWorkers)`

### `shaderSdCircle(px, py, cx, cy, r)`

### `shaderSdBox(px, py, cx, cy, hw, hh)`

### `shaderSdLine(px, py, ax, ay, bx, by)`

### `shaderSdRoundedBox(px, py, cx, cy, hw, hh, r)`

### `shaderSdfFill(d, color)`

### `shaderSdfSmoothFill(d, color, bg, edge)`

### `shaderSdfStroke(d, thickness, color)`

### `shaderSdfGlow(d, color, intensity, radius)`

### `shaderSdUnion(d1, d2)`

### `shaderSdIntersect(d1, d2)`

### `shaderSdSubtract(d1, d2)`

### `shaderSdSmoothUnion(d1, d2, k)`

### `shaderSdSmoothIntersect(d1, d2, k)`

### `shaderSdSmoothSubtract(d1, d2, k)`

### `shaderSdAnnular(d, r)`

### `shaderSdRepeat2(px, py, cx, cy)`

### `shaderCheckerboard(x, y, size)`

### `shaderStripes(x, y, angle, width)`

### `shaderGrid(x, y, size, thickness)`

### `shaderDots(x, y, spacing, radius)`

### `shaderVoronoi(x, y, scale_)`

### `shaderGlslVersion(ver?)`

### `shaderGlslPrecision(prec?, type?)`

### `shaderGlslStdUniforms()`

### `shaderGlslUniform(type, name)`

### `shaderGlslUniforms(uniforms)`

### `shaderGlslIn(type, name)`

### `shaderGlslOut(type, name)`

### `shaderGlslQuadVertex()`

### `shaderGlslQuadVertexCompat()`

### `shaderGlslFragmentWrap(body, version?)`

### `shaderGlslMathLib()`

### `shaderHlslStdCBuffer()`

### `shaderHlslCBuffer(name, uniforms)`

### `shaderHlslQuadVertex()`

### `shaderHlslFragmentWrap(body)`

### `shaderHlslMathLib()`

### `shaderSubmitWebGL(window, fragSource, vertSource?)`

### `shaderDrawWebGL(window, clearR?, clearG?, clearB?)`

### `shaderRenderWebGL(window, fragSource)`

### `shaderCompileGLSL(source, stage?, outputPath?)`

### `shaderCompileHLSL(source, profile?, entry?, outputPath?)`

### `shaderCompileDXC(source, profile?, entry?, outputPath?, spirv?)`

### `shaderAssembleGLSL(opts)`

### `shaderAssembleHLSL(opts)`

### `Shader(fragment?, opts?)`

### `shaderElapsed(shader)`

### `shaderPause(shader)`

### `shaderResume(shader)`

### `shaderReset(shader)`

### `shaderSetUniform(shader, key, value)`

### `shaderGetUniform(shader, key)`

### `shaderSetResolution(shader, res)`

### `shaderBeginFrame(shader)`

### `shaderEndFrame(shader)`

### `shaderDt(shader)`

### `shaderIsRunning(shader)`

### `shaderFrameCount(shader)`

### `shaderRegisteredCount()`

### `shaderUnregisterShader(shader)`

### `shaderClearAll()`

### `shaderDestroyAll()`

### `shaderRender(window, shader, x, y, w, h)`

### `shaderRenderShader(window, shader, x, y, w, h)`

### `shaderRenderShaderLines(window, shader, x, y, w, h)`

### `shaderRenderGradientBands(window, gradientFn, x, y, w, h, time, bands?)`

### `shaderRenderGradient(window, gradientFn, x, y, w, h, time)`

### `shaderRenderHorizontalBands(window, gradientFn, x, y, w, h, time, bands?)`

### `shaderRenderColumns(window, columns, ox, oy, h)`

### `shaderUpdateColumns(columns, h, t, rate?)`

### `ShaderPass(shader, x, y, w, h)`

### `shaderComposePasses(window, passes)`

### `shaderCreateBuffer(w, h)`

### `shaderClearBuffer(buf, color?)`

### `shaderSetPixel(buf, x, y, color)`

### `shaderGetPixel(buf, x, y)`

### `shaderRenderBuffer(window, buf, ox, oy)`

### `shaderRenderShaderToBuffer(buf, shader)`

- `fontFW_THIN` — constant
- `fontFW_EXTRALIGHT` — constant
- `fontFW_LIGHT` — constant
- `fontFW_NORMAL` — constant
- `fontFW_MEDIUM` — constant
- `fontFW_SEMIBOLD` — constant
- `fontFW_BOLD` — constant
- `fontFW_EXTRABOLD` — constant
- `fontFW_HEAVY` — constant
### `defaultFontSpec()`

### `createFontFromSpec(spec)`

### `deleteFontFromSpec(fontResult)`

### `fontKey(spec)`

### `cachedFont(windowOrDisplay, spec)`

### `releaseCachedFonts(windowOrDisplay)`

### `measureTextEx(windowOrDisplay, spec, text)`

### `selectFontEx(args)`

### `fontGetTextMetrics(hdc)`

### `fontLineHeight(hdcOrFontStruct)`

### `buildXLFD(spec)`

### `videoFrameToBuffer(frame)`

### `videoBufferToFrame(buf)`

### `videoRenderFrame(window, frame, ox, oy)`

### `videoRenderFrameScaled(window, frame, ox, oy, scale)`

### `videoBufferGrayscale(buf)`

### `videoBufferInvert(buf)`

### `videoBufferThreshold(buf, t)`

### `videoBufferBlend(bufA, bufB, alpha)`

### `videoBufferDiff(bufA, bufB)`

### `videoCaptureBuffer(buf)`

### `videoFrameToBmpPixels(frame)`

### `createCanvas(window, options)`

### `beginCanvas(canvas)`

### `endCanvas(canvas)`

### `moveCanvas(canvas, x, y)`

### `resizeCanvas(canvas, w, h)`

### `setCanvasVisible(canvas, vis)`

### `setCanvasZIndex(canvas, z)`

### `setCanvasOpacity(canvas, alpha)`

### `setCanvasTransparentColor(canvas, color)`

### `destroyCanvas(canvas)`

### `destroyAllCanvases(window)`

### `isCanvas?(obj)`

### `canvases(window)`

### `canvasCount(window)`

### `canvasAt(window, px, py)`

### `canvasHitTest?(canvas, px, py)`

### `canvasToLocal(canvas, px, py)`

### `addJumpListTask(title, path, arguments, iconPath, iconIndex, description)`

### `clearJumpList()`

### `addJumpListRecentFile(filePath)`

### `registerFileAssociation(ext, progId, desc, cmd, icon)`

### `unregisterFileAssociation(ext, progId)`

### `refreshShellAssociations()`

### `addSearchFolder(folderPath, scope)`

### `searchFiles(query, maxResults)`

### `searchFilesWithProperty(query, property, maxResults)`

### `spatialAudioSource(id, x, y, volume)`

### `spatialAudioUpdate(source, lx, ly, vw, vh, maxDist)`

### `spatialApplyToSamples(samples, pan, gain)`

### `spatialMixSources(sources, lx, ly, vw, vh, maxDist, bufLen)`

### `setAppVolumeName(displayName)`

### `getSystemVolume()`

### `setSystemVolume(level)`

### `enableMediaTransportControls(options)`

### `setMediaPlaybackStatus(status)`

### `updateMediaInfo(title, artist, albumTitle)`

### `testMouseMove(x, y)`

### `testMouseClick(x, y)`

### `testMouseRightClick(x, y)`

### `testMouseDoubleClick(x, y)`

### `testMouseDrag(x1, y1, x2, y2)`

### `testMouseWheel(delta)`

### `testKeyDown(vk)`

### `testKeyUp(vk)`

### `testKeyPress(vk)`

### `testTypeText(text)`

### `testKeyCombo(keys)`

### `testSuite(name)`

### `testCase(suite, name, testFn)`

### `testRun(suite)`

### `testAssert(result, condition, message)`

### `testAssertEqual(result, actual, expected, message)`

### `testReport(suite)`

### `testGetWindowRect(hwnd)`

### `testGetWindowText(hwnd)`

### `testIsWindowVisible?(hwnd)`

### `testGetForegroundWindow()`

### `testFindWindow(cls, name)`

### `testScreenshot(hwnd)`

### `testFreeScreenshot(ss)`

- `TEST_VK_RETURN` — constant
- `TEST_VK_ESCAPE` — constant
- `TEST_VK_TAB` — constant
- `TEST_VK_BACK` — constant
- `TEST_VK_SPACE` — constant
- `TEST_VK_LEFT` — constant
- `TEST_VK_UP` — constant
- `TEST_VK_RIGHT` — constant
- `TEST_VK_DOWN` — constant
- `TEST_VK_SHIFT` — constant
- `TEST_VK_CONTROL` — constant
- `TEST_VK_MENU` — constant
- `TEST_VK_DELETE` — constant
- `TEST_VK_HOME` — constant
- `TEST_VK_END` — constant
### `getGPUAdapters()`

### `getGPUAdaptersParsed()`

### `getDXGIAdapters()`

### `getD3DFeatureLevel()`

### `getDisplayModes()`

### `getMonitorInfo()`

### `gpuCapabilityDump()`

### `gpuCapabilityDumpParallel()`

### `getGDIObjectCount()`

### `getUserObjectCount()`

### `getGDIObjectPeak()`

### `getUserObjectPeak()`

### `getHandleCount()`

### `getWorkingSetSize()`

### `getPeakWorkingSetSize()`

### `leakDetectorState()`

### `leakSnapshot(state)`

### `leakSnapshotParallel(state)`

### `leakCheck(state)`

### `leakReport(state)`

### `leakSetThresholds(state, thresholds)`

### `leakResetBaseline(state)`

### `leakTrend(state)`

- `LEAK_GR_GDIOBJECTS` — constant
- `LEAK_GR_USEROBJECTS` — constant
- `LEAK_GR_GDIOBJECTS_PEAK` — constant
- `LEAK_GR_USEROBJECTS_PEAK` — constant
## Module: `linux`

- `sys` · `import(...)`
## Module: `linux-constants`

- `LibC` · `[4 items]`
- `LibDL` · `[4 items]`
- `LibX11` · `[4 items]`
- `PROT_NONE` · `0`
- `PROT_READ` · `1`
- `PROT_WRITE` · `2`
- `PROT_EXEC` · `4`
- `MAP_SHARED` · `1`
- `MAP_PRIVATE` · `2`
- `MAP_FIXED` · `16`
- `MAP_ANONYMOUS` · `32`
- `O_RDONLY` · `0`
- `O_WRONLY` · `1`
- `O_RDWR` · `2`
- `O_CREAT` · `64`
- `O_TRUNC` · `512`
- `O_APPEND` · `1024`
- `SEEK_SET` · `0`
- `SEEK_CUR` · `1`
- `SEEK_END` · `2`
- `F_OK` · `0`
- `X_OK` · `1`
- `W_OK` · `2`
- `R_OK` · `4`
- `RTLD_LAZY` · `1`
- `RTLD_NOW` · `2`
- `RTLD_GLOBAL` · `256`
- `RTLD_LOCAL` · `0`
- `_SC_PAGESIZE` · `30`
- `KeyPressMask` · `1`
- `ButtonPressMask` · `4`
- `ExposureMask` · `32768`
- `StructureNotifyMask` · `131072`
- `KeyPress` · `2`
- `Expose` · `12`
- `DestroyNotify` · `17`
- `ClientMessage` · `33`
## Module: `linux-core`

- `sys` · `import(...)`
### `isLinux?()`

### `cstr(s)`

### `_readCString(ptr, maxLen)`

### `_platformError(apiName)`

> returns `:object`

### `_resolveFirst(libraries, symbol, i)`

> returns `:object`

### `_callResolved(resolved, args...)`

### `resolve(symbol)`

### `resolveIn(library, symbol)`

### `call(target, args...)`

### `libc(symbol, args...)`

### `libdl(symbol, args...)`

## Module: `linux-libc`

- `sys` · `import(...)`
### `currentProcessId()`

### `parentProcessId()`

### `pageSize()`

### `errno()`

### `strerror(errorCode)`

### `lastErrorMessage()`

### `getLastError()`

### `formatMessage(errorCode)`

### `currentProcess()`

### `moduleHandle(name)`

### `imageBase()`

### `getuid()`

### `geteuid()`

### `getgid()`

### `getegid()`

### `gethostname(bufferPtr, size)`

### `getcwd(bufferPtr, size)`

### `chdir(path)`

### `access(path, mode)`

### `openFile(path, flags, mode)`

### `closeFile(fd)`

### `closeHandle(handle)`

### `readFileDescriptor(fd, bufferPtr, count)`

### `writeFileDescriptor(fd, bufferPtr, count)`

### `seek(fd, offset, whence)`

### `unlink(path)`

### `mmap(addr, length, prot, flags, fd, offset)`

### `munmap(addr, length)`

### `mprotect(addr, length, prot)`

### `allocPages(size, prot)`

### `freePages(address, size)`

### `protectPages(address, size, prot)`

### `virtualAlloc(baseAddress, size, allocationType, protection)`

### `virtualFree(address, size, freeType)`

### `virtualProtect(address, size, newProtect, oldProtectOutPtr)`

### `_compatNotImplemented(apiName)`

> returns `:object`

### `openProcess(desiredAccess, inheritHandle, processId)`

### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`

### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`

### `virtualAllocEx(process, baseAddress, size, allocationType, protection)`

### `virtualFreeEx(process, address, size, freeType)`

### `virtualQuery(address, mbiBufferPtr, mbiSize)`

### `virtualQueryEx(process, address, mbiBufferPtr, mbiSize)`

## Module: `linux-loader`

- `sys` · `import(...)`
- `_loadedLibraries` · `{}`
### `_libraryKey(library)`

### `_normalizeHandleResult(result, apiName, library)`

> returns `:object`

### `dlopen(path, flags)`

### `dlsym(handle, symbol)`

### `dlclose(handle)`

### `loadLibrary(path)`

### `freeLibrary(handle)`

### `procAddress(module, symbol)`

### `_loadDllCandidate(candidates, i, flags)`

> returns `:object`

### `loadDll(library)`

### `resolveInLoaded(library, symbol)`

> returns `:object`

### `callIn(library, symbol, args...)`

### `x11(symbol, args...)`

## Module: `linux-windowing`

- `sys` · `import(...)`
### `_r1Positive?(res)`

> returns `:bool`

### `callOk?(res)`

> returns `:bool`

### `_zeros(n, acc)`

### `xEventSize()`

> returns `:int`

### `createXEventBuffer()`

### `xEventType(eventPtr)`

### `openDisplay(displayName)`

### `closeDisplay(display)`

### `defaultScreen(display)`

### `rootWindow(display, screen)`

### `blackPixel(display, screen)`

### `whitePixel(display, screen)`

### `createSimpleWindow(display, parent, x, y, width, height, borderWidth, border, background)`

### `destroyWindow(display, window)`

### `storeName(display, window, title)`

### `selectInput(display, window, eventMask)`

### `mapWindow(display, window)`

### `unmapWindow(display, window)`

### `moveWindow(display, window, x, y)`

### `resizeWindow(display, window, width, height)`

### `displayWidth(display, screen)`

### `displayHeight(display, screen)`

### `createGC(display, drawable, valueMask, values)`

### `freeGC(display, gc)`

### `setForeground(display, gc, color)`

### `drawLine(display, window, gc, x1, y1, x2, y2)`

### `fillRectangle(display, window, gc, x, y, width, height)`

### `drawString(display, window, gc, x, y, text)`

### `flush(display)`

### `pending(display)`

### `nextEvent(display, eventPtr)`

### `_openWindowFromRoot(display, screen, root, black, white, title, width, height)`

> returns `:object`

### `_openWindowFromScreen(display, screen, title, width, height)`

### `_openWindowFromDisplay(display, title, width, height)`

### `openDefaultWindow(title, width, height)`

### `closeWindow(state)`

> returns `:int`

### `pumpWindowEvent(display, eventPtr)`

> returns `:object`

### `runWindowLoop(display, eventPtr)`

> returns `:int`

## Module: `math`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

### `orient(x0, y0, x1, y1)`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

### `median(xs)`

### `stddev(xs)`

### `round(n, decimals)`

## Module: `math-base`

- `Pi` · `3.141592653589793`
- `E` · `2.718281828459045`
### `sign(n)`

> returns `:int`

### `abs(n)`

### `sqrt(n)`

## Module: `math-geo`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

> returns `:list`

### `orient(x0, y0, x1, y1)`

> returns `:int`

## Module: `math-stats`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

> returns `?`

### `median(xs)`

> returns `?`

### `stddev(xs)`

### `pbatchMean(datasets)`

### `pbatchStddev(datasets)`

### `round(n, decimals)`

## Module: `sort`

### `sort!(xs, pred)`

### `sort(xs, pred)`

### `_mergeSorted(a, b, pred)`

### `psort(xs, pred)`

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

## Module: `sys`

### `_isObject?(v)`

### `ok?(result)`

> returns `:bool`

### `error?(result)`

> returns `:bool`

### `resolve(library, symbol)`

> returns `:object`

### `call(target, args...)`

### `resolveAndCall(library, symbol, args...)`

### `valueOr(result, fallback)`

## Module: `thread`

### `spawn(fnToRun, args...)`

### `makeChannel(size)`

### `send(ch, value, callback)`

### `recv(ch, callback)`

### `close(_ch)`

> returns `?`

### `cs Mutex()`

> returns `:object`

### `cs Semaphore(n)`

> returns `:object`

### `cs WaitGroup()`

> returns `:object`

### `cs Future(fnToRun)`

> returns `:object`

### `cs Pool(numWorkers)`

> returns `:object`

### `parallel(fns)`

### `pmap(list, fnToRun)`

### `pmapConcurrent(list, fnToRun, maxConcurrent)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

## Module: `video`

### `clampByte(n)`

> returns `:int`

### `clampUnit(n)`

> returns `:int`

### `_min(a, b)`

### `_max(a, b)`

### `_pixelIndexRaw(width, channels, x, y, ch)`

### `pixelIndex(frame, x, y, ch)`

### `frame(width, height, pixels, channels)`

> returns `:object`

### `blank(width, height, channels, value)`

### `cloneFrame(f)`

> returns `:object`

### `getPixel(f, x, y)`

### `setPixel(f, x, y, pixel)`

> returns `:object`

### `mapPixels(f, mapper)`

> returns `:object`

### `_luma(pixel)`

### `toGrayscale(f)`

### `invert(f, maxValue)`

### `threshold(f, t)`

### `rgbToYuv(pixel)`

> returns `:list`

### `yuvToRgb(pixel)`

> returns `:list`

### `crop(f, x, y, width, height)`

### `resizeNearest(f, newWidth, newHeight)`

### `blend(a, b, alpha)`

### `frameDiff(a, b)`

### `mapFrames(frames, f)`

### `sampleFrame(frames, timeSeconds, fps)`

> returns `?`

### `pmapPixels(f, mapper, numWorkers)`

> returns `:object`

### `presizeNearest(f, newWidth, newHeight, numWorkers)`

### `pmapFrames(frames, f)`

## Module: `win-common`

- `windows` · `import(...)`
### `_ptrRead(address)`

### `_comCall(comObj, methodIndex, args...)`

### `_readOutPointer(outBuf)`

### `_releaseOk?(res)`

> returns `:bool`

## Module: `windows`

- `sys` · `import(...)`
## Module: `windows-constants`

- `Kernel32` · `'kernel32.dll'`
- `Ntdll` · `'ntdll.dll'`
- `Psapi` · `'psapi.dll'`
- `User32` · `'user32.dll'`
- `Gdi32` · `'gdi32.dll'`
- `Advapi32` · `'advapi32.dll'`
- `Shell32` · `'shell32.dll'`
- `Ole32` · `'ole32.dll'`
- `Ws2_32` · `'ws2_32.dll'`
- `Comctl32` · `'comctl32.dll'`
- `Wininet` · `'wininet.dll'`
- `OpenGL32` · `'opengl32.dll'`
- `Vulkan1` · `'vulkan-1.dll'`
- `D3d9` · `'d3d9.dll'`
- `D3d11` · `'d3d11.dll'`
- `Dxgi` · `'dxgi.dll'`
- `Ddraw` · `'ddraw.dll'`
- `Msvcrt` · `'msvcrt.dll'`
- `Ucrtbase` · `'ucrtbase.dll'`
- `Vcruntime140` · `'vcruntime140.dll'`
- `ActionCenter` · `'actioncenter.dll'`
- `Aclui` · `'aclui.dll'`
- `Acledit` · `'acledit.dll'`
- `Acppage` · `'acppage.dll'`
- `Acproxy` · `'acproxy.dll'`
- `Adprovider` · `'adprovider.dll'`
- `Aeinv` · `'aeinv.dll'`
- `Aepic` · `'aepic.dll'`
- `Amstream` · `'amstream.dll'`
- `Adsldp` · `'adsldp.dll'`
- `Adsnt` · `'adsnt.dll'`
- `Adtschema` · `'adtschema.dll'`
- `Adsldpc` · `'adsldpc.dll'`
- `Adsmsext` · `'adsmsext.dll'`
- `Adhsvc` · `'adhsvc.dll'`
- `Advapi32res` · `'advapi32res.dll'`
- `Advpack` · `'advpack.dll'`
- `Aeevts` · `'aeevts.dll'`
- `Apds` · `'apds.dll'`
- `Winhttp` · `'winhttp.dll'`
- `Urlmon` · `'urlmon.dll'`
- `Crypt32` · `'crypt32.dll'`
- `Bcrypt` · `'bcrypt.dll'`
- `Secur32` · `'secur32.dll'`
- `Comdlg32` · `'comdlg32.dll'`
- `Oleaut32` · `'oleaut32.dll'`
- `Imm32` · `'imm32.dll'`
- `Shlwapi` · `'shlwapi.dll'`
- `Shcore` · `'shcore.dll'`
- `UxTheme` · `'uxtheme.dll'`
- `Dwmapi` · `'dwmapi.dll'`
- `Version` · `'version.dll'`
- `Setupapi` · `'setupapi.dll'`
- `Netapi32` · `'netapi32.dll'`
- `Winmm` · `'winmm.dll'`
- `Avrt` · `'avrt.dll'`
- `Mmdevapi` · `'mmdevapi.dll'`
- `Dsound` · `'dsound.dll'`
- `Mfplat` · `'mfplat.dll'`
- `Mfreadwrite` · `'mfreadwrite.dll'`
- `Mfuuid` · `'mfuuid.dll'`
- `Taskschd` · `'taskschd.dll'`
- `Wevtapi` · `'wevtapi.dll'`
- `Wlanapi` · `'wlanapi.dll'`
- `Mpr` · `'mpr.dll'`
- `Spoolss` · `'spoolss.dll'`
- `Wtsapi32` · `'wtsapi32.dll'`
- `Rasapi32` · `'rasapi32.dll'`
- `Msi` · `'msi.dll'`
- `Wimgapi` · `'wimgapi.dll'`
- `Cabinet` · `'cabinet.dll'`
- `Apphelp` · `'apphelp.dll'`
- `Wer` · `'wer.dll'`
- `Faultrep` · `'faultrep.dll'`
- `Dbghelp` · `'dbghelp.dll'`
- `Dbgeng` · `'dbgeng.dll'`
- `Pdh` · `'pdh.dll'`
- `Iphlpapi` · `'iphlpapi.dll'`
- `Wscapi` · `'wscapi.dll'`
- `Sensapi` · `'sensapi.dll'`
- `Ncrypt` · `'ncrypt.dll'`
- `Cryptui` · `'cryptui.dll'`
- `Wintrust` · `'wintrust.dll'`
- `Samlib` · `'samlib.dll'`
- `Netshell` · `'netshell.dll'`
- `Fwpuclnt` · `'fwpuclnt.dll'`
- `Dnsapi` · `'dnsapi.dll'`
- `Nlaapi` · `'nlaapi.dll'`
- `Httpapi` · `'httpapi.dll'`
- `Rpcrt4` · `'rpcrt4.dll'`
- `Srpapi` · `'srpapi.dll'`
- `Sxs` · `'sxs.dll'`
- `Msvcirt` · `'msvcirt.dll'`
- `ApiSetPrefix` · `'api-ms-win-'`
- `D3dx9Prefix` · `'d3dx9_'`
- `MsvcpPrefix` · `'msvcp'`
- `VcruntimePrefix` · `'vcruntime'`
- `AtlPrefix` · `'atl'`
- `MfcPrefix` · `'mfc'`
- `VcompPrefix` · `'vcomp'`
## Module: `windows-core`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `makeWord(low, high)`

### `resolve(symbol)`

### `resolveIn(library, symbol)`

### `call(target, args...)`

### `kernel32(symbol, args...)`

### `ntdll(symbol, args...)`

### `ntNative(symbol, args...)`

### `psapi(symbol, args...)`

## Module: `windows-flags`

- `PROCESS_TERMINATE` · `1`
- `PROCESS_VM_READ` · `16`
- `PROCESS_VM_WRITE` · `32`
- `PROCESS_VM_OPERATION` · `8`
- `PROCESS_QUERY_INFORMATION` · `1024`
- `PROCESS_QUERY_LIMITED_INFORMATION` · `4096`
- `PROCESS_ALL_ACCESS` · `2035711`
- `MEM_COMMIT` · `4096`
- `MEM_RESERVE` · `8192`
- `MEM_DECOMMIT` · `16384`
- `MEM_RELEASE` · `32768`
- `PAGE_NOACCESS` · `1`
- `PAGE_READONLY` · `2`
- `PAGE_READWRITE` · `4`
- `PAGE_EXECUTE` · `16`
- `PAGE_EXECUTE_READ` · `32`
- `PAGE_EXECUTE_READWRITE` · `64`
- `FORMAT_MESSAGE_IGNORE_INSERTS` · `512`
- `FORMAT_MESSAGE_FROM_SYSTEM` · `4096`
- `ERROR_SUCCESS` · `0`
- `AF_INET` · `2`
- `SOCK_STREAM` · `1`
- `SOCK_DGRAM` · `2`
- `IPPROTO_TCP` · `6`
- `IPPROTO_UDP` · `17`
- `SOCKET_ERROR` — constant
- `INVALID_SOCKET` — constant
- `SD_RECEIVE` · `0`
- `SD_SEND` · `1`
- `SD_BOTH` · `2`
- `INTERNET_OPEN_TYPE_PRECONFIG` · `0`
- `INTERNET_OPEN_TYPE_DIRECT` · `1`
- `INTERNET_OPEN_TYPE_PROXY` · `3`
- `INTERNET_DEFAULT_HTTP_PORT` · `80`
- `INTERNET_DEFAULT_HTTPS_PORT` · `443`
- `INTERNET_SERVICE_HTTP` · `3`
- `HKEY_CLASSES_ROOT` · `2147483648`
- `HKEY_CURRENT_USER` · `2147483649`
- `HKEY_LOCAL_MACHINE` · `2147483650`
- `HKEY_USERS` · `2147483651`
- `HKEY_CURRENT_CONFIG` · `2147483653`
- `KEY_QUERY_VALUE` · `1`
- `KEY_SET_VALUE` · `2`
- `KEY_CREATE_SUB_KEY` · `4`
- `KEY_ENUMERATE_SUB_KEYS` · `8`
- `KEY_READ` · `131097`
- `KEY_WRITE` · `131078`
- `REG_SZ` · `1`
- `REG_DWORD` · `4`
- `REG_QWORD` · `11`
- `CS_VREDRAW` · `1`
- `CS_HREDRAW` · `2`
- `CS_DBLCLKS` · `8`
- `CS_OWNDC` · `32`
- `WS_OVERLAPPED` · `0`
- `WS_CAPTION` · `12582912`
- `WS_SYSMENU` · `524288`
- `WS_THICKFRAME` · `262144`
- `WS_MINIMIZEBOX` · `131072`
- `WS_MAXIMIZEBOX` · `65536`
- `WS_VISIBLE` · `268435456`
- `WS_CLIPSIBLINGS` · `67108864`
- `WS_CLIPCHILDREN` · `33554432`
- `WS_OVERLAPPEDWINDOW` · `13565952`
- `CW_USEDEFAULT` — constant
- `WS_POPUP` · `2147483648`
- `WS_EX_APPWINDOW` · `262144`
- `GWL_STYLE` — constant
- `GWL_EXSTYLE` — constant
- `SM_CXSCREEN` · `0`
- `SM_CYSCREEN` · `1`
- `HWND_TOP` · `0`
- `HWND_TOPMOST` — constant
- `HWND_NOTOPMOST` — constant
- `WM_CREATE` · `1`
- `WM_DESTROY` · `2`
- `WM_PAINT` · `15`
- `WM_CLOSE` · `16`
- `WM_QUIT` · `18`
- `WM_COMMAND` · `273`
- `SW_HIDE` · `0`
- `SW_MAXIMIZE` · `3`
- `SW_SHOW` · `5`
- `SW_RESTORE` · `9`
- `PM_NOREMOVE` · `0`
- `PM_REMOVE` · `1`
- `MB_OK` · `0`
- `MB_ICONERROR` · `16`
- `MB_ICONWARNING` · `48`
- `MB_ICONINFORMATION` · `64`
- `IDC_ARROW` · `32512`
- `IDI_APPLICATION` · `32512`
## Module: `windows-gdi`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
- `_szBuf` · `bits(...)`
### `isWindows?()`

### `wstr(s)`

- `_gdiProcCache` · `{}`
- `_userProcCache` · `{}`
### `_cachedGdi32(symbol, args...)`

### `_cachedUser32(symbol, args...)`

### `user32(symbol, args...)`

### `gdi32(symbol, args...)`

### `beginPaint(hwnd, paintStructPtr)`

### `endPaint(hwnd, paintStructPtr)`

### `getDC(hwnd)`

### `releaseDC(hwnd, hdc)`

### `getStockObject(objectIndex)`

### `selectObject(hdc, gdiObject)`

### `setBkMode(hdc, mode)`

### `setTextColor(hdc, colorRef)`

### `textOut(hdc, x, y, text)`

### `createFont(height, width, escapement, orientation, weight, italic, underline, strikeOut, charSet, outPrecision, clipPrecision, quality, pitchAndFamily, faceName)`

### `rectangle(hdc, left, top, right, bottom)`

### `ellipse(hdc, left, top, right, bottom)`

### `createSolidBrush(colorRef)`

### `getTextExtentPoint32(hdc, text)`

> returns `:object`

### `deleteObject(gdiObject)`

## Module: `windows-kernel`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `_utf16leToString(buf)`

### `wstr(s)`

### `cstr(s)`

### `kernel32(symbol, args...)`

### `statusOk?(res)`

> returns `:bool`

### `ptrSize()`

> returns `:int`

### `writePtr(address, value)`

### `ptrInt(ptrOrInt)`

### `callValueOrZero(res)`

### `_zeros(n, acc)`

### `getLastError()`

### `formatMessage(errorCode)`

### `lastErrorMessage()`

### `currentProcessId()`

### `currentProcess()`

### `moduleHandle(name)`

### `imageBase()`

### `loadLibrary(path)`

### `freeLibrary(module)`

### `procAddress(module, symbol)`

### `openProcess(desiredAccess, inheritHandle, processId)`

### `closeHandle(handle)`

### `virtualAlloc(baseAddress, size, allocationType, protection)`

### `virtualAllocEx(process, baseAddress, size, allocationType, protection)`

### `virtualFree(address, size, freeType)`

### `virtualFreeEx(process, address, size, freeType)`

### `virtualProtect(address, size, newProtect, oldProtectOutPtr)`

### `readProcessMemory(process, address, outBufferPtr, size, bytesReadOutPtr)`

### `writeProcessMemory(process, address, inBufferPtr, size, bytesWrittenOutPtr)`

### `virtualQuery(address, mbiBufferPtr, mbiSize)`

### `virtualQueryEx(process, address, mbiBufferPtr, mbiSize)`

## Module: `windows-loader`

- `sys` · `import(...)`
- `_loadedLibraries` · `{}`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
### `isWindows?()`

- `_loaderProcCache` · `{}`
### `_cachedCallIn(library, symbol, args...)`

### `_normalizeHandleResult(result, apiName, library)`

> returns `:object`

### `loadDll(library)`

> returns `:object`

### `resolveInLoaded(library, symbol)`

> returns `:object`

### `callIn(library, symbol, args...)`

### `user32(symbol, args...)`

### `gdi32(symbol, args...)`

### `advapi32(symbol, args...)`

### `shell32(symbol, args...)`

### `ole32(symbol, args...)`

### `ws2_32(symbol, args...)`

### `comctl32(symbol, args...)`

### `wininet(symbol, args...)`

### `opengl32(symbol, args...)`

### `vulkan1(symbol, args...)`

### `d3d9(symbol, args...)`

### `ddraw(symbol, args...)`

### `d3d11(symbol, args...)`

### `dxgi(symbol, args...)`

### `directDrawCreateEx(guidPtr, outPtr, iidPtr, outerUnknown)`

### `directDrawCreate(guidPtr, outPtr, outerUnknown)`

### `direct3DCreate9(sdkVersion)`

### `d3dx9Dll(suffix)`

### `d3dx9(suffix, symbol, args...)`

### `apiSetDll(contract)`

### `apiSet(contract, symbol, args...)`

### `msvcrt(symbol, args...)`

### `ucrtbase(symbol, args...)`

### `vcruntime140(symbol, args...)`

### `actionCenter(symbol, args...)`

### `aclui(symbol, args...)`

### `acledit(symbol, args...)`

### `acppage(symbol, args...)`

### `acproxy(symbol, args...)`

### `adprovider(symbol, args...)`

### `aeinv(symbol, args...)`

### `aepic(symbol, args...)`

### `amstream(symbol, args...)`

### `adsldp(symbol, args...)`

### `adsnt(symbol, args...)`

### `adtschema(symbol, args...)`

### `adsldpc(symbol, args...)`

### `adsmsext(symbol, args...)`

### `adhsvc(symbol, args...)`

### `advapi32res(symbol, args...)`

### `advpack(symbol, args...)`

### `aeevts(symbol, args...)`

### `apds(symbol, args...)`

### `winhttp(symbol, args...)`

### `urlmon(symbol, args...)`

### `crypt32(symbol, args...)`

### `bcrypt(symbol, args...)`

### `secur32(symbol, args...)`

### `comdlg32(symbol, args...)`

### `oleaut32(symbol, args...)`

### `imm32(symbol, args...)`

### `shlwapi(symbol, args...)`

### `shcore(symbol, args...)`

### `uxTheme(symbol, args...)`

### `dwmapi(symbol, args...)`

### `versionDll(symbol, args...)`

### `setupapi(symbol, args...)`

### `netapi32(symbol, args...)`

### `winmm(symbol, args...)`

### `avrt(symbol, args...)`

### `mmdevapi(symbol, args...)`

### `dsound(symbol, args...)`

### `mfplat(symbol, args...)`

### `mfreadwrite(symbol, args...)`

### `mfuuid(symbol, args...)`

### `taskschd(symbol, args...)`

### `wevtapi(symbol, args...)`

### `wlanapi(symbol, args...)`

### `mpr(symbol, args...)`

### `spoolss(symbol, args...)`

### `wtsapi32(symbol, args...)`

### `rasapi32(symbol, args...)`

### `msi(symbol, args...)`

### `wimgapi(symbol, args...)`

### `cabinet(symbol, args...)`

### `apphelp(symbol, args...)`

### `wer(symbol, args...)`

### `faultrep(symbol, args...)`

### `dbghelp(symbol, args...)`

### `dbgeng(symbol, args...)`

### `pdh(symbol, args...)`

### `iphlpapi(symbol, args...)`

### `wscapi(symbol, args...)`

### `sensapi(symbol, args...)`

### `ncrypt(symbol, args...)`

### `cryptui(symbol, args...)`

### `wintrust(symbol, args...)`

### `samlib(symbol, args...)`

### `netshell(symbol, args...)`

### `fwpuclnt(symbol, args...)`

### `dnsapi(symbol, args...)`

### `nlaapi(symbol, args...)`

### `httpapi(symbol, args...)`

### `rpcrt4(symbol, args...)`

### `srpapi(symbol, args...)`

### `sxs(symbol, args...)`

### `msvcirt(symbol, args...)`

### `_familyDll(prefix, suffix)`

### `msvcpDll(suffix)`

### `msvcpFamily(suffix, symbol, args...)`

### `vcruntimeDll(suffix)`

### `vcruntimeFamily(suffix, symbol, args...)`

### `atlDll(suffix)`

### `atlFamily(suffix, symbol, args...)`

### `mfcDll(suffix)`

### `mfcFamily(suffix, symbol, args...)`

### `vcompDll(suffix)`

### `vcompFamily(suffix, symbol, args...)`

## Module: `windows-net`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `wstr(s)`

### `cstr(s)`

### `_zeros(n, acc)`

### `callOk?(res)`

> returns `:bool`

### `callValueOrZero(res)`

### `ws2_32(symbol, args...)`

### `wininet(symbol, args...)`

### `wsLastError()`

### `wsaStartup(version, wsaDataPtr)`

### `wsaCleanup()`

### `socket(af, socketType, protocol)`

### `bindSocket(sock, sockaddrPtr, sockaddrLen)`

### `connectSocket(sock, sockaddrPtr, sockaddrLen)`

### `listenSocket(sock, backlog)`

### `acceptSocket(sock, addrOutPtr, addrLenInOutPtr)`

### `sendSocket(sock, bufferPtr, size, flags)`

### `recvSocket(sock, bufferPtr, size, flags)`

### `shutdownSocket(sock, how)`

### `closeSocket(sock)`

### `htons(value)`

### `htonl(value)`

### `inetAddr(ipv4)`

### `internetOpen(agent, accessType, proxy, proxyBypass, flags)`

### `internetConnect(hInternet, serverName, serverPort, username, password, service, flags, context)`

### `internetOpenUrl(hInternet, url, headers, headersLen, flags, context)`

### `internetReadFile(hFile, outBufferPtr, bytesToRead, bytesReadOutPtr)`

### `internetCloseHandle(hInternet)`

### `_bytesToString(raw)`

### `sockaddrIn(ipv4, port)`

> returns `:object`

### `_internetReadAll(hInternetFile, chunkBuf, bytesReadBuf, chunkSize, out)`

> returns `:object`

### `internetSimpleGet(url, agent, chunkSize)`

## Module: `windows-registry`

- `sys` · `import(...)`
### `_platformError(apiName)`

> returns `:object`

### `isWindows?()`

### `ptrSize()`

> returns `:int`

### `_zeros(n, acc)`

### `_statusOk?(res)`

> returns `:bool`

### `_ptrRead(address)`

### `_utf16leToString(buf)`

### `wstr(s)`

### `advapi32(symbol, args...)`

### `regCloseKey(hKey)`

### `regOpenKeyEx(rootKey, subKey, options, samDesired, outKeyPtr)`

### `regCreateKeyEx(rootKey, subKey, reserved, className, options, samDesired, securityAttributesPtr, outKeyPtr, dispositionOutPtr)`

### `regQueryValueEx(hKey, valueName, reserved, typeOutPtr, dataOutPtr, dataLenInOutPtr)`

### `regSetValueEx(hKey, valueName, reserved, valueType, dataPtr, dataLen)`

### `regDeleteValue(hKey, valueName)`

### `regDeleteTree(rootKey, subKey)`

### `regReadDword(rootKey, subKey, valueName)`

> returns `:object`

### `regWriteDword(rootKey, subKey, valueName, value)`

> returns `:object`

### `regReadString(rootKey, subKey, valueName)`

> returns `:object`

### `regWriteString(rootKey, subKey, valueName, value)`

> returns `:object`

## Module: `windows-windowing`

- `sys` · `import(...)`
### `_toI32(u)`

- `_cachedDefProcAddr` · `?`
- `_cachedCursor` · `?`
- `_cachedIcon` · `?`
- `_cachedImageBase` · `?`
- `_regClassTimings` · `{}`
- `_registeredClasses` · `{}`
- `DefaultClassName` · `'MagnoliaGUIWindowClass'`
### `_platformError(apiName)`

> returns `:object`

- `_isWindowsPlatform` — constant
### `isWindows?()`

### `wstr(s)`

### `cstr(s)`

- `_cachedPtrSize` — constant
### `ptrSize()`

### `writePtr(address, value)`

### `ptrInt(ptrOrInt)`

### `callValueOrZero(res)`

### `_zeros(n, acc)`

### `callOk?(res)`

> returns `:bool`

### `noMessage?(res)`

> returns `:bool`

- `_k32ProcCache` · `{}`
- `_u32ProcCache` · `{}`
- `_shcoreProcCache` · `{}`
- `_dwmapiProcCache` · `{}`
### `kernel32(symbol, args...)`

### `user32(symbol, args...)`

### `shcore(symbol, args...)`

### `dwmapi(symbol, args...)`

### `moduleHandle(name)`

### `imageBase()`

### `registerClassEx(wndClassExPtr)`

### `createWindowEx(exStyle, className, windowName, style, x, y, width, height, parent, menu, instance, param)`

### `defWindowProc(hwnd, msg, wParam, lParam)`

### `showWindow(hwnd, cmdShow)`

### `updateWindow(hwnd)`

### `getWindowLongPtr(hwnd, index)`

### `setWindowLongPtr(hwnd, index, value)`

### `getSystemMetrics(idx)`

### `destroyWindow(hwnd)`

### `postQuitMessage(exitCode)`

### `getMessage(msgPtr, hwnd, msgFilterMin, msgFilterMax)`

### `peekMessage(msgPtr, hwnd, msgFilterMin, msgFilterMax, removeMsg)`

### `translateMessage(msgPtr)`

### `dispatchMessage(msgPtr)`

### `isWindow(hwnd)`

### `waitMessage()`

### `windowAlive?(hwnd)`

> returns `:bool`

### `msgStructSize()`

> returns `:int`

### `createMsgBuffer()`

### `pumpWindowMessage(hwnd, msgPtr)`

> returns `:object`

### `loadCursor(instance, cursorId)`

### `loadIcon(instance, iconId)`

### `registerWindowClassEx(className, iconHandle, smallIconHandle, cursorHandle, classStyle)`

- `_initTimings` · `?`
### `getRegClassTimings()`

### `getInitTimings()`

### `registerDefaultWindowClass(className)`

### `createTopLevelWindow(className, title, width, height, style)`

### `runWindowLoop(hwnd)`

### `runWindowLoopPeek(hwnd, msgPtr)`

### `messageBox(hwnd, text, caption, msgType)`

### `setWindowText(hwnd, text)`

### `getCursorPos()`

> returns `:object`

### `getWindowRect(hwnd)`

> returns `:object`

### `getWindowPlacement(hwnd)`

> returns `:object`

- `DPI_AWARENESS_CONTEXT_UNAWARE` — constant
- `DPI_AWARENESS_CONTEXT_SYSTEM_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE` — constant
- `DPI_AWARENESS_CONTEXT_PER_MONITOR_AWARE_V2` — constant
- `MDT_EFFECTIVE_DPI` · `0`
- `MDT_ANGULAR_DPI` · `1`
- `MDT_RAW_DPI` · `2`
### `setProcessDpiAwarenessContext(context)`

> returns `:bool`

### `getDpiForWindow(hwnd)`

### `getDpiForSystem()`

### `getDpiForMonitor(hMonitor, dpiType)`

> returns `:object`

- `MONITOR_DEFAULTTONULL` · `0`
- `MONITOR_DEFAULTTOPRIMARY` · `1`
- `MONITOR_DEFAULTTONEAREST` · `2`
### `monitorFromWindow(hwnd, flags)`

### `adjustWindowRectExForDpi(x, y, w, h, style, exStyle, dpi)`

> returns `:object`

### `enableNonClientDpiScaling(hwnd)`

> returns `:bool`

### `dpiScale(value, dpi)`

### `dpiUnscale(value, dpi)`

- `_MONITORINFOEX_SIZE` · `104`
### `getMonitorInfo(hMonitor)`

> returns `:object`

### `monitorFromPoint(x, y, flags)`

### `monitorFromRect(left, top, right, bottom, flags)`

### `getSystemMetricsForDpi(index, dpi)`

## Module: `writes`

### `_b0(v)`

### `_b1(v)`

### `_b2(v)`

### `_b3(v)`

### `_b4(v)`

### `_b5(v)`

### `_b6(v)`

### `_b7(v)`

### `readU32(address)`

### `writeU32(address, value)`

### `readU64(address)`

### `writeU64(address, value)`

