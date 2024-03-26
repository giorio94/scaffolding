// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	cmd := &cobra.Command{
		Use:   "cmapisrv-mock",
		Short: "Run the ClusterMesh apiserver mock",
	}

	cmd.AddCommand()

	if err := cmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
