GREEN := $(shell tput setaf 2)
RESET := $(shell tput sgr0)
build: 
	go build
rebuildbuild: clean build 
clean: 
	rm -rf myWs
rebuild: clean build
test: build
	@echo "$(GREEN)Running tests...$(RESET)"
	@echo "$(GREEN)test1: With -l flag$(RESET)"
	./myWs -l test/test1.txt
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test2: With -m flag$(RESET)"
	./myWs -m test/test2.txt
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test3: With -w flag$(RESET)"
	./myWs -w test/test1.txt
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test4: With -l many files$(RESET)"
	./myWs -l test/test1.txt test/test2.txt test/test3.txt test/test4.txt test/test5.txt test/test6.txt test/test7.txt test/test8.txt
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test5: With -m many files$(RESET)"
	./myWs -m test/test1.txt test/test2.txt test/test3.txt test/test4.txt test/test5.txt test/test6.txt test/test7.txt test/test8.txt
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)test5: With -w many files$(RESET)"
	./myWs -w test/test1.txt test/test2.txt test/test3.txt test/test4.txt test/test5.txt test/test6.txt test/test7.txt test/test8.txt
	@echo "--------------------------------------------------------------------------"
	@echo "$(GREEN)All tests completed!$(RESET)"