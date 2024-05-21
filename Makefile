up-mysql:
	DB_CLIENT=mysql docker-compose -f build/docker-compose.yml -f build/docker-compose.mysql.yml up

up-mysql-debug:
	DB_CLIENT=mysql docker-compose -f build/docker-compose.yml -f build/docker-compose.mysql.yml -f build/docker-compose.mysql.debug.yml up

up-redis:
	DB_CLIENT=redis docker-compose -f build/docker-compose.yml -f build/docker-compose.redis.yml up