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
