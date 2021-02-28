#!/bin/bash
# 环境 release-生产环境 debug-开发环境 test-测试环境
export env=release
# 应用域名
export app_domain=
# 数据库类型
export db_connection=mysql
# 数据库连接地址
export db_host=
# 数据库商品
export db_port=3306
# 数据库库名
export db_database=
# 数据库用户
export db_user=
# 数据库密码
export db_password=
# 数据表前缀
export db_table_prefix=
# 数据库连接池中最大闲置连接数
export db_max_idle_conn=200
# 数据库最大连接数量
export db_max_open_conn=50
# 数据库连接空闲超时时间(秒)
export db_conn_max_lifetime=300
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
# api token secret
export token_secret=