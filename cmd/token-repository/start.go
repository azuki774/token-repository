package main

import (
	"context"
	"time"
	"token-repository/internal/factory"

	"github.com/spf13/cobra"
	"go.uber.org/zap"
)

type StartOption struct {
	Logger *zap.Logger
	Port   string
	DBInfo struct {
		Host string
		Port string
		User string
		Pass string
		Name string
	}
}

var startOpt StartOption

func init() {
	jst, err := time.LoadLocation("Asia/Tokyo")
	if err != nil {
		panic(err)
	}
	time.Local = jst
}

// startCmd represents the start command
var startCmd = &cobra.Command{
	Use:   "start",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:
Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	SilenceErrors: true,
	SilenceUsage:  true,
	RunE: func(cmd *cobra.Command, args []string) error {
		return start(&startOpt)
	},
}

func start(opts *StartOption) error {
	l, err := factory.NewLogger()
	if err != nil {
		return err
	}

	repo, err := factory.NewOAuth2Repo(l, opts.DBInfo.User, opts.DBInfo.Pass, opts.DBInfo.Host, opts.DBInfo.Port, opts.DBInfo.Name)
	if err != nil {
		return err
	}

	svc := factory.NewTokenRepoService(l, repo)
	srv := factory.NewServer(l, opts.Port, svc)
	ctx := context.Background()
	return srv.Start(ctx)
}

func init() {
	rootCmd.AddCommand(startCmd)
	startCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	startCmd.Flags().StringVar(&startOpt.Port, "port", "80", "DB Host")
	startCmd.Flags().StringVar(&startOpt.DBInfo.Host, "db-host", "localhost", "DB Host")
	startCmd.Flags().StringVar(&startOpt.DBInfo.Port, "db-port", "3306", "DB Port")
	startCmd.Flags().StringVar(&startOpt.DBInfo.Name, "db-name", "tokenrepo", "DB Name")
	startCmd.Flags().StringVar(&startOpt.DBInfo.User, "db-user", "root", "DB User")
	startCmd.Flags().StringVar(&startOpt.DBInfo.Pass, "db-pass", "password", "DB Pass")
}
