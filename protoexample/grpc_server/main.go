package main

import (
	"context"
	"fmt"
	"net/http"

	"connectrpc.com/connect"
	connectcors "connectrpc.com/cors"
	"github.com/rs/cors"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"

	greetv1 "github.com/griggsca91/protodevtools/protoexample/gen/go" // generated by protoc-gen-go

	"github.com/griggsca91/protodevtools/protoexample/gen/go/greetv1connect" // generated by protoc-gen-connect-go
)

type GreetServer struct{}

func withCORS(h http.Handler) http.Handler {
	middleware := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: connectcors.AllowedMethods(),
		AllowedHeaders: connectcors.AllowedHeaders(),
		ExposedHeaders: connectcors.ExposedHeaders(),
	})
	return middleware.Handler(h)
}

// SendComplexObject implements greetv1connect.GreeterHandler.
func (s *GreetServer) SendComplexObject(context.Context, *connect.Request[greetv1.ComplexObjectRequest]) (*connect.Response[greetv1.ComplexObjectResponse], error) {
	panic("unimplemented")
}

func (s *GreetServer) SayHello(
	ctx context.Context,
	req *connect.Request[greetv1.HelloRequest],
) (*connect.Response[greetv1.HelloReply], error) {
	b := true

	reply := &greetv1.HelloReply{
		Message: fmt.Sprintf("Hello, %s!", req.Msg.Name),
		Response: &greetv1.ComplexObjectResponse{
			StringValue:         "string",
			RepeatedStringValue: []string{"hello", "world", "hello", "world", "hello", "mars", "bye", "felicia"},
			IntValue:            42,
			BoolValue:           true,
			OptionalBoolValue:   &b,
			FloatValue:          3.14,
			EnumValue:           greetv1.ComplexObjectResponse_CORPUS_LOCAL,
			RepeatedMessageValue: []*greetv1.Result{
				{
					Url:      "url",
					Title:    "title",
					Snippets: []string{"snippet 1", "snippet 2"},
				},
				{
					Url:      "url",
					Title:    "title",
					Snippets: []string{"snippet 1", "snippet 2"},
				},
				{
					Url:      "url",
					Title:    "title",
					Snippets: []string{"snippet 1", "snippet 2"},
				},
				{
					Url:      "url",
					Title:    "title",
					Snippets: []string{"snippet 1", "snippet 2"},
				},
				{
					Url:      "url",
					Title:    "title",
					Snippets: []string{"snippet 1", "snippet 2"},
				},
			},
			OneofValue: &greetv1.ComplexObjectResponse_SubMessage{
				SubMessage: &greetv1.Result{
					Url:      "url",
					Title:    "title",
					Snippets: []string{"snippet 1", "snippet 2"},
				},
			},
			MapValue: map[string]*greetv1.Result{
				"key": {
					Url:      "url",
					Title:    "title",
					Snippets: []string{"snippet 1", "snippet 2"},
				},
			},
		},
	}
	res := connect.NewResponse(reply)

	res.Header().Set("Greet-Version", "v1")
	return res, nil
}

func main() {
	greeter := &GreetServer{}
	mux := http.NewServeMux()
	path, handler := greetv1connect.NewGreeterHandler(greeter)
	mux.Handle(path, withCORS(handler))
	http.ListenAndServe(
		":8080",
		// Use h2c so we can serve HTTP/2 without TLS.
		h2c.NewHandler(mux, &http2.Server{}),
	)
}
