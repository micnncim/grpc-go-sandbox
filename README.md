# grpc-go-sample

## Usage

```sh
$ go run main.go &
$ grpcurl -plaintext -d '{ "name": "micnncim" }' localhost:8080 proto.GreetingService/Hello 
{
  "message": "Bye, micnncim!"
}
```
