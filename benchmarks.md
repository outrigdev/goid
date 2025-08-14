# Benchmark Results

Source: [GitHub Actions Run](https://github.com/outrigdev/goid/actions/runs/16952394384/job/48047934444)

**Note:** Linux armv7 results are from emulated environments and are significantly slower than native performance.

**See also:** [Test Results](testresults.md) for correctness testing across platforms.

| OS | Architecture | Go Version | Benchmark | Iterations | ns/op |
|---|---|---|---|---|---|
| darwin | amd64 | 1.21 | BenchmarkGet-4 | 281617 | 4246 |
| darwin | amd64 | 1.21 | BenchmarkGetFromStack-4 | 311266 | 3929 |
| darwin | amd64 | 1.22 | BenchmarkGet-4 | 353326 | 3426 |
| darwin | amd64 | 1.22 | BenchmarkGetFromStack-4 | 326725 | 3515 |
| darwin | amd64 | 1.23 | BenchmarkGet-4 | 510257491 | 2.354 |
| darwin | amd64 | 1.23 | BenchmarkGetFromStack-4 | 347304 | 3407 |
| darwin | amd64 | 1.24 | BenchmarkGet-4 | 438443386 | 2.978 |
| darwin | amd64 | 1.24 | BenchmarkGetFromStack-4 | 307723 | 3916 |
| darwin | amd64 | 1.25 | BenchmarkGet-4 | 480184009 | 2.504 |
| darwin | amd64 | 1.25 | BenchmarkGetFromStack-4 | 347685 | 3821 |
| darwin | arm64 | 1.21 | BenchmarkGet-3 | 542028 | 2374 |
| darwin | arm64 | 1.21 | BenchmarkGetFromStack-3 | 540763 | 2850 |
| darwin | arm64 | 1.22 | BenchmarkGet-3 | 499723 | 3351 |
| darwin | arm64 | 1.22 | BenchmarkGetFromStack-3 | 353186 | 3278 |
| darwin | arm64 | 1.23 | BenchmarkGet-3 | 563719041 | 2.191 |
| darwin | arm64 | 1.23 | BenchmarkGetFromStack-3 | 521589 | 2312 |
| darwin | arm64 | 1.24 | BenchmarkGet-3 | 456102579 | 2.501 |
| darwin | arm64 | 1.24 | BenchmarkGetFromStack-3 | 436033 | 2681 |
| darwin | arm64 | 1.25 | BenchmarkGet-3 | 463723759 | 2.737 |
| darwin | arm64 | 1.25 | BenchmarkGetFromStack-3 | 401673 | 3349 |
| linux | amd64 | 1.21 | BenchmarkGet-4 | 377533 | 3217 |
| linux | amd64 | 1.21 | BenchmarkGetFromStack-4 | 375058 | 3173 |
| linux | amd64 | 1.22 | BenchmarkGet-4 | 368794 | 3281 |
| linux | amd64 | 1.22 | BenchmarkGetFromStack-4 | 364500 | 3247 |
| linux | amd64 | 1.23 | BenchmarkGet-4 | 384602512 | 3.143 |
| linux | amd64 | 1.23 | BenchmarkGetFromStack-4 | 377143 | 3180 |
| linux | amd64 | 1.24 | BenchmarkGet-4 | 384162895 | 3.116 |
| linux | amd64 | 1.24 | BenchmarkGetFromStack-4 | 359095 | 3309 |
| linux | amd64 | 1.25 | BenchmarkGet-4 | 385248692 | 3.117 |
| linux | amd64 | 1.25 | BenchmarkGetFromStack-4 | 362942 | 3257 |
| linux | arm64 | 1.21 | BenchmarkGet-4 | 292501 | 4066 |
| linux | arm64 | 1.21 | BenchmarkGetFromStack-4 | 293904 | 4058 |
| linux | arm64 | 1.22 | BenchmarkGet-4 | 311624 | 3834 |
| linux | arm64 | 1.22 | BenchmarkGetFromStack-4 | 314026 | 3812 |
| linux | arm64 | 1.23 | BenchmarkGet-4 | 678060669 | 1.769 |
| linux | arm64 | 1.23 | BenchmarkGetFromStack-4 | 316370 | 3765 |
| linux | arm64 | 1.24 | BenchmarkGet-4 | 678067620 | 1.769 |
| linux | arm64 | 1.24 | BenchmarkGetFromStack-4 | 297493 | 4015 |
| linux | arm64 | 1.25 | BenchmarkGet-4 | 677853956 | 1.769 |
| linux | arm64 | 1.25 | BenchmarkGetFromStack-4 | 340588 | 3521 |
| linux | armv7 | 1.21.13 | BenchmarkGet-4 | 21975 | 54786 |
| linux | armv7 | 1.21.13 | BenchmarkGetFromStack-4 | 22383 | 54206 |
| linux | armv7 | 1.22.12 | BenchmarkGet-4 | 22575 | 53237 |
| linux | armv7 | 1.22.12 | BenchmarkGetFromStack-4 | 22734 | 52224 |
| linux | armv7 | 1.23.10 | BenchmarkGet-4 | 21726 | 55565 |
| linux | armv7 | 1.23.10 | BenchmarkGetFromStack-4 | 21351 | 55053 |
| linux | armv7 | 1.24.4 | BenchmarkGet-4 | 22716 | 52393 |
| linux | armv7 | 1.24.4 | BenchmarkGetFromStack-4 | 22808 | 52463 |
| linux | armv7 | 1.25.0 | BenchmarkGet-4 | 22480 | 53489 |
| linux | armv7 | 1.25.0 | BenchmarkGetFromStack-4 | 22203 | 53753 |
| windows | 386 | 1.21 | BenchmarkGet-4 | 336234 | 3312 |
| windows | 386 | 1.21 | BenchmarkGetFromStack-4 | 395820 | 3229 |
| windows | 386 | 1.22 | BenchmarkGet-4 | 345986 | 3296 |
| windows | 386 | 1.22 | BenchmarkGetFromStack-4 | 395380 | 3304 |
| windows | 386 | 1.23 | BenchmarkGet-4 | 356799754 | 3.195 |
| windows | 386 | 1.23 | BenchmarkGetFromStack-4 | 346082 | 3475 |
| windows | 386 | 1.24 | BenchmarkGet-4 | 383640788 | 3.137 |
| windows | 386 | 1.24 | BenchmarkGetFromStack-4 | 335522 | 3366 |
| windows | 386 | 1.25 | BenchmarkGet-4 | 383506902 | 3.195 |
| windows | 386 | 1.25 | BenchmarkGetFromStack-4 | 361395 | 3305 |
| windows | amd64 | 1.21 | BenchmarkGet-4 | 390070 | 3239 |
| windows | amd64 | 1.21 | BenchmarkGetFromStack-4 | 396336 | 3283 |
| windows | amd64 | 1.22 | BenchmarkGet-4 | 356541 | 3304 |
| windows | amd64 | 1.22 | BenchmarkGetFromStack-4 | 395510 | 3302 |
| windows | amd64 | 1.23 | BenchmarkGet-4 | 373290446 | 3.117 |
| windows | amd64 | 1.23 | BenchmarkGetFromStack-4 | 353804 | 3376 |
| windows | amd64 | 1.24 | BenchmarkGet-4 | 336515743 | 3.572 |
| windows | amd64 | 1.24 | BenchmarkGetFromStack-4 | 322154 | 3509 |
| windows | amd64 | 1.25 | BenchmarkGet-4 | 384100788 | 3.168 |
| windows | amd64 | 1.25 | BenchmarkGetFromStack-4 | 369338 | 3337 |