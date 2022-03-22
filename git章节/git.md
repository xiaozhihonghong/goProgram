# git连接远程教程
初始化仓库：git int

建立连接：git remote add origin 远程仓库地址

远程仓库pull到本地:git pull origin master.
此时出现错误，fatal: couldn't find remote ref master

解决方式：添加一个txt文件，git add test.txt

git commit -m "first"

然后首次连接仓库： git push --set-upstream origin master 完成