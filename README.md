# TerMImagenator

or TMI which also can be think of as Too Many Images

## Demo

![vhs/demo.gif](./vhs/demo.gif)

## Why

I couldn't find an easy and quick way to filter (Docker) container images that
can be deleted.

Sometimes your dev workflow generates lots of images or tags rapidly that
clutter the image list and it would be nice to select and delete (or untag)
unwanted ones quickly.

None worked for me from
[this](https://stackoverflow.com/questions/32490229/how-can-i-delete-docker-images-by-tag-preferably-with-wildcarding).

So why not build a tool that scratches my back (and potentially yours too)!

## Try

### With Nix

```bash
nix run github:ryuheechul/termimagenator
```

### With Go

```bash
go run github.com/ryuheechul/termimagenator/cmd/tmi@latest
```

## Install

### With Go

```bash
go install github.com/ryuheechul/termimagenator/cmd/tmi@latest
```

### Download Prebuilt Binaries

https://github.com/ryuheechul/termimagenator/releases/latest

## Core Dependencies

- https://github.com/docker/docker
- https://github.com/charmbracelet/bubbletea
  - https://github.com/charmbracelet/bubbles
  - https://github.com/charmbracelet/lipgloss
- https://github.com/samber/lo
- https://github.com/samber/mo
