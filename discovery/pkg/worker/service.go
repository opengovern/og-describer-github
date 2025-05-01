package worker

import (
	"github.com/spf13/cobra"
	"go.uber.org/zap"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func WorkerCommand() *cobra.Command {
	cmd := &cobra.Command{
		RunE: func(cmd *cobra.Command, args []string) error {
			setupCtx := cmd.Context()
			cmd.SilenceUsage = true
			logger, err := zap.NewProduction()
			if err != nil {
				return err
			}

			sigChan := make(chan os.Signal, 1)
			signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
			go func() {
				sig := <-sigChan
				logger.Warn("Received termination signal, worker will shut down when idle timeout is reached or on next check.", zap.String("signal", sig.String()))
			}()

			w, err := NewWorker(
				logger,
				setupCtx,
			)
			if err != nil {
				return err
			}

			// Define the idle timeout duration
			idleTimeout := 2 * time.Minute

			// Run the worker with a background context, independent of signals,
			// and pass the idle timeout.
			logger.Info("Starting worker with idle timeout", zap.Duration("timeout", idleTimeout))
			runErr := w.Run(context.Background(), idleTimeout) // Use context.Background()

			if runErr != nil && !errors.Is(runErr, context.Canceled) { // Avoid logging error for clean context cancellations if any were added back
				logger.Error("Worker Run exited with error", zap.Error(runErr))
				return runErr // Propagate actual errors
			}

			logger.Info("Worker Run finished.")
			return nil
		},
	}

	return cmd
}
