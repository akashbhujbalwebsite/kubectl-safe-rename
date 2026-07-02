package cmd

import (
	"github.com/spf13/cobra"
	"k8s.io/client-go/tools/clientcmd"
)

var (
	kubeconfig string
	namespace  string
	dryRun     bool
	yes        bool
)

var rootCmd = &cobra.Command{
	Use:   "kubectl-safe-rename",
	Short: "Safely rename ConfigMaps and Secrets",
	Long: `kubectl safe-rename renames ConfigMaps and Secrets — the only Kubernetes
resources where a rename is safe. These resources are stateless data stores
whose names carry no runtime identity; Kubernetes imposes no live binding on
their names at the API level.

Other resources (Pods, Deployments, Services, ServiceAccounts, PVCs) are
intentionally not supported: their names are tied to DNS, pod identity, or
runtime bindings where a create+delete cycle would break running workloads.`,
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	rootCmd.PersistentFlags().StringVar(&kubeconfig, "kubeconfig", clientcmd.RecommendedHomeFile, "path to kubeconfig")
	rootCmd.PersistentFlags().StringVarP(&namespace, "namespace", "n", "default", "namespace of the resource")
	rootCmd.PersistentFlags().BoolVar(&dryRun, "dry-run", false, "show what would happen without making changes")
	rootCmd.PersistentFlags().BoolVarP(&yes, "yes", "y", false, "skip confirmation prompt")

	rootCmd.AddCommand(newConfigMapCmd())
	rootCmd.AddCommand(newSecretCmd())
}
