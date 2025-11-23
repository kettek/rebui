# Templates
Templates provide an alternative way to creating widgets by providing a set of nodes that are to be dynamically inserted where a Template widget is defined.

For example, if we have the following layout:

```yaml
- ID: my_template
  Type: Template
  Source: button_pair
  Width: 100%
  Height: 100%
```

And then the following `button_pair` file defined (note that a Template loader will have to be defined):

```yaml
- ID: button1
  Type: Button
  Text: one
  Height: 100%
  Width: 50%
- ID: button2
  Type: Button
  Text: two
  Height: 100%
  Width: 50%
  X: after button1
```

Then you will have two buttons placed where `my_template` is defined and using `my_template`'s dimensions and positions for button1 and button2's sizing and position.

Additionally, template children will have their IDs modified to use the template parent's ID for fully unique identification. In the above example, `button1` becomes `my_template__button1` and `button2` becomes `my_template__button2`.

Finally, nested templating (e.g., templates in templates) is supported, however there are no checks in place for cyclical loops, such as templating the same template within itself.
