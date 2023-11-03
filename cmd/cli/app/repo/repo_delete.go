//
// Copyright 2023 Stacklok, Inc.
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

package repo

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	github "github.com/stacklok/minder/internal/providers/github"
	"github.com/stacklok/minder/internal/util"
	pb "github.com/stacklok/minder/pkg/api/protobuf/go/minder/v1"
)

var repoDeleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete repository",
	Long:  `Repo delete is used to delete a repository within the minder control plane`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if err := viper.BindPFlags(cmd.Flags()); err != nil {
			fmt.Fprintf(os.Stderr, "error binding flags: %s", err)
		}
	},
	RunE: func(cmd *cobra.Command, args []string) error {

		provider := util.GetConfigValue(viper.GetViper(), "provider", "provider", cmd, "").(string)
		repoid := viper.GetString("repo-id")
		format := viper.GetString("output")
		name := util.GetConfigValue(viper.GetViper(), "name", "name", cmd, "").(string)

		// if name is set, repo-id cannot be set
		if name != "" && repoid != "" {
			return fmt.Errorf("cannot set both name and repo-id")
		}

		// either name or repoid needs to be set
		if name == "" && repoid == "" {
			return fmt.Errorf("either name or repo-id needs to be set")
		}

		// if name is set, provider needs to be set
		if name != "" && provider == "" {
			return fmt.Errorf("provider needs to be set if name is set")
		}

		switch format {
		case formatJSON:
		case formatYAML:
		case formatDefault:
		default:
			return fmt.Errorf("invalid output format: %s", format)
		}

		conn, err := util.GrpcForCommand(cmd, viper.GetViper())
		util.ExitNicelyOnError(err, "Error getting grpc connection")
		defer conn.Close()

		client := pb.NewRepositoryServiceClient(conn)
		ctx, cancel := util.GetAppContext()
		defer cancel()

		deletedRepoID := &pb.DeleteRepositoryByIdResponse{}
		deletedRepoName := &pb.DeleteRepositoryByNameResponse{}
		// delete repo by id
		if repoid != "" {
			resp, err := client.DeleteRepositoryById(ctx, &pb.DeleteRepositoryByIdRequest{
				RepositoryId: repoid,
			})
			util.ExitNicelyOnError(err, "Error deleting repo by id")
			deletedRepoID = resp
		} else {
			if provider != github.Github {
				return fmt.Errorf("only %s is supported at this time", github.Github)
			}

			// delete repo by name
			resp, err := client.DeleteRepositoryByName(ctx, &pb.DeleteRepositoryByNameRequest{Provider: provider, Name: name})
			util.ExitNicelyOnError(err, "Error deleting repo by name")
			deletedRepoName = resp
		}

		status := util.GetConfigValue(viper.GetViper(), "status", "status", cmd, false).(bool)
		if status {
			// TODO: implement this
		} else {
			var out string
			var err error
			// print result just in JSON or YAML
			if format == "" || format == formatJSON {
				if repoid != "" {
					out, err = util.GetJsonFromProto(deletedRepoID)
				} else {
					out, err = util.GetJsonFromProto(deletedRepoName)
				}
				util.ExitNicelyOnError(err, "Error getting json from proto")
				fmt.Println(out)
			} else {
				if repoid != "" {
					out, err = util.GetYamlFromProto(deletedRepoID)
				} else {
					out, err = util.GetYamlFromProto(deletedRepoName)
				}
				util.ExitNicelyOnError(err, "Error getting json from proto")
				fmt.Println(out)
			}
		}
		return nil
	},
}

func init() {
	RepoCmd.AddCommand(repoDeleteCmd)
	repoDeleteCmd.Flags().StringP("output", "f", "", "Output format (json or yaml)")
	repoDeleteCmd.Flags().StringP("provider", "p", "", "Name of the enrolled provider")
	repoDeleteCmd.Flags().StringP("name", "n", "", "Name of the repository (owner/name format)")
	repoDeleteCmd.Flags().StringP("repo-id", "r", "", "ID of the repo to delete")
	repoDeleteCmd.Flags().BoolP("status", "s", false, "Only return the status of the profiles associated to this repo")
}
