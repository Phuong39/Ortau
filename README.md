#备忘录
最后完成效果是，用于隐藏C2服务器的、开箱即用的反向代理服务器。
省去繁琐的配置Nginx服务的过程。

~~~
隐藏Ortau的特征（伪装Nginx）
注册Linux服务（？），win直接使用exe
运行时询问配置，或直接指定参数，或在配置文件中修改。
反代随机端口
配置文件yaml
过滤器->根据header头过滤或放行响应
~~~