## mkid

mkid is a CLI for generating unique IDs

### Options

```
  -h, --help   help for mkid
```

### SEE ALSO

* [mkid completion](#mkid-completion)	 - Generate the autocompletion script for the specified shell
* [mkid docs](#mkid-docs)	 - Print documentation
* [mkid uuid](#mkid-uuid)	 - Generates a UUID of the given version

###### Auto generated by spf13/cobra on 4-Dec-2024
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

###### Auto generated by spf13/cobra on 4-Dec-2024
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

###### Auto generated by spf13/cobra on 4-Dec-2024
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

###### Auto generated by spf13/cobra on 4-Dec-2024
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

###### Auto generated by spf13/cobra on 4-Dec-2024
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

###### Auto generated by spf13/cobra on 4-Dec-2024
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

###### Auto generated by spf13/cobra on 4-Dec-2024
## mkid uuid

Generates a UUID of the given version

```
mkid uuid <version> [flags]
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

###### Auto generated by spf13/cobra on 4-Dec-2024
