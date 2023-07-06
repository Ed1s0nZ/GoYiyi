# GoYiyi
提供一个Golang 的bypass AV 思路。
## 文件介绍
1. AES加密 -> 将shellcode进行AES加密;
2. 内存加载shellcode -> 用来生成可执行木马;
3. 内存加载shellcode2 -> "内存加载shellcode"加载器重构版, 也是用来生成可执行木马, 能过的杀软比"内存加载shellcode"多;
4. 反虚拟机 -> 在"内存加载shellcode2"基础上增加了反虚拟机的功能。
## 流程
总的来说，就是AES加密shellcode -> 用加载器解密并加载shellcode:
1. 将Cobalt_Strike 生成的shellcode(C语言，格式替换为0x00,0x00)放入"AES"中加密;
2. 将加密过后的shellcode 放入文件"内存加载shellcode(内存加载shellcode2、反虚拟机)"的main.go中;
3. 用-ldflags="-H windowsgui" 打包内存加载shellcode(或者内存加载shellcode2、反虚拟机)。

(用"内存加载shellcode"打包的exe不能过火绒; 用"内存加载shellcode2"打包的exe过火绒、360和Windows Defender等)

## Star History

[![Star History Chart](https://api.star-history.com/svg?repos=Ed1s0nZ/GoYiyi&type=Timeline)](https://star-history.com/#Ed1s0nZ/GoYiyi&Timeline)
