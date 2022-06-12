### 第五周作业
* 参考 Hystrix 实现一个滑动窗口计数器

### 实现思路
* 定义一个放置请求结果的对象Bucket
* 定义一个滑动窗口对象Rolling
  * 核心属性是一个实现了FIFO方法的Bucket切片
  * Bucket切片的规定大小的上限