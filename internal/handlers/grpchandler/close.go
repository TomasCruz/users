package grpchandler

func (g GRPCHandler) Close() {
	g.server.GracefulStop()
}
