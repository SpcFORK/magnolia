# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ai-data.oak`

### `ModelConfig(params)`

### `oneHotEncode(label, numClasses)`

### `oneHotDecode(v)`

> returns `?`

### `shuffleIndices(n, seed)`

### `trainTestSplit(data, labels, ratio, seed)`

> returns `:object`

### `kFoldIndices(n, k)`

### `batchify(items, batchSize)`

### `padSequence(sequence, targetLength, padToken)`

### `labelEncode(labels)`

### `polynomialFeatures(features, degree)`

### `accuracy(predicted, actual)`

> returns `:int`

### `precisionScore(predicted, actual, positiveLabel)`

### `recallScore(predicted, actual, positiveLabel)`

### `f1Score(predicted, actual, positiveLabel)`

> returns `:int`

### `confusionMatrix(predicted, actual, numClasses)`

### `rSquared(actual, predicted)`

> returns `:int`

### `addNoise(v, scale, seed)`

### `mixup(sample1, sample2, lambda)`

### `quantize(values, numLevels)`

### `dequantize(levels, numLevels)`

### `trainLoop(params, data, batchSize, epochs, updateFn)`

