```bash
Running tool: /usr/local/go/bin/go test -benchmem -run=^$ -bench ^(BenchmarkOneEncrypt|BenchmarkOneDecrypt|BenchmarkOneEncryptDecrypt|BenchmarkEncrypt|BenchmarkDecrypt)$ eas

goos: linux
goarch: amd64
pkg: eas
cpu: Intel(R) Core(TM) i5-9500 CPU @ 3.00GHz
BenchmarkOneEncrypt
BenchmarkOneEncrypt-6              91299             12708 ns/op           26400 B/op         12 allocs/op
BenchmarkOneDecrypt
BenchmarkOneDecrypt-6             107198             11683 ns/op           16160 B/op         10 allocs/op
BenchmarkOneEncryptDecrypt
BenchmarkOneEncryptDecrypt-6       50672             29189 ns/op           42560 B/op         22 allocs/op
BenchmarkEncrypt
Run 0. Did 1000000 iterations in 15.050802735s
BenchmarkEncrypt-6                     1        29474572983 ns/op       52800045960 B/op        24000476 allocs/op
BenchmarkDecrypt
Run 0. Did 1000000 iterations in 13.455559085s
BenchmarkDecrypt-6                     1        26866821025 ns/op       32320034728 B/op        20000359 allocs/op
PASS
ok      eas     60.727s
```