
## Book

For the book entity, we will use the following schema:

# Assumptions: 

1. There could be multiple books with the same name, author, genre, edition, and description. In that case the publishing date would differentiate them. To ensure that we could uniquely identify a book I am using a UUID. This does lessen the user experience when doing operations like delete and update as the user would have to know the UUID of the book they want manipulate. However, I think that it does give more flexibility to the users as they can have multiple books with the same name, author, genre, edition, and description.

1. The majority of books will be in many collections. In that case it makes more sense in my opinion to have the collection schema reference the list of books associated with it, rather then the books storing a list of collections they are in. I could create an associative table that maps a book to a specific collection, but I think that would be a little out of scope for this project.

1. Book titles will be relatively brief. The longest book title on file is [3,777 words long](https://www.guinnessworldrecords.com/world-records/358711-longest-title-of-a-book). In my opinion it isn't reasonable to allocate that much space for a single column, so instead I am going to limit the title to 255 characters.
