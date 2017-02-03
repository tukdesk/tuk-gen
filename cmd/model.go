// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"log"

	"github.com/spf13/cobra"
	"github.com/tukdesk/base/sql/db"
	"github.com/tukdesk/tuk-gen/generator"
)

var opt generator.Option

// modelCmd represents the model command
var modelCmd = &cobra.Command{
	Use:   "model",
	Short: "generate model codes",
	Long:  `generate model codes from specified yaml files`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			cmd.Usage()
			log.Fatalln("input file required")
		}

		opt.Input = args[0]

		if err := generator.Generate(opt); err != nil {
			log.Fatalln(err)
		}
	},
}

func init() {
	modelCmd.Flags().StringVarP(&opt.OutputDir, "output", "o", "", "output dir")
	modelCmd.Flags().StringVarP(&opt.Package, "package", "p", "", "package name")
	modelCmd.Flags().StringVarP(&opt.DefaultEngine, "engine", "e", db.MYSQL, "default engine")
	RootCmd.AddCommand(modelCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// modelCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// modelCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

}
