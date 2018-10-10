package runner

// COMPLETED returned by command running successfully
const COMPLETED string = "completed"

//Runner set ups and runs a CameraCommand object
type Runner interface {
	Add(command CameraCommand)
	//Handle return messages from external process
	Handle() error
	Start() error
	Stop()
}
