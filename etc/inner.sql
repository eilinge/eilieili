use eilieili;

-- select * from account_content;
-- select a.content_hash,weight,a.title,b.token_id from content a, account_content b where a.content_hash = b.content_hash and address = "0xc8357fd9e82aa6366853d57e36156918eddb2929";
select a.content_hash,weight,a.title,b.token_id from content a, account_content b where a.content_hash = b.content_hash;
-- create table `group` 
-- (
-- 	id int(10) unsigned not null primary key auto_increment,
--     name varchar(256)
-- )ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4;

-- create table `user`
-- (
-- 	id int(10) unsigned not null primary key auto_increment,
--     name varchar(256),
--     groupid int(10),
--     Index(groupid)
-- )ENGINE = InnoDB AUTO_INCREMENT = 2 DEFAULT CHARSET = utf8mb4;

-- insert into `group`(name) values("A");
-- insert into `group`(name) values("B");

-- insert into `user`(name,groupid) values("eilinge", 1);
-- insert into `user`(name,groupid) values("lin", 2);
-- select * from `group`;