#!/bin/bash
# 环境 release-生产环境 debug-开发环境 test-测试环境
export env=release
# 应用域名
export app_domain=
# 数据库连接地址
export mysql_host=
# 数据库商品
export mysql_port=3306
# 数据库库名
export mysql_database=
# 数据库用户
export mysql_user=
# 数据库密码
export mysql_password=
# 数据表前缀
export mysql_table_prefix=
# 数据库连接池中最大闲置连接数
export mysql_max_idle_conn=2
# 数据库最大连接数量
export mysql_max_open_conn=5
# redis连接地址
export redis_host=127.0.0.1
# redis连接端口
export redis_port=6379
# redis密码
export redis_password=
# redis key前缀
export redis_key_prefix=
# etcd连接地址
export etcd_host=127.0.0.1
# etcd连接端口
export etcd_port=2379
# api token key
export token_key=