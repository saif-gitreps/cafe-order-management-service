package main

func main() {
	gRPCserver := NewGRPCServer(":9000")
	gRPCserver.Run()

}