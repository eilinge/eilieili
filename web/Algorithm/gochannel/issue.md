# Golang 性能调优

## 代码覆盖

    1. go test -coverprofile cover.out
    2. go tool cover -html=cover.out

## 性能调优

    1. go tool pprof
    2. go test -bench . -cpuprofile cpu.out
    3. go tool pprof cpu.out
    4. web

## 插件

    Graphviz
        1. package: https://graphviz.gitlab.io/_pages/Download/Download_windows.html
        2. windows: https://www.cnblogs.com/onemorepoint/p/8310996.html

## Golang常见性能调优

    1. https://www.jianshu.com/p/4e4ff6be6af9
    2. https://www.cnblogs.com/zhangboyu/p/7456609.html

    调优: https://segmentfault.com/a/1190000016354853

## 单元/集成/压力测试

    1.  TestAdd(t *testing.T)
        go test -v
        
        func TestAdd(t *testing.T) {
            s := Add(url)
            if s == "" {
                t.Errorf("Test.Add error!")
            }
        }

    2.  TestCase:
            BenchmarkAdd(b *testing.B)
        Run:
            go test -bench=. -cpuprofile=cpu.prof
        Watch:
            go tool pprof cpu.prof 
            (pprof) web

            go tool pprof -http=:8080 cpu.prof
            pprof -http=:8080 cpu.prof
        Eg:
            func BenchmarkAdd(b *testing.B) {
                fmt.Println("b.N ", b.N)
                for i := 0; i < b.N; i++ {
                    Add(url)
                }
            }
