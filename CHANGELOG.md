# Changelog

## 0.2.1

### Changed

- `cases.yaml` now says which Unicode versions its cases hold for, and why they
  hold across all of them. Each port reads grapheme clusters from whatever its
  language offers — a dependency in Go, the runtime in Swift, a generated table
  in Kotlin — and those sit on different versions of the standard, so the cases
  are written not to lean on the rules that moved between them. No case
  changed, and no answer changed.

## 0.2.0

### Changed

- How many colour slots there are is now yours to say. `Slot` moved onto a
  palette you build with the number your own theme paints, and the library no
  longer names a number of its own.

  ```go
  // was
  slot := lettermark.Slot(id) // 0 … lettermark.Slots-1

  // now
  marks := lettermark.NewPalette(12) // however many colours your theme paints
  slot := marks.Slot(id)             // 0 … marks.Slots()-1
  ```

  Pass 12 and every colour stays where it was: the rule behind `Slot` is
  unchanged, so a thing drawn yesterday is drawn in the same colour today.

  Build the palette once, where the rest of your setup happens, rather than at
  each call. `NewPalette` panics on a count below 1 — there is no error to
  handle and nothing to fall back to, because a palette that paints nothing
  means your code and your theme disagree, and that is worth hearing at startup
  instead of halfway through drawing a screen.

- `cases.yaml`, the test corpus this package answers, now names a palette size
  in every slot case and asks about palettes of 1, 5, 7, 8 and 256 colours.
  This concerns you only if you answer that corpus yourself, in your own
  implementation of these rules: one with a twelve left anywhere inside it
  passes every case of 0.1.0 and fails the new ones.

### Removed

- `Slots` and `Slot` at package level. Both moved onto the palette, and the
  constant became a method, so `lettermark.Slots` is now `marks.Slots()` —
  parentheses included. A method call is not a constant expression, so a `const`
  or an array length that stood on the old `Slots` needs rewriting rather than
  reparenthesising. Twelve was the first caller's palette, and the second one
  paints a different number.

## 0.1.0

### Added

- `Initials(name)` — the letters drawn in a square for a thing with no picture,
  and `Slot(id)` — which of the palette's `Slots` colours it is drawn on.
- `cases.yaml`, the contract every port answers: 92 cases across writing
  systems, with the rules stated in full at the top of the file.
