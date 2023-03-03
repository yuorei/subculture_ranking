build:
	@docker compose build
	@make up

up:
	@docker compose up -d
	@make api

api:
	@docker compose exec api bash

ps:
	@docker compose ps

stop api:
	@docker compose stop api

stop db:
	@docker compose stop db
rm api:
	@docker compose rm api
rm db:
	@docker compose rm db