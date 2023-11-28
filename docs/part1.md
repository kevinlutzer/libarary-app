# Describe the expected user experience

<!-- This should include all CLI commands and their options, think of this as writing the help or
manual page of your software ahead of time, including documented examples of user
interactions. -->


## Assumptions: 
1. The CLI does not have to support i18n. If it did I would use the `GNU gettext` tool with the `"github.com/gosexy/gettext"` go binding similar to what LXC does.

```
bookapp create book name --published= --edition= --author= --description= --genre=
bookapp delete book id
bookapp update book --books 
bookapp create collection name --books=[id1,id2, id3]
bookapp get book --author= --genre= --published= --edition=
```

``` bash
go run cli/main.go create book "The Great Gatsby" --author="F. Scott Fitzgerald" --genre="tragedy" --published="1925-04-10" --edition=1 --description="The Great Gatsby is a 1925 novel written by American author F. Scott Fitzgerald that follows a cast of characters living in the fictional towns of West Egg and East Egg on prosperous Long Island in the summer of 1922. Many literary critics consider The Great Gatsby to be one of the greatest novels ever written."
go run cli/main.go create book "Fellowship of the Ring" --author="J.R.R. Tolkien" --genre="fantasy" --published="1954-07-29" --edition=1 --description="The Fellowship of the Ring is the first of three volumes of the epic novel The Lord of the Rings by the English author J. R. R. Tolkien. It is followed by The Two Towers and The Return of the King. It takes place in the fictional universe of Middle-earth. It was originally published on 29 July 1954 in the United Kingdom."
go run cli/main.go create book "The Two Towers" --author="J.R.R. Tolkien" --genre="fantasy" --published="1954-11-11" --edition=1 --description="The Two Towers is the second volume of J. R. R. Tolkien's high fantasy novel The Lord of the Rings. It is preceded by The Fellowship of the Ring and followed by The Return of the King."
go run cli/main.go create book "The Return of the King" --author="J.R.R. Tolkien" --genre="fantasy" --published="1955-10-20" --edition=1 --description="The Return of the King is the third and final volume of J. R. R. Tolkien's The Lord of the Rings, following The Fellowship of the Ring and The Two Towers. The story begins in the kingdom of Gondor, which is soon to be attacked by the Dark Lord Sauron."
go run cli/main.go create book "The Hobbit" --author="J.R.R. Tolkien" --genre="fantasy" --published="1937-09-21" --edition=1 --description="The Hobbit, or There and Back Again is a children's fantasy novel by English author J. R. R. Tolkien. It was published on 21 September 1937 to wide critical acclaim, being nominated for the Carnegie Medal and awarded a prize from the New York Herald Tribune for best juvenile fiction."

go run cli/main.go create collection "The Lord of the Rings" --bookid="994dda1e-3c2e-42cf-84f0-b4383b34d8b8" --bookid="7aec852b-d0ba-4a89-a04d-a8f41685ffa2" --bookid="78c93c13-e0e4-40d4-b67b-5458fba7da0b"

```


