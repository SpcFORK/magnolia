# GUI Form Library (gui-form)

## Overview

`gui-form` contains stateless form helpers and form-widget drawing functions.

Draw functions receive an injected context as their first argument:

- `ctx.drawLine`
- `ctx.fillRect`
- `ctx.drawText`

## Import

```oak
form := import('gui-form')
```

## State helpers

- `formSetStatus`
- `formPopChar`
- `formClamp`
- `formHitListIndex`
- `formHitRectKey`
- `formToggleIfHit`
- `formSelectByHit`
- `formSetByAssignments`
- `formResetFlags`
- `formSetHoverFromRects`
- `formToggleKeysByHit`
- `formSetKeyByHit`
- `formTruncateText`
- `formIsPrintableChar?`
- `formAppendByFocus`
- `formBackspaceByFocus`
- `formNotesAppendChar`
- `formNotesBackspace`
- `formNotesNewLine`
- `formSliderValue`
- `formApplySliderDrag`
- `formNextField`
- `formMaskText`

## Drawing helpers

- `formDrawBorder`
- `formDrawField`
- `formDrawPasswordField`
- `formDrawCheckbox`
- `formDrawRadio`
- `formDrawPrimaryButton`
- `formDrawSecondaryButton`
- `formDrawSlider`
- `formDrawLabeledPercentSlider`
- `formDrawNotes`
- `formDrawStatusBanner`

## Notes

- `GUI` re-exports window-first wrappers for all form draw helpers.
