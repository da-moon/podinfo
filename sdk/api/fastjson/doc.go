// Package fastjson encodes/decodes payloads to/from
// JSON format.
// besides a convenient API, it is significantly
// more performant when compared with stdlib
// keep in mind that sometimes, an operation may fail
// when this package is used, as it uses some 'unsafe'
// code to gain it's speed advantage
// ─── ENCODE BENCHMARKS ──────────────────────────────────────────────────────────
// goos: linux
// goarch: amd64
// pkg: github.com/da-moon/podinfo/sdk/api/fastjson
// cpu: Intel(R) Xeon(R) CPU
// BenchmarkEncode
// BenchmarkEncode-16          	  766094	      1447 ns/op	     104 B/op	       2 allocs/op
// BenchmarkStdLibEncode
// BenchmarkStdLibEncode-16    	  468454	      2918 ns/op	     520 B/op	      12 allocs/op
// PASS
// ok  	github.com/da-moon/podinfo/sdk/api/fastjson	2.532s
// ─── DECODE BENCHMARKS ──────────────────────────────────────────────────────────
// goos: linux
// goarch: amd64
// pkg: github.com/da-moon/podinfo/sdk/api/fastjson
// cpu: Intel(R) Xeon(R) CPU
// BenchmarkDecode
// BenchmarkDecode-16                    	  537007	      2355 ns/op	     544 B/op	      16 allocs/op
// BenchmarkStdLibDecode
// BenchmarkStdLibDecode-16              	  485347	      3222 ns/op	     736 B/op	      17 allocs/op
// BenchmarkDecodeFromReader
// BenchmarkDecodeFromReader-16          	  453264	      2579 ns/op	    1240 B/op	      19 allocs/op
// BenchmarkStdlibDecodeFromReader
// BenchmarkStdlibDecodeFromReader-16    	  401690	      3486 ns/op	    1480 B/op	      20 allocs/op
// PASS
// ok  	github.com/da-moon/podinfo/sdk/api/fastjson	5.518s
// ────────────────────────────────────────────────────────────────────────────────
package fastjson
