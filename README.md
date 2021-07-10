# gRPC-test
スターティングgRPCの第3章から第5章までをGOで実装。  

gRPCサーバー起動
```shell
go run server/server.go
2021/07/11 04:08:46 start gRPC server port: 50051
Baked a pancake for God !
{"level":"info","ts":1625944135.289278,"caller":"zap/options.go:212","msg":"finished unary call with code OK","grpc.start_time":"2021-07-11T04:08:55+09:00","system":"grpc","span.kind":"server","grpc.service":"pancake.maker.PancakeBakerService","grpc.method":"Bake","grpc.code":"OK","grpc.time_ms":0.07400000095367432}
Baked a pancake for God !
{"level":"info","ts":1625944135.291156,"caller":"zap/options.go:212","msg":"finished unary call with code OK","grpc.start_time":"2021-07-11T04:08:55+09:00","system":"grpc","span.kind":"server","grpc.service":"pancake.maker.PancakeBakerService","grpc.method":"Bake","grpc.code":"OK","grpc.time_ms":0.1940000057220459}
{"level":"info","ts":1625944135.29218,"caller":"zap/options.go:212","msg":"finished unary call with code OK","grpc.start_time":"2021-07-11T04:08:55+09:00","system":"grpc","span.kind":"server","grpc.service":"pancake.maker.PancakeBakerService","grpc.method":"Report","grpc.code":"OK","grpc.time_ms":0.0430000014603138}
{"level":"info","ts":1625944135.296078,"caller":"zap/grpclogger.go:92","msg":"[transport]transport: loopyWriter.run returning. connection error: desc = \"transport is closing\"","system":"grpc","grpc_log":true}
```

gRPCクライアント
```shell
go run client/bakery.go
2021/07/11 04:08:55 Start bake
2021/07/11 04:08:55 Response from Server: pancake:{chef_name:"gami" menu:CLASSIC technical_score:0.11529837 create_time:{seconds:1625944135 nanos:289256000}}
2021/07/11 04:08:55 Start bake
2021/07/11 04:08:55 Response from Server: pancake:{chef_name:"gami" menu:BANANA_AND_WHIP technical_score:0.36213678 create_time:{seconds:1625944135 nanos:291124000}}
2021/07/11 04:08:55 Start report
2021/07/11 04:08:55 Response from Server: report:{bake_counts:{menu:CLASSIC count:1} bake_counts:{menu:BANANA_AND_WHIP count:1}}

```