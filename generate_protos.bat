protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/ApplicationRegisterService.proto
protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/Common.proto
protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/DiscoveryService.proto
protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/Downstream.proto
protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/JVMMetricsService.proto
protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/KeyWithIntegerValue.proto
protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/KeyWithStringValue.proto
protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/NetworkAddressRegisterService.proto
protoc -I collector-protocol/ --go_out=plugins=grpc:proto collector-protocol/TraceSegmentService.proto
