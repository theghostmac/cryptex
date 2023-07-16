package server

type StartRunner struct {
	ListenAddr string
}

func (r *StartRunner) Run() error {
	server := &GracefulShutdown{
		ListenAddr:  r.ListenAddr,
		BaseHandler: nil, // Replace nil with your actual HTTP handler if needed.
	}

	server.Start()
	return nil
}
