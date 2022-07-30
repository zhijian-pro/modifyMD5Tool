# modifyMD5Tool
通过修改文件 md5 值避免云盘标记资源

## 下载与编译
```shell
git clone https://github.com/zhijian-pro/modifyMD5Tool.git && cd modifyMD5Tool
make
```
## 使用方法
```shell
# 修改 dirPath 下所有普通文件的 md5 值 （可以多次执行，已经被修改过的文件不会再次修改）
$ ./modifyMD5Tool -m dirPath
# 恢复 dirPath 下所有被修改过的普通文件 （可以多次执行，已经被恢复过的文件不会再被恢复）
$ ./modifyMD5Tool -r dirPath

# 修改单个文件的 md5 值（可以多次执行，已经被修改过的文件不会再次修改）
$ ./modifyMD5Tool --modify-One filePath
# 恢复单个文件（可以多次执行，已经被恢复过的文件不会再被恢复）
$ ./modifyMD5Tool --recover-One filePath

```
