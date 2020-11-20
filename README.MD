## StatusCake

This is an uptime monitoring service written in golang.

### Database Schema

```
ID                              int
Url     			string
FailureThreshold 		int
FailureCount			int
Status				string
```


### APIs

The following are the API supported by the service.

#### /urls/
* `GET` : Get the list of all urls present in the database.

* `POST` : Create a new entry in the database.

#### /urls/:id
* `GET` : Check the health of the url
* `PUT` : Update a url
* `DELETE` : Delete a url

### Post Params
```
{
    "url": "http://httpbin.org/status/200",
    "failure_count": 0,
    "failure_threshold": 50
}
```
