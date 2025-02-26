# Defaults

This package provides sub-packages for setting default settings. This is done through the use of import side-effects, so only import these if you plan to use their values. For example, adding `import _ "github.com/kettek/rebui/defaults/font"` to a package will ensure that a default font is available to any widgets that use text rendering.
