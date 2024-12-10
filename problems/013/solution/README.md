# 13. Roman to Integer


## Compare Memory Use
```bash
go test -bench=. -benchmem -memprofile mem.out && \
go tool pprof mem.out 
```

Then see running `top`:
```text
Showing nodes accounting for 70MB, 100% of 70MB total
```