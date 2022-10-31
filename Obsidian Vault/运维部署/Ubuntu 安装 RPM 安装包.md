# Ubuntu 安装 RPM 安装包
## 描述
有些安装包基于 RPM 而非 DEB 因此需要现在转换成 DEB 安装包

## 步骤

> Step 1. 安装 alien

~~~Shell
sudo apt update
sudo apt -y install alien
~~~

>Step 2. 转换 RPM 至 DEB

~~~Shell
sudo alien software.rpm
~~~

>Step 3. 安装 DEB

~~~Shell
sudo dpkg -i software.deb
sudo apt -y install ./software.deb
~~~

