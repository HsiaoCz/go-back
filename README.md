## go语言之-再回首

### 1、go workspace

```bash
# 新建一个工作目录
# 一个工作目录可以放整个的项目
mkdir workspace

# 在workspace里面建两个项目 
cd workspace

mkdir hello

cd hello

go mod init

cd ../

# 再建一个项目
mkdir hello-world

cd hello-world

go mod init
```
现在我们在工作目录下使用go work init ./hello
然后将hello-world也加入进来 go work use ./hello-world