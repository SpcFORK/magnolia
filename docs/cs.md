# Classes (cs)

## Overview

Magnolia supports class constructor syntax sugar through the `cs` keyword.

Classes in Magnolia are constructor-like values that are called like functions and typically return objects. They are useful for grouping object construction, constructor arguments, and methods that close over constructor state.

## Syntax

```oak
cs Name(arg1, arg2, rest...,) {
    // constructor body expression
}
```

The constructor body is an expression, commonly an object literal.

## Basic Usage

### Class Without Parameters

```oak
cs Empty {
    {}
}

type(Empty()) // => :object
```

### Constructor Arguments

```oak
cs Pair(left, right) {
    {
        left: left
        right: right
    }
}

Pair(1, 2).right // => 2
```

### Methods Closing Over Constructor State

```oak
cs Counter(start) {
    {
        value: start
        add: fn(delta) start + delta
    }
}

Counter(4).add(3) // => 7
```

### Variadic Parameters

```oak
cs Bag(items...,) {
    items
}

len(Bag(1, 2, 3)) // => 3
```

## Empty Body Behavior

```oak
cs Empty {}
Empty() // => ?
```

If the class body is a bare empty block (`{}`), construction returns `?`.

## Inheritance and Static Members

Magnolia supports inheritance-style composition from one or more parent classes, and static members on the class value.

```oak
cs Parent1 { a := 2, b := 1 }
cs Parent2 { c := 3, b := 2 }

cs Hi(make) {
    testStaticVar := 2
    fn testStaticFn {}

    (Parent1, Parent2) -> {
        { make: make }
    }
}

Hi('sedan').a // => 2
Hi('sedan').b // => 2
Hi('sedan').c // => 3
Hi('sedan').make // => 'sedan'
Hi.testStaticVar // => 2
type(Hi.testStaticFn) // => :function
```

When parent objects share keys, values from later parents in the list take precedence.

## Notes

- Classes are syntax sugar for constructor-style behavior.
- Constructor arguments are available in the class body scope.
- Class values are invoked like functions (for example, `Pair(1, 2)`).