# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ai-vec.oak`

### `dotProduct(v1, v2)`

### `magnitude(v)`

### `normalize(v)`

> returns `?`

### `cosineSimilarity(v1, v2)`

> returns `?`

### `euclideanDistance(v1, v2)`

### `manhattanDistance(v1, v2)`

### `chebyshevDistance(v1, v2)`

> returns `:int`

### `hammingDistance(v1, v2)`

### `vecAdd(v1, v2)`

### `vecSub(v1, v2)`

### `vecScale(v, scalar)`

### `minMaxScale(data)`

> returns `:object`

### `zScore(data)`

> returns `:object`

### `mean(data)`

> returns `:int`

### `variance(data)`

> returns `:int`

### `stddev(data)`

### `entropy(probabilities)`

### `ema(data, alpha)`

> returns `:list`

### `pbatchNormalize(vectors)`

### `pbatchCosineSimilarity(vectors, query)`

### `cosineSimilarityMatrix(vectors)`

### `distanceMatrix(vectors, distFn)`

