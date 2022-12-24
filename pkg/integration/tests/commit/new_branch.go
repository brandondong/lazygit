package commit

import (
	"github.com/jesseduffield/lazygit/pkg/config"
	. "github.com/jesseduffield/lazygit/pkg/integration/components"
)

var NewBranch = NewIntegrationTest(NewIntegrationTestArgs{
	Description:  "Creating a new branch from a commit",
	ExtraCmdArgs: "",
	Skip:         false,
	SetupConfig:  func(config *config.AppConfig) {},
	SetupRepo: func(shell *Shell) {
		shell.
			EmptyCommit("commit 1").
			EmptyCommit("commit 2").
			EmptyCommit("commit 3")
	},
	Run: func(shell *Shell, input *Input, assert *Assert, keys config.KeybindingConfig) {
		assert.CommitCount(3)

		input.SwitchToCommitsWindow()
		assert.CurrentViewName("commits")
		input.NextItem()

		input.Press(keys.Universal.New)

		branchName := "my-branch-name"
		input.Prompt(Contains("New Branch Name"), branchName)

		assert.CommitCount(2)
		assert.HeadCommitMessage(Contains("commit 2"))
		assert.CurrentBranchName(branchName)
	},
})
