.PHONY: app sh
sh:
	docker exec -it cloud-track-backend_app_1 sh

.PHONY: db bash
bash:
	docker exec -it cloud-track-backend_db_1 bash