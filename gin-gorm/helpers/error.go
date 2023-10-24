package helpers

func PanicError (err error, msg string) {
  if err != nil {
    panic("Error: " + msg)
  }
}
