
# Cachet

**Cachet** is a high-performance, Redis-inspired in-memory key-value store written in Go. Built for speed and simplicity, Cachet delivers exceptional performance for basic key-value operations while maintaining a clean, extensible architecture.

## Features

- **High-Performance Operations**: Optimized in-memory storage with millions of operations per second
- **Redis-Compatible Commands**: Support for GET, SET, DELETE, EXISTS, and more
- **String Operations**: INCR, DECR, APPEND, STRLEN for numeric and text manipulation
- **Concurrent Access**: Thread-safe operations with efficient locking mechanisms
- **TCP Server**: Network-accessible server with simple text protocol
- **Comprehensive Benchmarking**: Built-in performance testing and comparison tools
- **Minimal Dependencies**: Pure Go implementation with standard library only

## Performance Highlights

Cachet significantly outperforms Redis for basic operations:

- **SET**: 2.7M ops/sec (38x faster than Redis)
- **GET**: 6.0M ops/sec (85x faster than Redis)
- **Concurrent Operations**: Up to 18.8M ops/sec for concurrent reads
- **Memory Efficient**: ~188 bytes per key-value pair average

## Getting Started

### Prerequisites

- Go 1.21 or higher
- Network access to port 6380 (configurable)

### Installation

1. Clone the repository:
```bash
git clone https://github.com/kelvinwambua/cachet.git
cd cachet
```

2. Initialize Go modules:
```bash
go mod tidy
```

3. Build the project:
```bash
go build -o cachet cmd/cachet/main.go
```

### Running the Server

Start the Cachet server:
```bash
go run cmd/cachet/main.go
```

The server will start on `localhost:6380` by default.

### Connecting to the Server

Connect using netcat:
```bash
nc localhost 6380
```

Or using telnet:
```bash
telnet localhost 6380
```

### Basic Commands

Once connected, try these commands:

```
SET username john_doe
GET username
SET counter 42
INCR counter
APPEND username _admin
EXISTS username
DELETE counter
KEYS
SIZE
CLEAR
PING
EXIT
```

## Supported Commands

### String Operations
- `GET key` - Retrieve value for key
- `SET key value` - Store key-value pair
- `INCR key` - Increment numeric value (creates if not exists)
- `DECR key` - Decrement numeric value
- `APPEND key value` - Append to existing string value
- `STRLEN key` - Get length of string value

### Key Management
- `EXISTS key` - Check if key exists (returns 1 or 0)
- `DELETE key` - Remove key-value pair
- `KEYS` - List all keys in store
- `SIZE` - Get total number of keys
- `CLEAR` - Remove all keys from store

### Server Commands
- `PING` - Test server connectivity (returns PONG)
- `EXIT` - Close connection

## Project Structure

```
Cachet/
├── cmd/cachet/
│   └── main.go              # Application entry point
├── internal/
│   ├── store/
│   │   ├── store.go         # Store interface definition
│   │   └── memory.go        # In-memory implementation
│   └── server/
│       ├── server.go        # TCP server implementation
│       └── handler.go       # Command processing logic
├── benchmark.go             # Performance testing suite
├── go.mod                   # Go module definition
└── README.md               # This file
```

## Performance Benchmarks

### Basic Operations
| Operation | Throughput | Performance |
|-----------|------------|-------------|
| SET | 2,745,955 ops/sec | 36.4ms for 100k items |
| GET | 6,089,059 ops/sec | 16.4ms for 100k items |
| EXISTS | 5,691,325 ops/sec | 17.6ms for 100k items |
| DELETE | 5,862,008 ops/sec | 8.5ms for 50k items |

### String Operations
| Operation | Throughput | Notes |
|-----------|------------|-------|
| INCR | 12,626,900 ops/sec | Counter increment operations |
| APPEND | 1,178,287 ops/sec | String concatenation |

### Concurrent Performance (10 Goroutines)
| Operation | Throughput | Performance |
|-----------|------------|-------------|
| Concurrent SET | 1,942,272 ops/sec | 100k total operations |
| Concurrent GET | 18,830,973 ops/sec | 100k total operations |

### Memory Efficiency
- Memory per item: 187.84 bytes average
- 1M items: 179.14 MB total memory usage
- Efficient Go map-based storage with minimal overhead

### Running Benchmarks

Execute the comprehensive benchmark suite:
```bash
go run benchmark.go
```

## Comparison with Redis

### Single-Threaded Operations
| Operation | Cachet | Redis (typical) | Advantage |
|-----------|--------|-----------------|-----------|
| SET | 2.7M ops/sec | ~70K ops/sec | 38x faster |
| GET | 6.0M ops/sec | ~71K ops/sec | 85x faster |
| EXISTS | 5.6M ops/sec | ~60K ops/sec | 93x faster |
| DELETE | 5.8M ops/sec | ~65K ops/sec | 89x faster |

### Concurrent Operations
| Operation | Cachet | Redis | Advantage |
|-----------|--------|-------|-----------|
| Concurrent SET | 1.9M ops/sec | ~500K ops/sec | 4x faster |
| Concurrent GET | 18.8M ops/sec | ~800K ops/sec | 24x faster |

Note: Performance advantages come from direct memory access, no network serialization overhead, and optimized Go implementation.

## Development

### Testing

Run the test suite:
```bash
go test ./...
```

### Contributing

1. Fork the repository
2. Create a feature branch
3. Make your changes
4. Add tests for new functionality
5. Run benchmarks to ensure performance
6. Submit a pull request

## Roadmap

### Phase 1: Core Functionality (Current)
- [x] Basic string operations (GET, SET, DELETE, EXISTS)
- [x] String manipulation (INCR, DECR, APPEND, STRLEN)
- [x] Key management operations
- [x] Concurrent access support
- [x] Performance benchmarking

### Phase 2: Advanced Data Types
- [ ] Lists (LPUSH, RPUSH, LPOP, RPOP, LLEN, LRANGE)
- [ ] Hashes (HSET, HGET, HDEL, HKEYS, HVALS)
- [ ] Sets (SADD, SREM, SMEMBERS, SINTER, SUNION)
- [ ] Sorted Sets (ZADD, ZREM, ZRANGE, ZSCORE)

### Phase 3: Persistence & Reliability
- [ ] AOF (Append-Only File) persistence
- [ ] RDB snapshots
- [ ] Memory optimization and eviction policies
- [ ] Transaction support (MULTI, EXEC)

### Phase 4: Production Features
- [ ] Replication and clustering
- [ ] Authentication and security
- [ ] HTTP REST API
- [ ] Client libraries for popular languages

## License

This project is licensed under the MIT License - see the LICENSE file for details.

## Acknowledgments

- Inspired by Redis and its elegant design
- Built with Go's excellent concurrency primitives
- Performance optimizations based on modern database techniques
  
