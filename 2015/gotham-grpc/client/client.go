// The client command issues RPCs to a Google server and prints the
// results.
//
// In "search" mode, client calls Search on the server and prints the
// results.
//
// In "watch" mode, client starts a Watch on the server and prints the
// result stream.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"time"

	"golang.org/x/net/context"
	pb "golang.org/x/talks/2015/gotham-grpc/search"
	"google.golang.org/grpc"
)

var (
	server = flag.String("server", "localhost:36060", "server address")
	mode   = flag.String("mode", "search", `one of "search" or "watch"`)
	query  = flag.String("query", "test", "query string")
)

func main() {
	flag.Parse()

	// Connect to the server.
	conn, err := grpc.Dial(*server, grpc.WithInsecure()) // HL
	if err != nil {
		log.Fatalf("fail to dial: %v", err)
	}
	defer conn.Close()
	client := pb.NewGoogleClient(conn) // HL

	// Run the RPC.
	switch *mode {
	case "search":
		search(client, *query) // HL
	case "watch":
		watch(client, *query)
	default:
		log.Fatalf("unknown mode: %q", *mode)
	}
}

// search issues a search for query and prints the result.
func search(client pb.GoogleClient, query string) {
	ctx, cancel := context.WithTimeout(context.Background(), 80*time.Millisecond) // HL
	defer cancel()
	req := &pb.Request{Query: query}    // HL
	res, err := client.Search(ctx, req) // HL
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res) // HL
}

// watch runs a Watch RPC and prints the result stream.
func watch(client pb.GoogleClient, query string) {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	req := &pb.Request{Query: query}      // HL
	stream, err := client.Watch(ctx, req) // HL
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv() // HL
		if err == io.EOF {        // HL
			fmt.Println("and now your watch is ended")
			return
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res) // HL
	}
}
