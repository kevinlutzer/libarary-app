# Describe the expected user experience

<!-- This should include all CLI commands and their options, think of this as writing the help or
manual page of your software ahead of time, including documented examples of user
interactions. -->

## Assumptions: 
1. The CLI does not have to support i18n. If it did, I would use the `GNU gettext` tool with the `"github.com/gosexy/gettext"` go binding similar to what LXC does.
1. The REST server does not require API authentication
1. The server is going to be served on localhost and port 8000. Connections will be done with the http protocal. If this was not the case additional parameters can be provided to the CLI to specify additional host information.

## Resources

The CLI supports managing and reading two different resources: `book` and `collection`. Book represents the information describing a book. This includes title, author, genre, published date, and description. Collection represents a collection of books. This includes the name of the collection and a list of ID references to the books associated with that collection. 

## CLI Commands

There are 4 different commands to interact with the two different resources available: create, update, delete, and get. 

The CLI uses nested commands to allow the user to create, update, delete or get resources. The two valid resources that the CLI and REST server support are book and collections. Book represents a singular book and collection represents books and book collections. The two avaliable resources are book and collection 

### Create A Resource

This command with the syntax of: 

`bookapp create resource name`

Would be used to create either a book or a book collection. With either command, the user would need to enter a name as the only required field. Additional optional fields would be provided to allow the user to enter more information about the resource. This command will return the ID of the resource that was created.

#### Book

For `bookapp create book "Hello World"` the additional options would be `--author`, `--genre`, `--published`, `--edition`, and `--description`. Bellow is a description of those options

``` bash
    --author string        The author of the book
    --description string   A brief description of the book
    --edition uint8        The edition of the book (default 1)
    --genre string         The genre of the book, valid genres are: science, history, philosophy, art, cooking, fantasy, tragedy
    -h, --help             help for book
    --id string            The id of the book, if not provided a new id will be generated
    --published string     The data the book was published (default "1970-01-01")
```

An example of creating a book would be: `bookapp create book "The Great Gatsby" --id="b87c1185-e7d7-4fe9-90d2-69c234909c1b" --author="F. Scott Fitzgerald" --genre="tragedy" --published="1925-04-10" --edition=1 --description="The Great Gatsby is a 1925 novel written by American author F. Scott Fitzgerald that follows a cast of characters living in the fictional towns of West Egg and East Egg on prosperous Long Island in the summer of 1922. Many literary critics consider The Great Gatsby to be one of the greatest novels ever written."`

#### Collection

The create command can also be use collections as a resource. To create a book collection run: `bookapp create collection "Hello World"`. For this command the only additional operation provided is `--bookid` which is used to specify the ID of a book to add to the collection. This option can be used multiple times to add multiple books to the collection. 

``` bash
    --bookid stringArray   The id of a book to add to the collection, this can be specified multiple times
    -h, --help             help for collection
```

An example of creating a collection of books would be: `bookapp create collection "Classics" --bookid="b87c1185-e7d7-4fe9-90d2-69c234909c1b"`

### Delete A Resource

Once a resource is deleted it can no longer be modified or read through the get command. This command can be used to delete a single item within a resource. The command can be used like:

`bookapp delete resource id`

There are no additional operations to this command. It deletes only the resource with the specified ID. 

#### Book

To delete a book, run `bookapp delete book "994dda1e-3c2e-42cf-84f0-b4383b34d8b8"`

#### Collection

To delete a collection, run `bookapp delete collection "7aec852b-d0ba-4a89-a04d-a8f41685ffa2"`

### Get a Resource

This command can be used to load a resource from the REST server. Note that resources that have been deleted cannot be read from this command anymore. The syntax for this command is:

`bookapp get resource`

#### Book

Additional options can be provided to filter the results, but only for books. Those options are `--author`, `--genre`, `--publishedStart`, `--publishedend`. Note that the `publishedstart` and `publishedend` arguments are inclusive. All of these additional arguments are optional 

``` bash
    --id stringArray       The id(s) of the books to get can be specified multiple times
    --author string        The author of the book
    --genre string         The genre of the book, valid genres are: science, history, philosophy, art, cooking, fantasy, tragedy
    -h, --help             help for book
    --publishedend string  The data the book was published (default "1970-01-01")
    --publishedstart string
                           The data the book was published (default "1970-01-01")
```

#### Collection

To list collections run `bookapp get collection`. You can specify `--includebooks=true` to list the books associated with the collection

### Update

This command can be used to update a resource. The syntax for this command is:

`bookapp update resource id`

#### Book

All fields on the book can be updated except for name. To update the name for a specific book I would recommend deleting that book and then creating a new one with the desired name. The first argument specified must be the book argument specified  The additional options for this command are similar: 

``` bash
    --author string        The author of the book
    --description string   A brief description of the book
    --edition uint8        The edition of the book (default 1)
    --genre string         The genre of the book, valid genres are: science, history, philosophy, art, cooking, fantasy, tragedy
    -h, --help             help for book
    --published string     The data the book was published (default "1970-01-01")
```

With this command 

```
bookapp create resource name 
bookapp delete book id
bookapp update book --books 
bookapp create collection name --books=[id1,id2, id3]
bookapp get book --author= --genre= --published= --edition=
```