.PHONY: wire.all
wire.all: wire.job wire.connector wire.logic

.PHONY: wire.job
wire.job:
	cd app/job/cmd/job && wire

.PHONY: wire.logic
wire.logic:
	cd app/logic/cmd/logic && wire

.PHONY: wire.connector
wire.connector:
	cd app/connector/cmd/connector && wire