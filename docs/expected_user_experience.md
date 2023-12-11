<!-- This should include all CLI commands and their options, think of this as writing the help or
manual page of your software ahead of time, including documented examples of user
interactions. -->

# Expected User Experience

The CLI supports managing and reading two different resources: `book` and `collection`. Book represents the information describing a book. This includes title, author, genre, published date, and description. Collection represents a collection of books. This includes the name of the collection and a list of ID references to the books associated with that collection. 

## CLI Command Structure

There are 4 different commands to interact with the two different resources available: `create`, `update`, `delete`, and `get`. To use any of these commands with a resource use the syntax: 

`bookapp command resource [arguments] [options]` 

For example to get all books run: `bookapp get book`. Similarily, to get all collections you would use: `bookapp get collection`. For all commands the option `--host` is supported to specify an alternative host then the default `localhost:8000`.

### Create A Resource

`bookapp create resource name`

This command would be used to create either a book or a book collection. With either command, the user would need to enter a name as the only required field. Additional optional fields would be provided to allow the user to enter more information about the resource. This command will return the ID of the resource that was created.

#### Creating a Book

For example to create a book, use: `bookapp create book "Hello World"`. The additional options would be `--author`, `--genre`, `--published`, `--edition`, and `--description`. Below is a description of those options.

``` bash
    --author string        The author of the book
    --description string   A brief description of the book
    --edition uint8        The edition of the book (default 1)
    --genre string         The genre of the book, valid genres are: science, history, philosophy, art, cooking, fantasy, tragedy
    -h, --help             help for book
    --id string            The id of the book, if not provided a new id will be generated
    --published string     The data the book was published (default "1970-01-01")
    --host string          The hostname of the server to connect to, this must include the port
```

Some additional examples for creating books would be: 

```bash
bookapp create book "The Great Gatsby" --id="b87c1185-e7d7-4fe9-90d2-69c234909c1b" --author="F. Scott Fitzgerald" --genre="tragedy" --published="1925-04-10" --edition=1 --description="The Great Gatsby is a 1925 novel written by American author F. Scott Fitzgerald that follows a cast of characters living in the fictional towns of West Egg and East Egg on prosperous Long Island in the summer of 1922. Many literary critics consider The Great Gatsby to be one of the greatest novels ever written."
bookapp create book "Fellowship of the Ring" --id="d81049f5-0018-441b-9d9a-601c75b0be99" --author="J.R.R. Tolkien" --genre="fantasy" --published="1954-07-29" --edition=1 --description="The Fellowship of the Ring is the first of three volumes of the epic novel The Lord of the Rings by the English author J. R. R. Tolkien. It is followed by The Two Towers and The Return of the King. It takes place in the fictional universe of Middle-earth. It was originally published on 29 July 1954 in the United Kingdom."
bookapp create book "The Two Towers" --id="417d8cd0-5250-467e-a4d1-cba1b3c5ecde" --author="J.R.R. Tolkien" --genre="fantasy" --published="1954-11-11" --edition=1 --description="The Two Towers is the second volume of J. R. R. Tolkien's high fantasy novel The Lord of the Rings. It is preceded by The Fellowship of the Ring and followed by The Return of the King."
bookapp create book "The Return of the King" --id="7321749e-001e-4f69-a7cf-89fe53dfd9e2" --author="J.R.R. Tolkien" --genre="fantasy" --published="1955-10-20" --edition=1 --description="The Return of the King is the third and final volume of J. R. R. Tolkien's The Lord of the Rings, following The Fellowship of the Ring and The Two Towers. The story begins in the kingdom of Gondor, which is soon to be attacked by the Dark Lord Sauron."
bookapp create book "The Hobbit" --id="b659067a-2714-4711-bb8a-88a2779d06ac" --author="J.R.R. Tolkien" --genre="fantasy" --published="1937-09-21" --edition=1 --description="The Hobbit, or There and Back Again is a children's fantasy novel by English author J. R. R. Tolkien. It was published on 21 September 1937 to wide critical acclaim, being nominated for the Carnegie Medal and awarded a prize from the New York Herald Tribune for best juvenile fiction."
```

#### Collection

For example to create a book collection run: `bookapp create collection "Hello World"`. For this command there is an additional argument provided to specify the ID of a book to associate with the collection. Below is a description of the additional options.

``` bash
    --bookid stringArray   The id of a book to add to the collection, this can be specified multiple times
    -h, --help             help for collection
    --host string          The hostname of the server to connect to, this must include the port
```

Some additional examples for creating a collection of books: 

``` bash
bookapp create collection "The Lord of the Rings" --bookid="d81049f5-0018-441b-9d9a-601c75b0be99" --bookid="417d8cd0-5250-467e-a4d1-cba1b3c5ecde" --bookid="7321749e-001e-4f69-a7cf-89fe53dfd9e2" --bookid="b659067a-2714-4711-bb8a-88a2779d06ac"
bookapp create collection "Classics" --bookid="b87c1185-e7d7-4fe9-90d2-69c234909c1b"
```

### Delete A Resource

Once a resource is deleted it can no longer be modified or read through the get command. This command can be used to delete a single item within a resource type. The command can be used like:

`bookapp delete resource id`

There are no additional operations to this command. Bellow are additional options for deleting a book or a book collection. 

``` bash
-h, --help          help for delete
    --host string   The hostname of the server to connect to, this must include the port (default "localhost:8080")
```

#### Book

To delete a book, run `bookapp delete book "994dda1e-3c2e-42cf-84f0-b4383b34d8b8"`

#### Collection

To delete a collection, run `bookapp delete collection "7aec852b-d0ba-4a89-a04d-a8f41685ffa2"`

### Get a Resource

This command can be used to load a resource from the REST server. Note that resources that have been deleted cannot be read from this command anymore. The syntax for this command is:

`bookapp get resource`

#### Book

To get books run `bookapp get book`. Additional options can be provided to filter the results, but only for books. Those options are `--id`, `--author`, `--genre`, `--publishedStart`, and `--publishedend`. Note that the `publishedstart` and `publishedend` arguments are inclusive. All of these additional arguments are optional. Below is a description of the additional options for this command: 

``` bash
    --id stringArray       The id(s) of the books to get can be specified multiple times
    --author string        The author of the book
    --genre string         The genre of the book, valid genres are: science, history, philosophy, art, cooking, fantasy, tragedy
    -h, --help             help for book
    --publishedend string  The data the book was published (default "1970-01-01")
    --publishedstart string
                           The data the book was published (default "1970-01-01")
    --host string          The hostname of the server to connect to, this must include the port (default "localhost:8080")
```

Some examples of this command would be: 

``` bash 
bookapp get book --genre="fantasy" # get all fantasy books
bookapp get book --author="J.R.R. Tolkien" # get all books by J.R.R. Tolkien
bookapp get book --publishedstart="1954-07-29" --publishedend="1955-10-20" # get all books published between 1954-07-29 and 1955-10-20
```

#### Collection

To get collections run `bookapp get collection`. You can specify `--includebooks=true` to nest the books associated with the collection under the collection details Below is a description of the additional options for this command:

``` bash
-h, --help           help for collection
    --host string    The hostname of the server to connect to, this must include the port (default "localhost:8080")
    --includebooks   Include the books in the collection
```

Some examples of this command would be: 

``` bash
bookapp get collection 
bookapp get collection --includebooks=true # include the books in the collection under the collection details
```

### Update

This command can be used to update a resource. Note that once a resource has been deleted it cannot be updated anymore The syntax for this command is:

`bookapp update resource id`

#### Book

To update a book run `bookapp update book id`. All fields on the book can be updated except for the name and ID. To update the name for a specific book you would have to delete that book and then creating a new one with the desired name. The first argument specified must be the book ID. The additional options for this command are similar to the create book command: 

``` bash
    --author string        The author of the book
    --description string   A brief description of the book
    --edition uint8        The edition of the book (default 1)
    --genre string         The genre of the book, valid genres are: science, history, philosophy, art, cooking, fantasy, tragedy
    -h, --help             help for book
    --published string     The data the book was published (default "1970-01-01")
    --host string          The hostname of the server to connect to, this must include the port (default "localhost:8080")
```

Some examples of this command would be: 

``` bash
bookapp create book "b87c1185-e7d7-4fe9-90d2-69c234909c1b" --genre="art" --edition=2 # change the edition and genre of the Great Gatsby 
bookapp create book "417d8cd0-5250-467e-a4d1-cba1b3c5ecde" --published="2021-01-2" --edition=1 --description="This is a really long book" # change the description and published date of the Two Towers
```

#### Collection

To update a book run `bookapp update collection id`. The only fields that can be modified on the collection are the name and the list of books associated with the collection. The first argument specified must be the collection ID. The additional options for this command are similar to the create collection command: 

``` bash
      --bookid stringArray   The id of a book to add to the collection, this can be specified multiple times
  -h, --help                 help for collection
      --host string          The hostname of the server to connect to, this must include the port (default "localhost:8080")
      --name string          The name of the collection
```

Some examples of this command would be: 

``` bash
bookapp update collection "7aec852b-d0ba-4a89-a04d-a8f41685ffa2" --name="Older Classics" # change the name of the Classics collection
bookapp update collection "7aec852b-d0ba-4a89-a04d-a8f41685ffa2" --bookid="b87c1185-e7d7-4fe9-90d2-69c234909c1b" # add the Great Gatsby to the collection
```