# Widgets

Widgets in rebui are simple structs that can implement a variety of interfaces. These interfaces are defined in the elements package under the setters, getters, and receivers sub-packages.

## Setters

Setters provide a way for rebui to tell the widget to set a particular field to a particular value. For example, if `SetBorderColor(color.Color)` is defined on a widget, then that widget will receive either the current theme's default border color, or it will receive a border color parsed from a node in a layout.

In general, most widgets will want to simply embed the `widgets.Basic` widget since that implements all the baseline functionality related to setting X, Y, Width, and Height, as well as implementing the `Hit` function to determine if pointer events intersect with the widget.

## Getters

Getters provides a way for rebui to access properties on a widget. These getters are used to compare a widget's current field value with a layout-calculated one, and if they differ, to not call the associated setter. What this means is that if a widget has its X, Y, Width, or Height manually changed, then any further layout calculations will not automatically adjust those values to the newly calculated ones. If the aforementioned fields are not manually changed, then they will automatically be set to the latest calculations during a layout change.

## Receivers

Receivers provide a way for rebui to send Events to a widget. All the widget has to do is implement a receiver interface, such as `HandlePointerIn(evt EventPointerIn)`, and the widget will receive a pointer in event when a pointer enters the widget's hit area. To note, there are also Node-level "On" event handlers that can be assigned to a func. This provides a mechanism to assign function pointers directly to a node in the event having such functionality on the widget-level doesn't satisfy requirements. Of course, this sort of "On" event handling could be defined directly on a widget by adding `OnPointerIn(func(EventPointerIn))` methods to the widget, storing the passed in callback, and triggering it within `HandlePointerIn(EventPointerIn)`.
