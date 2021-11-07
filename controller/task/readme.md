# 任务模块初步设计

## 核心：taskid映射的任务表（可拓展）

**标记在源码内，以const形式存在，可修正** 

| taskid | mission                           |
| ------ | --------------------------------- |
| 1      | selection                         |
| 2      | healing                           |
| 3      | moment                            |
| 4      | easteregg1(准备一个彩蛋任务,可选) |
| 5      | easteregg2(准备两个彩蛋任务,可选) |

## 使用

在对应controller下读取userid,后调用task模块的对象使用方法,实现接口约定的方法

## Mysql设计
**源码见models文件夹 task.go / task_table.go**

### 任务记录表

**每24h更新**

**counter记录任务完成部分，taskid标记任务库**

| 额外部分   | userid | taskid | counter |
| ---------- | ------ | ------ | ------- |
| gorm.model | int    | int    | int     |

### 任务库

**可扩展，需要在task模块实现对应的接口并约定好相应的任务序号** 

**text标记任务文本，max标记积分上限**

| 额外部分   | text   | max  |
| ---------- | ------ | ---- |
| gorm.Model | string | int  |

## Redis设计

**24h定时更新**

| key           | f1     | f2     | ...(根据用户情形拓展) |
| ------------- | ------ | ------ | --------------------- |
| {userid}/task | taskid | taskid |                       |



