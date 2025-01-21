.PHONY: all

up:	# Запуск контейнеров
	docker compose up -d --force-recreate --remove-orphans

down: # Остановка контейнеров
	docker compose down --remove-orphans

build: # Сборка образов и установка зависимостей
	make copy-env
	docker compose up --build -d wallet-app

copy-env:
	cp ./.env.example ./.env

help:	# Справка по командам
	@grep -E '^[a-zA-Z0-9 -]+:.*#'  Makefile | sort | while read -r l; do printf "\033[1;32m$$(echo $$l | cut -f 1 -d':')\033[00m:$$(echo $$l | cut -f 2- -d'#')\n"; done