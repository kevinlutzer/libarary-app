# Assumptions

1. The CLI does not have to support i18n. If it did, I would use the `GNU gettext` tool with the `"github.com/gosexy/gettext"` go binding similar to what LXC does.
1. The REST server does not require API authentication
1. The port 8000 is open on the host computer running the server. The server does take an environment variable to specify an additional port. The CLI also will take a flag to specify the hostname which will include the port to accommodate 8000 not being available.
1. There is no auth requirement
1. There is no other security requirements. I didn't worry about protecting against things like cross site request forgery (CSRF) or man in the middle (MIM) attacks. In other projects I have worked on, I have used a CRSF token in the cookie to validate that and SSL with HTTPs to help protect against MIM attacks. 
1. It's not really common to change the name of a book, so the API and CLI will not support it. If it did, I would add it to the `POST /v1/book` request arguments and the CLI's update book command. 
1. There could be multiple books with the same name, author, genre, edition, published date and description.
1. The majority of books will be in many collections. In that case it makes more sense in my opinion to have the collection schema reference the list of books associated with it, rather than the books storing a list of collections they are in. I could create an associative table that maps a book to a specific collection, but I think that would be a little out of scope for this project.
1. Book titles will be relatively brief. The longest book title on file is [3,777 words long](https://www.guinnessworldrecords.com/world-records/358711-longest-title-of-a-book). In my opinion it isn't reasonable to allocate that much space for a single column, so instead I am going to limit the title to 255 characters.

# Implementation Notes

LXD uses what appears to be a distributed form of SQLite. I am going to use SQLite for this project as it is a lightweight database that is easy to use and setup. I am going to use [GORM](https://gorm.io/) as an object relational mapper to interact with the database. It doesn't appear that SQLite supports text arrays, so I am going to use a blob to store the list of books associated with a collection.

Since I am assuming there can be multiple books with identical fields, I will use a UUID as the primary key. This does lessen the user experience when doing operations like delete and update as the user would have to know the UUID of the book they want to manipulate. However, I think that it does give more flexibility to the users as they can have multiple books with the same name, author, genre, edition, and description.

I also choose to use the same HTTP routing multiplexer lxc/icarus uses, [gorilla/mux](github.com/gorilla/mux). The APIs use a CRUD REST style. 