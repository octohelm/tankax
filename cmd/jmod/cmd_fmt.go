package main

import (
	"context"

	"github.com/jsonnetmod/jsonnetmod/pkg/jsonnet"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(
		cmdFmt(),
	)
}

func cmdFmt() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fmt",
		Short: "format jsonnet codes",
	}

	formatOpts := jsonnet.FormatOpts{}

	return setupRun(cmd, &formatOpts, func(ctx context.Context, args []string) error {
		baseDir := "./"
		if len(args) > 0 {
			baseDir = args[0]
		}

		files, err := mod.ListJsonnet(baseDir)
		if err != nil {
			return err
		}

		return jsonnet.FormatFiles(cmd.Context(), files, formatOpts)
	})
}
