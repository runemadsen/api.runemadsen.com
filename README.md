api.runemadsen.com
==================

This is the codebase for `api.runemadsen.com`, a hypermedia API that returns data about me, Rune Madsen. I was tired of updating my DB model everytime I wanted to update my webpage, and also just liked the idea of exposing an API with very personal data.

In a hypermedia'ish way, the API is explanatory. Do a `GET` on `api.runemadsen.com`, and it will list all resources available for unauthenticated users. If you have an authentication token (which you don't, because there's only one and that's for me), it will also list requests you can make while authenticated.

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