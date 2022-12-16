# Day 13 Notes

Originally I solved this using what I thought would be a relatively simple OO approach (see [day13.go](day13.go)).
But this is very hacky in Go. It includes a couple of casts which seem like code smells.
The input string parser is also very complex with this "simple" method.

I saw another solution using a tree structure and decided to try it (see [day13Tree.go](day13Tree.go)). This is a much
better solution. Also, the input string parser is much simpler!
