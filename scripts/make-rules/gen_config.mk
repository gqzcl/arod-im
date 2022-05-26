INTERNAL_PROTO_FILES1=$(shell find app/logic/internal -name *.proto)
INTERNAL_PROTO_FILES2=$(shell find app/job/internal -name *.proto)
INTERNAL_PROTO_FILES3=$(shell find app/connector/internal -name *.proto)

.PHONY: config.all
config.all: config.logic config.job config.connector


.PHONY: config.logic
config.logic:
	protoc --proto_path=./app/logic/internal \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:./app/logic/internal \
	       $(INTERNAL_PROTO_FILES1)

.PHONY: config.job
config.job:
	protoc --proto_path=./app/job/internal \
    	   --proto_path=./third_party \
     	   --go_out=paths=source_relative:./app/job/internal \
    	   $(INTERNAL_PROTO_FILES2)

.PHONY: config.connector
config.connector:
	protoc --proto_path=./app/connector/internal \
    	   --proto_path=./third_party \
     	   --go_out=paths=source_relative:./app/connector/internal \
    	   $(INTERNAL_PROTO_FILES3)