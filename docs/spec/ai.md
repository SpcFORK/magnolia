# API Documentation

_Auto-generated from Magnolia source._

---

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

