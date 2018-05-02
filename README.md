## 第五章
今天在写第五章代码计算从1加到100的和的时候遇到一个错误如下  
goto指令当前pc是17,操作数是FFF3,可是我的代码直接报错index out of range，这里面我犯了两个错误，
1. 第一我把FFF3读成了F3FF,于是就变成了十进制的62463
2. 第二goto的操作数类型是int16,我处理成了uint16,
更正错误后FFF3就被正确的处理为-13,所以goto指令执行的时候就会跳转到pc为4的地址去

## 第六章
第六章的代码遇到一个空指针异常，在Go中的提示是这样的：

```
panic: runtime error: invalid memory address or nil pointer dereference
[signal SIGSEGV: segmentation violation code=0x1 addr=0x8 pc=0x10bffcf]
```
由于对go语言还不够熟悉，当时没有看出来这个错误究竟是怎么回事，后来才发现。
