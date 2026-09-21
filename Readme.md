# lettermark

[![Go](https://github.com/botforge-pro/lettermark/actions/workflows/go.yml/badge.svg)](https://github.com/botforge-pro/lettermark/actions/workflows/go.yml)
[![Reference](https://pkg.go.dev/badge/github.com/botforge-pro/lettermark.svg)](https://pkg.go.dev/github.com/botforge-pro/lettermark)

The mark a thing gets when it has no picture of its own: the letters that go in
the square, and which of the palette's colour slots it is drawn on.

```go
letters := lettermark.Initials(name)
slot := lettermark.Slot(id) // 0 … lettermark.Slots-1
```

The palette is not here. A caller keeps its own — CSS custom properties, a
colour catalogue, a theme — with `lettermark.Slots` colours in it, and looks up
the slot this library returns. That keeps one thing in one place: the rule here,
the colours where the rest of the design lives.

`Initials` is given the name a reader sees. A caller holding markup strips it
first; this library does not know what markup its caller writes.

`cases.yaml` is the contract. Every port carries a copy of it and a test that
compares that copy with this repository byte for byte, so a case added here is
answered by all of them or fails loudly.

## Ports

- Go — this repository, and the one the others follow
- Swift — [lettermark-swift](https://github.com/botforge-pro/lettermark-swift)
- Kotlin — [lettermark-kotlin](https://github.com/botforge-pro/lettermark-kotlin)

## Lines of Code

<picture>
  <source media="(prefers-color-scheme: dark)" srcset=".github/loc-history-dark.svg">
  <source media="(prefers-color-scheme: light)" srcset=".github/loc-history-light.svg">
  <img alt="Lines of Code graph" src=".github/loc-history-light.svg">
</picture>
