#如果复制到mysql中执行时 加上
#DELIMITER ;;
drop database `doog_library`
create database `doog_library`;;
use `doog_library`;;

drop table if exists `sys_user`;;
drop table if exists `sys_book`;;

##user
create table `sys_user`(
`user_id` bigint not null auto_increment,
`username` varchar(100),
constraint `pk_sys_user` primary key(`user_id`)
) charset=utf8 ENGINE=InnoDB;;
alter table `sys_user` auto_increment=100000;;

##book
create table `sys_book`(
`book_id` bigint not null auto_increment,
`title` varchar(100) not null,
`isbn10` char(10),
`isbn13` char(13),
constraint `pk_sys_book` primary key(`book_id`)
) charset=utf8 ENGINE=InnoDB;;
alter table `sys_book` auto_increment=100000;;

##collection
create table `sys_collection`(
`book_id` bigint not null,
`user_id` bigint not null,
`state` varchar(10) not null,
constraint `pk_sys_collection` primary key(`book_id`, `user_id`),
constraint `fk_sys_collection_sys_user_user_id` foreign key (`user_id`) references `sys_user` (`user_id`),
constraint `fk_sys_collection_sys_book_book_id` foreign key (`book_id`) references `sys_book` (`book_id`)
) charset=utf8 ENGINE=InnoDB;;
alter table `sys_collection` auto_increment=100000;;