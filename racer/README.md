## test data races

tl;dr 2 writers is a race

output:

```
2020/03/27 11:20:56 0 starting
2020/03/27 11:20:56 1 starting
==================
WARNING: DATA RACE
Write at 0x00c00012e000 by goroutine 8:
2020/03/27 11:20:56 0 done
  main.main.func1()
      /home/schmichael/go/src/github.com/schmichael/play/racer/racer.go:16 +0xaa

Previous write at 0x00c00012e000 by goroutine 7:
  main.main.func1()
      /home/schmichael/go/src/github.com/schmichael/play/racer/racer.go:16 +0xaa

Goroutine 8 (running) created at:
  main.main()
      /home/schmichael/go/src/github.com/schmichael/play/racer/racer.go:13 +0xb6

2020/03/27 11:20:56 3 starting
Goroutine 7 (running) created at:
  main.main()
      /home/schmichael/go/src/github.com/schmichael/play/racer/racer.go:13 +0xb6
==================
2020/03/27 11:20:56 2 starting
2020/03/27 11:20:56 4 starting
2020/03/27 11:20:56 4 done
2020/03/27 11:20:56 2 done
2020/03/27 11:20:56 5 starting
2020/03/27 11:20:56 5 done
2020/03/27 11:20:56 1 done
2020/03/27 11:20:56 6 starting
2020/03/27 11:20:56 6 done
2020/03/27 11:20:56 3 done
2020/03/27 11:20:56 7 starting
2020/03/27 11:20:56 7 done
2020/03/27 11:20:56 8 starting
2020/03/27 11:20:56 8 done
2020/03/27 11:20:56 > waiting
2020/03/27 11:20:56 9 starting
2020/03/27 11:20:56 > done
2020/03/27 11:20:56 9 done
Found 1 data race(s)
exit status 66
```
