/*
Copyright © 2021 Tom Siewert

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// completionCmd represents the completion command
var completionCmd = &cobra.Command{
	Use:   "completion [bash|zsh|fish|powershell]",
	Short: "Generate completion script",
	Long: `To load completions:

Bash:

  $ source <(berichtsheft completion bash)

  # To load completions for each session, execute once:
  # Linux:
  $ berichtsheft completion bash > /etc/bash_completion.d/yourprogram
  # macOS:
  $ berichtsheft completion bash > /usr/local/etc/bash_completion.d/yourprogram

Zsh:

  # If shell completion is not already enabled in your environment,
  # you will need to enable it.  You can execute the following once:

  $ echo "autoload -U compinit; compinit" >> ~/.zshrc

  # To load completions for each session, execute once:
  $ berichtsheft completion zsh > "${fpath[1]}/_yourprogram"

  # You will need to start a new shell for this setup to take effect.

fish:

  $ berichtsheft completion fish | source

  # To load completions for each session, execute once:
  $ berichtsheft completion fish > ~/.config/fish/completions/yourprogram.fish

PowerShell:

  PS> berichtsheft completion powershell | Out-String | Invoke-Expression

  # To load completions for every new session, run:
  PS> berichtsheft completion powershell > yourprogram.ps1
  # and source this file from your PowerShell profile.
`,
	DisableFlagsInUseLine: true,
	ValidArgs:             []string{"bash", "zsh", "fish", "powershell"},
	Args:                  cobra.ExactValidArgs(1),
	SilenceUsage:          true,
	RunE: func(cmd *cobra.Command, args []string) error {
		var err error

		switch args[0] {
		case "bash":
			err = cmd.Root().GenBashCompletion(os.Stdout)
		case "zsh":
			err = cmd.Root().GenZshCompletion(os.Stdout)
		case "fish":
			err = cmd.Root().GenFishCompletion(os.Stdout, true)
		case "powershell":
			err = cmd.Root().GenPowerShellCompletionWithDesc(os.Stdout)
		default:
			err = fmt.Errorf("unsupported shell: %s", args)
		}
		return err
	},
}

func init() {
	RootCmd.AddCommand(completionCmd)
}
