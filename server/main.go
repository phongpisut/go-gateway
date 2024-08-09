package main

import (
	"context"
	"crypto/subtle"
	"fmt"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/alitto/pond"
	"github.com/danielgtaylor/huma/v2"
	"github.com/danielgtaylor/huma/v2/adapters/humaecho"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/soheilhy/cmux"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	_ "github.com/danielgtaylor/huma/v2/formats/cbor"
	gen "github.com/phongpisut/go-gateway/gen/hello/v1"
	math "github.com/phongpisut/go-gateway/gen/math/v1"
)

// GreetingOutput represents the greeting operation response.
type GreetingOutput struct {
	Body struct {
		Message string `json:"message" example:"Hello, world!" doc:"Greeting message"`
		Age     int    `json:"age" example:"30" doc:"Age of the person"`
	}
}

type GreetingRequest struct {
	Name string `path:"name" maxLength:"30" example:"world" doc:"Name to greet"`
	Body *struct {
		Age int `json:"age,omitempty" maxLength:"3" example:"30" doc:"Age of the person"`
	}
}

type GreeterServerImpl struct {
	gen.UnimplementedGreeterServiceServer
}

type MathServerImpl struct {
	math.UnimplementedMathServiceServer
}

func main() {
	// Create a new router & API
	logger, _ := zap.NewProduction()
	defer logger.Sync()
	sugar := logger.Sugar()

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	gen.RegisterGreeterServiceServer(grpcServer, &GreeterServerImpl{})
	math.RegisterMathServiceServer(grpcServer, &MathServerImpl{})

	mux := runtime.NewServeMux()

	router := echo.New()
	router.Use(middleware.Logger())

	config := huma.DefaultConfig("My API", "1.0.0")
	config.Components.SecuritySchemes = map[string]*huma.SecurityScheme{
		"bearer": {
			Type:         "http",
			Scheme:       "bearer",
			BearerFormat: "JWT",
		},
	}
	config.DocsPath = ""

	router.GET("/docs", func(c echo.Context) error {
		return c.HTML(http.StatusOK, `<!doctype html>
		<html lang="en">
		  <head>
			<meta charset="utf-8">
			<meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">
			<title>Elements in HTML</title>
			<!-- Embed elements Elements via Web Component -->
			<script src="https://unpkg.com/@stoplight/elements/web-components.min.js"></script>
			<link rel="stylesheet" href="https://unpkg.com/@stoplight/elements/styles.min.css">
		  </head>
		  <body>
		
			<elements-api
			  apiDescriptionUrl="/openapi.yaml"
			  router="hash"
			  layout="sidebar"
			/>
		
		  </body>
		</html>`)
	}, middleware.BasicAuth(func(username, password string, c echo.Context) (bool, error) {
		// Be careful to use constant time comparison to prevent timing attacks
		if subtle.ConstantTimeCompare([]byte(username), []byte("admin")) == 1 &&
			subtle.ConstantTimeCompare([]byte(password), []byte("1234")) == 1 {
			return true, nil
		}
		return false, nil
	}))

	api := humaecho.New(router, config)

	router.Group("/v1/*{grpc_gateway}").Any("", echo.WrapHandler(mux))

	err := gen.RegisterGreeterServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8888", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	if err != nil {
		sugar.Fatal(err)
	}
	err_math := math.RegisterMathServiceHandlerFromEndpoint(context.Background(), mux, "localhost:8888", []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	if err_math != nil {
		sugar.Fatal(err)
	}

	huma.Get(api, "/greeting/{name}", GreetingHandler)

	serverListener, err := net.Listen("tcp", ":8888")
	if err != nil {
		sugar.Fatal(err)
	}

	m := cmux.New(serverListener)

	httpL := m.Match(cmux.HTTP1Fast())
	grpcL := m.Match(cmux.HTTP2())

	go http.Serve(httpL, router)
	go grpcServer.Serve(grpcL)

	m.Serve()

	zap.L().Info("Server started on :8888")

}

func GreetingHandler(ctx context.Context, input *GreetingRequest) (*GreetingOutput, error) {
	resp := &GreetingOutput{}
	resp.Body.Message = fmt.Sprintf("hello, %s", input.Name)
	if input.Body != nil && input.Body.Age > 0 {
		resp.Body.Age = input.Body.Age
	}

	return resp, nil
}

func (g *GreeterServerImpl) SayHello(ctx context.Context, request *gen.SayHelloRequest) (*gen.SayHelloResponse, error) {
	return &gen.SayHelloResponse{
		Message: fmt.Sprintf("hello %s", request.Name),
	}, nil
}

func (g *GreeterServerImpl) StreamText(request *gen.StreamTextRequest, srv gen.GreeterService_StreamTextServer) error {

	pool := pond.New(5, 1000)

	zap.L().Info("Task Start : 🚀")

	for i := 0; i < 10; i++ {
		n := i
		pool.Submit(func() {
			time.Sleep(5 * time.Second)
			resp := gen.StreamTextResponse{Message: fmt.Sprintf("Request #%s TaskID :%d", request.Text, n)}
			if err := srv.Send(&resp); err != nil {
				log.Printf("send error %v", err)
			}
		})

	}
	pool.StopAndWait()
	fmt.Println("Task success : 🚀")
	return nil
}

func (g *MathServerImpl) Add(ctx context.Context, request *math.AddRequest) (*math.AddResponse, error) {
	return &math.AddResponse{
		Sum: request.Num1 + request.Num2,
	}, nil
}

func (g *MathServerImpl) Multiply(ctx context.Context, request *math.MultiplyRequest) (*math.MultiplyResponse, error) {
	return &math.MultiplyResponse{
		Sum: request.Num1 * request.Num2,
	}, nil
}
