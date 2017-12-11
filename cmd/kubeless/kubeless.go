/*
Copyright (c) 2016-2017 Bitnami

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

// Serverless framework for Kubernetes.
package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var globalUsage = `` //TODO: add explanation

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "kubeless",
		Short: "Serverless framework for Kubernetes",
		Long:  globalUsage,
	}

	cmd.AddCommand(functionCmd, topicCmd, versionCmd, ingressCmd, autoscaleCmd)
	return cmd
}

func main() {
	cmd := newRootCmd()
	fmt.Println("Cli starts here!!!")
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
