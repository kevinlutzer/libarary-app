

# Database and Server Repository Notes

LXD uses what appears to be a distributed form of SQLite. I am going to use SQLite for this project as it is a lightweight database that is easy to use and setup. I am going to use [GORM](https://gorm.io/) as an object relational mapper to interact with the database. It doesn't appear that SQLite supports text arrays, so I am going to use a blob to store the list of books associated with a collection. There will be further notes on this in the assumptions section.

# Assumptions: 

1. There could be multiple books with the same name, author, genre, edition, published date and description. To ensure that we could uniquely identify a book I am using a UUID as the primary key. This does lessen the user experience when doing operations like delete and update as the user would have to know the UUID of the book they want to manipulate. However, I think that it does give more flexibility to the users as they can have multiple books with the same name, author, genre, edition, and description.

1. The majority of books will be in many collections. In that case it makes more sense in my opinion to have the collection schema reference the list of books associated with it, rather than the books storing a list of collections they are in. I could create an associative table that maps a book to a specific collection, but I think that would be a little out of scope for this project.

1. Book titles will be relatively brief. The longest book title on file is [3,777 words long](https://www.guinnessworldrecords.com/world-records/358711-longest-title-of-a-book). In my opinion it isn't reasonable to allocate that much space for a single column, so instead I am going to limit the title to 255 characters.

# Book Schema

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

# Collection Schema

``` sql
CREATE TABLE `collection` (
    `id` varchar(36),
    `name` varchar(512),
    `book_ids` text,
    `deleted` numeric DEFAULT false,
    PRIMARY KEY (`id`)
)
```