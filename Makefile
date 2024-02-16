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

.PHONY: obu receiver dcalc aggregator
