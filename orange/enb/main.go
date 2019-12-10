


func main(){

	// Check connection
        // In order to get uConn
        // give name for GTP device, role(GGSN/SGSN), local/remote net.Addr, restart counter, channel to let background process pass the errors.
        uConn, err := v1.DialUPlaneKernel("gtp0", v1.RoleGGSN, laddr, raddr, 0, errCh)
        if err != nil {
            log.Fatal(err)
        }
	defer uConn.Close() // close user plane connection after return
        log.Printf("Connection established with %s", raddr.String())

	// Add handler ( what to do on receiving message )
	// AddHandler function has two parameters 1. message type 2. hand


}
