# Day 14 Notes

The key insight for me with this puzzle was regarding the data structure to represent the grid. Initially it seems
that an array would work - but there are two issues:

1. The array would be huge, with mainly empty space
2. We don't know the size of the array initially, so there would be a lot of work with resizing as we parsed the input

So it is better to use a map of points that have either sand or rocks. Is a point doesn't exist in the map, then 
we know it is air. This makes the parse much easier and also is far more memory efficient.

This is a similar insight to the solution for Day 12.
