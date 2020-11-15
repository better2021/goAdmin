
# SHELL
echo "执行shell"

if [ -n "$1" ] # 文件名后的第一个参数
then
    git add -A
    git commit -m"$1"
    git pull origin master
    git push origin master
    git status
    echo "完成add、commit、pull和push"
else
    echo "请添加注释再来一遍"
fi

echo '按下 <CTRL+D> 退出'
echo -n '输入你最喜欢的网站名: '
while read FILM
do
  if [ "$FILM" ]
  then
    echo "是的！$FILM 是一个好网站"
  else
    echo "是的！github 是一个好网站"
  fi
done
