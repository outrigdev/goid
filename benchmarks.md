# Benchmark Results

Source: [GitHub Actions Run](https://github.com/outrigdev/goid/actions/runs/16091915987/job/45409895718)

**Note:** Linux armv7 results are from emulated environments and are significantly slower than native performance.

**See also:** [Test Results](testresults.md) for correctness testing across platforms.

| OS | Architecture | Go Version | Benchmark | Iterations | ns/op |
|---|---|---|---|---|---|
| darwin | amd64 | 1.21 | BenchmarkGet-4 | 294028 | 3753 |
| darwin | amd64 | 1.21 | BenchmarkGetFromStack-4 | 331011 | 3609 |
| darwin | amd64 | 1.22 | BenchmarkGet-4 | 310675 | 3870 |
| darwin | amd64 | 1.22 | BenchmarkGetFromStack-4 | 297252 | 3906 |
| darwin | amd64 | 1.23 | BenchmarkGet-4 | 570482876 | 2.149 |
| darwin | amd64 | 1.23 | BenchmarkGetFromStack-4 | 366036 | 3425 |
| darwin | amd64 | 1.24 | BenchmarkGet-4 | 571967260 | 2.093 |
| darwin | amd64 | 1.24 | BenchmarkGetFromStack-4 | 353312 | 3328 |
| darwin | arm64 | 1.21 | BenchmarkGet-3 | 525079 | 2471 |
| darwin | arm64 | 1.21 | BenchmarkGetFromStack-3 | 484431 | 2530 |
| darwin | arm64 | 1.22 | BenchmarkGet-3 | 505934 | 2236 |
| darwin | arm64 | 1.22 | BenchmarkGetFromStack-3 | 501240 | 2338 |
| darwin | arm64 | 1.23 | BenchmarkGet-3 | 551324269 | 2.148 |
| darwin | arm64 | 1.23 | BenchmarkGetFromStack-3 | 529714 | 2259 |
| darwin | arm64 | 1.24 | BenchmarkGet-3 | 578396758 | 2.088 |
| darwin | arm64 | 1.24 | BenchmarkGetFromStack-3 | 495888 | 2360 |
| linux | amd64 | 1.21 | BenchmarkGet-4 | 364442 | 3204 |
| linux | amd64 | 1.21 | BenchmarkGetFromStack-4 | 378756 | 3189 |
| linux | amd64 | 1.22 | BenchmarkGet-4 | 361702 | 3257 |
| linux | amd64 | 1.22 | BenchmarkGetFromStack-4 | 365979 | 3283 |
| linux | amd64 | 1.23 | BenchmarkGet-4 | 384317960 | 3.123 |
| linux | amd64 | 1.23 | BenchmarkGetFromStack-4 | 366117 | 3279 |
| linux | amd64 | 1.24 | BenchmarkGet-4 | 385344177 | 3.119 |
| linux | amd64 | 1.24 | BenchmarkGetFromStack-4 | 372216 | 3235 |
| linux | arm64 | 1.21 | BenchmarkGet-4 | 293150 | 4072 |
| linux | arm64 | 1.21 | BenchmarkGetFromStack-4 | 294363 | 4065 |
| linux | arm64 | 1.22 | BenchmarkGet-4 | 310695 | 3812 |
| linux | arm64 | 1.22 | BenchmarkGetFromStack-4 | 314676 | 3820 |
| linux | arm64 | 1.23 | BenchmarkGet-4 | 678152516 | 1.769 |
| linux | arm64 | 1.23 | BenchmarkGetFromStack-4 | 318222 | 3753 |
| linux | arm64 | 1.24 | BenchmarkGet-4 | 678223470 | 1.769 |
| linux | arm64 | 1.24 | BenchmarkGetFromStack-4 | 294981 | 4027 |
| linux | armv7 | 1.21.13 | BenchmarkGet-4 | 22076 | 53649 |
| linux | armv7 | 1.21.13 | BenchmarkGetFromStack-4 | 22030 | 54560 |
| linux | armv7 | 1.22.12 | BenchmarkGet-4 | 22569 | 52879 |
| linux | armv7 | 1.22.12 | BenchmarkGetFromStack-4 | 23438 | 53174 |
| linux | armv7 | 1.23.10 | BenchmarkGet-4 | 21943 | 55462 |
| linux | armv7 | 1.23.10 | BenchmarkGetFromStack-4 | 21566 | 55758 |
| linux | armv7 | 1.24.4 | BenchmarkGet-4 | 22930 | 52757 |
| linux | armv7 | 1.24.4 | BenchmarkGetFromStack-4 | 22419 | 53334 |
| windows | 386 | 1.21 | BenchmarkGet-4 | 338602 | 3567 |
| windows | 386 | 1.21 | BenchmarkGetFromStack-4 | 345255 | 3498 |
| windows | 386 | 1.22 | BenchmarkGet-4 | 313646 | 3568 |
| windows | 386 | 1.22 | BenchmarkGetFromStack-4 | 348753 | 3573 |
| windows | 386 | 1.23 | BenchmarkGet-4 | 356015263 | 3.354 |
| windows | 386 | 1.23 | BenchmarkGetFromStack-4 | 299224 | 3601 |
| windows | 386 | 1.24 | BenchmarkGet-4 | 360921033 | 3.386 |
| windows | 386 | 1.24 | BenchmarkGetFromStack-4 | 351578 | 3563 |
| windows | amd64 | 1.21 | BenchmarkGet-4 | 320881 | 3476 |
| windows | amd64 | 1.21 | BenchmarkGetFromStack-4 | 331588 | 3491 |
| windows | amd64 | 1.22 | BenchmarkGet-4 | 373081 | 3291 |
| windows | amd64 | 1.22 | BenchmarkGetFromStack-4 | 376132 | 3301 |
| windows | amd64 | 1.23 | BenchmarkGet-4 | 358471441 | 3.348 |
| windows | amd64 | 1.23 | BenchmarkGetFromStack-4 | 346507 | 3629 |
| windows | amd64 | 1.24 | BenchmarkGet-4 | 354911352 | 3.326 |
| windows | amd64 | 1.24 | BenchmarkGetFromStack-4 | 310999 | 3658 |