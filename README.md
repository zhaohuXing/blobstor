## 图片、视频、文档存储 (BlobStor)

模块梳理

- 登录注册
- 个人管理
- 创建文件夹
- 上传文件
- 下载文件
- 图片的在线查看
- 视频的在线播放
- 分享
- 断点续传（亮点）
- 秒传 (亮点)


#### 登录注册

数据库设计:

用户表
```
create table users (
	id bigint not null auto_increment,
	phone varchar(11) not null,
	password varchar(255) not null,
	nickname varchar(16) not null,
	role varchar(8) defalut 'user',
	create_time timestamp not null,
	primary key(id)
)engine = innodb default charset = utf8 auto_increment=1000;
```

文件表
```
create table file_index (
	id bigint not null auto_increment,
	user_id bigint not null,
	filepath varchar(255) not null,
	primary key(id)
)engine = innodb default charset = utf8;
```

