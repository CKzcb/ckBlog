# RESTFUL API

- Get: read or search
- Post: add new and create
- Put: upload a whole resource, must be idempotence
- Patch: update, part of resource
- Delete: del

*idempotence: Methods can also have the property of "idempotence" in that (aside from error or expiration issues) the
side-effects of N > 0 identical requests is the same as for a single request.*

### design

#### tag

| func | method | path |
| --- |:---:| --- |
| new tag | POST | /tags|
| del tag | DELETE | /tags/:id |
| update tag | PUT | /tags/:id |
| get tags | GET | /tags |

#### article

| func                 | method | path          |
|----------------------|:---:|---------------|
| new article          | POST | /articles     |
| del articles         | DELETE | /articles/:id |
| update articles      | PUT | /articles/:id |
| get specified articles | GET | /articles/:id |
| get articles         | GET | /articles |




