# my-ls
(Still under progress)
## Description

**my-ls** is a custom implementation of the **ls** command, which displays the files and directories of a specified directory. If no directory is provided, it defaults to displaying the contents of the current directory.

The goal of this project is to replicate the behavior of the original **ls** command while implementing a subset of its options.

## Features

The **my-ls** command supports the following flags:

- **`-l`** : Displays detailed information about each file, similar to `ls -l`.
- **`-R`** : Recursively lists all subdirectories and their contents.
- **`-a`** : Shows hidden files (files that start with `.`).
- **`-r`** : Reverses the order of the output.
- **`-t`** : Sorts files by modification time, with the newest files first.

You can combine multiple flags just like in the standard **ls** command.

## Usage

User must first download "my-ls.exe" file using the following command:

```sh
go build -o my-ls main.go
```

```sh
./my-ls [options] [directory]
```

### Examples

- List files in the current directory:
  ```sh
  ./my-ls
  ```

- List files with detailed information:
  ```sh
  ./my-ls -l
  ```

- Show hidden files:
  ```sh
  ./my-ls -a
  ```

- List files in reverse order:
  ```sh
  ./my-ls -r
  ```

- Recursively list all files and directories:
  ```sh
  ./my-ls -R
  ```

- Sort files by modification time:
  ```sh
  ./my-ls -t
  ```

- Combine multiple flags:
  ```sh
  ./my-ls -la
  ```

## Notes

- The `-l` flag output must be identical to the system `ls -l`.
- Other flag outputs can be customized but should maintain logical consistency.

## License
This project is open-source and free to use.

