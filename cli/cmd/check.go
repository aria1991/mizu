package cmd

import (
	"github.com/creasty/defaults"
	"github.com/spf13/cobra"
	"github.com/up9inc/mizu/cli/config/configStructs"
	"github.com/up9inc/mizu/cli/telemetry"
	"github.com/up9inc/mizu/shared/logger"
)

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check the Mizu installation for potential problems",
	RunE: func(cmd *cobra.Command, args []string) error {
		go telemetry.ReportRun("check", nil)
		runMizuCheck()
		return nil
	},
}

func init() {
	rootCmd.AddCommand(checkCmd)

	defaultCheckConfig := configStructs.CheckConfig{}
	if err := defaults.Set(&defaultCheckConfig); err != nil {
		logger.Log.Debug(err)
	}

	checkCmd.Flags().Bool(configStructs.PreTapCheckName, defaultCheckConfig.PreTap, "Check pre-tap Mizu installation for potential problems")
}
