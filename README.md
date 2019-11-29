# bitcoinstreamgo
This code is an example of using the rgamba websocket library to stream bitcoin transaction events in go. In the source code, there is a limit on the size of the message:

```var msg = make([]byte, 512)```

You will need to change this to be larger if you want to stream a larger message (as in this case). 
