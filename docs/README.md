# kingpin

背景：想使用 prometheus 的 node_expoter 项目，但该项
     目里引用了kingpin.v2 库，由于我的项目使用了别的命
     令行库，会和 kingpin.v2 冲突；

目的：兼容 node_expoter项目引用的kingpin.v2 库；
     
# 使用

例: go mod

在自己项目的 go.mod 文件中添加如下语句（用来替换原生的kingpin）
```
replace gopkg.in/alecthomas/kingpin.v2 v2.2.6 => github.com/zhanglp92/kingpin master
```

执行下边命令，使其生效
```
go mod tidy -v
go mod vendor -v
```

获取flag 及默认值接口
```
kingpin.GetFlags()
```

修改flag 值
Set Values (to main.go init function)

例：
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
