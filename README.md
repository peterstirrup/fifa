# FIFA in Go

Go package that supports working with FIFA projects written in Go. Features include:
 * New player generation
 * Player representation
 * Name generation
 * Overall calculation
 * Typical player value calculation

Get started:
```
go get -u github.com/peterstirrup/fifa
```

Uses a little Testify in places. Designed to be as easy as possible to pick up and play without any prior knowledge
of the package.

### Generating a player
In the words of Commodores, that's why I'm easy.

```
p := fifa.GenPlayer()
```

### Generating n players
Again - soooooo easy!

```
p := fifa.GenPlayer(n)
```

### Adjusting generation

Uses a top player from FIFA 19 as a template and fluctuates stats.

Adjust four variables to get different results:
 * ```repeatRand``` Inversely proportional to the amount of chaos in stat generation.
 * ```maxDiff``` Again, the higher this is the more chaos we get in stat values.
 * ```qualityCo``` Inversely proportional to generated player 'quality' (i.e. A qualityCo closer to 0 will produce better players).
 * ```probOppositeFoot``` The probability a wide player will have a strong foot opposite to their side (i.e. if = 1 every right midfielder will have a left strong foot).
