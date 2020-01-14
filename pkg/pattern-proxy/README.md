<b>Realisation of Proxy pattern</b>

This implementation has 3 structures: dataBase(main struct), dataBaseNode(proxy struct) and cache.<br>
Cache allows DBNode to receive 'fresh' data without calling the main dataBase

The dataBase and DBNode are connected by one 'DBFunctions' interface

Each of structures has two <i>methods</i>:<br>
SendData() - sending new data from a client <br>
GetData() - returning data to the client

You can test this pattern in main.go
