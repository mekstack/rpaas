# NATaaS core Proof-of-Concept implementation

## Project structure

```txt
├── Makefile
├── proto
│   └── objects.proto
├── README.md
├── requirements.in
├── requirements.txt
└── src
    ├── configs
    ├── controllers
    ├── grpc_server
    ├── main.py
    ├── proto
    ├── repositories
    └── services
```

`Makefile` - file is designed to build Python code from Google protobuf objects.
To execute the input:

```shell
make proto_objects
```

`proto` - directory for storing *.proto files.
`requirements.in` - file declares Python packages versions.
To execute the input:

```shell
pip install pip-tools
pip-compile
pip-sync
```

> After `pip-compile` you get requirements.txt

`src` - main directory of __NATaaS Core Service__:

+ `configs` - directory stores the service configuration
+ `controllsers` - directory stores the business logic of service
+ `grpc_server` - directory stores the logic of the server
+ `main.py` - service starting point
+ `proto` - directory with compiled files from protobuf into Python code
+ `repositories` - directory which help `controllers` interacting with the database
+ `services` - directory which work with gRPC requests

## How to start?

Set environment variables:

```shell
export REDIS_HOST=YOUR_HOST_FOR_CONNECTING_REDIS
export REDIS_PORT=YOUR_PORT_FOR_CONNECTING_REDIS
export SERVER_HOST=YOUR_HOST_FOR_RUN_SERVER
export SERVER_PORT=YOUR_PORT_FOR_RUN_SERVER
```

After, run command:

```shell
python3.11 main.py
```