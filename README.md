# MFR 📂
Go program that renames all files in a directory using a specified pattern.

## Requirements 🔗
<b>OS</b>:
- Linux
- Windows
- macOS (née OS X, aka Darwin)
- OpenBSD
- DragonFly BSD
- FreeBSD
- NetBSD
- Solaris

<b>To build</b>:
- Go compiler

## Quick usage guide 📚
#### Build:
```
go build --o mfr cmd/main.go
```

<br>

> <b>NOTE</b>: Generate in the root directory of the project

---

#### Arguments:
```
./mfr [/target/directory/path] [pattern]
```

---

#### Example:
```
./mfr /home/user/vitu/test/ file-
```

- Before:
```
/home/user/vitu/test/
├── text.txt
├── image.png
├── document.pdf
└── program.java
```

- After:
```
/home/user/vitu/test/
├── file-1.txt
├── file-2.png
├── file-3.pdf
└── file-4.java
```

## Future plans 📌
- Write more tests
- Filter by file extension
- Include help commands
- Allow customized sorting
- Create a backup of the renamed files
- Ask for confirmation in important directories

## Do you want help me? 👥
If you have any ideas or wish to contribute to the project, contact me on X (<a href="https://x.com/ohtoaki" target="_blank">@ohtoaki</a>) or send me a pull request :)

## License 🏳️
```
MIT License

Copyright (c) 2024 Vitu Ohto

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
```
