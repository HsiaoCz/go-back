## cobra 命令行框架

cobra 建立在命令、参数和标志这三个结构之上。

- 命令(command)：命令表示要执行的操作
- 参数(arg):是命令的参数，一般表示操作的对象
- 标志(flag):是命令的修饰，可以调整操作的行为

好的命令行遵循的模式

```bash

appname verb noun --adjective

appname command arg --flag

```

verb 代表动词，noun 代表名词，adjective 代表形容词

例如

```bash
hugo server --port=1314
```

server 是一个子命令，port 是一个标志,1314 是标志的参数

### 1、快速开始

```bash
go get -u github.com/spf13/cobra/cobra
```

