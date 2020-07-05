```shell
docker ps # 運行中的容器。
docker ps -a #  列出運行中和曾經運行中的容器。
docker  ps -s #  查詢每個容器佔用多到記憶體，可以搭配  -a 使用。

docker start <container id> # 把已經停止的容器再次啟動。
docker stop <container id> # 停止容器。
docker restart <container id> # 已經停止的容器會直接啟動，如果是已經啟動狀態的容器則是會先停止再做啟動。
docker  rm <container id> # 移除已經停止的容器。
docker inspect <container id> # 查詢容器配置。
docker stats # 查詢每個容器的 cpu和記憶體消耗，可以搭配  -a 使用。
docker stats <container id> # 查詢容器的 cpu和記憶體消耗。
docker logs <container id> # 查詢容器內的 log，-f可以查詢即時 log。

docker exec -it <container id> /bin/bash # 進入容器並且使用 bash

docker images # 列出目前全部 image。
docker rmi <image id> # 移除 image。
docker inspect <image id> # 查詢 image配置。

docker build -t <iamge alias> . # 編譯目前路徑的 dockerfile產生 image。
docker build -t <iamge alias> <dockerfile path> # 編譯指定路徑的 dockerfile產生 image，範例路徑：./2-docker/ 。
docker run -d -p 8002:8080 <iamge alias> # 啟動容器，-d可以背景啟動，-p指定 port。
```
# 容器生命週期
ref: https://zhuanlan.zhihu.com/p/98783855

![容器生命週期](1.jpg)

## 圆形 代表容器的五种状态：
* created：初建状态
* running：运行状态
* stopped：停止状态
* paused： 暂停状态
* deleted：删除状态

## 长方形 代表容器在执行某种命令后进入的状态：
* docker create ： 创建容器后，不立即启动运行，容器进入初建状态；
* docker run ： 创建容器，并立即启动运行，进入运行状态；
* docker start ： 容器转为运行状态；
* docker stop ： 容器将转入停止状态；
* docker kill ： 容器在故障（死机）时，执行kill（断电），容器转入停止状态，这种操作容易丢失数据，除非必要，否则不建议使用；
* docker restart ： 重启容器，容器转入运行状态；
* docker pause ： 容器进入暂停状态；
* docker unpause ： 取消暂停状态，容器进入运行状态；
* docker rm ： 删除容器，容器转入删除状态（如果没有保存相应的数据库，则状态不可见）。

## 菱形 需要根据实际情况选择的操作
* killed by out-of-memory（因内存不足被终止）
* 宿主机内存被耗尽，也被称为OOM：非计划终止
* 这时需要杀死最吃内存的容器
* 然后进行选择操作
* container process exited（异常终止）
* 出现容器被终止后，将进入Should restart?选择操作：
* yes 需要重启，容器执行start命令，转为运行状态。
* no 不需要重启，容器转为停止状态。