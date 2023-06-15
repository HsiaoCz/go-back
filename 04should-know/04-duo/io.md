## linux I/O多路复用机制

### 1、阻塞与非阻塞

**流**

- 可以进行I/O操作的内核对象
- 文件、管道、套接字
- 流的入口：文件描述符(fd)

**I/O操作**

对流的读或者写，从流中读数据，也可以从另一端写如数据
流是有方向的

**什么时候会发生阻塞**

如果流中没有内容，读就会发生阻塞

举个例子：比如你要洗袜子，需要肥皂，但是刚好没有
你就洗不了袜子，你又比较死脑筋，那么你就只能等着
这中情况就是阻塞

**非阻塞，忙轮询**

还是那个例子，你在网上买了块肥皂，你不停的打电话让快递给你把肥皂送过来
这种情况下虽然也没有实现任务，但是处于一种忙碌的状态

**阻塞等待**

阻塞的时候，不占用CPU的时间片

非阻塞，忙轮询：占用CPU,系统资源

阻塞等待的缺点：如果同一时刻到达多个资源，只能处理一个，而别的就处理不了

阻塞等待不适合处理并发消息

多路I/O:既能够阻塞等待，不浪费资源，也能够同一时刻监听多个I/O请求的状态