package main

import (
	"os"

	"github.com/akashbhujbalwebsite/kubectl-safe-rename/cmd"
	_ "k8s.io/client-go/plugin/pkg/client/auth"
)

func main() {
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
