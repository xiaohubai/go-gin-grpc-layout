if [ -d "/usr/local/volumes" ];then
	echo "文件夹: volumes已存在,无需创建"
else
	cp -r volumes /usr/local
	chmod 777 /usr/local/volumes
	chmod 777 /usr/local/volumes/mysql
	chmod 777 /usr/local/volumes/redis/data
	echo "success"
fi