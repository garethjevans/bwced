package cmd

import (
	"fmt"
	"net/http"

	"github.com/garethjevans/bwced/pkg/bwced"
	"github.com/spf13/cobra"
)

var (
	BindAddress   string
	Port          int
	DocumentRoot  string
	MaxUploadSize int64
	EnableCORS    bool
)

// NewRunCmd creates a new run command.
func NewRunCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:     "run",
		Short:   "Run the webserver",
		Long:    "",
		Example: "bwced run",
		Aliases: []string{"r"},
		RunE: func(cmd *cobra.Command, args []string) error {
			server := bwced.NewServer(DocumentRoot, MaxUploadSize, EnableCORS, nil)
			a := fmt.Sprintf("%s:%d", BindAddress, Port)
			fmt.Printf("listeninig on %s\n", a)

			return http.ListenAndServe(a, server)
		},
		Args:         cobra.NoArgs,
		SilenceUsage: true,
	}

	cmd.Flags().StringVarP(&BindAddress, "bind-address", "", "localhost", "The address to bind to (default: localhost)")
	cmd.Flags().IntVarP(&Port, "port", "p", 8080, "The port to run the webserver on (default: 8080)")
	cmd.Flags().Int64VarP(&MaxUploadSize, "max-upload-size", "", 1024, "The max upload size in bytes (default: 1024)")
	cmd.Flags().BoolVarP(&EnableCORS, "enable-cors", "", false, "Whether to allow CORS requests (default: false)")
	cmd.Flags().StringVarP(&DocumentRoot, "document-root", "d", "", "The root to store all documents")

	return cmd
}
