.PHONY: run.all
run.all: run.logic run.connector run.job

.PHONY: run.logic
run.logic:
	@kratos run app/logic/cmd/logic

.PHONY: run.job
run.job:
	@kratos run app/job/cmd/job

.PHONY: run.connector
run.connector:
	@kratos run app/connector/cmd/connector