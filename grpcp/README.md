
Installation
-----------

```bash
go get -u leeson1995/go-common/grpcp
```

Usage
-----------

default

```go
import (
    "context"
    "fmt"

    "google.golang.org/grpc"
    pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
    var addr, name string
    conn, _ := grpc.Dial(addr, grpc.WithInsecure())
    defer conn.Close()
    client := pb.NewGreeterClient(conn)
    r, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    fmt.Println(r.GetMessage())
}
```

with grpcp

```go
import (
    "context"
    "fmt"

    "google.golang.org/grpc"
    pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
    var addr, name string

    conn, _ := grpcp.GetConn(addr)  // get conn with grpcp default pool
    // defer conn.Close()  // no close, close will disconnect
    client := pb.NewGreeterClient(conn)
    r, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    fmt.Println(r.GetMessage())
}
```

custom dial function

```go
import (
    "context"
    "fmt"

    "leeson1995/go-common/grpcp"
    "google.golang.org/grpc"
    pb "google.golang.org/grpc/examples/helloworld/helloworld"
)

func main() {
    var addr, name string

    pool := grpcp.New(func(addr string) (*grpc.ClientConn, error) {
        return grpc.Dial(
            addr,
            grpc.WithInsecure(),
        )
    })
    conn, _ := pool.GetConn(addr)

    client := pb.NewGreeterClient(conn)
    r, _ := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    fmt.Println(r.GetMessage())
}
```
