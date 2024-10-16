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

func runTemplate(fn func(_m *migrate.Migrate) error) error {
	m, err := newMigrate()
	if err != nil {
		util.FatalLog(fmt.Sprintf("Failed to make migrate instance: %v", err))
	}
	defer func() {
		srcErr, dbErr := m.Close()
		if srcErr != nil {
			util.ErrorLog(fmt.Sprintf("Source close error: %v", srcErr))
		}
		if dbErr != nil {
			util.ErrorLog(fmt.Sprintf("Database close error: %v", dbErr))
		}
	}()
	if err := fn(m); err != nil {
		return err
	}
	util.InfoLog("Migration finished successfully")
	return nil
}

func newMigrateUpCmd() *cobra.Command {
	migrateUpCmd := &cobra.Command{
		Use:   "up",
		Short: "Up migration command.",
		Long:  "Up migration command.",
		RunE: func(cmd *cobra.Command, args []string) error {
			return runTemplate(func(m *migrate.Migrate) error {
				return m.Up()
			})
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
			return runTemplate(func(m *migrate.Migrate) error {
				return m.Down()
			})
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

			return runTemplate(func(m *migrate.Migrate) error {
				return m.Migrate(uint(ver))
			})
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

			return runTemplate(func(m *migrate.Migrate) error {
				return m.Force(ver)
			})
		},
	}
	return migrateForceCmd
}
