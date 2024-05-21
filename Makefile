generate_code:
	protoc -I ./api ./api/endpoint_management.proto --go_out=./pkg/endpoint_management --go_opt=paths=source_relative --go-grpc_out=./pkg/endpoint_management --go-grpc_opt=paths=source_relative

run:
	go run ./cmd/xds --debug=true

profile:
	go run ./cmd/xds --profile=true --debug=true

mem-profile:
	go run ./cmd/xds --mem-profile=true --debug=true

cpu-profile:
	go run ./cmd/xds --cpu_profile=true --debug=true

bench:
	go test -run ^$ -bench . -benchmem > bencg.txt

.PHONY: generate_code, run, bench, profile, mem-profile, cpu-profile

.DEFAULT_GOAL=run