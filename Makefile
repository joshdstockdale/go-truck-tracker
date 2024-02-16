obu:
	@go build -o bin/obu obu/main.go ;
	@./bin/obu ;

receiver:
	@go build -o bin/receiver ./receiver ;
	@./bin/receiver ;

dcalc:
	@go build -o bin/dcalc ./distance_calculator ;
	@./bin/dcalc ;

aggregator:
	@go build -o bin/aggregator ./aggregator ;
	@./bin/aggregator ;

proto:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative types/ptypes.proto ;

.PHONY: obu receiver dcalc aggregator
