// The frontend command runs a Google server that combines results
// from multiple backends.
package main

import (
	"flag"
	"log"
	"net"
	"net/http"
	_ "net/http/pprof"
	"strings"
	"sync"

	"golang.org/x/net/context"
	pb "golang.org/x/talks/2015/gotham-grpc/search"
	"google.golang.org/grpc"
)

var (
	backends = flag.String("backends", "localhost:36061,localhost:36062", "comma-separated backend server addresses")
)

type server struct {
	backends []pb.GoogleClient
}

// Search issues Search RPCs in parallel to the backends and returns
// the first result.
func (s *server) Search(ctx context.Context, req *pb.Request) (*pb.Result, error) { // HL
	c := make(chan result, len(s.backends))
	for _, b := range s.backends {
		go func(backend pb.GoogleClient) { // HL
			res, err := backend.Search(ctx, req) // HL
			c <- result{res, err}                // HL
		}(b) // HL
	}
	first := <-c                // HL
	return first.res, first.err // HL
}

type result struct {
	res *pb.Result
	err error
}

// Watch runs Watch RPCs in parallel on the backends and returns a
// merged stream of results.
func (s *server) Watch(req *pb.Request, stream pb.Google_WatchServer) error { // HL
	ctx := stream.Context()
	c := make(chan result) // HL
	var wg sync.WaitGroup
	for _, b := range s.backends {
		wg.Add(1)
		go func(backend pb.GoogleClient) { // HL
			defer wg.Done()                    // HL
			watchBackend(ctx, backend, req, c) // HL
		}(b) // HL
	}
	go func() {
		wg.Wait()
		close(c) // HL
	}()
	for res := range c { // HL
		if res.err != nil {
			return res.err
		}
		if err := stream.Send(res.res); err != nil { // HL
			return err // HL
		} // HL
	}
	return nil
}

// watchBackend runs Watch on a single backend and sends results on c.
// watchBackend returns when ctx.Done is closed or stream.Recv fails.
func watchBackend(ctx context.Context, backend pb.GoogleClient, req *pb.Request, c chan<- result) {
	stream, err := backend.Watch(ctx, req) // HL
	if err != nil {
		select {
		case c <- result{err: err}: // HL
		case <-ctx.Done():
		}
		return
	}
	for {
		res, err := stream.Recv() // HL
		select {
		case c <- result{res, err}: // HL
			if err != nil {
				return
			}
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	flag.Parse()
	go http.ListenAndServe(":36660", nil)   // HTTP debugging
	lis, err := net.Listen("tcp", ":36060") // RPC port
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := new(server)
	for _, addr := range strings.Split(*backends, ",") {
		conn, err := grpc.Dial(addr, grpc.WithInsecure())
		if err != nil {
			log.Fatalf("fail to dial: %v", err)
		}
		client := pb.NewGoogleClient(conn)
		s.backends = append(s.backends, client)
	}
	g := grpc.NewServer()
	pb.RegisterGoogleServer(g, s)
	g.Serve(lis)
}
