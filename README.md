#flatten


## including dirs
A program written in Go to flatten a directory.

Assuming src is:
src
src/a.txt
src/x
src/x/b.txt

With the command:
`flatten src dst`

dst is now:
dst
dst/a.txt
dst/x-b.txt


## without including dirs

Assuming src is:
src
src/a.txt
src/x
src/x/b.txt

With the command:
`flatten -d=false src dst`

dst is now:
dst
dst/a.txt
dst/b.txt

