package cmd

import (
	"encoding/base64"
	"errors"
	"fmt"
	"math/big"
	"strings"

	"github.com/google/uuid"
	"github.com/mr-tron/base58"
	"github.com/spf13/cobra"
)

var uuidBase string
var uuidNoPad bool
var uuidCount int

var uuidCmd = &cobra.Command{
	Use:   "uuid <version>",
	Short: "Generates a UUID of the given version",
	Args:  cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flags().Changed("pad") && (uuidBase != "58" && uuidBase != "62") {
			return errors.New("padding is only supported for base-58 and base-62 encoding")
		}

		for range uuidCount {
			ustr, err := genUUID(args[0])
			if err != nil {
				return err
			}

			fmt.Fprintln(cmd.OutOrStdout(), ustr)
		}

		return nil
	},
}

func genUUID(vsn string) (string, error) {
	var u uuid.UUID
	switch vsn {
	case "4":
		var err error
		u, err = uuid.NewRandom()
		if err != nil {
			return "", fmt.Errorf("generate UUID: %w", err)
		}
	case "7":
		var err error
		u, err = uuid.NewV7()
		if err != nil {
			return "", fmt.Errorf("generate UUID: %w", err)
		}
	default:
		return "", fmt.Errorf("invalid UUID version: %s", vsn)
	}

	var ustr string
	switch uuidBase {
	case "58":
		ustr = base58.Encode(u[:])

		if !uuidNoPad && len(ustr) < 22 {
			ustr = padStr(ustr, 22, '1')
		}
	case "62":
		var i big.Int
		i.SetBytes(u[:])
		ustr = i.Text(62)

		if !uuidNoPad && len(ustr) < 22 {
			ustr = padStr(ustr, 22, '0')
		}
	case "64":
		ustr = base64.StdEncoding.EncodeToString(u[:])
	case "url64":
		ustr = base64.URLEncoding.EncodeToString(u[:])
	case "":
		ustr = u.String()
	default:
		return "", fmt.Errorf("invalid base-encoding: %s", uuidBase)
	}

	return ustr, nil
}

func padStr(s string, n int, pad rune) string {
	if len(s) >= n {
		return s
	}

	return strings.Repeat(string(pad), n-len(s)) + s
}

func init() {
	uuidCmd.Flags().StringVarP(&uuidBase, "base", "b", "", "base-encoding for the UUID (58, 62, 64, url64)")
	uuidCmd.Flags().BoolVarP(&uuidNoPad, "no-pad", "P", false, "do not pad base-58 or base-62 encoded UUIDs with zeroes")
	uuidCmd.Flags().IntVarP(&uuidCount, "count", "c", 1, "number of UUIDs to generate")
	rootCmd.AddCommand(uuidCmd)
}
