# a-very-golf
Riley Poole inspired anagram generator.

run `go build` to compile.

usage: binary-name "aefglorvy "

### example output
```
./a-glove-fry.exe "aefglorvyle "
frage volley | floe gravely | folly egrave | forgave yell | valley forge | yell forgave
flavor elegy | frage lovely | leafy grovel | voyager fell | gravely floe | gravelly foe
grovel leafy | egrave folly | elegy flavor | fell voyager | foe gravelly | forge valley
lovely frage | volley frage

Time elapsed: 6m5.5597571s
Iterations: 344058066
```

### notes and tips
add as many spaces to your anagram as you wish. it doesn't add spaces automatically.

inputting more than 10 characters will drastically increase anagram counts due to recursive calls.

you will start waiting minutes or hours depending on how many characters you go beyond 10.
