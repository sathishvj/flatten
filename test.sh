# very simple test. just file counts.

SRC="src"
DST="dst"
go build -o flatten
if [ $? -ne 0 ]; then
	echo "build failed."
	exit 1
fi

rm -r $DST
./flatten $SRC $DST
srcCnt=`find $SRC -type f | wc -l`
dstCnt=`find $DST -type f | wc -l`
printf "srcCnt=%d, dstCnt=%d\n" $srcCnt $dstCnt
if [ $srcCnt -ne $dstCnt ]; then
	echo "test 1: Failed"
else
	echo "test 1: Passed"
fi

# also test with the -d option
echo 
rm -r $DST
./flatten -d=false $SRC $DST
srcCnt=`find $SRC -type f | wc -l`
dstCnt=`find $DST -type f | wc -l`
printf "srcCnt=%d, dstCnt=%d\n" $srcCnt $dstCnt
if [ $srcCnt -ne $dstCnt ]; then
	echo "test 2 with -d: Failed"
else
	echo "test 2 with -d: Passed"
fi
