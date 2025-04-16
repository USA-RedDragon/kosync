package cmd

import (
	"fmt"
	"log/slog"
	"os"
	"syscall"

	"github.com/USA-RedDragon/configulator"
	"github.com/USA-RedDragon/kosync/internal/config"
	"github.com/USA-RedDragon/kosync/internal/server"
	"github.com/USA-RedDragon/kosync/internal/store"
	"github.com/lmittmann/tint"
	"github.com/spf13/cobra"
	"github.com/ztrue/shutdown"
)

func NewCommand(version, commit string) *cobra.Command {
	return &cobra.Command{
		Use:     "kosync",
		Version: fmt.Sprintf("%s - %s", version, commit),
		Annotations: map[string]string{
			"version": version,
			"commit":  commit,
		},
		RunE:              runRoot,
		SilenceErrors:     true,
		DisableAutoGenTag: true,
	}
}

func runRoot(cmd *cobra.Command, _ []string) error {
	fmt.Printf("kosync - %s (%s)\n", cmd.Annotations["version"], cmd.Annotations["commit"])

	c, err := configulator.FromContext[config.Config](cmd.Context())
	if err != nil {
		return fmt.Errorf("failed to get config from context")
	}

	cfg, err := c.Load()
	if err != nil {
		return err
	}

	var logger *slog.Logger
	switch cfg.LogLevel {
	case config.LogLevelDebug:
		logger = slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelDebug}))
	case config.LogLevelInfo:
		logger = slog.New(tint.NewHandler(os.Stdout, &tint.Options{Level: slog.LevelInfo}))
	case config.LogLevelWarn:
		logger = slog.New(tint.NewHandler(os.Stderr, &tint.Options{Level: slog.LevelWarn}))
	case config.LogLevelError:
		logger = slog.New(tint.NewHandler(os.Stderr, &tint.Options{Level: slog.LevelError}))
	}
	slog.SetDefault(logger)

	store, err := store.NewStore(cfg)
	if err != nil {
		return fmt.Errorf("failed to create store: %w", err)
	}

	slog.Info("Connected to datastore", "type", cfg.Storage.Type)

	server := server.NewServer(cfg, &store)

	if err := server.Start(); err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	slog.Info("Server started successfully")

	stop := func(sig os.Signal) {
		// Remove control codes from the current line in the terminal
		fmt.Println("")

		slog.Info("Received signal", "signal", sig)

		err := server.Stop()
		if err != nil {
			slog.Error("Failed to stop server", "error", err)
		} else {
			slog.Info("Server stopped gracefully")
		}
	}
	shutdown.AddWithParam(stop)
	shutdown.Listen(syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGHUP)

	return nil
}
