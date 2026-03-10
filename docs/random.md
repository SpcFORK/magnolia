# Random Library (random)

## Overview

`librandom` provides utilities for generating pseudorandom values using Oak's built-in `rand()` function. **Not suitable for cryptographic or security-sensitive applications** - use `srand()` or the `crypto` library instead.

## Import

```oak
random := import('random')
// or destructure specific functions
{ boolean: boolean, integer: integer, number: number, choice: choice } := import('random')
```

## Functions

### `boolean()`

Returns either `true` or `false` with equal probability (50/50).

```oak
{ boolean: boolean } := import('random')

boolean() // =>  true or false (50% chance each)

// Example: coin flip
if boolean() {
    true -> println('Heads!')
    _ -> println('Tails!')
}
```

### `integer(min, max)`

Returns a random integer in the range `[min, max)` (inclusive min, exclusive max) with uniform probability.

If only one argument is provided, it's treated as `max` with `min = 0`.

```oak
{ integer: integer } := import('random')

integer(1, 7) // => 1, 2, 3, 4, 5, or 6 (dice roll)
integer(0, 10) // => 0-9
integer(10) // => 0-9 (same as above)

integer(100, 200) // => 100-199
integer(-10, 10) // => -10 to 9
```

### `number(min, max)`

Returns a random floating-point number in the range `[min, max)` with uniform probability.

If only one argument is provided, it's treated as `max` with `min = 0`.

```oak
{ number: number } := import('random')

number(0, 1) // => 0.0 to 0.999...
number(1) // => 0.0 to 0.999... (same as above)

number(5.0, 10.0) // => 5.0 to 9.999...
number(-1.0, 1.0) // => -1.0 to 0.999...

// Random temperature
temp := number(20, 30) // => 20.0-29.9°C
```

### `choice(list)`

Returns a random element from a list, with each element having equal probability.

```oak
{ choice: choice } := import('random')

fruits := ['apple', 'banana', 'orange', 'grape']
choice(fruits) // => One of the fruits at random

colors := ['red', 'green', 'blue']
choice(colors) // => 'red', 'green', or 'blue'

// Empty list behavior
choice([]) // => ? (undefined behavior)
```

### `normal()`

Returns a sample from a standard normal distribution with mean μ = 0 and standard deviation σ = 1.

Uses the Box-Muller transform for generation.

```oak
{ normal: normal } := import('random')

normal() // => ~68% chance between -1 and 1
        // => ~95% chance between -2 and 2
        // => ~99.7% chance between -3 and 3

// Generate values around a custom mean and stddev
mean := 100
stddev := 15
value := mean + stddev * normal()
// => IQ score distribution (μ=100, σ=15)
```

## Examples

### Dice Rolling

```oak
{ integer: integer } := import('random')

fn rollDice(sides) integer(1, sides + 1)
fn roll2d6 rollDice(6) + rollDice(6)

println('D6: ' + string(rollDice(6)))
println('D20: ' + string(rollDice(20)))
println('2D6: ' + string(roll2d6()))
```

### Random Color Generation

```oak
{ integer: integer } := import('random')

fn randomColor {
    r := integer(0, 256)
    g := integer(0, 256)
    b := integer(0, 256)
    'rgb(' + string(r) + ', ' + string(g) + ', ' + string(b) + ')'
}

randomColor() // => 'rgb(142, 68, 233)'
```

### Shuffling a List

```oak
{ integer: integer } := import('random')

fn shuffle(list) {
    shuffled := clone(list)
    n := len(shuffled)
    
    // Fisher-Yates shuffle
    fn sub(i) if i < n {
        true -> {
            j := integer(i, n)
            tmp := shuffled.(i)
            shuffled.(i) := shuffled.(j)
            shuffled.(j) := tmp
            sub(i + 1)
        }
        _ -> shuffled
    }
    
    sub(0)
}

deck := ['A', '2', '3', '4', '5', '6', '7', '8', '9', '10', 'J', 'Q', 'K']
shuffled := shuffle(deck)
```

### Random String Generation

```oak
{ choice: choice, integer: integer } := import('random')

fn randomString(length) {
    chars := 'abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789'
    
    fn sub(i, acc) if i {
        length -> acc
        _ -> sub(i + 1, acc << chars.(integer(0, len(chars))))
    }
    
    sub(0, '')
}

randomString(10) // => 'aB3kDm9pQ2'
```

### Weighted Random Choice

```oak
{ number: number } := import('random')

fn weightedChoice(items, weights) {
    total := sum(weights...)
    r := number(0, total)
    cumulative := 0
    
    fn sub(i) {
        cumulative <- cumulative + weights.(i)
        if r < cumulative {
            true -> items.(i)
            _ -> sub(i + 1)
        }
    }
    
    sub(0)
}

// 50% common, 30% uncommon, 20% rare
items := ['common', 'uncommon', 'rare']
weights := [50, 30, 20]
weightedChoice(items, weights)
```

### Random Point in Circle

```oak
{ number: number, normal: normal } := import('random')
{ sqrt: sqrt, Pi: Pi, cos: cos, sin: sin } := import('math')

// Uniform distribution in circle
fn randomInCircle(radius) {
    r := sqrt(number(0, 1)) * radius
    theta := number(0, 2 * Pi)
    [
        r * cos(theta)
        r * sin(theta)
    ]
}

[x, y] := randomInCircle(10)
// => Point within circle of radius 10
```

### Monte Carlo Simulation

```oak
{ number: number, integer: integer } := import('random')

// Estimate π using Monte Carlo method
fn estimatePi(samples) {
    insideCircle := 0
    
    fn sub(i) if i < samples {
        true -> {
            x := number(-1, 1)
            y := number(-1, 1)
            if x * x + y * y < 1 -> insideCircle <- insideCircle + 1
            sub(i + 1)
        }
        _ -> (4.0 * insideCircle) / samples
    }
    
    sub(0)
}

estimate := estimatePi(100000)
println('π ≈ ' + string(estimate))
```

### Random Delays

```oak
{ number: number } := import('random')

fn randomDelay(min, max) {
    delay := number(min, max)
    // Use with wait() builtin
    wait(delay)
}

// Random delay between 1-3 seconds
randomDelay(1, 3)
```

### Gaussian/Normal Distribution Data

```oak
{ normal: normal } := import('random')
{ round: round } := import('math')

// Generate test scores (mean=75, stddev=10)
fn generateTestScore {
    score := 75 + 10 * normal()
    round(clamp(score, 0, 100))
}

scores := range(30) |> map(fn generateTestScore)
println('Average: ' + string(mean(scores)))
```

## Implementation Notes

- Uses Oak's `rand()` builtin for random number generation
- `rand()` returns a float in `[0, 1)`
- **Not cryptographically secure** - use `srand()` or `crypto` library for security
- `normal()` uses the Box-Muller transform algorithm
- Each call to `rand()` advances the PRNG state
- Seed control depends on Oak runtime implementation

## Probability Distributions

```oak
{ integer: integer, normal: normal } := import('random')

// Uniform distribution [0, 1)
uniform := rand()

// Uniform distribution [min, max)
uniform := number(min, max)

// Discrete uniform distribution [min, max)
discrete := integer(min, max)

// Standard normal distribution (μ=0, σ=1)
gaussian := normal()

// Custom normal distribution
customNormal := mean + stddev * normal()
```

## Limitations

- No seeding control (depends on runtime)
- No other probability distributions (exponential, Poisson, etc.)
- No random sampling without replacement
- No cryptographic randomness (use `srand()` for that)
- `choice()` doesn't handle empty lists gracefully
- No shuffle function (must implement manually)
- No weighted random choice (must implement manually)

## Security Warning

⚠️ **DO NOT use `librandom` for:**
- Session tokens
- Cryptographic keys
- Password generation
- Security-sensitive UUIDs
- CSRF tokens
- Authentication codes

**Use instead:**
- `srand()` builtin for secure random bytes
- `crypto.uuid()` for unique identifiers
- Custom implementations with cryptographic RNGs

## See Also

- `crypto` library - For cryptographically secure UUID generation
- Oak built-in `rand()` - Pseudorandom number (0-1)
- Oak built-in `srand(n)` - Secure random bytes
- `math` library - For mathematical functions used in random algorithms
