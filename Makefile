build-all:
	go build ./bin/master
	go build ./bin/master_slu
	go build ./bin/master_slr
	go build ./bin/master_tergum

update-deps:
	@echo "Updating sikalabs/slu..."; \
	REPO="sikalabs/slu"; \
	SLU_COMMIT=$$(slu github latest-commit $$REPO --short); \
	echo "  -> commit $$SLU_COMMIT"; \
	go get github.com/$$REPO@$$SLU_COMMIT; \
	echo "Updating sikalabs/slr..."; \
	REPO="sikalabs/slr"; \
	SLR_COMMIT=$$(slu github latest-commit $$REPO --short); \
	echo "  -> commit $$SLR_COMMIT"; \
	go get github.com/$$REPO@$$SLR_COMMIT; \
	echo "Updating sikalabs/tergum..."; \
	REPO="sikalabs/tergum"; \
	TERGUM_COMMIT=$$(slu github latest-commit $$REPO --short); \
	echo "  -> commit $$TERGUM_COMMIT"; \
	go get github.com/$$REPO@$$TERGUM_COMMIT; \
	echo "Running go mod tidy..."; \
	go mod tidy; \
	if [ -n "$$(git status --porcelain)" ]; then \
		echo "Committing changes..."; \
		git add go.mod go.sum; \
		git commit -m "deps: update slu@$$SLU_COMMIT, slr@$$SLR_COMMIT, tergum@$$TERGUM_COMMIT"; \
		echo "Tagging and pushing..."; \
		slu git tag-next-calver; \
		git push; \
		git push --tags; \
	else \
		echo "No changes to commit."; \
	fi
