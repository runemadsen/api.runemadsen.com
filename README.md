api.runemadsen.com
==================

This is the codebase for `api.runemadsen.com`, a hypermedia API that returns data about me, Rune Madsen. I was tired of updating my DB model every time I wanted to update my website, so I made this.... and made it public.

In a hypermedia'ish way, the API is explanatory. Do a `GET` on `api.runemadsen.com`, and it will list all resources available for unauthenticated users. If you have an authentication token (which you don't, because there's only one and that's for me), it will also list requests you can make while authenticated.


Schema
------

The API tries to conform to the [HAL specification](http://stateless.co/hal_specification.html) as much as possible. This means....

### _links

Every resource has a `_links` object that holds URI's for resource links. Here's a very simple (and fake) example of what that might look like.

```json
{
  "_links" : {
    "self" : { "href" : "/animals" },
    "next" : { "href" : "/animals?page=2" },
    "show" : { "href" : "/animals/{id}", "templated" : true }
  }
}
```

Each link object has at least a `href` attribute that can be a simple URI or a [URI template](http://tools.ietf.org/html/rfc6570), the latter always represented with the `templated` property set to `true`.

### _embedded

If the resource comes preloaded with any of the above `_links`, those resources can be found in the `_embedded` object. This is especially important when loading a resource collection.

```json
{
  "_links" : {
    "self" : { "href" : "/animals" },
    "next" : { "href" : "/animals?page=2" },
    "show" : { "href" : "/animals/{id}" }
  },
  "_embedded" : {
    "animals" : [
      {
        "name" : "lion",
        "type" : "cat",
        "_links": { "self": { "href": "/animals/lion" }}
      },
      {
        "name" : "tiger",
        "type" : "cat",
        "_links": { "self": { "href": "/animals/tiger" }}
      }
    ]
  }
}
```

### Pagination

You can paginate, yeah!


Endpoints
---------

### Articles

Status: **Not Implemented**

The writings endpoint returns shorter and longer essays written by me.

```bash
GET /articles
```

### Portfolio

Status: **Not Implemented**

The portfolio endpoint returns work done by me.

```bash
GET /portfolio
```

### Courses

Status: **Not Implemented**

The courses endpoint returns syllabi and lectures from classes taught by me.

```bash
GET /courses
```

### Presentations

Status: **Not Implemented**

The presentations endpoint returns slides from presentations given by me.

```bash
GET /presentations
```

### Bio

Status: **Not Implemented**

The bio endpoint returns biographical info, like recent professional experience and education.

```bash
GET /bio
```

### Info

Status: **Not Implemented**

The info endpoint returns contact info.

```bash
GET /bio
```

### Photos

Status: **Not Implemented**

```bash
GET /photos
```

### Activities

Status: **Not Implemented**

The activities endpoint returns recent real world activities, like Foursquare check-ins.

```bash
GET /activities
```

### Code

Status: **Not Implemented**

The code endpoint returns programming activities, like statistics around programming languages and Git commits.

```bash
GET /code
```

### Geo

Status: **Not Implemented**

The geo endpoint returns openpaths data

```bash
GET /geo
```

Local Development
-----------------

If you want to try run this api on a local machine, you can use vagrant do spin up a development box:

```bash
$ gem install berkshelf
$ vagrant plugin install vagrant-berkshelf
$ vagrant plugin install vagrant-omnibus
$ vagrant up
```

After chef has installed all the goodies, you can SSH into the dev box and run the api:

```bash
$ vagrant ssh
$ cd /vagrant
$ goreman start
```

You now have the api running on `localhost:3001` and rethinkdb running on `localhost:8081` on your host machine.

Tests
-----

Run the test by doing this:

```bash
$ cd test
$ go test
```

TODO
----

- application+json/hal as middleware https://github.com/codegangsta/martini#middleware-handlers