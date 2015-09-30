// Tha backend command runs a Google server that returns fake results.
package main

import (
	"flag"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"

	"golang.org/x/net/context"
	"golang.org/x/net/trace"
	pb "golang.org/x/talks/2015/gotham-grpc/search"
	"google.golang.org/grpc"
)

var (
	index = flag.Int("index", 0, "RPC port is 36061+index; debug port is 36661+index")
)

type server struct{}

// randomDuration returns a random duration up to max, at intervals of max/10.
func randomDuration(max time.Duration) time.Duration {
	return time.Duration(1+int64(rand.Intn(10))) * (max / 10)
}

// Search sleeps for a random interval then returns a string
// identifying the query and this backend.
func (s *server) Search(ctx context.Context, req *pb.Request) (*pb.Result, error) { // HL
	d := randomDuration(100 * time.Millisecond)
	logSleep(ctx, d) // HL
	select {
	case <-time.After(d):
		return &pb.Result{ // HL
			Title: fmt.Sprintf("result for [%s] from backend %d", req.Query, *index), // HL
		}, nil // HL
	case <-ctx.Done():
		return nil, ctx.Err()
	}
}

func logSleep(ctx context.Context, d time.Duration) {
	if tr, ok := trace.FromContext(ctx); ok { // HL
		tr.LazyPrintf("sleeping for %s", d) // HL
	}
}

// Watch returns a stream of results identifying the query and this
// backend, sleeping a random interval between each send.
func (s *server) Watch(req *pb.Request, stream pb.Google_WatchServer) error { // HL
	ctx := stream.Context()
	for i := 0; ; i++ {
		d := randomDuration(1 * time.Second)
		logSleep(ctx, d) // HL
		select {
		case <-time.After(d):
			err := stream.Send(&pb.Result{ // HL
				Title: fmt.Sprintf("result %d for [%s] from backend %d", i, req.Query, *index), // HL
			}) // HL
			if err != nil {
				return err
			}
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}

func main() {
	flag.Parse()
	rand.Seed(time.Now().UnixNano())
	go http.ListenAndServe(fmt.Sprintf(":%d", 36661+*index), nil)   // HTTP debugging
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 36061+*index)) // RPC port // HL
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	g := grpc.NewServer()                   // HL
	pb.RegisterGoogleServer(g, new(server)) // HL
	g.Serve(lis)                            // HL
}
