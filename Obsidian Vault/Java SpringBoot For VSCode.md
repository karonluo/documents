#编程语言/Java语言 #Java-SpringBoot #T200 #开发工具/IDE/VSCode #WEB/WEB后端

## 描述

## 安装步骤
### 1. 安装 JDK1.11

### 2. 在 VisualStudio Code 安装 Java 扩展包

#### Java Extension for Pack
#### Spring Boot Extension Pack

### 3. 配置 VSCode Java 和 Maven 
~~~JSON
{
    "java.jdt.ls.java.home":"/usr/lib/jvm/jdk-11",
    "maven.terminal.useJavaHome": true
}
~~~

### 3. 建立 MAVEN 项目
~~~Shell
# 远程 Linux SSH
# root 账号
# 项目名称 juwbwebapp
# 创建 项目目录
mkdir -p ~/juwbwebapp/.mvn/wrapper/
# 利用 Maven 生成空白项目
/root/.vscode-server/extensions/vscjava.vscode-maven-0.39.2/resources/maven-wrapper/mvnw org.apache.maven.plugins:maven-archetype-plugin:3.1.2:generate -DarchetypeArtifactId="elm-spring-boot-blank-archetype" -DarchetypeGroupId="am.ik.archetype" -DarchetypeVersion="0.0.3" -DgroupId="com.karonluo" -DartifactId="uwbwebapp" -DoutputDirectory="/root/juwbwebapp"

~~~

