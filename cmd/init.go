package cmd

import (
	"fmt"

	"github.com/flat35hd99/manesan/model"
	"github.com/flat35hd99/manesan/repository"
	"github.com/spf13/cobra"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// initCmd represents the init command
var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initilize Note",
	Long:  `Setting your manesan note.`,
	Run: func(cmd *cobra.Command, args []string) {
		db, err := gorm.Open(sqlite.Open("manesan.db"), &gorm.Config{})
		if err != nil {
			panic(err)
		}

		db.Migrator().CreateTable(model.Note{})
		db.Migrator().CreateTable(model.Experiment{})

		noteRepository := repository.NoteRepositoryImpl{DB: db}
		note := &model.Note{
			Name: "The first note",
			Experiments: []model.Experiment{
				{
					Name: "A",
				},
				{
					Name: "B",
				},
			},
		}
		err = noteRepository.Create(note)

		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("success!")
	},
}

func init() {
	rootCmd.AddCommand(initCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// initCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// initCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
