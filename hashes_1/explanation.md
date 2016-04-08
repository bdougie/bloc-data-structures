## Discoveries in GO

This assignment required me to rebuild a [murmur3](https://en.wikipedia.org/wiki/MurmurHash) hash using arrays. I had some
limitations trying to make my hash interact the way Ruby interacts with
it's hashes. I ended up bailing on that and just creating an insertion
funciton.

From all my research there is no abstraction to make this work:

```
hahs['key'] = 'value'
```

I really did not want to rebuild what is a Go [Map](https://blog.golang.org/go-maps-in-action), but I might do that later. For now I prefer to learn about Algorithms and Binary Trees.
