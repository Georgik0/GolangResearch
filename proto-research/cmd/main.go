package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"

	"github.com/soheilhy/cmux"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
)

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(err)
	}

	m := cmux.New(listener)
	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	g := new(errgroup.Group)
	g.Go(func() error { return grpcServe(grpcListener) })
	g.Go(func() error { return httpServe(httpListener) })
	g.Go(func() error { return m.Serve() })

	log.Println("run server: ", g.Wait())
}

type grpcServer struct {
	hw.UnimplementedHelloWorldServer
}

func (*grpcServer) Say(ctx context.Context, in *hw.Req) (*hw.Res, error) {
	return &hw.Res{Msg: fmt.Sprintf("Hello %s", in.Name)}, nil
}

func grpcServe(l net.Listener) error {
	s := grpc.NewServer()
	hw.RegisterHelloWorldServer(s, &grpcServer{})

	return s.Serve(l)
}

func httpServe(l net.Listener) error {
	mux := http.NewServeMux()
	mux.HandleFunc("/ping/", func(w http.ResponseWriter, _ *http.Request) {
		w.Write([]byte("pong"))
	})

	//conn, err := grpc.DialContext(
	//	context.Background(),
	//	"0.0.0.0:8080",
	//	grpc.WithBlock(),
	//	grpc.WithTransportCredentials(insecure.NewCredentials()),
	//)
	//
	//gwmux := runtime.NewServeMux()
	//err = hw.RegisterHelloWorldHandler(context.Background(), gwmux, conn)
	//if err != nil {
	//	log.Fatalln("Failed to register gateway:", err)
	//}
	//
	//mux.Handle("/", gwmux)

	s := &http.Server{Handler: mux}
	return s.Serve(l)
}
