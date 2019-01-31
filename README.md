[中文说明](docs/README.md)

# kingpin
for Prometheus node_expoter remover kingpin

# Use

eg: go mod

add to go.mod
```
replace gopkg.in/alecthomas/kingpin.v2 v2.2.6 => github.com/zhanglp92/kingpin master
```

continue 

```
go mod tidy -v
go mod vendor -v
```

Get all flags
```
kingpin.GetFlags()
```

Set Values (to main.go init function)

eg: 

```
// main.go

func init() {
  // ...
  
  kingpin.SetVal("name1", "v1", "v2")
  kingpin.SetVal("name2", "true")
  kingpin.SetVal("name3", "3.23ms")
  
  // ...
}
```
