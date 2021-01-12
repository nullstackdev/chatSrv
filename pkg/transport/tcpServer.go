package transport

type TCPServerConfig struct {

}

type TCPServer struct {

}

//PublicAPI

func (s *TCPServer) Listen()  {

}

func (s *TCPServer) Stop() {

}

//MAKE

func NewTCPServer(cfg *TCPServerConfig) (srv *TCPServer,cfgError error)  {
	srv = nil
	cfgError = nil

	return srv, cfgError
}