# HOME

HOME是一个专注在家庭网络应用设想与可能实现，它去中心，可以构建在手机、树莓派、任何Iot设备，它是构建在广域网中的逻辑私有网络应用群。它加密、不可篡改、不上报任何个人数据。同样，它只是个设想，估计会很费钱。

不知道大家是否有遇到过一个问题，我们的吃穿、住行、工作，甚至睡眠，都牢牢被大公司盯着，盯得紧紧的，让人感觉总有双眼睛盯着你。我不怀疑移动支付等新兴科技带来的便利，可是我怀疑的是我们是否需要把我们的数据交给他们？有没有一种可能，数据在我们自己的家庭服务器上，可以是任何一种设备，它以加密方式存储数据，不限数据库，也可以纯文本存储。

诚然，没有网络也是恐怖的，但是不是必须的，李子柒如果没有网络资本推手的渲染，我想她不会有这么多人知道，可是，没有网络她就不会过得这么好么？我想不是的，没有网络，她依旧可以每天穿着带泥的人字拖在田间采摘，惬意生活。

网络到底是拉近了我们的距离还是捅破了我们的窗户纸？腾讯存储我们每一条聊天信息，阿里京东肆意向合作伙伴分享你的搜索记录，还小米每天在分析我们的睡眠，甚至你啪啪啪的波动他们都躲在背后呵呵呵，...，没有一个大厂没有上报我们的行为数据，**我们都是在大街上裸奔的巨婴，每个巨头呵护着我们的生活**。

所以，我就在想，能不能构建一个网络，它没有中心，只有P2P，聊天、家庭的生活数据都维护在自己的树莓派上。

~~设想一个场景：我去街边的自动柜员机买瓶水，现在的流程是这样的：~~

> ~~打开支付宝->打开扫码->扫柜员机的码->支付宝付款->拿货->支付宝上报购物信息..(后面的已经不重要）~~。

~~那我们能不能这样解决：~~

> ~~打开我们自己的APP->扫码->扫柜员机的码->扣除我们的虚拟币->拿货~~

~~这个过程，柜员机不知道是谁买的，它只需要去虚拟币的地址买取币。我们不去讨论这个虚拟币的产生与消费，但是可行性是有的。~~

也可以说，**HOME**会是一个区块链应用，我不否定这一点。

## 家庭私有云

不同于云部署，我这里指的家庭私有云是自行部署在家里，通过我们的光纤宽带注册到Network上，可能是Go-micro的Network，也可能是其它的注册中心。但是只是注册一些服务的接口，服务逻辑仍然在我们本地，数据本身自然也是在本地。家庭自身的数据是保密的，其它平台是无法获取的，或者我们也可以辅助备份，加密备份，加密的密钥由主人持有，服务端不做任何存储，加密存储也不会做，目的只是保证数据在平台这一侧是无法解密的，自然也就不存在暴露的风险，暴露了也是乱码。（暴力破解可能性存在，但是这是安全方面的事，不是服务本身的事，暴力破解到的只是一个人的，而平台泄漏，则是全部的）。

试想，我们每个家庭都有自己的服务器，它可能是一台电脑，可能是一台单片机，甚至是我们淘汰下来的手机，都可能是我们的私有人中心。它们保护着我们的隐私，我们的自由。

我个人非常反感用户大数据，以其说是大数据，还不如说是大盗窃。

。。。持续构思中。

## 如何对在公网的设备通信

我临时觉得可能有两个思路

提供中央平台登录，这是公网的地址，登录后获取注册上来的设备地址，跳转到这个地址。此时两个思路为：

1. P2P包通信，设备与盒子直接对接通信。
2. 盒子分发域名，解析到盒子分到的公网IP，现在IPv6已经开始流程，广域设备都有一个IP成为趋势，如果只是IPv4，那可以使用花生壳的技术动态解析地址等等。

第一个方案复杂，但是自由，实现难度大。
第二个政策因素影响大。易实现。

## 家庭

### 计划

家庭的生活、工作等计划管理

### 财务

家庭收支管理

### 相册

家庭相册

### 记事

快捷家庭记事本

### 聊天

家庭私有小群聊天，可能可以做成P2P的家庭网络交流圈

## 5G

随着5G的到来，速度越来越快，流量越来越便宜，私有移动服务器的可能性越来越大，也就是说，我们的家庭私有云将打破现有的网络格局，不再依赖宽带，延时越来越小，受小区网络影响可能性越来越少。

## 盒子

可能会用一个树莓派作为家庭应用中心，在上面安装我们的服务。

## 代码目录

### [chat](./chat) 聊天

### [gateway](./gateway) 家庭轻量网关

### [common](./common) 公用包

### [doc](./doc) 文档目录

### [finance](./finance) 账务管理

#### [book](./finance/book) 账本

- 支出
- 收入
- 标签，自定义标签

### [note](./note) 记事本

### [open](./open) 开放平台，用于家庭私有云选择性开发部分能力，比如与其它家庭建立聊天群，共享计划等等

### [photo](./photo) 照片

### [plan](./plan) 家庭工作、生活、旅行等计划管理

### [platform](./platform) 核心平台，家庭网络中央服务节点平台，公有云部署

### [proto](./proto) 接口原型文件

### [web](./web) 私有云桌面端浏览器入口