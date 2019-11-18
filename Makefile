all: analysis unit_test build-artifact integration_test acceptance_test robot_test

analysis:
	golangci-lint run

unit_test:
	go test ./cmd/... && go test ./internal/timesheet/...

build-artifact:
	docker-compose build

integration_test:
	docker-compose up -d
	sleep 30
	docker exec -i my-mariadb mysql --user=root --password=root timesheet < atdd/data/prepare_timesheet.sql
	go test ./internal/repository
	docker-compose down
	
acceptance_test: 
	docker-compose up -d
	sleep 30
	docker exec -i my-mariadb mysql --user=root --password=root timesheet < atdd/data/prepare_timesheet.sql
	sleep 10
	newman run atdd/api/showSummaryTimesheetSuccess.json -e atdd/api/environment/local_environment.json -d atdd/api/data/show_summary_timesheet_success.json
	newman run atdd/api/showTimeSheetByEmployeeIDSuccess.json -e atdd/api/environment/local_environment.json -d atdd/api/data/show_timesheet_by_employee_id_success.json
	newman run atdd/api/updateStatusCheckingTransferSuccess.json -e atdd/api/environment/local_environment.json -d atdd/api/data/update_status_checking_transfer_success.json
	newman run atdd/api/deleteIncomeItemSuccess.json
	newman run atdd/api/showEmployeeDetailsByEmployeeIDSuccess.json
	newman run atdd/api/updateEmployeeDetailsSuccess.json
	newman run atdd/api/loginSuccess.json
	newman run atdd/api/showSummaryInYearSuccess.json
	docker-compose down

robot_test:
	docker-compose up -d
	sleep 30
	docker exec -i my-mariadb mysql --user=root --password=root timesheet < atdd/data/prepare_timesheet.sql
	sleep 10
	robot atdd/ui/summary_timesheet_success.robot
	robot atdd/ui/summary_timesheet_by_id_success.robot
	robot atdd/ui/edit_member_info_success.robot
	docker-compose down

down:
	docker-compose down

seed:
	docker exec -i my-mariadb mysql --user=root --password=root timesheet < atdd/data/prepare_timesheet.sql
