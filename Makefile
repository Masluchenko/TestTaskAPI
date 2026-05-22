include .env
export

export PROJECT_ROOT=$(shell pwd)

env-up:
	docker compose up -d depart-postgres

env-down:
	docker compose down depart-postgres

env-clen:
	@read -p "Очистить все volume файлы? [y/n]" ans; \
	if ["$$ans" = "y" ]; then \
		docker compose down depart-postgres && \
		rm -rf out/pgdata
		echo "Файлы окружения очищены"
	else \
		echo "Очистка окружения отменена"; \
	fi