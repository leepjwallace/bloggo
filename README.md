<p align="center">
  <img src="https://go.dev/blog/go-brand/Go-Logo/PNG/Go-Logo_LightBlue.png" alt="Bloggo Logo" width="200"/>
</p>


# Bloggo

[![GitHub release](https://img.shields.io/github/release/yourusername/bloggo.svg)](https://github.com/leepjwallace/bloggo/releases)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)

Bloggo is a lightweight command-line tool for creating and managing your notes. It's written in Go and uses Vim as the text editor. 

## Installation

To install Bloggo, simply run the following command:

```bash
curl -sSL https://raw.githubusercontent.com/yourusername/bloggo/main/install.sh | sh
```

This will download the latest binary, make it executable, and move it to /usr/local/bin.

## Usage

### Create a new note:
```bash
bloggo create my-note
```

### List all notes:
```bash
bloggo list
```

### Open an existing note:
```bash
bloggo open my-note
```

### Delete a note:
```bash
bloggo delete my-note
```

### Help information:
```bash
bloggo help
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

This project is licensed under the terms of the MIT license. See the LICENSE file for details.