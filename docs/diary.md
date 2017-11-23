## 编码日志

### 2017-12-11

- 完成 Register 接口
- 调整了日志输出形式
- 完成 Login 接口 

```
create table users (
	id bigint not null auto_increment,
	phone varchar(11) not null,
	password varchar(255) not null,
	nickname varchar(16) not null,
	role varchar(8) default 'user',
	create_time timestamp not null,
	primary key(id)
)engine = innodb default charset = utf8 auto_increment=1000;
```
