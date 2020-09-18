# xp
xp is my name,but the project is same like ansible playbook auto ops

# 介绍

插件接口流程：

* pipeline
    * User 用户管理
    * Host 主机清单
    * Playbook
        * YAML
        * module
    * Plugin
        * start
        * stop
        * status
        * init
    * input
        * host conn check
        * host env
        * yaml module 分析
    * filter
        * 执行各个module
            * 连接
                * ssh
                * docker
                * k8s
                * 网络设备
                * snmp等
            * 执行
                * RPC
                * RESTFULL
                * CLI
        * 管理执行的生命周期
            * prepare
            * before
            * runtime
            * after
    * output
        * 输出结果
        * 返回状态
