## mkid

mkid is a CLI for generating unique IDs

### Options

```
  -h, --help   help for mkid
```

### SEE ALSO

* [mkid completion](#mkid-completion)	 - Generate the autocompletion script for the specified shell
* [mkid docs](#mkid-docs)	 - Print documentation
* [mkid ulid](#mkid-ulid)	 - Generates a ULID
* [mkid uuid](#mkid-uuid)	 - Generates a UUID of the given version

## mkid completion

Generate the autocompletion script for the specified shell

### Synopsis

Generate the autocompletion script for mkid for the specified shell.
See each sub-command's help for details on how to use the generated script.


### Options

```
  -h, --help   help for completion
```

### SEE ALSO

* [mkid](#mkid)	 - mkid is a CLI for generating unique IDs
* [mkid completion bash](#mkid-completion-bash)	 - Generate the autocompletion script for bash
* [mkid completion fish](#mkid-completion-fish)	 - Generate the autocompletion script for fish
* [mkid completion powershell](#mkid-completion-powershell)	 - Generate the autocompletion script for powershell
* [mkid completion zsh](#mkid-completion-zsh)	 - Generate the autocompletion script for zsh

## mkid completion bash

Generate the autocompletion script for bash

### Synopsis

Generate the autocompletion script for the bash shell.

This script depends on the 'bash-completion' package.
If it is not installed already, you can install it via your OS's package manager.

To load completions in your current shell session:

	source <(mkid completion bash)

To load completions for every new session, execute once:

#### Linux:

	mkid completion bash > /etc/bash_completion.d/mkid

#### macOS:

	mkid completion bash > $(brew --prefix)/etc/bash_completion.d/mkid

You will need to start a new shell for this setup to take effect.


```
mkid completion bash
```

### Options

```
  -h, --help              help for bash
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [mkid completion](#mkid-completion)	 - Generate the autocompletion script for the specified shell

## mkid completion fish

Generate the autocompletion script for fish

### Synopsis

Generate the autocompletion script for the fish shell.

To load completions in your current shell session:

	mkid completion fish | source

To load completions for every new session, execute once:

	mkid completion fish > ~/.config/fish/completions/mkid.fish

You will need to start a new shell for this setup to take effect.


```
mkid completion fish [flags]
```

### Options

```
  -h, --help              help for fish
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [mkid completion](#mkid-completion)	 - Generate the autocompletion script for the specified shell

## mkid completion powershell

Generate the autocompletion script for powershell

### Synopsis

Generate the autocompletion script for powershell.

To load completions in your current shell session:

	mkid completion powershell | Out-String | Invoke-Expression

To load completions for every new session, add the output of the above command
to your powershell profile.


```
mkid completion powershell [flags]
```

### Options

```
  -h, --help              help for powershell
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [mkid completion](#mkid-completion)	 - Generate the autocompletion script for the specified shell

## mkid completion zsh

Generate the autocompletion script for zsh

### Synopsis

Generate the autocompletion script for the zsh shell.

If shell completion is not already enabled in your environment you will need
to enable it.  You can execute the following once:

	echo "autoload -U compinit; compinit" >> ~/.zshrc

To load completions in your current shell session:

	source <(mkid completion zsh)

To load completions for every new session, execute once:

#### Linux:

	mkid completion zsh > "${fpath[1]}/_mkid"

#### macOS:

	mkid completion zsh > $(brew --prefix)/share/zsh/site-functions/_mkid

You will need to start a new shell for this setup to take effect.


```
mkid completion zsh [flags]
```

### Options

```
  -h, --help              help for zsh
      --no-descriptions   disable completion descriptions
```

### SEE ALSO

* [mkid completion](#mkid-completion)	 - Generate the autocompletion script for the specified shell

## mkid docs

Print documentation

```
mkid docs [flags]
```

### Options

```
  -f, --format string   Output format (default "markdown")
  -h, --help            help for docs
```

### SEE ALSO

* [mkid](#mkid)	 - mkid is a CLI for generating unique IDs

## mkid ulid

Generates a ULID

```
mkid ulid [flags]
```

### Examples

```
  # Generate a ULID
  mkid ulid
```

### Options

```
  -c, --count int   number of ULIDs to generate (default 1)
  -h, --help        help for ulid
```

### SEE ALSO

* [mkid](#mkid)	 - mkid is a CLI for generating unique IDs

## mkid uuid

Generates a UUID of the given version

### Synopsis

Generates a UUID of the given version.

The UUID is encoded in the specified base-encoding, if provided. For base-58 and
base-62, a standard length of 22 is assumed. Unless the --no-pad/-P flag is
provided, the UUID will be left-padded with the encoding's zero value.

```
mkid uuid <version> [flags]
```

### Examples

```
  # Generate a UUID v4
  mkid uuid 4

  # Generate a base62-encoded UUID v7
  mkid uuid 7 -b62

  # Generate a non-padded base58-encoded UUID v4
  mkid uuid 4 -b58 -P
  
  # Generate 100 URL-safe base64-encoded UUID v4s
  mkid uuid 4 -burl64 -c100
```

### Options

```
  -b, --base string   base-encoding for the UUID (58, 62, 64, url64)
  -c, --count int     number of UUIDs to generate (default 1)
  -h, --help          help for uuid
  -P, --no-pad        do not pad base-58 or base-62 encoded UUIDs with zero-values
```

### SEE ALSO

* [mkid](#mkid)	 - mkid is a CLI for generating unique IDs

