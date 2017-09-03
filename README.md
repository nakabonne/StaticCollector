## Overview

Application to analyze static files of competing sites.
You can do the following.

- Confirm the change in ranking of competing websites
- Compare two HTML

![result](https://github.com/ryonakao/StaticCollector/blob/media_for_demo/media/NCA_demo.gif)

## Usage

Start mongoDB

```
$ sudo mongod --dbpath /var/lib/mongodb --logpath /var/log/moodb.log
```

Start Mysql

```
$ mysql.server restart
```

仮StaticFiles挿入

```
> db.static_files.insert({word_id:1, page_id:8, title:'仮タイトル', html:"<html></html>", rank:2, target_day:ISODate("2017-08-24T04:54:00.697Z")});
```