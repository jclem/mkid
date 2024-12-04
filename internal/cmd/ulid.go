package cmd

import (
	"fmt"

	"github.com/oklog/ulid/v2"
	"github.com/spf13/cobra"
)

var ulidCount int

var ulidCmd = &cobra.Command{
	Use:   "ulid",
	Short: "Generates a ULID",
	Example: `  # Generate a ULID
  mkid ulid`,
	RunE: func(cmd *cobra.Command, _ []string) error {
		for range ulidCount {
			u, err := ulid.New(ulid.Now(), ulid.DefaultEntropy())
			if err != nil {
				return fmt.Errorf("generate ULID: %w", err)
			}

			fmt.Fprintln(cmd.OutOrStdout(), u.String())
		}

		return nil
	},
}

func init() {
	ulidCmd.Flags().IntVarP(&ulidCount, "count", "c", 1, "number of ULIDs to generate")
	rootCmd.AddCommand(ulidCmd)
}
