package cmd

import (
	"fmt"
	"net/http"
	"strconv"

	migrate "github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/mysql" // nolint: golint // Necessary for Register MySQL Driver for migration
	"github.com/golang-migrate/migrate/v4/source/httpfs"
	"github.com/spf13/cobra"
	"github.com/takahiroaoki/grpc-sample/app/config"
	"github.com/takahiroaoki/grpc-sample/app/resource"
	"github.com/takahiroaoki/grpc-sample/app/util"
)

func newMigrateCmd() *cobra.Command {
	migrateCmd := &cobra.Command{
		Use: "migrate",
		Run: func(cmd *cobra.Command, args []string) {
			_ = cmd.Help()
		},
	}
	migrateCmd.AddCommand(
		newMigrateUpCmd(),
		newMigrateDownCmd(),
		newMigrateVerCmd(),
		newMigrateForceCmd(),
	)
	return migrateCmd
}

func newMigrate() (*migrate.Migrate, error) {
	rsc, err := httpfs.New(http.FS(resource.FS), "migration")
	if err != nil {
		return nil, err
	}

	return migrate.NewWithSourceInstance("httpfs", rsc, config.GetDataBaseURL())
}

func newMigrateUpCmd() *cobra.Command {
	migrateUpCmd := &cobra.Command{
		Use:   "up",
		Short: "Up migration command.",
		Long:  "Up migration command.",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := newMigrate()
			if err != nil {
				util.FatalLog(fmt.Sprintf("Failed to make migrate instance: %v", err))
			}
			if err := m.Up(); err != nil {
				return err
			}
			util.InfoLog("Up migration finished successfully")
			return nil
		},
	}
	return migrateUpCmd
}

func newMigrateDownCmd() *cobra.Command {
	migrateDownCmd := &cobra.Command{
		Use:   "down",
		Short: "Down migration command.",
		Long:  "Down migration command.",
		RunE: func(cmd *cobra.Command, args []string) error {
			m, err := newMigrate()
			if err != nil {
				util.FatalLog(fmt.Sprintf("Failed to make migrate instance: %v", err))
			}
			if err := m.Down(); err != nil {
				return err
			}
			util.InfoLog("Down migration finished successfully")
			return nil
		},
	}
	return migrateDownCmd
}

func newMigrateVerCmd() *cobra.Command {
	migrateVerCmd := &cobra.Command{
		Use:   "ver",
		Short: "Versioning migration command.",
		Long:  "Versioning migration command. The argument for version is necessary.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				util.FatalLog("The numeric argument is required for version")
			}

			ver, err := strconv.Atoi(args[0])
			if err != nil {
				util.FatalLog("The argument for version must be numeric")
			}

			m, err := newMigrate()
			if err != nil {
				util.FatalLog(fmt.Sprintf("Failed to make migrate instance: %v", err))
			}
			if err := m.Migrate(uint(ver)); err != nil {
				return err
			}

			util.InfoLog("Migration versioning finished successfully")
			return nil
		},
	}
	return migrateVerCmd
}

func newMigrateForceCmd() *cobra.Command {
	migrateForceCmd := &cobra.Command{
		Use:   "force",
		Short: "Force migration command.",
		Long:  "Force migration command. The argument for version is necessary.",
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				util.FatalLog("The numeric argument is required for version")
			}

			ver, err := strconv.Atoi(args[0])
			if err != nil {
				util.FatalLog("The argument for version must be numeric")
			}

			m, err := newMigrate()
			if err != nil {
				util.FatalLog(fmt.Sprintf("Failed to make migrate instance: %v", err))
			}
			if err := m.Force(ver); err != nil {
				return err
			}
			util.InfoLog("Force migration finished successfully")
			return nil
		},
	}
	return migrateForceCmd
}
