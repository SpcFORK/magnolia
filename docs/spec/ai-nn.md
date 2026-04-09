# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ai-nn.oak`

### `relu(x)`

### `leakyRelu(x, alpha)`

### `sigmoid(x)`

> returns `:int`

### `tanh(x)`

### `gelu(x)`

### `swish(x)`

### `mish(x)`

### `elu(x, alpha)`

### `selu(x)`

### `softmax(logits)`

> returns `:list`

### `pbatchSoftmax(vectors)`

### `crossEntropyLoss(predicted, actual)`

### `meanSquaredError(predicted, actual)`

> returns `:int`

### `meanAbsoluteError(predicted, actual)`

> returns `:int`

### `huberLoss(predicted, actual, delta)`

> returns `:int`

### `pbatchMeanSquaredError(predictedBatch, actualBatch)`

### `scaledDotProductAttention(query, key, value, scale)`

### `multiHeadAttention(queries, keys, values, dK)`

### `positionalEncoding(seqLen, dModel)`

### `ropeEmbedding(v, position, base)`

### `layerNorm(v, gamma, beta)`

### `causalMask(seqLen)`

### `paddingMask(sequence, padToken)`

### `applyMask(scores, mask)`

