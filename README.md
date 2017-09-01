## Overview

It's a web crawler to scrape Google search results

## Usage

mongo起動

```
$ sudo mongod --dbpath /var/lib/mongodb --logpath /var/log/moodb.log
```

Mysql起動

```
$ mysql.server restart
```

仮StaticFiles挿入

```
> db.static_files.insert({word_id:1, page_id:8, title:'引越しの見積もりは引越しの専業【サカイ引越センター】公式サイト', html:"<html></html>", rank:2, target_day:ISODate("2017-08-24T04:54:00.697Z")});
```