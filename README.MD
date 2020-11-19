## API


#### /urls/
* `GET` : Get all urls
* `POST` : Create a new url

#### /urls/:id
* `GET` : Check the health of the url
* `PUT` : Update a url
* `DELETE` : Delete a url

#Post Params
```
{
    "url": "http://httpbin.org/status/200",
    "frequency": 0,
    "failure_threshold": 50
}
```
