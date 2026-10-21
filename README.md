# imfmt

A small CLI for converting images between formats.

## Installation

Requires Go 1.25 or newer.

```bash
go install github.com/voittaneena/imfmt/cmd/imfmt@latest
```

## Usage

```bash
# Output goes next to the source file
imfmt convert stub.png --to jpg

# Custom quality and output path
imfmt convert stub.png -t jpg -q 90 -o out.jpg
```

Supported formats are listed in `imfmt convert --help`.

## Flags

| Flag | Description |
| --- | --- |
| `-t`, `--to` | Target format. Required. |
| `-q`, `--quality` | Quality, `0`–`100` (default `85`). Applies to JPEG; for PNG it only affects compression speed. |
| `-o`, `--out` | Output path. Default: next to the source, extension taken from `--to`. |

## License

[MIT](LICENSE)
