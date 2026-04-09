# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ai-ml.oak`

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

