Provides IsEmpty function for directory that are not present in
the stdlib. This is because I used f.Readdirnames(-1) (https://github.com/golang/go/issues/36197) in my code
base once and that lead to memory leak because of the huge number
of allocations required. Using this function will ensure that this
mistake is not repeated.

