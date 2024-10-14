package internal

import (
	"fmt"
	"sync"

	"github.com/AlecAivazis/survey/v2"
)

var databases = []string{"PostgreSQL", "MySQL", "SQLite", "MongoDB"}

type DatabaseCredentials struct {
	Username string
	Password string
	DBName   string `survey:"db"`
}

func HandleDatabaseSelection() {
	selectedDatabases := []string{}

	prompt := &survey.MultiSelect{
		Message: "Select the databases you want to use:",
		Options: databases,
	}

	err := survey.AskOne(prompt, &selectedDatabases)
	if err != nil {
		fmt.Println("Oops... Something occurred right here: ", err)
		return
	}

	if len(selectedDatabases) == 0 {
		fmt.Println("No databases selected.")
		return
	}

	fmt.Println("You selected the following databases:")
	for _, db := range selectedDatabases {
		fmt.Println("- ", db)
	}
	fmt.Print("\n")

	credentials := map[string]DatabaseCredentials{}

	var credentialsQuestion = []*survey.Question{
		{
			Name:     "username",
			Prompt:   &survey.Input{Message: "Enter the username:"},
			Validate: survey.Required,
		},
		{
			Name:     "password",
			Prompt:   &survey.Password{Message: "Enter the password:"},
			Validate: survey.Required,
		},
		{
			Name:     "db",
			Prompt:   &survey.Input{Message: "Enter the DB name:"},
			Validate: survey.Required,
		},
	}

	// Gather the credentials for each database
	for i, db := range selectedDatabases {
		fmt.Printf("%s\n", db)

		credential := DatabaseCredentials{}
		err := survey.Ask(credentialsQuestion, &credential)
		if err != nil {
			fmt.Println("Oopss... \nSomething happened: ", err.Error())
			return
		}

		// Workaround for printing the next DB name after the first question
		if i+1 < len(selectedDatabases) {
			fmt.Printf("\n%s\n", selectedDatabases[i+1])
		}

		credentials[db] = credential

	}
	fmt.Printf("\nStarting your backups...\n")

	var wg sync.WaitGroup

	// The backup itself
	for db, credential := range credentials {
		wg.Add(1)
		switch db {
		case "PostgreSQL":
			go BackupPostgres(credential, &wg)
			break
		case "MySQL":
			go BackupMySQL(credential, &wg)
			break
		case "MongoDB":
			go BackupMongoDB(credential, &wg)
			break
		case "SQLite":
			go BackupSQLite(credential, &wg)
			break
		}
	}
	wg.Wait()
	fmt.Print("Alrighty! Everything was backed up sucessfully!\n")

}
