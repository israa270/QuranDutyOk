package tms



const(
	TERMINAL_INFO  = "info"
	TERMINAL_LOG_CAT = "log"
	COMMAND_TASK = "task"  //for app | firmware
	COMMAND_UNINSTALL = "uninstall"
	TASK_ID = "taskId"
    COMMAND = "command"
	PACKAGE_NAME ="packageName"
    COMMAND_CONTROL= "control"
    COMMAND_RESTART= "restart"
	COMMAND_MESSAGE ="notification"
    TypeLogCat = "type"
	TERMINAL_MESSAGE = "message"
	MODULE = "module"
	VALUE  = "value"

	// with taskId
	APP_INSTALLED = "app_install"
	// APP_UNINSTALLED = "app_uninstall"
	// DOWNLOAD_FILE = "download_file"
	UPDATE_FIRMWARE = "update_firmware"
)