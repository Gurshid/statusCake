## API

#### /urls/
* `GET` : Get all urls
* `POST` : Create a new url

#### /urls/:id
* `GET` : Get a url
* `PUT` : Update a url
* `DELETE` : Delete a url

#Post Params
```
{
    "url": "http://httpbin.org/status/200",
    "crawl_timeout": 0,
    "failure_threshold": 50
}
```
