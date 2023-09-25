# sendy
A lightweaght and easy to use file sharing program written in go
# Tutorial
start by finding the file you want to send, then run
`sendy <file_name>`.
This will start an HTTP server hosting the file to your local network. Anyone on your network with the host's IP adress can run 
`curl -v -X POST http://<IP adress>/ > <desired name of file>` to download the file.
