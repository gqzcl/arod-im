
TOOLS ?=$(BLOCKER_TOOLS) $(CRITICAL_TOOLS) $(TRIVIAL_TOOLS)

.PHONY: tools.install
tools.install: $(addprefix tools.install., $(TOOLS))

.PHONY: tools.install.%
tools.install.%:
	@echo "======> Installing $*"
	@$(MAKE) install.$*

.PHONY: tools.verify.%
tools.verify.%:
	@if ! which $* &>/dev/null; then $(MAKE) tools.install.$*; fi

.PHONY: install.swagger
install.swagger:
	@go install github.com/go-swagger/go-swagger/cmd/swagger@latest

.PHONY: install.gsemver
install.gsemver:
	@go install github.com/arnaud-deprez/gsemver@latest

## 添加版权
.PHONY: install.addlicense
install.addlicense:
	@go install github.com/marmotedu/addlicense@latest