## shell 编程

### 1、helloWorld

```bash
# 1、新建一个helloworld.sh文件

touch helloworld.sh

# 2、给它增加一个执行权限

chmod +x helloworld.sh

# 3、编辑helloworld.sh文件，增加

echo "helloworld"

# 4、执行

./hellowrold.sh
```

### 2、shell 当中的变量

- 我们自己定义的变量（自定义变量）: 仅在当前 Shell 实例中有效，其他 Shell 启动的程序不能访问局部变量。

- Linux 已定义的环境变量（环境变量， 例如：PATH, ​HOME 等..., 这类变量我们可以直接使用），使用 env 命令可以查看所有的环境变量，而 set 命令既可以查看环境变量也可以查看自定义变量。

- Shell 变量：Shell 变量是由 Shell 程序设置的特殊变量。Shell 变量中有一部分是环境变量，有一部分是局部变量，这些变量保证了 Shell 的正常运行

**常见的环境变量**

- PATH 决定了 shell 将到哪些目录中寻找命令或程序
- HOME 当前用户主目录
- HISTSIZE 　历史记录数
- LOGNAME 当前用户的登录名
- HOSTNAME 　指主机的名称
- SHELL 当前用户 Shell 类型
- LANGUAGE 　语言相关的环境变量，多语言可以修改此环境变量
- MAIL 　当前用户的邮件存放目录
- PS1 　基本提示符，对于 root 用户是#，对于普通用户是$

使用 linux 已定义的环境变量

```sh
# hellowrold.sh
# 自定义变量
hello="helloworld"

echo $hello

echo "HelloWorld"
```

shell 里面的变量名使用的注意事项：

- 命名只能使用英文字母，数字和下划线，首个字符不能以数字开头，但是可以使用下划线（\_）开头。
- 中间不能有空格，可以使用下划线（\_）。
- 不能使用标点符号。
- 不能使用 bash 里的关键字（可用 help 命令查看保留关键字）

### 2、shell 的字符串
