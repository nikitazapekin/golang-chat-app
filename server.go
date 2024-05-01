/*package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
)
func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err !=nil {
		log.Fatalf("Failed to listen port 9000")
	}
	grpcServer :=grpc.NewServer()
	if err:= grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to  serve gRPC server over port  9000")
	}
}
 */

/*
 package main

import "fmt"

// Интерфейс для всех продуктов
type Product interface {
    GetName() string
}

// Конкретные продукты
type ConcreteProductA struct{}

func (p ConcreteProductA) GetName() string {
    return "Product A"
}

type ConcreteProductB struct{}

func (p ConcreteProductB) GetName() string {
    return "Product B"
}

// Интерфейс фабрики
type Factory interface {
    CreateProduct() Product
}

// Конкретные фабрики
type ConcreteFactoryA struct{}

func (f ConcreteFactoryA) CreateProduct() Product {
    return ConcreteProductA{}
}

type ConcreteFactoryB struct{}

func (f ConcreteFactoryB) CreateProduct() Product {
    return ConcreteProductB{}
}

func main() {
    // Использование фабрик
    factoryA := ConcreteFactoryA{}
    productA := factoryA.CreateProduct()
    fmt.Println(productA.GetName())

    factoryB := ConcreteFactoryB{}
    productB := factoryB.CreateProduct()
    fmt.Println(productB.GetName())
}

*/



// Package main serves as an example application that makes use of the observer pattern.
// Playground: https://play.golang.org/p/cr8jEmDmw0


/*
package main

import (
	"fmt"
	"time"
)

type (

	Event struct {

		Data int64
	}

	Observer interface {
 
		OnNotify(Event)
	}
 
	Notifier interface {
 
		Register(Observer)
 
		Deregister(Observer)
 
		Notify(Event)
	}
)

type (
	eventObserver struct{
		id int
	}

	eventNotifier struct{
 
		observers map[Observer]struct{}
	}
)

func (o *eventObserver) OnNotify(e Event) {
	fmt.Printf("*** Observer %d received: %d\n", o.id, e.Data)
}

func (o *eventNotifier) Register(l Observer) {
	o.observers[l] = struct{}{}
}

func (o *eventNotifier) Deregister(l Observer) {
	delete(o.observers, l)
}

func (p *eventNotifier) Notify(e Event) {
	for o := range p.observers {
		o.OnNotify(e)
	}
}

func main() {
 
	n := eventNotifier{
		observers: map[Observer]struct{}{},
	}

 
	n.Register(&eventObserver{id: 1})
	n.Register(&eventObserver{id: 2})

	 
	stop := time.NewTimer(10 * time.Second).C
	tick := time.NewTicker(time.Second).C
	for {
		select {
		case <- stop:
			return
		case t := <-tick:
			n.Notify(Event{Data: t.UnixNano()})
		}
	}
}
*/




//module github.com/nikita/go-microservices
 
package main

import (
	//"fmt"
	"log"
	//"net/http"
	//"server/db"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
 
 r "github.com/nikita/go-microservices/routes"

)
func main() {
	e := echo.New()
	//db.Connect()
	e.Use(middleware.CORS())



	r.InitRoutes(e)
/*	
*/
 r.InitWebsocketRoutes(e) 
	err := e.Start(":5000")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}


	//handleHistory()
}
 

// go mod init github.com/nikita/go-microservices
//go get -u github.com/golang/protobuf/protoc-gen-go
//go get -u github.com/golang/protobuf/protoc-gen   
//go get -u google.golang.org/grpc
// mkdir chat







//go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
//go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2