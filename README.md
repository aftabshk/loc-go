# LOC

LOC - Line of Code is a tool which displays number of lines in each file in your project. By default, it displays files in descending order of number of lines. 

The first line of the output displays the total number of files on which loc is applied

![example](https://github.com/affishaikh/loc-go/blob/main/images/example.png?raw=true)

## Ignoring files/directories

There are following two ways two ignore certain files and folders

### .locignore

If you want to ignore certain directories or files from project like .git or .idea, then put a list inside **$HOME/.locignore** file

```text
.idea
.git
```

### -ignore (cli option to ignore)

You can use cli option `-ignore` to skip certain files/directories. You just have to put names of files/directories

```text
    loc -ignore file.txt,test_dir
```

## Sorting the output

By default, the output is displayed in descending order of number of lines (loc).

You have two others options to apply custom sort on - name & loc. And you can choose two orders ASC & DESC

```shell
    loc -sort name ASC
```