### Benchmark Results

**Environment:**
- **GOOS**: windows
- **GOARCH**: amd64
- **Package**: go.uber.org/zap/benchmarks
- **CPU**: AMD Ryzen 7 PRO 6850U with Radeon Graphics

**Details:**
Benchmarking disabled logging fields, adding context at each log site as described in `scenario_bench_test.go:280`.

| Benchmark                                   | Operations             | ns/op      | B/op      | allocs/op |
|---------------------------------------------|-----------------------|------------|-----------|-----------|
| BenchmarkDisabledAddingFields/Zap.Check-16  | 1,000,000,000         | 0.7998      | 0         | 0         |
| BenchmarkDisabledAddingFields/rs/zerolog-16 | 68,445,896            | 18.27       | 24        | 1         |
| BenchmarkDisabledAddingFields/Zap.Sugar-16  | 14,662,755            | 77.43       | 136       | 6         |
| BenchmarkDisabledAddingFields/slog-16       | 13,676,038            | 80.61       | 136       | 6         |
| BenchmarkDisabledAddingFields/Zap-16        | 4,632,270             | 226.1       | 736       | 5         |
| BenchmarkDisabledAddingFields/slog.LogAttrs-16 | 5,318,572         | 227.6       | 512       | 5         |
| BenchmarkDisabledAddingFields/apex/log-16   | 3,382,635             | 329.9       | 886       | 10        |
| BenchmarkDisabledAddingFields/ssgreg/logf-16| 2,890,052             | 432.7       | 1016      | 6         |
| BenchmarkDisabledAddingFields/sirupsen/logrus-16 | 1,570,825       | 751.7       | 1525      | 12        |

**Result:** PASS

![image (1).png](image%20%281%29.png)
![image (2).png](image%20%282%29.png)
Here's the provided text formatted in Markdown for improved readability and structure:

### Motivation
- There is a race condition bug with our existing ssgreg/logf implementation library.
- The current logging library is limiting in terms of features. Certain features like pretty print (can be set only for development custom stand), sampling, etc., will allow for a better developer experience.

### Summary of Findings
- **Zap.Check**: It is the fastest with zero allocation and zero memory usage. However, it requires a very specific aligning of fields where every odd field is the key and even is the value, e.g., `["int", 1, "foo", "valueOfFoo", "key", "value]`.
- **Logrus**: This is the slowest of all, likely because it was one of the earliest logging packages and was not maintained as well as others. Newer libraries have learned from it and improvised.
- **ssgreg/logf**: This library, heavily used in Acronis, does not seem to add any value compared to other popular libraries.
- **ZapSugar**: Syntactically similar to our current setup and should be a drop-in replacement.
- **Future Development**: Works are in place to use Go 1.21 `slog` with Zap as a backend.

### Ranking
Libraries that don't have type safety checks like `log.Int("intField", 5)` are removed. Here's the ranking based on the evaluation:
1. **Zerolog**
2. **Zap Sugar**
3. **Slog Attributes**
4. **Apex**
5. **ssgreg/logf**
6. **logrus**