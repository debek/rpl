# rpl - Replace String CLI Tool

## Overview

This CLI tool, developed in Go, is designed for efficiently replacing a phrase or a string of characters within the current directory and all its subdirectories. It's an invaluable utility for quick and bulk text modifications in multiple files.

Works good on Linux and MacOs systems.

## Features

- Replace a single word or a phrase in all files within the current directory and subdirectories.
- Supports string replacements involving spaces by using quotes.

## Installation

ou can install **rpl** using the Go tool:

`go install github.com/debek/rpl@latest`

After installation, add the Go binary path to your system's PATH to access the `rpl` command from anywhere:

`export PATH=$PATH:~/go/bin`

## Usage

To use the tool, navigate to the directory where you want to perform the replacements and execute one of the following commands:

### Replacing a Single Word

```
rpl [old_word] [new_word]
```

- `old_word`: The word you want to replace.
- `new_word`: The word that will replace the old word.

### Replacing a Phrase

```
rpl "[old_string one]" "[new_string two]"
```

- `old_string one`: The phrase you want to replace. Enclose in quotes.
- `new_string two`: The phrase that will replace the old phrase. Enclose in quotes.

## Contribution

If you're interested in contributing to the osmon project, please check our [contribution guidelines](https://github.com/debek/osmon/blob/main/CONTRIBUTE.md). All contributions, from bug reporting to new feature suggestions, are highly appreciated.

## License

This project is available under the [MIT License](https://github.com/debek/osmon/blob/main/LICENSE).