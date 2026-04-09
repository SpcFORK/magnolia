# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ai.oak`

### `dotProduct(v1, v2)`

### `magnitude(v)`

### `normalize(v)`

> returns `?`

### `cosineSimilarity(v1, v2)`

> returns `?`

### `euclideanDistance(v1, v2)`

### `manhattanDistance(v1, v2)`

### `minMaxScale(data)`

> returns `:object`

### `zScore(data)`

> returns `:object`

### `softmax(logits)`

> returns `:list`

### `scaledDotProductAttention(query, key, value, scale)`

### `relu(x)`

### `leakyRelu(x, alpha)`

### `sigmoid(x)`

> returns `:int`

### `tanh(x)`

### `gelu(x)`

### `crossEntropyLoss(predicted, actual)`

### `meanSquaredError(predicted, actual)`

> returns `:int`

### `tokenizeSimple(text)`

> returns `:list`

### `tokenCount(text)`

### `ModelConfig(params)`

> returns `:object`

### `topK(items, k)`

### `kosineSimilaritySearch(embeddings, queryEmbedding, k)`

### `batchify(items, batchSize)`

### `padSequence(sequence, targetLength, padToken)`

### `mean(data)`

> returns `:int`

### `variance(data)`

> returns `:int`

### `stddev(data)`

### `entropy(probabilities)`

### `pbatchNormalize(vectors)`

### `pbatchCosineSimilarity(vectors, query)`

### `matZeros(rows, cols)`

### `matIdentity(n)`

### `matTranspose(mat)`

> returns `:list`

### `matMul(a, b)`

### `matAdd(a, b)`

### `matScale(mat, scalar)`

### `vecAdd(v1, v2)`

### `vecSub(v1, v2)`

### `vecScale(v, scalar)`

### `accuracy(predicted, actual)`

> returns `:int`

### `precisionScore(predicted, actual, positiveLabel)`

### `recallScore(predicted, actual, positiveLabel)`

### `f1Score(predicted, actual, positiveLabel)`

> returns `:int`

### `confusionMatrix(predicted, actual, numClasses)`

### `oneHotEncode(label, numClasses)`

### `oneHotDecode(v)`

> returns `?`

### `shuffleIndices(n, seed)`

### `trainTestSplit(data, labels, ratio, seed)`

> returns `:object`

### `kFoldIndices(n, k)`

### `knn(data, labels, query, k)`

### `positionalEncoding(seqLen, dModel)`

### `layerNorm(v, gamma, beta)`

### `lrConstant(baseLr, _step)`

### `lrStepDecay(baseLr, step, stepSize, decayFactor)`

### `lrExponentialDecay(baseLr, step, decayRate)`

### `lrCosineAnnealing(baseLr, step, totalSteps, minLr)`

### `lrWarmupLinear(baseLr, step, warmupSteps, totalSteps)`

### `initUniform(n, lo, hi, seed)`

### `initXavier(n, fanIn, fanOut, seed)`

### `initHe(n, fanIn, seed)`

### `sgdStep(params, grads, lr)`

### `sgdMomentumStep(params, grads, velocity, lr, momentum)`

### `clipGradNorm(grads, maxNorm)`

### `argmax(v)`

### `temperatureScale(logits, temperature)`

### `topPFilter(logits, p)`

### `embeddingLookup(table, indices)`

### `pmatMul(a, b)`

### `pbatchSoftmax(vectors)`

### `pbatchMeanSquaredError(predictedBatch, actualBatch)`

