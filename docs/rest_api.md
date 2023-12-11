# REST API Server
<!-- 
This should include the documentation of the REST API, all supported methods and what they’ll
be doing, examples of input and output data and what query parameters will be supported and
what they’ll do. See for example the actual REST API documentation of LXD. -->

## Status codes

The REST API will return one of the following status codes:

Code  | Meaning
:---  | :------
200   | Success
400   | Bad Request
404   | Not found
405   | Method Not Allowed
409   | Already exists
412   | Precondition failed
500   | Internal error

If a non 200 status code is returned no database insert or updates will have been preformed. 

## Return values

There are two different return values for the REST API. A success return and an error return. The response for both will be a JSON structure. For example, a successful response will look like:

```json
{
    "type": "success",
    "result": {}
}
```

Result can be any valid JSON type including strings, arrays and objects. The `result` value type is dependent on the endpoint being called and the HTTP method used.  

An error response will look like: 

```json
{
    "type": "error",
    "result": {
      "msg": "error message",
      "errorType": "NotFound"
    }
}
```

The status code returned from the API on error will correspond with the `errorType` field in the response. The values for `errorType` are an ENUM and include:

```js
"AlreadyExists" // 409
"NotFound" // 404
"Internal" // 500
"InvalidArguments" // 400
"PreconditionFailed" // 412
"MethodNotAllow" // 405
```

Error message can be any string, and will provide details about what went wrong.

## Filtering

Only the `GET /v1/book` API supports filtering. The filtering is contained withing the query parameters of the request. For this specific API the following query parameters are supported:

- ids
- author
- genre"
- rangeStart
- rangeEnd

Note that for the request to be valid the query paramters must be 'escaped'. An example of a request that would be successful is:

```bash
curl -X GET "http://localhost:8000/v1/book?author=John%20Doe&genre=Fantasy&&rangeEnd=2016-02-1" -H "accept: application/json"
```

## APIs 

### GET /v1/book

Returns a list of books. The list can be filtered by the query parameters listed above.

#### Example request

```bash
curl -X GET "http://localhost:8000/v1/book" -H "accept: application/json"
```

