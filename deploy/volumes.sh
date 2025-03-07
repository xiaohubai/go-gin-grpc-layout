if [ -d "/usr/local/volumes" ];then
	echo "文件夹: volumes已存在,无需创建"
else
	cp -r volumes /usr/local
	chmod -R 777 /usr/local/volumes
	chmod -R 777 /usr/local/volumes/mysql/data
	chmod -R 777 /usr/local/volumes/redis/data
	echo "success"
fi