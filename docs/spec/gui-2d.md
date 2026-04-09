# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\gui-2d.oak`

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

