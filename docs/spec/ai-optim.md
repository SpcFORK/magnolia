# API Documentation

_Auto-generated from Magnolia source._

---

## Module: `lib\ai-optim.oak`

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

