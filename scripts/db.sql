CREATE DATABASE IF NOT EXISTS ck_blog DEFAULT CHARACTER SET utf8mb4 DEFAULT COLLATE utf8mb4_general_ci;

use ck_blog;

CREATE TABLE blog_tag (
                          id int(10) unsigned NOT NULL AUTO_INCREMENT primary key ,
                          name varchar(100) DEFAULT '' COMMENT 'the name of tag',
                          state tinyint(3) unsigned DEFAULT 1 COMMENT 'state: 0-forbidden 1-normal'
)ENGINE = InnoDb DEFAULT CHARSET = utf8mb4 COMMENT = 'tag manager';

CREATE TABLE blog_article (
                              id int(10) unsigned NOT NULL PRIMARY KEY auto_increment,
                              title varchar(100) DEFAULT '' COMMENT 'the name of article',
                              doc varchar(255) DEFAULT '' COMMENT 'the describe of article',
                              cover_image_url varchar(255) DEFAULT '' COMMENT 'the image of home page',
                              content longtext COMMENT 'the content of article',
                              state tinyint(3) DEFAULT 1 COMMENT 'state: 0-forbidden 1-normal'
)ENGINE = InnoDb DEFAULT CHARSET = utf8mb4 COMMENT = 'article manager';


CREATE TABLE blog_article_tag (
                                  id int(10) unsigned NOT NULL AUTO_INCREMENT primary key ,
                                  article_id int(11) NOT NULL COMMENT 'the id of article',
                                  tag_id int(10) unsigned DEFAULT 0 COMMENT 'the id of tag'
)ENGINE = InnoDb DEFAULT CHARSET = utf8mb4 COMMENT = 'the relation between tag and article';


