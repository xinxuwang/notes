## RST介绍
* 用来异常关闭链接
* 发送RST关闭链接时，不必等待缓冲区的包全部发出，直接丢弃缓冲区的包，发送RST
* 接收RST包后，也不必发送ACK确认
## 发送RST的情况
* 端口没有打开，server没有打开端口，client链接时，server直接发送RST包，不同的系统也可能不理会，直接丢包
* 请求超时，client设置了SO_RCVTIMEO想要在短时间接收到响应，但是超时了，导致client直接发送RST
* 读缓冲区里还有数据时，提前关闭socket，导致发送RST
* 在已经关闭的socket上还收到了数据，发送RST
* 发送窗口一直都是0，超过了定时器，也会发送RST