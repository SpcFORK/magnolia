# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `ai-data`

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

## Module: `ai-decode`

### `argmax(v)`

### `temperatureScale(logits, temperature)`

### `topPFilter(logits, p)`

### `beamSearch(scoreFn, startSeq, beamWidth, maxLen)`

### `greedyDecode(scoreFn, startSeq, maxLen)`

### `topK(items, k)`

### `kosineSimilaritySearch(embeddings, queryEmbedding, k)`

### `embeddingLookup(table, indices)`

### `categoricalSample(probs, seed)`

### `multinomialSample(probs, n, seed)`

### `uniformRandom(n, seed)`

## Module: `ai-linalg`

### `matZeros(rows, cols)`

### `matIdentity(n)`

### `matTranspose(mat)`

> returns `:list`

### `matMul(a, b)`

### `matAdd(a, b)`

### `matScale(mat, scalar)`

### `matDiag(mat)`

### `matFromDiag(diag)`

### `matTrace(mat)`

### `outerProduct(v1, v2)`

### `matHadamard(a, b)`

### `matFlatten(mat)`

### `matReshape(flat, rows, cols)`

### `pmatMul(a, b)`

## Module: `ai-ml`

### `knn(data, labels, query, k)`

### `kMeansStep(data, centroids)`

> returns `:object`

### `kMeans(data, initialCentroids, iterations)`

### `centerData(data)`

> returns `:object`

### `covarianceMatrix(centered)`

### `projectData(data, basis)`

### `decisionStump(featureIdx, threshold, leftLabel, rightLabel)`

> **thunk** returns `:function`

### `majorityVote(predictions)`

### `ensemblePredict(classifiers, sample)`

### `rollingMean(data, windowSize)`

### `rollingStddev(data, windowSize)`

### `differencing(data)`

### `autoCorrelation(data, lag)`

> returns `:int`

### `convolve1d(signal, kernel)`

### `maxPool1d(data, windowSize, stride)`

### `avgPool1d(data, windowSize, stride)`

## Module: `ai-nn`

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

## Module: `ai-optim`

### `dropout(v, rate, seed)`

### `l2Regularization(weights, lambda)`

### `l1Regularization(weights, lambda)`

### `elasticNetRegularization(weights, lambda, alpha)`

### `sgdStep(params, grads, lr)`

### `sgdMomentumStep(params, grads, velocity, lr, momentum)`

### `clipGradNorm(grads, maxNorm)`

### `adamStep(params, grads, m, v, lr, beta1, beta2, t)`

### `rmspropStep(params, grads, cache, lr, decayRate)`

### `lrConstant(baseLr, _step)`

### `lrStepDecay(baseLr, step, stepSize, decayFactor)`

### `lrExponentialDecay(baseLr, step, decayRate)`

### `lrCosineAnnealing(baseLr, step, totalSteps, minLr)`

### `lrWarmupLinear(baseLr, step, warmupSteps, totalSteps)`

### `lrCosineWarmRestarts(baseLr, step, t0, tMult, minLr)`

### `initUniform(n, lo, hi, seed)`

### `initXavier(n, fanIn, fanOut, seed)`

### `initHe(n, fanIn, seed)`

### `numericalGradient(f, x, h)`

### `numericalGradientVec(f, x, h)`

## Module: `ai-text`

### `tokenizeSimple(text)`

> returns `:list`

### `tokenCount(text)`

### `nGrams(tokens, n)`

### `bagOfWords(tokens)`

### `tfidf(tf, df, n)`

## Module: `ai-vec`

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

## Module: `lib\ai.oak`

### `dotProduct(v1, v2)`

### `magnitude(v)`

### `normalize(v)`

### `cosineSimilarity(v1, v2)`

### `euclideanDistance(v1, v2)`

### `manhattanDistance(v1, v2)`

### `chebyshevDistance(v1, v2)`

### `hammingDistance(v1, v2)`

### `vecAdd(v1, v2)`

### `vecSub(v1, v2)`

### `vecScale(v, scalar)`

### `minMaxScale(data)`

### `zScore(data)`

### `mean(data)`

### `variance(data)`

### `stddev(data)`

### `entropy(probabilities)`

### `ema(data, alpha)`

### `pbatchNormalize(vectors)`

### `pbatchCosineSimilarity(vectors, query)`

### `cosineSimilarityMatrix(vectors)`

### `distanceMatrix(vectors, distFn)`

### `matZeros(rows, cols)`

### `matIdentity(n)`

### `matTranspose(mat)`

### `matMul(a, b)`

### `matAdd(a, b)`

### `matScale(mat, scalar)`

### `matDiag(mat)`

### `matFromDiag(diag)`

### `matTrace(mat)`

### `outerProduct(v1, v2)`

### `matHadamard(a, b)`

### `matFlatten(mat)`

### `matReshape(flat, rows, cols)`

### `pmatMul(a, b)`

### `relu(x)`

### `leakyRelu(x, alpha)`

### `sigmoid(x)`

### `tanh(x)`

### `gelu(x)`

### `swish(x)`

### `mish(x)`

### `elu(x, alpha)`

### `selu(x)`

### `softmax(logits)`

### `pbatchSoftmax(vectors)`

### `crossEntropyLoss(predicted, actual)`

### `meanSquaredError(predicted, actual)`

### `meanAbsoluteError(predicted, actual)`

### `huberLoss(predicted, actual, delta)`

### `pbatchMeanSquaredError(predictedBatch, actualBatch)`

### `scaledDotProductAttention(query, key, value, scale)`

### `multiHeadAttention(queries, keys, values, dK)`

### `positionalEncoding(seqLen, dModel)`

### `ropeEmbedding(v, position, base)`

### `layerNorm(v, gamma, beta)`

### `causalMask(seqLen)`

### `paddingMask(sequence, padToken)`

### `applyMask(scores, mask)`

### `dropout(v, rate, seed)`

### `l2Regularization(weights, lambda)`

### `l1Regularization(weights, lambda)`

### `elasticNetRegularization(weights, lambda, alpha)`

### `sgdStep(params, grads, lr)`

### `sgdMomentumStep(params, grads, velocity, lr, momentum)`

### `clipGradNorm(grads, maxNorm)`

### `adamStep(params, grads, m, v, lr, beta1, beta2, t)`

### `rmspropStep(params, grads, cache, lr, decayRate)`

### `lrConstant(baseLr, step)`

### `lrStepDecay(baseLr, step, stepSize, decayFactor)`

### `lrExponentialDecay(baseLr, step, decayRate)`

### `lrCosineAnnealing(baseLr, step, totalSteps, minLr)`

### `lrWarmupLinear(baseLr, step, warmupSteps, totalSteps)`

### `lrCosineWarmRestarts(baseLr, step, t0, tMult, minLr)`

### `initUniform(n, lo, hi, seed)`

### `initXavier(n, fanIn, fanOut, seed)`

### `initHe(n, fanIn, seed)`

### `numericalGradient(f, x, h)`

### `numericalGradientVec(f, x, h)`

### `ModelConfig(params)`

### `oneHotEncode(label, numClasses)`

### `oneHotDecode(v)`

### `shuffleIndices(n, seed)`

### `trainTestSplit(data, labels, ratio, seed)`

### `kFoldIndices(n, k)`

### `batchify(items, batchSize)`

### `padSequence(sequence, targetLength, padToken)`

### `labelEncode(labels)`

### `polynomialFeatures(features, degree)`

### `accuracy(predicted, actual)`

### `precisionScore(predicted, actual, positiveLabel)`

### `recallScore(predicted, actual, positiveLabel)`

### `f1Score(predicted, actual, positiveLabel)`

### `confusionMatrix(predicted, actual, numClasses)`

### `rSquared(actual, predicted)`

### `addNoise(v, scale, seed)`

### `mixup(sample1, sample2, lambda)`

### `quantize(values, numLevels)`

### `dequantize(levels, numLevels)`

### `trainLoop(params, data, batchSize, epochs, updateFn)`

### `tokenizeSimple(text)`

### `tokenCount(text)`

### `nGrams(tokens, n)`

### `bagOfWords(tokens)`

### `tfidf(tf, df, n)`

### `knn(data, labels, query, k)`

### `kMeansStep(data, centroids)`

### `kMeans(data, initialCentroids, iterations)`

### `centerData(data)`

### `covarianceMatrix(centered)`

### `projectData(data, basis)`

### `decisionStump(featureIdx, threshold, leftLabel, rightLabel)`

### `majorityVote(predictions)`

### `ensemblePredict(classifiers, sample)`

### `rollingMean(data, windowSize)`

### `rollingStddev(data, windowSize)`

### `differencing(data)`

### `autoCorrelation(data, lag)`

### `convolve1d(signal, kernel)`

### `maxPool1d(data, windowSize, stride)`

### `avgPool1d(data, windowSize, stride)`

### `argmax(v)`

### `temperatureScale(logits, temperature)`

### `topPFilter(logits, p)`

### `beamSearch(scoreFn, startSeq, beamWidth, maxLen)`

### `greedyDecode(scoreFn, startSeq, maxLen)`

### `topK(items, k)`

### `kosineSimilaritySearch(embeddings, queryEmbedding, k)`

### `embeddingLookup(table, indices)`

### `categoricalSample(probs, seed)`

### `multinomialSample(probs, n, seed)`

### `uniformRandom(n, seed)`

## Module: `math`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

### `orient(x0, y0, x1, y1)`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

### `median(xs)`

### `stddev(xs)`

### `round(n, decimals)`

## Module: `math-base`

- `Pi` · `3.141592653589793`
- `E` · `2.718281828459045`
### `sign(n)`

> returns `:int`

### `abs(n)`

### `sqrt(n)`

## Module: `math-geo`

### `hypot(x0, y0, x1, y1)`

### `scale(x, a, b, c, d)`

### `bearing(x, y, d, t)`

> returns `:list`

### `orient(x0, y0, x1, y1)`

> returns `:int`

## Module: `math-stats`

### `sum(xs...)`

### `prod(xs...)`

### `min(xs...)`

### `max(xs...)`

### `clamp(x, a, b)`

### `mean(xs)`

> returns `?`

### `median(xs)`

> returns `?`

### `stddev(xs)`

### `pbatchMean(datasets)`

### `pbatchStddev(datasets)`

### `round(n, decimals)`

## Module: `sort`

### `sort!(xs, pred)`

### `sort(xs, pred)`

### `_mergeSorted(a, b, pred)`

### `psort(xs, pred)`

## Module: `std`

### `identity(x)`

### `is(x)`

> **thunk** returns `:function`

### `constantly(x)`

> **thunk** returns `:function`

### `_baseIterator(v)`

> returns `:string`

### `_asPredicate(pred)`

> returns `:function`

### `default(x, base)`

- `_nToH` · `'0123456789abcdef'`
### `toHex(n)`

- `_hToN` · `{22 entries}`
### `fromHex(s)`

### `clamp(min, max, n, m)`

> returns `:list`

### `slice(xs, min, max)`

### `clone(x)`

> returns `:string`

### `range(start, end, step)`

> returns `:list`

### `reverse(xs)`

### `map(xs, f)`

### `each(xs, f)`

### `filter(xs, f)`

### `exclude(xs, f)`

### `separate(xs, f)`

### `reduce(xs, seed, f)`

### `flatten(xs)`

### `compact(xs)`

### `some(xs, pred)`

### `every(xs, pred)`

### `append(xs, ys)`

### `join(xs, ys)`

### `zip(xs, ys, zipper)`

### `partition(xs, by)`

### `uniq(xs, pred)`

### `first(xs)`

### `last(xs)`

### `take(xs, n)`

### `takeLast(xs, n)`

### `find(xs, pred)`

### `rfind(xs, pred)`

### `indexOf(xs, x)`

### `rindexOf(xs, x)`

### `contains?(xs, x)`

> returns `:bool`

### `values(obj)`

### `entries(obj)`

### `fromEntries(entries)`

### `merge(os...)`

> returns `?`

### `once(f)`

> **thunk** returns `:function`

### `loop(max, f)`

### `aloop(max, f, done)`

### `serial(xs, f, done)`

### `parallel(xs, f, done)`

### `debounce(duration, firstCall, f)`

> **thunk** returns `:function`

### `stdin()`

### `println(xs...)`

## Module: `str`

### `checkRange(lo, hi)`

> **thunk** returns `:function`

### `upper?(c)`

> returns `:bool`

### `lower?(c)`

> returns `:bool`

### `digit?(c)`

> returns `:bool`

### `space?(c)`

> returns `:bool`

### `letter?(c)`

> returns `:bool`

### `word?(c)`

> returns `:bool`

### `join(strings, joiner)`

> returns `:string`

### `startsWith?(s, prefix)`

### `endsWith?(s, suffix)`

### `_matchesAt?(s, substr, idx)`

> returns `:bool`

### `indexOf(s, substr)`

### `rindexOf(s, substr)`

### `contains?(s, substr)`

### `cut(s, sep)`

> returns `:list`

### `lower(s)`

### `upper(s)`

### `_replaceNonEmpty(s, old, new)`

### `replace(s, old, new)`

### `_splitNonEmpty(s, sep)`

### `split(s, sep)`

### `_extend(pad, n)`

### `padStart(s, n, pad)`

### `padEnd(s, n, pad)`

### `_trimStartSpace(s)`

### `_trimStartNonEmpty(s, prefix)`

### `trimStart(s, prefix)`

### `_trimEndSpace(s)`

### `_trimEndNonEmpty(s, suffix)`

### `trimEnd(s, suffix)`

### `trim(s, part)`

## Module: `thread`

### `spawn(fnToRun, args...)`

### `makeChannel(size)`

### `send(ch, value, callback)`

### `recv(ch, callback)`

### `close(_ch)`

> returns `?`

### `cs Mutex()`

> returns `:object`

### `cs Semaphore(n)`

> returns `:object`

### `cs WaitGroup()`

> returns `:object`

### `cs Future(fnToRun)`

> returns `:object`

### `cs Pool(numWorkers)`

> returns `:object`

### `parallel(fns)`

### `pmap(list, fnToRun)`

### `pmapConcurrent(list, fnToRun, maxConcurrent)`

### `race(fns)`

### `pipeline(input, stages...)`

### `retry(fnToRun, maxAttempts)`

### `debounce(fnToRun, waitTime)`

> **thunk** returns `:function`

### `throttle(fnToRun, waitTime)`

> **thunk** returns `:function`

