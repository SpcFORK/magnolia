# gui-theme — System Theme Detection

`import('gui-theme')` provides Windows system theme detection for dark mode, high contrast mode, and accent color via registry queries.

## Quick Start

```oak
theme := import('gui-theme')

if theme.isDarkMode?() -> {
    println('Dark mode is active')
    bg := [30, 30, 30]
} else {
    println('Light mode')
    bg := [255, 255, 255]
}

if theme.isHighContrast?() -> {
    println('High contrast mode — use system colors')
}

accent := theme.accentColor()
if accent != ? -> {
    println('Accent: rgb(' + string(accent.r) + ',' + string(accent.g) + ',' + string(accent.b) + ')')
}
```

## API Reference

### `isDarkMode?()`

Checks if dark mode is active by reading `AppsUseLightTheme` from the registry. Returns `true` if dark mode is enabled, `false` otherwise.

### `isHighContrast?()`

Checks if Windows high contrast mode is enabled via `SystemParametersInfo`. Returns `true`/`false`.

### `accentColor()`

Gets the user's accent color from the registry. Returns `{r, g, b}` or `?` if unavailable.

## Notes

- Theme detection reads from `HKEY_CURRENT_USER\Software\Microsoft\Windows\CurrentVersion\Themes\Personalize`.
- Accent color is read from `HKEY_CURRENT_USER\Software\Microsoft\Windows\DWM\AccentColor`.
- These are point-in-time queries; the module does not watch for theme changes. Re-query periodically or after `WM_SETTINGCHANGE` if you need live updates.
