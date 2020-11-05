# Agent(Golang)

构建 Agent 是指包含了 agent 进程监控和调度部分逻辑的代码，不包含与流水线交互的构建类业务逻辑代码，需要与另外一个 worker(kotlin) 一起整合才能是完整的 Agent 包。

## Agent 二进制程序编译

在根目录下分别有 3 个操作系统的编译脚本：

- build_linux.sh
- build_macos.sh
- build_windows.bat

只需要直接执行即可，比如 Linux 包将会在 bin 目录下生成对应 devopsDaemon_linux,devopsAgent_linux ，其他系统依此类推。

- devopsDaemon： 用于守护 agent 进程，监控和拉起 agent 进程
- devopsAgent: 用于和调度服务通信，以及拉起构建进程 worker

## Agent 控制脚本

举例 Linux, 其他系统依此类推。

- scripts/linux/install.sh： agent 安装脚本
- scripts/linux/start.sh： agent 启动脚本
- scripts/linux/stop.sh： agent 停止脚本
- scripts/linux/uninstall.sh： agent 卸载脚本
