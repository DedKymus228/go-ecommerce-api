package server

//func Run() {
//	srv := gin.Default()
//
//}
//
//func Shutdown(srv *gin.Engine) {
//	quit := make(chan os.Signal, 1)
//	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
//	<-quit
//
//	ctx, cancel := context.WithTimeout(context.Background(), constants.ShutdownTime)
//	defer cancel()
//
//	srv.Shutdown(ctx)
//}
