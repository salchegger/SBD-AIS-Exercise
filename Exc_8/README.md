# SBD Exercise 8
In today's exercise we are going to create and implement a gRPC client and server.
Before starting with the exercise, please install the latest version of `protoc`,
which can be found here: https://protobuf.dev/installation/.
`protoc` is required to generate a gRPC client / server from a Protobuf definition.

Create a OrderService `service` with `rpc` definitions and suitable `messages`.
Store orders in-memory (no need for a database) and return them when requested.
The server should listen on any free port on your machine (i.e. `:4000`)

The gRPC server should be able to serve the following routes:
- `OrderDrink`
- `GetDrinks`
- `GetOrders`

- [ ] Create a protobuf definition
- [ ] Generate a Go client and server
- [ ] Implement the `OrderServiceServer` interface with all it's functions (`OrderDrink`, `GetDrinks`, `GetOrders`)
  - Make sure to embed `pb.UnimplementedOrderServiceServer` in your server struct for forward compatibility
  - Prepopulate drinks and store orders in-memory
- [ ] Create a client that does the following tasks sequentially
  - List drinks
  - Order a few drinks
  - Order more drinks
  - Get order total
- [ ] Start the server in a go routine and then run the client

The final output of your program should look something like this:
```
Requesting drinks 🍹🍺☕
Available drinks:
	> id:1  name:"Spritzer"  price:2  description:"Wine with soda"
	> id:2  name:"Beer"  price:3  description:"Hagenberger Gold"
	> id:3  name:"Coffee"  description:"Mifare isn't that secure"
Ordering drinks 👨‍🍳⏱️🍻🍻
	> Ordering: 2 x Spritzer
	> Ordering: 2 x Beer
	> Ordering: 2 x Coffee
Ordering another round of drinks 👨‍🍳⏱️🍻🍻
	> Ordering: 6 x Spritzer
	> Ordering: 6 x Beer
	> Ordering: 6 x Coffee
Getting the bill 💹💹💹
	> Total: 8 x Spritzer
	> Total: 8 x Beer
	> Total: 8 x Coffee
Orders complete!
```

## Tips and Tricks
Have a look at the gRPC quick start (https://grpc.io/docs/languages/go/quickstart/)
and the tutorial (https://grpc.io/docs/languages/go/basics/).

You can use `google.protobuf.BoolValue` to return booleans 
and `google.protobuf.Empty` to accept empty requests.



