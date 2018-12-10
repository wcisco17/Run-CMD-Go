# Run-CMD-Go

CMD ShortHand to run applications in your directory.

Project is written in go:

* Make sure you have go installed, to run the cmd program.
* After running this program kill all ports by running kill -9 $(lsof -i:3000 -t) 2> /dev/null
  (Still in the process of creating a cmd that does it directly)

First enter your own directory path containing all your projects (line 30): 	

   Enter your own source path example: "*", "src/projectName"
   
	`dirname := "../../../Desktop/"` 

Run this command in the root project directory: 

`go run main.go pow`

It will prompt a message like so: 

Enter First Part of The Command To Run: 

`yarn`

Enter Second Part of The Command To Run: 

`start`


Enter Desktop Dir Path: 

`enterProjectName`

then it should run the application directly in your file!

(If not please consult the error log or personally ask on the issue page)

The original Purpose of this project was to create an easy interactive program to quickly run my projects.

Feel Free to update it, fork it, and use it for your own program. 


Command Line shorthand to traverse through directories and run applications directly using Go lang.  
