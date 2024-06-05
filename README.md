# termpal

✨ Overview
---

`termpal` *(short for terminal palette)* is a dead simple CLI tool that shows
you how colors will look in your terminal.

![Usage](assets/termpal.png)

💾 Installation
---

**go**:

```sh
go install github.com/dhth/termpal@latest
```

⚡️ Usage
---

```bash
echo '#fb4934\n#83a598\n#d3869b' | termpal
# or
termpal -colors='#fb4934 #83a598 #d3869b'
```

```
Flags:
  -1    to print in a one column
  -bg
        show usage as a background color (default true)
  -c string
        space separated list of hex colors
  -cols int
        number of columns to show (default 8)
  -fg
        show usage as a foreground color
  -fgc string
        hex color to use for foreground (default "#282828")
  -txt string
        text to show in color
```
