# GoYiyi
提供一个Golang 的bypass AV 思路。
## 流程
1. 将Cobalt_Strike 生成的shellcode(C语言，格式替换为0x00,0x00)放入AES 中加密;
2. 将加密过后的shellcode 放入文件内存加载shellcode的main.go中;
3. 用-ldflags="-H windowsgui" 打包内存加载shellcode。

(亲测过火绒、360和Windows Defender)
