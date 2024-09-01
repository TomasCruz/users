package httphandler

import (
	"context"
	"time"
)

func (h HTTPHandler) Close() error {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// shut down HTTP server
	if err := h.e.Shutdown(shutdownCtx); err != nil {
		h.e.Logger.Error(err)
		return err
	}

	return nil
}
