api.runemadsen.com
==================

This is the codebase for `api.runemadsen.com`, a hypermedia API that returns data about me, Rune Madsen. I was tired of updating my DB model everytime I wanted to update my webpage, and also just liked the idea of exposing an API with very personal data.

In a hypermedia'ish way, the API is explanatory. Do a `GET` on `api.runemadsen.com`, and it will list all resources available for unauthenticated users. If you have an authentication token (which you don't, because there's only one and that's for me), it will also list requests you can make while authenticated.

TODO
----

application/hal+json

Schema
------

The API tries to conform to the [HAL specification](http://stateless.co/hal_specification.html) as much as possible. This means....

### _links

Every resource has a `_links` object that holds URI's that holds resource links. This is a very simple (and fake) example of what that might look like.

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


Resources
---------

### Works

```bash
GET /works
```

### Experiences

```bash
GET /experiences
```

### Writings

```bash
GET /writings
```

### Teachings

```bash
GET /teachings
```

### Info

```bash
GET /info
```

### Photos

```bash
GET /photos
```



Random Ideas
------------

Foursquare data?

GitHub data?