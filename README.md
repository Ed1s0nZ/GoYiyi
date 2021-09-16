# GoYiyi
提供一个Golang 的bypass AV 思路。
## 文件介绍
1. AES加密 -> 将shellcode进行AES加密;
2. 内存加载shellcode -> 用来生成可执行木马;
3. 内存加载shellcode2 -> "内存加载shellcode"加载器重构版, 也是用来生成可执行木马, 能过的杀软比"内存加载shellcode"多。
## 流程
1. 将Cobalt_Strike 生成的shellcode(C语言，格式替换为0x00,0x00)放入"AES"中加密;
2. 将加密过后的shellcode 放入文件"内存加载shellcode(或者内存加载shellcode2)"的main.go中;
3. 用-ldflags="-H windowsgui" 打包内存加载shellcode(或者内存加载shellcode2)。

(用"内存加载shellcode"打包的exe不能过火绒; 用"内存加载shellcode2"打包的exe过火绒、360和Windows Defender等)
