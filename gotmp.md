go工具分析目前看来主要是trace和pprof 标准库 sync.pool runtime container context bufio bytes net-socket编程 


程序运行先run或者build 先 -race一下竞争检测
atomic能够以很底层的加锁机制来同步访问整型变量和指针 但是只能同步一个值 详见go实战


gc触发机制等fabric看完后再研究 runtime/mgc.go 源程序中还推荐去看垃圾回收算法手册那本书