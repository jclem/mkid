package cmd

import "github.com/spf13/cobra"

var rootCmd = &cobra.Command{
	Use:               "mkid",
	Short:             "mkid is a CLI for generating unique IDs",
	DisableAutoGenTag: true,
	PersistentPreRun: func(cmd *cobra.Command, _ []string) {
		// Ensures that usage isn't printed for errors such as non-zero exits.
		// SEE: https://github.com/spf13/cobra/issues/340#issuecomment-378726225
		cmd.SilenceUsage = true
	},
}
