# tcp 粘包

## 原因

主要原因是 tcp 数据传递模式是流模式，在保持长连接的时候可以进行多次的收、发。

“粘包”在服务端和客户端都可能发生

+ 由 Nagle 算法造成的发送端的粘包：Nagle 算法是一种改善网络传输效率的算法。当提交一段数据给tcp发送时，tcp 并不立刻发送此段数据，而是等待一小段时间，看看在此期间还有没有需要发送的数据。如果有，会一次性把多段数据发送出去

+ 接收端未及时接收造成的粘包：tcp 会把接收到的数据存放在自己的缓冲区中，然后通知应用层取数据。当应用层由于某些原因不能及时把tcp数据取出来，就会造成tcp缓冲区中存放了多段数据。

## 解决办法

“粘包”的关键在于接收方不确定将要传输的数据包的大小，因此我们可以对数据包进行封包和拆包的操作。

+ 封包：给一段数据加上包头，这样一来数据包就分为*包头*和*包体*两部分内容了（过滤非法包时，封包会加入*包尾*内容）。包头部分的长度是固定的，并且它存储了包体的长度。

+ 拆包：根据固定长度获取包头，再根据包头存储的值获取包体。
