# LOC

LOC - Line of Code is a tool which displays number of lines in each file in your project. It displays files in descending order of number of lines.

![example](https://github.com/affishaikh/loc-go/blob/main/images/example.png?raw=true)

## Ignoring files/directories

### .locignore

If you want to ignore certain directories or files from project like .git or .idea, then put a list inside **$HOME/.locignore** file

```text
.idea
.git
```

### cli option to ignore

You can use cli option `ignore` to skip certain files/directories. You just have to put names of files/directories

```text
    loc ignore="file.txt,test_dir"
```