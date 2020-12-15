<p align="center">
  <a href="https://github.com/uncleyeung">
   <img alt="Uncle-Yeong-Logo" src="https://raw.githubusercontent.com/uncleyeung/uncleyeung.github.io/master/web/img/logo1.jpg">
  </a>
</p>

<p align="center">
  为简化开发工作、提高生产率而生
</p>

<p align="center">
  
  <a href="https://github.com/996icu/996.ICU/blob/master/LICENSE">
    <img alt="996icu" src="https://img.shields.io/badge/license-NPL%20(The%20996%20Prohibited%20License)-blue.svg">
  </a>

  <a href="https://www.apache.org/licenses/LICENSE-2.0">
    <img alt="code style" src="https://img.shields.io/badge/license-Apache%202-4EB1BA.svg?style=flat-square">
  </a>
</p>

# uncleyeung's Repo For Cydia
> * Source: https://github.com/uncleyeung/uncleyeung.github.io/
> * Twitter: https://twitter.com/uncle_yeung
> * Tumblr: https://www.tumblr.com/blog/uncleyeung
# yeung-go-user-center
* goctl api -o user.api(生成rest api 文件)
* goctl api go -api user.api -dir .(生成rest文件)
* goctl rpc template -o user.proto(生成proto文件)
* goctl rpc proto -src user.proto -dir .(生成rpc文件)
* goctl model mysql ddl -c -src user.sql -dir .(生成db文件根据ddl文件)
* goctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/test" -table="user*" -dir ./model(通配符匹配所有user
开头的表)
* goctl model mysql datasource -url="user:password@tcp(127.0.0.1:3306)/test" -table="user" -dir ./model(指定表名user)
* goctl docker -go main.go(docker)
#### 文档
https://www.yuque.com/tal-tech/go-zero
https://mp.weixin.qq.com/s/VLBiIbZStKhb7uth1ndgQQ