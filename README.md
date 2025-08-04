#  Cachet

**Cachet** is a lightweight, Redis-inspired in-memory key-value store written in Go. 

---

##  Features

- Simple TCP server  
- `SET` and `GET` commands (more to come)  
- In-memory data store using `map[string]string`  
- Basic benchmarking tool  
- Minimal dependencies  

---

##  Getting Started

### 1. Clone the repo

```bash
git clone https://github.com/your-username/cachet.git
cd cachet
```

### 2. Run the server
```bash
go run cmd/main.go
```
### 3. Connect via nc or telnet
```bash
nc localhost 6380
```
or
```bash
telnet localhost 6380
```

### 4. Try Commands 

```sql
SET name Kelvin
GET name
```





