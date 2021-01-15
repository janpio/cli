package branch

import (
	"context"
	"fmt"
	"os"

	"github.com/lensesio/tableprinter"
	"github.com/pkg/browser"
	"github.com/planetscale/cli/cmdutil"
	"github.com/planetscale/cli/config"
	"github.com/planetscale/cli/printer"
	"github.com/spf13/cobra"
)

func GetCmd(cfg *config.Config) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "get <source_name> <branch_name>",
		Short: "Get a specific branch of a database",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			if len(args) != 2 {
				return cmd.Usage()
			}

			source := args[0]
			branch := args[1]

			web, err := cmd.Flags().GetBool("web")
			if err != nil {
				return err
			}

			if web {
				fmt.Println("🌐  Redirecting you to your database branch in your web browser.")
				err := browser.OpenURL(fmt.Sprintf("%s/%s/%s/branches/%s", cmdutil.ApplicationURL, cfg.Organization, source, branch))
				if err != nil {
					return err
				}
				return nil
			}

			client, err := cfg.NewClientFromConfig()
			if err != nil {
				return err
			}

			b, err := client.DatabaseBranches.Get(ctx, cfg.Organization, source, branch)
			if err != nil {
				return err
			}

			tableprinter.Print(os.Stdout, printer.NewDatabaseBranchPrinter(b))

			return nil
		},
	}

	cmd.Flags().BoolP("web", "w", false, "Show a database branch in your web browser.")
	return cmd
}
