TMP_COVER_FILE="/tmp/go-cover.tmp"
TMP_COVER_HTML_FILE="/tmp/index.html"

run_tests:
	@echo -e "\n======================="
	@echo -e 	   "Running full test suite"
	@echo -e 	   "=======================\n"
	@go test -cover ./... | sed ''/^ok/s//$$(printf "\033[32mok\033[0m")/'' | sed ''/^?/s//$$(printf "\033[33m-\033[0m")/'' | sed ''/^FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

coverage:
	@echo -e "\n\n\n================================"
	@echo -e 	   "Evaluating coverage of $$(printf "\033[33meach file\033[0m")"
	@echo -e 	   "================================\n"
	@go test -cover ./... | sed ''/^ok/s//$$(printf "\033[32mok\033[0m")/'' | sed ''/^?/s//$$(printf "\033[33m-\033[0m")/'' | sed ''/^FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''

	@echo -e "\n============================================"
	@echo -e   "Evaluating coverage for $$(printf "\033[33mfile as part of repo\033[0m")"
	@echo -e   "============================================\n"
	@go test -coverpkg=./cmd/...,./internal/... -coverprofile=$(TMP_COVER_FILE) ./... | sed ''/^ok/s//$$(printf "\033[32mok\033[0m")/'' | sed ''/^?/s//$$(printf "\033[33m-\033[0m")/'' | sed ''/^FAIL/s//$$(printf "\033[31mFAIL\033[0m")/''
	
	@echo -e "\n====================="
	@echo -e   "$$(printf "\033[33mBreakdown by function\033[0m")"
	@echo -e   "=====================\n"
	@go tool cover -func=$(TMP_COVER_FILE)
	@go tool cover -o $(TMP_COVER_HTML_FILE) -html=$(TMP_COVER_FILE)

	@echo -e "\nFor a full visual breakdown, see $$(printf "\033[32mhttp://localhost:8081\033[0m") for more details"

	@echo -e "\n"
