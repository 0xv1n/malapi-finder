# malapi-finder
POC Code that scans an arbitrary file for known well-defined structures of commonly used Win32 API.


# Usage

`go run finder.go <filename>`

You can test the project by using the included `examples` file which should output the following:

```
Found matching function signature: VirtualAllocEx
Parsed parameters: [HANDLE hProcess LPVOID lpAddress SIZE_T dwSize DWORD  flAllocationType DWORD  flProtect]
Found matching function signature: NotVirtualAlloc
Parsed parameters: [HANDLE hProcess LPVOID lpAddress SIZE_T dwSize DWORD  flAllocationType DWORD  flProtect]
```
