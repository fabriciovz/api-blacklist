# Black List Generator 

To generate the black list, black list without exclude items, and just exclude items

![alt text](https://github.com/fabriciovz/api-walmart/blob/master/bl_postman_photo.png?raw=true)

## Dependencies

Run locally an instance of MySQL

```bash
docker run --name mysql8 \
-e MYSQL_ROOT_PASSWORD=root -e MYSQL_USER=homestead -e MYSQL_DATABASE=bl_db -e MYSQL_PASSWORD=secret \
-p 3306:3306 mysql:8.0.18 --default-authentication-plugin=mysql_native_password
``` 

Then you have to access to mysql

```bash
docker exec -it mysql8 mysql -u homestead -p
```

Then select bl_db
```bash
use bl_db
```

Create the black_list table
```bash
DROP TABLE IF EXISTS `black_list`;
CREATE TABLE `black_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sku` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2400 DEFAULT CHARSET=utf8;
```

And create the black_list_exclude table
```bash
DROP TABLE IF EXISTS `black_list_exclude`;
CREATE TABLE `black_list_exclude` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `sku` int(11) DEFAULT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2401 DEFAULT CHARSET=utf8;
```

Load the black list from bl_list.sql file 

## Info

1.- Create Environment Variables:

- **MYSQL56_USER** = homestead
- **MYSQL56_PASSWORD** = secret
- **MYSQL56_HOST** = "127.0.0.1" (local)
- **MYSQL56_PORT** = 3306

2.- Run locally api-blacklist

3.- Remember you have a postman collection into the root folder called: bl_generator.json