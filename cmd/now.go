package cmd

import (
	"go-gituser/internal/connectors/git"
	"go-gituser/internal/logger"
	"go-gituser/internal/models"
	"go-gituser/internal/services"
	"go-gituser/internal/storage"
	"os"

	"github.com/spf13/cobra"
)

var nowCmd = &cobra.Command{
	Use:   "now",
	Short: "Print the current git account",
	Long:  "Print the current git account you are using",
	Run: func(cmd *cobra.Command, args []string) {
		accountStorage := storage.NewAccountJSONStorage(storage.AccountsStorageFile)
		gitConnector := git.NewGitConnector()
		accountService := services.NewAccountService(accountStorage, gitConnector)

		savedAccounts, err := accountService.ReadSavedAccounts()
		if err != nil {
			logger.PrintErrorExecutingMode()
			os.Exit(1)
		}

		currGitAccount := accountService.ReadCurrentGitAccount()
		if currGitAccount.Username == "" || currGitAccount.Email == "" {
			logger.PrintNoActiveMode()
			return
		}

		if savedAccounts.Personal.Username == (currGitAccount.Username) && savedAccounts.Personal.Email == (currGitAccount.Email) {
			logger.ReadCurrentAccountData(currGitAccount, models.PersonalMode)
			return
		}

		if savedAccounts.School.Username == (currGitAccount.Username) && savedAccounts.School.Email == (currGitAccount.Email) {
			logger.ReadCurrentAccountData(currGitAccount, models.SchoolMode)
			return
		}

		if savedAccounts.Work.Username == (currGitAccount.Username) && savedAccounts.Work.Email == (currGitAccount.Email) {
			logger.ReadCurrentAccountData(currGitAccount, models.WorkMode)
			return
		}

		isAccountSaved, err := accountService.CheckSavedAccount(currGitAccount)
		if err != nil {
			logger.PrintErrorExecutingMode()
			return
		}

		if !isAccountSaved {
			logger.ReadUnsavedGitAccount(currGitAccount)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(nowCmd)
}
