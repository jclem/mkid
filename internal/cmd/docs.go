package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/cobra/doc"
)

var docsCmd = &cobra.Command{
	Use:   "docs",
	Short: "Print documentation",
	RunE: func(cmd *cobra.Command, _ []string) error {
		formatFlag, err := cmd.Flags().GetString("format")
		if err != nil {
			return fmt.Errorf("get format: %w", err)
		}

		var format string

		switch formatFlag {
		case "markdown":
			format = formatFlag
		default:
			return fmt.Errorf("unsupported format: %s", formatFlag)
		}

		switch format {
		case "markdown":
			md, err := genMarkdown(rootCmd)
			if err != nil {
				return fmt.Errorf("generate markdown: %w", err)
			}

			fmt.Fprintf(cmd.OutOrStdout(), "%s", md)
		}

		return nil
	},
}

func init() { //nolint:gochecknoinits
	docsCmd.Flags().StringP("format", "f", "markdown", "Output format")
	rootCmd.AddCommand(docsCmd)
}

func genMarkdown(cmd *cobra.Command) (string, error) {
	b := strings.Builder{}

	var gen func(cmd *cobra.Command) error

	linkHandler := func(s string) string {
		s = strings.ReplaceAll(s, "_", "-")
		s = strings.TrimSuffix(s, ".md")
		s = "#" + s
		return s
	}

	gen = func(cmd *cobra.Command) error {
		if err := doc.GenMarkdownCustom(cmd, &b, linkHandler); err != nil {
			return fmt.Errorf("generate markdown: %w", err)
		}

		for _, c := range cmd.Commands() {
			if !c.IsAvailableCommand() || c.IsAdditionalHelpTopicCommand() {
				continue
			}

			if err := gen(c); err != nil {
				return err
			}
		}

		return nil
	}

	if err := gen(cmd); err != nil {
		return "", err
	}

	return b.String(), nil
}
