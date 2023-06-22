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

字符串可以使用单引号，也可以使用双引号

```sh
#!/bin/bash
name='SnailClimb'
hello='Hello, I  am '$name'!'
echo $hello
```

双引号

```sh
#!/bin/bash
name='SnailClimb'
hello="Hello, I  am "$name"!"
echo $hello
```

shell 中字符串的常见操作

```sh
#!/bin/bash
name="SnailClimb"

# 使用双引号拼接
greeting="hello, "$name" !"
greeting_1="hello, ${name} !"
echo $greeting  $greeting_1

# 使用单引号拼接
greeting_2='hello, '$name' !'
greeting_3='hello, ${name} !'
echo $greeting_2  $greeting_3

# 获取字符串长度
name="SnailClimb"
# 第一种方式
echo ${#name} #输出 10
# 第二种方式
expr length "$name";

# 有一点需要注意的是
# 使用 expr 命令时，表达式中的运算符左右必须包含空格，
# 如果不包含空格，将会输出表达式本身
expr 5+6    // 直接输出 5+6
expr 5 + 6       // 输出 11

# 还可以使用转义字符
expr 5 * 6       // 输出错误
expr 5 \* 6      // 输出30

# 截取字符串
str="SnailClimb is a great man"
echo ${str:0:10} #输出:SnailClimb

# 根据表达式截取字符串
var="https://www.runoob.com/linux/linux-shell-variable.html"
# %表示删除从后匹配, 最短结果
# %%表示删除从后匹配, 最长匹配结果
# #表示删除从头匹配, 最短结果
# ##表示删除从头匹配, 最长匹配结果
# 注: *为通配符, 意为匹配任意数量的任意字符
s1=${var%%t*} #h
s2=${var%t*}  #https://www.runoob.com/linux/linux-shell-variable.h
s3=${var%%.*} #http://www
s4=${var#*/}  #/www.runoob.com/linux/linux-shell-variable.html
s5=${var##*/} #linux-shell-variable.html
```

### 3、shell 中的数组

bash 中支持一维数组，并不支持多维数组

```sh
array=(1 2 3 4 5);
# 获取数组长度
length=${#array[@]}
# 或者
length2=${#array[*]}
#输出数组长度
echo $length #输出：5
echo $length2 #输出：5
# 输出数组第三个元素
echo ${array[2]} #输出：3
unset array[1]# 删除下标为1的元素也就是删除第二个元素
for i in ${array[@]};do echo $i ;done # 遍历数组，输出：1 3 4 5
unset array; # 删除数组中的所有元素
for i in ${array[@]};do echo $i ;done # 遍历数组，数组元素为空，没有任何输出内容
```

### 4、shell 基本运算符

**1、算术运算符**

```sh
a=3;b=3;
val=`expr $a + $b`
#输出：Total value : 6
echo "Total value : $val"
```

**2、关系运算符**

- -eq 检测是否相等
- -ne 是否不等
- -le 小于等于
- -ge 大于等于
- -lt 小于
- -gt 大于

```sh
score=90;
maxscore=100;
if [ $score -eq $maxscore ]
then
   echo "A"
else
   echo "B"
fi
```

**3、逻辑运算**

- && and
- || or

```sh
a=$(( 1 && 0))
# 输出：0；逻辑与运算只有相与的两边都是1，与的结果才是1；否则与的结果是0
echo $a;
```

**4、布尔运算符**

- ! 非运算
- -o 或运算
- -a 与运算

**5、字符串运算**

- = 检测两个字符串是否相等,是返回 true
- != 检测两个字符串是否相等,不是返回 true
- -z 检测字符串长度是否为 0，为 0 返回 true
- -n 检测字符串长度是否为 0，不为 0 返回 true
- str 检测字符串是否为空，不为空返回 true

```sh
a="abc";
b="efg";
if [ $a = $b ]
then
   echo "a 等于 b"
else
   echo "a 不等于 b"
fi
```

**6、文件相关运算符**

- -b file 检测文件是否为块设备文件，是返回 true
- -c file 字符设备
- -d file 目录文件
- -f file 普通文件 既不是目录，也不是设备
- -g file 是否设置了 SGID 位
- -k file 是否设置了粘着位
- -p file 是否是有名管道
- -u file 是否设置了 SUID 位
- -r file 是否可读
- -w file 是否可写
- -x file 是否可执行
- -s file 是否为空
- -e file 是否存在(包括目录)

```sh
# 比如我们定义好了一个文件路径file="/usr/learnshell/test.sh"
# 如果我们想判断这个文件是否可读
file= "/usr/local/bin/kratos"
if [ -r $file]
```

### 5、shell 流程控制

**if 条件语句**

```sh
a=3;
b=9;
if [ $a -eq $b ]
then
   echo "a 等于 b"
elif [ $a -gt $b ]
then
   echo "a 大于 b"
else
   echo "a 小于 b"
fi

```

shell if 条件语句中不能包含空语句也就是什么都不做的语句。

**for 循环语句**

输出当前列表中的数据

```sh
for loop in 1 2 3 4 5

do

  echo "the value is : $loop"

done
```

产生 10 个随机数

```sh
for i in {0..9};
do
   echo $RANDOM;
done
```

输出 1 到 5

```sh
# 通常情况下 shell 变量调用需要加 $,但是 for 的 (()) 中不需要,下面来看一个例子
for((i=1;i<=5;i++));do
    echo $i;
done;
```

**while 语句**

基本 while 循环语句

```sh
int=1
while(( $int<=5 ))
do
    echo $int
    let "int++"
done
```

while 循环可用于读取键盘信息

```sh
echo '按下 <CTRL-D> 退出'
echo -n '输入你最喜欢的电影: '
while read FILM
do
    echo "是的！$FILM 是一个好电影"
done
```

无限循环

```sh
while true
do
    command
done
```
