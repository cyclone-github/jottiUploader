# jottiUploader

```
$ ./jottiUploader_amd64.bin jottiUploader_amd64.exe
SHA1 Checksum: c62b74f0b4632bfae7ac01e160912b3eff6cb0a3
File jottiUploader_amd64.exe found on Jotti:
https://virusscan.jotti.org/en-US/search/hash/c62b74f0b4632bfae7ac01e160912b3eff6cb0a3
```
### About:
- This tool is a CLI file uploader for Jotti https://virusscan.jotti.org
- Jotti is a lesser-known alternative to VirusTotal
- Jotti enforces a rate limit which this tool honors once it has been reached. If you need to scan more files, consider supporting the Jotti project by purchasing an API key. 
### Usage Instructions:
```
./jottiUploader {file_to_scan}
./jottiUploader -help
./jottiUploader -version
```
### Change Log:
- https://github.com/cyclone-github/jottiUploader/blob/main/CHANGELOG.md
### Compile from source code info:
- https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt