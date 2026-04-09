# AI Library: 2026 Standards

The `ai` library provides modern machine learning and artificial intelligence utilities for Magnolia, designed according to 2026 standards. It includes vector operations, neural network primitives, embedding utilities, attention mechanisms, loss functions, and data processing tools.

## Quick Start

```js
ai := import('ai')

// Compute cosine similarity between embeddings
sim := ai.cosineSimilarity([1, 0, 0], [0.9, 0.1, 0])

// Normalize an embedding
normalized := ai.normalize([3, 4])  // => [0.6, 0.8]

// Apply softmax for attention
probs := ai.softmax([1, 2, 3])

// Tokenize and count tokens
tokens := ai.tokenizeSimple('Hello world AI')
count := ai.tokenCount('Hello world AI')

// Create model configuration
config := ai.ModelConfig({
    name: 'my-model'
    hiddenSize: 4096
    numLayers: 48
})
```

## Table of Contents

- [Vector Operations & Embeddings](#vector-operations--embeddings)
- [Normalization & Standardization](#normalization--standardization)
- [Attention & Transformer Operations](#attention--transformer-operations)
- [Neural Network Activations](#neural-network-activations)
- [Loss Functions](#loss-functions)
- [Tokenization & Text Processing](#tokenization--text-processing)
- [Embeddings & Retrieval](#embeddings--retrieval)
- [Batch Processing](#batch-processing)
- [Statistical Utilities](#statistical-utilities)
- [Model Configuration](#model-configuration)

## Vector Operations & Embeddings

### dotProduct(v1, v2)

Computes the dot product of two vectors. Essential for similarity computation and matrix operations.

```js
ai.dotProduct([1, 2, 3], [4, 5, 6])  // => 32 (1*4 + 2*5 + 3*6)
ai.dotProduct([1, 0], [0, 1])         // => 0 (orthogonal vectors)
```

### magnitude(v)

Computes the Euclidean (L2) norm of a vector. Used to measure vector length.

```js
ai.magnitude([3, 4])    // => 5
ai.magnitude([0, 0])    // => 0
```

### normalize(v)

Returns a unit vector (magnitude 1) in the same direction as the input. Essential for embedding normalization.

```js
ai.normalize([3, 4])    // => [0.6, 0.8]
ai.normalize([0, 0])    // => ? (undefined for zero vector)
```

### cosineSimilarity(v1, v2)

Computes cosine similarity between two vectors. Returns a value between -1 and 1, where 1 indicates identical direction. This is the **standard similarity metric for embeddings** in modern AI applications.

```js
ai.cosineSimilarity([1, 0, 0], [1, 0, 0])  // => 1.0 (identical)
ai.cosineSimilarity([1, 0], [0, 1])        // => 0.0 (orthogonal)
ai.cosineSimilarity([1, 0], [-1, 0])       // => -1.0 (opposite)
```

**Use case**: Finding similar embeddings for semantic search, RAG, and nearest-neighbor retrieval.

### euclideanDistance(v1, v2)

Computes the L2 distance between two vectors. Used for distance-based clustering and search.

```js
ai.euclideanDistance([0, 0], [3, 4])  // => 5
ai.euclideanDistance([1, 2], [4, 6])  // => 5
```

### manhattanDistance(v1, v2)

Computes the L1 (Manhattan) distance between two vectors. Faster than Euclidean distance, used in some optimization contexts.

```js
ai.manhattanDistance([0, 0], [3, 4])  // => 7 (|3| + |4|)
```

## Normalization & Standardization

### minMaxScale(data)

Normalizes values to the range [0, 1] using min-max scaling. Returns an object with `min`, `max`, and `data` fields.

```js
result := ai.minMaxScale([1, 2, 3, 4, 5])
// => { min: 1, max: 5, data: [0, 0.25, 0.5, 0.75, 1] }
```

**Use case**: Feature normalization before feeding data to neural networks.

### zScore(data)

Standardizes values using z-score normalization (standard normal distribution). Returns `mean`, `stddev`, and normalized `data`.

```js
result := ai.zScore([1, 2, 3, 4, 5])
// => { mean: 3, stddev: 1.414, data: [...] }
```

**Use case**: Standardizing features to have zero mean and unit variance, common in machine learning pipelines.

## Attention & Transformer Operations

### softmax(logits)

Applies the softmax function to convert logits to probabilities. Essential for attention mechanisms and output layers.

```js
ai.softmax([1, 2, 3])   // => [0.09, 0.24, 0.67] (approximately)
ai.softmax([0, 0, 0])   // => [0.333, 0.333, 0.333]
```

**Use case**: Converting attention scores to probability weights in transformer models.

### scaledDotProductAttention(query, key, value, scale)

Computes scaled dot-product attention, a key mechanism in transformer architectures.

```js
attention := ai.scaledDotProductAttention(
    [1, 0],     // query
    [1, 0],     // key
    [5, 10],    // value
    0.5         // scale factor (typically 1/sqrt(d_k))
)
```

**Use case**: Core attention computation in transformer-based models.

## Neural Network Activations

### relu(x)

Rectified Linear Unit: `max(0, x)`. Most common activation in modern neural networks.

```js
ai.relu(-5)   // => 0
ai.relu(3)    // => 3
ai.relu(0)    // => 0
```

### leakyRelu(x, alpha)

Leaky ReLU with small negative slope for negative inputs. Helps prevent dying ReLU problem.

```js
ai.leakyRelu(-5, 0.01)  // => -0.05
ai.leakyRelu(3, 0.01)   // => 3
```

### sigmoid(x)

Sigmoid activation: `1 / (1 + e^(-x))`. Maps inputs to (0, 1), used for binary classification.

```js
ai.sigmoid(0)    // => 0.5
ai.sigmoid(2)    // => 0.88 (approximately)
ai.sigmoid(-2)   // => 0.12 (approximately)
```

### tanh(x)

Hyperbolic tangent activation. Maps inputs to (-1, 1). Often preferred over sigmoid for hidden layers.

```js
ai.tanh(0)     // => 0
ai.tanh(0.5)   // => 0.46 (approximately)
```

### gelu(x)

Gaussian Error Linear Unit. Standard activation in modern models (BERT, GPT, etc.).

```js
ai.gelu(0)    // => 0
ai.gelu(0.5)  // => 0.345 (approximately)
ai.gelu(2)    // => 1.95 (approximately)
```

**2026 Standard**: GELU and SiLU have largely replaced ReLU as the preferred activations in state-of-the-art models.

## Loss Functions

### crossEntropyLoss(predicted, actual)

Computes binary cross-entropy loss. Measures distance between predicted probability and true label (0 or 1).

```js
ai.crossEntropyLoss(0.9, 1)  // Small error, good prediction
ai.crossEntropyLoss(0.1, 1)  // Large error, bad prediction
```

**Use case**: Classification tasks, training with discrete labels.

### meanSquaredError(predicted, actual)

Computes MSE loss. Measures average squared difference between predictions and targets.

```js
ai.meanSquaredError([1, 2, 3], [1, 2, 3])      // => 0 (perfect)
ai.meanSquaredError([1, 2, 3], [1.1, 2.2, 2.9])  // => 0.0466
```

**Use case**: Regression tasks, continuous value prediction.

## Tokenization & Text Processing

### tokenizeSimple(text)

Basic whitespace tokenization. Splits text on spaces and returns tokens.

```js
ai.tokenizeSimple('hello world AI model')
// => ['hello', 'world', 'AI', 'model']
```

**Note**: For production use with LLMs, use specialized tokenizers (BPE, WordPiece, SentencePiece).

### tokenCount(text)

Estimates token count using a heuristic: `wordCount * 1.3`. Useful for quick token budgeting.

```js
ai.tokenCount('hello world AI')     // => 4 (approximately)
ai.tokenCount('The quick brown fox') // => 5 (approximately)
```

**Use case**: Estimating token usage for API calls to LLMs.

## Embeddings & Retrieval

### topK(items, k)

Returns the top K items from a list of `[item, score]` pairs, sorted by score (highest first).

```js
items := [['a', 0.9], ['b', 0.5], ['c', 0.8]]
ai.topK(items, 2)  // => [['a', 0.9], ['c', 0.8]]
```

**Use case**: Retrieval-augmented generation (RAG), recommendation systems.

### kosineSimilaritySearch(embeddings, queryEmbedding, k)

Finds the top K most similar embeddings using cosine similarity. Standard for semantic search.

```js
embeddings := [[1, 0], [0, 1], [0.9, 0.1]]
query := [1, 0]
results := ai.kosineSimilaritySearch(embeddings, query, 2)
// Returns indices and similarity scores of top 2 matches
```

**Use case**: Vector database search, semantic similarity matching, RAG systems.

## Batch Processing

### batchify(items, batchSize)

Divides a list into batches of specified size. Useful for processing data in mini-batches.

```js
ai.batchify([1, 2, 3, 4, 5], 2)  // => [[1, 2], [3, 4], [5]]
ai.batchify(data, 32)             // Create batches of 32 samples
```

**Use case**: Training loops, batch inference.

### padSequence(sequence, targetLength, padToken)

Pads a sequence to a target length with a padding token. Essential for variable-length input handling.

```js
ai.padSequence([1, 2, 3], 5, 0)    // => [1, 2, 3, 0, 0]
ai.padSequence([1], 4, -1)         // => [1, -1, -1, -1]
```

**Use case**: Preparing sequences for transformer models, NLP preprocessing.

## Statistical Utilities

### mean(data)

Computes the arithmetic mean of a dataset.

```js
ai.mean([1, 2, 3, 4, 5])  // => 3
ai.mean([])                // => 0
```

### variance(data)

Computes the variance (spread of data around mean).

```js
ai.variance([1, 2, 3, 4, 5])  // => 2
```

### stddev(data)

Computes the standard deviation (square root of variance).

```js
ai.stddev([1, 2, 3, 4, 5])  // => 1.414...
ai.stddev([5, 5, 5])         // => 0 (no variance)
```

### entropy(probabilities)

Computes Shannon entropy of a probability distribution. Measures uncertainty.

```js
ai.entropy([0.5, 0.5])   // => 1.0 (maximum entropy)
ai.entropy([1, 0])       // => 0   (no entropy)
```

**Use case**: Information theory, decision trees, uncertainty quantification.

## Model Configuration

### ModelConfig(params)

Creates a configuration object for AI models. Centralizes hyperparameters and architecture details.

```js
config := ai.ModelConfig({
    name: 'gpt-4-turbo'
    vocabSize: 100000
    hiddenSize: 8192
    numLayers: 96
    numHeads: 96
    maxSequenceLength: 128000
    learningRate: 0.0001
    precision: 'bfloat16'
})
```

**Default values:**
- `vocabSize`: 50,000
- `hiddenSize`: 768
- `numLayers`: 12
- `numHeads`: 12
- `maxSequenceLength`: 2048
- `dropoutRate`: 0.1
- `learningRate`: 0.0001
- `batchSize`: 32
- `trainable`: true
- `precision`: 'float32'

## 2026 Standards & Best Practices

### Embedding Normalization

Always normalize embeddings before similarity computation:

```js
normalized := ai.normalize(embedding)
similarity := ai.cosineSimilarity(normalized, query)
```

### Attention Mechanisms

The library supports scaled dot-product attention, the foundation of all modern transformers:

```js
scale := 1 / ai.magnitude([ai.hiddenSize]) // or 1/sqrt(d_k)
attention := ai.scaledDotProductAttention(Q, K, V, scale)
```

### Data Preprocessing Pipeline

Standard 2026 preprocessing flow:

```js
// 1. Tokenize
tokens := ai.tokenizeSimple(text)

// 2. Normalize
normalized := ai.minMaxScale(tokens)

// 3. Batch
batches := ai.batchify(normalized.data, 32)

// 4. Pad to sequence length
batches := batches |> map(fn(b) ai.padSequence(b, 512, 0))
```

### Loss & Optimization

Use GELU activations and modern loss functions:

```js
// Forward pass with GELU
hidden := ai.gelu(x)

// Compute loss
loss := ai.meanSquaredError(predictions, targets)
```

### Retrieval-Augmented Generation (RAG)

Efficient similarity search for context retrieval:

```js
docs := [[...], [...], ...]  // Document embeddings
queryEmbed := computeEmbedding(userQuery)

topDocs := ai.kosineSimilaritySearch(docs, queryEmbed, 5)
// Use top 5 docs for context
```

## Performance Considerations

1. **Cosine Similarity** over Euclidean distance for normalized embeddings (faster)
2. **Batch Processing** for throughput: `batchify(data, 32)`
3. **GELU** activations preferred over ReLU (2026 standard)
4. **Min-max scaling** for bounded outputs, z-score for unbounded
5. **Vector normalization** before similarity operations

## See Also

- [std](std.md) - Standard library
- [math](math.md) - Mathematical functions
- [json](json.md) - JSON serialization
- [build](build.md) - Build system

## Matrix Operations

### matZeros(rows, cols)

Creates an `m × n` matrix filled with zeros.

```oak
ai.matZeros(2, 3)  // => [[0, 0, 0], [0, 0, 0]]
```

### matIdentity(n)

Creates an `n × n` identity matrix.

```oak
ai.matIdentity(3)  // => [[1,0,0],[0,1,0],[0,0,1]]
```

### matTranspose(mat)

Transposes a matrix (swaps rows and columns).

```oak
ai.matTranspose([[1, 2], [3, 4], [5, 6]])  // => [[1, 3, 5], [2, 4, 6]]
```

### matMul(a, b)

Multiplies two matrices `A (m × n)` and `B (n × p)`, returning an `m × p` matrix.

```oak
ai.matMul([[1, 2], [3, 4]], [[5, 6], [7, 8]])  // => [[19, 22], [43, 50]]
```

### matAdd(a, b)

Adds two matrices element-wise.

```oak
ai.matAdd([[1, 2], [3, 4]], [[5, 6], [7, 8]])  // => [[6, 8], [10, 12]]
```

### matScale(mat, scalar)

Multiplies every element in a matrix by a scalar.

```oak
ai.matScale([[1, 2], [3, 4]], 3)  // => [[3, 6], [9, 12]]
```

### vecAdd(v1, v2)

Adds two vectors element-wise.

```oak
ai.vecAdd([1, 2, 3], [4, 5, 6])  // => [5, 7, 9]
```

### vecSub(v1, v2)

Subtracts `v2` from `v1` element-wise.

```oak
ai.vecSub([4, 5, 6], [1, 2, 3])  // => [3, 3, 3]
```

### vecScale(v, scalar)

Multiplies every element of a vector by a scalar.

```oak
ai.vecScale([1, 2, 3], 2)  // => [2, 4, 6]
```

## Classification Metrics

### accuracy(predicted, actual)

Computes the fraction of correct predictions.

```oak
ai.accuracy([1, 0, 1, 1], [1, 0, 0, 1])  // => 0.75
```

### precisionScore(predicted, actual, positiveLabel)

Computes precision for a binary classifier: `TP / (TP + FP)`.

```oak
ai.precisionScore([1, 1, 0, 1], [1, 0, 0, 1], 1)  // => 0.666...
```

### recallScore(predicted, actual, positiveLabel)

Computes recall for a binary classifier: `TP / (TP + FN)`.

```oak
ai.recallScore([1, 1, 0, 1], [1, 0, 0, 1], 1)  // => 1.0
```

### f1Score(predicted, actual, positiveLabel)

Computes the F1 score (harmonic mean of precision and recall).

```oak
ai.f1Score([1, 1, 0, 1], [1, 0, 0, 1], 1)  // => 0.8
```

### confusionMatrix(predicted, actual, numClasses)

Builds a confusion matrix for integer class labels `0..numClasses-1`. Returns a `numClasses × numClasses` matrix where `mat[actual][predicted] = count`.

```oak
ai.confusionMatrix([0, 1, 1, 2], [0, 1, 2, 2], 3)
// => [[1,0,0],[0,1,0],[0,1,1]]
```

## Data Preprocessing

### oneHotEncode(label, numClasses)

Encodes an integer label as a one-hot vector of given length.

```oak
ai.oneHotEncode(2, 5)  // => [0, 0, 1, 0, 0]
```

### oneHotDecode(v)

Returns the index of the maximum value in a vector (argmax).

```oak
ai.oneHotDecode([0, 0, 1, 0])    // => 2
ai.oneHotDecode([0.1, 0.7, 0.2]) // => 1
```

### shuffleIndices(n, seed)

Returns a list of indices `[0..n-1]` in pseudo-random order using a deterministic LCG seeded by `seed`.

```oak
ai.shuffleIndices(5, 42)  // => [3, 0, 4, 1, 2] (deterministic)
```

### trainTestSplit(data, labels, ratio, seed)

Splits data and labels into train and test sets. `ratio` is the fraction used for training (e.g. 0.8). Returns `{trainData, trainLabels, testData, testLabels}`.

```oak
ai.trainTestSplit([[1],[2],[3],[4],[5]], [0,1,0,1,0], 0.6, 42)
```

### kFoldIndices(n, k)

Generates `k` folds of train/test index splits for cross-validation. Returns a list of `{train: [...], test: [...]}` objects.

```oak
ai.kFoldIndices(10, 5)  // => [{train: [2..9], test: [0,1]}, ...]
```

## K-Nearest Neighbors

### knn(data, labels, query, k)

Classifies a query point using k-nearest neighbors from a labeled dataset. Returns the most common label among the `k` nearest neighbors (by Euclidean distance).

```oak
ai.knn([[0,0],[1,1],[2,2],[3,3]], [:a,:a,:b,:b], [1.5,1.5], 3)  // => :a
```

## Positional Encoding (Transformer)

### positionalEncoding(seqLen, dModel)

Generates sinusoidal positional encodings as used in the original Transformer architecture ("Attention Is All You Need"). Returns a `seqLen × dModel` matrix.

```oak
ai.positionalEncoding(4, 8)  // => 4x8 matrix of positional encodings
```

## Layer Normalization

### layerNorm(v, gamma, beta)

Applies layer normalization to a vector. Normalizes to zero mean and unit variance, then applies scale (`gamma`) and shift (`beta`) parameters.

```oak
ai.layerNorm([1, 2, 3, 4], 1, 0)  // => [-1.34, -0.45, 0.45, 1.34] (approx)
```

## Learning Rate Scheduling

### lrConstant(baseLr, step)

Returns a constant learning rate regardless of step.

### lrStepDecay(baseLr, step, stepSize, decayFactor)

Returns a learning rate that decays by `decayFactor` every `stepSize` steps.

```oak
ai.lrStepDecay(0.01, 25, 10, 0.5)  // => 0.0025
```

### lrExponentialDecay(baseLr, step, decayRate)

Returns a learning rate with exponential decay.

```oak
ai.lrExponentialDecay(0.01, 100, 0.96)  // => 0.01 * 0.96^100
```

### lrCosineAnnealing(baseLr, step, totalSteps, minLr)

Returns a learning rate following a cosine annealing schedule. Smoothly decreases from `baseLr` to `minLr` over `totalSteps`.

```oak
ai.lrCosineAnnealing(0.01, 50, 100, 0.0001)  // => ~0.005
```

### lrWarmupLinear(baseLr, step, warmupSteps, totalSteps)

Linear warmup followed by linear decay. During warmup, `lr` increases linearly from 0 to `baseLr`.

```oak
ai.lrWarmupLinear(0.01, 5, 10, 100)   // => 0.005 (in warmup)
ai.lrWarmupLinear(0.01, 50, 10, 100)  // => ~0.0055 (in decay)
```

## Weight Initialization

### initUniform(n, lo, hi, seed)

Generates a vector of `n` values uniformly distributed in `[lo, hi]` using a deterministic LCG.

```oak
ai.initUniform(5, -1, 1, 42)
```

### initXavier(n, fanIn, fanOut, seed)

Xavier/Glorot uniform initialization. Suitable for layers with `fanIn` inputs and `fanOut` outputs.

```oak
ai.initXavier(10, 128, 64, 42)
```

### initHe(n, fanIn, seed)

He (Kaiming) uniform initialization. Suitable for ReLU-activated layers.

```oak
ai.initHe(10, 128, 42)
```

## Gradient Descent Helpers

### sgdStep(params, grads, lr)

Performs a single SGD update on a parameter vector.

```oak
ai.sgdStep([1, 2, 3], [0.1, 0.2, 0.1], 0.01)  // => [0.999, 1.998, 2.999]
```

### sgdMomentumStep(params, grads, velocity, lr, momentum)

Performs an SGD step with momentum. Returns `{params, velocity}`.

```oak
ai.sgdMomentumStep([1,2], [0.1,0.2], [0,0], 0.01, 0.9)
```

### clipGradNorm(grads, maxNorm)

Scales a gradient vector so its L2 norm does not exceed `maxNorm`.

```oak
ai.clipGradNorm([3, 4], 1)  // => [0.6, 0.8]
```

## Sampling & Decoding

### argmax(v)

Returns the index of the largest element in a vector.

```oak
ai.argmax([0.1, 0.7, 0.2])  // => 1
```

### temperatureScale(logits, temperature)

Applies temperature scaling to logits before softmax. Lower temperature sharpens the distribution.

```oak
ai.temperatureScale([1, 2, 3], 0.5)  // => [2, 4, 6]
```

### topPFilter(logits, p)

Nucleus sampling: filters logits to keep only tokens whose cumulative softmax probability mass is ≤ `p`. Rejected positions are set to a large negative value.

```oak
ai.topPFilter([1, 2, 3, 0.5], 0.9)
```

## Embedding Lookup

### embeddingLookup(table, indices)

Retrieves embedding vectors for a list of token indices from an embedding matrix.

```oak
table := [[0.1, 0.2], [0.3, 0.4], [0.5, 0.6]]
ai.embeddingLookup(table, [2, 0])  // => [[0.5, 0.6], [0.1, 0.2]]
```

## Parallel Utilities

### pmatMul(a, b)

Multiplies two matrices with parallelism across rows.

```oak
ai.pmatMul([[1,2],[3,4]], [[5,6],[7,8]])  // => [[19,22],[43,50]]
```

### pbatchSoftmax(vectors)

Applies softmax to each vector in a list, in parallel.

```oak
ai.pbatchSoftmax([[1,2,3],[4,5,6]])
```

### pbatchMeanSquaredError(predictedBatch, actualBatch)

Computes MSE between corresponding pairs from two lists of vectors, in parallel.

```oak
ai.pbatchMeanSquaredError([[1,2],[3,4]], [[1.1,2.1],[3.2,3.8]])
```

