[![Readme Card](https://github-readme-stats.vercel.app/api/pin/?username=cyclone-github&repo=jottiUploader&theme=gruvbox)](https://github.com/cyclone-github/jottiUploader/)

[![Go Report Card](https://goreportcard.com/badge/github.com/cyclone-github/jottiUploader)](https://goreportcard.com/report/github.com/cyclone-github/jottiUploader)
[![GitHub issues](https://img.shields.io/github/issues/cyclone-github/jottiUploader.svg)](https://github.com/cyclone-github/jottiUploader/issues)
[![License](https://img.shields.io/github/license/cyclone-github/jottiUploader.svg)](LICENSE)
[![GitHub release](https://img.shields.io/github/release/cyclone-github/jottiUploader.svg)](https://github.com/cyclone-github/jottiUploader/releases)

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
### Compile from source:
- If you want the latest features, compiling from source is the best option since the release version may run several revisions behind the source code.
- This assumes you have Go and Git installed
  - `git clone https://github.com/cyclone-github/jottiUploader.git`
  - `cd jottiUploader`
  - `go mod init jottiUploader`
  - `go mod tidy`
  - `go build -ldflags="-s -w" .`
- Compile from source code how-to:
  - https://github.com/cyclone-github/scripts/blob/main/intro_to_go.txt
