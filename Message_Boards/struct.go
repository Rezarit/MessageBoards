package main

//CREATE TABLE `users` (
//`id` INT AUTO_INCREMENT PRIMARY KEY,    -- 用户的唯一标识
//`nickname` VARCHAR(255)                   -- 用户名
//`username` VARCHAR(255) NOT NULL UNIQUE,    -- 账号，确保唯一
//`password` VARCHAR(255) NOT NULL,        -- 用户密码
//);

//CREATE TABLE `messages` (
//`id` INT AUTO_INCREMENT PRIMARY KEY,    -- 留言的唯一标识
//`user_id` INT NOT NULL,                 -- 留言的用户ID，外键引用 `users` 表的 `id`
//`content` TEXT NOT NULL,                -- 留言内容
//);

type Users struct {
	Id       int    `json:"id"`
	Nickname string `json:"nickname"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type Messages struct {
	Id      int    `json:"id"`
	UserId  int    `json:"user_id"`
	Content string `json:"content"`
}
