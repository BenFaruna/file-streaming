# File streaming
This is a basic [server](main.go) that implements file streaming and a [script](client.js) that acts as the client.

The server is implemented using Golang stdlib (net/http). It features streaming and a simple file upload via two endpoints. This simple implementation is an experiment of the advantages of using streaming over simple file upload.

## Prerequisites
- **Go 1.21+** (or latest stable version)
- **Git** (for cloning the repository)

## Installation
- **Clone the repository**:
```shell
git clone https://github.com/BenFaruna/file-streaming
cd file-streaming
```

## Build & Run
- Build the binary:
```shell
go build -o app
```

- Run the application:
```shell
./app
```

Or run directly without building:
```shell
go run main.go
```

## ENDPOINTS
| ROUTE          | METHOD | DESCRIPTION                                                                      | RESPONSE                                        |
|----------------|--------|----------------------------------------------------------------------------------|-------------------------------------------------|
| /stream-upload | POST   | accepts data stream and ends when the stream closes or when it encounters an EOF | { status: 201, msg: 'file uploaded', ok: true } |
| /file-upload   | POST   | accepts file inputs and save them on the server                                  | { status: 201, msg: 'file uploaded', ok: true } |

## Testing endpoint with client script
### Prerequisite
- Node
- File streaming server running

### Create big file (optional)
- Create big file with about 1.9 GB size using [create_big_file.js](./example/create_big_file.js) in `example` folder.
- Move file from example folder into the root folder
```shell
mv ./example/big.csv .
```

### Executing the script
- If a different file is used from the big file created above, edit `filePath` variable in [client.js](client.js) to the correct file path.

- Execute client script
```shell
node client.js
```

## Result from client test with large files
Using a file greater than 2 GB with the script causes node to fail with the error below.
This was the immediate pointer to the advantage of file streaming over simple upload.
```shell
<--- Last few GCs --->

[251184:0x2b8525c0]    26005 ms: Scavenge 2038.8 (2075.1) -> 2038.1 (2075.8) MB, 7.88 / 0.00 ms  (average mu = 0.166, current mu = 0.133) allocation failure;
[251184:0x2b8525c0]    26016 ms: Scavenge 2039.4 (2075.8) -> 2038.7 (2080.8) MB, 7.62 / 0.00 ms  (average mu = 0.166, current mu = 0.133) allocation failure;
[251184:0x2b8525c0]    26032 ms: Scavenge 2041.9 (2080.8) -> 2040.2 (2080.8) MB, 8.53 / 0.00 ms  (average mu = 0.166, current mu = 0.133) allocation failure;


<--- JS stacktrace --->

FATAL ERROR: Ineffective mark-compacts near heap limit Allocation failed - JavaScript heap out of memory
----- Native stack trace -----

1: 0xb82c28 node::OOMErrorHandler(char const*, v8::OOMDetails const&) [node]
2: 0xeed540 v8::Utils::ReportOOMFailure(v8::internal::Isolate*, char const*, v8::OOMDetails const&) [node]
3: 0xeed827 v8::internal::V8::FatalProcessOutOfMemory(v8::internal::Isolate*, char const*, v8::OOMDetails const&) [node]
4: 0x10ff3c5  [node]
5: 0x10ff954 v8::internal::Heap::RecomputeLimits(v8::internal::GarbageCollector) [node]
6: 0x1116844 v8::internal::Heap::PerformGarbageCollection(v8::internal::GarbageCollector, v8::internal::GarbageCollectionReason, char const*) [node]
7: 0x111705c v8::internal::Heap::CollectGarbage(v8::internal::AllocationSpace, v8::internal::GarbageCollectionReason, v8::GCCallbackFlags) [node]
8: 0x11191ba v8::internal::Heap::HandleGCRequest() [node]
9: 0x1084827 v8::internal::StackGuard::HandleInterrupts() [node]
10: 0x1526f7a v8::internal::Runtime_StackGuard(int, unsigned long*, v8::internal::Isolate*) [node]
11: 0x1960ef6  [node]
Aborted (core dumped)
```

## Challenge
- Sending data over the stream that is not a file causes the connection to freeze since there is no EOF marker.
- Identifying file types over the stream does not work, the `Content-Type` is empty even after setting it on the client script.

## Contributing
If you think of any improvement I can make to this, kindly create an issue. I'm more than happy to explore ways to make this better.