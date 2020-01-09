package gapi

import (
	"context"
	"fmt"
	"net"

	"github.com/eviot/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ActivePipe interface {
	InitPipe(req *InitPipeReq)
	DestroyPipe(req *DestroyPipeReq)
}

type InputPipe interface {
	ReceiveMsg(msg *ReceiveMsgReq)
}

type OutputPipe interface {
}

// map[pipe plugin URI]connection
// TODO if necessary, then implement a length check
var sendingConns map[string]*grpc.ClientConn

func init() {
	sendingConns = make(map[string]*grpc.ClientConn)
}

type PluginService struct {
	Port     int
	listener net.Listener
	server   *grpc.Server
	service  *pluginReceiveServiceServer

	activePipes map[string]ActivePipe
	inputPipes  map[string]InputPipe
}

// NewPluginService create PluginService on default port 5550
func NewPluginServiceWithDefaultPort() (*PluginService, error) {
	return NewPluginService( /*Default port*/ 5550)
}

// NewPluginServiceWithPort create PluginService on custom port
func NewPluginService(port int) (*PluginService, error) {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		return nil, fmt.Errorf("Error open listen server %v", err)
	}

	s := grpc.NewServer()
	go s.Serve(lis)
	p := &PluginService{
		Port:     port,
		listener: lis,
		server:   s,

		activePipes: make(map[string]ActivePipe),
		inputPipes:  make(map[string]InputPipe),
	}
	service := newPluginClientServiceServer(p)
	RegisterPluginReceiveServiceServer(s, service)
	p.service = service
	return p, nil
}

func (p *PluginService) RegisterActivePipe(pipeDescriptorName string, pipe ActivePipe) bool {
	_, ok := p.activePipes[pipeDescriptorName]
	if ok {
		return false // this pipeDescriptorName was registered
	}
	p.activePipes[pipeDescriptorName] = pipe
	return true
}

func (p *PluginService) UnregisterActivePipe(pipeDescriptorName string) bool {
	_, ok := p.activePipes[pipeDescriptorName]
	if ok {
		delete(p.activePipes, pipeDescriptorName)
		return true
	}
	return false
}

func (p *PluginService) RegisterInputPipe(pipeDescriptorName string, pipe InputPipe) bool {
	_, ok := p.inputPipes[pipeDescriptorName]
	if ok {
		return false // this pipeDescriptorName was registered
	}
	p.inputPipes[pipeDescriptorName] = pipe
	return true
}

func (p *PluginService) UnregisterInputPipe(pipeDescriptorName string) bool {
	_, ok := p.inputPipes[pipeDescriptorName]
	if ok {
		delete(p.inputPipes, pipeDescriptorName)
		return true
	}
	return false
}

func (p *PluginService) WaitForClose() {
	ch := make(chan bool)
	<-ch // never ends
}

type pluginReceiveServiceServer struct {
	plugin *PluginService
}

func newPluginClientServiceServer(plugin *PluginService) *pluginReceiveServiceServer {
	return &pluginReceiveServiceServer{
		plugin: plugin,
	}
}

func (p *pluginReceiveServiceServer) InitPipe(ctx context.Context, req *InitPipeReq) (*InitPipeRes, error) {
	if pipe, ok := p.plugin.activePipes[req.Pipe.PipeDescriptorName]; ok {
		pipe.InitPipe(req)
	}
	return &InitPipeRes{}, nil
}

func (p *pluginReceiveServiceServer) DestroyPipe(ctx context.Context, req *DestroyPipeReq) (*DestroyPipeRes, error) {
	if pipe, ok := p.plugin.activePipes[req.PipeDescriptorName]; ok {
		pipe.DestroyPipe(req)
	}
	return &DestroyPipeRes{}, nil
}

func (p *pluginReceiveServiceServer) ReceiveMsg(ctx context.Context, req *ReceiveMsgReq) (*ReceiveMsgRes, error) {
	if pipe, ok := p.plugin.inputPipes[req.Pipe.PipeDescriptorName]; ok {
		go pipe.ReceiveMsg(req)
	}
	return &ReceiveMsgRes{}, nil
}

func (*pluginReceiveServiceServer) ReceiveStreamMsg(srv PluginReceiveService_ReceiveStreamMsgServer) error {
	return status.Errorf(codes.Unimplemented, "method ReceiveStreamMsg not implemented")
}

func SendMsg(ctx context.Context, fromPipe *PipeExt, fromOutputID string, msg map[string]*Any) {
	for _, bind := range fromPipe.Binds {
		if bind.FromOutputId != fromOutputID {
			continue
		}
		uri := bind.ToPipe.PluginUri
		log.Info(uri)
		if len(uri) == 0 {
			continue
		}
		conn, ok := sendingConns[uri]
		if !ok {
			var err error
			conn, err = grpc.Dial(uri, grpc.WithInsecure())
			if err != nil {
				continue
			}
			sendingConns[uri] = conn
		}
		log.Info(conn, ok)
		req := &ReceiveMsgReq{
			Pipe:    bind.ToPipe,
			InputId: bind.ToInputId,
			Message: msg,
		}
		service := NewPluginReceiveServiceClient(conn)
		res, err := service.ReceiveMsg(ctx, req)
		log.Info(res, err)
	}
}

func SendStreamMsg(ctx context.Context, fromPipe *Pipe, fromOutputID string, msg map[string]*Any) {
	// TODO: implement
}
