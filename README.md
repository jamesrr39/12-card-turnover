# 12-card-turnover

An implementation of a simulation of the card game where:

12 cards are picked at random from a shuffled deck. The first is turned over, and the player guesses if the number on the next card will be higher or lower.

If the player guesses correctly (or the number is the same), this process is repeated. The process carries on until all cards are revealed (player wins) or the player guesses incorrectly (player loses).

## implementation

The computer player here takes the strategy of counting whether there are more cards higher or lower left each iteration, and guesses in that direction.

Make sure you run this with enough iterations! On my laptop 1000000 (one million) takes a minute or two and gives consistent results.

### Plot spoiler

[Don't click if you don't want to know!](./plot-spoiler.txt)