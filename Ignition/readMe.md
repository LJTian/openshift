# 使用说明


## 修改 example.bu 文件内容
```shell
vim example.bu
```
> 具体内容修改规则查看官网 [https://coreos.github.io/ignition/examples/]

## 生成 ign 文件
```shell
./butane.sh example.bu > out.ign
```

## 检阅文件内容
```shell
cat out.ign
```