# Changelog

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

  Build the palette once, where the rest of your setup happens, rather than at
  each call: `NewPalette` refuses a count below 1, and that refusal is worth
  having at startup instead of halfway through drawing a screen.

- `cases.yaml` carries the palette size in every slot case, and adds cases for
  palettes of 1, 5, 7, 8 and 256 colours. A port holding a twelve of its own
  anywhere answers the old cases correctly and these wrongly.

### Removed

- `Slots`, the package-level constant. Twelve was the first caller's palette,
  and the second one paints a different number.

## 0.1.0

### Added

- `Initials(name)` — the letters drawn in a square for a thing with no picture,
  and `Slot(id)` — which of the palette's `Slots` colours it is drawn on.
- `cases.yaml`, the contract every port answers: 92 cases across writing
  systems, with the rules stated in full at the top of the file.
