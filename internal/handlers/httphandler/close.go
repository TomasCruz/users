package httphandler

import (
	"context"
	"time"

	"github.com/pkg/errors"
)

func (h HTTPHandler) Close() error {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// shut down HTTP server
	if err := h.e.Shutdown(shutdownCtx); err != nil {
		return errors.WithStack(err)
	}

	return nil
}
