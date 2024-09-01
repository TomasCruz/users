package httphandler

import (
	"context"
	"time"
)

func (h HTTPHandler) Close() {
	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdownCancel()

	// shut down HTTP server
	if err := h.e.Shutdown(shutdownCtx); err != nil {
		h.e.Logger.Fatal(err)
	}
}
