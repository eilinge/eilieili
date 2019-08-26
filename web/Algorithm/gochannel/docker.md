# Dockerfile

## web daemon

    https://www.runoob.com/docker/docker-install-nginx.html
    https://blog.csdn.net/tsummerb/article/details/79248015#commentBox
    https://blog.csdn.net/harleylau/article/details/80150375

## params

    FROM  FROM debian:stretch表示以debian:stretch作为基础镜像进行构建

    RUN  可以看出RUN后面跟的其实就是一些shell命令，通过&&将这些脚本连接在了一行执行，这么做的原因是为了减少镜像的层数，每多一行RUN都会给镜像增加一层，所以这里选择将所有命令联结在一起执行以减少层数

    ARG  特地将这个指令放在RUN之后讲解，这个指令可以进行一些宏定义，比如我定义ENV JAVA_HOME=/opt/jdk，之后RUN后面的shell命令中的${JAVA_HOME}都会被/opt/jdk代替

    ENV  可以看出这个指令的作用是在shell中设置一些环境变量（其实就是export）

    FROM…AS…  这是Docker 17.05及以上版本新出来的指令，其实就是给这个阶段的镜像起个别名：FROM ...(基础镜像) AS ...(别名)，在后面引用这个阶段的镜像时直接使用别名就可以了

    COPY  顾名思义，就是用来来回复制文件的，COPY . /root/workspace/agent表示将当前文件夹（.表示当前文件夹，即Dockerfile所在文件夹）的所以文件拷贝到容器的/root/workspace/agent文件夹中。通过--from参数也可以从前面阶段的镜像中拷贝文件过来，比如--from=builder表示文件来源不是本地文件系统，而是之前的别名为builder的容器

    WORKDIR  在执行RUN后面的shell命令前会先cd进WORKDIR后面的目录

    ENTRYPOINT  这个参数表示镜像的“入口”，镜像打包完成之后，使用docker run命令运行这个镜像时，其实就是执行这个ENTRYPOINT后面的可执行文件（一般是一个shell脚本文件），也可以通过["可执行文件", "参数1", "参数2"]这种方式来赋予可执行文件的执行参数，这个“入口”执行的工作目录也是WORKDIR后面的那个目录
