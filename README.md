# Description

The following is quite a naïve attempt at solving the [Penniless Pilgrim problem](https://www.youtube.com/watch?v=6sBB-gRhfjE) in Golang.

In this solution, the program stores all possible states the pilgrim could be in as he moves through the city, and the number of states already explodes when the city is comprised of 6 streets. A less memory-intensive solution would be, for instance, only having one state and having some sort of "memory" of the blocks the pilgrim has not visited yet.