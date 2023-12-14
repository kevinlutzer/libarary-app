IDs for both schemas are a UUID v4 string. I also created indices on the columns that will be filtered on within gorm.

## Book Schema

``` sql
CREATE TABLE `book` ( 
    `id` varchar(36), 
    `title` varchar(512),
    `author` varchar(512),
    `description` varchar(4096),
    `published_at` datetime,
    `genre` varchar(128),
    `edition` integer,
    `deleted` numeric DEFAULT false,
    PRIMARY KEY (`id`)
)
```

## Collection Schema

Note that the book_ids column is a blob that stores a comma separated list of book IDs. It doesn't appear that SQLite supports text arrays.

``` sql
CREATE TABLE `collection` (
    `id` varchar(36),
    `name` varchar(512),
    `book_ids` text,
    `deleted` numeric DEFAULT false,
    PRIMARY KEY (`id`)
)
```